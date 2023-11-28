package pkg

import (
	"errors"
	"fmt"
	"net/http"
)

type ServerService struct {
	token string
}

//go:generate mockgen -destination=mocks/mock_server.go -package=mocks github.com/cloud-club/Aviator-service/pkg ServerInterface
type ServerInterface interface {
	Get(url string) error
	List(url string) error
	Create(url string, payload interface{}) error
	Update(url string) error
	Delete(url string) error
}

func NewServerService(token string) ServerInterface {
	return &ServerService{token: token}
}

func (server *ServerService) GetToken() string {
	return server.token
}

func (server *ServerService) Create(url string, payload interface{}) error {
	return nil
}

func (server *ServerService) Get(url string) error {
	return nil
}

func (server *ServerService) List(url string) error {
	if len(url) == 0 {
		return errors.New("please input url")
	}
	return nil
}

func (server *ServerService) Delete(regionCode, serverInstanceNo, url string) error {
	// apiURL 정의
	url += fmt.Sprintf("?regionCode=%s&serverInstanceNoList.1=%d&responseFormatType=json",
	 regionCode, serverInstanceNo,
	)
	
	// httpRequest 생성
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("Error creating request: ", err)
	}
	// httpRequest header 설정
	GetCommonHeader(request)
	
	// httpRequest token 설정
	SetAuthToken(request, server.token)

	// HTTP 클라이언트 생성
	client := &http.Client{}

	// 요청 보내기
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("Error sending request:", err)
	}
	defer response.Body.Close()
	
	// 결과 반환
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected response status code:", err)
	}
	
	return nil
}

func (server *ServerService) Update(url string) error {
	return nil
}
