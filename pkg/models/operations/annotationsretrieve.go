package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type AnnotationsRetrievePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type AnnotationsRetrieveRequest struct {
	PathParams AnnotationsRetrievePathParams
}

type AnnotationsRetrieveResponse struct {
	Annotation  *shared.Annotation
	ContentType string
	StatusCode  int
}
