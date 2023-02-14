package shared

import (
	"time"
)

type ActivityLog struct {
	Activity       string                 `json:"activity"`
	CreatedAt      *time.Time             `json:"created_at,omitempty"`
	Detail         map[string]interface{} `json:"detail,omitempty"`
	ID             string                 `json:"id"`
	IsSystem       *bool                  `json:"is_system,omitempty"`
	ItemID         *string                `json:"item_id,omitempty"`
	OrganizationID *string                `json:"organization_id,omitempty"`
	Scope          string                 `json:"scope"`
	Unread         bool                   `json:"unread"`
	User           UserBasic              `json:"user"`
}

type ActivityLogInput struct {
	Activity       string                 `json:"activity" form:"name=activity" multipartForm:"name=activity"`
	CreatedAt      *time.Time             `json:"created_at,omitempty" form:"name=created_at" multipartForm:"name=created_at"`
	Detail         map[string]interface{} `json:"detail,omitempty" form:"name=detail,json" multipartForm:"name=detail,json"`
	IsSystem       *bool                  `json:"is_system,omitempty" form:"name=is_system" multipartForm:"name=is_system"`
	ItemID         *string                `json:"item_id,omitempty" form:"name=item_id" multipartForm:"name=item_id"`
	OrganizationID *string                `json:"organization_id,omitempty" form:"name=organization_id" multipartForm:"name=organization_id"`
	Scope          string                 `json:"scope" form:"name=scope" multipartForm:"name=scope"`
	User           UserBasicInput         `json:"user" form:"name=user,json" multipartForm:"name=user,json"`
}
