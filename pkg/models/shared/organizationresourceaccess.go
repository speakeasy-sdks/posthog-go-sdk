package shared

import (
	"time"
)

type OrganizationResourceAccessResourceEnum string

const (
	OrganizationResourceAccessResourceEnumFeatureFlags      OrganizationResourceAccessResourceEnum = "feature flags"
	OrganizationResourceAccessResourceEnumExperiments       OrganizationResourceAccessResourceEnum = "experiments"
	OrganizationResourceAccessResourceEnumCohorts           OrganizationResourceAccessResourceEnum = "cohorts"
	OrganizationResourceAccessResourceEnumDataManagement    OrganizationResourceAccessResourceEnum = "data management"
	OrganizationResourceAccessResourceEnumSessionRecordings OrganizationResourceAccessResourceEnum = "session recordings"
	OrganizationResourceAccessResourceEnumInsights          OrganizationResourceAccessResourceEnum = "insights"
	OrganizationResourceAccessResourceEnumDashboards        OrganizationResourceAccessResourceEnum = "dashboards"
)

type OrganizationResourceAccessInput struct {
	AccessLevel *int64                                 `json:"access_level,omitempty" form:"name=access_level" multipartForm:"name=access_level"`
	Resource    OrganizationResourceAccessResourceEnum `json:"resource" form:"name=resource" multipartForm:"name=resource"`
}

type OrganizationResourceAccess struct {
	AccessLevel  *int64                                 `json:"access_level,omitempty"`
	CreatedAt    time.Time                              `json:"created_at"`
	CreatedBy    int64                                  `json:"created_by"`
	ID           int64                                  `json:"id"`
	Organization string                                 `json:"organization"`
	Resource     OrganizationResourceAccessResourceEnum `json:"resource"`
	UpdatedAt    time.Time                              `json:"updated_at"`
}
