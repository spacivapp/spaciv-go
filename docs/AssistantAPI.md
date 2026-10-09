# \AssistantAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChat**](AssistantAPI.md#CreateChat) | **Post** /v1/assistant/chat | Send a chat message
[**DeleteConversations**](AssistantAPI.md#DeleteConversations) | **Delete** /v1/assistant/conversations/{id} | Delete a conversation
[**GetConfig**](AssistantAPI.md#GetConfig) | **Get** /v1/assistant/config | Get the assistant configuration
[**GetConversationsTranscript**](AssistantAPI.md#GetConversationsTranscript) | **Get** /v1/assistant/conversations/{id}/transcript | Get a conversation&#39;s transcript
[**ListConversations**](AssistantAPI.md#ListConversations) | **Get** /v1/assistant/conversations | List conversations
[**ListModels**](AssistantAPI.md#ListModels) | **Get** /v1/assistant/models | List available assistant models
[**PollConversations**](AssistantAPI.md#PollConversations) | **Get** /v1/assistant/conversations/{id}/poll | Poll a conversation
[**StopConversations**](AssistantAPI.md#StopConversations) | **Post** /v1/assistant/conversations/{id}/stop | Stop the running turn
[**UpdateConfig**](AssistantAPI.md#UpdateConfig) | **Put** /v1/assistant/config | Set the assistant configuration



## CreateChat

> AssistantChatRes CreateChat(ctx, assistantChatReq)

Send a chat message



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
	assistantChatReq :=  // AssistantChatReq | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Assistant.CreateChat(context.Background(), assistantChatReq)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Assistant.CreateChat`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Assistant.CreateChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assistantChatReq** | [**AssistantChatReq**](AssistantChatReq.md) |  | 

### Return type

[**AssistantChatRes**](AssistantChatRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConversations

> DeleteConversations(ctx, id)

Delete a conversation



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
	id :=  // string | UUID of the conversation to delete.

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.Assistant.DeleteConversations(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Assistant.DeleteConversations`: %v\n", err)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UUID of the conversation to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConversationsRequest struct via the builder pattern


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


## GetConfig

> AssistantConfigBody GetConfig(ctx)

Get the assistant configuration



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
	resp, err := client.Assistant.GetConfig(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Assistant.GetConfig`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Assistant.GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


### Return type

[**AssistantConfigBody**](AssistantConfigBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsTranscript

> AssistantTranscriptBody GetConversationsTranscript(ctx, id)

Get a conversation's transcript



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
	id :=  // string | UUID of the conversation to read.

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Assistant.GetConversationsTranscript(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Assistant.GetConversationsTranscript`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Assistant.GetConversationsTranscript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UUID of the conversation to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsTranscriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssistantTranscriptBody**](AssistantTranscriptBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversations

> AssistantConversationListRes ListConversations(ctx)

List conversations



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
	resp, err := client.Assistant.ListConversations(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Assistant.ListConversations`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Assistant.ListConversations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListConversationsRequest struct via the builder pattern


### Return type

[**AssistantConversationListRes**](AssistantConversationListRes.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListModels

> AssistantModelListBody ListModels(ctx)

List available assistant models



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
	resp, err := client.Assistant.ListModels(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Assistant.ListModels`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Assistant.ListModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListModelsRequest struct via the builder pattern


### Return type

[**AssistantModelListBody**](AssistantModelListBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PollConversations

> AssistantChatPollBody PollConversations(ctx, id, params)

Poll a conversation



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
	id :=  // string | UUID of the conversation to poll.

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Assistant.PollConversations(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Assistant.PollConversations`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Assistant.PollConversations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UUID of the conversation to poll. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPollConversationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **string** | Id of the newest message the caller already has. Omit to read the conversation from its start. | 

### Return type

[**AssistantChatPollBody**](AssistantChatPollBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopConversations

> StopConversations(ctx, id)

Stop the running turn



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
	id :=  // string | UUID of the conversation whose turn to stop.

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	err := client.Assistant.StopConversations(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Assistant.StopConversations`: %v\n", err)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UUID of the conversation whose turn to stop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopConversationsRequest struct via the builder pattern


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


## UpdateConfig

> map[string]interface{} UpdateConfig(ctx, assistantConfigPutReq)

Set the assistant configuration



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
	assistantConfigPutReq :=  // AssistantConfigPutReq | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Assistant.UpdateConfig(context.Background(), assistantConfigPutReq)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Assistant.UpdateConfig`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Assistant.UpdateConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assistantConfigPutReq** | [**AssistantConfigPutReq**](AssistantConfigPutReq.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

