package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type CohortsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type CohortsRetrieveRequest struct {
	PathParams CohortsRetrievePathParams
}

type CohortsRetrieveResponse struct {
	Cohort      *shared.Cohort
	ContentType string
	StatusCode  int
}
