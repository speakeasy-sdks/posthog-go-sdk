package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ExperimentsSecondaryResultsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ExperimentsSecondaryResultsRetrieveRequest struct {
	PathParams ExperimentsSecondaryResultsRetrievePathParams
}

type ExperimentsSecondaryResultsRetrieveResponse struct {
	ContentType string
	Experiment  *shared.Experiment
	StatusCode  int
}
