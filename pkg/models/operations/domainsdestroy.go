package operations

type DomainsDestroyPathParams struct {
	ID                         string `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type DomainsDestroyRequest struct {
	PathParams DomainsDestroyPathParams
}

type DomainsDestroyResponse struct {
	ContentType string
	StatusCode  int64
}
