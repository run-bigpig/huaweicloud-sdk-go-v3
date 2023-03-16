package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreatePolicyRequest struct {

	// 选择接口返回的信息的语言，默认为\"zh-cn\"中文，zh-cn中文，en-us英文
	XLanguage *CreatePolicyRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreatePolicyReqBody `json:"body,omitempty"`
}

func (o CreatePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyRequest struct{}"
	}

	return strings.Join([]string{"CreatePolicyRequest", string(data)}, " ")
}

type CreatePolicyRequestXLanguage struct {
	value string
}

type CreatePolicyRequestXLanguageEnum struct {
	ZH_CN CreatePolicyRequestXLanguage
	EN_US CreatePolicyRequestXLanguage
}

func GetCreatePolicyRequestXLanguageEnum() CreatePolicyRequestXLanguageEnum {
	return CreatePolicyRequestXLanguageEnum{
		ZH_CN: CreatePolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreatePolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreatePolicyRequestXLanguage) Value() string {
	return c.value
}

func (c CreatePolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
