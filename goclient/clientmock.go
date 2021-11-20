package goclient

import (
	"net/http"
)

type httpClientMock struct {
	GetFn    func(url string, headers http.Header) (*response, error)
	PostFn   func(url string, headers http.Header, body interface{}) (*response, error)
	PutFn    func(url string, headers http.Header, body interface{}) (*response, error)
	PatchFn  func(url string, headers http.Header, body interface{}) (*response, error)
	DeleteFn func(url string, headers http.Header) (*response, error)
}

func (hcm *httpClientMock) Get(url string, headers http.Header) (*response, error) {
	return hcm.GetFn(url, headers)
}

func (hcm *httpClientMock) Post(url string, headers http.Header, body interface{}) (*response, error) {
	return hcm.PostFn(url, headers, body)
}

func (hcm *httpClientMock) Put(url string, headers http.Header, body interface{}) (*response, error) {
	return hcm.PutFn(url, headers, body)
}

func (hcm *httpClientMock) Patch(url string, headers http.Header, body interface{}) (*response, error) {
	return hcm.PatchFn(url, headers, body)
}

func (hcm *httpClientMock) Delete(url string, headers http.Header) (*response, error) {
	return hcm.DeleteFn(url, headers)
}
