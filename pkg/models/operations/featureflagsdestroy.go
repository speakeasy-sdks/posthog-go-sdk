package operations

type FeatureFlagsDestroyPathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsDestroyRequest struct {
	PathParams FeatureFlagsDestroyPathParams
}

type FeatureFlagsDestroyResponse struct {
	ContentType string
	StatusCode  int
}
