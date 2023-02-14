package operations

type SessionRecordingPlaylistsDestroyPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
	ShortID   string `pathParam:"style=simple,explode=false,name=short_id"`
}

type SessionRecordingPlaylistsDestroyRequest struct {
	PathParams SessionRecordingPlaylistsDestroyPathParams
}

type SessionRecordingPlaylistsDestroyResponse struct {
	ContentType string
	StatusCode  int64
}
