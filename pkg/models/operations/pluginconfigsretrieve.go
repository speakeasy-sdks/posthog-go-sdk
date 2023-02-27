package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginConfigsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PluginConfigsRetrieveRequest struct {
	PathParams PluginConfigsRetrievePathParams
}

type PluginConfigsRetrieveResponse struct {
	ContentType  string
	PluginConfig *shared.PluginConfig
	StatusCode   int
}
