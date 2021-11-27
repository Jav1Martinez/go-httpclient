# Go Http-Client
___

Simple HTTP client library for Go.


## What's included
___
A simple client library with:

- Really easy http-client for use.
- Sensible operation naming.
- All the power and flexibility of Go!
- Mock server included to create unit tests without using real environments.

## Getting started
___

## Installation

```bash
# Go Modules
require github.com/Jav1Martinez/go-httpclient
```

## Usage
In order to use the library for making HTTP calls you need to import the library and set your client first:

```go
import "github.com/Jav1Martinez/go-httpclient/goclient"

var Client = goclient.NewBuilder().BuildClient()
```

## Client calls
The ``Client`` interface provides convenient methods that you can use to perform different HTTP calls.

**Important:** It is not necessary to read & close the response since the client handle everything for you. You just need to get the response and start using it!

### Example

```go
// Use client to call any of the following methods.
// Client.method

Get(url string, headers http.Header) (*response, error)
Post(url string, headers http.Header, body interface{}) (*response, error)
Put(url string, headers http.Header, body interface{}) (*response, error)
Patch(url string, headers http.Header, body interface{}) (*response, error)
Delete(url string, headers http.Header) (*response, error)
```

## Mock server usage

The library provides a convenient class for mocking requests and getting a particular response.

### Mock server example:
```go
// Prepare the body response
r := ioutil.NopCloser(bytes.NewReader([]byte(``)))

// Create mock response
HttpClient = &Mock{
    DoFn: func(*http.Request) (*http.Response, error) {
        return &http.Response{
            StatusCode: http.StatusOK,
            Body:       r,
        }, nil
    },
}

// Do method will return the mocked response.
res, _ := HttpClient.Do(&http.Request{})
```

### Attributions

My thanks to [Federico Leon](https://github.com/federicoleon) for his examples and courses to create my first http-client library.