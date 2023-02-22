package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type GroupsRelatedRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type GroupsRelatedRetrieveQueryParams struct {
	GroupTypeIndex int64  `queryParam:"style=form,explode=true,name=group_type_index"`
	ID             string `queryParam:"style=form,explode=true,name=id"`
}

type GroupsRelatedRetrieveRequest struct {
	PathParams  GroupsRelatedRetrievePathParams
	QueryParams GroupsRelatedRetrieveQueryParams
}

type GroupsRelatedRetrieveResponse struct {
	ContentType string
	Group       *shared.Group
	StatusCode  int
}
