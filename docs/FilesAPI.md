# \FilesAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountDirectories**](FilesAPI.md#CreateAccountDirectories) | **Post** /v1/account/directories | Create a directory in the account&#39;s files
[**CreateAccountFiles**](FilesAPI.md#CreateAccountFiles) | **Post** /v1/account/files | Upload a file to the account
[**CreateAccountTextFiles**](FilesAPI.md#CreateAccountTextFiles) | **Post** /v1/account/text-files | Create an empty text file in the account&#39;s files
[**DeleteAccountDirectories**](FilesAPI.md#DeleteAccountDirectories) | **Delete** /v1/account/directories/{id} | Delete a directory
[**DeleteAccountFiles**](FilesAPI.md#DeleteAccountFiles) | **Delete** /v1/account/files/{id} | Delete an account file
[**GetAccountFilesContent**](FilesAPI.md#GetAccountFilesContent) | **Get** /v1/account/files/{id}/content | Download an account file
[**ListAccountDirectories**](FilesAPI.md#ListAccountDirectories) | **Get** /v1/account/directories | List the account&#39;s directories
[**ListAccountFiles**](FilesAPI.md#ListAccountFiles) | **Get** /v1/account/files | List the account&#39;s files
[**UpdateAccountDirectories**](FilesAPI.md#UpdateAccountDirectories) | **Patch** /v1/account/directories/{id} | Rename or move a directory
[**UpdateAccountFiles**](FilesAPI.md#UpdateAccountFiles) | **Patch** /v1/account/files/{id} | Rename or move an account file
[**UpdateAccountFilesContent**](FilesAPI.md#UpdateAccountFilesContent) | **Put** /v1/account/files/{id}/content | Replace an account file&#39;s content



## CreateAccountDirectories

> Directory CreateAccountDirectories(ctx, accountDirectoryPostBody)

Create a directory in the account's files



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
	accountDirectoryPostBody :=  // AccountDirectoryPostBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Files.CreateAccountDirectories(context.Background(), accountDirectoryPostBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Files.CreateAccountDirectories`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Files.CreateAccountDirectories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountDirectoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountDirectoryPostBody** | [**AccountDirectoryPostBody**](AccountDirectoryPostBody.md) |  | 

### Return type

[**Directory**](Directory.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## CreateAccountTextFiles

> Record CreateAccountTextFiles(ctx, accountTextFilePostBody)

Create an empty text file in the account's files



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
	accountTextFilePostBody :=  // AccountTextFilePostBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Files.CreateAccountTextFiles(context.Background(), accountTextFilePostBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Files.CreateAccountTextFiles`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Files.CreateAccountTextFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountTextFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountTextFilePostBody** | [**AccountTextFilePostBody**](AccountTextFilePostBody.md) |  | 

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


## DeleteAccountDirectories

> map[string]interface{} DeleteAccountDirectories(ctx, id)

Delete a directory



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
	resp, err := client.Files.DeleteAccountDirectories(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Files.DeleteAccountDirectories`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Files.DeleteAccountDirectories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountDirectoriesRequest struct via the builder pattern


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


## ListAccountDirectories

> DirectoryPage ListAccountDirectories(ctx, params)

List the account's directories



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
	resp, err := client.Files.ListAccountDirectories(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Files.ListAccountDirectories`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Files.ListAccountDirectories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountDirectoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of directories to return in this page. | [default to 50]
 **cursor** | **string** | Opaque pagination cursor returned as nextCursor by a previous response. A cursor is only valid for the account it was issued to. | 

### Return type

[**DirectoryPage**](DirectoryPage.md)

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
 **cursor** | **string** | Opaque pagination cursor returned as nextCursor by a previous response. A cursor is only valid for the account, and the parent, it was issued for. | 
 **parent** | **string** | List only what is directly inside this directory: a directory identifier, or root for the root. Leave it out to list every file of the account, whatever directory it is in. | 

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


## UpdateAccountDirectories

> Directory UpdateAccountDirectories(ctx, id, accountDirectoryPatchBody)

Rename or move a directory



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
	accountDirectoryPatchBody :=  // AccountDirectoryPatchBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Files.UpdateAccountDirectories(context.Background(), id, accountDirectoryPatchBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Files.UpdateAccountDirectories`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Files.UpdateAccountDirectories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountDirectoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountDirectoryPatchBody** | [**AccountDirectoryPatchBody**](AccountDirectoryPatchBody.md) |  | 

### Return type

[**Directory**](Directory.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountFiles

> Record UpdateAccountFiles(ctx, id, accountFilePatchBody)

Rename or move an account file



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


## UpdateAccountFilesContent

> Record UpdateAccountFilesContent(ctx, id, params)

Replace an account file's content



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
	resp, err := client.Files.UpdateAccountFilesContent(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Files.UpdateAccountFilesContent`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Files.UpdateAccountFilesContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountFilesContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | The ETag returned when the file&#39;s content was read. The content is replaced only if it has not changed since. | 
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

