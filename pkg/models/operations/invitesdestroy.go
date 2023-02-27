package operations

type InvitesDestroyPathParams struct {
	ID                         string `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type InvitesDestroyRequest struct {
	PathParams InvitesDestroyPathParams
}

type InvitesDestroyResponse struct {
	ContentType string
	StatusCode  int
}
