package shared

type PatchedCohortInput struct {
	Deleted     *bool                  `json:"deleted,omitempty" form:"name=deleted" multipartForm:"name=deleted"`
	Description *string                `json:"description,omitempty" form:"name=description" multipartForm:"name=description"`
	Filters     map[string]interface{} `json:"filters,omitempty" form:"name=filters,json" multipartForm:"name=filters,json"`
	Groups      map[string]interface{} `json:"groups,omitempty" form:"name=groups,json" multipartForm:"name=groups,json"`
	IsStatic    *bool                  `json:"is_static,omitempty" form:"name=is_static" multipartForm:"name=is_static"`
	Name        *string                `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
}
