package shared

type PaginatedPluginList struct {
	Count    *int64   `json:"count,omitempty"`
	Next     *string  `json:"next,omitempty"`
	Previous *string  `json:"previous,omitempty"`
	Results  []Plugin `json:"results,omitempty"`
}
