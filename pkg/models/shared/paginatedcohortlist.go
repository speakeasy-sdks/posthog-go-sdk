package shared

type PaginatedCohortList struct {
	Count    *int64   `json:"count,omitempty"`
	Next     *string  `json:"next,omitempty"`
	Previous *string  `json:"previous,omitempty"`
	Results  []Cohort `json:"results,omitempty"`
}
