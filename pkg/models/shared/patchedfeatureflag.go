package shared

import (
	"time"
)

// PatchedFeatureFlagInput
// Serializer mixin that resolves appropriate response for tags depending on license.
type PatchedFeatureFlagInput struct {
	Active                     *bool                  `json:"active,omitempty" form:"name=active" multipartForm:"name=active"`
	CreatedAt                  *time.Time             `json:"created_at,omitempty" form:"name=created_at" multipartForm:"name=created_at"`
	Deleted                    *bool                  `json:"deleted,omitempty" form:"name=deleted" multipartForm:"name=deleted"`
	EnsureExperienceContinuity *bool                  `json:"ensure_experience_continuity,omitempty" form:"name=ensure_experience_continuity" multipartForm:"name=ensure_experience_continuity"`
	Filters                    map[string]interface{} `json:"filters,omitempty" form:"name=filters,json" multipartForm:"name=filters,json"`
	Key                        *string                `json:"key,omitempty" form:"name=key" multipartForm:"name=key"`
	Name                       *string                `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
	PerformedRollback          *bool                  `json:"performed_rollback,omitempty" form:"name=performed_rollback" multipartForm:"name=performed_rollback"`
	RollbackConditions         map[string]interface{} `json:"rollback_conditions,omitempty" form:"name=rollback_conditions,json" multipartForm:"name=rollback_conditions,json"`
	Tags                       []interface{}          `json:"tags,omitempty" form:"name=tags,json" multipartForm:"name=tags,json"`
}
