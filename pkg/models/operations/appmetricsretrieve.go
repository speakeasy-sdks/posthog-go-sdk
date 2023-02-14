package operations

type AppMetricsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type AppMetricsRetrieveRequest struct {
	PathParams AppMetricsRetrievePathParams
}

type AppMetricsRetrieveResponse struct {
	ContentType string
	StatusCode  int64
}
