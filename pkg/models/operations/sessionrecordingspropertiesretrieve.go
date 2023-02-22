package operations

type SessionRecordingsPropertiesRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type SessionRecordingsPropertiesRetrieveRequest struct {
	PathParams SessionRecordingsPropertiesRetrievePathParams
}

type SessionRecordingsPropertiesRetrieveResponse struct {
	ContentType string
	StatusCode  int
}
