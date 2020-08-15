/*
 * EPS
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
type MigrateResourceRequest struct {
	EnterpriseProjectId string           `json:"enterprise_project_id"`
	Body                *MigrateResource `json:"body,omitempty"`
}

func (o MigrateResourceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"MigrateResourceRequest", string(data)}, " ")
}