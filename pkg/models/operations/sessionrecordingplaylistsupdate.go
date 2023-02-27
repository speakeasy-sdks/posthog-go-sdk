package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type SessionRecordingPlaylistsUpdatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
	ShortID   string `pathParam:"style=simple,explode=false,name=short_id"`
}

type SessionRecordingPlaylistsUpdateRequests struct {
	SessionRecordingPlaylist  *shared.SessionRecordingPlaylistInput `request:"mediaType=application/json"`
	SessionRecordingPlaylist1 *shared.SessionRecordingPlaylistInput `request:"mediaType=application/x-www-form-urlencoded"`
	SessionRecordingPlaylist2 *shared.SessionRecordingPlaylistInput `request:"mediaType=multipart/form-data"`
}

type SessionRecordingPlaylistsUpdateRequest struct {
	PathParams SessionRecordingPlaylistsUpdatePathParams
	Request    *SessionRecordingPlaylistsUpdateRequests
}

type SessionRecordingPlaylistsUpdateResponse struct {
	ContentType              string
	SessionRecordingPlaylist *shared.SessionRecordingPlaylist
	StatusCode               int
}
