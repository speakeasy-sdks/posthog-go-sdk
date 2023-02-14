package shared

type FunnelStepsResult struct {
	ActionID              string  `json:"action_id"`
	AverageConversionTime float64 `json:"average_conversion_time"`
	ConvertedPeopleURL    string  `json:"converted_people_url"`
	Count                 int64   `json:"count"`
	DroppedPeopleURL      string  `json:"dropped_people_url"`
	MedianConversionTime  float64 `json:"median_conversion_time"`
	Order                 string  `json:"order"`
}
