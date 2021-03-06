package runtime

import (
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"reflect"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/pensando/grpc-gateway/runtime/internal"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// FowardResponseStreamToWebSocket forwards stream response to a websocket
func FowardResponseStreamToWebSocket(ctx context.Context, marshaler Marshaler, w http.ResponseWriter, req *http.Request, conn *websocket.Conn, recv func() (proto.Message, error), opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	// The HTTP 1.1 response has already been sent by the time we get here. Handle the data stream here.
	for {
		resp, err := recv()
		if err == io.EOF {
			cmsg := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Backend connection terminated")
			conn.WriteControl(websocket.CloseMessage, cmsg, time.Time{})
			conn.Close()
			return
		}
		if err != nil {
			cmsg := websocket.FormatCloseMessage(websocket.CloseInternalServerErr, fmt.Sprintf("reading from backed had error (%s)", err))
			conn.WriteControl(websocket.CloseMessage, cmsg, time.Time{})
			conn.Close()
			return
		}
		err = conn.WriteJSON(resp)
		if err != nil {
			cmsg := websocket.FormatCloseMessage(websocket.CloseProtocolError, fmt.Sprintf("writing message failed with error (%s)", err))
			conn.WriteControl(websocket.CloseMessage, cmsg, time.Time{})
			conn.Close()
			return
		}
	}
}

// ForwardBinaryResponseStream forwards the response as a binary stream
func ForwardBinaryResponseStream(ctx context.Context, marshaler Marshaler, w http.ResponseWriter, req *http.Request, recv func() (proto.Message, error), opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	md, ok := ServerMetadataFromContext(ctx)
	if !ok {
		grpclog.Printf("Failed to extract ServerMetadata from context")
		http.Error(w, "unexpected error", http.StatusInternalServerError)
		return
	}
	handleForwardResponseServerMetadata(w, md)

	w.Header().Set("Content-Disposition", "attachment")
	w.Header().Set("Content-Type", "application/octet-stream; charset=utf-8")
	if err := handleForwardResponseOptions(ctx, w, nil, opts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	for {
		resp, err := recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			handleForwardResponseStreamError(marshaler, w, err)
			return
		}
		if err := handleForwardResponseOptions(ctx, w, resp, opts); err != nil {
			handleForwardResponseStreamError(marshaler, w, err)
			return
		}
		v := reflect.ValueOf(resp)
		if v.Kind() == reflect.Ptr {
			v = reflect.Indirect(v)
		}
		fv := v.FieldByName("Content").Interface()
		if _, err = w.Write(fv.([]byte)); err != nil {
			grpclog.Printf("Failed to send response: %v", err)
			return
		}
	}

}

// ForwardResponseStream forwards the stream from gRPC server to REST client.
func ForwardResponseStream(ctx context.Context, marshaler Marshaler, w http.ResponseWriter, req *http.Request, recv func() (proto.Message, error), opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	f, ok := w.(http.Flusher)
	if !ok {
		grpclog.Printf("Flush not supported in %T", w)
		http.Error(w, "unexpected type of web server", http.StatusInternalServerError)
		return
	}

	md, ok := ServerMetadataFromContext(ctx)
	if !ok {
		grpclog.Printf("Failed to extract ServerMetadata from context")
		http.Error(w, "unexpected error", http.StatusInternalServerError)
		return
	}
	handleForwardResponseServerMetadata(w, md)

	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", marshaler.ContentType())
	if err := handleForwardResponseOptions(ctx, w, nil, opts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	f.Flush()
	for {
		resp, err := recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			handleForwardResponseStreamError(marshaler, w, err)
			return
		}
		if err := handleForwardResponseOptions(ctx, w, resp, opts); err != nil {
			handleForwardResponseStreamError(marshaler, w, err)
			return
		}

		buf, err := marshaler.Marshal(streamChunk(resp, nil))
		if err != nil {
			grpclog.Printf("Failed to marshal response chunk: %v", err)
			return
		}
		if _, err = fmt.Fprintf(w, "%s\n", buf); err != nil {
			grpclog.Printf("Failed to send response chunk: %v", err)
			return
		}
		f.Flush()
	}
}

func handleForwardResponseServerMetadata(w http.ResponseWriter, md ServerMetadata) {
	for k, vs := range md.HeaderMD {
		hKey := fmt.Sprintf("%s%s", MetadataHeaderPrefix, k)
		for i := range vs {
			w.Header().Add(hKey, vs[i])
		}
	}
}

func handleForwardResponseTrailerHeader(w http.ResponseWriter, md ServerMetadata) {
	for k := range md.TrailerMD {
		tKey := textproto.CanonicalMIMEHeaderKey(fmt.Sprintf("%s%s", MetadataTrailerPrefix, k))
		w.Header().Add("Trailer", tKey)
	}
}

func handleForwardResponseTrailer(w http.ResponseWriter, md ServerMetadata) {
	for k, vs := range md.TrailerMD {
		tKey := fmt.Sprintf("%s%s", MetadataTrailerPrefix, k)
		for i := range vs {
			w.Header().Add(tKey, vs[i])
		}
	}
}

// ForwardResponseMessage forwards the message "resp" from gRPC server to REST client.
func ForwardResponseMessage(ctx context.Context, marshaler Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	md, ok := ServerMetadataFromContext(ctx)
	if !ok {
		grpclog.Printf("Failed to extract ServerMetadata from context")
	}

	handleForwardResponseServerMetadata(w, md)
	handleForwardResponseTrailerHeader(w, md)
	w.Header().Set("Content-Type", marshaler.ContentType())
	if err := handleForwardResponseOptions(ctx, w, resp, opts); err != nil {
		HTTPError(ctx, marshaler, w, req, err)
		return
	}

	buf, err := marshaler.Marshal(resp)
	if err != nil {
		grpclog.Printf("Marshal error: %v", err)
		HTTPError(ctx, marshaler, w, req, err)
		return
	}

	if _, err = w.Write(buf); err != nil {
		grpclog.Printf("Failed to write response: %v", err)
	}

	handleForwardResponseTrailer(w, md)
}

func handleForwardResponseOptions(ctx context.Context, w http.ResponseWriter, resp proto.Message, opts []func(context.Context, http.ResponseWriter, proto.Message) error) error {
	if len(opts) == 0 {
		return nil
	}
	for _, opt := range opts {
		if err := opt(ctx, w, resp); err != nil {
			grpclog.Printf("Error handling ForwardResponseOptions: %v", err)
			return err
		}
	}
	return nil
}

func handleForwardResponseStreamError(marshaler Marshaler, w http.ResponseWriter, err error) {
	buf, merr := marshaler.Marshal(streamChunk(nil, err))
	if merr != nil {
		grpclog.Printf("Failed to marshal an error: %v", merr)
		return
	}
	if _, werr := fmt.Fprintf(w, "%s\n", buf); werr != nil {
		grpclog.Printf("Failed to notify error to client: %v", werr)
		return
	}
}

func streamChunk(result proto.Message, err error) map[string]proto.Message {
	if err != nil {
		grpcCode := grpc.Code(err)
		httpCode := HTTPStatusFromCode(grpcCode)
		return map[string]proto.Message{
			"error": &internal.StreamError{
				GrpcCode:   int32(grpcCode),
				HttpCode:   int32(httpCode),
				Message:    err.Error(),
				HttpStatus: http.StatusText(httpCode),
			},
		}
	}
	if result == nil {
		return streamChunk(nil, fmt.Errorf("empty response"))
	}
	return map[string]proto.Message{"result": result}
}
