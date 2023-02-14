package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsViewedCreatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsViewedCreateFormatEnum string

const (
	InsightsViewedCreateFormatEnumCsv  InsightsViewedCreateFormatEnum = "csv"
	InsightsViewedCreateFormatEnumJSON InsightsViewedCreateFormatEnum = "json"
)

type InsightsViewedCreateQueryParams struct {
	Format *InsightsViewedCreateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsViewedCreateRequest struct {
	PathParams  InsightsViewedCreatePathParams
	QueryParams InsightsViewedCreateQueryParams
	Request     *shared.InsightInput `request:"mediaType=application/json"`
}

type InsightsViewedCreateResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int64
}
