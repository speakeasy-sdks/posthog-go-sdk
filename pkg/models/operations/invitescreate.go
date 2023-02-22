package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InvitesCreatePathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type InvitesCreateRequests struct {
	OrganizationInvite  *shared.OrganizationInviteInput `request:"mediaType=application/json"`
	OrganizationInvite1 *shared.OrganizationInviteInput `request:"mediaType=application/x-www-form-urlencoded"`
	OrganizationInvite2 *shared.OrganizationInviteInput `request:"mediaType=multipart/form-data"`
}

type InvitesCreateRequest struct {
	PathParams InvitesCreatePathParams
	Request    InvitesCreateRequests
}

type InvitesCreateResponse struct {
	ContentType        string
	OrganizationInvite *shared.OrganizationInvite
	StatusCode         int
}
