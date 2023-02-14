package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type SessionRecordingPlaylistsCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type SessionRecordingPlaylistsCreateRequests struct {
	SessionRecordingPlaylist  *shared.SessionRecordingPlaylistInput `request:"mediaType=application/json"`
	SessionRecordingPlaylist1 *shared.SessionRecordingPlaylistInput `request:"mediaType=application/x-www-form-urlencoded"`
	SessionRecordingPlaylist2 *shared.SessionRecordingPlaylistInput `request:"mediaType=multipart/form-data"`
}

type SessionRecordingPlaylistsCreateRequest struct {
	PathParams SessionRecordingPlaylistsCreatePathParams
	Request    *SessionRecordingPlaylistsCreateRequests
}

type SessionRecordingPlaylistsCreateResponse struct {
	ContentType              string
	SessionRecordingPlaylist *shared.SessionRecordingPlaylist
	StatusCode               int64
}
