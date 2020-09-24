/*
 * ELB
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
type UpdateHealthMonitorRequest struct {
	HealthmonitorId string                          `json:"healthmonitor_id"`
	Body            *UpdateHealthMonitorRequestBody `json:"body,omitempty"`
}

func (o UpdateHealthMonitorRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateHealthMonitorRequest", string(data)}, " ")
}