package shared

import (
	"time"
)

type CohortInput struct {
	Deleted     *bool                  `json:"deleted,omitempty" form:"name=deleted" multipartForm:"name=deleted"`
	Description *string                `json:"description,omitempty" form:"name=description" multipartForm:"name=description"`
	Filters     map[string]interface{} `json:"filters,omitempty" form:"name=filters,json" multipartForm:"name=filters,json"`
	Groups      map[string]interface{} `json:"groups,omitempty" form:"name=groups,json" multipartForm:"name=groups,json"`
	IsStatic    *bool                  `json:"is_static,omitempty" form:"name=is_static" multipartForm:"name=is_static"`
	Name        *string                `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
}

type CohortCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

type Cohort struct {
	Count             int64                  `json:"count"`
	CreatedAt         time.Time              `json:"created_at"`
	CreatedBy         CohortCreatedBy        `json:"created_by"`
	Deleted           *bool                  `json:"deleted,omitempty"`
	Description       *string                `json:"description,omitempty"`
	ErrorsCalculating int64                  `json:"errors_calculating"`
	Filters           map[string]interface{} `json:"filters,omitempty"`
	Groups            map[string]interface{} `json:"groups,omitempty"`
	ID                int64                  `json:"id"`
	IsCalculating     bool                   `json:"is_calculating"`
	IsStatic          *bool                  `json:"is_static,omitempty"`
	LastCalculation   time.Time              `json:"last_calculation"`
	Name              *string                `json:"name,omitempty"`
}
