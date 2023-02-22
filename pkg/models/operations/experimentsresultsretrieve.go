package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ExperimentsResultsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ExperimentsResultsRetrieveRequest struct {
	PathParams ExperimentsResultsRetrievePathParams
}

type ExperimentsResultsRetrieveResponse struct {
	ContentType string
	Experiment  *shared.Experiment
	StatusCode  int
}
