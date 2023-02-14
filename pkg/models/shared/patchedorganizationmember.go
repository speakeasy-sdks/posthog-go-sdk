package shared

type PatchedOrganizationMemberInput struct {
	Level *int64 `json:"level,omitempty" form:"name=level" multipartForm:"name=level"`
}
