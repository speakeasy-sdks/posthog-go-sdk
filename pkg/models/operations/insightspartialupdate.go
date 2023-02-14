package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsPartialUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsPartialUpdateFormatEnum string

const (
	InsightsPartialUpdateFormatEnumCsv  InsightsPartialUpdateFormatEnum = "csv"
	InsightsPartialUpdateFormatEnumJSON InsightsPartialUpdateFormatEnum = "json"
)

type InsightsPartialUpdateQueryParams struct {
	Format *InsightsPartialUpdateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsPartialUpdateRequest struct {
	PathParams  InsightsPartialUpdatePathParams
	QueryParams InsightsPartialUpdateQueryParams
	Request     *shared.PatchedInsightInput `request:"mediaType=application/json"`
}

type InsightsPartialUpdateResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int64
}
