package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type EventDefinitionsPartialUpdatePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type EventDefinitionsPartialUpdateRequests struct {
	PatchedEventDefinition  *shared.PatchedEventDefinitionInput `request:"mediaType=application/json"`
	PatchedEventDefinition1 *shared.PatchedEventDefinitionInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedEventDefinition2 *shared.PatchedEventDefinitionInput `request:"mediaType=multipart/form-data"`
}

type EventDefinitionsPartialUpdateRequest struct {
	PathParams EventDefinitionsPartialUpdatePathParams
	Request    *EventDefinitionsPartialUpdateRequests
}

type EventDefinitionsPartialUpdateResponse struct {
	ContentType     string
	EventDefinition *shared.EventDefinition
	StatusCode      int64
}
