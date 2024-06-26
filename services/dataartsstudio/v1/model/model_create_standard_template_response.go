package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStandardTemplateResponse Response Object
type CreateStandardTemplateResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateStandardTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStandardTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateStandardTemplateResponse", string(data)}, " ")
}
