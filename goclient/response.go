package goclient

import (
	"encoding/json"
	"net/http"
)

type response struct {
	status     string
	statusCode int
	headers    http.Header
	body       []byte
}

func (r *response) Status() string {
	return r.status
}
func (r *response) StatusCode() int {
	return r.statusCode
}
func (r *response) Headers() http.Header {
	return r.headers
}
func (r *response) BodyString() string {
	return string(r.body)
}
func (r *response) BodyJson(target interface{}) error {
	return json.Unmarshal(r.body, target)
}
