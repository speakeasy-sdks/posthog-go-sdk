package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type HooksRetrievePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type HooksRetrieveRequest struct {
	PathParams HooksRetrievePathParams
}

type HooksRetrieveResponse struct {
	ContentType string
	Hook        *shared.Hook
	StatusCode  int64
}
