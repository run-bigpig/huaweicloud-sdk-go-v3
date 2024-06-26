package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBizMetricByIdResponse Response Object
type ShowBizMetricByIdResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowBizMetricByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBizMetricByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowBizMetricByIdResponse", string(data)}, " ")
}
