package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type CreateRequests struct {
	Team  *shared.TeamInput `request:"mediaType=application/json"`
	Team1 *shared.TeamInput `request:"mediaType=application/x-www-form-urlencoded"`
	Team2 *shared.TeamInput `request:"mediaType=multipart/form-data"`
}

type CreateRequest struct {
	Request *CreateRequests
}

type CreateResponse struct {
	ContentType string
	StatusCode  int
	Team        *shared.Team
}
