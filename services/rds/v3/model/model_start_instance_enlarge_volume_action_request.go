package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type StartInstanceEnlargeVolumeActionRequest struct {
	XLanguage  *StartInstanceEnlargeVolumeActionRequestXLanguage `json:"X-Language,omitempty"`
	InstanceId string                                            `json:"instance_id"`
	Body       *EnlargeVolume                                    `json:"body,omitempty"`
}

func (o StartInstanceEnlargeVolumeActionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartInstanceEnlargeVolumeActionRequest struct{}"
	}

	return strings.Join([]string{"StartInstanceEnlargeVolumeActionRequest", string(data)}, " ")
}

type StartInstanceEnlargeVolumeActionRequestXLanguage struct {
	value string
}

type StartInstanceEnlargeVolumeActionRequestXLanguageEnum struct {
	ZH_CN StartInstanceEnlargeVolumeActionRequestXLanguage
	EN_US StartInstanceEnlargeVolumeActionRequestXLanguage
}

func GetStartInstanceEnlargeVolumeActionRequestXLanguageEnum() StartInstanceEnlargeVolumeActionRequestXLanguageEnum {
	return StartInstanceEnlargeVolumeActionRequestXLanguageEnum{
		ZH_CN: StartInstanceEnlargeVolumeActionRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartInstanceEnlargeVolumeActionRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartInstanceEnlargeVolumeActionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *StartInstanceEnlargeVolumeActionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
