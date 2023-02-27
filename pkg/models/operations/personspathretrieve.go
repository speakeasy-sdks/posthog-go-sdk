package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsPathRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsPathRetrieveFormatEnum string

const (
	PersonsPathRetrieveFormatEnumCsv  PersonsPathRetrieveFormatEnum = "csv"
	PersonsPathRetrieveFormatEnumJSON PersonsPathRetrieveFormatEnum = "json"
)

type PersonsPathRetrieveQueryParams struct {
	Format *PersonsPathRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsPathRetrieveRequest struct {
	PathParams  PersonsPathRetrievePathParams
	QueryParams PersonsPathRetrieveQueryParams
}

type PersonsPathRetrieveResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int
}
