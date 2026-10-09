# \ObjectGroupsAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ObjectGroupsAPI.md#Create) | **Post** /v1/object-groups | Create an object group
[**Delete**](ObjectGroupsAPI.md#Delete) | **Delete** /v1/object-groups/{id} | Delete or archive an object group
[**DeleteEntries**](ObjectGroupsAPI.md#DeleteEntries) | **Delete** /v1/object-groups/{id}/entries | Delete entries from an object group
[**DeleteEntriesImage**](ObjectGroupsAPI.md#DeleteEntriesImage) | **Delete** /v1/object-groups/{id}/entries/{entryId}/image | Delete an entry&#39;s image
[**Duplicate**](ObjectGroupsAPI.md#Duplicate) | **Post** /v1/object-groups/{id}/duplicate | Duplicate an object group
[**Get**](ObjectGroupsAPI.md#Get) | **Get** /v1/object-groups/{id} | Get an object group
[**GetEntriesImage**](ObjectGroupsAPI.md#GetEntriesImage) | **Get** /v1/object-groups/{id}/entries/{entryId}/image | Download an entry&#39;s image
[**GetFile**](ObjectGroupsAPI.md#GetFile) | **Get** /v1/object-groups/{id}/file | Download an object group&#39;s file
[**List**](ObjectGroupsAPI.md#List) | **Get** /v1/object-groups | List object groups
[**ListEntries**](ObjectGroupsAPI.md#ListEntries) | **Get** /v1/object-groups/{id}/entries | List entries in an object group
[**ReplaceEntries**](ObjectGroupsAPI.md#ReplaceEntries) | **Put** /v1/object-groups/{id}/entries | Upsert entries in an object group
[**SearchEntries**](ObjectGroupsAPI.md#SearchEntries) | **Get** /v1/object-groups/{id}/entries/search | Search entries in an object group
[**Update**](ObjectGroupsAPI.md#Update) | **Patch** /v1/object-groups/{id} | Update an object group
[**UpdateEntries**](ObjectGroupsAPI.md#UpdateEntries) | **Patch** /v1/object-groups/{id}/entries | Partially update entries in an object group
[**UpdateEntriesImage**](ObjectGroupsAPI.md#UpdateEntriesImage) | **Put** /v1/object-groups/{id}/entries/{entryId}/image | Upload an entry image
[**UpdateFile**](ObjectGroupsAPI.md#UpdateFile) | **Put** /v1/object-groups/{id}/file | Upload or replace an object group&#39;s file



## Create

> string Create(ctx, objectGroupChangelogBody)

Create an object group



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
	objectGroupChangelogBody :=  // ObjectGroupChangelogBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.ObjectGroups.Create(context.Background(), objectGroupChangelogBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.Create`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `ObjectGroups.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectGroupChangelogBody** | [**ObjectGroupChangelogBody**](ObjectGroupChangelogBody.md) |  | 

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

Delete or archive an object group



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
	resp, err := client.ObjectGroups.Delete(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.Delete`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `ObjectGroups.Delete`: %v\n", resp)
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

 **action** | **string** | Lifecycle transition to apply: \&quot;archive\&quot; moves the group to the archive, \&quot;restore\&quot; returns it to active, omitting the parameter (or \&quot;purge\&quot;) permanently deletes the group and all its entries. | 

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


## DeleteEntries

> DeleteEntries(ctx, id, objectGroupEntryUpdate)

Delete entries from an object group



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
	objectGroupEntryUpdate :=  // []ObjectGroupEntryUpdate | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.ObjectGroups.DeleteEntries(context.Background(), id, objectGroupEntryUpdate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.DeleteEntries`: %v\n", err)
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

 **objectGroupEntryUpdate** | [**[]ObjectGroupEntryUpdate**](ObjectGroupEntryUpdate.md) |  | 

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


## DeleteEntriesImage

> map[string]interface{} DeleteEntriesImage(ctx, id, entryId)

Delete an entry's image



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
	entryId :=  // string | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.ObjectGroups.DeleteEntriesImage(context.Background(), id, entryId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.DeleteEntriesImage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `ObjectGroups.DeleteEntriesImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntriesImageRequest struct via the builder pattern


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


## Duplicate

> string Duplicate(ctx, id)

Duplicate an object group



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
	resp, err := client.ObjectGroups.Duplicate(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.Duplicate`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `ObjectGroups.Duplicate`: %v\n", resp)
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

> ObjectGroupChangelogBody Get(ctx, id)

Get an object group



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
	resp, err := client.ObjectGroups.Get(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.Get`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `ObjectGroups.Get`: %v\n", resp)
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

[**ObjectGroupChangelogBody**](ObjectGroupChangelogBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntriesImage

> string GetEntriesImage(ctx, id, entryId, params)

Download an entry's image



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
	entryId :=  // string | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.ObjectGroups.GetEntriesImage(context.Background(), id, entryId, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.GetEntriesImage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `ObjectGroups.GetEntriesImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntriesImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withDefault** | **string** | When set to a recognised default image key, the server falls back to that default image if no entry-specific image exists, rather than returning an error. | 

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


## GetFile

> string GetFile(ctx, id)

Download an object group's file



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
	resp, err := client.ObjectGroups.GetFile(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.GetFile`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `ObjectGroups.GetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileRequest struct via the builder pattern


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

> ObjectGroupListRes List(ctx, params)

List object groups



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
	resp, err := client.ObjectGroups.List(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.List`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `ObjectGroups.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showArchived** | **bool** | When true, returns archived groups instead of active ones. | 

### Return type

[**ObjectGroupListRes**](ObjectGroupListRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntries

> ObjectGroupEntriesRes ListEntries(ctx, id, params)

List entries in an object group



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
	resp, err := client.ObjectGroups.ListEntries(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.ListEntries`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `ObjectGroups.ListEntries`: %v\n", resp)
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

 **start** | **int32** | Zero-based index of the first entry to return. Defaults to 0. Prefer cursor: reaching a later index costs a read of every entry before it. | 
 **limit** | **int32** | Maximum number of entries to return. Defaults to 100 when omitted or 0. | 
 **cursor** | **string** | Opaque cursor from a previous response. Takes precedence over start. | 

### Return type

[**ObjectGroupEntriesRes**](ObjectGroupEntriesRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceEntries

> ReplaceEntries(ctx, id, objectGroupEntryUpdate)

Upsert entries in an object group



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
	objectGroupEntryUpdate :=  // []ObjectGroupEntryUpdate | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.ObjectGroups.ReplaceEntries(context.Background(), id, objectGroupEntryUpdate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.ReplaceEntries`: %v\n", err)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectGroupEntryUpdate** | [**[]ObjectGroupEntryUpdate**](ObjectGroupEntryUpdate.md) |  | 

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


## SearchEntries

> ObjectGroupEntrySearchRes SearchEntries(ctx, id, params)

Search entries in an object group



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
	resp, err := client.ObjectGroups.SearchEntries(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.SearchEntries`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `ObjectGroups.SearchEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Matches entries having, for each word of the query, a word containing it either from its start or from a point with at least three characters left, within its first 40 characters. So \&quot;oo\&quot; and \&quot;oom\&quot; find \&quot;Meeting Room\&quot; but \&quot;om\&quot; does not. Word order does not matter. Omit it to browse the group from the start. | 
 **ids** | **[]string** | Entry ids to resolve by name instead of searching, repeated once per id. At most 200 per request; resolve a longer selection in batches. | 
 **limit** | **int32** | Maximum number of entries to return. Defaults to 25 when omitted or 0, and is capped at 100. | 
 **cursor** | **string** | Opaque cursor from a previous response. | 

### Return type

[**ObjectGroupEntrySearchRes**](ObjectGroupEntrySearchRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(ctx, id, objectGroupChangelogBody)

Update an object group



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
	objectGroupChangelogBody :=  // ObjectGroupChangelogBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.ObjectGroups.Update(context.Background(), id, objectGroupChangelogBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.Update`: %v\n", err)
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

 **objectGroupChangelogBody** | [**ObjectGroupChangelogBody**](ObjectGroupChangelogBody.md) |  | 

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

> UpdateEntries(ctx, id, entryPatch)

Partially update entries in an object group



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
	entryPatch :=  // []EntryPatch | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.ObjectGroups.UpdateEntries(context.Background(), id, entryPatch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.UpdateEntries`: %v\n", err)
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

 **entryPatch** | [**[]EntryPatch**](EntryPatch.md) |  | 

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


## UpdateEntriesImage

> map[string]interface{} UpdateEntriesImage(ctx, id, entryId, params)

Upload an entry image



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
	entryId :=  // string | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.ObjectGroups.UpdateEntriesImage(context.Background(), id, entryId, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.UpdateEntriesImage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `ObjectGroups.UpdateEntriesImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntriesImageRequest struct via the builder pattern


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


## UpdateFile

> map[string]interface{} UpdateFile(ctx, id, params)

Upload or replace an object group's file



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
	resp, err := client.ObjectGroups.UpdateFile(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroups.UpdateFile`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `ObjectGroups.UpdateFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileRequest struct via the builder pattern


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

