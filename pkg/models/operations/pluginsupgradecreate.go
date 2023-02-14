package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginsUpgradeCreatePathParams struct {
	ID                         int64  `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type PluginsUpgradeCreateRequests struct {
	Plugin  *shared.PluginInput `request:"mediaType=application/json"`
	Plugin1 *shared.PluginInput `request:"mediaType=application/x-www-form-urlencoded"`
	Plugin2 *shared.PluginInput `request:"mediaType=multipart/form-data"`
}

type PluginsUpgradeCreateRequest struct {
	PathParams PluginsUpgradeCreatePathParams
	Request    *PluginsUpgradeCreateRequests
}

type PluginsUpgradeCreateResponse struct {
	ContentType string
	Plugin      *shared.Plugin
	StatusCode  int64
}
