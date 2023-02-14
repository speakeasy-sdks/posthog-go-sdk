package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ActionsPartialUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ActionsPartialUpdateFormatEnum string

const (
	ActionsPartialUpdateFormatEnumCsv  ActionsPartialUpdateFormatEnum = "csv"
	ActionsPartialUpdateFormatEnumJSON ActionsPartialUpdateFormatEnum = "json"
)

type ActionsPartialUpdateQueryParams struct {
	Format *ActionsPartialUpdateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type ActionsPartialUpdateRequests struct {
	PatchedAction  *shared.PatchedActionInput `request:"mediaType=application/json"`
	PatchedAction1 *shared.PatchedActionInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedAction2 *shared.PatchedActionInput `request:"mediaType=multipart/form-data"`
}

type ActionsPartialUpdateRequest struct {
	PathParams  ActionsPartialUpdatePathParams
	QueryParams ActionsPartialUpdateQueryParams
	Request     *ActionsPartialUpdateRequests
}

type ActionsPartialUpdateResponse struct {
	Action      *shared.Action
	Body        []byte
	ContentType string
	StatusCode  int64
}
