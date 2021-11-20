package goclient

import (
	"net/http"
)

type httpClientMock struct {
	getFn    func(url string, headers http.Header) (*response, error)
	postFn   func(url string, headers http.Header, body interface{}) (*response, error)
	putFn    func(url string, headers http.Header, body interface{}) (*response, error)
	patchFn  func(url string, headers http.Header, body interface{}) (*response, error)
	deleteFn func(url string, headers http.Header) (*response, error)
}

func (hcm *httpClientMock) Get(url string, headers http.Header) (*response, error) {
	return hcm.getFn(url, headers)
}

func (hcm *httpClientMock) Post(url string, headers http.Header, body interface{}) (*response, error) {
	return hcm.postFn(url, headers, body)
}

func (hcm *httpClientMock) Put(url string, headers http.Header, body interface{}) (*response, error) {
	return hcm.putFn(url, headers, body)
}

func (hcm *httpClientMock) Patch(url string, headers http.Header, body interface{}) (*response, error) {
	return hcm.patchFn(url, headers, body)
}

func (hcm *httpClientMock) Delete(url string, headers http.Header) (*response, error) {
	return hcm.deleteFn(url, headers)
}
