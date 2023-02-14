package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsCohortsRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsCohortsRetrieveFormatEnum string

const (
	PersonsCohortsRetrieveFormatEnumCsv  PersonsCohortsRetrieveFormatEnum = "csv"
	PersonsCohortsRetrieveFormatEnumJSON PersonsCohortsRetrieveFormatEnum = "json"
)

type PersonsCohortsRetrieveQueryParams struct {
	Format *PersonsCohortsRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsCohortsRetrieveRequest struct {
	PathParams  PersonsCohortsRetrievePathParams
	QueryParams PersonsCohortsRetrieveQueryParams
}

type PersonsCohortsRetrieveResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int64
}
