package network

import "github.com/cloud-club/Aviator-service/types"

type AccessControlGroupNoList []string
type SecondaryIpList []string

type ResponseBodyNetworkInterface struct {
	NetworkInterfaceNo          string
	NetworkInterfaceName        string
	SubnetNo                    string
	DeleteOnTermination         bool
	IsDefault                   bool
	DeviceName                  string
	NetworkInterfaceStatus      types.CommonCode
	InstanceType                types.CommonCode
	InstanceNo                  string
	IP                          string
	MacAddress                  string
	AccessControlGroupNoList    AccessControlGroupNoList
	NetworkInterfaceDescription string
	SecondaryIpList             SecondaryIpList
}
