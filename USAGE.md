<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/posthog-go-sdk"
    "github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/operations"
)

func main() {
    s := posthog.New()
    
    req := operations.ActionsCountRetrieveRequest{
        PathParams: operations.ActionsCountRetrievePathParams{
            ID: 548814,
            ProjectID: "deserunt",
        },
        QueryParams: operations.ActionsCountRetrieveQueryParams{
            Format: "undefined",
        },
    }

    ctx := context.Background()
    res, err := s.Actions.ActionsCountRetrieve(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Action != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->