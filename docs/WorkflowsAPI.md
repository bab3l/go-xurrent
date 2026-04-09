# \WorkflowsAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorkflows**](WorkflowsAPI.md#GetWorkflows) | **Get** /v1/workflows | GetWorkflows
[**GetWorkflowsId**](WorkflowsAPI.md#GetWorkflowsId) | **Get** /v1/workflows/{id} | GetWorkflowProperties
[**PatchWorkflowsId**](WorkflowsAPI.md#PatchWorkflowsId) | **Patch** /v1/workflows/{id} | Update a workflow
[**PostWorkflows**](WorkflowsAPI.md#PostWorkflows) | **Post** /v1/workflows | Create a workflow
[**PostWorkflowsIdArchive**](WorkflowsAPI.md#PostWorkflowsIdArchive) | **Post** /v1/workflows/{id}/archive | Archive a workflow
[**PostWorkflowsIdRestore**](WorkflowsAPI.md#PostWorkflowsIdRestore) | **Post** /v1/workflows/{id}/restore | Restore a workflow
[**PostWorkflowsIdTasks**](WorkflowsAPI.md#PostWorkflowsIdTasks) | **Post** /v1/workflows/{id}/tasks | Create a task on a workflow
[**PostWorkflowsIdTrash**](WorkflowsAPI.md#PostWorkflowsIdTrash) | **Post** /v1/workflows/{id}/trash | Trash a workflow



## GetWorkflows

> map[string]interface{} GetWorkflows(ctx).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()

GetWorkflows

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xurrent/go-xurrent"
)

func main() {
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)
	xXurrentLanguage := "xXurrentLanguage_example" // string | Override response language for enums/errors (e.g. nl, fr). See API introduction. (optional)
	perPage := int32(56) // int32 | Page size (max 100). See Pagination docs. (optional)
	searchAfter := "searchAfter_example" // string | Cursor for next page (from Link rel=next). (optional)
	searchBefore := "searchBefore_example" // string | Cursor for previous page (from Link rel=prev). (optional)
	fields := "fields_example" // string | Comma-separated fields for collections (field selection). (optional)
	sort := "sort_example" // string | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. (optional)
	state := "state_example" // string | Predefined filter name (alternative to path /state). See Filtering docs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflows(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflows`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **xXurrentLanguage** | **string** | Override response language for enums/errors (e.g. nl, fr). See API introduction. | 
 **perPage** | **int32** | Page size (max 100). See Pagination docs. | 
 **searchAfter** | **string** | Cursor for next page (from Link rel&#x3D;next). | 
 **searchBefore** | **string** | Cursor for previous page (from Link rel&#x3D;prev). | 
 **fields** | **string** | Comma-separated fields for collections (field selection). | 
 **sort** | **string** | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. | 
 **state** | **string** | Predefined filter name (alternative to path /state). See Filtering docs. | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowsId

> map[string]interface{} GetWorkflowsId(ctx, id).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).IfNoneMatch(ifNoneMatch).Execute()

GetWorkflowProperties

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xurrent/go-xurrent"
)

func main() {
	id := int32(56) // int32 | 
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)
	xXurrentLanguage := "xXurrentLanguage_example" // string | Override response language for enums/errors (e.g. nl, fr). See API introduction. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Conditional GET; returns 304 Not Modified when unchanged. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflowsId(context.Background(), id).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflowsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowsId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflowsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **xXurrentLanguage** | **string** | Override response language for enums/errors (e.g. nl, fr). See API introduction. | 
 **ifNoneMatch** | **string** | Conditional GET; returns 304 Not Modified when unchanged. | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWorkflowsId

> map[string]interface{} PatchWorkflowsId(ctx, id).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()

Update a workflow

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xurrent/go-xurrent"
)

func main() {
	id := int32(56) // int32 | 
	authorization := "authorization_example" // string |  (optional)
	x4meAccount := "x4meAccount_example" // string |  (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.PatchWorkflowsId(context.Background(), id).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.PatchWorkflowsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchWorkflowsId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.PatchWorkflowsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWorkflowsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWorkflows

> map[string]interface{} PostWorkflows(ctx).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()

Create a workflow

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xurrent/go-xurrent"
)

func main() {
	authorization := "authorization_example" // string |  (optional)
	x4meAccount := "x4meAccount_example" // string |  (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.PostWorkflows(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.PostWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWorkflows`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.PostWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWorkflowsIdArchive

> map[string]interface{} PostWorkflowsIdArchive(ctx, id).Authorization(authorization).X4meAccount(x4meAccount).Execute()

Archive a workflow

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xurrent/go-xurrent"
)

func main() {
	id := int32(56) // int32 | 
	authorization := "authorization_example" // string |  (optional)
	x4meAccount := "x4meAccount_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.PostWorkflowsIdArchive(context.Background(), id).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.PostWorkflowsIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWorkflowsIdArchive`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.PostWorkflowsIdArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWorkflowsIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWorkflowsIdRestore

> map[string]interface{} PostWorkflowsIdRestore(ctx, id).Authorization(authorization).X4meAccount(x4meAccount).Execute()

Restore a workflow

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xurrent/go-xurrent"
)

func main() {
	id := int32(56) // int32 | 
	authorization := "authorization_example" // string |  (optional)
	x4meAccount := "x4meAccount_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.PostWorkflowsIdRestore(context.Background(), id).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.PostWorkflowsIdRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWorkflowsIdRestore`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.PostWorkflowsIdRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWorkflowsIdRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWorkflowsIdTasks

> map[string]interface{} PostWorkflowsIdTasks(ctx, id).Body(body).Authorization(authorization).X4meAccount(x4meAccount).Execute()

Create a task on a workflow

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xurrent/go-xurrent"
)

func main() {
	id := int32(56) // int32 | Workflow id
	body := map[string]interface{}{ ... } // map[string]interface{} | 
	authorization := "authorization_example" // string |  (optional)
	x4meAccount := "x4meAccount_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.PostWorkflowsIdTasks(context.Background(), id).Body(body).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.PostWorkflowsIdTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWorkflowsIdTasks`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.PostWorkflowsIdTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Workflow id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWorkflowsIdTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWorkflowsIdTrash

> map[string]interface{} PostWorkflowsIdTrash(ctx, id).Authorization(authorization).X4meAccount(x4meAccount).Execute()

Trash a workflow

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xurrent/go-xurrent"
)

func main() {
	id := int32(56) // int32 | 
	authorization := "authorization_example" // string |  (optional)
	x4meAccount := "x4meAccount_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.PostWorkflowsIdTrash(context.Background(), id).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.PostWorkflowsIdTrash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWorkflowsIdTrash`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.PostWorkflowsIdTrash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWorkflowsIdTrashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

