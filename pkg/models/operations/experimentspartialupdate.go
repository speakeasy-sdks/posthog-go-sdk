package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ExperimentsPartialUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ExperimentsPartialUpdateRequests struct {
	PatchedExperiment  *shared.PatchedExperimentInput `request:"mediaType=application/json"`
	PatchedExperiment1 *shared.PatchedExperimentInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedExperiment2 *shared.PatchedExperimentInput `request:"mediaType=multipart/form-data"`
}

type ExperimentsPartialUpdateRequest struct {
	PathParams ExperimentsPartialUpdatePathParams
	Request    *ExperimentsPartialUpdateRequests
}

type ExperimentsPartialUpdateResponse struct {
	ContentType string
	Experiment  *shared.Experiment
	StatusCode  int
}
