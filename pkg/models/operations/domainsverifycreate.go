package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type DomainsVerifyCreatePathParams struct {
	ID                         string `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type DomainsVerifyCreateRequests struct {
	OrganizationDomain  *shared.OrganizationDomainInput `request:"mediaType=application/json"`
	OrganizationDomain1 *shared.OrganizationDomainInput `request:"mediaType=application/x-www-form-urlencoded"`
	OrganizationDomain2 *shared.OrganizationDomainInput `request:"mediaType=multipart/form-data"`
}

type DomainsVerifyCreateRequest struct {
	PathParams DomainsVerifyCreatePathParams
	Request    DomainsVerifyCreateRequests
}

type DomainsVerifyCreateResponse struct {
	ContentType        string
	OrganizationDomain *shared.OrganizationDomain
	StatusCode         int64
}
