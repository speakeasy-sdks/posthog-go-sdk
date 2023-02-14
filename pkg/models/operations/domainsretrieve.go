package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type DomainsRetrievePathParams struct {
	ID                         string `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type DomainsRetrieveRequest struct {
	PathParams DomainsRetrievePathParams
}

type DomainsRetrieveResponse struct {
	ContentType        string
	OrganizationDomain *shared.OrganizationDomain
	StatusCode         int64
}
