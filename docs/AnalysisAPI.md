# \AnalysisAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnalyse**](AnalysisAPI.md#CreateAnalyse) | **Post** /v1/analyse | Run analysis



## CreateAnalyse

> CalcRes CreateAnalyse(ctx, calcBody)

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
	resp, err := client.Analysis.CreateAnalyse(context.Background(), calcBody)
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

