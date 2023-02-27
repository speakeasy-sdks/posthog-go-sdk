package operations

type PluginConfigsDestroyPathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PluginConfigsDestroyRequest struct {
	PathParams PluginConfigsDestroyPathParams
}

type PluginConfigsDestroyResponse struct {
	ContentType string
	StatusCode  int
}
