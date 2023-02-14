package shared

type PatchedHookInput struct {
	Event      *string `json:"event,omitempty" form:"name=event" multipartForm:"name=event"`
	ID         *string `json:"id,omitempty" form:"name=id" multipartForm:"name=id"`
	ResourceID *int64  `json:"resource_id,omitempty" form:"name=resource_id" multipartForm:"name=resource_id"`
	Target     *string `json:"target,omitempty" form:"name=target" multipartForm:"name=target"`
}
