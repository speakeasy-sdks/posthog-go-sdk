package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type ListRequest struct {
	QueryParams ListQueryParams
}

type ListResponse struct {
	ContentType            string
	PaginatedTeamBasicList *shared.PaginatedTeamBasicList
	StatusCode             int64
}
