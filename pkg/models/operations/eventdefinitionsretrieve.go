package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type EventDefinitionsRetrievePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type EventDefinitionsRetrieveRequest struct {
	PathParams EventDefinitionsRetrievePathParams
}

type EventDefinitionsRetrieveResponse struct {
	ContentType     string
	EventDefinition *shared.EventDefinition
	StatusCode      int64
}
