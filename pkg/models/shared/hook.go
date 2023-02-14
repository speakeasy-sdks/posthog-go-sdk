package shared

import (
	"time"
)

type HookInput struct {
	Event      string  `json:"event" form:"name=event" multipartForm:"name=event"`
	ID         *string `json:"id,omitempty" form:"name=id" multipartForm:"name=id"`
	ResourceID *int64  `json:"resource_id,omitempty" form:"name=resource_id" multipartForm:"name=resource_id"`
	Target     string  `json:"target" form:"name=target" multipartForm:"name=target"`
}

type Hook struct {
	Created    time.Time `json:"created"`
	Event      string    `json:"event"`
	ID         *string   `json:"id,omitempty"`
	ResourceID *int64    `json:"resource_id,omitempty"`
	Target     string    `json:"target"`
	Team       int64     `json:"team"`
	Updated    time.Time `json:"updated"`
}
