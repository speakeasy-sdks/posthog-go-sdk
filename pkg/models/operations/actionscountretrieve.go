package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ActionsCountRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ActionsCountRetrieveFormatEnum string

const (
	ActionsCountRetrieveFormatEnumCsv  ActionsCountRetrieveFormatEnum = "csv"
	ActionsCountRetrieveFormatEnumJSON ActionsCountRetrieveFormatEnum = "json"
)

type ActionsCountRetrieveQueryParams struct {
	Format *ActionsCountRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type ActionsCountRetrieveRequest struct {
	PathParams  ActionsCountRetrievePathParams
	QueryParams ActionsCountRetrieveQueryParams
}

type ActionsCountRetrieveResponse struct {
	Action      *shared.Action
	Body        []byte
	ContentType string
	StatusCode  int64
}
