package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type SubscriptionsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type SubscriptionsRetrieveRequest struct {
	PathParams SubscriptionsRetrievePathParams
}

type SubscriptionsRetrieveResponse struct {
	ContentType  string
	StatusCode   int64
	Subscription *shared.Subscription
}
