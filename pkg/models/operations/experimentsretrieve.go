package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ExperimentsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ExperimentsRetrieveRequest struct {
	PathParams ExperimentsRetrievePathParams
}

type ExperimentsRetrieveResponse struct {
	ContentType string
	Experiment  *shared.Experiment
	StatusCode  int64
}
