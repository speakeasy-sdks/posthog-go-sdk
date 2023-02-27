package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type AnnotationsCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type AnnotationsCreateRequests struct {
	Annotation  *shared.AnnotationInput `request:"mediaType=application/json"`
	Annotation1 *shared.AnnotationInput `request:"mediaType=application/x-www-form-urlencoded"`
	Annotation2 *shared.AnnotationInput `request:"mediaType=multipart/form-data"`
}

type AnnotationsCreateRequest struct {
	PathParams AnnotationsCreatePathParams
	Request    *AnnotationsCreateRequests
}

type AnnotationsCreateResponse struct {
	Annotation  *shared.Annotation
	ContentType string
	StatusCode  int
}
