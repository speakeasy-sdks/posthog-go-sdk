package operations

type SessionRecordingPlaylistsRecordingsDestroyPathParams struct {
	ProjectID          string `pathParam:"style=simple,explode=false,name=project_id"`
	SessionRecordingID string `pathParam:"style=simple,explode=false,name=session_recording_id"`
	ShortID            string `pathParam:"style=simple,explode=false,name=short_id"`
}

type SessionRecordingPlaylistsRecordingsDestroyRequest struct {
	PathParams SessionRecordingPlaylistsRecordingsDestroyPathParams
}

type SessionRecordingPlaylistsRecordingsDestroyResponse struct {
	ContentType string
	StatusCode  int64
}
