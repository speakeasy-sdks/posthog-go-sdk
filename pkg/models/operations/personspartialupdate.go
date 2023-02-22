package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsPartialUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsPartialUpdateFormatEnum string

const (
	PersonsPartialUpdateFormatEnumCsv  PersonsPartialUpdateFormatEnum = "csv"
	PersonsPartialUpdateFormatEnumJSON PersonsPartialUpdateFormatEnum = "json"
)

type PersonsPartialUpdateQueryParams struct {
	Format *PersonsPartialUpdateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsPartialUpdateRequests struct {
	PatchedPerson  *shared.PatchedPersonInput `request:"mediaType=application/json"`
	PatchedPerson1 *shared.PatchedPersonInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedPerson2 *shared.PatchedPersonInput `request:"mediaType=multipart/form-data"`
}

type PersonsPartialUpdateRequest struct {
	PathParams  PersonsPartialUpdatePathParams
	QueryParams PersonsPartialUpdateQueryParams
	Request     *PersonsPartialUpdateRequests
}

type PersonsPartialUpdateResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int
}
