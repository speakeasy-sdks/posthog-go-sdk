package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsActivityRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsActivityRetrieveFormatEnum string

const (
	PersonsActivityRetrieveFormatEnumCsv  PersonsActivityRetrieveFormatEnum = "csv"
	PersonsActivityRetrieveFormatEnumJSON PersonsActivityRetrieveFormatEnum = "json"
)

type PersonsActivityRetrieveQueryParams struct {
	Format *PersonsActivityRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsActivityRetrieveRequest struct {
	PathParams  PersonsActivityRetrievePathParams
	QueryParams PersonsActivityRetrieveQueryParams
}

type PersonsActivityRetrieveResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int64
}
