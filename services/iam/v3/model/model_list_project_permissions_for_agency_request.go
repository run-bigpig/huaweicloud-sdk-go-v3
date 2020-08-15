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
type ListProjectPermissionsForAgencyRequest struct {
	ProjectId string `json:"project_id"`
	AgencyId  string `json:"agency_id"`
}

func (o ListProjectPermissionsForAgencyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectPermissionsForAgencyRequest", string(data)}, " ")
}