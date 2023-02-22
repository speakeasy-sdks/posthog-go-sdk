package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type UpdatePathParams struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateRequests struct {
	Team  *shared.TeamInput `request:"mediaType=application/json"`
	Team1 *shared.TeamInput `request:"mediaType=application/x-www-form-urlencoded"`
	Team2 *shared.TeamInput `request:"mediaType=multipart/form-data"`
}

type UpdateRequest struct {
	PathParams UpdatePathParams
	Request    *UpdateRequests
}

type UpdateResponse struct {
	ContentType string
	StatusCode  int
	Team        *shared.Team
}
