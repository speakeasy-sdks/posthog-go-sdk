package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsFunnelRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsFunnelRetrieveFormatEnum string

const (
	InsightsFunnelRetrieveFormatEnumCsv  InsightsFunnelRetrieveFormatEnum = "csv"
	InsightsFunnelRetrieveFormatEnumJSON InsightsFunnelRetrieveFormatEnum = "json"
)

type InsightsFunnelRetrieveQueryParams struct {
	Format *InsightsFunnelRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsFunnelRetrieveRequest struct {
	PathParams  InsightsFunnelRetrievePathParams
	QueryParams InsightsFunnelRetrieveQueryParams
}

type InsightsFunnelRetrieveResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int
}
