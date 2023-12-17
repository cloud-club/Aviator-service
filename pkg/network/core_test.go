package network_test

import (
	"encoding/json"
	"github.com/cloud-club/Aviator-service/types/network"
	"io"
	"log"
)

func (suite *NcpSuite) TestGetNetworkInterfaceList() {

	resp, err := suite.ncp.Network.GetNetworkInterfaceList(&network.RequestBodyNetworkInterfaceList{
		RegionCode: "KR",
	})
	defer resp.Body.Close()
	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		suite.T().Fatalf("%v", err)
	}
	log.Println(resp.Status)
	data := new(ResponseOfGetNetworkInterfaceList)
	if err = json.Unmarshal(rawData, data); err != nil {
		suite.T().Fatalf("%v", err)
	}
	log.Println(data.GetNetworkInterfaceListResponse)

	if err != nil {
		suite.T().Fatalf("%v", err)
	}
}
