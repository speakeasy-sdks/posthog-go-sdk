package operations

type MembersDestroyPathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
	UserUUID                   string `pathParam:"style=simple,explode=false,name=user__uuid"`
}

type MembersDestroyRequest struct {
	PathParams MembersDestroyPathParams
}

type MembersDestroyResponse struct {
	ContentType string
	StatusCode  int64
}
