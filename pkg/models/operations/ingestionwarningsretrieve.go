package operations

type IngestionWarningsRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type IngestionWarningsRetrieveRequest struct {
	PathParams IngestionWarningsRetrievePathParams
}

type IngestionWarningsRetrieveResponse struct {
	ContentType string
	StatusCode  int64
}
