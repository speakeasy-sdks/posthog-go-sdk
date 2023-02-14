package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsPathCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsPathCreateFormatEnum string

const (
	InsightsPathCreateFormatEnumCsv  InsightsPathCreateFormatEnum = "csv"
	InsightsPathCreateFormatEnumJSON InsightsPathCreateFormatEnum = "json"
)

type InsightsPathCreateQueryParams struct {
	Format *InsightsPathCreateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsPathCreateRequest struct {
	PathParams  InsightsPathCreatePathParams
	QueryParams InsightsPathCreateQueryParams
	Request     *shared.InsightInput `request:"mediaType=application/json"`
}

type InsightsPathCreateResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int64
}
