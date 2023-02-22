package shared

import (
	"time"
)

// PatchedEventDefinitionInput
// Serializer mixin that resolves appropriate response for tags depending on license.
type PatchedEventDefinitionInput struct {
	CreatedAt       *time.Time    `json:"created_at,omitempty" form:"name=created_at" multipartForm:"name=created_at"`
	LastSeenAt      *time.Time    `json:"last_seen_at,omitempty" form:"name=last_seen_at" multipartForm:"name=last_seen_at"`
	Name            *string       `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
	PostToSlack     *bool         `json:"post_to_slack,omitempty" form:"name=post_to_slack" multipartForm:"name=post_to_slack"`
	QueryUsage30Day *int64        `json:"query_usage_30_day,omitempty" form:"name=query_usage_30_day" multipartForm:"name=query_usage_30_day"`
	Tags            []interface{} `json:"tags,omitempty" form:"name=tags" multipartForm:"name=tags"`
	Volume30Day     *int64        `json:"volume_30_day,omitempty" form:"name=volume_30_day" multipartForm:"name=volume_30_day"`
}
