package shared

// PatchedInsightInput
// Simplified serializer to speed response times when loading large amounts of objects.
type PatchedInsightInput struct {
	Dashboards  []int64                `json:"dashboards,omitempty"`
	Deleted     *bool                  `json:"deleted,omitempty"`
	DerivedName *string                `json:"derived_name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Favorited   *bool                  `json:"favorited,omitempty"`
	Filters     map[string]interface{} `json:"filters,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Order       *int64                 `json:"order,omitempty"`
	Query       map[string]interface{} `json:"query,omitempty"`
	Saved       *bool                  `json:"saved,omitempty"`
	Tags        []interface{}          `json:"tags,omitempty"`
}
