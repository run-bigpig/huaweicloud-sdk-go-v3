package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowStoragePoolResponse struct {
	StoragePool    *StoragePool `json:"storage_pool,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowStoragePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStoragePoolResponse struct{}"
	}

	return strings.Join([]string{"ShowStoragePoolResponse", string(data)}, " ")
}
