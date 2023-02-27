package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type DashboardsCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type DashboardsCreateRequests struct {
	Dashboard  *shared.DashboardInput `request:"mediaType=application/json"`
	Dashboard1 *shared.DashboardInput `request:"mediaType=application/x-www-form-urlencoded"`
	Dashboard2 *shared.DashboardInput `request:"mediaType=multipart/form-data"`
}

type DashboardsCreateRequest struct {
	PathParams DashboardsCreatePathParams
	Request    *DashboardsCreateRequests
}

type DashboardsCreateResponse struct {
	ContentType string
	Dashboard   *shared.DashboardOutput
	StatusCode  int
}
