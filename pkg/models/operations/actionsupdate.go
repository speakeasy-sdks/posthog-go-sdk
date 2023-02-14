package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ActionsUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ActionsUpdateFormatEnum string

const (
	ActionsUpdateFormatEnumCsv  ActionsUpdateFormatEnum = "csv"
	ActionsUpdateFormatEnumJSON ActionsUpdateFormatEnum = "json"
)

type ActionsUpdateQueryParams struct {
	Format *ActionsUpdateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type ActionsUpdateRequests struct {
	Action  *shared.ActionInput `request:"mediaType=application/json"`
	Action1 *shared.ActionInput `request:"mediaType=application/x-www-form-urlencoded"`
	Action2 *shared.ActionInput `request:"mediaType=multipart/form-data"`
}

type ActionsUpdateRequest struct {
	PathParams  ActionsUpdatePathParams
	QueryParams ActionsUpdateQueryParams
	Request     *ActionsUpdateRequests
}

type ActionsUpdateResponse struct {
	Action      *shared.Action
	Body        []byte
	ContentType string
	StatusCode  int64
}
