package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginConfigsLogsListPathParams struct {
	ParentLookupPluginConfigID string `pathParam:"style=simple,explode=false,name=parent_lookup_plugin_config_id"`
	ProjectID                  string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PluginConfigsLogsListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type PluginConfigsLogsListRequest struct {
	PathParams  PluginConfigsLogsListPathParams
	QueryParams PluginConfigsLogsListQueryParams
}

type PluginConfigsLogsListResponse struct {
	ContentType                 string
	PaginatedPluginLogEntryList *shared.PaginatedPluginLogEntryList
	StatusCode                  int64
}
