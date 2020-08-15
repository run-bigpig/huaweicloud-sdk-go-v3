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

// Request Object
type KeystoneListProjectsRequest struct {
	Name     string `json:"name,omitempty"`
	ParentId string `json:"parent_id,omitempty"`
	Enabled  bool   `json:"enabled,omitempty"`
	IsDomain bool   `json:"is_domain,omitempty"`
	Page     int32  `json:"page,omitempty"`
	PerPage  int32  `json:"per_page,omitempty"`
}

func (o KeystoneListProjectsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneListProjectsRequest", string(data)}, " ")
}