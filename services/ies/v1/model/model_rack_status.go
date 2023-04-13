package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 机柜状态。 - CREATING：创建中 - AVAILABLE：可用 - EXPANDING：扩容中 - PENDING_PAYMENT: 待支付
type RackStatus struct {
	value string
}

type RackStatusEnum struct {
	CREATING        RackStatus
	AVAILABLE       RackStatus
	EXPANDING       RackStatus
	PENDING_PAYMENT RackStatus
}

func GetRackStatusEnum() RackStatusEnum {
	return RackStatusEnum{
		CREATING: RackStatus{
			value: "CREATING",
		},
		AVAILABLE: RackStatus{
			value: "AVAILABLE",
		},
		EXPANDING: RackStatus{
			value: "EXPANDING",
		},
		PENDING_PAYMENT: RackStatus{
			value: "PENDING_PAYMENT",
		},
	}
}

func (c RackStatus) Value() string {
	return c.value
}

func (c RackStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RackStatus) UnmarshalJSON(b []byte) error {
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
