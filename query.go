package posthog

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/utils"
	"net/http"
)

type query struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newQuery(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *query {
	return &query{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

func (s *query) QueryRetrieve(ctx context.Context, request operations.QueryRetrieveRequest) (*operations.QueryRetrieveResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/api/projects/{project_id}/query/", request.PathParams)

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

	res := &operations.QueryRetrieveResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}
