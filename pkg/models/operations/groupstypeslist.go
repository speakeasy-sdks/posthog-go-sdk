package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type GroupsTypesListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type GroupsTypesListRequest struct {
	PathParams GroupsTypesListPathParams
}

type GroupsTypesListResponse struct {
	ContentType string
	GroupTypes  []shared.GroupType
	StatusCode  int64
}
