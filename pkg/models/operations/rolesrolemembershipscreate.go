package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type RolesRoleMembershipsCreatePathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
	ParentLookupRoleID         string `pathParam:"style=simple,explode=false,name=parent_lookup_role_id"`
}

type RolesRoleMembershipsCreateRequests struct {
	RoleMembership  *shared.RoleMembershipInput `request:"mediaType=application/json"`
	RoleMembership1 *shared.RoleMembershipInput `request:"mediaType=application/x-www-form-urlencoded"`
	RoleMembership2 *shared.RoleMembershipInput `request:"mediaType=multipart/form-data"`
}

type RolesRoleMembershipsCreateRequest struct {
	PathParams RolesRoleMembershipsCreatePathParams
	Request    RolesRoleMembershipsCreateRequests
}

type RolesRoleMembershipsCreateResponse struct {
	ContentType    string
	RoleMembership *shared.RoleMembershipOutput
	StatusCode     int64
}
