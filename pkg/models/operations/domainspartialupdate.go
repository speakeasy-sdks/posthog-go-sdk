package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type DomainsPartialUpdatePathParams struct {
	ID                         string `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type DomainsPartialUpdateRequests struct {
	PatchedOrganizationDomain  *shared.PatchedOrganizationDomainInput `request:"mediaType=application/json"`
	PatchedOrganizationDomain1 *shared.PatchedOrganizationDomainInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedOrganizationDomain2 *shared.PatchedOrganizationDomainInput `request:"mediaType=multipart/form-data"`
}

type DomainsPartialUpdateRequest struct {
	PathParams DomainsPartialUpdatePathParams
	Request    *DomainsPartialUpdateRequests
}

type DomainsPartialUpdateResponse struct {
	ContentType        string
	OrganizationDomain *shared.OrganizationDomain
	StatusCode         int
}
