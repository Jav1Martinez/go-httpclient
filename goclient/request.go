package goclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func (c *client) newRequest(method string, url string, headers http.Header, body interface{}) (*http.Request, error) {

	fullHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(body, fullHeaders.Get("Content-Type"))
	if err != nil {
		return nil, errors.New("unable to create the request body")
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	request.Header = fullHeaders

	return request, err
}

func (c *client) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// Add common headers
	for key, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(key, value[0])
		}
	}
	// Add custom headers
	for key, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(key, value[0])
		}
	}

	return result
}

func (c *client) getRequestBody(body interface{}, contentType string) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {

	case "application/json":
		return json.Marshal(body)

	default:
		return json.Marshal(body)
	}
}
