package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ResourceAccessUpdatePathParams struct {
	ID                         int64  `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type ResourceAccessUpdateRequests struct {
	OrganizationResourceAccess  *shared.OrganizationResourceAccessInput `request:"mediaType=application/json"`
	OrganizationResourceAccess1 *shared.OrganizationResourceAccessInput `request:"mediaType=application/x-www-form-urlencoded"`
	OrganizationResourceAccess2 *shared.OrganizationResourceAccessInput `request:"mediaType=multipart/form-data"`
}

type ResourceAccessUpdateRequest struct {
	PathParams ResourceAccessUpdatePathParams
	Request    ResourceAccessUpdateRequests
}

type ResourceAccessUpdateResponse struct {
	ContentType                string
	OrganizationResourceAccess *shared.OrganizationResourceAccess
	StatusCode                 int
}
