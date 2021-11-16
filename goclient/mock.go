package goclient

import (
	"net/http"
)

type MockDoFn func(request *http.Request) (*http.Response, error)

type Mock struct {
	DoFn MockDoFn
}

func (m *Mock) Do(request *http.Request) (*http.Response, error) {
	return m.DoFn(request)
}
