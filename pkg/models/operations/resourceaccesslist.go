package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ResourceAccessListPathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type ResourceAccessListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type ResourceAccessListRequest struct {
	PathParams  ResourceAccessListPathParams
	QueryParams ResourceAccessListQueryParams
}

type ResourceAccessListResponse struct {
	ContentType                             string
	PaginatedOrganizationResourceAccessList *shared.PaginatedOrganizationResourceAccessList
	StatusCode                              int64
}
