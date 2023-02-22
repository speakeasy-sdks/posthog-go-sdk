package operations

type FeatureFlagsRoleAccessDestroyPathParams struct {
	ID                        int64  `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupFeatureFlagID string `pathParam:"style=simple,explode=false,name=parent_lookup_feature_flag_id"`
	ProjectID                 string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsRoleAccessDestroyRequest struct {
	PathParams FeatureFlagsRoleAccessDestroyPathParams
}

type FeatureFlagsRoleAccessDestroyResponse struct {
	ContentType string
	StatusCode  int
}
