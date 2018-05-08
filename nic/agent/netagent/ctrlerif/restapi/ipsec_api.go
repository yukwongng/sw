// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package netproto is a auto generated package.
Input file: ipsec.proto
*/
package restapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pensando/sw/nic/agent/httputils"
	"github.com/pensando/sw/venice/ctrler/npm/rpcserver/netproto"
)

// addIPSecPolicyAPIRoutes adds IPSecPolicy routes
func addIPSecPolicyAPIRoutes(r *mux.Router, srv *RestServer) {

	r.Methods("GET").Subrouter().HandleFunc("/", httputils.MakeHTTPHandler(srv.listIPSecPolicyHandler))

	r.Methods("POST").Subrouter().HandleFunc("/", httputils.MakeHTTPHandler(srv.postIPSecPolicyHandler))

	r.Methods("PUT").Subrouter().HandleFunc("/{ObjectMeta.Tenant}/{ObjectMeta.NameSpace}/{ObjectMeta.Name}", httputils.MakeHTTPHandler(srv.putIPSecPolicyHandler))

	r.Methods("DELETE").Subrouter().HandleFunc("/{ObjectMeta.Tenant}/{ObjectMeta.NameSpace}/{ObjectMeta.Name}", httputils.MakeHTTPHandler(srv.deleteIPSecPolicyHandler))

}

func (s *RestServer) listIPSecPolicyHandler(r *http.Request) (interface{}, error) {
	return s.agent.ListIPSecPolicy(), nil
}

func (s *RestServer) postIPSecPolicyHandler(r *http.Request) (interface{}, error) {
	var o netproto.IPSecPolicy
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	return nil, s.agent.CreateIPSecPolicy(&o)

}

func (s *RestServer) putIPSecPolicyHandler(r *http.Request) (interface{}, error) {
	var o netproto.IPSecPolicy
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	return nil, s.agent.UpdateIPSecPolicy(&o)

}

func (s *RestServer) deleteIPSecPolicyHandler(r *http.Request) (interface{}, error) {
	var o netproto.IPSecPolicy
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	return nil, s.agent.DeleteIPSecPolicy(&o)

}

// addIPSecSADecryptAPIRoutes adds IPSecSADecrypt routes
func addIPSecSADecryptAPIRoutes(r *mux.Router, srv *RestServer) {

	r.Methods("GET").Subrouter().HandleFunc("/", httputils.MakeHTTPHandler(srv.listIPSecSADecryptHandler))

	r.Methods("POST").Subrouter().HandleFunc("/", httputils.MakeHTTPHandler(srv.postIPSecSADecryptHandler))

	r.Methods("PUT").Subrouter().HandleFunc("/{ObjectMeta.Tenant}/{ObjectMeta.Namespace}/{ObjectMeta.Name}", httputils.MakeHTTPHandler(srv.putIPSecSADecryptHandler))

	r.Methods("DELETE").Subrouter().HandleFunc("/{ObjectMeta.Tenant}/{ObjectMeta.Namespace}/{ObjectMeta.Name}", httputils.MakeHTTPHandler(srv.deleteIPSecSADecryptHandler))

}

func (s *RestServer) listIPSecSADecryptHandler(r *http.Request) (interface{}, error) {
	return s.agent.ListIPSecSADecrypt(), nil
}

func (s *RestServer) postIPSecSADecryptHandler(r *http.Request) (interface{}, error) {
	var o netproto.IPSecSADecrypt
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	return nil, s.agent.CreateIPSecSADecrypt(&o)

}

func (s *RestServer) putIPSecSADecryptHandler(r *http.Request) (interface{}, error) {
	var o netproto.IPSecSADecrypt
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	return nil, s.agent.UpdateIPSecSADecrypt(&o)

}

func (s *RestServer) deleteIPSecSADecryptHandler(r *http.Request) (interface{}, error) {
	var o netproto.IPSecSADecrypt
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	return nil, s.agent.DeleteIPSecSADecrypt(&o)

}

// addIPSecSAEncryptAPIRoutes adds IPSecSAEncrypt routes
func addIPSecSAEncryptAPIRoutes(r *mux.Router, srv *RestServer) {

	r.Methods("GET").Subrouter().HandleFunc("/", httputils.MakeHTTPHandler(srv.listIPSecSAEncryptHandler))

	r.Methods("POST").Subrouter().HandleFunc("/", httputils.MakeHTTPHandler(srv.postIPSecSAEncryptHandler))

	r.Methods("PUT").Subrouter().HandleFunc("/{ObjectMeta.Tenant}/{ObjectMeta.NameSpace}/{ObjectMeta.Name}", httputils.MakeHTTPHandler(srv.putIPSecSAEncryptHandler))

	r.Methods("DELETE").Subrouter().HandleFunc("/{ObjectMeta.Tenant}/{ObjectMeta.NameSpace}/{ObjectMeta.Name}", httputils.MakeHTTPHandler(srv.deleteIPSecSAEncryptHandler))

}

func (s *RestServer) listIPSecSAEncryptHandler(r *http.Request) (interface{}, error) {
	return s.agent.ListIPSecSAEncrypt(), nil
}

func (s *RestServer) postIPSecSAEncryptHandler(r *http.Request) (interface{}, error) {
	var o netproto.IPSecSAEncrypt
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	return nil, s.agent.CreateIPSecSAEncrypt(&o)

}

func (s *RestServer) putIPSecSAEncryptHandler(r *http.Request) (interface{}, error) {
	var o netproto.IPSecSAEncrypt
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	return nil, s.agent.UpdateIPSecSAEncrypt(&o)

}

func (s *RestServer) deleteIPSecSAEncryptHandler(r *http.Request) (interface{}, error) {
	var o netproto.IPSecSAEncrypt
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	return nil, s.agent.DeleteIPSecSAEncrypt(&o)

}
