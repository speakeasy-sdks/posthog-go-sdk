package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginConfigsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PluginConfigsListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type PluginConfigsListRequest struct {
	PathParams  PluginConfigsListPathParams
	QueryParams PluginConfigsListQueryParams
}

type PluginConfigsListResponse struct {
	ContentType               string
	PaginatedPluginConfigList *shared.PaginatedPluginConfigList
	StatusCode                int64
}
