package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsTrendRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsTrendRetrieveFormatEnum string

const (
	InsightsTrendRetrieveFormatEnumCsv  InsightsTrendRetrieveFormatEnum = "csv"
	InsightsTrendRetrieveFormatEnumJSON InsightsTrendRetrieveFormatEnum = "json"
)

type InsightsTrendRetrieveQueryParams struct {
	Format *InsightsTrendRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsTrendRetrieveRequest struct {
	PathParams  InsightsTrendRetrievePathParams
	QueryParams InsightsTrendRetrieveQueryParams
}

type InsightsTrendRetrieveResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int
}
