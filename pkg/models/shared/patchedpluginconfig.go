package shared

type PatchedPluginConfigInput struct {
	Enabled *bool                  `json:"enabled,omitempty" form:"name=enabled" multipartForm:"name=enabled"`
	Error   map[string]interface{} `json:"error,omitempty" form:"name=error,json" multipartForm:"name=error,json"`
	Order   *int64                 `json:"order,omitempty" form:"name=order" multipartForm:"name=order"`
	Plugin  *int64                 `json:"plugin,omitempty" form:"name=plugin" multipartForm:"name=plugin"`
}
