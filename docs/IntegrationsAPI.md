# \IntegrationsAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSlack**](IntegrationsAPI.md#CreateSlack) | **Post** /v1/integrations/slack | Connect Slack



## CreateSlack

> string CreateSlack(ctx, integrationCreateSlackParamsWrapperBody)

Connect Slack



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
	integrationCreateSlackParamsWrapperBody :=  // IntegrationCreateSlackParamsWrapperBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Integrations.CreateSlack(context.Background(), integrationCreateSlackParamsWrapperBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Integrations.CreateSlack`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Integrations.CreateSlack`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSlackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationCreateSlackParamsWrapperBody** | [**IntegrationCreateSlackParamsWrapperBody**](IntegrationCreateSlackParamsWrapperBody.md) |  | 

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

