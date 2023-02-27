package shared

import (
	"time"
)

// PatchedActionInput
// Serializer mixin that resolves appropriate response for tags depending on license.
type PatchedActionInput struct {
	Deleted            *bool         `json:"deleted,omitempty" form:"name=deleted" multipartForm:"name=deleted"`
	Description        *string       `json:"description,omitempty" form:"name=description" multipartForm:"name=description"`
	LastCalculatedAt   *time.Time    `json:"last_calculated_at,omitempty" form:"name=last_calculated_at" multipartForm:"name=last_calculated_at"`
	Name               *string       `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
	PostToSlack        *bool         `json:"post_to_slack,omitempty" form:"name=post_to_slack" multipartForm:"name=post_to_slack"`
	SlackMessageFormat *string       `json:"slack_message_format,omitempty" form:"name=slack_message_format" multipartForm:"name=slack_message_format"`
	Steps              []ActionStep  `json:"steps,omitempty" form:"name=steps,json" multipartForm:"name=steps,json"`
	Tags               []interface{} `json:"tags,omitempty" form:"name=tags" multipartForm:"name=tags"`
}
