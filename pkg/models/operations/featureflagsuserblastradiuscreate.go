package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FeatureFlagsUserBlastRadiusCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsUserBlastRadiusCreateRequests struct {
	FeatureFlag  *shared.FeatureFlagInput `request:"mediaType=application/json"`
	FeatureFlag1 *shared.FeatureFlagInput `request:"mediaType=application/x-www-form-urlencoded"`
	FeatureFlag2 *shared.FeatureFlagInput `request:"mediaType=multipart/form-data"`
}

type FeatureFlagsUserBlastRadiusCreateRequest struct {
	PathParams FeatureFlagsUserBlastRadiusCreatePathParams
	Request    FeatureFlagsUserBlastRadiusCreateRequests
}

type FeatureFlagsUserBlastRadiusCreateResponse struct {
	ContentType string
	FeatureFlag *shared.FeatureFlag
	StatusCode  int
}
