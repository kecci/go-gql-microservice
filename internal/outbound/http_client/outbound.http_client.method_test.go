package http_client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestHttpClientDoRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	httpClientMock := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "http://example.com/some/path")
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	s := &HttpClientOutbound{
		httpClientOutbound: httpClientMock,
	}
	mockURL, _ := url.Parse("http://example.com/some/path")

	t.Run("WHEN function success", func(t *testing.T) {
		body, _, err := s.Send(context.Background(), &http.Request{Method: "GET", URL: mockURL})
		assert.NoError(t, err)
		assert.Equal(t, []byte("OK"), body)
	})

	t.Run("WHEN failed do request", func(t *testing.T) {
		body, _, _ := s.Send(context.Background(), &http.Request{Method: "GET"})
		assert.NotEqual(t, []byte("OK"), body)
	})
}

func TestHttpClientCreateRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	httpClientMock := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "http://example.com/some/path")
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	s := &HttpClientOutbound{
		httpClientOutbound: httpClientMock,
	}

	t.Run("WHEN function success", func(t *testing.T) {
		_, err := s.CreateRequest("GET", "http://example.com/some/path", nil)
		assert.NoError(t, err)
	})
}
