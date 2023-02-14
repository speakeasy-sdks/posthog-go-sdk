package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type RolesUpdatePathParams struct {
	ID                         string `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type RolesUpdateRequests struct {
	Role  *shared.RoleInput `request:"mediaType=application/json"`
	Role1 *shared.RoleInput `request:"mediaType=application/x-www-form-urlencoded"`
	Role2 *shared.RoleInput `request:"mediaType=multipart/form-data"`
}

type RolesUpdateRequest struct {
	PathParams RolesUpdatePathParams
	Request    RolesUpdateRequests
}

type RolesUpdateResponse struct {
	ContentType string
	Role        *shared.Role
	StatusCode  int64
}
