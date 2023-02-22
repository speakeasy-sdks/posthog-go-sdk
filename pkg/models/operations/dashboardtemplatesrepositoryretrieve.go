package operations

type DashboardTemplatesRepositoryRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type DashboardTemplatesRepositoryRetrieveRequest struct {
	PathParams DashboardTemplatesRepositoryRetrievePathParams
}

type DashboardTemplatesRepositoryRetrieveResponse struct {
	ContentType string
	StatusCode  int
}
