package shared

type PaginatedPropertyDefinitionList struct {
	Count    *int64               `json:"count,omitempty"`
	Next     *string              `json:"next,omitempty"`
	Previous *string              `json:"previous,omitempty"`
	Results  []PropertyDefinition `json:"results,omitempty"`
}
