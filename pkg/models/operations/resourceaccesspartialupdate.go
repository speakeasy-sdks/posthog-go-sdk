package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ResourceAccessPartialUpdatePathParams struct {
	ID                         int64  `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type ResourceAccessPartialUpdateRequests struct {
	PatchedOrganizationResourceAccess  *shared.PatchedOrganizationResourceAccessInput `request:"mediaType=application/json"`
	PatchedOrganizationResourceAccess1 *shared.PatchedOrganizationResourceAccessInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedOrganizationResourceAccess2 *shared.PatchedOrganizationResourceAccessInput `request:"mediaType=multipart/form-data"`
}

type ResourceAccessPartialUpdateRequest struct {
	PathParams ResourceAccessPartialUpdatePathParams
	Request    *ResourceAccessPartialUpdateRequests
}

type ResourceAccessPartialUpdateResponse struct {
	ContentType                string
	OrganizationResourceAccess *shared.OrganizationResourceAccess
	StatusCode                 int
}
