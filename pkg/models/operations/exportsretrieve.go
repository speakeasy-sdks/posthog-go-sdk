package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ExportsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ExportsRetrieveRequest struct {
	PathParams ExportsRetrievePathParams
}

type ExportsRetrieveResponse struct {
	ContentType   string
	ExportedAsset *shared.ExportedAsset
	StatusCode    int
}
