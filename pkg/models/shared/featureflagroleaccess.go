package shared

import (
	"time"
)

type FeatureFlagRoleAccessFeatureFlagCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

// FeatureFlagRoleAccessFeatureFlag
// Serializer mixin that resolves appropriate response for tags depending on license.
type FeatureFlagRoleAccessFeatureFlag struct {
	Active                     *bool                                     `json:"active,omitempty"`
	CanEdit                    bool                                      `json:"can_edit"`
	CreatedAt                  *time.Time                                `json:"created_at,omitempty"`
	CreatedBy                  FeatureFlagRoleAccessFeatureFlagCreatedBy `json:"created_by"`
	Deleted                    *bool                                     `json:"deleted,omitempty"`
	EnsureExperienceContinuity *bool                                     `json:"ensure_experience_continuity,omitempty"`
	ExperimentSet              []int64                                   `json:"experiment_set"`
	Filters                    map[string]interface{}                    `json:"filters,omitempty"`
	ID                         int64                                     `json:"id"`
	IsSimpleFlag               bool                                      `json:"is_simple_flag"`
	Key                        string                                    `json:"key"`
	Name                       *string                                   `json:"name,omitempty"`
	PerformedRollback          *bool                                     `json:"performed_rollback,omitempty"`
	RollbackConditions         map[string]interface{}                    `json:"rollback_conditions,omitempty"`
	RolloutPercentage          int64                                     `json:"rollout_percentage"`
	Tags                       []interface{}                             `json:"tags,omitempty"`
}

type FeatureFlagRoleAccessRoleCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

type FeatureFlagRoleAccessRole struct {
	AssociatedFlags         string                             `json:"associated_flags"`
	CreatedAt               time.Time                          `json:"created_at"`
	CreatedBy               FeatureFlagRoleAccessRoleCreatedBy `json:"created_by"`
	FeatureFlagsAccessLevel *int64                             `json:"feature_flags_access_level,omitempty"`
	ID                      string                             `json:"id"`
	Members                 string                             `json:"members"`
	Name                    string                             `json:"name"`
}

type FeatureFlagRoleAccessOutput struct {
	AddedAt     time.Time                        `json:"added_at"`
	FeatureFlag FeatureFlagRoleAccessFeatureFlag `json:"feature_flag"`
	ID          int64                            `json:"id"`
	Role        FeatureFlagRoleAccessRole        `json:"role"`
	UpdatedAt   time.Time                        `json:"updated_at"`
}

type FeatureFlagRoleAccessInput struct {
	RoleID string `json:"role_id" form:"name=role_id" multipartForm:"name=role_id"`
}
