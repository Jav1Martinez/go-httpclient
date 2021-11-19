package goclient

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewRequest(t *testing.T) {
	newBuilder := &builder{}
	client := httpClient{
		builder: newBuilder,
	}

	t.Run("new request created", func(t *testing.T) {
		method := "GET"
		url := "www.test.com"
		headers := http.Header{}
		requestBody := []string{"with", "json"}
		httpRequest := &http.Request{}

		request, err := client.newRequest(method, url, headers, requestBody)

		assert.IsType(t, httpRequest, request, "It was expected a httpRequest match")
		assert.Equal(t, nil, err, "No error expected")
	})
}

func TestGetRequestHeaders(t *testing.T) {
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "go-httpclient")
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Test", "TEST-VALUE")
	client := httpClient{}
	client.builder = &builder{}
	client.builder.headers = commonHeaders

	fullHeaders := client.getRequestHeaders(requestHeaders)

	assert.EqualValues(t, 3, len(fullHeaders), "It was expected 3 header")
}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}

	t.Run("noBodyNilResponse", func(t *testing.T) {
		body, err := client.getRequestBody(nil, "")

		assert.Equal(t, nil, err, "No error expected")
		assert.Empty(t, body, "No body expected")
	})

	t.Run("bodyWithJson", func(t *testing.T) {
		requestBody := []string{"with", "json"}

		body, err := client.getRequestBody(requestBody, "application/json")

		assert.Equal(t, nil, err, "No error expected after Json marshall")
		assert.Equal(t, `["with","json"]`, string(body), "Invalid json body obtained")
	})

	t.Run("bodyWithJsonAsDefault", func(t *testing.T) {
		requestBody := []string{"with", "json", "default"}

		body, err := client.getRequestBody(requestBody, "")

		assert.Equal(t, nil, err, "No error expected after Json marshall")
		assert.Equal(t, `["with","json","default"]`, string(body), "Invalid json body obtained")
	})
}
