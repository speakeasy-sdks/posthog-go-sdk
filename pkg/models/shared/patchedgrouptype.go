package shared

type PatchedGroupTypeInput struct {
	NamePlural   *string `json:"name_plural,omitempty" form:"name=name_plural" multipartForm:"name=name_plural"`
	NameSingular *string `json:"name_singular,omitempty" form:"name=name_singular" multipartForm:"name=name_singular"`
}
