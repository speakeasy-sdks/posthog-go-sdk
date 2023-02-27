package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsStickinessRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsStickinessRetrieveFormatEnum string

const (
	PersonsStickinessRetrieveFormatEnumCsv  PersonsStickinessRetrieveFormatEnum = "csv"
	PersonsStickinessRetrieveFormatEnumJSON PersonsStickinessRetrieveFormatEnum = "json"
)

type PersonsStickinessRetrieveQueryParams struct {
	Format *PersonsStickinessRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsStickinessRetrieveRequest struct {
	PathParams  PersonsStickinessRetrievePathParams
	QueryParams PersonsStickinessRetrieveQueryParams
}

type PersonsStickinessRetrieveResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int
}
