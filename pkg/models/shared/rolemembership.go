package shared

import (
	"time"
)

type RoleMembershipUser struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

type RoleMembershipOutput struct {
	ID        string             `json:"id"`
	JoinedAt  time.Time          `json:"joined_at"`
	RoleID    string             `json:"role_id"`
	UpdatedAt time.Time          `json:"updated_at"`
	User      RoleMembershipUser `json:"user"`
}

type RoleMembershipInput struct {
	UserUUID string `json:"user_uuid" form:"name=user_uuid" multipartForm:"name=user_uuid"`
}
