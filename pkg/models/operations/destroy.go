package operations

type DestroyPathParams struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type DestroyRequest struct {
	PathParams DestroyPathParams
}

type DestroyResponse struct {
	ContentType string
	StatusCode  int64
}
