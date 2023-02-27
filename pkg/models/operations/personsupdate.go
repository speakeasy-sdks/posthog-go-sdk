package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsUpdateFormatEnum string

const (
	PersonsUpdateFormatEnumCsv  PersonsUpdateFormatEnum = "csv"
	PersonsUpdateFormatEnumJSON PersonsUpdateFormatEnum = "json"
)

type PersonsUpdateQueryParams struct {
	Format *PersonsUpdateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsUpdateRequests struct {
	Person  *shared.PersonInput `request:"mediaType=application/json"`
	Person1 *shared.PersonInput `request:"mediaType=application/x-www-form-urlencoded"`
	Person2 *shared.PersonInput `request:"mediaType=multipart/form-data"`
}

type PersonsUpdateRequest struct {
	PathParams  PersonsUpdatePathParams
	QueryParams PersonsUpdateQueryParams
	Request     *PersonsUpdateRequests
}

type PersonsUpdateResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int
}
