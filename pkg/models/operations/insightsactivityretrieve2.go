package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsActivityRetrieve2PathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsActivityRetrieve2FormatEnum string

const (
	InsightsActivityRetrieve2FormatEnumCsv  InsightsActivityRetrieve2FormatEnum = "csv"
	InsightsActivityRetrieve2FormatEnumJSON InsightsActivityRetrieve2FormatEnum = "json"
)

type InsightsActivityRetrieve2QueryParams struct {
	Format *InsightsActivityRetrieve2FormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type InsightsActivityRetrieve2Request struct {
	PathParams  InsightsActivityRetrieve2PathParams
	QueryParams InsightsActivityRetrieve2QueryParams
}

type InsightsActivityRetrieve2Response struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int
}
