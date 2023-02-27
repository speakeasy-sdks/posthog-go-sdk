package operations

type AppMetricsErrorDetailsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type AppMetricsErrorDetailsRetrieveRequest struct {
	PathParams AppMetricsErrorDetailsRetrievePathParams
}

type AppMetricsErrorDetailsRetrieveResponse struct {
	ContentType string
	StatusCode  int
}
