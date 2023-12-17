package pkg

import (
	"bytes"
	"fmt"
	"github.com/cloud-club/Aviator-service/pkg/common"
	"net/http"
)

type ServerService struct {
	Interface   ServerInterface
	HttpService common.HttpService
	token       string
}

type ServerInterface interface {
	Get(url string) (*http.Response, error)
	List(url string) (*http.Response, error)
	Create(url string, payload []byte) (*http.Response, error)
	Update(url string, payload []byte) (*http.Response, error)
	Delete(url string) (*http.Response, error)
}

func Init(i ServerInterface) ServerService {
	return ServerService{Interface: i}
}

func (server *ServerService) Create(url string, payload []byte) (*http.Response, error) {
	return server.do(http.MethodPost, url, payload)
}

func (server *ServerService) Get(url string) (*http.Response, error) {
	return server.do(http.MethodGet, url, nil)
}

func (server *ServerService) List(url string) (*http.Response, error) {
	if len(url) == 0 {
		return nil, fmt.Errorf("please input url")
	}
	//return nil, nil
	return server.do(http.MethodGet, url, nil)
}

func (server *ServerService) Delete(url string) (*http.Response, error) {
	return server.do(http.MethodDelete, url, nil)
}

func (server *ServerService) Update(url string, payload []byte) (*http.Response, error) {
	return server.do(http.MethodPatch, url, payload)
}

func (server *ServerService) do(method string, url string, payload []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	common.GetCommonHeader(req)

	return server.HttpService.Client.Do(req)
}
