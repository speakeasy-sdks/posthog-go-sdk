package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type EventsRetrievePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type EventsRetrieveFormatEnum string

const (
	EventsRetrieveFormatEnumCsv  EventsRetrieveFormatEnum = "csv"
	EventsRetrieveFormatEnumJSON EventsRetrieveFormatEnum = "json"
)

type EventsRetrieveQueryParams struct {
	Format *EventsRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type EventsRetrieveRequest struct {
	PathParams  EventsRetrievePathParams
	QueryParams EventsRetrieveQueryParams
}

type EventsRetrieveResponse struct {
	Body            []byte
	ClickhouseEvent *shared.ClickhouseEvent
	ContentType     string
	StatusCode      int64
}
