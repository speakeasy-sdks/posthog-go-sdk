package posthog

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/utils"
	"net/http"
)

type performanceEvents struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newPerformanceEvents(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *performanceEvents {
	return &performanceEvents{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

func (s *performanceEvents) PerformanceEventsList(ctx context.Context, request operations.PerformanceEventsListRequest) (*operations.PerformanceEventsListResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/api/projects/{project_id}/performance_events/", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

	client := s.defaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PerformanceEventsListResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.PaginatedPerformanceEventList
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PaginatedPerformanceEventList = out
		}
	}

	return res, nil
}

func (s *performanceEvents) PerformanceEventsRecentPageviewsRetrieve(ctx context.Context, request operations.PerformanceEventsRecentPageviewsRetrieveRequest) (*operations.PerformanceEventsRecentPageviewsRetrieveResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/api/projects/{project_id}/performance_events/recent_pageviews/", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.defaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PerformanceEventsRecentPageviewsRetrieveResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.PerformanceEvent
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PerformanceEvent = out
		}
	}

	return res, nil
}
