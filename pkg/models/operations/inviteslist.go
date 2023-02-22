package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InvitesListPathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type InvitesListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type InvitesListRequest struct {
	PathParams  InvitesListPathParams
	QueryParams InvitesListQueryParams
}

type InvitesListResponse struct {
	ContentType                     string
	PaginatedOrganizationInviteList *shared.PaginatedOrganizationInviteList
	StatusCode                      int
}
