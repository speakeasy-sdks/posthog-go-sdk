package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsCreateFormatEnum string

const (
	InsightsCreateFormatEnumCsv  InsightsCreateFormatEnum = "csv"
	InsightsCreateFormatEnumJSON InsightsCreateFormatEnum = "json"
)

type InsightsCreateQueryParams struct {
	Format *InsightsCreateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsCreateRequest struct {
	PathParams  InsightsCreatePathParams
	QueryParams InsightsCreateQueryParams
	Request     *shared.InsightInput `request:"mediaType=application/json"`
}

type InsightsCreateResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int64
}
