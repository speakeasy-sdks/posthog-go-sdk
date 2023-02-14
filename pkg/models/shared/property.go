package shared

type PropertyTypeEnum string

const (
	PropertyTypeEnumAnd PropertyTypeEnum = "AND"
	PropertyTypeEnumOr  PropertyTypeEnum = "OR"
)

type Property struct {
	Type   *PropertyTypeEnum `json:"type,omitempty"`
	Values []PropertyItem    `json:"values"`
}
