package operations

type UploadedMediaCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type UploadedMediaCreateRequest struct {
	PathParams UploadedMediaCreatePathParams
}

type UploadedMediaCreateResponse struct {
	ContentType string
	StatusCode  int
}
