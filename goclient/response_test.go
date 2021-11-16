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

	status := response.Status()

	assert.EqualValues(t, response.status, status, "Expected to have same string value")
}

func TestStatusCode(t *testing.T) {
	response := response{
		statusCode: 200,
	}

	statusCode := response.StatusCode()

	assert.EqualValues(t, response.statusCode, statusCode, "Expected to have same int value")
}

func TestHeaders(t *testing.T) {
	response := response{
		headers: make(http.Header),
	}

	headers := response.Headers()

	assert.EqualValues(t, response.headers, headers, "Expected to have same string value")
}

func TestBodyString(t *testing.T) {
	response := response{
		body: []byte("Test text"),
	}

	bodyString := response.BodyString()

	assert.EqualValues(t, string(response.body), bodyString, "Expected to have same string value")
}

func TestBodyJson(t *testing.T) {
	response := response{
		body: []byte("Test text json"),
	}

	bodyJson := response.BodyJson(response)

	assert.EqualValues(t, json.Unmarshal(response.body, &response), bodyJson, "Expected to have same string value")
}
