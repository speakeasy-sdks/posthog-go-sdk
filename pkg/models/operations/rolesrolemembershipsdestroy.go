package operations

type RolesRoleMembershipsDestroyPathParams struct {
	ID                         string `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
	ParentLookupRoleID         string `pathParam:"style=simple,explode=false,name=parent_lookup_role_id"`
}

type RolesRoleMembershipsDestroyRequest struct {
	PathParams RolesRoleMembershipsDestroyPathParams
}

type RolesRoleMembershipsDestroyResponse struct {
	ContentType string
	StatusCode  int
}
