package shared

import (
	"time"
)

type Person struct {
	CreatedAt   time.Time              `json:"created_at"`
	DistinctIds []string               `json:"distinct_ids"`
	ID          int64                  `json:"id"`
	Name        string                 `json:"name"`
	Properties  map[string]interface{} `json:"properties,omitempty"`
	UUID        string                 `json:"uuid"`
}

type PersonInput struct {
	Properties map[string]interface{} `json:"properties,omitempty" form:"name=properties,json" multipartForm:"name=properties,json"`
}
