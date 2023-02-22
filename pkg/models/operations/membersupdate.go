package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type MembersUpdatePathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
	UserUUID                   string `pathParam:"style=simple,explode=false,name=user__uuid"`
}

type MembersUpdateRequests struct {
	OrganizationMember  *shared.OrganizationMemberInput `request:"mediaType=application/json"`
	OrganizationMember1 *shared.OrganizationMemberInput `request:"mediaType=application/x-www-form-urlencoded"`
	OrganizationMember2 *shared.OrganizationMemberInput `request:"mediaType=multipart/form-data"`
}

type MembersUpdateRequest struct {
	PathParams MembersUpdatePathParams
	Request    *MembersUpdateRequests
}

type MembersUpdateResponse struct {
	ContentType        string
	OrganizationMember *shared.OrganizationMember
	StatusCode         int
}
