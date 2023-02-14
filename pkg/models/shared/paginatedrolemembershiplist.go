package shared

type PaginatedRoleMembershipList struct {
	Count    *int64                 `json:"count,omitempty"`
	Next     *string                `json:"next,omitempty"`
	Previous *string                `json:"previous,omitempty"`
	Results  []RoleMembershipOutput `json:"results,omitempty"`
}
