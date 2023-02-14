package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type CohortsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type CohortsListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type CohortsListRequest struct {
	PathParams  CohortsListPathParams
	QueryParams CohortsListQueryParams
}

type CohortsListResponse struct {
	ContentType         string
	PaginatedCohortList *shared.PaginatedCohortList
	StatusCode          int64
}
