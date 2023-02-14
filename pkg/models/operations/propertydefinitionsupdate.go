package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PropertyDefinitionsUpdatePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PropertyDefinitionsUpdateRequests struct {
	PropertyDefinition  *shared.PropertyDefinitionInput `request:"mediaType=application/json"`
	PropertyDefinition1 *shared.PropertyDefinitionInput `request:"mediaType=application/x-www-form-urlencoded"`
	PropertyDefinition2 *shared.PropertyDefinitionInput `request:"mediaType=multipart/form-data"`
}

type PropertyDefinitionsUpdateRequest struct {
	PathParams PropertyDefinitionsUpdatePathParams
	Request    PropertyDefinitionsUpdateRequests
}

type PropertyDefinitionsUpdateResponse struct {
	ContentType        string
	PropertyDefinition *shared.PropertyDefinition
	StatusCode         int64
}
