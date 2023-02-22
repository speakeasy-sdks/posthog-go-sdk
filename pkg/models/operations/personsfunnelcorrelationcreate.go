package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsFunnelCorrelationCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsFunnelCorrelationCreateFormatEnum string

const (
	PersonsFunnelCorrelationCreateFormatEnumCsv  PersonsFunnelCorrelationCreateFormatEnum = "csv"
	PersonsFunnelCorrelationCreateFormatEnumJSON PersonsFunnelCorrelationCreateFormatEnum = "json"
)

type PersonsFunnelCorrelationCreateQueryParams struct {
	Format *PersonsFunnelCorrelationCreateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsFunnelCorrelationCreateRequests struct {
	Person  *shared.PersonInput `request:"mediaType=application/json"`
	Person1 *shared.PersonInput `request:"mediaType=application/x-www-form-urlencoded"`
	Person2 *shared.PersonInput `request:"mediaType=multipart/form-data"`
}

type PersonsFunnelCorrelationCreateRequest struct {
	PathParams  PersonsFunnelCorrelationCreatePathParams
	QueryParams PersonsFunnelCorrelationCreateQueryParams
	Request     *PersonsFunnelCorrelationCreateRequests
}

type PersonsFunnelCorrelationCreateResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int
}
