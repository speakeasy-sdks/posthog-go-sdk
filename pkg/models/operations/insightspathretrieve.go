package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsPathRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsPathRetrieveFormatEnum string

const (
	InsightsPathRetrieveFormatEnumCsv  InsightsPathRetrieveFormatEnum = "csv"
	InsightsPathRetrieveFormatEnumJSON InsightsPathRetrieveFormatEnum = "json"
)

type InsightsPathRetrieveQueryParams struct {
	Format *InsightsPathRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsPathRetrieveRequest struct {
	PathParams  InsightsPathRetrievePathParams
	QueryParams InsightsPathRetrieveQueryParams
}

type InsightsPathRetrieveResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int64
}
