package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type SessionRecordingPlaylistsRecordingsRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
	ShortID   string `pathParam:"style=simple,explode=false,name=short_id"`
}

type SessionRecordingPlaylistsRecordingsRetrieveRequest struct {
	PathParams SessionRecordingPlaylistsRecordingsRetrievePathParams
}

type SessionRecordingPlaylistsRecordingsRetrieveResponse struct {
	ContentType              string
	SessionRecordingPlaylist *shared.SessionRecordingPlaylist
	StatusCode               int
}
