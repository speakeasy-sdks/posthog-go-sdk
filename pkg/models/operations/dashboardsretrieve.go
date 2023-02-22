package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type DashboardsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type DashboardsRetrieveRequest struct {
	PathParams DashboardsRetrievePathParams
}

type DashboardsRetrieveResponse struct {
	ContentType string
	Dashboard   *shared.DashboardOutput
	StatusCode  int
}
