# \SpaceRulesetsAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](SpaceRulesetsAPI.md#Create) | **Post** /v1/space-rulesets | Create a space ruleset
[**CreateImage**](SpaceRulesetsAPI.md#CreateImage) | **Post** /v1/space-rulesets/{id}/image | Upload a space ruleset image
[**Delete**](SpaceRulesetsAPI.md#Delete) | **Delete** /v1/space-rulesets/{id} | Delete a space ruleset
[**DeleteImage**](SpaceRulesetsAPI.md#DeleteImage) | **Delete** /v1/space-rulesets/{id}/image | Delete a space ruleset image
[**Get**](SpaceRulesetsAPI.md#Get) | **Get** /v1/space-rulesets/{id} | Get a space ruleset
[**GetImage**](SpaceRulesetsAPI.md#GetImage) | **Get** /v1/space-rulesets/{id}/image | Get a space ruleset image
[**List**](SpaceRulesetsAPI.md#List) | **Get** /v1/space-rulesets | List space rulesets
[**ListPresets**](SpaceRulesetsAPI.md#ListPresets) | **Get** /v1/space-rulesets/presets | List preset space rulesets
[**Update**](SpaceRulesetsAPI.md#Update) | **Patch** /v1/space-rulesets/{id} | Update a space ruleset



## Create

> map[string]interface{} Create(ctx, spaceRulesetCreate)

Create a space ruleset



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
	spaceRulesetCreate :=  // SpaceRulesetCreate | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.SpaceRulesets.Create(context.Background(), spaceRulesetCreate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRulesets.Create`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `SpaceRulesets.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spaceRulesetCreate** | [**SpaceRulesetCreate**](SpaceRulesetCreate.md) |  | 

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


## CreateImage

> map[string]interface{} CreateImage(ctx, id, params)

Upload a space ruleset image



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
	resp, err := client.SpaceRulesets.CreateImage(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRulesets.CreateImage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `SpaceRulesets.CreateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageRequest struct via the builder pattern


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


## Delete

> map[string]interface{} Delete(ctx, id)

Delete a space ruleset



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
	resp, err := client.SpaceRulesets.Delete(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRulesets.Delete`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `SpaceRulesets.Delete`: %v\n", resp)
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


## DeleteImage

> map[string]interface{} DeleteImage(ctx, id)

Delete a space ruleset image



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
	resp, err := client.SpaceRulesets.DeleteImage(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRulesets.DeleteImage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `SpaceRulesets.DeleteImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern


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

> SpaceRuleSetJSON Get(ctx, id)

Get a space ruleset



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
	resp, err := client.SpaceRulesets.Get(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRulesets.Get`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `SpaceRulesets.Get`: %v\n", resp)
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

[**SpaceRuleSetJSON**](SpaceRuleSetJSON.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImage

> string GetImage(ctx, id, params)

Get a space ruleset image



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
	resp, err := client.SpaceRulesets.GetImage(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRulesets.GetImage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `SpaceRulesets.GetImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noFallback** | **bool** | When true, returns 404 if no custom image has been uploaded rather than generating a default image from the ruleset name. | 

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

> []SpaceRulesetListEntry List(ctx)

List space rulesets



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
	resp, err := client.SpaceRulesets.List(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRulesets.List`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `SpaceRulesets.List`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


### Return type

[**[]SpaceRulesetListEntry**](SpaceRulesetListEntry.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPresets

> []SpaceRuleSet ListPresets(ctx)

List preset space rulesets



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
	resp, err := client.SpaceRulesets.ListPresets(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRulesets.ListPresets`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `SpaceRulesets.ListPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPresetsRequest struct via the builder pattern


### Return type

[**[]SpaceRuleSet**](SpaceRuleSet.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> map[string]interface{} Update(ctx, id, spaceRulesetUpdate)

Update a space ruleset



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
	spaceRulesetUpdate :=  // SpaceRulesetUpdate | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.SpaceRulesets.Update(context.Background(), id, spaceRulesetUpdate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceRulesets.Update`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `SpaceRulesets.Update`: %v\n", resp)
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

 **spaceRulesetUpdate** | [**SpaceRulesetUpdate**](SpaceRulesetUpdate.md) |  | 

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

