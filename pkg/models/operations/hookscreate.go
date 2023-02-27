package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type HooksCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type HooksCreateRequests struct {
	Hook  *shared.HookInput `request:"mediaType=application/json"`
	Hook1 *shared.HookInput `request:"mediaType=application/x-www-form-urlencoded"`
	Hook2 *shared.HookInput `request:"mediaType=multipart/form-data"`
}

type HooksCreateRequest struct {
	PathParams HooksCreatePathParams
	Request    HooksCreateRequests
}

type HooksCreateResponse struct {
	ContentType string
	Hook        *shared.Hook
	StatusCode  int
}
