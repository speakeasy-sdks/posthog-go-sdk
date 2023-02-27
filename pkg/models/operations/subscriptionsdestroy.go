package operations

type SubscriptionsDestroyPathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type SubscriptionsDestroyRequest struct {
	PathParams SubscriptionsDestroyPathParams
}

type SubscriptionsDestroyResponse struct {
	ContentType string
	StatusCode  int
}
