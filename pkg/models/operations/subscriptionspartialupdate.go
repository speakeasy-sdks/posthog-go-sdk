package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type SubscriptionsPartialUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type SubscriptionsPartialUpdateRequests struct {
	PatchedSubscription  *shared.PatchedSubscriptionInput `request:"mediaType=application/json"`
	PatchedSubscription1 *shared.PatchedSubscriptionInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedSubscription2 *shared.PatchedSubscriptionInput `request:"mediaType=multipart/form-data"`
}

type SubscriptionsPartialUpdateRequest struct {
	PathParams SubscriptionsPartialUpdatePathParams
	Request    *SubscriptionsPartialUpdateRequests
}

type SubscriptionsPartialUpdateResponse struct {
	ContentType  string
	StatusCode   int
	Subscription *shared.Subscription
}
