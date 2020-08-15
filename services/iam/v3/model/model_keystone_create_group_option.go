/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

//
type KeystoneCreateGroupOption struct {
	// 用户组描述信息，长度小于等于255字节。
	Description string `json:"description,omitempty"`
	// 用户组所属账号ID，获取方式请参见：[获取账号ID](https://support.huaweicloud.com/api-iam/zh-cn_topic_0057845624.html)。
	DomainId string `json:"domain_id,omitempty"`
	// 用户组名，长度小于等于64字节。
	Name string `json:"name"`
}

func (o KeystoneCreateGroupOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneCreateGroupOption", string(data)}, " ")
}