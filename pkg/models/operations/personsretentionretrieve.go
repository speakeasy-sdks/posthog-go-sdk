package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsRetentionRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsRetentionRetrieveFormatEnum string

const (
	PersonsRetentionRetrieveFormatEnumCsv  PersonsRetentionRetrieveFormatEnum = "csv"
	PersonsRetentionRetrieveFormatEnumJSON PersonsRetentionRetrieveFormatEnum = "json"
)

type PersonsRetentionRetrieveQueryParams struct {
	Format *PersonsRetentionRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsRetentionRetrieveRequest struct {
	PathParams  PersonsRetentionRetrievePathParams
	QueryParams PersonsRetentionRetrieveQueryParams
}

type PersonsRetentionRetrieveResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int64
}
