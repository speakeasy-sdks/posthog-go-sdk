package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type SubscriptionsCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type SubscriptionsCreateRequests struct {
	Subscription  *shared.SubscriptionInput `request:"mediaType=application/json"`
	Subscription1 *shared.SubscriptionInput `request:"mediaType=application/x-www-form-urlencoded"`
	Subscription2 *shared.SubscriptionInput `request:"mediaType=multipart/form-data"`
}

type SubscriptionsCreateRequest struct {
	PathParams SubscriptionsCreatePathParams
	Request    SubscriptionsCreateRequests
}

type SubscriptionsCreateResponse struct {
	ContentType  string
	StatusCode   int64
	Subscription *shared.Subscription
}
