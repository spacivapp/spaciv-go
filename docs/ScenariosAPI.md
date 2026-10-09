# \ScenariosAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ScenariosAPI.md#Create) | **Post** /v1/scenarios | Create a scenario
[**Delete**](ScenariosAPI.md#Delete) | **Delete** /v1/scenarios/{id} | Delete or archive a scenario
[**Duplicate**](ScenariosAPI.md#Duplicate) | **Post** /v1/scenarios/{id}/duplicate | Duplicate a scenario
[**Get**](ScenariosAPI.md#Get) | **Get** /v1/scenarios/{id} | Get a scenario
[**List**](ScenariosAPI.md#List) | **Get** /v1/scenarios | List scenarios
[**ListLayerVariables**](ScenariosAPI.md#ListLayerVariables) | **Get** /v1/layer-variables | List account variables used by scenario layers
[**Update**](ScenariosAPI.md#Update) | **Patch** /v1/scenarios/{id} | Update a scenario



## Create

> string Create(ctx, scenarioCreate)

Create a scenario



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
	scenarioCreate :=  // ScenarioCreate | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Scenarios.Create(context.Background(), scenarioCreate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scenarios.Create`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scenarios.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scenarioCreate** | [**ScenarioCreate**](ScenarioCreate.md) |  | 

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

> map[string]interface{} Delete(ctx, id, params)

Delete or archive a scenario



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
	resp, err := client.Scenarios.Delete(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scenarios.Delete`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scenarios.Delete`: %v\n", resp)
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

 **action** | **string** | Lifecycle action: \&quot;archive\&quot; moves the scenario to the archive, \&quot;restore\&quot; returns it to active, and \&quot;purge\&quot; permanently deletes it. Defaults to \&quot;purge\&quot; when omitted. | 

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


## Duplicate

> string Duplicate(ctx, id)

Duplicate a scenario



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
	resp, err := client.Scenarios.Duplicate(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scenarios.Duplicate`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scenarios.Duplicate`: %v\n", resp)
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

> ScenarioChangelogBody Get(ctx, id)

Get a scenario



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
	resp, err := client.Scenarios.Get(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scenarios.Get`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scenarios.Get`: %v\n", resp)
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

[**ScenarioChangelogBody**](ScenarioChangelogBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ScenarioListRes List(ctx, params)

List scenarios



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
	resp, err := client.Scenarios.List(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scenarios.List`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scenarios.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showArchived** | **bool** |  | 

### Return type

[**ScenarioListRes**](ScenarioListRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLayerVariables

> LayerVariablesRes ListLayerVariables(ctx)

List account variables used by scenario layers



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
	resp, err := client.Scenarios.ListLayerVariables(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scenarios.ListLayerVariables`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scenarios.ListLayerVariables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLayerVariablesRequest struct via the builder pattern


### Return type

[**LayerVariablesRes**](LayerVariablesRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> map[string]interface{} Update(ctx, id, scenarioPatch)

Update a scenario



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
	scenarioPatch :=  // ScenarioPatch | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Scenarios.Update(context.Background(), id, scenarioPatch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scenarios.Update`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scenarios.Update`: %v\n", resp)
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

 **scenarioPatch** | [**ScenarioPatch**](ScenarioPatch.md) |  | 

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

