package shared

import (
	"time"
)

type Group struct {
	CreatedAt       time.Time              `json:"created_at"`
	GroupKey        string                 `json:"group_key"`
	GroupProperties map[string]interface{} `json:"group_properties,omitempty"`
	GroupTypeIndex  int64                  `json:"group_type_index"`
}
