package shared

type GroupType struct {
	GroupType      string  `json:"group_type"`
	GroupTypeIndex int64   `json:"group_type_index"`
	NamePlural     *string `json:"name_plural,omitempty"`
	NameSingular   *string `json:"name_singular,omitempty"`
}
