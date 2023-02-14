package shared

import (
	"time"
)

type OrganizationMemberInput struct {
	Level *int64 `json:"level,omitempty" form:"name=level" multipartForm:"name=level"`
}

type OrganizationMemberUser struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

type OrganizationMember struct {
	ID        string                 `json:"id"`
	JoinedAt  time.Time              `json:"joined_at"`
	Level     *int64                 `json:"level,omitempty"`
	UpdatedAt time.Time              `json:"updated_at"`
	User      OrganizationMemberUser `json:"user"`
}
