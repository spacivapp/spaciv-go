# \ContributionsAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataCollectionsContribution**](ContributionsAPI.md#CreateDataCollectionsContribution) | **Post** /v1/data-collections/{id}/contribution | Save a contribution
[**CreateDataCollectionsUsersContribution**](ContributionsAPI.md#CreateDataCollectionsUsersContribution) | **Post** /v1/data-collections/{id}/users/{userId}/contribution | Create a contribution for a user
[**GetDataCollectionsContribution**](ContributionsAPI.md#GetDataCollectionsContribution) | **Get** /v1/data-collections/{id}/contribution | Get the caller&#39;s contribution for a survey



## CreateDataCollectionsContribution

> map[string]interface{} CreateDataCollectionsContribution(ctx, id, contributionPost, params)

Save a contribution



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
	contributionPost :=  // ContributionPost | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Contributions.CreateDataCollectionsContribution(context.Background(), id, contributionPost, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Contributions.CreateDataCollectionsContribution`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Contributions.CreateDataCollectionsContribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataCollectionsContributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datacollectionId** | **string** | The survey the contribution belongs to. Required when submitting on behalf of an anonymous or public contributor whose contribution ID differs from the survey ID. | 
 **contributionPost** | [**ContributionPost**](ContributionPost.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDataCollectionsUsersContribution

> map[string]interface{} CreateDataCollectionsUsersContribution(ctx, userId, id)

Create a contribution for a user



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
	userId :=  // int64 | 
	id :=  // string | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Contributions.CreateDataCollectionsUsersContribution(context.Background(), userId, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Contributions.CreateDataCollectionsUsersContribution`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Contributions.CreateDataCollectionsUsersContribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataCollectionsUsersContributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataCollectionsContribution

> GetContributionRes GetDataCollectionsContribution(ctx, id)

Get the caller's contribution for a survey



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

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Contributions.GetDataCollectionsContribution(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Contributions.GetDataCollectionsContribution`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Contributions.GetDataCollectionsContribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataCollectionsContributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContributionRes**](GetContributionRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

