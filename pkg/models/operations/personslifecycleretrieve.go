package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsLifecycleRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsLifecycleRetrieveFormatEnum string

const (
	PersonsLifecycleRetrieveFormatEnumCsv  PersonsLifecycleRetrieveFormatEnum = "csv"
	PersonsLifecycleRetrieveFormatEnumJSON PersonsLifecycleRetrieveFormatEnum = "json"
)

type PersonsLifecycleRetrieveQueryParams struct {
	Format *PersonsLifecycleRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsLifecycleRetrieveRequest struct {
	PathParams  PersonsLifecycleRetrievePathParams
	QueryParams PersonsLifecycleRetrieveQueryParams
}

type PersonsLifecycleRetrieveResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int
}
