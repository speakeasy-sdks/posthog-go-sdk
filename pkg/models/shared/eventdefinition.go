package shared

import (
	"time"
)

// EventDefinitionInput
// Serializer mixin that resolves appropriate response for tags depending on license.
type EventDefinitionInput struct {
	CreatedAt       *time.Time    `json:"created_at,omitempty" form:"name=created_at" multipartForm:"name=created_at"`
	LastSeenAt      *time.Time    `json:"last_seen_at,omitempty" form:"name=last_seen_at" multipartForm:"name=last_seen_at"`
	Name            string        `json:"name" form:"name=name" multipartForm:"name=name"`
	PostToSlack     *bool         `json:"post_to_slack,omitempty" form:"name=post_to_slack" multipartForm:"name=post_to_slack"`
	QueryUsage30Day *int64        `json:"query_usage_30_day,omitempty" form:"name=query_usage_30_day" multipartForm:"name=query_usage_30_day"`
	Tags            []interface{} `json:"tags,omitempty" form:"name=tags" multipartForm:"name=tags"`
	Volume30Day     *int64        `json:"volume_30_day,omitempty" form:"name=volume_30_day" multipartForm:"name=volume_30_day"`
}

type EventDefinitionCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

// EventDefinition
// Serializer mixin that resolves appropriate response for tags depending on license.
type EventDefinition struct {
	ActionID         int64                    `json:"action_id"`
	CreatedAt        *time.Time               `json:"created_at,omitempty"`
	CreatedBy        EventDefinitionCreatedBy `json:"created_by"`
	ID               string                   `json:"id"`
	IsAction         string                   `json:"is_action"`
	IsCalculating    bool                     `json:"is_calculating"`
	LastCalculatedAt time.Time                `json:"last_calculated_at"`
	LastSeenAt       *time.Time               `json:"last_seen_at,omitempty"`
	LastUpdatedAt    time.Time                `json:"last_updated_at"`
	Name             string                   `json:"name"`
	PostToSlack      *bool                    `json:"post_to_slack,omitempty"`
	QueryUsage30Day  *int64                   `json:"query_usage_30_day,omitempty"`
	Tags             []interface{}            `json:"tags,omitempty"`
	Volume30Day      *int64                   `json:"volume_30_day,omitempty"`
}
