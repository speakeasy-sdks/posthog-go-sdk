package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type GroupsPropertyDefinitionsRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type GroupsPropertyDefinitionsRetrieveRequest struct {
	PathParams GroupsPropertyDefinitionsRetrievePathParams
}

type GroupsPropertyDefinitionsRetrieveResponse struct {
	ContentType string
	Group       *shared.Group
	StatusCode  int
}
