package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDirectoryResponse Response Object
type DeleteDirectoryResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteDirectoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDirectoryResponse struct{}"
	}

	return strings.Join([]string{"DeleteDirectoryResponse", string(data)}, " ")
}
