package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StandElementValueVo 标准属性值。
type StandElementValueVo struct {

	// 属性名称。
	FdName string `json:"fd_name"`

	// 属性值。
	FdValue *string `json:"fd_value,omitempty"`

	// 属性定义的ID。
	FdId *int64 `json:"fd_id,omitempty"`

	// 标准所属目录。
	DirectoryId *int64 `json:"directory_id,omitempty"`

	// 标准所属行。
	RowId *int64 `json:"row_id,omitempty"`

	// 数据标准的ID。
	Id *int64 `json:"id,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o StandElementValueVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StandElementValueVo struct{}"
	}

	return strings.Join([]string{"StandElementValueVo", string(data)}, " ")
}
