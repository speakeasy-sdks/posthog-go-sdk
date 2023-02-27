package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FeatureFlagsActivityRetrieve2PathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsActivityRetrieve2Request struct {
	PathParams FeatureFlagsActivityRetrieve2PathParams
}

type FeatureFlagsActivityRetrieve2Response struct {
	ContentType string
	FeatureFlag *shared.FeatureFlag
	StatusCode  int
}
