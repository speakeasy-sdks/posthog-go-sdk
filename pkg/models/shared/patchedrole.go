package shared

type PatchedRoleInput struct {
	FeatureFlagsAccessLevel *int64  `json:"feature_flags_access_level,omitempty" form:"name=feature_flags_access_level" multipartForm:"name=feature_flags_access_level"`
	Name                    *string `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
}
