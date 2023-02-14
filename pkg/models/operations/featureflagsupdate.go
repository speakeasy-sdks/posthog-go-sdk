package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FeatureFlagsUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsUpdateRequests struct {
	FeatureFlag  *shared.FeatureFlagInput `request:"mediaType=application/json"`
	FeatureFlag1 *shared.FeatureFlagInput `request:"mediaType=application/x-www-form-urlencoded"`
	FeatureFlag2 *shared.FeatureFlagInput `request:"mediaType=multipart/form-data"`
}

type FeatureFlagsUpdateRequest struct {
	PathParams FeatureFlagsUpdatePathParams
	Request    FeatureFlagsUpdateRequests
}

type FeatureFlagsUpdateResponse struct {
	ContentType string
	FeatureFlag *shared.FeatureFlag
	StatusCode  int64
}
