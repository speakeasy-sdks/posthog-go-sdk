package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginConfigsRearrangePartialUpdatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PluginConfigsRearrangePartialUpdateRequests struct {
	PatchedPluginConfig  *shared.PatchedPluginConfigInput `request:"mediaType=application/json"`
	PatchedPluginConfig1 *shared.PatchedPluginConfigInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedPluginConfig2 *shared.PatchedPluginConfigInput `request:"mediaType=multipart/form-data"`
}

type PluginConfigsRearrangePartialUpdateRequest struct {
	PathParams PluginConfigsRearrangePartialUpdatePathParams
	Request    *PluginConfigsRearrangePartialUpdateRequests
}

type PluginConfigsRearrangePartialUpdateResponse struct {
	ContentType  string
	PluginConfig *shared.PluginConfig
	StatusCode   int
}
