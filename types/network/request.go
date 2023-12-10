package network

type RequestBodyNetworkInterfaceList struct {
	RegionCode                 string   `json:"RegionCode"`
	SubnetName                 string   `json:"SubnetName"`
	NetworkInterfaceNoList     []string `json:"NetworkInterfaceNoList"`
	NetworkInterfaceName       string   `json:"NetworkInterfaceName"`
	NetworkInterfaceStatusCode string   `json:"NetworkInterfaceStatusCode"`
	SecondaryIpList            []string `json:"SecondaryIpList"`
	Ip                         string   `json:"Ip"`
	InstanceNo                 string   `json:"InstanceNo"`
	IsDefault                  bool     `json:"IsDefault"`
	DeviceName                 string   `json:"DeviceName"`
	ServerName                 string   `json:"ServerName"`
}
