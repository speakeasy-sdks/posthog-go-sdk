package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ExportsContentRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ExportsContentRetrieveRequest struct {
	PathParams ExportsContentRetrievePathParams
}

type ExportsContentRetrieveResponse struct {
	ContentType   string
	ExportedAsset *shared.ExportedAsset
	StatusCode    int64
}
