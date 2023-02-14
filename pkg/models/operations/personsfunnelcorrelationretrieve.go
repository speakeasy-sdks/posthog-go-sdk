package operations

import (
	"github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
)

type PersonsFunnelCorrelationRetrievePathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type PersonsFunnelCorrelationRetrieveFormatEnum string

const (
	PersonsFunnelCorrelationRetrieveFormatEnumCsv  PersonsFunnelCorrelationRetrieveFormatEnum = "csv"
	PersonsFunnelCorrelationRetrieveFormatEnumJSON PersonsFunnelCorrelationRetrieveFormatEnum = "json"
)

type PersonsFunnelCorrelationRetrieveQueryParams struct {
	Format *PersonsFunnelCorrelationRetrieveFormatEnum `queryParam:"style=form,explode=true,name=format"`
}

type PersonsFunnelCorrelationRetrieveRequest struct {
	PathParams  PersonsFunnelCorrelationRetrievePathParams
	QueryParams PersonsFunnelCorrelationRetrieveQueryParams
}

type PersonsFunnelCorrelationRetrieveResponse struct {
	Body        []byte
	ContentType string
	Person      *shared.Person
	StatusCode  int64
}
