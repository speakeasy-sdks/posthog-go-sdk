package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type RetrievePathParams struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type RetrieveRequest struct {
	PathParams RetrievePathParams
}

type RetrieveResponse struct {
	ContentType string
	StatusCode  int
	Team        *shared.Team
}
