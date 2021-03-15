package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteRoutingRuleResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRoutingRuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRoutingRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteRoutingRuleResponse", string(data)}, " ")
}
