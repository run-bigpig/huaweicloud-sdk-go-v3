package model

import (
	"encoding/json"

	"strings"
)

// 查询上传的批量任务文件结构体。
type BatchTaskFile struct {
	// 上传的批量任务文件ID，由平台自动生成。
	FileId *string `json:"file_id,omitempty"`
	// 上传的批量任务文件名称。
	FileName *string `json:"file_name,omitempty"`
	// 在物联网平台上传文件的时间。格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。
	UploadTime *string `json:"upload_time,omitempty"`
}

func (o BatchTaskFile) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchTaskFile struct{}"
	}

	return strings.Join([]string{"BatchTaskFile", string(data)}, " ")
}
