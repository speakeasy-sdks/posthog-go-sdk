package shared

type PaginatedOrganizationInviteList struct {
	Count    *int64               `json:"count,omitempty"`
	Next     *string              `json:"next,omitempty"`
	Previous *string              `json:"previous,omitempty"`
	Results  []OrganizationInvite `json:"results,omitempty"`
}
