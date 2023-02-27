package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type RolesListPathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type RolesListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type RolesListRequest struct {
	PathParams  RolesListPathParams
	QueryParams RolesListQueryParams
}

type RolesListResponse struct {
	ContentType       string
	PaginatedRoleList *shared.PaginatedRoleList
	StatusCode        int
}
