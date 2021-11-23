package goclient

import (
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
	defaultConnectionTimeout  = 1 * time.Second
)

type client struct {
	client     *http.Client
	builder    *builder
	clientOnce sync.Once
}

type ClientInterface interface {
	Get(url string, headers http.Header) (*response, error)
	Post(url string, headers http.Header, body interface{}) (*response, error)
	Put(url string, headers http.Header, body interface{}) (*response, error)
	Patch(url string, headers http.Header, body interface{}) (*response, error)
	Delete(url string, headers http.Header) (*response, error)
}

func (c *client) Get(url string, headers http.Header) (*response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}
func (c *client) Post(url string, headers http.Header, body interface{}) (*response, error) {
	return c.do(http.MethodPost, url, headers, body)
}
func (c *client) Put(url string, headers http.Header, body interface{}) (*response, error) {
	return c.do(http.MethodPut, url, headers, body)
}
func (c *client) Patch(url string, headers http.Header, body interface{}) (*response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}
func (c *client) Delete(url string, headers http.Header) (*response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}

func (c *client) getMaxIdleConnections() int {
	if c.builder.GetMaxIdleConnections() > 0 {
		return c.builder.GetMaxIdleConnections()
	}
	return defaultMaxIdleConnections
}

func (c *client) getResponseTimeout() time.Duration {
	if c.builder.GetResponseTimeout() > 0 {
		return c.builder.GetResponseTimeout()
	}
	if c.builder.GetDisableTimeout() {
		return 0
	}
	return defaultResponseTimeout
}

func (c *client) getConnectionTimeout() time.Duration {
	if c.builder.GetConnectionTimeout() > 0 {
		return c.builder.GetConnectionTimeout()
	}
	if c.builder.GetDisableTimeout() {
		return 0
	}
	return defaultConnectionTimeout
}

func (c *client) do(method string, url string, headers http.Header, body interface{}) (*response, error) {

	request, err := c.newRequest(method, url, headers, body)
	if err != nil {
		return nil, errors.New("unable to do the request")
	}

	httpResponse, err := HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close()
	httpResponseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	return &response{
		status:     httpResponse.Status,
		statusCode: httpResponse.StatusCode,
		headers:    httpResponse.Header,
		body:       httpResponseBody,
	}, nil
}
