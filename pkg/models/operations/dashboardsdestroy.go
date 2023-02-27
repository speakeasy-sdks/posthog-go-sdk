package operations

type DashboardsDestroyPathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type DashboardsDestroyRequest struct {
	PathParams DashboardsDestroyPathParams
}

type DashboardsDestroyResponse struct {
	ContentType string
	StatusCode  int
}
