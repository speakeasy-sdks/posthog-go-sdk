package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FunnelsPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FunnelsFormatEnum string

const (
	FunnelsFormatEnumCsv  FunnelsFormatEnum = "csv"
	FunnelsFormatEnumJSON FunnelsFormatEnum = "json"
)

type FunnelsQueryParams struct {
	Format *FunnelsFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type FunnelsRequest struct {
	PathParams  FunnelsPathParams
	QueryParams FunnelsQueryParams
	Request     *shared.Funnel `request:"mediaType=application/json"`
}

type FunnelsResponse struct {
	Body               []byte
	ContentType        string
	FunnelStepsResults *shared.FunnelStepsResults
	StatusCode         int64
}
