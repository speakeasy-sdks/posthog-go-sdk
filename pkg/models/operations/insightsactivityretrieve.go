package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsActivityRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsActivityRetrieveFormatEnum string

const (
	InsightsActivityRetrieveFormatEnumCsv  InsightsActivityRetrieveFormatEnum = "csv"
	InsightsActivityRetrieveFormatEnumJSON InsightsActivityRetrieveFormatEnum = "json"
)

type InsightsActivityRetrieveQueryParams struct {
	Format *InsightsActivityRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsActivityRetrieveRequest struct {
	PathParams  InsightsActivityRetrievePathParams
	QueryParams InsightsActivityRetrieveQueryParams
}

type InsightsActivityRetrieveResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int
}
