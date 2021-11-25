package goclient

import (
	"net/http"
	"time"
)

type builder struct {
	headers            http.Header
	maxIdleConnections int
	responseTimeout    time.Duration
	connectionTimeout  time.Duration
	disableTimeout     bool
}

type BuilderInterface interface {
	SetHeaders(headers http.Header) BuilderInterface
	SetMaxIdleConnections(timeout int) BuilderInterface
	SetResponseTimeout(timeout time.Duration) BuilderInterface
	SetConnectionTimeout(timeout time.Duration) BuilderInterface
	DisableTimeout(disable bool) BuilderInterface
	BuildClient() ClientInterface
}

func NewBuilder() BuilderInterface {
	return &builder{}
}

func (b *builder) BuildClient() ClientInterface {
	return &Client{
		client:  &http.Client{},
		builder: b,
	}
}

func (b *builder) GetHeaders() http.Header {
	return b.headers
}
func (b *builder) SetHeaders(headers http.Header) BuilderInterface {
	b.headers = headers
	return b
}
func (b *builder) GetMaxIdleConnections() int {
	return b.maxIdleConnections
}
func (b *builder) SetMaxIdleConnections(timeout int) BuilderInterface {
	b.maxIdleConnections = timeout
	return b
}
func (b *builder) GetResponseTimeout() time.Duration {
	return b.responseTimeout
}
func (b *builder) SetResponseTimeout(timeout time.Duration) BuilderInterface {
	b.responseTimeout = timeout
	return b
}
func (b *builder) GetConnectionTimeout() time.Duration {
	return b.connectionTimeout
}
func (b *builder) SetConnectionTimeout(timeout time.Duration) BuilderInterface {
	b.connectionTimeout = timeout
	return b
}
func (b *builder) GetDisableTimeout() bool {
	return b.disableTimeout
}
func (b *builder) DisableTimeout(disable bool) BuilderInterface {
	b.disableTimeout = disable
	return b
}
