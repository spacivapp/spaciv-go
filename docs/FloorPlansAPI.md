# \FloorPlansAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](FloorPlansAPI.md#Create) | **Post** /v1/floorplans | Create a floor plan
[**Delete**](FloorPlansAPI.md#Delete) | **Delete** /v1/floorplans/{id} | Delete a floor plan
[**Get**](FloorPlansAPI.md#Get) | **Get** /v1/floorplans/{id} | Get a floor plan
[**GetImages**](FloorPlansAPI.md#GetImages) | **Get** /v1/floorplans/{id}/images/{imageType} | Get a floor plan image
[**List**](FloorPlansAPI.md#List) | **Get** /v1/floorplans | List floor plans
[**Update**](FloorPlansAPI.md#Update) | **Patch** /v1/floorplans/{id} | Update a floor plan
[**UpdateImages**](FloorPlansAPI.md#UpdateImages) | **Put** /v1/floorplans/{id}/images/{imageType} | Upload a floor plan image



## Create

> string Create(ctx, floorplanPostBody)

Create a floor plan



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
	floorplanPostBody :=  // FloorplanPostBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.FloorPlans.Create(context.Background(), floorplanPostBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FloorPlans.Create`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `FloorPlans.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **floorplanPostBody** | [**FloorplanPostBody**](FloorplanPostBody.md) |  | 

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

> map[string]interface{} Delete(ctx, id)

Delete a floor plan



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
	resp, err := client.FloorPlans.Delete(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FloorPlans.Delete`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `FloorPlans.Delete`: %v\n", resp)
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


## Get

> FloorplanGetBody Get(ctx, id)

Get a floor plan



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
	resp, err := client.FloorPlans.Get(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FloorPlans.Get`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `FloorPlans.Get`: %v\n", resp)
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

[**FloorplanGetBody**](FloorplanGetBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImages

> string GetImages(ctx, id, imageType, params)

Get a floor plan image



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
	imageType :=  // string | Which image layer to retrieve: the base drawing (image) or the foreground overlay (foreground).

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.FloorPlans.GetImages(context.Background(), id, imageType, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FloorPlans.GetImages`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `FloorPlans.GetImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**imageType** | **string** | Which image layer to retrieve: the base drawing (image) or the foreground overlay (foreground). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **noFallback** | **bool** | When true, returns 404 instead of a generated placeholder if no image has been uploaded. | 

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


## List

> []FloorplanBrief List(ctx)

List floor plans



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
	resp, err := client.FloorPlans.List(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FloorPlans.List`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `FloorPlans.List`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


### Return type

[**[]FloorplanBrief**](FloorplanBrief.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> map[string]interface{} Update(ctx, id, floorplanPatchParamsBody)

Update a floor plan



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
	floorplanPatchParamsBody :=  // FloorplanPatchParamsBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.FloorPlans.Update(context.Background(), id, floorplanPatchParamsBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FloorPlans.Update`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `FloorPlans.Update`: %v\n", resp)
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

 **floorplanPatchParamsBody** | [**FloorplanPatchParamsBody**](FloorplanPatchParamsBody.md) |  | 

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


## UpdateImages

> map[string]interface{} UpdateImages(ctx, id, imageType, params)

Upload a floor plan image



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
	imageType :=  // string | Which image layer to upload: the base drawing (image) or the foreground overlay (foreground).

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.FloorPlans.UpdateImages(context.Background(), id, imageType, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FloorPlans.UpdateImages`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `FloorPlans.UpdateImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**imageType** | **string** | Which image layer to upload: the base drawing (image) or the foreground overlay (foreground). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filename** | ***os.File** | filename of the file being uploaded | 
 **name** | **string** | general purpose name for multipart form value | 

### Return type

**map[string]interface{}**

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

