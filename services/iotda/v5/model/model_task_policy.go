package model

import (
	"encoding/json"

	"strings"
)

type TaskPolicy struct {
	// 批量任务指定执行时间，取值范围为7天内，不传入此参数表示立即执行，格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。
	ScheduleTime *string `json:"schedule_time,omitempty"`
	// 批量任务子任务自动重试次数，如果传入retry_interval参数，则需传入该参数，最大支持重试5次。
	RetryCount *int32 `json:"retry_count,omitempty"`
	// 批量任务子任务失败后，自动重试时间间隔，单位：分钟，最大1440(24小时)，不传入此参数表示不重试，如果传入retry_count参数则需要传入该参数。
	RetryInterval *int32 `json:"retry_interval,omitempty"`
}

func (o TaskPolicy) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TaskPolicy struct{}"
	}

	return strings.Join([]string{"TaskPolicy", string(data)}, " ")
}
