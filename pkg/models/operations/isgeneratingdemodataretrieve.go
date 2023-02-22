package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type IsGeneratingDemoDataRetrievePathParams struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type IsGeneratingDemoDataRetrieveRequest struct {
	PathParams IsGeneratingDemoDataRetrievePathParams
}

type IsGeneratingDemoDataRetrieveResponse struct {
	ContentType string
	StatusCode  int
	Team        *shared.Team
}
