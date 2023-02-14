package shared

type TrendBreakdownTypeEnum string

const (
	TrendBreakdownTypeEnumEvent   TrendBreakdownTypeEnum = "event"
	TrendBreakdownTypeEnumPerson  TrendBreakdownTypeEnum = "person"
	TrendBreakdownTypeEnumCohort  TrendBreakdownTypeEnum = "cohort"
	TrendBreakdownTypeEnumGroup   TrendBreakdownTypeEnum = "group"
	TrendBreakdownTypeEnumSession TrendBreakdownTypeEnum = "session"
	TrendBreakdownTypeEnumHogql   TrendBreakdownTypeEnum = "hogql"
)

type TrendDisplayEnum string

const (
	TrendDisplayEnumActionsLineGraph           TrendDisplayEnum = "ActionsLineGraph"
	TrendDisplayEnumActionsLineGraphCumulative TrendDisplayEnum = "ActionsLineGraphCumulative"
	TrendDisplayEnumActionsTable               TrendDisplayEnum = "ActionsTable"
	TrendDisplayEnumActionsPie                 TrendDisplayEnum = "ActionsPie"
	TrendDisplayEnumActionsBar                 TrendDisplayEnum = "ActionsBar"
	TrendDisplayEnumActionsBarValue            TrendDisplayEnum = "ActionsBarValue"
	TrendDisplayEnumWorldMap                   TrendDisplayEnum = "WorldMap"
	TrendDisplayEnumBoldNumber                 TrendDisplayEnum = "BoldNumber"
)

type Trend struct {
	Actions            []FilterAction          `json:"actions,omitempty"`
	Breakdown          *string                 `json:"breakdown,omitempty"`
	BreakdownType      *TrendBreakdownTypeEnum `json:"breakdown_type,omitempty"`
	Compare            *bool                   `json:"compare,omitempty"`
	DateFrom           *string                 `json:"date_from,omitempty"`
	DateTo             *string                 `json:"date_to,omitempty"`
	Display            *TrendDisplayEnum       `json:"display,omitempty"`
	Events             []FilterEvent           `json:"events,omitempty"`
	FilterTestAccounts *bool                   `json:"filter_test_accounts,omitempty"`
	Formula            *string                 `json:"formula,omitempty"`
	Properties         *Property               `json:"properties,omitempty"`
}
