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
type ShowVaultProjectTagResponse struct {
	// 标签列表
	Tags           *[]TagsResp `json:"tags,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowVaultProjectTagResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVaultProjectTagResponse", string(data)}, " ")
}