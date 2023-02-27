package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type AnnotationsUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type AnnotationsUpdateRequests struct {
	Annotation  *shared.AnnotationInput `request:"mediaType=application/json"`
	Annotation1 *shared.AnnotationInput `request:"mediaType=application/x-www-form-urlencoded"`
	Annotation2 *shared.AnnotationInput `request:"mediaType=multipart/form-data"`
}

type AnnotationsUpdateRequest struct {
	PathParams AnnotationsUpdatePathParams
	Request    *AnnotationsUpdateRequests
}

type AnnotationsUpdateResponse struct {
	Annotation  *shared.Annotation
	ContentType string
	StatusCode  int
}
