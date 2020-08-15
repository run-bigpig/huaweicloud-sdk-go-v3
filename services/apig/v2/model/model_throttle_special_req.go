/*
 * APIG
 *
 * API网关（API Gateway）是为开发者、合作伙伴提供的高性能、高可用、高安全的API托管服务，帮助用户轻松构建、管理和发布任意规模的API。
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

type ThrottleSpecialReq struct {
	// 流控时间内特殊对象能够访问API的最大次数限制
	CallLimits int32 `json:"call_limits"`
	// 特殊APP的编号或特殊租户的账号ID
	ObjectId string `json:"object_id"`
	// 特殊对象类型
	ObjectType ThrottleSpecialReqObjectType `json:"object_type"`
}

func (o ThrottleSpecialReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ThrottleSpecialReq", string(data)}, " ")
}

type ThrottleSpecialReqObjectType struct {
	value string
}

type ThrottleSpecialReqObjectTypeEnum struct {
	APP  ThrottleSpecialReqObjectType
	USER ThrottleSpecialReqObjectType
}

func GetThrottleSpecialReqObjectTypeEnum() ThrottleSpecialReqObjectTypeEnum {
	return ThrottleSpecialReqObjectTypeEnum{
		APP: ThrottleSpecialReqObjectType{
			value: "APP",
		},
		USER: ThrottleSpecialReqObjectType{
			value: "USER",
		},
	}
}

func (c ThrottleSpecialReqObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ThrottleSpecialReqObjectType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}