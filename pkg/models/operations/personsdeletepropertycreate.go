package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsDeletePropertyCreatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsDeletePropertyCreateFormatEnum string

const (
	PersonsDeletePropertyCreateFormatEnumCsv  PersonsDeletePropertyCreateFormatEnum = "csv"
	PersonsDeletePropertyCreateFormatEnumJSON PersonsDeletePropertyCreateFormatEnum = "json"
)

type PersonsDeletePropertyCreateQueryParams struct {
	DollarUnset string                                 `queryParam:"style=form,explode=true,name=$unset"`
	Format      *PersonsDeletePropertyCreateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsDeletePropertyCreateRequests struct {
	Person  *shared.PersonInput `request:"mediaType=application/json"`
	Person1 *shared.PersonInput `request:"mediaType=application/x-www-form-urlencoded"`
	Person2 *shared.PersonInput `request:"mediaType=multipart/form-data"`
}

type PersonsDeletePropertyCreateRequest struct {
	PathParams  PersonsDeletePropertyCreatePathParams
	QueryParams PersonsDeletePropertyCreateQueryParams
	Request     *PersonsDeletePropertyCreateRequests
}

type PersonsDeletePropertyCreateResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int
}
