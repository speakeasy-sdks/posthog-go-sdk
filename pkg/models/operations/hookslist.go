package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type HooksListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type HooksListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type HooksListRequest struct {
	PathParams  HooksListPathParams
	QueryParams HooksListQueryParams
}

type HooksListResponse struct {
	ContentType       string
	PaginatedHookList *shared.PaginatedHookList
	StatusCode        int
}
