package shared

type PatchedPluginPluginTypeEnum string

const (
	PatchedPluginPluginTypeEnumLocal      PatchedPluginPluginTypeEnum = "local"
	PatchedPluginPluginTypeEnumCustom     PatchedPluginPluginTypeEnum = "custom"
	PatchedPluginPluginTypeEnumRepository PatchedPluginPluginTypeEnum = "repository"
	PatchedPluginPluginTypeEnumSource     PatchedPluginPluginTypeEnum = "source"
	PatchedPluginPluginTypeEnumUnknown    PatchedPluginPluginTypeEnum = ""
	PatchedPluginPluginTypeEnumNull       PatchedPluginPluginTypeEnum = "null"
)

type PatchedPluginInput struct {
	Capabilities map[string]interface{}       `json:"capabilities,omitempty" form:"name=capabilities,json" multipartForm:"name=capabilities,json"`
	ConfigSchema map[string]interface{}       `json:"config_schema,omitempty" form:"name=config_schema,json" multipartForm:"name=config_schema,json"`
	Description  *string                      `json:"description,omitempty" form:"name=description" multipartForm:"name=description"`
	Icon         *string                      `json:"icon,omitempty" form:"name=icon" multipartForm:"name=icon"`
	IsGlobal     *bool                        `json:"is_global,omitempty" form:"name=is_global" multipartForm:"name=is_global"`
	Metrics      map[string]interface{}       `json:"metrics,omitempty" form:"name=metrics,json" multipartForm:"name=metrics,json"`
	Name         *string                      `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
	PluginType   *PatchedPluginPluginTypeEnum `json:"plugin_type,omitempty" form:"name=plugin_type" multipartForm:"name=plugin_type"`
	PublicJobs   map[string]interface{}       `json:"public_jobs,omitempty" form:"name=public_jobs,json" multipartForm:"name=public_jobs,json"`
	Tag          *string                      `json:"tag,omitempty" form:"name=tag" multipartForm:"name=tag"`
}
