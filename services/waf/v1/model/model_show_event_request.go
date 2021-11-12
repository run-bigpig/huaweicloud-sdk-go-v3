package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowEventRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 事件ID

	Eventid string `json:"eventid"`
}

func (o ShowEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventRequest struct{}"
	}

	return strings.Join([]string{"ShowEventRequest", string(data)}, " ")
}