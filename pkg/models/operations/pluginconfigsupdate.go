package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginConfigsUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PluginConfigsUpdateRequests struct {
	PluginConfig  *shared.PluginConfigInput `request:"mediaType=application/json"`
	PluginConfig1 *shared.PluginConfigInput `request:"mediaType=application/x-www-form-urlencoded"`
	PluginConfig2 *shared.PluginConfigInput `request:"mediaType=multipart/form-data"`
}

type PluginConfigsUpdateRequest struct {
	PathParams PluginConfigsUpdatePathParams
	Request    PluginConfigsUpdateRequests
}

type PluginConfigsUpdateResponse struct {
	ContentType  string
	PluginConfig *shared.PluginConfig
	StatusCode   int
}
