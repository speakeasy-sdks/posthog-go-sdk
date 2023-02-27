package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsUpdateFormatEnum string

const (
	InsightsUpdateFormatEnumCsv  InsightsUpdateFormatEnum = "csv"
	InsightsUpdateFormatEnumJSON InsightsUpdateFormatEnum = "json"
)

type InsightsUpdateQueryParams struct {
	Format *InsightsUpdateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsUpdateRequest struct {
	PathParams  InsightsUpdatePathParams
	QueryParams InsightsUpdateQueryParams
	Request     *shared.InsightInput `request:"mediaType=application/json"`
}

type InsightsUpdateResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int
}
