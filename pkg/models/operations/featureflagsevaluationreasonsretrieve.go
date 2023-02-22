package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FeatureFlagsEvaluationReasonsRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsEvaluationReasonsRetrieveRequest struct {
	PathParams FeatureFlagsEvaluationReasonsRetrievePathParams
}

type FeatureFlagsEvaluationReasonsRetrieveResponse struct {
	ContentType string
	FeatureFlag *shared.FeatureFlag
	StatusCode  int
}
