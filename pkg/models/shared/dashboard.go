package shared

import (
	"time"
)

// DashboardInput
// Serializer mixin that resolves appropriate response for tags depending on license.
type DashboardInput struct {
	DeleteInsights   *bool                  `json:"delete_insights,omitempty" form:"name=delete_insights" multipartForm:"name=delete_insights"`
	Deleted          *bool                  `json:"deleted,omitempty" form:"name=deleted" multipartForm:"name=deleted"`
	Description      *string                `json:"description,omitempty" form:"name=description" multipartForm:"name=description"`
	Filters          map[string]interface{} `json:"filters,omitempty" form:"name=filters,json" multipartForm:"name=filters,json"`
	Name             *string                `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
	Pinned           *bool                  `json:"pinned,omitempty" form:"name=pinned" multipartForm:"name=pinned"`
	RestrictionLevel *int64                 `json:"restriction_level,omitempty" form:"name=restriction_level" multipartForm:"name=restriction_level"`
	Tags             []interface{}          `json:"tags,omitempty" form:"name=tags,json" multipartForm:"name=tags,json"`
	UseDashboard     *int64                 `json:"use_dashboard,omitempty" form:"name=use_dashboard" multipartForm:"name=use_dashboard"`
	UseTemplate      *string                `json:"use_template,omitempty" form:"name=use_template" multipartForm:"name=use_template"`
}

type DashboardCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

type DashboardCreationModeEnum string

const (
	DashboardCreationModeEnumDefault   DashboardCreationModeEnum = "default"
	DashboardCreationModeEnumTemplate  DashboardCreationModeEnum = "template"
	DashboardCreationModeEnumDuplicate DashboardCreationModeEnum = "duplicate"
)

// DashboardOutput
// Serializer mixin that resolves appropriate response for tags depending on license.
type DashboardOutput struct {
	CreatedAt                 time.Time                 `json:"created_at"`
	CreatedBy                 DashboardCreatedBy        `json:"created_by"`
	CreationMode              DashboardCreationModeEnum `json:"creation_mode"`
	Deleted                   *bool                     `json:"deleted,omitempty"`
	Description               *string                   `json:"description,omitempty"`
	EffectivePrivilegeLevel   int64                     `json:"effective_privilege_level"`
	EffectiveRestrictionLevel int64                     `json:"effective_restriction_level"`
	Filters                   map[string]interface{}    `json:"filters,omitempty"`
	ID                        int64                     `json:"id"`
	IsShared                  bool                      `json:"is_shared"`
	Name                      *string                   `json:"name,omitempty"`
	Pinned                    *bool                     `json:"pinned,omitempty"`
	RestrictionLevel          *int64                    `json:"restriction_level,omitempty"`
	Tags                      []interface{}             `json:"tags,omitempty"`
	Tiles                     string                    `json:"tiles"`
}
