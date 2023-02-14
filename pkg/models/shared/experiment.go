package shared

import (
	"time"
)

type ExperimentInput struct {
	Archived         *bool                  `json:"archived,omitempty" form:"name=archived" multipartForm:"name=archived"`
	Description      *string                `json:"description,omitempty" form:"name=description" multipartForm:"name=description"`
	EndDate          *time.Time             `json:"end_date,omitempty" form:"name=end_date" multipartForm:"name=end_date"`
	FeatureFlagKey   string                 `json:"feature_flag_key" form:"name=feature_flag_key" multipartForm:"name=feature_flag_key"`
	Filters          map[string]interface{} `json:"filters,omitempty" form:"name=filters,json" multipartForm:"name=filters,json"`
	Name             string                 `json:"name" form:"name=name" multipartForm:"name=name"`
	Parameters       map[string]interface{} `json:"parameters,omitempty" form:"name=parameters,json" multipartForm:"name=parameters,json"`
	SecondaryMetrics map[string]interface{} `json:"secondary_metrics,omitempty" form:"name=secondary_metrics,json" multipartForm:"name=secondary_metrics,json"`
	StartDate        *time.Time             `json:"start_date,omitempty" form:"name=start_date" multipartForm:"name=start_date"`
}

type ExperimentCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

type Experiment struct {
	Archived         *bool                  `json:"archived,omitempty"`
	CreatedAt        time.Time              `json:"created_at"`
	CreatedBy        ExperimentCreatedBy    `json:"created_by"`
	Description      *string                `json:"description,omitempty"`
	EndDate          *time.Time             `json:"end_date,omitempty"`
	FeatureFlag      int64                  `json:"feature_flag"`
	FeatureFlagKey   string                 `json:"feature_flag_key"`
	Filters          map[string]interface{} `json:"filters,omitempty"`
	ID               int64                  `json:"id"`
	Name             string                 `json:"name"`
	Parameters       map[string]interface{} `json:"parameters,omitempty"`
	SecondaryMetrics map[string]interface{} `json:"secondary_metrics,omitempty"`
	StartDate        *time.Time             `json:"start_date,omitempty"`
	UpdatedAt        time.Time              `json:"updated_at"`
}
