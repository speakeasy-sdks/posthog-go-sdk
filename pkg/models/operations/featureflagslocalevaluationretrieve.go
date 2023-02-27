package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FeatureFlagsLocalEvaluationRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsLocalEvaluationRetrieveRequest struct {
	PathParams FeatureFlagsLocalEvaluationRetrievePathParams
}

type FeatureFlagsLocalEvaluationRetrieveResponse struct {
	ContentType string
	FeatureFlag *shared.FeatureFlag
	StatusCode  int
}
