/*
 * Cbr
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
type ListPoliciesResponse struct {
	//
	Policies []Policy `json:"policies,omitempty"`
}

func (o ListPoliciesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPoliciesResponse", string(data)}, " ")
}