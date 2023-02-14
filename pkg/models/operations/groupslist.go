package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type GroupsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type GroupsListQueryParams struct {
	Cursor         *int64 `queryParam:"style=form,explode=true,name=cursor"`
	GroupTypeIndex int64  `queryParam:"style=form,explode=true,name=group_type_index"`
}

type GroupsListRequest struct {
	PathParams  GroupsListPathParams
	QueryParams GroupsListQueryParams
}

type GroupsListResponse struct {
	ContentType        string
	PaginatedGroupList *shared.PaginatedGroupList
	StatusCode         int64
}
