package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsValuesRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsValuesRetrieveFormatEnum string

const (
	PersonsValuesRetrieveFormatEnumCsv  PersonsValuesRetrieveFormatEnum = "csv"
	PersonsValuesRetrieveFormatEnumJSON PersonsValuesRetrieveFormatEnum = "json"
)

type PersonsValuesRetrieveQueryParams struct {
	Format *PersonsValuesRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsValuesRetrieveRequest struct {
	PathParams  PersonsValuesRetrievePathParams
	QueryParams PersonsValuesRetrieveQueryParams
}

type PersonsValuesRetrieveResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int64
}
