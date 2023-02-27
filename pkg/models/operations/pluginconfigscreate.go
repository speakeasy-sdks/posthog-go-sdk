package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginConfigsCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PluginConfigsCreateRequests struct {
	PluginConfig  *shared.PluginConfigInput `request:"mediaType=application/json"`
	PluginConfig1 *shared.PluginConfigInput `request:"mediaType=application/x-www-form-urlencoded"`
	PluginConfig2 *shared.PluginConfigInput `request:"mediaType=multipart/form-data"`
}

type PluginConfigsCreateRequest struct {
	PathParams PluginConfigsCreatePathParams
	Request    PluginConfigsCreateRequests
}

type PluginConfigsCreateResponse struct {
	ContentType  string
	PluginConfig *shared.PluginConfig
	StatusCode   int
}
