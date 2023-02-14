package shared

type PaginatedTeamBasicList struct {
	Count    *int64      `json:"count,omitempty"`
	Next     *string     `json:"next,omitempty"`
	Previous *string     `json:"previous,omitempty"`
	Results  []TeamBasic `json:"results,omitempty"`
}
