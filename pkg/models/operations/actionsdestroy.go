package operations

type ActionsDestroyPathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ActionsDestroyFormatEnum string

const (
	ActionsDestroyFormatEnumCsv  ActionsDestroyFormatEnum = "csv"
	ActionsDestroyFormatEnumJSON ActionsDestroyFormatEnum = "json"
)

type ActionsDestroyQueryParams struct {
	Format *ActionsDestroyFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type ActionsDestroyRequest struct {
	PathParams  ActionsDestroyPathParams
	QueryParams ActionsDestroyQueryParams
}

type ActionsDestroyResponse struct {
	ContentType string
	StatusCode  int
}
