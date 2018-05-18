// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package securityGwService is a auto generated package.
Input file: app.proto
*/
package securityGwService

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"

	"github.com/GeertJohan/go.rice"
	"github.com/go-openapi/analysis"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"

	"github.com/pensando/grpc-gateway/runtime"

	"github.com/pensando/sw/venice/utils/log"
)

var muxMutex sync.Mutex
var mux *runtime.ServeMux

const codecSize = 1024 * 1024

func registerSwaggerDef(m *http.ServeMux, logger log.Logger) error {
	box, err := rice.FindBox("../../../../api/protos/../generated/security/swagger")
	if err != nil {
		err = errors.Wrap(err, "error opening rice.Box")
		return err
	}
	var docs []*spec.Swagger
	walkFn := func(path string, info os.FileInfo, err error) error {
		var e error
		var f []byte
		if !info.IsDir() {
			f, e = box.Bytes(info.Name())
			if e == nil {
				jrm := json.RawMessage(string(f))
				sdoc, e := loads.Analyzed(jrm, "")
				if e == nil {
					docs = append(docs, sdoc.Spec())
				}
			}
		}
		return e
	}
	err = box.Walk("", walkFn)
	if err != nil {
		err = errors.Wrap(err, "walk of files failed")
		return err
	}
	if len(docs) > 1 {
		_ = analysis.Mixin(docs[0], docs[1:]...)
		analysis.FixEmptyResponseDescriptions(docs[0])
	}
	content, err := json.MarshalIndent(docs[0], "", "  ")
	if err != nil {
		err = errors.Wrap(err, "error marshalling swagger file")
		return err
	}
	m.HandleFunc("/swagger/security/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(content)
	})
	return nil
}
