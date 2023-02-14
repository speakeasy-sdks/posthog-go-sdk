package shared

type PaginatedExperimentList struct {
	Count    *int64       `json:"count,omitempty"`
	Next     *string      `json:"next,omitempty"`
	Previous *string      `json:"previous,omitempty"`
	Results  []Experiment `json:"results,omitempty"`
}
