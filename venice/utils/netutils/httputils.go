package netutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"bytes"
	"net/url"

	"github.com/pensando/sw/venice/utils/log"
)

// RestAPIFunc defines a rest handler
type RestAPIFunc func(r *http.Request) (interface{}, error)

// MakeHTTPHandler is a utility that wraps a rest handler
func MakeHTTPHandler(handlerFunc RestAPIFunc) http.HandlerFunc {
	// Create a closure and return an anonymous function
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the handler
		resp, err := handlerFunc(r)
		if err != nil {
			// Log error
			log.Errorf("Handler for %s %s returned error: %s", r.Method, r.URL, err)

			if resp == nil {
				// Send HTTP response
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				// Send HTTP response as Json
				content, err1 := json.Marshal(resp)
				if err1 != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.WriteHeader(http.StatusInternalServerError)
				w.Write(content)
			}
		} else {
			// Send HTTP response as Json
			content, err := json.Marshal(resp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(content)
		}
	}
}

// HTTPPost performs http POST operation
func HTTPPost(url string, req interface{}, resp interface{}) error {
	// Convert the req to json
	jsonStr, err := json.Marshal(req)
	if err != nil {
		log.Errorf("Error converting request data(%#v) to Json. Err: %v", req, err)
		return err
	}

	// Perform HTTP POST operation
	res, err := http.Post(url, "application/json", strings.NewReader(string(jsonStr)))
	if err != nil {
		log.Errorf("Error during http POST. Err: %v", err)
		return err
	}

	// Check the response code
	if res.StatusCode == http.StatusInternalServerError {
		eBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.New("HTTP StatusInternalServerError" + err.Error())
		}
		res.Body.Close()
		return errors.New(string(eBody))
	}

	if res.StatusCode != http.StatusOK {
		log.Errorf("HTTP error response. Status: %s, StatusCode: %d", res.Status, res.StatusCode)
		return fmt.Errorf("HTTP error response. Status: %s, StatusCode: %d", res.Status, res.StatusCode)
	}

	// Read the entire response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error during ioutil readall. Err: %v", err)
		return err
	}
	defer res.Body.Close()

	if resp == nil {
		return nil
	}

	// Convert response json to struct
	err = json.Unmarshal(body, resp)
	if err != nil {
		log.Errorf("Error during json unmarshall. Err: %v", err)
		return err
	}

	return nil
}

// HTTPGetRaw gets the raw data from an http request, i.e. it basically execute curl
func HTTPGetRaw(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer r.Body.Close()

	log.Debugf("HTTP Get: %s", url)
	if r.StatusCode != 200 {
		return []byte{}, errors.New(r.Status)
	}

	response, err := ioutil.ReadAll(r.Body)

	return response, err
}

// HTTPGet fetches json object from an http get request
func HTTPGet(url string, jdata interface{}) error {
	response, err := HTTPGetRaw(url)
	if err != nil {
		return err
	}

	return json.Unmarshal(response, jdata)
}

// HTTPPut provides wrapper for http PUT operations.
func HTTPPut(url string, req interface{}, resp interface{}) error {
	client := &http.Client{}
	// Convert the req to json
	jsonStr, err := json.Marshal(req)
	if err != nil {
		log.Errorf("Error converting request data(%#v) to Json. Err: %v", req, err)
		return err
	}

	request, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(jsonStr)))
	if err != nil {
		log.Errorf("Error during http PUT. Err: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	res, err := client.Do(request)

	// Perform HTTP POST operation
	//res, err := http.Post(url, "application/json", strings.NewReader(string(jsonStr)))
	if err != nil {
		log.Errorf("Error during http PUT. Err: %v", err)
		return err
	}

	// Check the response code
	if res.StatusCode == http.StatusInternalServerError {
		eBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.New("HTTP StatusInternalServerError" + err.Error())
		}
		res.Body.Close()
		return errors.New(string(eBody))
	}

	if res.StatusCode != http.StatusOK {
		log.Errorf("HTTP error response. Status: %s, StatusCode: %d", res.Status, res.StatusCode)
		return errors.New("HTTP Error response")
	}

	if resp == nil {
		return nil
	}

	// Read the entire response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error during ioutil readall. Err: %v", err)
		return err
	}
	defer res.Body.Close()

	// Convert response json to struct
	err = json.Unmarshal(body, resp)
	if err != nil {
		log.Errorf("Error during json unmarshall. Err: %v", err)
		return err
	}

	return nil
}

// HTTPDelete provides wrapper for http DELETE operations.
func HTTPDelete(url string, req interface{}, resp interface{}) error {
	client := &http.Client{}
	jsonStr := []byte{}
	var err error
	if req != nil {
		// Convert the req to json
		jsonStr, err = json.Marshal(req)
		if err != nil {
			log.Errorf("Error converting request data(%#v) to Json. Err: %v", req, err)
			return err
		}
	}

	request, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(string(jsonStr)))
	if err != nil {
		log.Errorf("Error during http DELETE. Err: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	res, err := client.Do(request)

	// Perform HTTP POST operation
	//res, err := http.Post(url, "application/json", strings.NewReader(string(jsonStr)))
	if err != nil {
		log.Errorf("Error during http DELETE. Err: %v", err)
		return err
	}

	// Check the response code
	if res.StatusCode == http.StatusInternalServerError {
		eBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.New("HTTP StatusInternalServerError" + err.Error())
		}
		res.Body.Close()
		return errors.New(string(eBody))
	}

	if res.StatusCode != http.StatusOK {
		log.Errorf("HTTP error response. Status: %s, StatusCode: %d", res.Status, res.StatusCode)
		return errors.New("HTTP Error response")
	}

	// Read the entire response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error during ioutil readall. Err: %v", err)
		return err
	}
	defer res.Body.Close()

	if resp == nil {
		return nil
	}

	// Convert response json to struct
	err = json.Unmarshal(body, resp)
	if err != nil {
		log.Errorf("Error during json unmarshall. Err: %v", err)
		return err
	}

	return nil
}

func encodeHTTPRequest(req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

// CreateHTTPRequest creates a http request
func CreateHTTPRequest(instance string, in interface{}, method, path string) (*http.Request, error) {
	target, err := url.Parse(instance)
	if err != nil {
		return nil, fmt.Errorf("invalid instance %s", instance)
	}
	target.Path = path
	req, err := http.NewRequest(method, target.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request (%s)", err)
	}
	if err = encodeHTTPRequest(req, in); err != nil {
		return nil, fmt.Errorf("could not encode request (%s)", err)
	}
	return req, nil
}

// CopyHTTPDefaultTransport creates a http.Transport from http.DefaultTransport
func CopyHTTPDefaultTransport() *http.Transport {
	ht := &http.Transport{}
	dt := http.DefaultTransport.(*http.Transport)
	ht.Proxy = dt.Proxy
	ht.DialContext = dt.DialContext
	ht.MaxIdleConns = dt.MaxIdleConns
	ht.IdleConnTimeout = dt.IdleConnTimeout
	ht.TLSHandshakeTimeout = dt.TLSHandshakeTimeout
	ht.ExpectContinueTimeout = dt.ExpectContinueTimeout
	ht.MaxIdleConnsPerHost = dt.MaxIdleConnsPerHost
	return ht
}

// StatusWriter is used to record http status code
type StatusWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader records http status code before delegating call to http.ResponseWriter
func (s *StatusWriter) WriteHeader(statusCode int) {
	s.statusCode = statusCode
	s.ResponseWriter.WriteHeader(statusCode)
}

// GetStatusCode returns recorded http status code. If it is not set it returns an error.
func (s *StatusWriter) GetStatusCode() (int, error) {
	if s.statusCode == 0 {
		return s.statusCode, errors.New("status code unknown")
	}
	return s.statusCode, nil
}

// NewStatusWriter returns an instance of StatusWriter to record returned http status code
func NewStatusWriter(w http.ResponseWriter) *StatusWriter {
	return &StatusWriter{ResponseWriter: w}
}
