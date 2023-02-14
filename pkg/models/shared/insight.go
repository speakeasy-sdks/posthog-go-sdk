package shared

import (
	"time"
)

// InsightInput
// Simplified serializer to speed response times when loading large amounts of objects.
type InsightInput struct {
	Dashboards  []int64                `json:"dashboards,omitempty"`
	Deleted     *bool                  `json:"deleted,omitempty"`
	DerivedName *string                `json:"derived_name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Favorited   *bool                  `json:"favorited,omitempty"`
	Filters     map[string]interface{} `json:"filters,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Order       *int64                 `json:"order,omitempty"`
	Query       map[string]interface{} `json:"query,omitempty"`
	Saved       *bool                  `json:"saved,omitempty"`
	Tags        []interface{}          `json:"tags,omitempty"`
}

type InsightCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

type InsightLastModifiedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

// Insight
// Simplified serializer to speed response times when loading large amounts of objects.
type Insight struct {
	CreatedAt                 time.Time              `json:"created_at"`
	CreatedBy                 InsightCreatedBy       `json:"created_by"`
	DashboardTiles            []DashboardTileBasic   `json:"dashboard_tiles"`
	Dashboards                []int64                `json:"dashboards,omitempty"`
	Deleted                   *bool                  `json:"deleted,omitempty"`
	DerivedName               *string                `json:"derived_name,omitempty"`
	Description               *string                `json:"description,omitempty"`
	EffectivePrivilegeLevel   int64                  `json:"effective_privilege_level"`
	EffectiveRestrictionLevel int64                  `json:"effective_restriction_level"`
	Favorited                 *bool                  `json:"favorited,omitempty"`
	Filters                   map[string]interface{} `json:"filters,omitempty"`
	ID                        int64                  `json:"id"`
	IsCached                  string                 `json:"is_cached"`
	IsSample                  bool                   `json:"is_sample"`
	LastModifiedAt            time.Time              `json:"last_modified_at"`
	LastModifiedBy            InsightLastModifiedBy  `json:"last_modified_by"`
	LastRefresh               string                 `json:"last_refresh"`
	Name                      *string                `json:"name,omitempty"`
	Order                     *int64                 `json:"order,omitempty"`
	Query                     map[string]interface{} `json:"query,omitempty"`
	Result                    string                 `json:"result"`
	Saved                     *bool                  `json:"saved,omitempty"`
	ShortID                   string                 `json:"short_id"`
	Tags                      []interface{}          `json:"tags,omitempty"`
	Timezone                  string                 `json:"timezone"`
	UpdatedAt                 time.Time              `json:"updated_at"`
}
