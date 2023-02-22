package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ExperimentsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ExperimentsListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type ExperimentsListRequest struct {
	PathParams  ExperimentsListPathParams
	QueryParams ExperimentsListQueryParams
}

type ExperimentsListResponse struct {
	ContentType             string
	PaginatedExperimentList *shared.PaginatedExperimentList
	StatusCode              int
}
