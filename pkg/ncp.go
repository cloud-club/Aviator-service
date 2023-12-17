package pkg

import (
	"github.com/cloud-club/Aviator-service/pkg/common"
	"github.com/cloud-club/Aviator-service/pkg/network"
)

type NcpService struct {
	Server      ServerService
	HttpService common.HttpService
	Network     network.Network
}

func New() *NcpService {
	ncp := newNcpService()
	ncp.Network = network.Init(ncp.Network.Interface)
	ncp.Server = Init(ncp.Server.Interface)
	ncp.Network.HttpService = common.NewHttpClient(ncp.HttpService.Interface)
	ncp.Server.HttpService = common.NewHttpClient(ncp.HttpService.Interface)

	return ncp
}

func newNcpService() *NcpService {
	return &NcpService{}
}
