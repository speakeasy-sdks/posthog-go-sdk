package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsRetrieveFormatEnum string

const (
	PersonsRetrieveFormatEnumCsv  PersonsRetrieveFormatEnum = "csv"
	PersonsRetrieveFormatEnumJSON PersonsRetrieveFormatEnum = "json"
)

type PersonsRetrieveQueryParams struct {
	Format *PersonsRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsRetrieveRequest struct {
	PathParams  PersonsRetrievePathParams
	QueryParams PersonsRetrieveQueryParams
}

type PersonsRetrieveResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int64
}
