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
type CreatePolicyResponse struct {
	Policy *Policy `json:"policy,omitempty"`
}

func (o CreatePolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePolicyResponse", string(data)}, " ")
}