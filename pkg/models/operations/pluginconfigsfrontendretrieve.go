package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginConfigsFrontendRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PluginConfigsFrontendRetrieveRequest struct {
	PathParams PluginConfigsFrontendRetrievePathParams
}

type PluginConfigsFrontendRetrieveResponse struct {
	ContentType  string
	PluginConfig *shared.PluginConfig
	StatusCode   int
}
