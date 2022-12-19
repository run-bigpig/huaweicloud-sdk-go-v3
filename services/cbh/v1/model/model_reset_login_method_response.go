package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResetLoginMethodResponse struct {

	// 操作结果
	Code *int32 `json:"code,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 任务 id
	TaskId *string `json:"task_id,omitempty"`

	// 订单 id
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetLoginMethodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetLoginMethodResponse struct{}"
	}

	return strings.Join([]string{"ResetLoginMethodResponse", string(data)}, " ")
}