package shared

type PaginatedGroupList struct {
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  []Group `json:"results,omitempty"`
}
