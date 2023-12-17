package common

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

type HttpService struct {
	Client    *http.Client
	Interface HttpInterface
}

type HttpInterface interface {
	Get(url string) (*http.Response, error)
	List(url string) (*http.Response, error)
	Create(url string, payload []byte) (*http.Response, error)
	Update(url string, payload []byte) (*http.Response, error)
	Delete(url string) (*http.Response, error)
}

func NewHttpClient(i HttpInterface) HttpService {
	return HttpService{Client: &http.Client{Timeout: time.Second * 10}, Interface: i}
}

func (h *HttpService) Create(url string, payload []byte) (*http.Response, error) {
	return h.do(http.MethodPost, url, payload)
}

func (h *HttpService) Get(url string) (*http.Response, error) {
	return h.do(http.MethodGet, url, nil)
}

func (h *HttpService) List(url string) (*http.Response, error) {
	if len(url) == 0 {
		return nil, fmt.Errorf("please input url")
	}
	//return nil, nil
	return h.do(http.MethodGet, url, nil)
}

func (h *HttpService) Delete(url string) (*http.Response, error) {
	return h.do(http.MethodDelete, url, nil)
}

func (h *HttpService) Update(url string, payload []byte) (*http.Response, error) {
	return h.do(http.MethodPatch, url, payload)
}

func (h *HttpService) do(method string, url string, payload []byte) (*http.Response, error) {
	var req *http.Request
	var err error
	if method == http.MethodGet {
		if req, err = http.NewRequest(method, url, nil); err != nil {
			return nil, err
		}
	} else {
		if req, err = http.NewRequest(method, url, bytes.NewReader(payload)); err != nil {
			return nil, err
		}
	}

	return h.Client.Do(req)
}
