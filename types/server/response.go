package types

import (
	"github.com/cloud-club/Aviator-service/types"
	"time"
)

// ServerInstanceResponse is same as compute server in Naver Cloud (such as AWS EC2)
type ServerInstanceResponse struct {
	ServerInstanceNo               string
	ServerName                     string
	ServerDescription              string
	CpuCount                       int
	MemorySize                     int64
	PlatformType                   types.CommonCode
	LoginKeyName                   string
	PublicIpInstanceNo             string
	PublicIp                       string
	ServerInstanceStatus           types.CommonCode
	ServerInstanceOperation        types.CommonCode
	ServerInstanceStatusName       string
	CreateDate                     time.Time
	Uptime                         time.Time
	ServerImageProductCode         string
	ServerProductCode              string
	IsProtectServerTermination     bool
	ZoneCode                       string
	RegionCode                     string
	VpcNo                          string
	SubnetNo                       string
	NetworkInterfaceNoList         types.NetworkInterfaceNoList
	InitScriptNo                   string
	ServerInstanceType             types.CommonCode
	BaseBlockStorageDiskType       types.CommonCode
	BaseBlockStorageDiskDetailType types.CommonCode
	PlacementGroupNo               string
	PlacementGroupName             string
	MemberServerImageInstanceNo    string
	BlockDevicePartitionList       []types.BlockDevicePartition
	HypervisorType                 types.CommonCode
	ServerImageNo                  string
	ServerSpecCode                 string
}

type ServerInstanceListResponse struct {
	TotalCount   int
	InstanceList []ServerInstanceResponse
}
