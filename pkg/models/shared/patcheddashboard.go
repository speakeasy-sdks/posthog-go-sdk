package shared

// PatchedDashboardInput
// Serializer mixin that resolves appropriate response for tags depending on license.
type PatchedDashboardInput struct {
	DeleteInsights   *bool                  `json:"delete_insights,omitempty" form:"name=delete_insights" multipartForm:"name=delete_insights"`
	Deleted          *bool                  `json:"deleted,omitempty" form:"name=deleted" multipartForm:"name=deleted"`
	Description      *string                `json:"description,omitempty" form:"name=description" multipartForm:"name=description"`
	Filters          map[string]interface{} `json:"filters,omitempty" form:"name=filters,json" multipartForm:"name=filters,json"`
	Name             *string                `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
	Pinned           *bool                  `json:"pinned,omitempty" form:"name=pinned" multipartForm:"name=pinned"`
	RestrictionLevel *int64                 `json:"restriction_level,omitempty" form:"name=restriction_level" multipartForm:"name=restriction_level"`
	Tags             []interface{}          `json:"tags,omitempty" form:"name=tags,json" multipartForm:"name=tags,json"`
	UseDashboard     *int64                 `json:"use_dashboard,omitempty" form:"name=use_dashboard" multipartForm:"name=use_dashboard"`
	UseTemplate      *string                `json:"use_template,omitempty" form:"name=use_template" multipartForm:"name=use_template"`
}
