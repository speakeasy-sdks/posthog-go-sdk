package operations

type CohortsDestroyPathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type CohortsDestroyRequest struct {
	PathParams CohortsDestroyPathParams
}

type CohortsDestroyResponse struct {
	ContentType string
	StatusCode  int
}
