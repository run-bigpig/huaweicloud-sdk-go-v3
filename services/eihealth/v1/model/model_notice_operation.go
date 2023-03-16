package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NoticeOperation struct {
	value string
}

type NoticeOperationEnum struct {
	READ NoticeOperation
}

func GetNoticeOperationEnum() NoticeOperationEnum {
	return NoticeOperationEnum{
		READ: NoticeOperation{
			value: "read",
		},
	}
}

func (c NoticeOperation) Value() string {
	return c.value
}

func (c NoticeOperation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NoticeOperation) UnmarshalJSON(b []byte) error {
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
