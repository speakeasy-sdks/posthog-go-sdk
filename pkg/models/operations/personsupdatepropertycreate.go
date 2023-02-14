package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsUpdatePropertyCreatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsUpdatePropertyCreateFormatEnum string

const (
	PersonsUpdatePropertyCreateFormatEnumCsv  PersonsUpdatePropertyCreateFormatEnum = "csv"
	PersonsUpdatePropertyCreateFormatEnumJSON PersonsUpdatePropertyCreateFormatEnum = "json"
)

type PersonsUpdatePropertyCreateQueryParams struct {
	Format *PersonsUpdatePropertyCreateFormatEnum `queryParam:"style=form,explode=true,name=format"`
	Key    string                                 `queryParam:"style=form,explode=true,name=key"`
	Value  interface{}                            `queryParam:"style=form,explode=true,name=value"`
}

type PersonsUpdatePropertyCreateRequests struct {
	Person  *shared.PersonInput `request:"mediaType=application/json"`
	Person1 *shared.PersonInput `request:"mediaType=application/x-www-form-urlencoded"`
	Person2 *shared.PersonInput `request:"mediaType=multipart/form-data"`
}

type PersonsUpdatePropertyCreateRequest struct {
	PathParams  PersonsUpdatePropertyCreatePathParams
	QueryParams PersonsUpdatePropertyCreateQueryParams
	Request     *PersonsUpdatePropertyCreateRequests
}

type PersonsUpdatePropertyCreateResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int64
}
