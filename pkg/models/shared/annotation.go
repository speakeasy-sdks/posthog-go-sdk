package shared

import (
	"time"
)

type AnnotationCreationTypeEnum string

const (
	AnnotationCreationTypeEnumUsr AnnotationCreationTypeEnum = "USR"
	AnnotationCreationTypeEnumGit AnnotationCreationTypeEnum = "GIT"
)

type AnnotationScopeEnum string

const (
	AnnotationScopeEnumDashboardItem AnnotationScopeEnum = "dashboard_item"
	AnnotationScopeEnumProject       AnnotationScopeEnum = "project"
	AnnotationScopeEnumOrganization  AnnotationScopeEnum = "organization"
)

type AnnotationInput struct {
	Content       *string                     `json:"content,omitempty" form:"name=content" multipartForm:"name=content"`
	CreationType  *AnnotationCreationTypeEnum `json:"creation_type,omitempty" form:"name=creation_type" multipartForm:"name=creation_type"`
	DashboardItem *int64                      `json:"dashboard_item,omitempty" form:"name=dashboard_item" multipartForm:"name=dashboard_item"`
	DateMarker    *time.Time                  `json:"date_marker,omitempty" form:"name=date_marker" multipartForm:"name=date_marker"`
	Deleted       *bool                       `json:"deleted,omitempty" form:"name=deleted" multipartForm:"name=deleted"`
	Scope         *AnnotationScopeEnum        `json:"scope,omitempty" form:"name=scope" multipartForm:"name=scope"`
}

type AnnotationCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

type Annotation struct {
	Content        *string                     `json:"content,omitempty"`
	CreatedAt      time.Time                   `json:"created_at"`
	CreatedBy      AnnotationCreatedBy         `json:"created_by"`
	CreationType   *AnnotationCreationTypeEnum `json:"creation_type,omitempty"`
	DashboardItem  *int64                      `json:"dashboard_item,omitempty"`
	DateMarker     *time.Time                  `json:"date_marker,omitempty"`
	Deleted        *bool                       `json:"deleted,omitempty"`
	ID             int64                       `json:"id"`
	InsightName    string                      `json:"insight_name"`
	InsightShortID string                      `json:"insight_short_id"`
	Scope          *AnnotationScopeEnum        `json:"scope,omitempty"`
	UpdatedAt      time.Time                   `json:"updated_at"`
}
