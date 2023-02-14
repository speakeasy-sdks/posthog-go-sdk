package shared

type PatchedSessionRecordingPlaylistInput struct {
	Deleted     *bool                  `json:"deleted,omitempty" form:"name=deleted" multipartForm:"name=deleted"`
	DerivedName *string                `json:"derived_name,omitempty" form:"name=derived_name" multipartForm:"name=derived_name"`
	Description *string                `json:"description,omitempty" form:"name=description" multipartForm:"name=description"`
	Filters     map[string]interface{} `json:"filters,omitempty" form:"name=filters,json" multipartForm:"name=filters,json"`
	Name        *string                `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
	Pinned      *bool                  `json:"pinned,omitempty" form:"name=pinned" multipartForm:"name=pinned"`
}
