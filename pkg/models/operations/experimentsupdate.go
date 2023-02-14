package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ExperimentsUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ExperimentsUpdateRequests struct {
	Experiment  *shared.ExperimentInput `request:"mediaType=application/json"`
	Experiment1 *shared.ExperimentInput `request:"mediaType=application/x-www-form-urlencoded"`
	Experiment2 *shared.ExperimentInput `request:"mediaType=multipart/form-data"`
}

type ExperimentsUpdateRequest struct {
	PathParams ExperimentsUpdatePathParams
	Request    ExperimentsUpdateRequests
}

type ExperimentsUpdateResponse struct {
	ContentType string
	Experiment  *shared.Experiment
	StatusCode  int64
}
