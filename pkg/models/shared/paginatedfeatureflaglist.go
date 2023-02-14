package shared

type PaginatedFeatureFlagList struct {
	Count    *int64        `json:"count,omitempty"`
	Next     *string       `json:"next,omitempty"`
	Previous *string       `json:"previous,omitempty"`
	Results  []FeatureFlag `json:"results,omitempty"`
}
