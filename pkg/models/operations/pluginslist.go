package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginsListPathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type PluginsListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type PluginsListRequest struct {
	PathParams  PluginsListPathParams
	QueryParams PluginsListQueryParams
}

type PluginsListResponse struct {
	ContentType         string
	PaginatedPluginList *shared.PaginatedPluginList
	StatusCode          int
}
