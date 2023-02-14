package shared

type PaginatedPluginLogEntryList struct {
	Count    *int64           `json:"count,omitempty"`
	Next     *string          `json:"next,omitempty"`
	Previous *string          `json:"previous,omitempty"`
	Results  []PluginLogEntry `json:"results,omitempty"`
}
