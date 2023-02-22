package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsSplitCreatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsSplitCreateFormatEnum string

const (
	PersonsSplitCreateFormatEnumCsv  PersonsSplitCreateFormatEnum = "csv"
	PersonsSplitCreateFormatEnumJSON PersonsSplitCreateFormatEnum = "json"
)

type PersonsSplitCreateQueryParams struct {
	Format *PersonsSplitCreateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsSplitCreateRequests struct {
	Person  *shared.PersonInput `request:"mediaType=application/json"`
	Person1 *shared.PersonInput `request:"mediaType=application/x-www-form-urlencoded"`
	Person2 *shared.PersonInput `request:"mediaType=multipart/form-data"`
}

type PersonsSplitCreateRequest struct {
	PathParams  PersonsSplitCreatePathParams
	QueryParams PersonsSplitCreateQueryParams
	Request     *PersonsSplitCreateRequests
}

type PersonsSplitCreateResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int
}
