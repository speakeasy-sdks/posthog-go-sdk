package shared

import (
	"time"
)

// ActionInput
// Serializer mixin that resolves appropriate response for tags depending on license.
type ActionInput struct {
	Deleted            *bool         `json:"deleted,omitempty" form:"name=deleted" multipartForm:"name=deleted"`
	Description        *string       `json:"description,omitempty" form:"name=description" multipartForm:"name=description"`
	LastCalculatedAt   *time.Time    `json:"last_calculated_at,omitempty" form:"name=last_calculated_at" multipartForm:"name=last_calculated_at"`
	Name               *string       `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
	PostToSlack        *bool         `json:"post_to_slack,omitempty" form:"name=post_to_slack" multipartForm:"name=post_to_slack"`
	SlackMessageFormat *string       `json:"slack_message_format,omitempty" form:"name=slack_message_format" multipartForm:"name=slack_message_format"`
	Steps              []ActionStep  `json:"steps,omitempty" form:"name=steps,json" multipartForm:"name=steps,json"`
	Tags               []interface{} `json:"tags,omitempty" form:"name=tags,json" multipartForm:"name=tags,json"`
}

type ActionCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

// Action
// Serializer mixin that resolves appropriate response for tags depending on license.
type Action struct {
	CreatedAt          time.Time       `json:"created_at"`
	CreatedBy          ActionCreatedBy `json:"created_by"`
	Deleted            *bool           `json:"deleted,omitempty"`
	Description        *string         `json:"description,omitempty"`
	ID                 int64           `json:"id"`
	IsAction           bool            `json:"is_action"`
	IsCalculating      bool            `json:"is_calculating"`
	LastCalculatedAt   *time.Time      `json:"last_calculated_at,omitempty"`
	Name               *string         `json:"name,omitempty"`
	PostToSlack        *bool           `json:"post_to_slack,omitempty"`
	SlackMessageFormat *string         `json:"slack_message_format,omitempty"`
	Steps              []ActionStep    `json:"steps,omitempty"`
	Tags               []interface{}   `json:"tags,omitempty"`
	TeamID             int64           `json:"team_id"`
}
