package goclient

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMock(t *testing.T) {

	r := ioutil.NopCloser(bytes.NewReader([]byte(``)))

	HttpClient = &Mock{
		DoFn: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       r,
			}, nil
		},
	}

	res, _ := HttpClient.Do(&http.Request{})

	assert.Equal(t, http.StatusOK, res.StatusCode, "StatusCode 200 expected")
}
