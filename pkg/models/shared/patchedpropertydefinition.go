package shared

type PatchedPropertyDefinitionPropertyTypeEnum string

const (
	PatchedPropertyDefinitionPropertyTypeEnumDateTime PatchedPropertyDefinitionPropertyTypeEnum = "DateTime"
	PatchedPropertyDefinitionPropertyTypeEnumString   PatchedPropertyDefinitionPropertyTypeEnum = "String"
	PatchedPropertyDefinitionPropertyTypeEnumNumeric  PatchedPropertyDefinitionPropertyTypeEnum = "Numeric"
	PatchedPropertyDefinitionPropertyTypeEnumBoolean  PatchedPropertyDefinitionPropertyTypeEnum = "Boolean"
	PatchedPropertyDefinitionPropertyTypeEnumUnknown  PatchedPropertyDefinitionPropertyTypeEnum = ""
	PatchedPropertyDefinitionPropertyTypeEnumNull     PatchedPropertyDefinitionPropertyTypeEnum = "null"
)

// PatchedPropertyDefinitionInput
// Serializer mixin that resolves appropriate response for tags depending on license.
type PatchedPropertyDefinitionInput struct {
	IsNumerical     *bool                                      `json:"is_numerical,omitempty" form:"name=is_numerical" multipartForm:"name=is_numerical"`
	Name            *string                                    `json:"name,omitempty" form:"name=name" multipartForm:"name=name"`
	PropertyType    *PatchedPropertyDefinitionPropertyTypeEnum `json:"property_type,omitempty" form:"name=property_type" multipartForm:"name=property_type"`
	QueryUsage30Day *int64                                     `json:"query_usage_30_day,omitempty" form:"name=query_usage_30_day" multipartForm:"name=query_usage_30_day"`
	Tags            []interface{}                              `json:"tags,omitempty" form:"name=tags,json" multipartForm:"name=tags,json"`
}
