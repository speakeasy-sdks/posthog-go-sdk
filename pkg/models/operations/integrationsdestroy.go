package operations

type IntegrationsDestroyPathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type IntegrationsDestroyRequest struct {
	PathParams IntegrationsDestroyPathParams
}

type IntegrationsDestroyResponse struct {
	ContentType string
	StatusCode  int
}
