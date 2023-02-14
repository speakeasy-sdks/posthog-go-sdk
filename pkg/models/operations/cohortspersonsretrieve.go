package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type CohortsPersonsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type CohortsPersonsRetrieveFormatEnum string

const (
	CohortsPersonsRetrieveFormatEnumCsv  CohortsPersonsRetrieveFormatEnum = "csv"
	CohortsPersonsRetrieveFormatEnumJSON CohortsPersonsRetrieveFormatEnum = "json"
)

type CohortsPersonsRetrieveQueryParams struct {
	Format *CohortsPersonsRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type CohortsPersonsRetrieveRequest struct {
	PathParams  CohortsPersonsRetrievePathParams
	QueryParams CohortsPersonsRetrieveQueryParams
}

type CohortsPersonsRetrieveResponse struct {
	Body        []byte
	Cohort      *shared.Cohort
	ContentType string
	StatusCode  int64
}
