package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginConfigsPartialUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PluginConfigsPartialUpdateRequests struct {
	PatchedPluginConfig  *shared.PatchedPluginConfigInput `request:"mediaType=application/json"`
	PatchedPluginConfig1 *shared.PatchedPluginConfigInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedPluginConfig2 *shared.PatchedPluginConfigInput `request:"mediaType=multipart/form-data"`
}

type PluginConfigsPartialUpdateRequest struct {
	PathParams PluginConfigsPartialUpdatePathParams
	Request    *PluginConfigsPartialUpdateRequests
}

type PluginConfigsPartialUpdateResponse struct {
	ContentType  string
	PluginConfig *shared.PluginConfig
	StatusCode   int
}
