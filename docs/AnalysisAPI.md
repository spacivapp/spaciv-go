# \AnalysisAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnalyse**](AnalysisAPI.md#CreateAnalyse) | **Post** /v1/analyse | Run analysis
[**CreateGraphLayersPreview**](AnalysisAPI.md#CreateGraphLayersPreview) | **Post** /v1/graph-layers/{id}/preview | Preview a graph layer



## CreateAnalyse

> CalcRes CreateAnalyse(ctx, calcBody, params)

Run analysis



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
	calcBody :=  // CalcBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Analysis.CreateAnalyse(context.Background(), calcBody, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Analysis.CreateAnalyse`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Analysis.CreateAnalyse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnalyseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCalcReason** | **string** | Why the app asks, counted to measure the load it causes: data-change is a visualisation refreshing after the account&#39;s data changed. | 
 **calcBody** | [**CalcBody**](CalcBody.md) |  | 

### Return type

[**CalcRes**](CalcRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGraphLayersPreview

> GraphPreviewResult CreateGraphLayersPreview(ctx, id, graphPreviewBody)

Preview a graph layer



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
	id :=  // string | 
	graphPreviewBody :=  // GraphPreviewBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Analysis.CreateGraphLayersPreview(context.Background(), id, graphPreviewBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Analysis.CreateGraphLayersPreview`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Analysis.CreateGraphLayersPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGraphLayersPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **graphPreviewBody** | [**GraphPreviewBody**](GraphPreviewBody.md) |  | 

### Return type

[**GraphPreviewResult**](GraphPreviewResult.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

