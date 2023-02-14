package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginsSourceRetrievePathParams struct {
	ID                         int64  `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type PluginsSourceRetrieveRequest struct {
	PathParams PluginsSourceRetrievePathParams
}

type PluginsSourceRetrieveResponse struct {
	ContentType string
	Plugin      *shared.Plugin
	StatusCode  int64
}
