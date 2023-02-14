package shared

type DashboardTileBasic struct {
	DashboardID int64 `json:"dashboard_id"`
	Deleted     *bool `json:"deleted,omitempty"`
	ID          int64 `json:"id"`
}
