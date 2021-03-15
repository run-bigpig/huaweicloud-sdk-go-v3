package model

import (
	"encoding/json"

	"strings"
)

// 项目信息
type GetProjectInfoV4ResultProject struct {
	// 项目numId
	ProjectNumId *int32 `json:"project_num_id,omitempty"`
	// 项目uuid
	ProjectId *string `json:"project_id,omitempty"`
	// 项目名称
	Name *string `json:"name,omitempty"`
	// 项目创建时间
	CreatedOn float32 `json:"created_on,omitempty"`
	// 项目更新时间
	UpdatedOn float32 `json:"updated_on,omitempty"`
	// 项目类型
	ProjectType *string `json:"project_type,omitempty"`
	// 是否归档
	Archive *bool `json:"archive,omitempty"`
	// 企业项目id
	EnterpriseId *string                               `json:"enterprise_id,omitempty"`
	Creator      *GetProjectInfoV4ResultProjectCreator `json:"creator,omitempty"`
}

func (o GetProjectInfoV4ResultProject) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GetProjectInfoV4ResultProject struct{}"
	}

	return strings.Join([]string{"GetProjectInfoV4ResultProject", string(data)}, " ")
}
