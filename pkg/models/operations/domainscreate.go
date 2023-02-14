package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type DomainsCreatePathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type DomainsCreateRequests struct {
	OrganizationDomain  *shared.OrganizationDomainInput `request:"mediaType=application/json"`
	OrganizationDomain1 *shared.OrganizationDomainInput `request:"mediaType=application/x-www-form-urlencoded"`
	OrganizationDomain2 *shared.OrganizationDomainInput `request:"mediaType=multipart/form-data"`
}

type DomainsCreateRequest struct {
	PathParams DomainsCreatePathParams
	Request    DomainsCreateRequests
}

type DomainsCreateResponse struct {
	ContentType        string
	OrganizationDomain *shared.OrganizationDomain
	StatusCode         int64
}
