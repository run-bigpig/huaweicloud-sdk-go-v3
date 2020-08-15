/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type KeystoneListServicesResponse struct {
	// 服务信息列表。
	Services []Service `json:"services,omitempty"`
	Links    *Links    `json:"links,omitempty"`
}

func (o KeystoneListServicesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneListServicesResponse", string(data)}, " ")
}