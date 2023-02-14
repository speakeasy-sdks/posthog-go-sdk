package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type RolesRetrievePathParams struct {
	ID                         string `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type RolesRetrieveRequest struct {
	PathParams RolesRetrievePathParams
}

type RolesRetrieveResponse struct {
	ContentType string
	Role        *shared.Role
	StatusCode  int64
}
