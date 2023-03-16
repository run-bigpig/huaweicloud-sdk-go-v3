package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AttachPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AttachPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachPolicyResponse struct{}"
	}

	return strings.Join([]string{"AttachPolicyResponse", string(data)}, " ")
}
