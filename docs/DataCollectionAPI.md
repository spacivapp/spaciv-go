# \DataCollectionAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](DataCollectionAPI.md#Create) | **Post** /v1/data-collections | Create a data collection survey
[**Delete**](DataCollectionAPI.md#Delete) | **Delete** /v1/data-collections/{id} | Delete a data collection survey
[**Get**](DataCollectionAPI.md#Get) | **Get** /v1/data-collections/{id} | Get a data collection survey
[**GetLink**](DataCollectionAPI.md#GetLink) | **Get** /v1/data-collections/{id}/link | Get a survey&#39;s invitation link
[**GetLinkQrCode**](DataCollectionAPI.md#GetLinkQrCode) | **Get** /v1/data-collections/{id}/link-qr-code | Get a survey&#39;s QR code
[**List**](DataCollectionAPI.md#List) | **Get** /v1/data-collections | List data collection surveys
[**ListContributors**](DataCollectionAPI.md#ListContributors) | **Get** /v1/data-collections/{id}/contributors | List contributors for a survey
[**Update**](DataCollectionAPI.md#Update) | **Patch** /v1/data-collections/{id} | Update a data collection survey



## Create

> DataCollectionAddResponseBody Create(ctx)

Create a data collection survey



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
	resp, err := client.DataCollection.Create(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataCollection.Create`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `DataCollection.Create`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


### Return type

[**DataCollectionAddResponseBody**](DataCollectionAddResponseBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> map[string]interface{} Delete(ctx, id)

Delete a data collection survey



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
	resp, err := client.DataCollection.Delete(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataCollection.Delete`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `DataCollection.Delete`: %v\n", resp)
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

> GetDatacollectionRes Get(ctx, id)

Get a data collection survey



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
	resp, err := client.DataCollection.Get(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataCollection.Get`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `DataCollection.Get`: %v\n", resp)
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

[**GetDatacollectionRes**](GetDatacollectionRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLink

> string GetLink(ctx, id)

Get a survey's invitation link



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
	resp, err := client.DataCollection.GetLink(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataCollection.GetLink`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `DataCollection.GetLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkRequest struct via the builder pattern


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


## GetLinkQrCode

> string GetLinkQrCode(ctx, id)

Get a survey's QR code



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
	resp, err := client.DataCollection.GetLinkQrCode(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataCollection.GetLinkQrCode`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `DataCollection.GetLinkQrCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkQrCodeRequest struct via the builder pattern


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


## List

> GetDatacollectionListRes List(ctx)

List data collection surveys



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
	resp, err := client.DataCollection.List(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataCollection.List`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `DataCollection.List`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


### Return type

[**GetDatacollectionListRes**](GetDatacollectionListRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContributors

> GetContributionListRes ListContributors(ctx, id)

List contributors for a survey



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
	resp, err := client.DataCollection.ListContributors(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataCollection.ListContributors`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `DataCollection.ListContributors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContributionListRes**](GetContributionListRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> map[string]interface{} Update(ctx, id, datacollectionBody)

Update a data collection survey



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
	datacollectionBody :=  // DatacollectionBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.DataCollection.Update(context.Background(), id, datacollectionBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataCollection.Update`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `DataCollection.Update`: %v\n", resp)
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

 **datacollectionBody** | [**DatacollectionBody**](DatacollectionBody.md) |  | 

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

