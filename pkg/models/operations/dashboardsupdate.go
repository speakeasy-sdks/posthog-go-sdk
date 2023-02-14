package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type DashboardsUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type DashboardsUpdateRequests struct {
	Dashboard  *shared.DashboardInput `request:"mediaType=application/json"`
	Dashboard1 *shared.DashboardInput `request:"mediaType=application/x-www-form-urlencoded"`
	Dashboard2 *shared.DashboardInput `request:"mediaType=multipart/form-data"`
}

type DashboardsUpdateRequest struct {
	PathParams DashboardsUpdatePathParams
	Request    *DashboardsUpdateRequests
}

type DashboardsUpdateResponse struct {
	ContentType string
	Dashboard   *shared.DashboardOutput
	StatusCode  int64
}
