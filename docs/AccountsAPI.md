# \AccountsAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](AccountsAPI.md#Create) | **Post** /v1/accounts | Create an account
[**CreateOidcProvider**](AccountsAPI.md#CreateOidcProvider) | **Post** /v1/account/oidc-provider | Create the current account&#39;s OIDC provider
[**Delete**](AccountsAPI.md#Delete) | **Delete** /v1/account | Delete the current account
[**DeleteOidcProvider**](AccountsAPI.md#DeleteOidcProvider) | **Delete** /v1/account/oidc-provider | Delete the current account&#39;s OIDC provider
[**Get**](AccountsAPI.md#Get) | **Get** /v1/account | Get the current account
[**GetAuthOptions**](AccountsAPI.md#GetAuthOptions) | **Get** /v1/auth/options/{hostname} | List sign-in methods for a hostname
[**GetHostname**](AccountsAPI.md#GetHostname) | **Get** /v1/accounts/{id}/hostname | Get an account&#39;s hostname
[**GetOidcProvider**](AccountsAPI.md#GetOidcProvider) | **Get** /v1/account/oidc-provider | Get the current account&#39;s OIDC provider
[**GetProfile**](AccountsAPI.md#GetProfile) | **Get** /v1/account/profile | Get the current account&#39;s profile
[**GetSetupStatus**](AccountsAPI.md#GetSetupStatus) | **Get** /v1/accounts/{id}/setup-status | Wait for account setup to finish
[**GetUsage**](AccountsAPI.md#GetUsage) | **Get** /v1/account/usage | Get the current account&#39;s usage
[**List**](AccountsAPI.md#List) | **Get** /v1/accounts | List the accounts you belong to
[**ListAuthOptions**](AccountsAPI.md#ListAuthOptions) | **Get** /v1/account/auth-options | List the current account&#39;s sign-in methods
[**ListUsageCalculationCredits**](AccountsAPI.md#ListUsageCalculationCredits) | **Get** /v1/account/usage/calculation-credits | Get the account&#39;s credit usage
[**Update**](AccountsAPI.md#Update) | **Patch** /v1/account | Update the current account
[**UpdateAuthOptions**](AccountsAPI.md#UpdateAuthOptions) | **Patch** /v1/account/auth-options | Update the current account&#39;s sign-in methods
[**UpdateOidcProvider**](AccountsAPI.md#UpdateOidcProvider) | **Patch** /v1/account/oidc-provider | Update the current account&#39;s OIDC provider
[**UpdateOidcProviderClientSecret**](AccountsAPI.md#UpdateOidcProviderClientSecret) | **Put** /v1/account/oidc-provider/client-secret | Replace the current account&#39;s OIDC client secret



## Create

> int64 Create(ctx, accountPostReqBody)

Create an account



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
	accountPostReqBody :=  // AccountPostReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Accounts.Create(context.Background(), accountPostReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.Create`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Accounts.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountPostReqBody** | [**AccountPostReqBody**](AccountPostReqBody.md) |  | 

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


## CreateOidcProvider

> AccountOIDCProvider CreateOidcProvider(ctx, accountPostOIDCProviderReqBody)

Create the current account's OIDC provider



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
	accountPostOIDCProviderReqBody :=  // AccountPostOIDCProviderReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Accounts.CreateOidcProvider(context.Background(), accountPostOIDCProviderReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.CreateOidcProvider`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Accounts.CreateOidcProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOidcProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountPostOIDCProviderReqBody** | [**AccountPostOIDCProviderReqBody**](AccountPostOIDCProviderReqBody.md) |  | 

### Return type

[**AccountOIDCProvider**](AccountOIDCProvider.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx)

Delete the current account



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
	err := client.Accounts.Delete(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.Delete`: %v\n", err)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


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


## DeleteOidcProvider

> DeleteOidcProvider(ctx)

Delete the current account's OIDC provider



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
	err := client.Accounts.DeleteOidcProvider(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.DeleteOidcProvider`: %v\n", err)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOidcProviderRequest struct via the builder pattern


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

> Account Get(ctx)

Get the current account



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
	resp, err := client.Accounts.Get(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.Get`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Accounts.Get`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


### Return type

[**Account**](Account.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthOptions

> []AccountAuthOption GetAuthOptions(ctx, hostname)

List sign-in methods for a hostname



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
	resp, err := client.Accounts.GetAuthOptions(context.Background(), hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.GetAuthOptions`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Accounts.GetAuthOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostname** | **string** | Subdomain the user is signing in on. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AccountAuthOption**](AccountAuthOption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostname

> string GetHostname(ctx, id)

Get an account's hostname



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
	id :=  // int64 | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Accounts.GetHostname(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.GetHostname`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Accounts.GetHostname`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostnameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetOidcProvider

> AccountOIDCProvider GetOidcProvider(ctx)

Get the current account's OIDC provider



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
	resp, err := client.Accounts.GetOidcProvider(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.GetOidcProvider`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Accounts.GetOidcProvider`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOidcProviderRequest struct via the builder pattern


### Return type

[**AccountOIDCProvider**](AccountOIDCProvider.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfile

> AccountDetailsResBody GetProfile(ctx)

Get the current account's profile



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
	resp, err := client.Accounts.GetProfile(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.GetProfile`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Accounts.GetProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


### Return type

[**AccountDetailsResBody**](AccountDetailsResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSetupStatus

> GetSetupStatus(ctx, id)

Wait for account setup to finish



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
	id :=  // int64 | Identifier returned by POST /v1/accounts.

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.Accounts.GetSetupStatus(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.GetSetupStatus`: %v\n", err)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Identifier returned by POST /v1/accounts. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSetupStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetUsage

> AccountUsageResBody GetUsage(ctx)

Get the current account's usage



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
	resp, err := client.Accounts.GetUsage(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.GetUsage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Accounts.GetUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageRequest struct via the builder pattern


### Return type

[**AccountUsageResBody**](AccountUsageResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []ListAccountsForUserRow List(ctx)

List the accounts you belong to



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
	resp, err := client.Accounts.List(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.List`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Accounts.List`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


### Return type

[**[]ListAccountsForUserRow**](ListAccountsForUserRow.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthOptions

> AccountGetAuthOptionsResBody ListAuthOptions(ctx)

List the current account's sign-in methods



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
	resp, err := client.Accounts.ListAuthOptions(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.ListAuthOptions`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Accounts.ListAuthOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthOptionsRequest struct via the builder pattern


### Return type

[**AccountGetAuthOptionsResBody**](AccountGetAuthOptionsResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageCalculationCredits

> AccountCalculationUsageBody ListUsageCalculationCredits(ctx, params)

Get the account's credit usage



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
	resp, err := client.Accounts.ListUsageCalculationCredits(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.ListUsageCalculationCredits`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Accounts.ListUsageCalculationCredits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsageCalculationCreditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int64** | How many days of usage to return, ending today (UTC). | [default to 90]

### Return type

[**AccountCalculationUsageBody**](AccountCalculationUsageBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(ctx, accountPutReqBody)

Update the current account



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
	accountPutReqBody :=  // AccountPutReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.Accounts.Update(context.Background(), accountPutReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.Update`: %v\n", err)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountPutReqBody** | [**AccountPutReqBody**](AccountPutReqBody.md) |  | 

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


## UpdateAuthOptions

> []AccountAuthOption UpdateAuthOptions(ctx, accountPatchAuthOptionsReqBody)

Update the current account's sign-in methods



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
	accountPatchAuthOptionsReqBody :=  // AccountPatchAuthOptionsReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Accounts.UpdateAuthOptions(context.Background(), accountPatchAuthOptionsReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.UpdateAuthOptions`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Accounts.UpdateAuthOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountPatchAuthOptionsReqBody** | [**AccountPatchAuthOptionsReqBody**](AccountPatchAuthOptionsReqBody.md) |  | 

### Return type

[**[]AccountAuthOption**](AccountAuthOption.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOidcProvider

> AccountOIDCProvider UpdateOidcProvider(ctx, accountOIDCProviderPatch)

Update the current account's OIDC provider



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
	accountOIDCProviderPatch :=  // AccountOIDCProviderPatch | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Accounts.UpdateOidcProvider(context.Background(), accountOIDCProviderPatch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.UpdateOidcProvider`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Accounts.UpdateOidcProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOidcProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountOIDCProviderPatch** | [**AccountOIDCProviderPatch**](AccountOIDCProviderPatch.md) |  | 

### Return type

[**AccountOIDCProvider**](AccountOIDCProvider.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOidcProviderClientSecret

> UpdateOidcProviderClientSecret(ctx, accountPutOIDCClientSecretReqBody)

Replace the current account's OIDC client secret



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
	accountPutOIDCClientSecretReqBody :=  // AccountPutOIDCClientSecretReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.Accounts.UpdateOidcProviderClientSecret(context.Background(), accountPutOIDCClientSecretReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.UpdateOidcProviderClientSecret`: %v\n", err)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOidcProviderClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountPutOIDCClientSecretReqBody** | [**AccountPutOIDCClientSecretReqBody**](AccountPutOIDCClientSecretReqBody.md) |  | 

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

