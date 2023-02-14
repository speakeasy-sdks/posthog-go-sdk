package shared

import (
	"time"
)

type SessionRecordingPlaylistInput struct {
	Deleted     *bool                  `json:"deleted,omitempty" form:"name=deleted" multipartForm:"name=deleted"`
	DerivedName *string                `json:"derived_name,omitempty" form:"name=derived_name" multipartForm:"name=derived_name"`
	Description *string                `json:"description,omitempty" form:"name=description" multipartForm:"name=description"`
	Filters     map[string]interface{} `json:"filters,omitempty" form:"name=filters,json" multipartForm:"name=filters,json"`
	Name        *string                `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
	Pinned      *bool                  `json:"pinned,omitempty" form:"name=pinned" multipartForm:"name=pinned"`
}

type SessionRecordingPlaylistCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

type SessionRecordingPlaylistLastModifiedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

type SessionRecordingPlaylist struct {
	CreatedAt      time.Time                              `json:"created_at"`
	CreatedBy      SessionRecordingPlaylistCreatedBy      `json:"created_by"`
	Deleted        *bool                                  `json:"deleted,omitempty"`
	DerivedName    *string                                `json:"derived_name,omitempty"`
	Description    *string                                `json:"description,omitempty"`
	Filters        map[string]interface{}                 `json:"filters,omitempty"`
	ID             int64                                  `json:"id"`
	LastModifiedAt time.Time                              `json:"last_modified_at"`
	LastModifiedBy SessionRecordingPlaylistLastModifiedBy `json:"last_modified_by"`
	Name           *string                                `json:"name,omitempty"`
	Pinned         *bool                                  `json:"pinned,omitempty"`
	ShortID        string                                 `json:"short_id"`
}
