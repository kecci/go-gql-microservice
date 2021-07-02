package http_client

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
)

// Send func to send http REST
func (h *HttpClientOutbound) Send(ctx context.Context, req *http.Request) (resByte []byte, resHttpCode int, err error) {
	res, err := h.httpClientOutbound.Do(req)
	if err != nil {
		log.Println("[HTTPREST] error: ", err)
		return
	}

	defer res.Body.Close()
	resByte, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("[HTTPREST] error read response: ", err)
		return
	}
	resHttpCode = res.StatusCode
	return
}

// CreateRequest to create new http request
func (h *HttpClientOutbound) CreateRequest(method string, url string, body []byte) (httpRequest *http.Request, err error) {
	httpRequest, err = http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	return
}
