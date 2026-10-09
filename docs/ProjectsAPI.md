# \ProjectsAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ProjectsAPI.md#Create) | **Post** /v1/projects | Create a project
[**Delete**](ProjectsAPI.md#Delete) | **Delete** /v1/project | Delete or archive a project
[**DeleteImage**](ProjectsAPI.md#DeleteImage) | **Delete** /v1/project/image | Delete the project image
[**Get**](ProjectsAPI.md#Get) | **Get** /v1/project | Get a project
[**GetImage**](ProjectsAPI.md#GetImage) | **Get** /v1/project/image | Get the project image
[**List**](ProjectsAPI.md#List) | **Get** /v1/projects | List projects
[**Update**](ProjectsAPI.md#Update) | **Patch** /v1/project | Update a project
[**UpdateImage**](ProjectsAPI.md#UpdateImage) | **Put** /v1/project/image | Upload a project image



## Create

> int64 Create(ctx, projectCreateBody)

Create a project



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
	projectCreateBody :=  // ProjectCreateBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Projects.Create(context.Background(), projectCreateBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Projects.Create`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Projects.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectCreateBody** | [**ProjectCreateBody**](ProjectCreateBody.md) |  | 

### Return type

**int64**

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, params)

Delete or archive a project



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
	err := client.Projects.Delete(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Projects.Delete`: %v\n", err)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Lifecycle action: \&quot;archive\&quot; moves the project to the archive, \&quot;restore\&quot; makes it active again. Omit (or pass \&quot;purge\&quot;) to permanently delete the project and all its dashboards. | 

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


## DeleteImage

> map[string]interface{} DeleteImage(ctx)

Delete the project image



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
	resp, err := client.Projects.DeleteImage(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Projects.DeleteImage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Projects.DeleteImage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern


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

> ProjectChangelogBody Get(ctx, params)

Get a project



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
	resp, err := client.Projects.Get(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Projects.Get`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Projects.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **int64** | When non-zero, fetches this built-in template project instead of the caller&#39;s current project. | 

### Return type

[**ProjectChangelogBody**](ProjectChangelogBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImage

> string GetImage(ctx, params)

Get the project image



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
	resp, err := client.Projects.GetImage(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Projects.GetImage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Projects.GetImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noFallback** | **bool** | When true, returns 404 if no custom image has been uploaded rather than generating a default image from the project name. | 

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

> ProjectListRes List(ctx, params)

List projects



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
	resp, err := client.Projects.List(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Projects.List`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Projects.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showTemplate** | **bool** | When true, lists the built-in template projects shared across all accounts instead of the caller&#39;s own projects. | 
 **showArchived** | **bool** | When true, returns only archived projects; omit or set false to return only active projects. | 

### Return type

[**ProjectListRes**](ProjectListRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(ctx, projectPatchBody)

Update a project



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
	projectPatchBody :=  // ProjectPatchBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.Projects.Update(context.Background(), projectPatchBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Projects.Update`: %v\n", err)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectPatchBody** | [**ProjectPatchBody**](ProjectPatchBody.md) |  | 

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


## UpdateImage

> map[string]interface{} UpdateImage(ctx, params)

Upload a project image



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
	resp, err := client.Projects.UpdateImage(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Projects.UpdateImage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Projects.UpdateImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageRequest struct via the builder pattern


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

