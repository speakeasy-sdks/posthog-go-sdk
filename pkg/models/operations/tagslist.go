package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type TagsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type TagsListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type TagsListRequest struct {
	PathParams  TagsListPathParams
	QueryParams TagsListQueryParams
}

type TagsListResponse struct {
	ContentType             string
	PaginatedTaggedItemList *shared.PaginatedTaggedItemList
	StatusCode              int
}
