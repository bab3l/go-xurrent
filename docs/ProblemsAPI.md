# \ProblemsAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProblems**](ProblemsAPI.md#GetProblems) | **Get** /v1/problems | GetProblemProperties
[**PatchProblemsId**](ProblemsAPI.md#PatchProblemsId) | **Patch** /v1/problems/{id} | SetNewStatus (solved)
[**PostProblems**](ProblemsAPI.md#PostProblems) | **Post** /v1/problems | CreateNewProblem
[**PostProblemsProblemIdRequestsRequestId**](ProblemsAPI.md#PostProblemsProblemIdRequestsRequestId) | **Post** /v1/problems/{problem_id}/requests/{request_id} | UpdateProblem - AddRequestToProblem



## GetProblems

> GetProblems(ctx).X4meAccount(x4meAccount).Execute()

GetProblemProperties

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bab3l/go-xurrent"
)

func main() {
	x4meAccount := "school" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProblemsAPI.GetProblems(context.Background()).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProblemsAPI.GetProblems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProblemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x4meAccount** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProblemsId

> PatchProblemsId(ctx, id).X4meAccount(x4meAccount).Body(body).Execute()

SetNewStatus (solved)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bab3l/go-xurrent"
)

func main() {
	id := int32(56) // int32 | 
	x4meAccount := "school" // string |  (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProblemsAPI.PatchProblemsId(context.Background(), id).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProblemsAPI.PatchProblemsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProblemsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **x4meAccount** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProblems

> PostProblems(ctx).X4meAccount(x4meAccount).Body(body).Execute()

CreateNewProblem

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bab3l/go-xurrent"
)

func main() {
	x4meAccount := "school" // string |  (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProblemsAPI.PostProblems(context.Background()).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProblemsAPI.PostProblems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProblemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x4meAccount** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProblemsProblemIdRequestsRequestId

> PostProblemsProblemIdRequestsRequestId(ctx, problemId, requestId).X4meAccount(x4meAccount).Body(body).Execute()

UpdateProblem - AddRequestToProblem

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bab3l/go-xurrent"
)

func main() {
	problemId := int32(56) // int32 | 
	requestId := int32(56) // int32 | 
	x4meAccount := "school" // string |  (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProblemsAPI.PostProblemsProblemIdRequestsRequestId(context.Background(), problemId, requestId).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProblemsAPI.PostProblemsProblemIdRequestsRequestId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **int32** |  | 
**requestId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProblemsProblemIdRequestsRequestIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **x4meAccount** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

