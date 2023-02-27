package operations

type QueryRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type QueryRetrieveQueryParams struct {
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

type QueryRetrieveRequest struct {
	PathParams  QueryRetrievePathParams
	QueryParams QueryRetrieveQueryParams
}

type QueryRetrieveResponse struct {
	ContentType string
	StatusCode  int
}
