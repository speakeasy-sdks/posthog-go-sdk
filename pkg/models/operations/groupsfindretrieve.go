package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type GroupsFindRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type GroupsFindRetrieveQueryParams struct {
	GroupKey       string `queryParam:"style=form,explode=true,name=group_key"`
	GroupTypeIndex int64  `queryParam:"style=form,explode=true,name=group_type_index"`
}

type GroupsFindRetrieveRequest struct {
	PathParams  GroupsFindRetrievePathParams
	QueryParams GroupsFindRetrieveQueryParams
}

type GroupsFindRetrieveResponse struct {
	ContentType string
	Group       *shared.Group
	StatusCode  int64
}
