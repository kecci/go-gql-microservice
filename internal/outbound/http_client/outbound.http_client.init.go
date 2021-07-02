package http_client

import "net/http"

// HttpClientOutbound structure repository for special site
type HttpClientOutbound struct {
	httpClientOutbound *http.Client
}

// NewHttpClientOutbound initialize repository special site
func NewHttpClientOutbound() *HttpClientOutbound {
	httpClientOutbound := &http.Client{}
	return &HttpClientOutbound{
		httpClientOutbound: httpClientOutbound,
	}
}
