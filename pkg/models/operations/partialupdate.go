package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PartialUpdatePathParams struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type PartialUpdateRequests struct {
	PatchedTeam  *shared.PatchedTeamInput `request:"mediaType=application/json"`
	PatchedTeam1 *shared.PatchedTeamInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedTeam2 *shared.PatchedTeamInput `request:"mediaType=multipart/form-data"`
}

type PartialUpdateRequest struct {
	PathParams PartialUpdatePathParams
	Request    *PartialUpdateRequests
}

type PartialUpdateResponse struct {
	ContentType string
	StatusCode  int
	Team        *shared.Team
}
