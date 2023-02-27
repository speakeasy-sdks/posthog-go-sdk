package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FeatureFlagsRoleAccessCreatePathParams struct {
	ParentLookupFeatureFlagID string `pathParam:"style=simple,explode=false,name=parent_lookup_feature_flag_id"`
	ProjectID                 string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsRoleAccessCreateRequests struct {
	FeatureFlagRoleAccess  *shared.FeatureFlagRoleAccessInput `request:"mediaType=application/json"`
	FeatureFlagRoleAccess1 *shared.FeatureFlagRoleAccessInput `request:"mediaType=application/x-www-form-urlencoded"`
	FeatureFlagRoleAccess2 *shared.FeatureFlagRoleAccessInput `request:"mediaType=multipart/form-data"`
}

type FeatureFlagsRoleAccessCreateRequest struct {
	PathParams FeatureFlagsRoleAccessCreatePathParams
	Request    FeatureFlagsRoleAccessCreateRequests
}

type FeatureFlagsRoleAccessCreateResponse struct {
	ContentType           string
	FeatureFlagRoleAccess *shared.FeatureFlagRoleAccessOutput
	StatusCode            int
}
