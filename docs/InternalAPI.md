# \InternalAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdminFeatureGroups**](InternalAPI.md#CreateAdminFeatureGroups) | **Post** /v1/admin/feature-groups | List feature groups
[**CreateAdminFlushDeadLetters**](InternalAPI.md#CreateAdminFlushDeadLetters) | **Post** /v1/admin/flush-dead-letters | Flush dead-letter events
[**CreateAdminPopulate**](InternalAPI.md#CreateAdminPopulate) | **Post** /v1/admin/populate | Populate account from template
[**CreateAdminReconcileFileStorage**](InternalAPI.md#CreateAdminReconcileFileStorage) | **Post** /v1/admin/reconcile-file-storage | Reconcile account file storage
[**CreateAdminRecreateProperties**](InternalAPI.md#CreateAdminRecreateProperties) | **Post** /v1/admin/recreate-properties | Recreate properties from datasets
[**CreateAdminRepairFloorplanReferences**](InternalAPI.md#CreateAdminRepairFloorplanReferences) | **Post** /v1/admin/repair-floorplan-references | Repair floorplan references
[**CreateAdminRetries**](InternalAPI.md#CreateAdminRetries) | **Post** /v1/admin/retries | List pending retries
[**CreateAdminUpdateFeatureGroups**](InternalAPI.md#CreateAdminUpdateFeatureGroups) | **Post** /v1/admin/update-feature-groups | Update a feature flag group assignment
[**CreateTranslationsExport**](InternalAPI.md#CreateTranslationsExport) | **Post** /v1/translations/export | Export translations as CSV
[**CreateTranslationsImport**](InternalAPI.md#CreateTranslationsImport) | **Post** /v1/translations/import | Import translations from CSV
[**ListAdminRegeneratePresets**](InternalAPI.md#ListAdminRegeneratePresets) | **Get** /v1/admin/regenerate-presets | Regenerate customer preset assets



## CreateAdminFeatureGroups

> FeatureGroupsResBody CreateAdminFeatureGroups(ctx)

List feature groups



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
	resp, err := client.Internal.CreateAdminFeatureGroups(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Internal.CreateAdminFeatureGroups`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Internal.CreateAdminFeatureGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdminFeatureGroupsRequest struct via the builder pattern


### Return type

[**FeatureGroupsResBody**](FeatureGroupsResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdminFlushDeadLetters

> string CreateAdminFlushDeadLetters(ctx)

Flush dead-letter events



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
	resp, err := client.Internal.CreateAdminFlushDeadLetters(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Internal.CreateAdminFlushDeadLetters`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Internal.CreateAdminFlushDeadLetters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdminFlushDeadLettersRequest struct via the builder pattern


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


## CreateAdminPopulate

> CreateAdminPopulate(ctx, body)

Populate account from template



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
	body :=  // int32 | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.Internal.CreateAdminPopulate(context.Background(), body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Internal.CreateAdminPopulate`: %v\n", err)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdminPopulateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **int32** |  | 

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


## CreateAdminReconcileFileStorage

> ReconcileReport CreateAdminReconcileFileStorage(ctx, params)

Reconcile account file storage



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
	resp, err := client.Internal.CreateAdminReconcileFileStorage(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Internal.CreateAdminReconcileFileStorage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Internal.CreateAdminReconcileFileStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdminReconcileFileStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apply** | **bool** | Delete the orphaned objects that qualify. Omit to report what would be deleted and change nothing. | 
 **account** | **int64** | Reconcile only this account. Omit to reconcile every account and every object in the bucket. | 

### Return type

[**ReconcileReport**](ReconcileReport.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdminRecreateProperties

> CreateAdminRecreateProperties(ctx)

Recreate properties from datasets



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
	err := client.Internal.CreateAdminRecreateProperties(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Internal.CreateAdminRecreateProperties`: %v\n", err)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdminRecreatePropertiesRequest struct via the builder pattern


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


## CreateAdminRepairFloorplanReferences

> RepairFloorplansResBody CreateAdminRepairFloorplanReferences(ctx)

Repair floorplan references



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
	resp, err := client.Internal.CreateAdminRepairFloorplanReferences(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Internal.CreateAdminRepairFloorplanReferences`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Internal.CreateAdminRepairFloorplanReferences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdminRepairFloorplanReferencesRequest struct via the builder pattern


### Return type

[**RepairFloorplansResBody**](RepairFloorplansResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdminRetries

> RetriesResBody CreateAdminRetries(ctx)

List pending retries



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
	resp, err := client.Internal.CreateAdminRetries(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Internal.CreateAdminRetries`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Internal.CreateAdminRetries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdminRetriesRequest struct via the builder pattern


### Return type

[**RetriesResBody**](RetriesResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdminUpdateFeatureGroups

> CreateAdminUpdateFeatureGroups(ctx, sysAdminUpdateFeatureGroupsParamsWrapperBody)

Update a feature flag group assignment



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
	sysAdminUpdateFeatureGroupsParamsWrapperBody :=  // SysAdminUpdateFeatureGroupsParamsWrapperBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.Internal.CreateAdminUpdateFeatureGroups(context.Background(), sysAdminUpdateFeatureGroupsParamsWrapperBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Internal.CreateAdminUpdateFeatureGroups`: %v\n", err)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdminUpdateFeatureGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sysAdminUpdateFeatureGroupsParamsWrapperBody** | [**SysAdminUpdateFeatureGroupsParamsWrapperBody**](SysAdminUpdateFeatureGroupsParamsWrapperBody.md) |  | 

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


## CreateTranslationsExport

> string CreateTranslationsExport(ctx, i18nTranslationsExportParamsWrapperBody)

Export translations as CSV



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
	i18nTranslationsExportParamsWrapperBody :=  // I18nTranslationsExportParamsWrapperBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Internal.CreateTranslationsExport(context.Background(), i18nTranslationsExportParamsWrapperBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Internal.CreateTranslationsExport`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Internal.CreateTranslationsExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTranslationsExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **i18nTranslationsExportParamsWrapperBody** | [**I18nTranslationsExportParamsWrapperBody**](I18nTranslationsExportParamsWrapperBody.md) |  | 

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


## CreateTranslationsImport

> map[string]interface{} CreateTranslationsImport(ctx, i18nTranslationsImportParamsWrapperBody)

Import translations from CSV



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
	i18nTranslationsImportParamsWrapperBody :=  // I18nTranslationsImportParamsWrapperBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Internal.CreateTranslationsImport(context.Background(), i18nTranslationsImportParamsWrapperBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Internal.CreateTranslationsImport`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Internal.CreateTranslationsImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTranslationsImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **i18nTranslationsImportParamsWrapperBody** | [**I18nTranslationsImportParamsWrapperBody**](I18nTranslationsImportParamsWrapperBody.md) |  | 

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


## ListAdminRegeneratePresets

> ListAdminRegeneratePresets(ctx)

Regenerate customer preset assets



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
	err := client.Internal.ListAdminRegeneratePresets(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Internal.ListAdminRegeneratePresets`: %v\n", err)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAdminRegeneratePresetsRequest struct via the builder pattern


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

