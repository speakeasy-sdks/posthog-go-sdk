package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type ActivityLogBookmarkActivityNotificationCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type ActivityLogBookmarkActivityNotificationCreateRequests struct {
	ActivityLog  *shared.ActivityLogInput `request:"mediaType=application/json"`
	ActivityLog1 *shared.ActivityLogInput `request:"mediaType=application/x-www-form-urlencoded"`
	ActivityLog2 *shared.ActivityLogInput `request:"mediaType=multipart/form-data"`
}

type ActivityLogBookmarkActivityNotificationCreateRequest struct {
	PathParams ActivityLogBookmarkActivityNotificationCreatePathParams
	Request    ActivityLogBookmarkActivityNotificationCreateRequests
}

type ActivityLogBookmarkActivityNotificationCreateResponse struct {
	ActivityLog *shared.ActivityLog
	ContentType string
	StatusCode  int64
}
