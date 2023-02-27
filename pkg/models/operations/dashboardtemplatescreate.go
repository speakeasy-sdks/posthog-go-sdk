package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type DashboardTemplatesCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type DashboardTemplatesCreateRequests struct {
	DashboardTemplate  *shared.DashboardTemplate `request:"mediaType=application/json"`
	DashboardTemplate1 *shared.DashboardTemplate `request:"mediaType=application/x-www-form-urlencoded"`
	DashboardTemplate2 *shared.DashboardTemplate `request:"mediaType=multipart/form-data"`
}

type DashboardTemplatesCreateRequest struct {
	PathParams DashboardTemplatesCreatePathParams
	Request    DashboardTemplatesCreateRequests
}

type DashboardTemplatesCreateResponse struct {
	ContentType string
	StatusCode  int
}
