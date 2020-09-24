/*
 * CES
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
type ShowEventDataResponse struct {
	//
	Datapoints *[]EventDataInfo `json:"datapoints,omitempty"`
}

func (o ShowEventDataResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowEventDataResponse", string(data)}, " ")
}