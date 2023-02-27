package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginsPartialUpdatePathParams struct {
	ID                         int64  `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type PluginsPartialUpdateRequests struct {
	PatchedPlugin  *shared.PatchedPluginInput `request:"mediaType=application/json"`
	PatchedPlugin1 *shared.PatchedPluginInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedPlugin2 *shared.PatchedPluginInput `request:"mediaType=multipart/form-data"`
}

type PluginsPartialUpdateRequest struct {
	PathParams PluginsPartialUpdatePathParams
	Request    *PluginsPartialUpdateRequests
}

type PluginsPartialUpdateResponse struct {
	ContentType string
	Plugin      *shared.Plugin
	StatusCode  int
}
