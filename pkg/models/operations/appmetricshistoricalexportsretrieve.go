package operations

type AppMetricsHistoricalExportsRetrievePathParams struct {
	ParentLookupPluginConfigID string `pathParam:"style=simple,explode=false,name=parent_lookup_plugin_config_id"`
	ProjectID                  string `pathParam:"style=simple,explode=false,name=project_id"`
}

type AppMetricsHistoricalExportsRetrieveRequest struct {
	PathParams AppMetricsHistoricalExportsRetrievePathParams
}

type AppMetricsHistoricalExportsRetrieveResponse struct {
	ContentType string
	StatusCode  int64
}
