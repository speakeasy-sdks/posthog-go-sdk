package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type SubscriptionsUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type SubscriptionsUpdateRequests struct {
	Subscription  *shared.SubscriptionInput `request:"mediaType=application/json"`
	Subscription1 *shared.SubscriptionInput `request:"mediaType=application/x-www-form-urlencoded"`
	Subscription2 *shared.SubscriptionInput `request:"mediaType=multipart/form-data"`
}

type SubscriptionsUpdateRequest struct {
	PathParams SubscriptionsUpdatePathParams
	Request    SubscriptionsUpdateRequests
}

type SubscriptionsUpdateResponse struct {
	ContentType  string
	StatusCode   int64
	Subscription *shared.Subscription
}
