package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InsightsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type InsightsRetrieveFormatEnum string

const (
	InsightsRetrieveFormatEnumCsv  InsightsRetrieveFormatEnum = "csv"
	InsightsRetrieveFormatEnumJSON InsightsRetrieveFormatEnum = "json"
)

type InsightsRetrieveQueryParams struct {
	Format        *InsightsRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
	FromDashboard *int64                      `queryParam:"style=form,explode=true,name=from_dashboard"`
	Refresh       *bool                       `queryParam:"style=form,explode=true,name=refresh"`
}

type InsightsRetrieveRequest struct {
	PathParams  InsightsRetrievePathParams
	QueryParams InsightsRetrieveQueryParams
}

type InsightsRetrieveResponse struct {
	Body        []byte
	ContentType string
	Insight     *shared.Insight
	StatusCode  int
}
