package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FeatureFlagsActivityRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsActivityRetrieveRequest struct {
	PathParams FeatureFlagsActivityRetrievePathParams
}

type FeatureFlagsActivityRetrieveResponse struct {
	ContentType string
	FeatureFlag *shared.FeatureFlag
	StatusCode  int
}
