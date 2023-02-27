package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsPathCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsPathCreateFormatEnum string

const (
	PersonsPathCreateFormatEnumCsv  PersonsPathCreateFormatEnum = "csv"
	PersonsPathCreateFormatEnumJSON PersonsPathCreateFormatEnum = "json"
)

type PersonsPathCreateQueryParams struct {
	Format *PersonsPathCreateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsPathCreateRequests struct {
	Person  *shared.PersonInput `request:"mediaType=application/json"`
	Person1 *shared.PersonInput `request:"mediaType=application/x-www-form-urlencoded"`
	Person2 *shared.PersonInput `request:"mediaType=multipart/form-data"`
}

type PersonsPathCreateRequest struct {
	PathParams  PersonsPathCreatePathParams
	QueryParams PersonsPathCreateQueryParams
	Request     *PersonsPathCreateRequests
}

type PersonsPathCreateResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int
}
