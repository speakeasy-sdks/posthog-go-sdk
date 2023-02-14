package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsPropertiesRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsPropertiesRetrieveFormatEnum string

const (
	PersonsPropertiesRetrieveFormatEnumCsv  PersonsPropertiesRetrieveFormatEnum = "csv"
	PersonsPropertiesRetrieveFormatEnumJSON PersonsPropertiesRetrieveFormatEnum = "json"
)

type PersonsPropertiesRetrieveQueryParams struct {
	Format *PersonsPropertiesRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsPropertiesRetrieveRequest struct {
	PathParams  PersonsPropertiesRetrievePathParams
	QueryParams PersonsPropertiesRetrieveQueryParams
}

type PersonsPropertiesRetrieveResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int64
}
