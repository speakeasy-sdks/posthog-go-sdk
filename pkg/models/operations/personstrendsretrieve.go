package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsTrendsRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsTrendsRetrieveFormatEnum string

const (
	PersonsTrendsRetrieveFormatEnumCsv  PersonsTrendsRetrieveFormatEnum = "csv"
	PersonsTrendsRetrieveFormatEnumJSON PersonsTrendsRetrieveFormatEnum = "json"
)

type PersonsTrendsRetrieveQueryParams struct {
	Format *PersonsTrendsRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsTrendsRetrieveRequest struct {
	PathParams  PersonsTrendsRetrievePathParams
	QueryParams PersonsTrendsRetrieveQueryParams
}

type PersonsTrendsRetrieveResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int
}
