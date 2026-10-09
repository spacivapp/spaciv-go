# \FilesAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountFiles**](FilesAPI.md#CreateAccountFiles) | **Post** /v1/account/files | Upload a file to the account
[**DeleteAccountFiles**](FilesAPI.md#DeleteAccountFiles) | **Delete** /v1/account/files/{id} | Delete an account file
[**GetAccountFilesContent**](FilesAPI.md#GetAccountFilesContent) | **Get** /v1/account/files/{id}/content | Download an account file
[**ListAccountFiles**](FilesAPI.md#ListAccountFiles) | **Get** /v1/account/files | List the account&#39;s files
[**UpdateAccountFiles**](FilesAPI.md#UpdateAccountFiles) | **Patch** /v1/account/files/{id} | Rename an account file



## CreateAccountFiles

> Record CreateAccountFiles(ctx, params)

Upload a file to the account



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
	resp, err := client.Files.CreateAccountFiles(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Files.CreateAccountFiles`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Files.CreateAccountFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filename** | ***os.File** | filename of the file being uploaded | 
 **name** | **string** | general purpose name for multipart form value | 

### Return type

[**Record**](Record.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountFiles

> map[string]interface{} DeleteAccountFiles(ctx, id)

Delete an account file



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
	resp, err := client.Files.DeleteAccountFiles(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Files.DeleteAccountFiles`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Files.DeleteAccountFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountFilesRequest struct via the builder pattern


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


## GetAccountFilesContent

> string GetAccountFilesContent(ctx, id)

Download an account file



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
	resp, err := client.Files.GetAccountFilesContent(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Files.GetAccountFilesContent`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Files.GetAccountFilesContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountFilesContentRequest struct via the builder pattern


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


## ListAccountFiles

> ListPage ListAccountFiles(ctx, params)

List the account's files



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
	resp, err := client.Files.ListAccountFiles(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Files.ListAccountFiles`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Files.ListAccountFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of files to return in this page. | [default to 50]
 **cursor** | **string** | Opaque pagination cursor returned as nextCursor by a previous response. A cursor is only valid for the account it was issued to. | 

### Return type

[**ListPage**](ListPage.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountFiles

> Record UpdateAccountFiles(ctx, id, accountFilePatchBody)

Rename an account file



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
	accountFilePatchBody :=  // AccountFilePatchBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Files.UpdateAccountFiles(context.Background(), id, accountFilePatchBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Files.UpdateAccountFiles`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Files.UpdateAccountFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountFilePatchBody** | [**AccountFilePatchBody**](AccountFilePatchBody.md) |  | 

### Return type

[**Record**](Record.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

