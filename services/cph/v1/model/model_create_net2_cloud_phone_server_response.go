package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateNet2CloudPhoneServerResponse struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id"`

	// 订单ID，不超过64个字节
	OrderId string `json:"order_id"`

	// 产品ID，不超过64个字节
	ProductId      string `json:"product_id"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateNet2CloudPhoneServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNet2CloudPhoneServerResponse struct{}"
	}

	return strings.Join([]string{"CreateNet2CloudPhoneServerResponse", string(data)}, " ")
}