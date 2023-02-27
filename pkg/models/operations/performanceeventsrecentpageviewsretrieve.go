package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PerformanceEventsRecentPageviewsRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PerformanceEventsRecentPageviewsRetrieveRequest struct {
	PathParams PerformanceEventsRecentPageviewsRetrievePathParams
}

type PerformanceEventsRecentPageviewsRetrieveResponse struct {
	ContentType      string
	PerformanceEvent *shared.PerformanceEvent
	StatusCode       int
}
