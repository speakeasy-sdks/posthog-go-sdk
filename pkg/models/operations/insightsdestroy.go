package operations

type InsightsDestroyPathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsDestroyFormatEnum string

const (
	InsightsDestroyFormatEnumCsv  InsightsDestroyFormatEnum = "csv"
	InsightsDestroyFormatEnumJSON InsightsDestroyFormatEnum = "json"
)

type InsightsDestroyQueryParams struct {
	Format *InsightsDestroyFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsDestroyRequest struct {
	PathParams  InsightsDestroyPathParams
	QueryParams InsightsDestroyQueryParams
}

type InsightsDestroyResponse struct {
	ContentType string
	StatusCode  int64
}
