# \BrandingAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccountLoginBackground**](BrandingAPI.md#DeleteAccountLoginBackground) | **Delete** /v1/account/login-background | Delete the sign-in page background
[**DeleteAccountLogo**](BrandingAPI.md#DeleteAccountLogo) | **Delete** /v1/account/logo | Delete the account logo
[**Get**](BrandingAPI.md#Get) | **Get** /v1/branding/{hostname} | Get sign-in page branding for a hostname
[**GetAccountLoginBackground**](BrandingAPI.md#GetAccountLoginBackground) | **Get** /v1/account/login-background | Get the sign-in page background
[**GetAccountLogo**](BrandingAPI.md#GetAccountLogo) | **Get** /v1/account/logo | Get the account logo
[**GetLoginBackground**](BrandingAPI.md#GetLoginBackground) | **Get** /v1/branding/{hostname}/login-background | Get the sign-in page background for a hostname
[**GetLogo**](BrandingAPI.md#GetLogo) | **Get** /v1/branding/{hostname}/logo | Get the account logo for a hostname
[**UpdateAccountLoginBackground**](BrandingAPI.md#UpdateAccountLoginBackground) | **Put** /v1/account/login-background | Upload the sign-in page background
[**UpdateAccountLogo**](BrandingAPI.md#UpdateAccountLogo) | **Put** /v1/account/logo | Upload the account logo



## DeleteAccountLoginBackground

> DeleteAccountLoginBackground(ctx)

Delete the sign-in page background



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
	err := client.Branding.DeleteAccountLoginBackground(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Branding.DeleteAccountLoginBackground`: %v\n", err)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountLoginBackgroundRequest struct via the builder pattern


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


## DeleteAccountLogo

> DeleteAccountLogo(ctx, params)

Delete the account logo



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
	err := client.Branding.DeleteAccountLogo(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Branding.DeleteAccountLogo`: %v\n", err)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 

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


## Get

> Branding Get(ctx, hostname)

Get sign-in page branding for a hostname



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
	hostname :=  // string | Subdomain the user is signing in on.

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Branding.Get(context.Background(), hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Branding.Get`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Branding.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostname** | **string** | Subdomain the user is signing in on. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Branding**](Branding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountLoginBackground

> string GetAccountLoginBackground(ctx)

Get the sign-in page background



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
	resp, err := client.Branding.GetAccountLoginBackground(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Branding.GetAccountLoginBackground`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Branding.GetAccountLoginBackground`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountLoginBackgroundRequest struct via the builder pattern


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


## GetAccountLogo

> string GetAccountLogo(ctx, params)

Get the account logo



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
	resp, err := client.Branding.GetAccountLogo(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Branding.GetAccountLogo`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Branding.GetAccountLogo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Which logo variant to retrieve. Falls back to the opposite variant unless noFallback is true. | 
 **id** | **int64** | Account ID to retrieve the logo for. Only honoured when the caller has the *Manage accounts* permission; otherwise the caller&#39;s own account is used. | 
 **noFallback** | **bool** | When true, returns 404 instead of falling back to the opposite variant if the requested type has no logo. | 

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


## GetLoginBackground

> string GetLoginBackground(ctx, hostname, params)

Get the sign-in page background for a hostname



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
	hostname :=  // string | Subdomain the user is signing in on.

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Branding.GetLoginBackground(context.Background(), hostname, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Branding.GetLoginBackground`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Branding.GetLoginBackground`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostname** | **string** | Subdomain the user is signing in on. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginBackgroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v** | **string** | The backgroundVersion from GET /v1/branding/:hostname, used only to bust caches. | 
 **ifNoneMatch** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogo

> string GetLogo(ctx, hostname, params)

Get the account logo for a hostname



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
	hostname :=  // string | Subdomain the user is signing in on.

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Branding.GetLogo(context.Background(), hostname, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Branding.GetLogo`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Branding.GetLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostname** | **string** | Subdomain the user is signing in on. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Which logo variant to retrieve. Falls back to the opposite variant when the requested one has no logo. | 
 **v** | **string** | The logoVersion from GET /v1/branding/:hostname, used only to bust caches. | 
 **ifNoneMatch** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountLoginBackground

> UpdateAccountLoginBackground(ctx, params)

Upload the sign-in page background



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
	err := client.Branding.UpdateAccountLoginBackground(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Branding.UpdateAccountLoginBackground`: %v\n", err)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountLoginBackgroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filename** | ***os.File** | filename of the file being uploaded | 
 **name** | **string** | general purpose name for multipart form value | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountLogo

> UpdateAccountLogo(ctx, params)

Upload the account logo



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
	err := client.Branding.UpdateAccountLogo(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Branding.UpdateAccountLogo`: %v\n", err)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filename** | ***os.File** | filename of the file being uploaded | 
 **name** | **string** | general purpose name for multipart form value | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

