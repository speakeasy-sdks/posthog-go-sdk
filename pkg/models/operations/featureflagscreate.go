package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FeatureFlagsCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsCreateRequests struct {
	FeatureFlag  *shared.FeatureFlagInput `request:"mediaType=application/json"`
	FeatureFlag1 *shared.FeatureFlagInput `request:"mediaType=application/x-www-form-urlencoded"`
	FeatureFlag2 *shared.FeatureFlagInput `request:"mediaType=multipart/form-data"`
}

type FeatureFlagsCreateRequest struct {
	PathParams FeatureFlagsCreatePathParams
	Request    FeatureFlagsCreateRequests
}

type FeatureFlagsCreateResponse struct {
	ContentType string
	FeatureFlag *shared.FeatureFlag
	StatusCode  int64
}
