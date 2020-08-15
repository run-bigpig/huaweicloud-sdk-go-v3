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
type KeystoneShowSecurityComplianceResponse struct {
	Config *Config `json:"config,omitempty"`
}

func (o KeystoneShowSecurityComplianceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneShowSecurityComplianceResponse", string(data)}, " ")
}