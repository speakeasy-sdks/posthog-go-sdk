package shared

type PropertyItemOperatorEnum string

const (
	PropertyItemOperatorEnumExact        PropertyItemOperatorEnum = "exact"
	PropertyItemOperatorEnumIsNot        PropertyItemOperatorEnum = "is_not"
	PropertyItemOperatorEnumIcontains    PropertyItemOperatorEnum = "icontains"
	PropertyItemOperatorEnumNotIcontains PropertyItemOperatorEnum = "not_icontains"
	PropertyItemOperatorEnumRegex        PropertyItemOperatorEnum = "regex"
	PropertyItemOperatorEnumNotRegex     PropertyItemOperatorEnum = "not_regex"
	PropertyItemOperatorEnumGt           PropertyItemOperatorEnum = "gt"
	PropertyItemOperatorEnumLt           PropertyItemOperatorEnum = "lt"
	PropertyItemOperatorEnumGte          PropertyItemOperatorEnum = "gte"
	PropertyItemOperatorEnumLte          PropertyItemOperatorEnum = "lte"
	PropertyItemOperatorEnumIsSet        PropertyItemOperatorEnum = "is_set"
	PropertyItemOperatorEnumIsNotSet     PropertyItemOperatorEnum = "is_not_set"
	PropertyItemOperatorEnumIsDateExact  PropertyItemOperatorEnum = "is_date_exact"
	PropertyItemOperatorEnumIsDateAfter  PropertyItemOperatorEnum = "is_date_after"
	PropertyItemOperatorEnumIsDateBefore PropertyItemOperatorEnum = "is_date_before"
	PropertyItemOperatorEnumUnknown      PropertyItemOperatorEnum = ""
	PropertyItemOperatorEnumNull         PropertyItemOperatorEnum = "null"
)

type PropertyItemTypeEnum string

const (
	PropertyItemTypeEnumEvent               PropertyItemTypeEnum = "event"
	PropertyItemTypeEnumPerson              PropertyItemTypeEnum = "person"
	PropertyItemTypeEnumCohort              PropertyItemTypeEnum = "cohort"
	PropertyItemTypeEnumElement             PropertyItemTypeEnum = "element"
	PropertyItemTypeEnumStaticCohort        PropertyItemTypeEnum = "static-cohort"
	PropertyItemTypeEnumPrecalculatedCohort PropertyItemTypeEnum = "precalculated-cohort"
	PropertyItemTypeEnumGroup               PropertyItemTypeEnum = "group"
	PropertyItemTypeEnumRecording           PropertyItemTypeEnum = "recording"
	PropertyItemTypeEnumBehavioral          PropertyItemTypeEnum = "behavioral"
	PropertyItemTypeEnumSession             PropertyItemTypeEnum = "session"
	PropertyItemTypeEnumHogql               PropertyItemTypeEnum = "hogql"
	PropertyItemTypeEnumUnknown             PropertyItemTypeEnum = ""
)

type PropertyItem struct {
	Key      string                    `json:"key"`
	Operator *PropertyItemOperatorEnum `json:"operator,omitempty"`
	Type     *PropertyItemTypeEnum     `json:"type,omitempty"`
	Value    string                    `json:"value"`
}
