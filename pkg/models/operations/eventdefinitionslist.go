package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type EventDefinitionsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type EventDefinitionsListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type EventDefinitionsListRequest struct {
	PathParams  EventDefinitionsListPathParams
	QueryParams EventDefinitionsListQueryParams
}

type EventDefinitionsListResponse struct {
	ContentType                  string
	PaginatedEventDefinitionList *shared.PaginatedEventDefinitionList
	StatusCode                   int64
}
