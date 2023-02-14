package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PropertyDefinitionsPartialUpdatePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PropertyDefinitionsPartialUpdateRequests struct {
	PatchedPropertyDefinition  *shared.PatchedPropertyDefinitionInput `request:"mediaType=application/json"`
	PatchedPropertyDefinition1 *shared.PatchedPropertyDefinitionInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedPropertyDefinition2 *shared.PatchedPropertyDefinitionInput `request:"mediaType=multipart/form-data"`
}

type PropertyDefinitionsPartialUpdateRequest struct {
	PathParams PropertyDefinitionsPartialUpdatePathParams
	Request    *PropertyDefinitionsPartialUpdateRequests
}

type PropertyDefinitionsPartialUpdateResponse struct {
	ContentType        string
	PropertyDefinition *shared.PropertyDefinition
	StatusCode         int64
}
