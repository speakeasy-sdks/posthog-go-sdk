package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ResetTokenPartialUpdatePathParams struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type ResetTokenPartialUpdateRequests struct {
	PatchedTeam  *shared.PatchedTeamInput `request:"mediaType=application/json"`
	PatchedTeam1 *shared.PatchedTeamInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedTeam2 *shared.PatchedTeamInput `request:"mediaType=multipart/form-data"`
}

type ResetTokenPartialUpdateRequest struct {
	PathParams ResetTokenPartialUpdatePathParams
	Request    *ResetTokenPartialUpdateRequests
}

type ResetTokenPartialUpdateResponse struct {
	ContentType string
	StatusCode  int64
	Team        *shared.Team
}
