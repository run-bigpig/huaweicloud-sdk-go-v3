package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListImageTagsRequest struct {
	ImageId string `json:"image_id"`
}

func (o ListImageTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListImageTagsRequest struct{}"
	}

	return strings.Join([]string{"ListImageTagsRequest", string(data)}, " ")
}
