package shared

type FilterEventMathEnum string

const (
	FilterEventMathEnumTotal               FilterEventMathEnum = "total"
	FilterEventMathEnumDau                 FilterEventMathEnum = "dau"
	FilterEventMathEnumWeeklyActive        FilterEventMathEnum = "weekly_active"
	FilterEventMathEnumMonthlyActive       FilterEventMathEnum = "monthly_active"
	FilterEventMathEnumUniqueGroup         FilterEventMathEnum = "unique_group"
	FilterEventMathEnumUniqueSession       FilterEventMathEnum = "unique_session"
	FilterEventMathEnumSum                 FilterEventMathEnum = "sum"
	FilterEventMathEnumMin                 FilterEventMathEnum = "min"
	FilterEventMathEnumMax                 FilterEventMathEnum = "max"
	FilterEventMathEnumAvg                 FilterEventMathEnum = "avg"
	FilterEventMathEnumMedian              FilterEventMathEnum = "median"
	FilterEventMathEnumP90                 FilterEventMathEnum = "p90"
	FilterEventMathEnumP95                 FilterEventMathEnum = "p95"
	FilterEventMathEnumP99                 FilterEventMathEnum = "p99"
	FilterEventMathEnumMinCountPerActor    FilterEventMathEnum = "min_count_per_actor"
	FilterEventMathEnumMaxCountPerActor    FilterEventMathEnum = "max_count_per_actor"
	FilterEventMathEnumAvgCountPerActor    FilterEventMathEnum = "avg_count_per_actor"
	FilterEventMathEnumMedianCountPerActor FilterEventMathEnum = "median_count_per_actor"
	FilterEventMathEnumP90CountPerActor    FilterEventMathEnum = "p90_count_per_actor"
	FilterEventMathEnumP95CountPerActor    FilterEventMathEnum = "p95_count_per_actor"
	FilterEventMathEnumP99CountPerActor    FilterEventMathEnum = "p99_count_per_actor"
)

type FilterEvent struct {
	ID         string               `json:"id"`
	Math       *FilterEventMathEnum `json:"math,omitempty"`
	Properties []Property           `json:"properties,omitempty"`
}
