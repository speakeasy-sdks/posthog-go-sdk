package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type DomainsListPathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type DomainsListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type DomainsListRequest struct {
	PathParams  DomainsListPathParams
	QueryParams DomainsListQueryParams
}

type DomainsListResponse struct {
	ContentType                     string
	PaginatedOrganizationDomainList *shared.PaginatedOrganizationDomainList
	StatusCode                      int64
}
