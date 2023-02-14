package shared

type FunnelExclusion struct {
	FunnelFromStep *int64    `json:"funnel_from_step,omitempty"`
	FunnelToStep   *int64    `json:"funnel_to_step,omitempty"`
	ID             string    `json:"id"`
	Properties     *Property `json:"properties,omitempty"`
}
