package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type IntegrationsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type IntegrationsRetrieveRequest struct {
	PathParams IntegrationsRetrievePathParams
}

type IntegrationsRetrieveResponse struct {
	ContentType string
	Integration *shared.Integration
	StatusCode  int64
}
