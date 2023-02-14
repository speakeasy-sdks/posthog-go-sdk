package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type SubscriptionsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type SubscriptionsListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type SubscriptionsListRequest struct {
	PathParams  SubscriptionsListPathParams
	QueryParams SubscriptionsListQueryParams
}

type SubscriptionsListResponse struct {
	ContentType               string
	PaginatedSubscriptionList *shared.PaginatedSubscriptionList
	StatusCode                int64
}
