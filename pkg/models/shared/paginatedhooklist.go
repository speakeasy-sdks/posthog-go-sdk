package shared

type PaginatedHookList struct {
	Count    *int64  `json:"count,omitempty"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  []Hook  `json:"results,omitempty"`
}
