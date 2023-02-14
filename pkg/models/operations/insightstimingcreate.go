package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsTimingCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsTimingCreateFormatEnum string

const (
	InsightsTimingCreateFormatEnumCsv  InsightsTimingCreateFormatEnum = "csv"
	InsightsTimingCreateFormatEnumJSON InsightsTimingCreateFormatEnum = "json"
)

type InsightsTimingCreateQueryParams struct {
	Format *InsightsTimingCreateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsTimingCreateRequest struct {
	PathParams  InsightsTimingCreatePathParams
	QueryParams InsightsTimingCreateQueryParams
	Request     *shared.InsightInput `request:"mediaType=application/json"`
}

type InsightsTimingCreateResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int64
}
