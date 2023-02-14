package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginsCreatePathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type PluginsCreateRequests struct {
	Plugin  *shared.PluginInput `request:"mediaType=application/json"`
	Plugin1 *shared.PluginInput `request:"mediaType=application/x-www-form-urlencoded"`
	Plugin2 *shared.PluginInput `request:"mediaType=multipart/form-data"`
}

type PluginsCreateRequest struct {
	PathParams PluginsCreatePathParams
	Request    *PluginsCreateRequests
}

type PluginsCreateResponse struct {
	ContentType string
	Plugin      *shared.Plugin
	StatusCode  int64
}
