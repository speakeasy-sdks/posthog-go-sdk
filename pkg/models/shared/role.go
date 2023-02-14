package shared

import (
	"time"
)

type RoleInput struct {
	FeatureFlagsAccessLevel *int64 `json:"feature_flags_access_level,omitempty" form:"name=feature_flags_access_level" multipartForm:"name=feature_flags_access_level"`
	Name                    string `json:"name" form:"name=name" multipartForm:"name=name"`
}

type RoleCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

type Role struct {
	AssociatedFlags         string        `json:"associated_flags"`
	CreatedAt               time.Time     `json:"created_at"`
	CreatedBy               RoleCreatedBy `json:"created_by"`
	FeatureFlagsAccessLevel *int64        `json:"feature_flags_access_level,omitempty"`
	ID                      string        `json:"id"`
	Members                 string        `json:"members"`
	Name                    string        `json:"name"`
}
