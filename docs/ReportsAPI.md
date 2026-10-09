# \ReportsAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ReportsAPI.md#Create) | **Post** /v1/reports | Generate a report



## Create

> string Create(ctx, reportParamsWrapperBody)

Generate a report



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
	reportParamsWrapperBody :=  // ReportParamsWrapperBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Reports.Create(context.Background(), reportParamsWrapperBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Reports.Create`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Reports.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportParamsWrapperBody** | [**ReportParamsWrapperBody**](ReportParamsWrapperBody.md) |  | 

### Return type

**string**

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

