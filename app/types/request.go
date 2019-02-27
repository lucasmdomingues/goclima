package types

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

type Request struct {
	ResponseWriter http.ResponseWriter
	Route          string
	Method         string
	Values         []byte
}

func (r *Request) SendRequest() ([]byte, error) {

	r.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	r.ResponseWriter.Header().Add("Access-Control-Allow-Credentials", "true")

	var values io.Reader
	if r.Values == nil {
		values = nil
	} else {
		values = bytes.NewBuffer(r.Values)
	}

	req, err := http.NewRequest(r.Method, r.Route, values)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func NewRequest(w http.ResponseWriter, route, method string, values []byte) *Request {
	return &Request{
		ResponseWriter: w,
		Route:          route,
		Method:         "GET",
		Values:         nil,
	}
}
