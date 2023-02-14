package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsFunnelCreatePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsFunnelCreateFormatEnum string

const (
	PersonsFunnelCreateFormatEnumCsv  PersonsFunnelCreateFormatEnum = "csv"
	PersonsFunnelCreateFormatEnumJSON PersonsFunnelCreateFormatEnum = "json"
)

type PersonsFunnelCreateQueryParams struct {
	Format *PersonsFunnelCreateFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsFunnelCreateRequests struct {
	Person  *shared.PersonInput `request:"mediaType=application/json"`
	Person1 *shared.PersonInput `request:"mediaType=application/x-www-form-urlencoded"`
	Person2 *shared.PersonInput `request:"mediaType=multipart/form-data"`
}

type PersonsFunnelCreateRequest struct {
	PathParams  PersonsFunnelCreatePathParams
	QueryParams PersonsFunnelCreateQueryParams
	Request     *PersonsFunnelCreateRequests
}

type PersonsFunnelCreateResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int64
}
