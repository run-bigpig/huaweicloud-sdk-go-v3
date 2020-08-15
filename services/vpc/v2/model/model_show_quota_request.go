/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type ShowQuotaRequest struct {
	Type ShowQuotaRequestType `json:"type,omitempty"`
}

func (o ShowQuotaRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowQuotaRequest", string(data)}, " ")
}

type ShowQuotaRequestType struct {
	value string
}

type ShowQuotaRequestTypeEnum struct {
	VPC                       ShowQuotaRequestType
	SUBNET                    ShowQuotaRequestType
	SECURITY_GROUP            ShowQuotaRequestType
	SECURITY_GROUP_RULE       ShowQuotaRequestType
	PUBLIC_IP                 ShowQuotaRequestType
	VPN                       ShowQuotaRequestType
	VPNGW                     ShowQuotaRequestType
	VPC_PEER                  ShowQuotaRequestType
	FIREWALL                  ShowQuotaRequestType
	SHARE_BANDWIDTH           ShowQuotaRequestType
	SHARE_BANDWIDTH_IP        ShowQuotaRequestType
	LOADBALANCER              ShowQuotaRequestType
	LISTENER                  ShowQuotaRequestType
	PHYSICAL_CONNECT          ShowQuotaRequestType
	VIRTUAL_INTERFACE         ShowQuotaRequestType
	VPC_CONTAIN_ROUTETABLE    ShowQuotaRequestType
	ROUTETABLE_CONTAIN_ROUTES ShowQuotaRequestType
}

func GetShowQuotaRequestTypeEnum() ShowQuotaRequestTypeEnum {
	return ShowQuotaRequestTypeEnum{
		VPC: ShowQuotaRequestType{
			value: "vpc",
		},
		SUBNET: ShowQuotaRequestType{
			value: "subnet",
		},
		SECURITY_GROUP: ShowQuotaRequestType{
			value: "securityGroup",
		},
		SECURITY_GROUP_RULE: ShowQuotaRequestType{
			value: "securityGroupRule",
		},
		PUBLIC_IP: ShowQuotaRequestType{
			value: "publicIp",
		},
		VPN: ShowQuotaRequestType{
			value: "vpn",
		},
		VPNGW: ShowQuotaRequestType{
			value: "vpngw",
		},
		VPC_PEER: ShowQuotaRequestType{
			value: "vpcPeer",
		},
		FIREWALL: ShowQuotaRequestType{
			value: "firewall",
		},
		SHARE_BANDWIDTH: ShowQuotaRequestType{
			value: "shareBandwidth",
		},
		SHARE_BANDWIDTH_IP: ShowQuotaRequestType{
			value: "shareBandwidthIP",
		},
		LOADBALANCER: ShowQuotaRequestType{
			value: "loadbalancer",
		},
		LISTENER: ShowQuotaRequestType{
			value: "listener",
		},
		PHYSICAL_CONNECT: ShowQuotaRequestType{
			value: "physicalConnect",
		},
		VIRTUAL_INTERFACE: ShowQuotaRequestType{
			value: "virtualInterface",
		},
		VPC_CONTAIN_ROUTETABLE: ShowQuotaRequestType{
			value: "vpcContainRoutetable",
		},
		ROUTETABLE_CONTAIN_ROUTES: ShowQuotaRequestType{
			value: "routetableContainRoutes",
		},
	}
}

func (c ShowQuotaRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowQuotaRequestType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}