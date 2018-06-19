// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package netproto is a auto generated package.
Input file: network.proto
*/
package restapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/gorilla/mux"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/httputils"
	agentTypes "github.com/pensando/sw/nic/agent/netagent/state/types"
	"github.com/pensando/sw/venice/ctrler/npm/rpcserver/netproto"
)

// addNetworkAPIRoutes adds Network routes
func addNetworkAPIRoutes(r *mux.Router, srv *RestServer) {

	r.Methods("GET").Subrouter().HandleFunc("/", httputils.MakeHTTPHandler(srv.listNetworkHandler))

	r.Methods("POST").Subrouter().HandleFunc("/", httputils.MakeHTTPHandler(srv.postNetworkHandler))

	r.Methods("PUT").Subrouter().HandleFunc("/{ObjectMeta.Tenant}/{ObjectMeta.Namespace}/{ObjectMeta.Name}", httputils.MakeHTTPHandler(srv.putNetworkHandler))

	r.Methods("DELETE").Subrouter().HandleFunc("/{ObjectMeta.Tenant}/{ObjectMeta.Namespace}/{ObjectMeta.Name}", httputils.MakeHTTPHandler(srv.deleteNetworkHandler))

}

func (s *RestServer) listNetworkHandler(r *http.Request) (interface{}, error) {
	return s.agent.ListNetwork(), nil
}

func (s *RestServer) postNetworkHandler(r *http.Request) (interface{}, error) {
	var res Response
	var o netproto.Network
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	c, _ := types.TimestampProto(time.Now())
	o.CreationTime = api.Timestamp{
		Timestamp: *c,
	}
	o.ModTime = api.Timestamp{
		Timestamp: *c,
	}

	err = s.agent.CreateNetwork(&o)

	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()

		return res, err

	}

	res.References = []string{fmt.Sprintf("%s%s/%s/%s", r.RequestURI, o.Tenant, o.Namespace, o.Name)}

	res.StatusCode = http.StatusOK
	return res, err
}

func (s *RestServer) putNetworkHandler(r *http.Request) (interface{}, error) {
	var res Response
	var o netproto.Network
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	m, _ := types.TimestampProto(time.Now())
	o.ModTime = api.Timestamp{
		Timestamp: *m,
	}
	err = s.agent.UpdateNetwork(&o)

	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()

		return res, err

	}

	res.References = []string{r.RequestURI}

	res.StatusCode = http.StatusOK
	return res, err
}

func (s *RestServer) deleteNetworkHandler(r *http.Request) (interface{}, error) {
	var res Response
	var o netproto.Network
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	err = s.agent.DeleteNetwork(&o)

	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()

		// check if its a cannot delete type err
		delErr, ok := err.(*agentTypes.ErrCannotDelete)
		if ok {
			res.References = delErr.References
			return res, err
		}

	}

	res.References = []string{r.RequestURI}

	res.StatusCode = http.StatusOK
	return res, err
}
