package operations

type ResourceAccessDestroyPathParams struct {
	ID                         int64  `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type ResourceAccessDestroyRequest struct {
	PathParams ResourceAccessDestroyPathParams
}

type ResourceAccessDestroyResponse struct {
	ContentType string
	StatusCode  int64
}
