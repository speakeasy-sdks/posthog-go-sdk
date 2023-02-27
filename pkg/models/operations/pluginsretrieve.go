package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginsRetrievePathParams struct {
	ID                         int64  `pathParam:"style=simple,explode=false,name=id"`
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type PluginsRetrieveRequest struct {
	PathParams PluginsRetrievePathParams
}

type PluginsRetrieveResponse struct {
	ContentType string
	Plugin      *shared.Plugin
	StatusCode  int
}
