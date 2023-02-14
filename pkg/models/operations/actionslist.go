package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ActionsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ActionsListFormatEnum string

const (
	ActionsListFormatEnumCsv  ActionsListFormatEnum = "csv"
	ActionsListFormatEnumJSON ActionsListFormatEnum = "json"
)

type ActionsListQueryParams struct {
	Format *ActionsListFormatEnum `queryParam:"style=form,explode=true,name=format"`
	Limit  *int64                 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64                 `queryParam:"style=form,explode=true,name=offset"`
}

type ActionsListRequest struct {
	PathParams  ActionsListPathParams
	QueryParams ActionsListQueryParams
}

type ActionsListResponse struct {
	Body                []byte
	ContentType         string
	PaginatedActionList *shared.PaginatedActionList
	StatusCode          int64
}
