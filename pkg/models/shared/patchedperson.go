package shared

type PatchedPersonInput struct {
	Properties map[string]interface{} `json:"properties,omitempty" form:"name=properties,json" multipartForm:"name=properties,json"`
}
