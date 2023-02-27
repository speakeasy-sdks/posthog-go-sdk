package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type RolesCreatePathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type RolesCreateRequests struct {
	Role  *shared.RoleInput `request:"mediaType=application/json"`
	Role1 *shared.RoleInput `request:"mediaType=application/x-www-form-urlencoded"`
	Role2 *shared.RoleInput `request:"mediaType=multipart/form-data"`
}

type RolesCreateRequest struct {
	PathParams RolesCreatePathParams
	Request    RolesCreateRequests
}

type RolesCreateResponse struct {
	ContentType string
	Role        *shared.Role
	StatusCode  int
}
