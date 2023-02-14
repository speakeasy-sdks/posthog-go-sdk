package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type MembersListPathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type MembersListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type MembersListRequest struct {
	PathParams  MembersListPathParams
	QueryParams MembersListQueryParams
}

type MembersListResponse struct {
	ContentType                     string
	PaginatedOrganizationMemberList *shared.PaginatedOrganizationMemberList
	StatusCode                      int64
}
