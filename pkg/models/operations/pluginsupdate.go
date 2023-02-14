package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginsUpdatePathParams struct {
	ID                         int64  `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type PluginsUpdateRequests struct {
	Plugin  *shared.PluginInput `request:"mediaType=application/json"`
	Plugin1 *shared.PluginInput `request:"mediaType=application/x-www-form-urlencoded"`
	Plugin2 *shared.PluginInput `request:"mediaType=multipart/form-data"`
}

type PluginsUpdateRequest struct {
	PathParams PluginsUpdatePathParams
	Request    *PluginsUpdateRequests
}

type PluginsUpdateResponse struct {
	ContentType string
	Plugin      *shared.Plugin
	StatusCode  int64
}
