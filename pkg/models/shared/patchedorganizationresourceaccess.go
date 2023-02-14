package shared

type PatchedOrganizationResourceAccessResourceEnum string

const (
	PatchedOrganizationResourceAccessResourceEnumFeatureFlags      PatchedOrganizationResourceAccessResourceEnum = "feature flags"
	PatchedOrganizationResourceAccessResourceEnumExperiments       PatchedOrganizationResourceAccessResourceEnum = "experiments"
	PatchedOrganizationResourceAccessResourceEnumCohorts           PatchedOrganizationResourceAccessResourceEnum = "cohorts"
	PatchedOrganizationResourceAccessResourceEnumDataManagement    PatchedOrganizationResourceAccessResourceEnum = "data management"
	PatchedOrganizationResourceAccessResourceEnumSessionRecordings PatchedOrganizationResourceAccessResourceEnum = "session recordings"
	PatchedOrganizationResourceAccessResourceEnumInsights          PatchedOrganizationResourceAccessResourceEnum = "insights"
	PatchedOrganizationResourceAccessResourceEnumDashboards        PatchedOrganizationResourceAccessResourceEnum = "dashboards"
)

type PatchedOrganizationResourceAccessInput struct {
	AccessLevel *int64                                         `json:"access_level,omitempty" form:"name=access_level" multipartForm:"name=access_level"`
	Resource    *PatchedOrganizationResourceAccessResourceEnum `json:"resource,omitempty" form:"name=resource" multipartForm:"name=resource"`
}
