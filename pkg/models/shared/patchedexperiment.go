package shared

import (
	"time"
)

type PatchedExperimentInput struct {
	Archived         *bool                  `json:"archived,omitempty" form:"name=archived" multipartForm:"name=archived"`
	Description      *string                `json:"description,omitempty" form:"name=description" multipartForm:"name=description"`
	EndDate          *time.Time             `json:"end_date,omitempty" form:"name=end_date" multipartForm:"name=end_date"`
	FeatureFlagKey   *string                `json:"feature_flag_key,omitempty" form:"name=feature_flag_key" multipartForm:"name=feature_flag_key"`
	Filters          map[string]interface{} `json:"filters,omitempty" form:"name=filters,json" multipartForm:"name=filters,json"`
	Name             *string                `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
	Parameters       map[string]interface{} `json:"parameters,omitempty" form:"name=parameters,json" multipartForm:"name=parameters,json"`
	SecondaryMetrics map[string]interface{} `json:"secondary_metrics,omitempty" form:"name=secondary_metrics,json" multipartForm:"name=secondary_metrics,json"`
	StartDate        *time.Time             `json:"start_date,omitempty" form:"name=start_date" multipartForm:"name=start_date"`
}
