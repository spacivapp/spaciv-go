# \LogicLayersAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGraphLayers**](LogicLayersAPI.md#CreateGraphLayers) | **Post** /v1/graph-layers | Create a graph layer
[**CreateSpaceRulesets**](LogicLayersAPI.md#CreateSpaceRulesets) | **Post** /v1/space-rulesets | Create a logic layer
[**CreateSpaceRulesetsImage**](LogicLayersAPI.md#CreateSpaceRulesetsImage) | **Post** /v1/space-rulesets/{id}/image | Upload a space ruleset image
[**DeleteGraphLayers**](LogicLayersAPI.md#DeleteGraphLayers) | **Delete** /v1/graph-layers/{id} | Delete a graph layer
[**DeleteManyGraphLayers**](LogicLayersAPI.md#DeleteManyGraphLayers) | **Delete** /v1/graph-layers | Delete graph layers
[**DeleteSpaceRulesets**](LogicLayersAPI.md#DeleteSpaceRulesets) | **Delete** /v1/space-rulesets/{id} | Delete a logic layer
[**DeleteSpaceRulesetsImage**](LogicLayersAPI.md#DeleteSpaceRulesetsImage) | **Delete** /v1/space-rulesets/{id}/image | Delete a space ruleset image
[**DuplicateGraphLayers**](LogicLayersAPI.md#DuplicateGraphLayers) | **Post** /v1/graph-layers/{id}/duplicate | Duplicate a graph layer
[**DuplicateSpaceRulesets**](LogicLayersAPI.md#DuplicateSpaceRulesets) | **Post** /v1/space-rulesets/{id}/duplicate | Duplicate a logic layer
[**GetGraphLayers**](LogicLayersAPI.md#GetGraphLayers) | **Get** /v1/graph-layers/{id} | Get a graph layer
[**GetGraphLayersUsedBy**](LogicLayersAPI.md#GetGraphLayersUsedBy) | **Get** /v1/graph-layers/{id}/used-by | List the graph layers that use a graph layer
[**GetSpaceRulesets**](LogicLayersAPI.md#GetSpaceRulesets) | **Get** /v1/space-rulesets/{id} | Get a logic layer
[**GetSpaceRulesetsImage**](LogicLayersAPI.md#GetSpaceRulesetsImage) | **Get** /v1/space-rulesets/{id}/image | Get a space ruleset image
[**ListGraphLayers**](LogicLayersAPI.md#ListGraphLayers) | **Get** /v1/graph-layers | List graph layers
[**ListSpaceRulesets**](LogicLayersAPI.md#ListSpaceRulesets) | **Get** /v1/space-rulesets | List logic layers
[**ListSpaceRulesetsPresets**](LogicLayersAPI.md#ListSpaceRulesetsPresets) | **Get** /v1/space-rulesets/presets | List preset logic layers
[**UpdateGraphLayers**](LogicLayersAPI.md#UpdateGraphLayers) | **Patch** /v1/graph-layers/{id} | Update a graph layer
[**UpdateSpaceRulesets**](LogicLayersAPI.md#UpdateSpaceRulesets) | **Patch** /v1/space-rulesets/{id} | Update a logic layer



## CreateGraphLayers

> PostResponseBody CreateGraphLayers(ctx, graphLayerCreate)

Create a graph layer



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
	graphLayerCreate :=  // GraphLayerCreate | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.LogicLayers.CreateGraphLayers(context.Background(), graphLayerCreate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.CreateGraphLayers`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.CreateGraphLayers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGraphLayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphLayerCreate** | [**GraphLayerCreate**](GraphLayerCreate.md) |  | 

### Return type

[**PostResponseBody**](PostResponseBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpaceRulesets

> map[string]interface{} CreateSpaceRulesets(ctx, spaceRulesetCreate)

Create a logic layer



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
	resp, err := client.LogicLayers.CreateSpaceRulesets(context.Background(), spaceRulesetCreate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.CreateSpaceRulesets`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.CreateSpaceRulesets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpaceRulesetsRequest struct via the builder pattern


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


## CreateSpaceRulesetsImage

> map[string]interface{} CreateSpaceRulesetsImage(ctx, id, params)

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
	resp, err := client.LogicLayers.CreateSpaceRulesetsImage(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.CreateSpaceRulesetsImage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.CreateSpaceRulesetsImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpaceRulesetsImageRequest struct via the builder pattern


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


## DeleteGraphLayers

> map[string]interface{} DeleteGraphLayers(ctx, id)

Delete a graph layer



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
	resp, err := client.LogicLayers.DeleteGraphLayers(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.DeleteGraphLayers`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.DeleteGraphLayers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGraphLayersRequest struct via the builder pattern


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


## DeleteManyGraphLayers

> GraphLayersDeleted DeleteManyGraphLayers(ctx, graphLayersToDelete)

Delete graph layers



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
	graphLayersToDelete :=  // GraphLayersToDelete | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.LogicLayers.DeleteManyGraphLayers(context.Background(), graphLayersToDelete)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.DeleteManyGraphLayers`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.DeleteManyGraphLayers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManyGraphLayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphLayersToDelete** | [**GraphLayersToDelete**](GraphLayersToDelete.md) |  | 

### Return type

[**GraphLayersDeleted**](GraphLayersDeleted.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSpaceRulesets

> map[string]interface{} DeleteSpaceRulesets(ctx, id)

Delete a logic layer



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
	resp, err := client.LogicLayers.DeleteSpaceRulesets(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.DeleteSpaceRulesets`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.DeleteSpaceRulesets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpaceRulesetsRequest struct via the builder pattern


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


## DeleteSpaceRulesetsImage

> map[string]interface{} DeleteSpaceRulesetsImage(ctx, id)

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
	resp, err := client.LogicLayers.DeleteSpaceRulesetsImage(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.DeleteSpaceRulesetsImage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.DeleteSpaceRulesetsImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpaceRulesetsImageRequest struct via the builder pattern


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


## DuplicateGraphLayers

> DuplicateResponseBody DuplicateGraphLayers(ctx, id)

Duplicate a graph layer



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
	resp, err := client.LogicLayers.DuplicateGraphLayers(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.DuplicateGraphLayers`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.DuplicateGraphLayers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuplicateGraphLayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DuplicateResponseBody**](DuplicateResponseBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DuplicateSpaceRulesets

> string DuplicateSpaceRulesets(ctx, id)

Duplicate a logic layer



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
	resp, err := client.LogicLayers.DuplicateSpaceRulesets(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.DuplicateSpaceRulesets`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.DuplicateSpaceRulesets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuplicateSpaceRulesetsRequest struct via the builder pattern


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


## GetGraphLayers

> GraphLayerJSON GetGraphLayers(ctx, id)

Get a graph layer



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
	resp, err := client.LogicLayers.GetGraphLayers(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.GetGraphLayers`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.GetGraphLayers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGraphLayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GraphLayerJSON**](GraphLayerJSON.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGraphLayersUsedBy

> []GraphLayerUser GetGraphLayersUsedBy(ctx, id)

List the graph layers that use a graph layer



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
	resp, err := client.LogicLayers.GetGraphLayersUsedBy(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.GetGraphLayersUsedBy`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.GetGraphLayersUsedBy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGraphLayersUsedByRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GraphLayerUser**](GraphLayerUser.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceRulesets

> SpaceRuleSetJSON GetSpaceRulesets(ctx, id)

Get a logic layer



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
	resp, err := client.LogicLayers.GetSpaceRulesets(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.GetSpaceRulesets`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.GetSpaceRulesets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceRulesetsRequest struct via the builder pattern


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


## GetSpaceRulesetsImage

> string GetSpaceRulesetsImage(ctx, id, params)

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
	resp, err := client.LogicLayers.GetSpaceRulesetsImage(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.GetSpaceRulesetsImage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.GetSpaceRulesetsImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceRulesetsImageRequest struct via the builder pattern


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


## ListGraphLayers

> []GraphLayerListEntry ListGraphLayers(ctx)

List graph layers



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
	resp, err := client.LogicLayers.ListGraphLayers(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.ListGraphLayers`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.ListGraphLayers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGraphLayersRequest struct via the builder pattern


### Return type

[**[]GraphLayerListEntry**](GraphLayerListEntry.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpaceRulesets

> []SpaceRulesetListEntry ListSpaceRulesets(ctx)

List logic layers



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
	resp, err := client.LogicLayers.ListSpaceRulesets(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.ListSpaceRulesets`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.ListSpaceRulesets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSpaceRulesetsRequest struct via the builder pattern


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


## ListSpaceRulesetsPresets

> []SpaceRuleSet ListSpaceRulesetsPresets(ctx)

List preset logic layers



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
	resp, err := client.LogicLayers.ListSpaceRulesetsPresets(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.ListSpaceRulesetsPresets`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.ListSpaceRulesetsPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSpaceRulesetsPresetsRequest struct via the builder pattern


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


## UpdateGraphLayers

> map[string]interface{} UpdateGraphLayers(ctx, id, graphLayerUpdate)

Update a graph layer



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
	graphLayerUpdate :=  // GraphLayerUpdate | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.LogicLayers.UpdateGraphLayers(context.Background(), id, graphLayerUpdate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.UpdateGraphLayers`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.UpdateGraphLayers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGraphLayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **graphLayerUpdate** | [**GraphLayerUpdate**](GraphLayerUpdate.md) |  | 

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


## UpdateSpaceRulesets

> map[string]interface{} UpdateSpaceRulesets(ctx, id, spaceRulesetUpdate)

Update a logic layer



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
	resp, err := client.LogicLayers.UpdateSpaceRulesets(context.Background(), id, spaceRulesetUpdate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicLayers.UpdateSpaceRulesets`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `LogicLayers.UpdateSpaceRulesets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpaceRulesetsRequest struct via the builder pattern


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

