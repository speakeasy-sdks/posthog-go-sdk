package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type AnnotationsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type AnnotationsListQueryParams struct {
	Limit  *int64  `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64  `queryParam:"style=form,explode=true,name=offset"`
	Search *string `queryParam:"style=form,explode=true,name=search"`
}

type AnnotationsListRequest struct {
	PathParams  AnnotationsListPathParams
	QueryParams AnnotationsListQueryParams
}

type AnnotationsListResponse struct {
	ContentType             string
	PaginatedAnnotationList *shared.PaginatedAnnotationList
	StatusCode              int64
}
