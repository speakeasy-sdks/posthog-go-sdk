package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type GroupsTypesUpdateMetadataPartialUpdatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type GroupsTypesUpdateMetadataPartialUpdateRequests struct {
	PatchedGroupType  *shared.PatchedGroupTypeInput `request:"mediaType=application/json"`
	PatchedGroupType1 *shared.PatchedGroupTypeInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedGroupType2 *shared.PatchedGroupTypeInput `request:"mediaType=multipart/form-data"`
}

type GroupsTypesUpdateMetadataPartialUpdateRequest struct {
	PathParams GroupsTypesUpdateMetadataPartialUpdatePathParams
	Request    *GroupsTypesUpdateMetadataPartialUpdateRequests
}

type GroupsTypesUpdateMetadataPartialUpdateResponse struct {
	ContentType string
	GroupType   *shared.GroupType
	StatusCode  int
}
