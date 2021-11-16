package goclient

import (
	"net/http"
)

type HttpClientInterface interface {
	Do(request *http.Request) (*http.Response, error)
}

var (
	HttpClient HttpClientInterface
)

func init() {
	HttpClient = &http.Client{}
}
