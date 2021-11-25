package goclient

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestNewBuilder(t *testing.T) {
	builderInterface := &builder{}

	builder := NewBuilder()

	assert.IsType(t, builderInterface, builder, "It was expected a builderInterface match")
}

func TestBuildClient(t *testing.T) {
	builder := &builder{}
	httpClientInterface := &Client{}

	httpClient := builder.BuildClient()

	assert.IsType(t, httpClientInterface, httpClient, "It was expected a httpClientInterface match")
}

func TestSetHeaders(t *testing.T) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	builder := &builder{}

	builder.SetHeaders(headers)

	assert.EqualValues(t, len(headers), len(builder.GetHeaders()), "It was expected 1 header")
}

func TestSetMaxIdleConnections(t *testing.T) {
	maxIdleConnections := 2
	builder := &builder{}

	builder.SetMaxIdleConnections(maxIdleConnections)

	assert.EqualValues(t, maxIdleConnections, builder.GetMaxIdleConnections(), "Expected to have same number value")
}

func TestSetResponseTimeout(t *testing.T) {
	responseTimeout := 2 * time.Second
	builder := &builder{}

	builder.SetResponseTimeout(responseTimeout)

	assert.EqualValues(t, responseTimeout, builder.GetResponseTimeout(), "Expected to have same time value")
}

func TestSetConnectionTimeout(t *testing.T) {
	connectionTimeout := 4 * time.Second
	builder := &builder{}

	builder.SetConnectionTimeout(connectionTimeout)

	assert.EqualValues(t, connectionTimeout, builder.GetConnectionTimeout(), "Expected to have same time value")
}

func TestDisableTimeout(t *testing.T) {
	DisableTimeout := true
	builder := &builder{}

	builder.DisableTimeout(DisableTimeout)

	assert.EqualValues(t, DisableTimeout, builder.GetDisableTimeout(), "Expected to have same bool value")
}
