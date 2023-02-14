package operations

type SessionRecordingsSnapshotsRetrievePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type SessionRecordingsSnapshotsRetrieveRequest struct {
	PathParams SessionRecordingsSnapshotsRetrievePathParams
}

type SessionRecordingsSnapshotsRetrieveResponse struct {
	ContentType string
	StatusCode  int64
}
