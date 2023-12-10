package network_test

import (
	"github.com/cloud-club/Aviator-service/types/network"
	"log"
)

func (suite *NcpSuite) TestGetNetworkInterfaceList() {
	resp, err := suite.ncp.Network.GetNetworkInterfaceList(&network.RequestBodyNetworkInterfaceList{})
	defer resp.Body.Close()
	if err != nil {
		suite.T().Fatalf("%v", err)
	}
	log.Println(resp.Status)
	log.Println(resp.Body)
}
