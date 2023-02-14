package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type DashboardsPartialUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type DashboardsPartialUpdateRequests struct {
	PatchedDashboard  *shared.PatchedDashboardInput `request:"mediaType=application/json"`
	PatchedDashboard1 *shared.PatchedDashboardInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedDashboard2 *shared.PatchedDashboardInput `request:"mediaType=multipart/form-data"`
}

type DashboardsPartialUpdateRequest struct {
	PathParams DashboardsPartialUpdatePathParams
	Request    *DashboardsPartialUpdateRequests
}

type DashboardsPartialUpdateResponse struct {
	ContentType string
	Dashboard   *shared.DashboardOutput
	StatusCode  int64
}
