package shared

type PaginatedDashboardList struct {
	Count    *int64            `json:"count,omitempty"`
	Next     *string           `json:"next,omitempty"`
	Previous *string           `json:"previous,omitempty"`
	Results  []DashboardOutput `json:"results,omitempty"`
}
