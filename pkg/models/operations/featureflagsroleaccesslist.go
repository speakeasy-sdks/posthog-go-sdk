package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type FeatureFlagsRoleAccessListPathParams struct {
	ParentLookupFeatureFlagID string `pathParam:"style=simple,explode=false,name=parent_lookup_feature_flag_id"`
	ProjectID                 string `pathParam:"style=simple,explode=false,name=project_id"`
}

type FeatureFlagsRoleAccessListQueryParams struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type FeatureFlagsRoleAccessListRequest struct {
	PathParams  FeatureFlagsRoleAccessListPathParams
	QueryParams FeatureFlagsRoleAccessListQueryParams
}

type FeatureFlagsRoleAccessListResponse struct {
	ContentType                        string
	PaginatedFeatureFlagRoleAccessList *shared.PaginatedFeatureFlagRoleAccessList
	StatusCode                         int64
}
