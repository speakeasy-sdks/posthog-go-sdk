package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PropertyDefinitionsListPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PropertyDefinitionsListTypeEnum string

const (
	PropertyDefinitionsListTypeEnumEvent  PropertyDefinitionsListTypeEnum = "event"
	PropertyDefinitionsListTypeEnumPerson PropertyDefinitionsListTypeEnum = "person"
	PropertyDefinitionsListTypeEnumGroup  PropertyDefinitionsListTypeEnum = "group"
)

type PropertyDefinitionsListQueryParams struct {
	EventNames         *string                          `queryParam:"style=form,explode=true,name=event_names"`
	ExcludedProperties *string                          `queryParam:"style=form,explode=true,name=excluded_properties"`
	FilterByEventNames *bool                            `queryParam:"style=form,explode=true,name=filter_by_event_names"`
	GroupTypeIndex     *int64                           `queryParam:"style=form,explode=true,name=group_type_index"`
	IsFeatureFlag      *bool                            `queryParam:"style=form,explode=true,name=is_feature_flag"`
	IsNumerical        *bool                            `queryParam:"style=form,explode=true,name=is_numerical"`
	Limit              *int64                           `queryParam:"style=form,explode=true,name=limit"`
	Offset             *int64                           `queryParam:"style=form,explode=true,name=offset"`
	Properties         *string                          `queryParam:"style=form,explode=true,name=properties"`
	Search             *string                          `queryParam:"style=form,explode=true,name=search"`
	Type               *PropertyDefinitionsListTypeEnum `queryParam:"style=form,explode=true,name=type"`
}

type PropertyDefinitionsListRequest struct {
	PathParams  PropertyDefinitionsListPathParams
	QueryParams PropertyDefinitionsListQueryParams
}

type PropertyDefinitionsListResponse struct {
	ContentType                     string
	PaginatedPropertyDefinitionList *shared.PaginatedPropertyDefinitionList
	StatusCode                      int64
}
