# \TopDownModelsAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](TopDownModelsAPI.md#Create) | **Post** /v1/top-down-models | Create a top-down model
[**Delete**](TopDownModelsAPI.md#Delete) | **Delete** /v1/top-down-models/{id} | Delete or archive a top-down model
[**DeleteEntries**](TopDownModelsAPI.md#DeleteEntries) | **Delete** /v1/top-down-models/{id}/entries | Delete top-down model entries
[**Duplicate**](TopDownModelsAPI.md#Duplicate) | **Post** /v1/top-down-models/{id}/duplicate | Duplicate a top-down model
[**Get**](TopDownModelsAPI.md#Get) | **Get** /v1/top-down-models/{id} | Get a top-down model
[**List**](TopDownModelsAPI.md#List) | **Get** /v1/top-down-models | List top-down models
[**ListEntries**](TopDownModelsAPI.md#ListEntries) | **Get** /v1/top-down-models/{id}/entries | List entries of a top-down model
[**Update**](TopDownModelsAPI.md#Update) | **Patch** /v1/top-down-models/{id} | Update a top-down model
[**UpdateEntries**](TopDownModelsAPI.md#UpdateEntries) | **Put** /v1/top-down-models/{id}/entries | Create or update a top-down model entry



## Create

> string Create(ctx, topDownModelChangelogBody)

Create a top-down model



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
	topDownModelChangelogBody :=  // TopDownModelChangelogBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.TopDownModels.Create(context.Background(), topDownModelChangelogBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopDownModels.Create`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `TopDownModels.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topDownModelChangelogBody** | [**TopDownModelChangelogBody**](TopDownModelChangelogBody.md) |  | 

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


## Delete

> Delete(ctx, id, params)

Delete or archive a top-down model



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
	err := client.TopDownModels.Delete(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopDownModels.Delete`: %v\n", err)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **string** | Lifecycle action: \&quot;archive\&quot; moves the model to the archive, \&quot;restore\&quot; returns it to active, and \&quot;purge\&quot; permanently deletes it along with all its entries. Defaults to \&quot;purge\&quot; when omitted. | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEntries

> DeleteEntries(ctx, id, requestBody)

Delete top-down model entries



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
	requestBody :=  // []string | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.TopDownModels.DeleteEntries(context.Background(), id, requestBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopDownModels.DeleteEntries`: %v\n", err)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

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


## Duplicate

> string Duplicate(ctx, id)

Duplicate a top-down model



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
	resp, err := client.TopDownModels.Duplicate(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopDownModels.Duplicate`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `TopDownModels.Duplicate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> TopDownModelChangelogBody Get(ctx, id)

Get a top-down model



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
	resp, err := client.TopDownModels.Get(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopDownModels.Get`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `TopDownModels.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TopDownModelChangelogBody**](TopDownModelChangelogBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []TopDownModelEntry List(ctx, params)

List top-down models



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

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.TopDownModels.List(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopDownModels.List`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `TopDownModels.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showArchived** | **bool** |  | 

### Return type

[**[]TopDownModelEntry**](TopDownModelEntry.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntries

> GetTopDownModelEntriesRes ListEntries(ctx, id, params)

List entries of a top-down model



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
	resp, err := client.TopDownModels.ListEntries(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopDownModels.ListEntries`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `TopDownModels.ListEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | Zero-based index of the first entry to return. Defaults to 0. | 
 **limit** | **int32** | Maximum number of entries to return. Returns every entry from start when omitted or 0. | 

### Return type

[**GetTopDownModelEntriesRes**](GetTopDownModelEntriesRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(ctx, id, topDownModelChangelogBody)

Update a top-down model



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
	topDownModelChangelogBody :=  // TopDownModelChangelogBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.TopDownModels.Update(context.Background(), id, topDownModelChangelogBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopDownModels.Update`: %v\n", err)
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

 **topDownModelChangelogBody** | [**TopDownModelChangelogBody**](TopDownModelChangelogBody.md) |  | 

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


## UpdateEntries

> UpdateEntries(ctx, id, topDownModelEntryUpdate)

Create or update a top-down model entry



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
	topDownModelEntryUpdate :=  // TopDownModelEntryUpdate | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.TopDownModels.UpdateEntries(context.Background(), id, topDownModelEntryUpdate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopDownModels.UpdateEntries`: %v\n", err)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topDownModelEntryUpdate** | [**TopDownModelEntryUpdate**](TopDownModelEntryUpdate.md) |  | 

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

