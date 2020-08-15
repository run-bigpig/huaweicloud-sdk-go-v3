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
type KeystoneShowPermissionResponse struct {
	Role *RoleResult `json:"role,omitempty"`
}

func (o KeystoneShowPermissionResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneShowPermissionResponse", string(data)}, " ")
}