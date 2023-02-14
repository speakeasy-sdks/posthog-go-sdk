package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PropertyDefinitionsRetrievePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PropertyDefinitionsRetrieveRequest struct {
	PathParams PropertyDefinitionsRetrievePathParams
}

type PropertyDefinitionsRetrieveResponse struct {
	ContentType        string
	PropertyDefinition *shared.PropertyDefinition
	StatusCode         int64
}
