package shared

import (
	"time"
)

type PerformanceEvent struct {
	ConnectEnd                       *float64  `json:"connect_end,omitempty"`
	ConnectStart                     *float64  `json:"connect_start,omitempty"`
	CurrentURL                       string    `json:"current_url"`
	DecodedBodySize                  *int64    `json:"decoded_body_size,omitempty"`
	DistinctID                       string    `json:"distinct_id"`
	DomComplete                      *float64  `json:"dom_complete,omitempty"`
	DomContentLoadedEvent            *float64  `json:"dom_content_loaded_event,omitempty"`
	DomInteractive                   *float64  `json:"dom_interactive,omitempty"`
	DomainLookupEnd                  *float64  `json:"domain_lookup_end,omitempty"`
	DomainLookupStart                *float64  `json:"domain_lookup_start,omitempty"`
	Duration                         *float64  `json:"duration,omitempty"`
	EncodedBodySize                  *int64    `json:"encoded_body_size,omitempty"`
	EntryType                        string    `json:"entry_type"`
	FetchStart                       *float64  `json:"fetch_start,omitempty"`
	InitiatorType                    *string   `json:"initiator_type,omitempty"`
	LargestContentfulPaintElement    *string   `json:"largest_contentful_paint_element,omitempty"`
	LargestContentfulPaintID         *string   `json:"largest_contentful_paint_id,omitempty"`
	LargestContentfulPaintLoadTime   *float64  `json:"largest_contentful_paint_load_time,omitempty"`
	LargestContentfulPaintRenderTime *float64  `json:"largest_contentful_paint_render_time,omitempty"`
	LargestContentfulPaintSize       *float64  `json:"largest_contentful_paint_size,omitempty"`
	LargestContentfulPaintURL        *string   `json:"largest_contentful_paint_url,omitempty"`
	LoadEventEnd                     *float64  `json:"load_event_end,omitempty"`
	LoadEventStart                   *float64  `json:"load_event_start,omitempty"`
	Name                             string    `json:"name"`
	NavigationType                   *string   `json:"navigation_type,omitempty"`
	NextHopProtocol                  *string   `json:"next_hop_protocol,omitempty"`
	PageviewID                       *string   `json:"pageview_id,omitempty"`
	RedirectCount                    *int64    `json:"redirect_count,omitempty"`
	RedirectEnd                      *float64  `json:"redirect_end,omitempty"`
	RedirectStart                    *float64  `json:"redirect_start,omitempty"`
	RenderBlockingStatus             *string   `json:"render_blocking_status,omitempty"`
	RequestStart                     *float64  `json:"request_start,omitempty"`
	ResponseEnd                      *float64  `json:"response_end,omitempty"`
	ResponseStart                    *float64  `json:"response_start,omitempty"`
	ResponseStatus                   *int64    `json:"response_status,omitempty"`
	SecureConnectionStart            *float64  `json:"secure_connection_start,omitempty"`
	SessionID                        string    `json:"session_id"`
	StartTime                        *float64  `json:"start_time,omitempty"`
	TimeOrigin                       time.Time `json:"time_origin"`
	Timestamp                        time.Time `json:"timestamp"`
	TransferSize                     *int64    `json:"transfer_size,omitempty"`
	UnloadEventEnd                   *float64  `json:"unload_event_end,omitempty"`
	UnloadEventStart                 *float64  `json:"unload_event_start,omitempty"`
	UUID                             string    `json:"uuid"`
	WindowID                         string    `json:"window_id"`
	WorkerStart                      *float64  `json:"worker_start,omitempty"`
}
