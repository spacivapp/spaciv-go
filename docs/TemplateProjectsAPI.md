# \TemplateProjectsAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMapping**](TemplateProjectsAPI.md#GetMapping) | **Get** /v1/template-projects/{id}/mapping | Get template project mapping



## GetMapping

> GetMappingRes GetMapping(ctx, id)

Get template project mapping



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spacivapp/spaciv-go"
)

func main() {
	id :=  // int64 | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.TemplateProjects.GetMapping(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateProjects.GetMapping`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `TemplateProjects.GetMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMappingRes**](GetMappingRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

