package goclient

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestStatus(t *testing.T) {
	response := response{
		status: "ok",
	}

	status := response.GetStatus()

	assert.EqualValues(t, response.status, status, "Expected to have same string value")
}

func TestStatusCode(t *testing.T) {
	response := response{
		statusCode: 200,
	}

	statusCode := response.GetStatusCode()

	assert.EqualValues(t, response.statusCode, statusCode, "Expected to have same int value")
}

func TestHeaders(t *testing.T) {
	response := response{
		headers: make(http.Header),
	}

	headers := response.GetHeaders()

	assert.EqualValues(t, response.headers, headers, "Expected to have same string value")
}

func TestBodyString(t *testing.T) {
	response := response{
		body: []byte("Test text"),
	}

	bodyString := response.GetBodyString()

	assert.EqualValues(t, string(response.body), bodyString, "Expected to have same string value")
}

func TestBodyJson(t *testing.T) {
	response := response{
		body: []byte("Test text json"),
	}

	bodyJson := response.GetBodyJson(response)

	assert.EqualValues(t, json.Unmarshal(response.body, &response), bodyJson, "Expected to have same string value")
}
