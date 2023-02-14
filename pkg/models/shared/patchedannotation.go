package shared

import (
	"time"
)

type PatchedAnnotationCreationTypeEnum string

const (
	PatchedAnnotationCreationTypeEnumUsr PatchedAnnotationCreationTypeEnum = "USR"
	PatchedAnnotationCreationTypeEnumGit PatchedAnnotationCreationTypeEnum = "GIT"
)

type PatchedAnnotationScopeEnum string

const (
	PatchedAnnotationScopeEnumDashboardItem PatchedAnnotationScopeEnum = "dashboard_item"
	PatchedAnnotationScopeEnumProject       PatchedAnnotationScopeEnum = "project"
	PatchedAnnotationScopeEnumOrganization  PatchedAnnotationScopeEnum = "organization"
)

type PatchedAnnotationInput struct {
	Content       *string                            `json:"content,omitempty" form:"name=content" multipartForm:"name=content"`
	CreationType  *PatchedAnnotationCreationTypeEnum `json:"creation_type,omitempty" form:"name=creation_type" multipartForm:"name=creation_type"`
	DashboardItem *int64                             `json:"dashboard_item,omitempty" form:"name=dashboard_item" multipartForm:"name=dashboard_item"`
	DateMarker    *time.Time                         `json:"date_marker,omitempty" form:"name=date_marker" multipartForm:"name=date_marker"`
	Deleted       *bool                              `json:"deleted,omitempty" form:"name=deleted" multipartForm:"name=deleted"`
	Scope         *PatchedAnnotationScopeEnum        `json:"scope,omitempty" form:"name=scope" multipartForm:"name=scope"`
}
