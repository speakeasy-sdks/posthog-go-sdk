package shared

type PaginatedEventDefinitionList struct {
	Count    *int64            `json:"count,omitempty"`
	Next     *string           `json:"next,omitempty"`
	Previous *string           `json:"previous,omitempty"`
	Results  []EventDefinition `json:"results,omitempty"`
}
