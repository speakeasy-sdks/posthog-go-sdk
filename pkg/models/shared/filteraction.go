package shared

type FilterActionMathEnum string

const (
	FilterActionMathEnumTotal               FilterActionMathEnum = "total"
	FilterActionMathEnumDau                 FilterActionMathEnum = "dau"
	FilterActionMathEnumWeeklyActive        FilterActionMathEnum = "weekly_active"
	FilterActionMathEnumMonthlyActive       FilterActionMathEnum = "monthly_active"
	FilterActionMathEnumUniqueGroup         FilterActionMathEnum = "unique_group"
	FilterActionMathEnumUniqueSession       FilterActionMathEnum = "unique_session"
	FilterActionMathEnumSum                 FilterActionMathEnum = "sum"
	FilterActionMathEnumMin                 FilterActionMathEnum = "min"
	FilterActionMathEnumMax                 FilterActionMathEnum = "max"
	FilterActionMathEnumAvg                 FilterActionMathEnum = "avg"
	FilterActionMathEnumMedian              FilterActionMathEnum = "median"
	FilterActionMathEnumP90                 FilterActionMathEnum = "p90"
	FilterActionMathEnumP95                 FilterActionMathEnum = "p95"
	FilterActionMathEnumP99                 FilterActionMathEnum = "p99"
	FilterActionMathEnumMinCountPerActor    FilterActionMathEnum = "min_count_per_actor"
	FilterActionMathEnumMaxCountPerActor    FilterActionMathEnum = "max_count_per_actor"
	FilterActionMathEnumAvgCountPerActor    FilterActionMathEnum = "avg_count_per_actor"
	FilterActionMathEnumMedianCountPerActor FilterActionMathEnum = "median_count_per_actor"
	FilterActionMathEnumP90CountPerActor    FilterActionMathEnum = "p90_count_per_actor"
	FilterActionMathEnumP95CountPerActor    FilterActionMathEnum = "p95_count_per_actor"
	FilterActionMathEnumP99CountPerActor    FilterActionMathEnum = "p99_count_per_actor"
)

type FilterAction struct {
	ID         string                `json:"id"`
	Math       *FilterActionMathEnum `json:"math,omitempty"`
	Properties []Property            `json:"properties,omitempty"`
}
