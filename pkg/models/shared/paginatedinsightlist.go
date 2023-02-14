package shared

type PaginatedInsightList struct {
	Count    *int64    `json:"count,omitempty"`
	Next     *string   `json:"next,omitempty"`
	Previous *string   `json:"previous,omitempty"`
	Results  []Insight `json:"results,omitempty"`
}
