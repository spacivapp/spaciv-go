# \NeedsAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEstimations**](NeedsAPI.md#GetEstimations) | **Get** /v1/needs/{id}/estimations/{dataset}/{group}/{node} | Get a need estimation
[**Update**](NeedsAPI.md#Update) | **Patch** /v1/needs/{id} | Update a need estimation



## GetEstimations

> Contribution GetEstimations(ctx, id, dataset, group, node)

Get a need estimation



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
	dataset :=  // string | Dataset identifier; must match the top-down model's propertyId.
	group :=  // string | Aggregation scope; currently only \"all\" is supported.
	node :=  // string | Entry ID within the top-down model whose contribution values are returned.

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Needs.GetEstimations(context.Background(), id, dataset, group, node)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Needs.GetEstimations`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Needs.GetEstimations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**dataset** | **string** | Dataset identifier; must match the top-down model&#39;s propertyId. | 
**group** | **string** | Aggregation scope; currently only \&quot;all\&quot; is supported. | 
**node** | **string** | Entry ID within the top-down model whose contribution values are returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEstimationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Contribution**](Contribution.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(ctx, id, estimationPatch, params)

Update a need estimation



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
	estimationPatch :=  // EstimationPatch | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.Needs.Update(context.Background(), id, estimationPatch, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Needs.Update`: %v\n", err)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **node** | **string** | Entry ID within the top-down model to apply the estimation to. | 
 **estimationPatch** | [**EstimationPatch**](EstimationPatch.md) |  | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

