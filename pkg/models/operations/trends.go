package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type TrendsPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type TrendsFormatEnum string

const (
	TrendsFormatEnumCsv  TrendsFormatEnum = "csv"
	TrendsFormatEnumJSON TrendsFormatEnum = "json"
)

type TrendsQueryParams struct {
	Format *TrendsFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type TrendsRequest struct {
	PathParams  TrendsPathParams
	QueryParams TrendsQueryParams
	Request     *shared.Trend `request:"mediaType=application/json"`
}

type TrendsResponse struct {
	Body         []byte
	ContentType  string
	StatusCode   int64
	TrendResults *shared.TrendResults
}
