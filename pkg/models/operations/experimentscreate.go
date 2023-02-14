package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ExperimentsCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ExperimentsCreateRequests struct {
	Experiment  *shared.ExperimentInput `request:"mediaType=application/json"`
	Experiment1 *shared.ExperimentInput `request:"mediaType=application/x-www-form-urlencoded"`
	Experiment2 *shared.ExperimentInput `request:"mediaType=multipart/form-data"`
}

type ExperimentsCreateRequest struct {
	PathParams ExperimentsCreatePathParams
	Request    ExperimentsCreateRequests
}

type ExperimentsCreateResponse struct {
	ContentType string
	Experiment  *shared.Experiment
	StatusCode  int64
}
