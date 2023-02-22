package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FeatureFlagsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type FeatureFlagsListRequest struct {
	PathParams  FeatureFlagsListPathParams
	QueryParams FeatureFlagsListQueryParams
}

type FeatureFlagsListResponse struct {
	ContentType              string
	PaginatedFeatureFlagList *shared.PaginatedFeatureFlagList
	StatusCode               int
}
