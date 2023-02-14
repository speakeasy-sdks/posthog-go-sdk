package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type CohortsUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type CohortsUpdateRequests struct {
	Cohort  *shared.CohortInput `request:"mediaType=application/json"`
	Cohort1 *shared.CohortInput `request:"mediaType=application/x-www-form-urlencoded"`
	Cohort2 *shared.CohortInput `request:"mediaType=multipart/form-data"`
}

type CohortsUpdateRequest struct {
	PathParams CohortsUpdatePathParams
	Request    *CohortsUpdateRequests
}

type CohortsUpdateResponse struct {
	Cohort      *shared.Cohort
	ContentType string
	StatusCode  int64
}
