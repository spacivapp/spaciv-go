# \LocalisationAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTranslations**](LocalisationAPI.md#CreateTranslations) | **Post** /v1/translations | Get translations
[**ListLocales**](LocalisationAPI.md#ListLocales) | **Get** /v1/locales | List available locales



## CreateTranslations

> map[string]string CreateTranslations(ctx, i18nTranslationsParamsWrapperBody)

Get translations



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
	i18nTranslationsParamsWrapperBody :=  // I18nTranslationsParamsWrapperBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Localisation.CreateTranslations(context.Background(), i18nTranslationsParamsWrapperBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Localisation.CreateTranslations`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Localisation.CreateTranslations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTranslationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **i18nTranslationsParamsWrapperBody** | [**I18nTranslationsParamsWrapperBody**](I18nTranslationsParamsWrapperBody.md) |  | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocales

> []IDNameStringString ListLocales(ctx)

List available locales



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
	resp, err := client.Localisation.ListLocales(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Localisation.ListLocales`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Localisation.ListLocales`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLocalesRequest struct via the builder pattern


### Return type

[**[]IDNameStringString**](IDNameStringString.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

