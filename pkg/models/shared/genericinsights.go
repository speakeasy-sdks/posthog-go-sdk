package shared

type GenericInsights struct {
	Actions            []FilterAction `json:"actions,omitempty"`
	DateFrom           *string        `json:"date_from,omitempty"`
	DateTo             *string        `json:"date_to,omitempty"`
	Events             []FilterEvent  `json:"events,omitempty"`
	FilterTestAccounts *bool          `json:"filter_test_accounts,omitempty"`
	Properties         *Property      `json:"properties,omitempty"`
}
