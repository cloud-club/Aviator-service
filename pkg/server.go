package pkg

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/cloud-club/Aviator-service/types/auth"
	types "github.com/cloud-club/Aviator-service/types/server"
)

var errorNotStopped = errors.New("Server is not stopped before update or deletion")

type ServerService struct {
	KeyService *auth.KeyService
}

func NewServerService(keyService *auth.KeyService) *ServerService {
	return &ServerService{KeyService: keyService}
}

type ServerInterface interface {
	List(url string, request *types.ListServerRequest) (*types.ListServerResponse, error)
	Create(url string, request *types.CreateServerRequest, params []int) (*types.CreateServerResponse, error)
	Update(url string, request *types.UpdateServerRequest) (*types.UpdateServerResponse, error)
	Start(url string, request *types.StartServerRequest) (*types.StartServerResponse, error)
	Stop(url string, request *types.StopServerRequest) (*types.StopServerResponse, error)
	Delete(url string, request *types.DeleteServerRequest) (*types.DeleteServerResponse, error)
}

func checkStatus(server *ServerService, condition string, repeat int) bool {
	for i := 0; i < repeat; i++ {
		lsr := &types.ListServerRequest{RegionCode: "KR"}
		resp, _ := server.List(API_URL+GET_SERVER_INSTANCE_PATH, lsr)

		serverStatus := resp.ServerInstanceList[0].ServerInstanceStatus.Code
		if serverStatus == condition {
			return true
		}
		time.Sleep(time.Second)
	}

	return false
}

func (server *ServerService) List(url string, request *types.ListServerRequest) (*types.ListServerResponse, error) {
	requestParams := types.RequestString(request)

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, url+requestParams, nil)
	if err != nil {
		return nil, err
	}
	// Set HTTP header for NCP authorization
	SetNCPHeader(req, server.KeyService.GetAccessKey(), server.KeyService.GetSecretKey())

	// Make the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		// Read the response body and show the body message in error.
		responseByteData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", responseByteData)
	}

	responseByteData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var csr *types.ListServerResponse
	responseInterface, err := types.MapResponse(responseByteData, &csr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responseStruct := responseInterface.(**types.ListServerResponse)

	return *responseStruct, err
}

func (server *ServerService) Create(url string, request *types.CreateServerRequest, params []int) (*types.CreateServerResponse, error) {
	// Set url with query parameters
	networkInterfaceIndex := params[0]
	acgNoList := params[1]
	requestParams := types.CreateRequestString(request, networkInterfaceIndex, acgNoList)

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, url+requestParams, nil)
	if err != nil {
		return nil, err
	}
	// Set HTTP header for NCP authorization
	SetNCPHeader(req, server.KeyService.GetAccessKey(), server.KeyService.GetSecretKey())

	// Make the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		// Read the response body and show the body message in error.
		responseByteData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", responseByteData)
	}

	responseByteData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var csr *types.CreateServerResponse
	responseInterface, err := types.MapResponse(responseByteData, &csr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// interface{} 타입으로 변환된 responseInterface를 다시 CreateServerResponse 타입으로 변환
	responseStruct := responseInterface.(**types.CreateServerResponse)

	return *responseStruct, err
}

func (server *ServerService) Update(url string, request *types.UpdateServerRequest) (*types.UpdateServerResponse, error) {
	var usr *types.UpdateServerResponse
	if serverStatus := checkStatus(server, "NSTOP", 25); !serverStatus {
		return usr, errorNotStopped
	}
	requestParams := types.RequestString(request)

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, url+requestParams, nil)
	if err != nil {
		return nil, err
	}
	// Set HTTP header for NCP authorization
	SetNCPHeader(req, server.KeyService.GetAccessKey(), server.KeyService.GetSecretKey())

	// Make the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		// Read the response body and show the body message in error.
		responseByteData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", responseByteData)
	}

	responseByteData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responseInterface, err := types.MapResponse(responseByteData, &usr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responseStruct := responseInterface.(**types.UpdateServerResponse)

	return *responseStruct, err
}

func (server *ServerService) Start(url string, request *types.StartServerRequest) (*types.StartServerResponse, error) {
	requestParams := types.RequestString(request)

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, url+requestParams, nil)
	if err != nil {
		return nil, err
	}
	// Set HTTP header for NCP authorization
	SetNCPHeader(req, server.KeyService.GetAccessKey(), server.KeyService.GetSecretKey())

	// Make the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		// Read the response body and show the body message in error.
		responseByteData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", responseByteData)
	}

	responseByteData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var ssr *types.StartServerResponse
	responseInterface, err := types.MapResponse(responseByteData, &ssr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responseStruct := responseInterface.(**types.StartServerResponse)

	return *responseStruct, err
}

func (server *ServerService) Stop(url string, request *types.StopServerRequest) (*types.StopServerResponse, error) {
	requestParams := types.RequestString(request)

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, url+requestParams, nil)
	if err != nil {
		return nil, err
	}
	// Set HTTP header for NCP authorization
	SetNCPHeader(req, server.KeyService.GetAccessKey(), server.KeyService.GetSecretKey())

	// Make the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		// Read the response body and show the body message in error.
		responseByteData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", responseByteData)
	}

	responseByteData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var csr *types.StopServerResponse
	responseInterface, err := types.MapResponse(responseByteData, &csr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responseStruct := responseInterface.(**types.StopServerResponse)

	return *responseStruct, err
}

func (server *ServerService) Delete(url string, request *types.DeleteServerRequest) (*types.DeleteServerResponse, error) {
	var dsr *types.DeleteServerResponse
	if serverStatus := checkStatus(server, "NSTOP", 20); !serverStatus {
		return dsr, errorNotStopped
	}
	requestParams := types.RequestString(request)

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, url+requestParams, nil)
	if err != nil {
		return nil, err
	}
	// Set HTTP header for NCP authorization
	SetNCPHeader(req, server.KeyService.GetAccessKey(), server.KeyService.GetSecretKey())

	// Make the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		// Read the response body and show the body message in error.
		responseByteData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", responseByteData)
	}

	responseByteData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var csr *types.DeleteServerResponse
	responseInterface, err := types.MapResponse(responseByteData, &csr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responseStruct := responseInterface.(**types.DeleteServerResponse)

	return *responseStruct, err
}
