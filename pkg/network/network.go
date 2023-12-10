package network

import (
	"fmt"
	"github.com/cloud-club/Aviator-service/pkg/common"
	"github.com/cloud-club/Aviator-service/types/network"
	"net/http"
	"strconv"
)

func (n Network) getNetworkInterfaceList(dto *network.RequestBodyNetworkInterfaceList) (*http.Response, error) {
	url := fmt.Sprintf("%s/getNetworkInterfaceList?regionCode=%s&subnetName=%s,&networkInterfaceNoList.1=%s&networkInterfaceName=%s&networkInterfaceStatusCode=%s&ip=%s&secondaryIpList.1=%s&instanceNo%s&isDefault=%s&deviceName=%s&serverName=%s",
		common.BaseUrl,
		dto.RegionCode, dto.SubnetName, dto.NetworkInterfaceNoList, dto.NetworkInterfaceName, dto.NetworkInterfaceStatusCode,
		dto.Ip, dto.SecondaryIpList, dto.InstanceNo, strconv.FormatBool(dto.IsDefault), dto.DeviceName, dto.ServerName)

	return n.HttpService.Get(url)
}
