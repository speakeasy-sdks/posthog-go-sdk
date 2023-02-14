package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsMyLastViewedRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsMyLastViewedRetrieveFormatEnum string

const (
	InsightsMyLastViewedRetrieveFormatEnumCsv  InsightsMyLastViewedRetrieveFormatEnum = "csv"
	InsightsMyLastViewedRetrieveFormatEnumJSON InsightsMyLastViewedRetrieveFormatEnum = "json"
)

type InsightsMyLastViewedRetrieveQueryParams struct {
	Format *InsightsMyLastViewedRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsMyLastViewedRetrieveRequest struct {
	PathParams  InsightsMyLastViewedRetrievePathParams
	QueryParams InsightsMyLastViewedRetrieveQueryParams
}

type InsightsMyLastViewedRetrieveResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int64
}
