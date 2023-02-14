package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type DashboardsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type DashboardsListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type DashboardsListRequest struct {
	PathParams  DashboardsListPathParams
	QueryParams DashboardsListQueryParams
}

type DashboardsListResponse struct {
	ContentType            string
	PaginatedDashboardList *shared.PaginatedDashboardList
	StatusCode             int64
}
