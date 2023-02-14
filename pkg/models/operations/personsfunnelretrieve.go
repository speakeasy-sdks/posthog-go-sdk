package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsFunnelRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsFunnelRetrieveFormatEnum string

const (
	PersonsFunnelRetrieveFormatEnumCsv  PersonsFunnelRetrieveFormatEnum = "csv"
	PersonsFunnelRetrieveFormatEnumJSON PersonsFunnelRetrieveFormatEnum = "json"
)

type PersonsFunnelRetrieveQueryParams struct {
	Format *PersonsFunnelRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsFunnelRetrieveRequest struct {
	PathParams  PersonsFunnelRetrievePathParams
	QueryParams PersonsFunnelRetrieveQueryParams
}

type PersonsFunnelRetrieveResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int64
}
