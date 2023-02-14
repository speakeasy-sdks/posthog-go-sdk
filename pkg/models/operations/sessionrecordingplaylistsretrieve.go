package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type SessionRecordingPlaylistsRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
	ShortID   string `pathParam:"style=simple,explode=false,name=short_id"`
}

type SessionRecordingPlaylistsRetrieveRequest struct {
	PathParams SessionRecordingPlaylistsRetrievePathParams
}

type SessionRecordingPlaylistsRetrieveResponse struct {
	ContentType              string
	SessionRecordingPlaylist *shared.SessionRecordingPlaylist
	StatusCode               int64
}
