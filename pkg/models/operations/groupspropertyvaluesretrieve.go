package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type GroupsPropertyValuesRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type GroupsPropertyValuesRetrieveQueryParams struct {
	GroupTypeIndex int64  `queryParam:"style=form,explode=true,name=group_type_index"`
	Key            string `queryParam:"style=form,explode=true,name=key"`
}

type GroupsPropertyValuesRetrieveRequest struct {
	PathParams  GroupsPropertyValuesRetrievePathParams
	QueryParams GroupsPropertyValuesRetrieveQueryParams
}

type GroupsPropertyValuesRetrieveResponse struct {
	ContentType string
	Group       *shared.Group
	StatusCode  int
}
