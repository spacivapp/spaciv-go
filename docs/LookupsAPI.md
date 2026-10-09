# \LookupsAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAvailabilities**](LookupsAPI.md#ListAvailabilities) | **Get** /v1/lookups/availabilities | List availabilities
[**ListCompanySizes**](LookupsAPI.md#ListCompanySizes) | **Get** /v1/lookups/company-sizes | List company sizes
[**ListCountries**](LookupsAPI.md#ListCountries) | **Get** /v1/lookups/countries | List countries
[**ListIndustries**](LookupsAPI.md#ListIndustries) | **Get** /v1/lookups/industries | List industries
[**ListJobFunctions**](LookupsAPI.md#ListJobFunctions) | **Get** /v1/lookups/job-functions | List job functions
[**ListProducts**](LookupsAPI.md#ListProducts) | **Get** /v1/lookups/products | List products



## ListAvailabilities

> []string ListAvailabilities(ctx)

List availabilities



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
	resp, err := client.Lookups.ListAvailabilities(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Lookups.ListAvailabilities`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Lookups.ListAvailabilities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailabilitiesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanySizes

> []IDNameStringString ListCompanySizes(ctx)

List company sizes



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
	resp, err := client.Lookups.ListCompanySizes(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Lookups.ListCompanySizes`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Lookups.ListCompanySizes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCompanySizesRequest struct via the builder pattern


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


## ListCountries

> []IDNameStringString ListCountries(ctx)

List countries



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
	resp, err := client.Lookups.ListCountries(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Lookups.ListCountries`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Lookups.ListCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCountriesRequest struct via the builder pattern


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


## ListIndustries

> []IDNameStringString ListIndustries(ctx)

List industries



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
	resp, err := client.Lookups.ListIndustries(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Lookups.ListIndustries`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Lookups.ListIndustries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIndustriesRequest struct via the builder pattern


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


## ListJobFunctions

> []IDNameStringString ListJobFunctions(ctx)

List job functions



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
	resp, err := client.Lookups.ListJobFunctions(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Lookups.ListJobFunctions`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Lookups.ListJobFunctions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListJobFunctionsRequest struct via the builder pattern


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


## ListProducts

> []string ListProducts(ctx)

List products



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
	resp, err := client.Lookups.ListProducts(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Lookups.ListProducts`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Lookups.ListProducts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListProductsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

