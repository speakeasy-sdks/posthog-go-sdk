package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type InvitesBulkCreatePathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type InvitesBulkCreateRequests struct {
	OrganizationInvite  *shared.OrganizationInviteInput `request:"mediaType=application/json"`
	OrganizationInvite1 *shared.OrganizationInviteInput `request:"mediaType=application/x-www-form-urlencoded"`
	OrganizationInvite2 *shared.OrganizationInviteInput `request:"mediaType=multipart/form-data"`
}

type InvitesBulkCreateRequest struct {
	PathParams InvitesBulkCreatePathParams
	Request    InvitesBulkCreateRequests
}

type InvitesBulkCreateResponse struct {
	ContentType        string
	OrganizationInvite *shared.OrganizationInvite
	StatusCode         int
}
