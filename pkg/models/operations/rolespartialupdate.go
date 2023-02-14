package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type RolesPartialUpdatePathParams struct {
	ID                         string `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type RolesPartialUpdateRequests struct {
	PatchedRole  *shared.PatchedRoleInput `request:"mediaType=application/json"`
	PatchedRole1 *shared.PatchedRoleInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedRole2 *shared.PatchedRoleInput `request:"mediaType=multipart/form-data"`
}

type RolesPartialUpdateRequest struct {
	PathParams RolesPartialUpdatePathParams
	Request    *RolesPartialUpdateRequests
}

type RolesPartialUpdateResponse struct {
	ContentType string
	Role        *shared.Role
	StatusCode  int64
}
