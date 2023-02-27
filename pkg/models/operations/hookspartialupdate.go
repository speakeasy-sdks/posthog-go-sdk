package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type HooksPartialUpdatePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type HooksPartialUpdateRequests struct {
	PatchedHook  *shared.PatchedHookInput `request:"mediaType=application/json"`
	PatchedHook1 *shared.PatchedHookInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedHook2 *shared.PatchedHookInput `request:"mediaType=multipart/form-data"`
}

type HooksPartialUpdateRequest struct {
	PathParams HooksPartialUpdatePathParams
	Request    *HooksPartialUpdateRequests
}

type HooksPartialUpdateResponse struct {
	ContentType string
	Hook        *shared.Hook
	StatusCode  int
}
