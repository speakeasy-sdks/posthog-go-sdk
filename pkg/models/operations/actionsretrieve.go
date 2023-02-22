package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ActionsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ActionsRetrieveFormatEnum string

const (
	ActionsRetrieveFormatEnumCsv  ActionsRetrieveFormatEnum = "csv"
	ActionsRetrieveFormatEnumJSON ActionsRetrieveFormatEnum = "json"
)

type ActionsRetrieveQueryParams struct {
	Format *ActionsRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type ActionsRetrieveRequest struct {
	PathParams  ActionsRetrievePathParams
	QueryParams ActionsRetrieveQueryParams
}

type ActionsRetrieveResponse struct {
	Action      *shared.Action
	Body        []byte
	ContentType string
	StatusCode  int
}
