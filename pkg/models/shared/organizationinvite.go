package shared

import (
	"time"
)

type OrganizationInviteCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

type OrganizationInvite struct {
	CreatedAt           time.Time                   `json:"created_at"`
	CreatedBy           OrganizationInviteCreatedBy `json:"created_by"`
	EmailingAttemptMade bool                        `json:"emailing_attempt_made"`
	FirstName           *string                     `json:"first_name,omitempty"`
	ID                  string                      `json:"id"`
	IsExpired           bool                        `json:"is_expired"`
	Message             *string                     `json:"message,omitempty"`
	TargetEmail         string                      `json:"target_email"`
	UpdatedAt           time.Time                   `json:"updated_at"`
}

type OrganizationInviteInput struct {
	FirstName   *string `json:"first_name,omitempty" form:"name=first_name" multipartForm:"name=first_name"`
	Message     *string `json:"message,omitempty" form:"name=message" multipartForm:"name=message"`
	TargetEmail string  `json:"target_email" form:"name=target_email" multipartForm:"name=target_email"`
}
