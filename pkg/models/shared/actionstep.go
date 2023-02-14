package shared

type ActionStepURLMatchingEnum string

const (
	ActionStepURLMatchingEnumContains ActionStepURLMatchingEnum = "contains"
	ActionStepURLMatchingEnumRegex    ActionStepURLMatchingEnum = "regex"
	ActionStepURLMatchingEnumExact    ActionStepURLMatchingEnum = "exact"
	ActionStepURLMatchingEnumUnknown  ActionStepURLMatchingEnum = ""
	ActionStepURLMatchingEnumNull     ActionStepURLMatchingEnum = "null"
)

type ActionStep struct {
	Event       *string                    `json:"event,omitempty"`
	Href        *string                    `json:"href,omitempty"`
	ID          *string                    `json:"id,omitempty"`
	Name        *string                    `json:"name,omitempty"`
	Properties  map[string]interface{}     `json:"properties,omitempty"`
	Selector    *string                    `json:"selector,omitempty"`
	TagName     *string                    `json:"tag_name,omitempty"`
	Text        *string                    `json:"text,omitempty"`
	URL         *string                    `json:"url,omitempty"`
	URLMatching *ActionStepURLMatchingEnum `json:"url_matching,omitempty"`
}
