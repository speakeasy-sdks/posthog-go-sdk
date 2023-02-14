package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ExperimentsRequiresFlagImplementationRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ExperimentsRequiresFlagImplementationRetrieveRequest struct {
	PathParams ExperimentsRequiresFlagImplementationRetrievePathParams
}

type ExperimentsRequiresFlagImplementationRetrieveResponse struct {
	ContentType string
	Experiment  *shared.Experiment
	StatusCode  int64
}
