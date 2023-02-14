package operations

type RolesDestroyPathParams struct {
	ID                         string `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type RolesDestroyRequest struct {
	PathParams RolesDestroyPathParams
}

type RolesDestroyResponse struct {
	ContentType string
	StatusCode  int64
}
