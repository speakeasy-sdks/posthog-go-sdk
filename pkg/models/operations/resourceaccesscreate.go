package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ResourceAccessCreatePathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type ResourceAccessCreateRequests struct {
	OrganizationResourceAccess  *shared.OrganizationResourceAccessInput `request:"mediaType=application/json"`
	OrganizationResourceAccess1 *shared.OrganizationResourceAccessInput `request:"mediaType=application/x-www-form-urlencoded"`
	OrganizationResourceAccess2 *shared.OrganizationResourceAccessInput `request:"mediaType=multipart/form-data"`
}

type ResourceAccessCreateRequest struct {
	PathParams ResourceAccessCreatePathParams
	Request    ResourceAccessCreateRequests
}

type ResourceAccessCreateResponse struct {
	ContentType                string
	OrganizationResourceAccess *shared.OrganizationResourceAccess
	StatusCode                 int64
}
