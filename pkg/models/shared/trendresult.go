package shared

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/types"
)

type TrendResult struct {
	Data   []int64         `json:"data"`
	Days   []types.Date    `json:"days"`
	Filter GenericInsights `json:"filter"`
	Label  string          `json:"label"`
	Labels []string        `json:"labels"`
}
