package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type MembersPartialUpdatePathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
	UserUUID                   string `pathParam:"style=simple,explode=false,name=user__uuid"`
}

type MembersPartialUpdateRequests struct {
	PatchedOrganizationMember  *shared.PatchedOrganizationMemberInput `request:"mediaType=application/json"`
	PatchedOrganizationMember1 *shared.PatchedOrganizationMemberInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedOrganizationMember2 *shared.PatchedOrganizationMemberInput `request:"mediaType=multipart/form-data"`
}

type MembersPartialUpdateRequest struct {
	PathParams MembersPartialUpdatePathParams
	Request    *MembersPartialUpdateRequests
}

type MembersPartialUpdateResponse struct {
	ContentType        string
	OrganizationMember *shared.OrganizationMember
	StatusCode         int
}
