package operations

type PersonsDestroyPathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsDestroyFormatEnum string

const (
	PersonsDestroyFormatEnumCsv  PersonsDestroyFormatEnum = "csv"
	PersonsDestroyFormatEnumJSON PersonsDestroyFormatEnum = "json"
)

type PersonsDestroyQueryParams struct {
	DeleteEvents *bool                     `queryParam:"style=form,explode=true,name=delete_events"`
	Format       *PersonsDestroyFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsDestroyRequest struct {
	PathParams  PersonsDestroyPathParams
	QueryParams PersonsDestroyQueryParams
}

type PersonsDestroyResponse struct {
	ContentType string
	StatusCode  int64
}
