package shared

type PaginatedRoleList struct {
	Count    *int64  `json:"count,omitempty"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  []Role  `json:"results,omitempty"`
}
