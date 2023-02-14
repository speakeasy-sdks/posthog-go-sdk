package shared

type PropertyDefinitionPropertyTypeEnum string

const (
	PropertyDefinitionPropertyTypeEnumDateTime PropertyDefinitionPropertyTypeEnum = "DateTime"
	PropertyDefinitionPropertyTypeEnumString   PropertyDefinitionPropertyTypeEnum = "String"
	PropertyDefinitionPropertyTypeEnumNumeric  PropertyDefinitionPropertyTypeEnum = "Numeric"
	PropertyDefinitionPropertyTypeEnumBoolean  PropertyDefinitionPropertyTypeEnum = "Boolean"
	PropertyDefinitionPropertyTypeEnumUnknown  PropertyDefinitionPropertyTypeEnum = ""
	PropertyDefinitionPropertyTypeEnumNull     PropertyDefinitionPropertyTypeEnum = "null"
)

// PropertyDefinitionInput
// Serializer mixin that resolves appropriate response for tags depending on license.
type PropertyDefinitionInput struct {
	IsNumerical     *bool                               `json:"is_numerical,omitempty" form:"name=is_numerical" multipartForm:"name=is_numerical"`
	Name            string                              `json:"name" form:"name=name" multipartForm:"name=name"`
	PropertyType    *PropertyDefinitionPropertyTypeEnum `json:"property_type,omitempty" form:"name=property_type" multipartForm:"name=property_type"`
	QueryUsage30Day *int64                              `json:"query_usage_30_day,omitempty" form:"name=query_usage_30_day" multipartForm:"name=query_usage_30_day"`
	Tags            []interface{}                       `json:"tags,omitempty" form:"name=tags,json" multipartForm:"name=tags,json"`
}

// PropertyDefinition
// Serializer mixin that resolves appropriate response for tags depending on license.
type PropertyDefinition struct {
	ID                     string                              `json:"id"`
	IsNumerical            *bool                               `json:"is_numerical,omitempty"`
	IsSeenOnFilteredEvents string                              `json:"is_seen_on_filtered_events"`
	Name                   string                              `json:"name"`
	PropertyType           *PropertyDefinitionPropertyTypeEnum `json:"property_type,omitempty"`
	QueryUsage30Day        *int64                              `json:"query_usage_30_day,omitempty"`
	Tags                   []interface{}                       `json:"tags,omitempty"`
}
