package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsRetentionRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsRetentionRetrieveFormatEnum string

const (
	InsightsRetentionRetrieveFormatEnumCsv  InsightsRetentionRetrieveFormatEnum = "csv"
	InsightsRetentionRetrieveFormatEnumJSON InsightsRetentionRetrieveFormatEnum = "json"
)

type InsightsRetentionRetrieveQueryParams struct {
	Format *InsightsRetentionRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsRetentionRetrieveRequest struct {
	PathParams  InsightsRetentionRetrievePathParams
	QueryParams InsightsRetentionRetrieveQueryParams
}

type InsightsRetentionRetrieveResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int
}
