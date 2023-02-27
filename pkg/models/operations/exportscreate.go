package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ExportsCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ExportsCreateRequests struct {
	ExportedAsset  *shared.ExportedAssetInput `request:"mediaType=application/json"`
	ExportedAsset1 *shared.ExportedAssetInput `request:"mediaType=application/x-www-form-urlencoded"`
	ExportedAsset2 *shared.ExportedAssetInput `request:"mediaType=multipart/form-data"`
}

type ExportsCreateRequest struct {
	PathParams ExportsCreatePathParams
	Request    ExportsCreateRequests
}

type ExportsCreateResponse struct {
	ContentType   string
	ExportedAsset *shared.ExportedAsset
	StatusCode    int
}
