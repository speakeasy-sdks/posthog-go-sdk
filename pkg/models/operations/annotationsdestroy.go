package operations

type AnnotationsDestroyPathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type AnnotationsDestroyRequest struct {
	PathParams AnnotationsDestroyPathParams
}

type AnnotationsDestroyResponse struct {
	ContentType string
	StatusCode  int64
}
