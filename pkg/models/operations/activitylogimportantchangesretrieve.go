package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ActivityLogImportantChangesRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ActivityLogImportantChangesRetrieveRequest struct {
	PathParams ActivityLogImportantChangesRetrievePathParams
}

type ActivityLogImportantChangesRetrieveResponse struct {
	ActivityLog *shared.ActivityLog
	ContentType string
	StatusCode  int64
}
