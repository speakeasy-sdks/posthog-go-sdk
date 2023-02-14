package shared

type PaginatedPluginConfigList struct {
	Count    *int64         `json:"count,omitempty"`
	Next     *string        `json:"next,omitempty"`
	Previous *string        `json:"previous,omitempty"`
	Results  []PluginConfig `json:"results,omitempty"`
}
