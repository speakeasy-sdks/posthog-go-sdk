package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type AnnotationsPartialUpdatePathParams struct {
	ID        int64  `pathParam:"style=simple,explode=false,name=id"`
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type AnnotationsPartialUpdateRequests struct {
	PatchedAnnotation  *shared.PatchedAnnotationInput `request:"mediaType=application/json"`
	PatchedAnnotation1 *shared.PatchedAnnotationInput `request:"mediaType=application/x-www-form-urlencoded"`
	PatchedAnnotation2 *shared.PatchedAnnotationInput `request:"mediaType=multipart/form-data"`
}

type AnnotationsPartialUpdateRequest struct {
	PathParams AnnotationsPartialUpdatePathParams
	Request    *AnnotationsPartialUpdateRequests
}

type AnnotationsPartialUpdateResponse struct {
	Annotation  *shared.Annotation
	ContentType string
	StatusCode  int
}
