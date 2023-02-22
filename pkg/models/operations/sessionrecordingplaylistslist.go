package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type SessionRecordingPlaylistsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type SessionRecordingPlaylistsListQueryParams struct {
	CreatedBy *int64  `queryParam:"style=form,explode=true,name=created_by"`
	Limit     *int64  `queryParam:"style=form,explode=true,name=limit"`
	Offset    *int64  `queryParam:"style=form,explode=true,name=offset"`
	ShortID   *string `queryParam:"style=form,explode=true,name=short_id"`
}

type SessionRecordingPlaylistsListRequest struct {
	PathParams  SessionRecordingPlaylistsListPathParams
	QueryParams SessionRecordingPlaylistsListQueryParams
}

type SessionRecordingPlaylistsListResponse struct {
	ContentType                           string
	PaginatedSessionRecordingPlaylistList *shared.PaginatedSessionRecordingPlaylistList
	StatusCode                            int
}
