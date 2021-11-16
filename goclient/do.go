package goclient

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*response, error) {

	request, err := c.newRequest(method, url, headers, body)
	if err != nil {
		return nil, errors.New("unable to do the request")
	}

	requestResponse, err := c.getHttpClient().Do(request)
	if err != nil {
		return nil, err
	}

	defer requestResponse.Body.Close()
	requestResponseBody, err := ioutil.ReadAll(requestResponse.Body)
	if err != nil {
		return nil, err
	}

	return &response{
		status:     requestResponse.Status,
		statusCode: requestResponse.StatusCode,
		headers:    requestResponse.Header,
		body:       requestResponseBody,
	}, nil
}

func (c *httpClient) getHttpClient() *http.Client {

	c.clientOnce.Do(func() {
		c.client = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})

	return c.client
}
