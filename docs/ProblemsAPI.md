# \ProblemsAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ProblemsGet**](ProblemsAPI.md#V1ProblemsGet) | **Get** /v1/problems | GetProblemProperties
[**V1ProblemsIdPatch**](ProblemsAPI.md#V1ProblemsIdPatch) | **Patch** /v1/problems/{id} | SetNewStatus (solved)
[**V1ProblemsIdRequestsIdPost**](ProblemsAPI.md#V1ProblemsIdRequestsIdPost) | **Post** /v1/problems/{id}/requests/{id} | UpdateProblem - AddRequestToProblem
[**V1ProblemsPost**](ProblemsAPI.md#V1ProblemsPost) | **Post** /v1/problems | CreateNewProblem



## V1ProblemsGet

> V1ProblemsGet(ctx).X4meAccount(x4meAccount).Execute()

GetProblemProperties

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xurrent/go-xurrent/xurrent"
)

func main() {
	x4meAccount := "school" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProblemsAPI.V1ProblemsGet(context.Background()).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProblemsAPI.V1ProblemsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ProblemsGetRequest struct via the builder pattern


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


## V1ProblemsIdPatch

> V1ProblemsIdPatch(ctx, id).X4meAccount(x4meAccount).Body(body).Execute()

SetNewStatus (solved)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xurrent/go-xurrent/xurrent"
)

func main() {
	id := int32(56) // int32 | 
	x4meAccount := "school" // string |  (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProblemsAPI.V1ProblemsIdPatch(context.Background(), id).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProblemsAPI.V1ProblemsIdPatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1ProblemsIdPatchRequest struct via the builder pattern


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


## V1ProblemsIdRequestsIdPost

> V1ProblemsIdRequestsIdPost(ctx, id).X4meAccount(x4meAccount).Body(body).Execute()

UpdateProblem - AddRequestToProblem

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xurrent/go-xurrent/xurrent"
)

func main() {
	id := int32(56) // int32 | 
	x4meAccount := "school" // string |  (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProblemsAPI.V1ProblemsIdRequestsIdPost(context.Background(), id).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProblemsAPI.V1ProblemsIdRequestsIdPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1ProblemsIdRequestsIdPostRequest struct via the builder pattern


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


## V1ProblemsPost

> V1ProblemsPost(ctx).X4meAccount(x4meAccount).Body(body).Execute()

CreateNewProblem

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xurrent/go-xurrent/xurrent"
)

func main() {
	x4meAccount := "school" // string |  (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProblemsAPI.V1ProblemsPost(context.Background()).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProblemsAPI.V1ProblemsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ProblemsPostRequest struct via the builder pattern


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

