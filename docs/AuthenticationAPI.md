# \AuthenticationAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiKeys**](AuthenticationAPI.md#CreateApiKeys) | **Post** /v1/api-keys | Create an API key
[**CreateAuthApiKeyToken**](AuthenticationAPI.md#CreateAuthApiKeyToken) | **Post** /v1/auth/api-key-token | Exchange an API key for an access token
[**CreateAuthLoginPassword**](AuthenticationAPI.md#CreateAuthLoginPassword) | **Post** /v1/auth/login-password | Sign in with email and password
[**CreateAuthLoginPublic**](AuthenticationAPI.md#CreateAuthLoginPublic) | **Post** /v1/auth/login-public | Sign in as a public contributor
[**CreateAuthLoginToken**](AuthenticationAPI.md#CreateAuthLoginToken) | **Post** /v1/auth/login-token | Sign in with a magic link token
[**CreateAuthLogout**](AuthenticationAPI.md#CreateAuthLogout) | **Post** /v1/auth/logout | Sign out
[**CreateAuthRefreshToken**](AuthenticationAPI.md#CreateAuthRefreshToken) | **Post** /v1/auth/refresh-token | Refresh an access token
[**CreateAuthSendPasswordReset**](AuthenticationAPI.md#CreateAuthSendPasswordReset) | **Post** /v1/auth/send-password-reset | Request a password reset
[**CreateAuthSendToken**](AuthenticationAPI.md#CreateAuthSendToken) | **Post** /v1/auth/send-token | Send a magic link sign-in email
[**DeleteApiKeys**](AuthenticationAPI.md#DeleteApiKeys) | **Delete** /v1/api-keys/{id} | Revoke an API key
[**GetAuthLoginOidc**](AuthenticationAPI.md#GetAuthLoginOidc) | **Get** /v1/auth/login-oidc | Sign in via OIDC
[**ListApiKeys**](AuthenticationAPI.md#ListApiKeys) | **Get** /v1/api-keys | List API keys
[**UpdateAuthPassword**](AuthenticationAPI.md#UpdateAuthPassword) | **Put** /v1/auth/password | Set a new password



## CreateApiKeys

> ApiKeyPostResBody CreateApiKeys(ctx, apiKeyPostReqBody)

Create an API key



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
	apiKeyPostReqBody :=  // ApiKeyPostReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Authentication.CreateApiKeys(context.Background(), apiKeyPostReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Authentication.CreateApiKeys`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Authentication.CreateApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKeyPostReqBody** | [**ApiKeyPostReqBody**](ApiKeyPostReqBody.md) |  | 

### Return type

[**ApiKeyPostResBody**](ApiKeyPostResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthApiKeyToken

> AuthAPIKeyTokenResBody CreateAuthApiKeyToken(ctx, authAPIKeyTokenReqBody)

Exchange an API key for an access token



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
	authAPIKeyTokenReqBody :=  // AuthAPIKeyTokenReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Authentication.CreateAuthApiKeyToken(context.Background(), authAPIKeyTokenReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Authentication.CreateAuthApiKeyToken`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Authentication.CreateAuthApiKeyToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthApiKeyTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authAPIKeyTokenReqBody** | [**AuthAPIKeyTokenReqBody**](AuthAPIKeyTokenReqBody.md) |  | 

### Return type

[**AuthAPIKeyTokenResBody**](AuthAPIKeyTokenResBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthLoginPassword

> AuthLoginResBody CreateAuthLoginPassword(ctx, authLoginPassReqBody)

Sign in with email and password



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
	authLoginPassReqBody :=  // AuthLoginPassReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Authentication.CreateAuthLoginPassword(context.Background(), authLoginPassReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Authentication.CreateAuthLoginPassword`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Authentication.CreateAuthLoginPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthLoginPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authLoginPassReqBody** | [**AuthLoginPassReqBody**](AuthLoginPassReqBody.md) |  | 

### Return type

[**AuthLoginResBody**](AuthLoginResBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthLoginPublic

> AuthLoginResBody CreateAuthLoginPublic(ctx, authLoginPublicReqBody)

Sign in as a public contributor



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
	authLoginPublicReqBody :=  // AuthLoginPublicReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Authentication.CreateAuthLoginPublic(context.Background(), authLoginPublicReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Authentication.CreateAuthLoginPublic`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Authentication.CreateAuthLoginPublic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthLoginPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authLoginPublicReqBody** | [**AuthLoginPublicReqBody**](AuthLoginPublicReqBody.md) |  | 

### Return type

[**AuthLoginResBody**](AuthLoginResBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthLoginToken

> AuthLoginResBody CreateAuthLoginToken(ctx, authLoginTokenReqBody)

Sign in with a magic link token



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
	authLoginTokenReqBody :=  // AuthLoginTokenReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Authentication.CreateAuthLoginToken(context.Background(), authLoginTokenReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Authentication.CreateAuthLoginToken`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Authentication.CreateAuthLoginToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthLoginTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authLoginTokenReqBody** | [**AuthLoginTokenReqBody**](AuthLoginTokenReqBody.md) |  | 

### Return type

[**AuthLoginResBody**](AuthLoginResBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthLogout

> CreateAuthLogout(ctx, params)

Sign out



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/spacivapp/spaciv-go"
)

func main() {

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.Authentication.CreateAuthLogout(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Authentication.CreateAuthLogout`: %v\n", err)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthLogoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostToken** | [**Cookie**](Cookie.md) |  | 
 **token** | [**Cookie**](Cookie.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthRefreshToken

> AuthRefreshTokenResBody CreateAuthRefreshToken(ctx, authRefreshTokenReqBody, params)

Refresh an access token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/spacivapp/spaciv-go"
)

func main() {
	authRefreshTokenReqBody :=  // AuthRefreshTokenReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Authentication.CreateAuthRefreshToken(context.Background(), authRefreshTokenReqBody, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Authentication.CreateAuthRefreshToken`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Authentication.CreateAuthRefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authRefreshTokenReqBody** | [**AuthRefreshTokenReqBody**](AuthRefreshTokenReqBody.md) |  | 
 **hostToken** | [**Cookie**](Cookie.md) |  | 
 **token** | [**Cookie**](Cookie.md) |  | 

### Return type

[**AuthRefreshTokenResBody**](AuthRefreshTokenResBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthSendPasswordReset

> CreateAuthSendPasswordReset(ctx, sendPasswordPutReqBody)

Request a password reset



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
	sendPasswordPutReqBody :=  // SendPasswordPutReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.Authentication.CreateAuthSendPasswordReset(context.Background(), sendPasswordPutReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Authentication.CreateAuthSendPasswordReset`: %v\n", err)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthSendPasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendPasswordPutReqBody** | [**SendPasswordPutReqBody**](SendPasswordPutReqBody.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthSendToken

> CreateAuthSendToken(ctx, authSendTokenReqBody)

Send a magic link sign-in email



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
	authSendTokenReqBody :=  // AuthSendTokenReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.Authentication.CreateAuthSendToken(context.Background(), authSendTokenReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Authentication.CreateAuthSendToken`: %v\n", err)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthSendTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authSendTokenReqBody** | [**AuthSendTokenReqBody**](AuthSendTokenReqBody.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiKeys

> DeleteApiKeys(ctx, id)

Revoke an API key



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
	err := client.Authentication.DeleteApiKeys(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Authentication.DeleteApiKeys`: %v\n", err)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeysRequest struct via the builder pattern


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


## GetAuthLoginOidc

> AuthLoginResBody GetAuthLoginOidc(ctx, params)

Sign in via OIDC



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
	resp, err := client.Authentication.GetAuthLoginOidc(context.Background(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Authentication.GetAuthLoginOidc`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Authentication.GetAuthLoginOidc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthLoginOidcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | 
 **hostname** | **string** | The account hostname that determines which OIDC provider to authenticate against. | 

### Return type

[**AuthLoginResBody**](AuthLoginResBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiKeys

> []APIKeyDetails ListApiKeys(ctx)

List API keys



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
	resp, err := client.Authentication.ListApiKeys(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Authentication.ListApiKeys`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Authentication.ListApiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiKeysRequest struct via the builder pattern


### Return type

[**[]APIKeyDetails**](APIKeyDetails.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthPassword

> UpdateAuthPassword(ctx, passwordPutReqBody)

Set a new password



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
	passwordPutReqBody :=  // PasswordPutReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.Authentication.UpdateAuthPassword(context.Background(), passwordPutReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Authentication.UpdateAuthPassword`: %v\n", err)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordPutReqBody** | [**PasswordPutReqBody**](PasswordPutReqBody.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

