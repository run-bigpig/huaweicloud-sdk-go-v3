package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type StartInstanceSingleToHaActionRequest struct {
	XLanguage  *StartInstanceSingleToHaActionRequestXLanguage `json:"X-Language,omitempty"`
	InstanceId string                                         `json:"instance_id"`
	Body       *Single2Ha                                     `json:"body,omitempty"`
}

func (o StartInstanceSingleToHaActionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartInstanceSingleToHaActionRequest struct{}"
	}

	return strings.Join([]string{"StartInstanceSingleToHaActionRequest", string(data)}, " ")
}

type StartInstanceSingleToHaActionRequestXLanguage struct {
	value string
}

type StartInstanceSingleToHaActionRequestXLanguageEnum struct {
	ZH_CN StartInstanceSingleToHaActionRequestXLanguage
	EN_US StartInstanceSingleToHaActionRequestXLanguage
}

func GetStartInstanceSingleToHaActionRequestXLanguageEnum() StartInstanceSingleToHaActionRequestXLanguageEnum {
	return StartInstanceSingleToHaActionRequestXLanguageEnum{
		ZH_CN: StartInstanceSingleToHaActionRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartInstanceSingleToHaActionRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartInstanceSingleToHaActionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *StartInstanceSingleToHaActionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
