package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ActionsPeopleRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ActionsPeopleRetrieveFormatEnum string

const (
	ActionsPeopleRetrieveFormatEnumCsv  ActionsPeopleRetrieveFormatEnum = "csv"
	ActionsPeopleRetrieveFormatEnumJSON ActionsPeopleRetrieveFormatEnum = "json"
)

type ActionsPeopleRetrieveQueryParams struct {
	Format *ActionsPeopleRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type ActionsPeopleRetrieveRequest struct {
	PathParams  ActionsPeopleRetrievePathParams
	QueryParams ActionsPeopleRetrieveQueryParams
}

type ActionsPeopleRetrieveResponse struct {
	Action      *shared.Action
	Body        []byte
	ContentType string
	StatusCode  int64
}
