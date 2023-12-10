package network

import (
	"github.com/cloud-club/Aviator-service/pkg/common"
	"github.com/cloud-club/Aviator-service/types/network"
	"net/http"
)

type Network struct {
	HttpService common.HttpService
	Interface   NetServiceInterface
}

type NetServiceInterface interface {
	GetNetworkInterfaceList(dto *network.RequestBodyNetworkInterfaceList) (*http.Response, error)
}

func Init(i NetServiceInterface) Network {
	return Network{Interface: i}
}

func (n Network) GetNetworkInterfaceList(dto *network.RequestBodyNetworkInterfaceList) (*http.Response, error) {
	return n.getNetworkInterfaceList(dto)
}
