package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsFunnelCorrelationRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsFunnelCorrelationRetrieveFormatEnum string

const (
	InsightsFunnelCorrelationRetrieveFormatEnumCsv  InsightsFunnelCorrelationRetrieveFormatEnum = "csv"
	InsightsFunnelCorrelationRetrieveFormatEnumJSON InsightsFunnelCorrelationRetrieveFormatEnum = "json"
)

type InsightsFunnelCorrelationRetrieveQueryParams struct {
	Format *InsightsFunnelCorrelationRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsFunnelCorrelationRetrieveRequest struct {
	PathParams  InsightsFunnelCorrelationRetrievePathParams
	QueryParams InsightsFunnelCorrelationRetrieveQueryParams
}

type InsightsFunnelCorrelationRetrieveResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int
}
