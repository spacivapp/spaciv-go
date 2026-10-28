# \ScriptsAPI

All URIs are relative to *https://api.spaciv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ScriptsAPI.md#Create) | **Post** /v1/scripts | Create a script
[**CreateReauthorise**](ScriptsAPI.md#CreateReauthorise) | **Post** /v1/scripts/{id}/reauthorise | Re-authorise a script
[**CreateRegenerateWebhook**](ScriptsAPI.md#CreateRegenerateWebhook) | **Post** /v1/scripts/{id}/regenerate-webhook | Regenerate a script&#39;s webhook URL
[**CreateSchedulePreview**](ScriptsAPI.md#CreateSchedulePreview) | **Post** /v1/scripts/schedule-preview | Preview a script schedule
[**CreateScratchpadBuilds**](ScriptsAPI.md#CreateScratchpadBuilds) | **Post** /v1/scripts/scratchpad/builds | Build scratchpad code with its dependencies
[**CreateScratchpadRuns**](ScriptsAPI.md#CreateScratchpadRuns) | **Post** /v1/scripts/scratchpad/runs | Run scratchpad code
[**Delete**](ScriptsAPI.md#Delete) | **Delete** /v1/scripts/{id} | Delete a script
[**Get**](ScriptsAPI.md#Get) | **Get** /v1/scripts/{id} | Get a script
[**GetAudit**](ScriptsAPI.md#GetAudit) | **Get** /v1/scripts/{id}/audit | List a script&#39;s activity
[**GetExport**](ScriptsAPI.md#GetExport) | **Get** /v1/scripts/export | Export scripts
[**GetScratchpadBuilds**](ScriptsAPI.md#GetScratchpadBuilds) | **Get** /v1/scripts/scratchpad/builds/{buildId} | Get a scratchpad build
[**GetScratchpadRuns**](ScriptsAPI.md#GetScratchpadRuns) | **Get** /v1/scripts/scratchpad/runs/{runId} | Get a scratchpad run
[**GetUsage**](ScriptsAPI.md#GetUsage) | **Get** /v1/scripts/usage | Get script usage
[**List**](ScriptsAPI.md#List) | **Get** /v1/scripts | List scripts
[**ListRuns**](ScriptsAPI.md#ListRuns) | **Get** /v1/scripts/{id}/runs | List a script&#39;s runs
[**ListRunsLogs**](ScriptsAPI.md#ListRunsLogs) | **Get** /v1/scripts/{id}/runs/{runId}/logs | Get a script run&#39;s logs
[**Update**](ScriptsAPI.md#Update) | **Patch** /v1/scripts/{id} | Update a script
[**UpdateSecrets**](ScriptsAPI.md#UpdateSecrets) | **Put** /v1/scripts/{id}/secrets | Set script secrets



## Create

> ScriptDetail Create(ctx, scriptPostReqBody)

Create a script



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
	scriptPostReqBody :=  // ScriptPostReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Scripts.Create(context.Background(), scriptPostReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.Create`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scriptPostReqBody** | [**ScriptPostReqBody**](ScriptPostReqBody.md) |  | 

### Return type

[**ScriptDetail**](ScriptDetail.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReauthorise

> ScriptDetail CreateReauthorise(ctx, id)

Re-authorise a script



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
	resp, err := client.Scripts.CreateReauthorise(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.CreateReauthorise`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.CreateReauthorise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReauthoriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScriptDetail**](ScriptDetail.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRegenerateWebhook

> ScriptWebhookResBody CreateRegenerateWebhook(ctx, id)

Regenerate a script's webhook URL



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
	resp, err := client.Scripts.CreateRegenerateWebhook(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.CreateRegenerateWebhook`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.CreateRegenerateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegenerateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScriptWebhookResBody**](ScriptWebhookResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSchedulePreview

> ScriptSchedulePreviewResBody CreateSchedulePreview(ctx, scriptSchedulePreviewReqBody)

Preview a script schedule



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
	scriptSchedulePreviewReqBody :=  // ScriptSchedulePreviewReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Scripts.CreateSchedulePreview(context.Background(), scriptSchedulePreviewReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.CreateSchedulePreview`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.CreateSchedulePreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSchedulePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scriptSchedulePreviewReqBody** | [**ScriptSchedulePreviewReqBody**](ScriptSchedulePreviewReqBody.md) |  | 

### Return type

[**ScriptSchedulePreviewResBody**](ScriptSchedulePreviewResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScratchpadBuilds

> ScratchpadBuild CreateScratchpadBuilds(ctx, scratchpadBuildPostReqBody)

Build scratchpad code with its dependencies



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
	scratchpadBuildPostReqBody :=  // ScratchpadBuildPostReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Scripts.CreateScratchpadBuilds(context.Background(), scratchpadBuildPostReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.CreateScratchpadBuilds`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.CreateScratchpadBuilds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScratchpadBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scratchpadBuildPostReqBody** | [**ScratchpadBuildPostReqBody**](ScratchpadBuildPostReqBody.md) |  | 

### Return type

[**ScratchpadBuild**](ScratchpadBuild.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScratchpadRuns

> ScratchpadRun CreateScratchpadRuns(ctx, scratchpadRunPostReqBody)

Run scratchpad code



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
	scratchpadRunPostReqBody :=  // ScratchpadRunPostReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Scripts.CreateScratchpadRuns(context.Background(), scratchpadRunPostReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.CreateScratchpadRuns`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.CreateScratchpadRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScratchpadRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scratchpadRunPostReqBody** | [**ScratchpadRunPostReqBody**](ScratchpadRunPostReqBody.md) |  | 

### Return type

[**ScratchpadRun**](ScratchpadRun.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, id)

Delete a script



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
	err := client.Scripts.Delete(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.Delete`: %v\n", err)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


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


## Get

> ScriptDetail Get(ctx, id)

Get a script



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
	resp, err := client.Scripts.Get(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.Get`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScriptDetail**](ScriptDetail.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudit

> ScriptAuditResBody GetAudit(ctx, id, params)

List a script's activity



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
	resp, err := client.Scripts.GetAudit(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.GetAudit`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.GetAudit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 25]
 **cursor** | **string** | nextCursor from the previous page. | 

### Return type

[**ScriptAuditResBody**](ScriptAuditResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExport

> ScriptExportResBody GetExport(ctx)

Export scripts



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
	resp, err := client.Scripts.GetExport(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.GetExport`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.GetExport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportRequest struct via the builder pattern


### Return type

[**ScriptExportResBody**](ScriptExportResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScratchpadBuilds

> ScratchpadBuild GetScratchpadBuilds(ctx, buildId, params)

Get a scratchpad build



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
	buildId :=  // string | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Scripts.GetScratchpadBuilds(context.Background(), buildId, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.GetScratchpadBuilds`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.GetScratchpadBuilds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScratchpadBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **waitSeconds** | **int32** | Longest to wait for a build in progress to finish before responding. Defaults to 0, returning at once. | 

### Return type

[**ScratchpadBuild**](ScratchpadBuild.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScratchpadRuns

> ScratchpadRun GetScratchpadRuns(ctx, runId, params)

Get a scratchpad run



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
	runId :=  // string | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Scripts.GetScratchpadRuns(context.Background(), runId, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.GetScratchpadRuns`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.GetScratchpadRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScratchpadRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **waitSeconds** | **int32** | Longest to wait for a run in progress to finish before responding. Defaults to 0, returning at once. | 

### Return type

[**ScratchpadRun**](ScratchpadRun.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsage

> ScriptUsageResBody GetUsage(ctx)

Get script usage



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
	resp, err := client.Scripts.GetUsage(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.GetUsage`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.GetUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageRequest struct via the builder pattern


### Return type

[**ScriptUsageResBody**](ScriptUsageResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ScriptListResBody List(ctx)

List scripts



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
	resp, err := client.Scripts.List(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.List`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.List`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


### Return type

[**ScriptListResBody**](ScriptListResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRuns

> ScriptRunsResBody ListRuns(ctx, id, params)

List a script's runs



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
	resp, err := client.Scripts.ListRuns(context.Background(), id, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.ListRuns`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.ListRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 25]
 **cursor** | **string** | nextCursor from the previous page. | 

### Return type

[**ScriptRunsResBody**](ScriptRunsResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRunsLogs

> ScriptRunLogsResBody ListRunsLogs(ctx, id, runId)

Get a script run's logs



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
	runId :=  // string | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Scripts.ListRunsLogs(context.Background(), id, runId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.ListRunsLogs`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.ListRunsLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRunsLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ScriptRunLogsResBody**](ScriptRunLogsResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> ScriptDetail Update(ctx, id, scriptPatchReqBody)

Update a script



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
	scriptPatchReqBody :=  // ScriptPatchReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Scripts.Update(context.Background(), id, scriptPatchReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.Update`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scriptPatchReqBody** | [**ScriptPatchReqBody**](ScriptPatchReqBody.md) |  | 

### Return type

[**ScriptDetail**](ScriptDetail.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecrets

> ScriptSecretsResBody UpdateSecrets(ctx, id, scriptSecretsReqBody)

Set script secrets



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
	scriptSecretsReqBody :=  // ScriptSecretsReqBody | 

	client := openapiclient.NewAPIKeyClient(os.Getenv("SPACIV_API_KEY"))
	resp, err := client.Scripts.UpdateSecrets(context.Background(), id, scriptSecretsReqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Scripts.UpdateSecrets`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `Scripts.UpdateSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scriptSecretsReqBody** | [**ScriptSecretsReqBody**](ScriptSecretsReqBody.md) |  | 

### Return type

[**ScriptSecretsResBody**](ScriptSecretsResBody.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

