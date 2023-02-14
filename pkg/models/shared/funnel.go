package shared

type FunnelBreakdownTypeEnum string

const (
	FunnelBreakdownTypeEnumEvent   FunnelBreakdownTypeEnum = "event"
	FunnelBreakdownTypeEnumPerson  FunnelBreakdownTypeEnum = "person"
	FunnelBreakdownTypeEnumCohort  FunnelBreakdownTypeEnum = "cohort"
	FunnelBreakdownTypeEnumGroup   FunnelBreakdownTypeEnum = "group"
	FunnelBreakdownTypeEnumSession FunnelBreakdownTypeEnum = "session"
	FunnelBreakdownTypeEnumHogql   FunnelBreakdownTypeEnum = "hogql"
)

type FunnelFunnelOrderTypeEnum string

const (
	FunnelFunnelOrderTypeEnumStrict    FunnelFunnelOrderTypeEnum = "strict"
	FunnelFunnelOrderTypeEnumUnordered FunnelFunnelOrderTypeEnum = "unordered"
	FunnelFunnelOrderTypeEnumOrdered   FunnelFunnelOrderTypeEnum = "ordered"
)

type FunnelFunnelVizTypeEnum string

const (
	FunnelFunnelVizTypeEnumTrends        FunnelFunnelVizTypeEnum = "trends"
	FunnelFunnelVizTypeEnumTimeToConvert FunnelFunnelVizTypeEnum = "time_to_convert"
	FunnelFunnelVizTypeEnumSteps         FunnelFunnelVizTypeEnum = "steps"
)

type FunnelFunnelWindowIntervalTypeEnum string

const (
	FunnelFunnelWindowIntervalTypeEnumDay    FunnelFunnelWindowIntervalTypeEnum = "DAY"
	FunnelFunnelWindowIntervalTypeEnumMinute FunnelFunnelWindowIntervalTypeEnum = "MINUTE"
	FunnelFunnelWindowIntervalTypeEnumHour   FunnelFunnelWindowIntervalTypeEnum = "HOUR"
	FunnelFunnelWindowIntervalTypeEnumWeek   FunnelFunnelWindowIntervalTypeEnum = "WEEK"
	FunnelFunnelWindowIntervalTypeEnumMonth  FunnelFunnelWindowIntervalTypeEnum = "MONTH"
)

type Funnel struct {
	Actions                   []FilterAction                      `json:"actions,omitempty"`
	AggregationGroupTypeIndex *int64                              `json:"aggregation_group_type_index,omitempty"`
	Breakdown                 *string                             `json:"breakdown,omitempty"`
	BreakdownLimit            *int64                              `json:"breakdown_limit,omitempty"`
	BreakdownType             *FunnelBreakdownTypeEnum            `json:"breakdown_type,omitempty"`
	DateFrom                  *string                             `json:"date_from,omitempty"`
	DateTo                    *string                             `json:"date_to,omitempty"`
	Events                    []FilterEvent                       `json:"events,omitempty"`
	Exclusions                []FunnelExclusion                   `json:"exclusions,omitempty"`
	FilterTestAccounts        *bool                               `json:"filter_test_accounts,omitempty"`
	FunnelOrderType           *FunnelFunnelOrderTypeEnum          `json:"funnel_order_type,omitempty"`
	FunnelVizType             *FunnelFunnelVizTypeEnum            `json:"funnel_viz_type,omitempty"`
	FunnelWindowDays          *int64                              `json:"funnel_window_days,omitempty"`
	FunnelWindowInterval      *int64                              `json:"funnel_window_interval,omitempty"`
	FunnelWindowIntervalType  *FunnelFunnelWindowIntervalTypeEnum `json:"funnel_window_interval_type,omitempty"`
	Properties                *Property                           `json:"properties,omitempty"`
}
