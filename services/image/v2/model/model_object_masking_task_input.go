package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObjectMaskingTaskInput struct {

	// 任务的输入类型。可选类型有obs（对象存储服务存储的文件），url（指定的文件地址）
	Type string `json:"type"`

	// 任务的输入详情。针对不同的输入类型有不同的配置。
	Data []ObjectMaskingTaskInputData `json:"data"`
}

func (o ObjectMaskingTaskInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectMaskingTaskInput struct{}"
	}

	return strings.Join([]string{"ObjectMaskingTaskInput", string(data)}, " ")
}