package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsListFormatEnum string

const (
	InsightsListFormatEnumCsv  InsightsListFormatEnum = "csv"
	InsightsListFormatEnumJSON InsightsListFormatEnum = "json"
)

type InsightsListQueryParams struct {
	CreatedBy *int64                  `queryParam:"style=form,explode=true,name=created_by"`
	Format    *InsightsListFormatEnum `queryParam:"style=form,explode=true,name=format"`
	Limit     *int64                  `queryParam:"style=form,explode=true,name=limit"`
	Offset    *int64                  `queryParam:"style=form,explode=true,name=offset"`
	ShortID   *string                 `queryParam:"style=form,explode=true,name=short_id"`
}

type InsightsListRequest struct {
	PathParams  InsightsListPathParams
	QueryParams InsightsListQueryParams
}

type InsightsListResponse struct {
	Body                 []byte
	ContentType          string
	PaginatedInsightList *shared.PaginatedInsightList
	StatusCode           int
}
