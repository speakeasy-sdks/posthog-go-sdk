package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type CohortsPartialUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type CohortsPartialUpdateRequests struct {
	PatchedCohort  *shared.PatchedCohortInput `request:"mediaType=application/json"`
	PatchedCohort1 *shared.PatchedCohortInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedCohort2 *shared.PatchedCohortInput `request:"mediaType=multipart/form-data"`
}

type CohortsPartialUpdateRequest struct {
	PathParams CohortsPartialUpdatePathParams
	Request    *CohortsPartialUpdateRequests
}

type CohortsPartialUpdateResponse struct {
	Cohort      *shared.Cohort
	ContentType string
	StatusCode  int64
}
