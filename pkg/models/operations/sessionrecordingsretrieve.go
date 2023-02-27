package operations

type SessionRecordingsRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type SessionRecordingsRetrieveRequest struct {
	PathParams SessionRecordingsRetrievePathParams
}

type SessionRecordingsRetrieveResponse struct {
	ContentType string
	StatusCode  int
}
