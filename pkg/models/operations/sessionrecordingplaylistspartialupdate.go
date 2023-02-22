package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type SessionRecordingPlaylistsPartialUpdatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
	ShortID   string `pathParam:"style=simple,explode=false,name=short_id"`
}

type SessionRecordingPlaylistsPartialUpdateRequests struct {
	PatchedSessionRecordingPlaylist  *shared.PatchedSessionRecordingPlaylistInput `request:"mediaType=application/json"`
	PatchedSessionRecordingPlaylist1 *shared.PatchedSessionRecordingPlaylistInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedSessionRecordingPlaylist2 *shared.PatchedSessionRecordingPlaylistInput `request:"mediaType=multipart/form-data"`
}

type SessionRecordingPlaylistsPartialUpdateRequest struct {
	PathParams SessionRecordingPlaylistsPartialUpdatePathParams
	Request    *SessionRecordingPlaylistsPartialUpdateRequests
}

type SessionRecordingPlaylistsPartialUpdateResponse struct {
	ContentType              string
	SessionRecordingPlaylist *shared.SessionRecordingPlaylist
	StatusCode               int
}
