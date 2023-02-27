package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type IntegrationsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type IntegrationsListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type IntegrationsListRequest struct {
	PathParams  IntegrationsListPathParams
	QueryParams IntegrationsListQueryParams
}

type IntegrationsListResponse struct {
	ContentType              string
	PaginatedIntegrationList *shared.PaginatedIntegrationList
	StatusCode               int
}
