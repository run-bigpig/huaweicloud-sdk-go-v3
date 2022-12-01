package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAppResponse struct {

	// 对象id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppResponse struct{}"
	}

	return strings.Join([]string{"CreateAppResponse", string(data)}, " ")
}