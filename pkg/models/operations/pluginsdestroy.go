package operations

type PluginsDestroyPathParams struct {
	ID                         int64  `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type PluginsDestroyRequest struct {
	PathParams PluginsDestroyPathParams
}

type PluginsDestroyResponse struct {
	ContentType string
	StatusCode  int
}
