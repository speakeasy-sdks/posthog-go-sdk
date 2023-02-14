package shared

type PaginatedFeatureFlagRoleAccessList struct {
	Count    *int64                        `json:"count,omitempty"`
	Next     *string                       `json:"next,omitempty"`
	Previous *string                       `json:"previous,omitempty"`
	Results  []FeatureFlagRoleAccessOutput `json:"results,omitempty"`
}
