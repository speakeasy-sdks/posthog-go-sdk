package shared

import (
	"time"
)

// FeatureFlagInput
// Serializer mixin that resolves appropriate response for tags depending on license.
type FeatureFlagInput struct {
	Active                     *bool                  `json:"active,omitempty" form:"name=active" multipartForm:"name=active"`
	CreatedAt                  *time.Time             `json:"created_at,omitempty" form:"name=created_at" multipartForm:"name=created_at"`
	Deleted                    *bool                  `json:"deleted,omitempty" form:"name=deleted" multipartForm:"name=deleted"`
	EnsureExperienceContinuity *bool                  `json:"ensure_experience_continuity,omitempty" form:"name=ensure_experience_continuity" multipartForm:"name=ensure_experience_continuity"`
	Filters                    map[string]interface{} `json:"filters,omitempty" form:"name=filters,json" multipartForm:"name=filters,json"`
	Key                        string                 `json:"key" form:"name=key" multipartForm:"name=key"`
	Name                       *string                `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
	PerformedRollback          *bool                  `json:"performed_rollback,omitempty" form:"name=performed_rollback" multipartForm:"name=performed_rollback"`
	RollbackConditions         map[string]interface{} `json:"rollback_conditions,omitempty" form:"name=rollback_conditions,json" multipartForm:"name=rollback_conditions,json"`
	Tags                       []interface{}          `json:"tags,omitempty" form:"name=tags" multipartForm:"name=tags"`
}

type FeatureFlagCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

// FeatureFlag
// Serializer mixin that resolves appropriate response for tags depending on license.
type FeatureFlag struct {
	Active                     *bool                  `json:"active,omitempty"`
	CanEdit                    bool                   `json:"can_edit"`
	CreatedAt                  *time.Time             `json:"created_at,omitempty"`
	CreatedBy                  FeatureFlagCreatedBy   `json:"created_by"`
	Deleted                    *bool                  `json:"deleted,omitempty"`
	EnsureExperienceContinuity *bool                  `json:"ensure_experience_continuity,omitempty"`
	ExperimentSet              []int64                `json:"experiment_set"`
	Filters                    map[string]interface{} `json:"filters,omitempty"`
	ID                         int64                  `json:"id"`
	IsSimpleFlag               bool                   `json:"is_simple_flag"`
	Key                        string                 `json:"key"`
	Name                       *string                `json:"name,omitempty"`
	PerformedRollback          *bool                  `json:"performed_rollback,omitempty"`
	RollbackConditions         map[string]interface{} `json:"rollback_conditions,omitempty"`
	RolloutPercentage          int64                  `json:"rollout_percentage"`
	Tags                       []interface{}          `json:"tags,omitempty"`
}
