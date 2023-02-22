package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type CohortsCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type CohortsCreateRequests struct {
	Cohort  *shared.CohortInput `request:"mediaType=application/json"`
	Cohort1 *shared.CohortInput `request:"mediaType=application/x-www-form-urlencoded"`
	Cohort2 *shared.CohortInput `request:"mediaType=multipart/form-data"`
}

type CohortsCreateRequest struct {
	PathParams CohortsCreatePathParams
	Request    *CohortsCreateRequests
}

type CohortsCreateResponse struct {
	Cohort      *shared.Cohort
	ContentType string
	StatusCode  int
}
