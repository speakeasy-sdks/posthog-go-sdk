package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FeatureFlagsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsRetrieveRequest struct {
	PathParams FeatureFlagsRetrievePathParams
}

type FeatureFlagsRetrieveResponse struct {
	ContentType string
	FeatureFlag *shared.FeatureFlag
	StatusCode  int64
}
