package operations

type ExperimentsDestroyPathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ExperimentsDestroyRequest struct {
	PathParams ExperimentsDestroyPathParams
}

type ExperimentsDestroyResponse struct {
	ContentType string
	StatusCode  int
}
