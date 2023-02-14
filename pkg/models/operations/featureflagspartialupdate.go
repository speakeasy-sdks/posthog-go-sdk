package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FeatureFlagsPartialUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsPartialUpdateRequests struct {
	PatchedFeatureFlag  *shared.PatchedFeatureFlagInput `request:"mediaType=application/json"`
	PatchedFeatureFlag1 *shared.PatchedFeatureFlagInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedFeatureFlag2 *shared.PatchedFeatureFlagInput `request:"mediaType=multipart/form-data"`
}

type FeatureFlagsPartialUpdateRequest struct {
	PathParams FeatureFlagsPartialUpdatePathParams
	Request    *FeatureFlagsPartialUpdateRequests
}

type FeatureFlagsPartialUpdateResponse struct {
	ContentType string
	FeatureFlag *shared.FeatureFlag
	StatusCode  int64
}
