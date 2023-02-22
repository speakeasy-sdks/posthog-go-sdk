package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ResourceAccessRetrievePathParams struct {
	ID                         int64  `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type ResourceAccessRetrieveRequest struct {
	PathParams ResourceAccessRetrievePathParams
}

type ResourceAccessRetrieveResponse struct {
	ContentType                string
	OrganizationResourceAccess *shared.OrganizationResourceAccess
	StatusCode                 int
}
