package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginsActivityRetrievePathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type PluginsActivityRetrieveRequest struct {
	PathParams PluginsActivityRetrievePathParams
}

type PluginsActivityRetrieveResponse struct {
	ContentType string
	Plugin      *shared.Plugin
	StatusCode  int
}
