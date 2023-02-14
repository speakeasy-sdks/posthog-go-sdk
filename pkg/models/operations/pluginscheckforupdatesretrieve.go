package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginsCheckForUpdatesRetrievePathParams struct {
	ID                         int64  `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type PluginsCheckForUpdatesRetrieveRequest struct {
	PathParams PluginsCheckForUpdatesRetrievePathParams
}

type PluginsCheckForUpdatesRetrieveResponse struct {
	ContentType string
	Plugin      *shared.Plugin
	StatusCode  int64
}
