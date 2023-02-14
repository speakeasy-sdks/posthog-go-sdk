package posthog

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/operations"
	"net/http"
	"strings"
)

type prompts struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newPrompts(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *prompts {
	return &prompts{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// PromptsMyPromptsPartialUpdate - Create, read, update and delete prompt sequences state for a person.
func (s *prompts) PromptsMyPromptsPartialUpdate(ctx context.Context) (*operations.PromptsMyPromptsPartialUpdateResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/api/prompts/my_prompts/"

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, nil)
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

	res := &operations.PromptsMyPromptsPartialUpdateResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// PromptsMyPromptsPartialUpdate - Create, read, update and delete prompt sequences state for a person.
func (s *prompts) PromptsMyPromptsPartialUpdate(ctx context.Context) (*operations.PromptsMyPromptsPartialUpdateResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/api/prompts/my_prompts/"

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, nil)
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

	res := &operations.PromptsMyPromptsPartialUpdateResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}
