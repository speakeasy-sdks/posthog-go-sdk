package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsActivityRetrieve2PathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsActivityRetrieve2FormatEnum string

const (
	PersonsActivityRetrieve2FormatEnumCsv  PersonsActivityRetrieve2FormatEnum = "csv"
	PersonsActivityRetrieve2FormatEnumJSON PersonsActivityRetrieve2FormatEnum = "json"
)

type PersonsActivityRetrieve2QueryParams struct {
	Format *PersonsActivityRetrieve2FormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsActivityRetrieve2Request struct {
	PathParams  PersonsActivityRetrieve2PathParams
	QueryParams PersonsActivityRetrieve2QueryParams
}

type PersonsActivityRetrieve2Response struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int64
}
