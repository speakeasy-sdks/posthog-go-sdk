package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginsUpdateSourcePartialUpdatePathParams struct {
	ID                         int64  `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type PluginsUpdateSourcePartialUpdateRequests struct {
	PatchedPlugin  *shared.PatchedPluginInput `request:"mediaType=application/json"`
	PatchedPlugin1 *shared.PatchedPluginInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedPlugin2 *shared.PatchedPluginInput `request:"mediaType=multipart/form-data"`
}

type PluginsUpdateSourcePartialUpdateRequest struct {
	PathParams PluginsUpdateSourcePartialUpdatePathParams
	Request    *PluginsUpdateSourcePartialUpdateRequests
}

type PluginsUpdateSourcePartialUpdateResponse struct {
	ContentType string
	Plugin      *shared.Plugin
	StatusCode  int
}
