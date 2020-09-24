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
type ListFlavorsRequest struct {
	Marker      *string   `json:"marker,omitempty"`
	Limit       *int32    `json:"limit,omitempty"`
	PageReverse *bool     `json:"page_reverse,omitempty"`
	Id          *[]string `json:"id,omitempty"`
	Name        *[]string `json:"name,omitempty"`
	Type        *[]string `json:"type,omitempty"`
	Shared      *bool     `json:"shared,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}