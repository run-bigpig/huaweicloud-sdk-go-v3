/*
 * cts
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type ListTrackersRequest struct {
	TrackerName string                         `json:"tracker_name,omitempty"`
	TrackerType ListTrackersRequestTrackerType `json:"tracker_type,omitempty"`
}

func (o ListTrackersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListTrackersRequest", string(data)}, " ")
}

type ListTrackersRequestTrackerType struct {
	value string
}

type ListTrackersRequestTrackerTypeEnum struct {
	SYSTEM ListTrackersRequestTrackerType
	DATA   ListTrackersRequestTrackerType
}

func GetListTrackersRequestTrackerTypeEnum() ListTrackersRequestTrackerTypeEnum {
	return ListTrackersRequestTrackerTypeEnum{
		SYSTEM: ListTrackersRequestTrackerType{
			value: "system",
		},
		DATA: ListTrackersRequestTrackerType{
			value: "data",
		},
	}
}

func (c ListTrackersRequestTrackerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListTrackersRequestTrackerType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}