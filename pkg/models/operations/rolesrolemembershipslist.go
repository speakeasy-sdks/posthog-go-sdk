package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type RolesRoleMembershipsListPathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
	ParentLookupRoleID         string `pathParam:"style=simple,explode=false,name=parent_lookup_role_id"`
}

type RolesRoleMembershipsListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type RolesRoleMembershipsListRequest struct {
	PathParams  RolesRoleMembershipsListPathParams
	QueryParams RolesRoleMembershipsListQueryParams
}

type RolesRoleMembershipsListResponse struct {
	ContentType                 string
	PaginatedRoleMembershipList *shared.PaginatedRoleMembershipList
	StatusCode                  int64
}
