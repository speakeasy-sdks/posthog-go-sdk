package shared

type PaginatedSubscriptionList struct {
	Count    *int64         `json:"count,omitempty"`
	Next     *string        `json:"next,omitempty"`
	Previous *string        `json:"previous,omitempty"`
	Results  []Subscription `json:"results,omitempty"`
}
