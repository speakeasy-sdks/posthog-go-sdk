package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FeatureFlagsMyFlagsRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsMyFlagsRetrieveRequest struct {
	PathParams FeatureFlagsMyFlagsRetrievePathParams
}

type FeatureFlagsMyFlagsRetrieveResponse struct {
	ContentType string
	FeatureFlag *shared.FeatureFlag
	StatusCode  int
}
