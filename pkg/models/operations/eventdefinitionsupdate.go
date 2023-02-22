package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type EventDefinitionsUpdatePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type EventDefinitionsUpdateRequests struct {
	EventDefinition  *shared.EventDefinitionInput `request:"mediaType=application/json"`
	EventDefinition1 *shared.EventDefinitionInput `request:"mediaType=application/x-www-form-urlencoded"`
	EventDefinition2 *shared.EventDefinitionInput `request:"mediaType=multipart/form-data"`
}

type EventDefinitionsUpdateRequest struct {
	PathParams EventDefinitionsUpdatePathParams
	Request    EventDefinitionsUpdateRequests
}

type EventDefinitionsUpdateResponse struct {
	ContentType     string
	EventDefinition *shared.EventDefinition
	StatusCode      int
}
