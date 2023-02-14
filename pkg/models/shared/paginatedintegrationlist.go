package shared

type PaginatedIntegrationList struct {
	Count    *int64        `json:"count,omitempty"`
	Next     *string       `json:"next,omitempty"`
	Previous *string       `json:"previous,omitempty"`
	Results  []Integration `json:"results,omitempty"`
}
