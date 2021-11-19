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
	BuildClient() HttpClientInterface
	BuildClientMock() HttpClientInterface
}

func NewBuilder() BuilderInterface {
	return &builder{}
}

func (b *builder) BuildClient() HttpClientInterface {
	return &httpClient{
		client:  &http.Client{},
		builder: b,
	}
}

func (b *builder) BuildClientMock() HttpClientInterface {
	return &httpClientMock{
		builder: b,
	}
}

func (b *builder) SetHeaders(headers http.Header) BuilderInterface {
	b.headers = headers
	return b
}
func (b *builder) SetMaxIdleConnections(timeout int) BuilderInterface {
	b.maxIdleConnections = timeout
	return b
}
func (b *builder) SetResponseTimeout(timeout time.Duration) BuilderInterface {
	b.responseTimeout = timeout
	return b
}
func (b *builder) SetConnectionTimeout(timeout time.Duration) BuilderInterface {
	b.connectionTimeout = timeout
	return b
}
func (b *builder) DisableTimeout(disable bool) BuilderInterface {
	b.disableTimeout = disable
	return b
}
