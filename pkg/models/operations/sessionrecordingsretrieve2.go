package operations

type SessionRecordingsRetrieve2PathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type SessionRecordingsRetrieve2Request struct {
	PathParams SessionRecordingsRetrieve2PathParams
}

type SessionRecordingsRetrieve2Response struct {
	ContentType string
	StatusCode  int64
}
