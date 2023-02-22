package operations

type AppMetricsHistoricalExportsRetrieve2PathParams struct {
	ID                         string `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupPluginConfigID string `pathParam:"style=simple,explode=false,name=parent_lookup_plugin_config_id"`
	ProjectID                  string `pathParam:"style=simple,explode=false,name=project_id"`
}

type AppMetricsHistoricalExportsRetrieve2Request struct {
	PathParams AppMetricsHistoricalExportsRetrieve2PathParams
}

type AppMetricsHistoricalExportsRetrieve2Response struct {
	ContentType string
	StatusCode  int
}
