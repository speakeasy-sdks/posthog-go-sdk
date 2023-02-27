package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsFunnelCorrelationCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsFunnelCorrelationCreateFormatEnum string

const (
	InsightsFunnelCorrelationCreateFormatEnumCsv  InsightsFunnelCorrelationCreateFormatEnum = "csv"
	InsightsFunnelCorrelationCreateFormatEnumJSON InsightsFunnelCorrelationCreateFormatEnum = "json"
)

type InsightsFunnelCorrelationCreateQueryParams struct {
	Format *InsightsFunnelCorrelationCreateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsFunnelCorrelationCreateRequest struct {
	PathParams  InsightsFunnelCorrelationCreatePathParams
	QueryParams InsightsFunnelCorrelationCreateQueryParams
	Request     *shared.InsightInput `request:"mediaType=application/json"`
}

type InsightsFunnelCorrelationCreateResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int
}
