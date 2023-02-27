package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type SessionRecordingPlaylistsRecordingsCreatePathParams struct {
	ProjectID          string `pathParam:"style=simple,explode=false,name=project_id"`
	SessionRecordingID string `pathParam:"style=simple,explode=false,name=session_recording_id"`
	ShortID            string `pathParam:"style=simple,explode=false,name=short_id"`
}

type SessionRecordingPlaylistsRecordingsCreateRequests struct {
	SessionRecordingPlaylist  *shared.SessionRecordingPlaylistInput `request:"mediaType=application/json"`
	SessionRecordingPlaylist1 *shared.SessionRecordingPlaylistInput `request:"mediaType=application/x-www-form-urlencoded"`
	SessionRecordingPlaylist2 *shared.SessionRecordingPlaylistInput `request:"mediaType=multipart/form-data"`
}

type SessionRecordingPlaylistsRecordingsCreateRequest struct {
	PathParams SessionRecordingPlaylistsRecordingsCreatePathParams
	Request    *SessionRecordingPlaylistsRecordingsCreateRequests
}

type SessionRecordingPlaylistsRecordingsCreateResponse struct {
	ContentType              string
	SessionRecordingPlaylist *shared.SessionRecordingPlaylist
	StatusCode               int
}
