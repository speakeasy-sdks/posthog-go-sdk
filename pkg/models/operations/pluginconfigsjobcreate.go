package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginConfigsJobCreatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PluginConfigsJobCreateRequests struct {
	PluginConfig  *shared.PluginConfigInput `request:"mediaType=application/json"`
	PluginConfig1 *shared.PluginConfigInput `request:"mediaType=application/x-www-form-urlencoded"`
	PluginConfig2 *shared.PluginConfigInput `request:"mediaType=multipart/form-data"`
}

type PluginConfigsJobCreateRequest struct {
	PathParams PluginConfigsJobCreatePathParams
	Request    PluginConfigsJobCreateRequests
}

type PluginConfigsJobCreateResponse struct {
	ContentType  string
	PluginConfig *shared.PluginConfig
	StatusCode   int
}
