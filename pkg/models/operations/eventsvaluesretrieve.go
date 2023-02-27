package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type EventsValuesRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type EventsValuesRetrieveFormatEnum string

const (
	EventsValuesRetrieveFormatEnumCsv  EventsValuesRetrieveFormatEnum = "csv"
	EventsValuesRetrieveFormatEnumJSON EventsValuesRetrieveFormatEnum = "json"
)

type EventsValuesRetrieveQueryParams struct {
	Format *EventsValuesRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type EventsValuesRetrieveRequest struct {
	PathParams  EventsValuesRetrievePathParams
	QueryParams EventsValuesRetrieveQueryParams
}

type EventsValuesRetrieveResponse struct {
	Body            []byte
	ClickhouseEvent *shared.ClickhouseEvent
	ContentType     string
	StatusCode      int
}
