package operations

type HooksDestroyPathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type HooksDestroyRequest struct {
	PathParams HooksDestroyPathParams
}

type HooksDestroyResponse struct {
	ContentType string
	StatusCode  int64
}
