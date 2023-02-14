package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FeatureFlagsRoleAccessRetrievePathParams struct {
	ID                        int64  `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupFeatureFlagID string `pathParam:"style=simple,explode=false,name=parent_lookup_feature_flag_id"`
	ProjectID                 string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsRoleAccessRetrieveRequest struct {
	PathParams FeatureFlagsRoleAccessRetrievePathParams
}

type FeatureFlagsRoleAccessRetrieveResponse struct {
	ContentType           string
	FeatureFlagRoleAccess *shared.FeatureFlagRoleAccessOutput
	StatusCode            int64
}
