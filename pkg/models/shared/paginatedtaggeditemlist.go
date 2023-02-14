package shared

type PaginatedTaggedItemList struct {
	Count    *int64       `json:"count,omitempty"`
	Next     *string      `json:"next,omitempty"`
	Previous *string      `json:"previous,omitempty"`
	Results  []TaggedItem `json:"results,omitempty"`
}
