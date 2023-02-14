package shared

type PluginPluginTypeEnum string

const (
	PluginPluginTypeEnumLocal      PluginPluginTypeEnum = "local"
	PluginPluginTypeEnumCustom     PluginPluginTypeEnum = "custom"
	PluginPluginTypeEnumRepository PluginPluginTypeEnum = "repository"
	PluginPluginTypeEnumSource     PluginPluginTypeEnum = "source"
	PluginPluginTypeEnumUnknown    PluginPluginTypeEnum = ""
	PluginPluginTypeEnumNull       PluginPluginTypeEnum = "null"
)

type PluginInput struct {
	Capabilities map[string]interface{} `json:"capabilities,omitempty" form:"name=capabilities,json" multipartForm:"name=capabilities,json"`
	ConfigSchema map[string]interface{} `json:"config_schema,omitempty" form:"name=config_schema,json" multipartForm:"name=config_schema,json"`
	Description  *string                `json:"description,omitempty" form:"name=description" multipartForm:"name=description"`
	Icon         *string                `json:"icon,omitempty" form:"name=icon" multipartForm:"name=icon"`
	IsGlobal     *bool                  `json:"is_global,omitempty" form:"name=is_global" multipartForm:"name=is_global"`
	Metrics      map[string]interface{} `json:"metrics,omitempty" form:"name=metrics,json" multipartForm:"name=metrics,json"`
	Name         *string                `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
	PluginType   *PluginPluginTypeEnum  `json:"plugin_type,omitempty" form:"name=plugin_type" multipartForm:"name=plugin_type"`
	PublicJobs   map[string]interface{} `json:"public_jobs,omitempty" form:"name=public_jobs,json" multipartForm:"name=public_jobs,json"`
	Tag          *string                `json:"tag,omitempty" form:"name=tag" multipartForm:"name=tag"`
}

type Plugin struct {
	Capabilities     map[string]interface{} `json:"capabilities,omitempty"`
	ConfigSchema     map[string]interface{} `json:"config_schema,omitempty"`
	Description      *string                `json:"description,omitempty"`
	Icon             *string                `json:"icon,omitempty"`
	ID               int64                  `json:"id"`
	IsGlobal         *bool                  `json:"is_global,omitempty"`
	LatestTag        string                 `json:"latest_tag"`
	Metrics          map[string]interface{} `json:"metrics,omitempty"`
	Name             *string                `json:"name,omitempty"`
	OrganizationID   string                 `json:"organization_id"`
	OrganizationName string                 `json:"organization_name"`
	PluginType       *PluginPluginTypeEnum  `json:"plugin_type,omitempty"`
	PublicJobs       map[string]interface{} `json:"public_jobs,omitempty"`
	Tag              *string                `json:"tag,omitempty"`
	URL              string                 `json:"url"`
}
