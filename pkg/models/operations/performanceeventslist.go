package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PerformanceEventsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PerformanceEventsListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type PerformanceEventsListRequest struct {
	PathParams  PerformanceEventsListPathParams
	QueryParams PerformanceEventsListQueryParams
}

type PerformanceEventsListResponse struct {
	ContentType                   string
	PaginatedPerformanceEventList *shared.PaginatedPerformanceEventList
	StatusCode                    int
}
