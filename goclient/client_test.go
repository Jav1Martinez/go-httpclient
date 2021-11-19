package goclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetMaxIdleConnections(t *testing.T) {
	newBuilder := &builder{}
	client := httpClient{
		builder: newBuilder,
	}

	t.Run("When maxIdleConnection value Is higher Than 0", func(t *testing.T) {
		client.builder.maxIdleConnections = 1

		maxIdleConnection := client.getMaxIdleConnections()

		assert.Greater(t, maxIdleConnection, 0, "maxIdleConnections have to be higher than 0")
	})

	t.Run("When maxIdleConnection value Is equal Than defaultMaxIdleConnections value", func(t *testing.T) {
		client.builder.maxIdleConnections = 0

		maxIdleConnection := client.getMaxIdleConnections()

		assert.EqualValues(t, defaultMaxIdleConnections, maxIdleConnection, "maxIdleConnections have to be equal than defaultMaxIdleConnections value")
	})
}

func TestGetResponseTimeout(t *testing.T) {
	newBuilder := &builder{}
	client := httpClient{
		builder: newBuilder,
	}

	t.Run("When responseTimeout value Is higher Than 0", func(t *testing.T) {
		client.builder.responseTimeout = 1 * time.Second

		responseTimeout := client.getResponseTimeout()

		assert.Greater(t, responseTimeout, 0*time.Second, "responseTimeout have to be higher than 0")
	})

	t.Run("When responseTimeout value Is disable", func(t *testing.T) {
		client.builder.responseTimeout = 0 * time.Second
		client.builder.DisableTimeout(true)

		responseTimeout := client.getResponseTimeout()

		assert.Equal(t, 0*time.Second, responseTimeout, "responseTimeout have to be equal than 0")
	})

	t.Run("When responseTimeout value Is enable and Is equal Than 0", func(t *testing.T) {
		client.builder.responseTimeout = 0 * time.Second
		client.builder.DisableTimeout(false)

		responseTimeout := client.getResponseTimeout()

		assert.Equal(t, defaultResponseTimeout, responseTimeout, "responseTimeout have to be equal than defaultResponseTimeout value")
	})
}

func TestGetConnectionTimeout(t *testing.T) {
	newBuilder := &builder{}
	client := httpClient{
		builder: newBuilder,
	}

	t.Run("When connectionTimeout value Is higher Than 0", func(t *testing.T) {
		client.builder.connectionTimeout = 1 * time.Second

		connectionTimeout := client.getConnectionTimeout()

		assert.Greater(t, connectionTimeout, 0*time.Second, "connectionTimeout have to be higher than 0")
	})

	t.Run("When connectionTimeout Is disable", func(t *testing.T) {
		client.builder.connectionTimeout = 0 * time.Second
		client.builder.DisableTimeout(true)

		connectionTimeout := client.getConnectionTimeout()

		assert.Equal(t, 0*time.Second, connectionTimeout, "connectionTimeout have to be disable")
	})

	t.Run("When connectionTimeout Is enable and Is equal Than 0", func(t *testing.T) {
		client.builder.connectionTimeout = 0 * time.Second
		client.builder.DisableTimeout(false)

		connectionTimeout := client.getConnectionTimeout()

		assert.Equal(t, defaultConnectionTimeout, connectionTimeout, "connectionTimeout have to be equal than defaultConnectionTimeout value ")
	})
}
