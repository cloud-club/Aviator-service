package pkg

import (
	"bytes"
	"errors"
	"net/http"
)

type ServerService struct {
	client *http.Client
	token  string
}

//go:generate mockgen -destination=mocks/mock_server.go -package=mocks github.com/cloud-club/Aviator-service/pkg ServerInterface
type ServerInterface interface {
	Get(url string) (*http.Response, error)
	List(url string) (*http.Response, error)
	Create(url string, payload []byte) (*http.Response, error)
	Update(url string, payload []byte) (*http.Response, error)
	Delete(url string) (*http.Response, error)
}

func NewServerService(token string) ServerInterface {
	return &ServerService{token: token}
}

func (server *ServerService) GetToken() string {
	return server.token
}

func (server *ServerService) Create(url string, payload []byte) (*http.Response, error) {
	return server.do(http.MethodPost, url, payload)
}

func (server *ServerService) Get(url string) (*http.Response, error) {
	return server.do(http.MethodGet, url, nil)
}

func (server *ServerService) List(url string) (*http.Response, error) {
	if len(url) == 0 {
		return nil, errors.New("please input url")
	}
	return nil, nil
	//return server.do(http.MethodGet, url, nil)
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
	GetCommonHeader(req)
	SetAuthToken(req, server.token)

	return server.client.Do(req)
}
