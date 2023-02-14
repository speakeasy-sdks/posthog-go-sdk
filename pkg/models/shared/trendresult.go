package shared

import (
	"time"
)

type TrendResult struct {
	Data   []int64         `json:"data"`
	Days   []time.Time     `json:"days"`
	Filter GenericInsights `json:"filter"`
	Label  string          `json:"label"`
	Labels []string        `json:"labels"`
}
