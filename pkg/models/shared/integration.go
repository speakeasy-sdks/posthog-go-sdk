package shared

import (
	"time"
)

type IntegrationCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

type IntegrationKindEnum string

const (
	IntegrationKindEnumSlack IntegrationKindEnum = "slack"
)

// Integration
// Standard Integration serializer.
type Integration struct {
	Config    map[string]interface{} `json:"config,omitempty"`
	CreatedAt time.Time              `json:"created_at"`
	CreatedBy IntegrationCreatedBy   `json:"created_by"`
	Errors    string                 `json:"errors"`
	ID        int64                  `json:"id"`
	Kind      IntegrationKindEnum    `json:"kind"`
}

// IntegrationInput
// Standard Integration serializer.
type IntegrationInput struct {
	Config map[string]interface{} `json:"config,omitempty" form:"name=config,json" multipartForm:"name=config,json"`
	Kind   IntegrationKindEnum    `json:"kind" form:"name=kind" multipartForm:"name=kind"`
}
