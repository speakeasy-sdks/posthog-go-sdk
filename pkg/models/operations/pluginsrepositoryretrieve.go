package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PluginsRepositoryRetrievePathParams struct {
	ParentLookupOrganizationID string `pathParam:"style=simple,explode=false,name=parent_lookup_organization_id"`
}

type PluginsRepositoryRetrieveRequest struct {
	PathParams PluginsRepositoryRetrievePathParams
}

type PluginsRepositoryRetrieveResponse struct {
	ContentType string
	Plugin      *shared.Plugin
	StatusCode  int64
}
