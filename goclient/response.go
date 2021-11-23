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

func (r *response) GetStatus() string {
	return r.status
}
func (r *response) GetStatusCode() int {
	return r.statusCode
}
func (r *response) GetHeaders() http.Header {
	return r.headers
}
func (r *response) GetBodyString() string {
	return string(r.body)
}
func (r *response) GetBodyJson(target interface{}) error {
	return json.Unmarshal(r.body, target)
}
