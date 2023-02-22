package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsCancelCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsCancelCreateFormatEnum string

const (
	InsightsCancelCreateFormatEnumCsv  InsightsCancelCreateFormatEnum = "csv"
	InsightsCancelCreateFormatEnumJSON InsightsCancelCreateFormatEnum = "json"
)

type InsightsCancelCreateQueryParams struct {
	Format *InsightsCancelCreateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsCancelCreateRequest struct {
	PathParams  InsightsCancelCreatePathParams
	QueryParams InsightsCancelCreateQueryParams
	Request     *shared.InsightInput `request:"mediaType=application/json"`
}

type InsightsCancelCreateResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int
}
