package goclient

import (
	"net/http"
	"sync"
	"time"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
	defaultConnectionTimeout  = 1 * time.Second
)

type httpClient struct {
	client     *http.Client
	builder    *builder
	clientOnce sync.Once
}

type HttpClientInterface interface {
	Get(url string, headers http.Header) (*response, error)
	Post(url string, headers http.Header, body interface{}) (*response, error)
	Put(url string, headers http.Header, body interface{}) (*response, error)
	Patch(url string, headers http.Header, body interface{}) (*response, error)
	Delete(url string, headers http.Header) (*response, error)
}

func (c *httpClient) Get(url string, headers http.Header) (*response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}
func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*response, error) {
	return c.do(http.MethodPost, url, headers, body)
}
func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*response, error) {
	return c.do(http.MethodPut, url, headers, body)
}
func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}
func (c *httpClient) Delete(url string, headers http.Header) (*response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}
	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	if c.builder.disableTimeout {
		return 0
	}
	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}
	if c.builder.disableTimeout {
		return 0
	}
	return defaultConnectionTimeout
}
