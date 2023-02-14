package shared

type PaginatedOrganizationResourceAccessList struct {
	Count    *int64                       `json:"count,omitempty"`
	Next     *string                      `json:"next,omitempty"`
	Previous *string                      `json:"previous,omitempty"`
	Results  []OrganizationResourceAccess `json:"results,omitempty"`
}
