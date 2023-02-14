package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type HooksUpdatePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type HooksUpdateRequests struct {
	Hook  *shared.HookInput `request:"mediaType=application/json"`
	Hook1 *shared.HookInput `request:"mediaType=application/x-www-form-urlencoded"`
	Hook2 *shared.HookInput `request:"mediaType=multipart/form-data"`
}

type HooksUpdateRequest struct {
	PathParams HooksUpdatePathParams
	Request    HooksUpdateRequests
}

type HooksUpdateResponse struct {
	ContentType string
	Hook        *shared.Hook
	StatusCode  int64
}
