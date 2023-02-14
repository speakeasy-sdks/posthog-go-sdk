package shared

type ClickhouseEvent struct {
	DistinctID    string `json:"distinct_id"`
	Elements      string `json:"elements"`
	ElementsChain string `json:"elements_chain"`
	Event         string `json:"event"`
	ID            string `json:"id"`
	Person        string `json:"person"`
	Properties    string `json:"properties"`
	Timestamp     string `json:"timestamp"`
}
