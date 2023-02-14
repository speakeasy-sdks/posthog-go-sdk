package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ActionsCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ActionsCreateFormatEnum string

const (
	ActionsCreateFormatEnumCsv  ActionsCreateFormatEnum = "csv"
	ActionsCreateFormatEnumJSON ActionsCreateFormatEnum = "json"
)

type ActionsCreateQueryParams struct {
	Format *ActionsCreateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type ActionsCreateRequests struct {
	Action  *shared.ActionInput `request:"mediaType=application/json"`
	Action1 *shared.ActionInput `request:"mediaType=application/x-www-form-urlencoded"`
	Action2 *shared.ActionInput `request:"mediaType=multipart/form-data"`
}

type ActionsCreateRequest struct {
	PathParams  ActionsCreatePathParams
	QueryParams ActionsCreateQueryParams
	Request     *ActionsCreateRequests
}

type ActionsCreateResponse struct {
	Action      *shared.Action
	Body        []byte
	ContentType string
	StatusCode  int64
}
