package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type IntegrationsCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type IntegrationsCreateRequests struct {
	Integration  *shared.IntegrationInput `request:"mediaType=application/json"`
	Integration1 *shared.IntegrationInput `request:"mediaType=application/x-www-form-urlencoded"`
	Integration2 *shared.IntegrationInput `request:"mediaType=multipart/form-data"`
}

type IntegrationsCreateRequest struct {
	PathParams IntegrationsCreatePathParams
	Request    IntegrationsCreateRequests
}

type IntegrationsCreateResponse struct {
	ContentType string
	Integration *shared.Integration
	StatusCode  int
}
