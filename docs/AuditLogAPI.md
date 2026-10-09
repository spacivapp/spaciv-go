# \AuditLogAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](AuditLogAPI.md#List) | **Get** /v1/audit-logs | Query the audit log



## List

> GetResponseBody List(ctx, params)

Query the audit log



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/spacivapp/spaciv-go"
)

func main() {

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.AuditLog.List(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLog.List`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `AuditLog.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **string** | Entity type to query. One of: dashboard, property, scenario, floorplan, space-ruleset, top-down-model, object-group, datacollection, contribution, i18n, user, file, variable, upload-quarantine. Pass \&quot;all\&quot; to retrieve entries across all entity types the caller has permission to see (entityId must be omitted). | 
 **entityId** | **string** | Restrict results to changes made to this specific entity. Cannot be used when entity&#x3D;all. | 
 **from** | **time.Time** | Return only entries at or after this timestamp (RFC 3339). Omit for no lower bound. | 
 **to** | **time.Time** | Return only entries at or before this timestamp (RFC 3339). Omit for no upper bound. | 
 **limit** | **int32** |  | [default to 100]
 **cursor** | **string** | Opaque pagination cursor returned as nextCursor by a previous response. | 

### Return type

[**GetResponseBody**](GetResponseBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

