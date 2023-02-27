package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsPropertiesTimelineRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsPropertiesTimelineRetrieveFormatEnum string

const (
	PersonsPropertiesTimelineRetrieveFormatEnumCsv  PersonsPropertiesTimelineRetrieveFormatEnum = "csv"
	PersonsPropertiesTimelineRetrieveFormatEnumJSON PersonsPropertiesTimelineRetrieveFormatEnum = "json"
)

type PersonsPropertiesTimelineRetrieveQueryParams struct {
	Format *PersonsPropertiesTimelineRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsPropertiesTimelineRetrieveRequest struct {
	PathParams  PersonsPropertiesTimelineRetrievePathParams
	QueryParams PersonsPropertiesTimelineRetrieveQueryParams
}

type PersonsPropertiesTimelineRetrieveResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int
}
