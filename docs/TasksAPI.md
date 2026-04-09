# \TasksAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTasks**](TasksAPI.md#GetTasks) | **Get** /v1/tasks | GetListOfTasks (by workflow id)
[**GetTasksApprovalByMe**](TasksAPI.md#GetTasksApprovalByMe) | **Get** /v1/tasks/approval_by_me | List approval tasks assigned to API user (non-registered status)
[**GetTasksAssignedByMe**](TasksAPI.md#GetTasksAssignedByMe) | **Get** /v1/tasks/assigned_by_me | GetListOfTasks (managed_by_me) Copy 3
[**GetTasksAssignedToMe**](TasksAPI.md#GetTasksAssignedToMe) | **Get** /v1/tasks/assigned_to_me | GetListOfTasks (managed_by_me) Copy 2
[**GetTasksAssignedToMyTeam**](TasksAPI.md#GetTasksAssignedToMyTeam) | **Get** /v1/tasks/assigned_to_my_team | GetListOfTasks (managed_by_me) Copy
[**GetTasksFinished**](TasksAPI.md#GetTasksFinished) | **Get** /v1/tasks/finished | List finished tasks
[**GetTasksId**](TasksAPI.md#GetTasksId) | **Get** /v1/tasks/{id} | GetTasksProperties
[**GetTasksManagedByMe**](TasksAPI.md#GetTasksManagedByMe) | **Get** /v1/tasks/managed_by_me | GetListOfTasks (managed_by_me)
[**GetTasksOpen**](TasksAPI.md#GetTasksOpen) | **Get** /v1/tasks/open | GetListOfTasks (opened)
[**PatchTasksId**](TasksAPI.md#PatchTasksId) | **Patch** /v1/tasks/{id} | SetNewStatus Copy
[**PostTasksIdNotes**](TasksAPI.md#PostTasksIdNotes) | **Post** /v1/tasks/{id}/notes | AddCommentToTask



## GetTasks

> GetTasks(ctx).Authorization(authorization).X4meAccount(x4meAccount).Workflow(workflow).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()

GetListOfTasks (by workflow id)

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
	workflow := "{{some workflow id}}" // string |  (optional)
	xXurrentLanguage := "xXurrentLanguage_example" // string | Override response language for enums/errors (e.g. nl, fr). See API introduction. (optional)
	perPage := int32(56) // int32 | Page size (max 100). See Pagination docs. (optional)
	searchAfter := "searchAfter_example" // string | Cursor for next page (from Link rel=next). (optional)
	searchBefore := "searchBefore_example" // string | Cursor for previous page (from Link rel=prev). (optional)
	fields := "fields_example" // string | Comma-separated fields for collections (field selection). (optional)
	sort := "sort_example" // string | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. (optional)
	state := "state_example" // string | Predefined filter name (alternative to path /state). See Filtering docs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TasksAPI.GetTasks(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Workflow(workflow).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.GetTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **workflow** | **string** |  | 
 **xXurrentLanguage** | **string** | Override response language for enums/errors (e.g. nl, fr). See API introduction. | 
 **perPage** | **int32** | Page size (max 100). See Pagination docs. | 
 **searchAfter** | **string** | Cursor for next page (from Link rel&#x3D;next). | 
 **searchBefore** | **string** | Cursor for previous page (from Link rel&#x3D;prev). | 
 **fields** | **string** | Comma-separated fields for collections (field selection). | 
 **sort** | **string** | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. | 
 **state** | **string** | Predefined filter name (alternative to path /state). See Filtering docs. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksApprovalByMe

> []map[string]interface{} GetTasksApprovalByMe(ctx).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()

List approval tasks assigned to API user (non-registered status)

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
	xXurrentLanguage := "xXurrentLanguage_example" // string | Override response language for enums/errors (e.g. nl, fr). See API introduction. (optional)
	perPage := int32(56) // int32 | Page size (max 100). See Pagination docs. (optional)
	searchAfter := "searchAfter_example" // string | Cursor for next page (from Link rel=next). (optional)
	searchBefore := "searchBefore_example" // string | Cursor for previous page (from Link rel=prev). (optional)
	fields := "fields_example" // string | Comma-separated fields for collections (field selection). (optional)
	sort := "sort_example" // string | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. (optional)
	state := "state_example" // string | Predefined filter name (alternative to path /state). See Filtering docs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.GetTasksApprovalByMe(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.GetTasksApprovalByMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTasksApprovalByMe`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.GetTasksApprovalByMe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksApprovalByMeRequest struct via the builder pattern


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

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksAssignedByMe

> GetTasksAssignedByMe(ctx).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()

GetListOfTasks (managed_by_me) Copy 3

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
	xXurrentLanguage := "xXurrentLanguage_example" // string | Override response language for enums/errors (e.g. nl, fr). See API introduction. (optional)
	perPage := int32(56) // int32 | Page size (max 100). See Pagination docs. (optional)
	searchAfter := "searchAfter_example" // string | Cursor for next page (from Link rel=next). (optional)
	searchBefore := "searchBefore_example" // string | Cursor for previous page (from Link rel=prev). (optional)
	fields := "fields_example" // string | Comma-separated fields for collections (field selection). (optional)
	sort := "sort_example" // string | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. (optional)
	state := "state_example" // string | Predefined filter name (alternative to path /state). See Filtering docs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TasksAPI.GetTasksAssignedByMe(context.Background()).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.GetTasksAssignedByMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksAssignedByMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xXurrentLanguage** | **string** | Override response language for enums/errors (e.g. nl, fr). See API introduction. | 
 **perPage** | **int32** | Page size (max 100). See Pagination docs. | 
 **searchAfter** | **string** | Cursor for next page (from Link rel&#x3D;next). | 
 **searchBefore** | **string** | Cursor for previous page (from Link rel&#x3D;prev). | 
 **fields** | **string** | Comma-separated fields for collections (field selection). | 
 **sort** | **string** | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. | 
 **state** | **string** | Predefined filter name (alternative to path /state). See Filtering docs. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksAssignedToMe

> GetTasksAssignedToMe(ctx).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()

GetListOfTasks (managed_by_me) Copy 2

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
	xXurrentLanguage := "xXurrentLanguage_example" // string | Override response language for enums/errors (e.g. nl, fr). See API introduction. (optional)
	perPage := int32(56) // int32 | Page size (max 100). See Pagination docs. (optional)
	searchAfter := "searchAfter_example" // string | Cursor for next page (from Link rel=next). (optional)
	searchBefore := "searchBefore_example" // string | Cursor for previous page (from Link rel=prev). (optional)
	fields := "fields_example" // string | Comma-separated fields for collections (field selection). (optional)
	sort := "sort_example" // string | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. (optional)
	state := "state_example" // string | Predefined filter name (alternative to path /state). See Filtering docs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TasksAPI.GetTasksAssignedToMe(context.Background()).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.GetTasksAssignedToMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksAssignedToMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xXurrentLanguage** | **string** | Override response language for enums/errors (e.g. nl, fr). See API introduction. | 
 **perPage** | **int32** | Page size (max 100). See Pagination docs. | 
 **searchAfter** | **string** | Cursor for next page (from Link rel&#x3D;next). | 
 **searchBefore** | **string** | Cursor for previous page (from Link rel&#x3D;prev). | 
 **fields** | **string** | Comma-separated fields for collections (field selection). | 
 **sort** | **string** | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. | 
 **state** | **string** | Predefined filter name (alternative to path /state). See Filtering docs. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksAssignedToMyTeam

> GetTasksAssignedToMyTeam(ctx).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()

GetListOfTasks (managed_by_me) Copy

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
	xXurrentLanguage := "xXurrentLanguage_example" // string | Override response language for enums/errors (e.g. nl, fr). See API introduction. (optional)
	perPage := int32(56) // int32 | Page size (max 100). See Pagination docs. (optional)
	searchAfter := "searchAfter_example" // string | Cursor for next page (from Link rel=next). (optional)
	searchBefore := "searchBefore_example" // string | Cursor for previous page (from Link rel=prev). (optional)
	fields := "fields_example" // string | Comma-separated fields for collections (field selection). (optional)
	sort := "sort_example" // string | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. (optional)
	state := "state_example" // string | Predefined filter name (alternative to path /state). See Filtering docs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TasksAPI.GetTasksAssignedToMyTeam(context.Background()).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.GetTasksAssignedToMyTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksAssignedToMyTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xXurrentLanguage** | **string** | Override response language for enums/errors (e.g. nl, fr). See API introduction. | 
 **perPage** | **int32** | Page size (max 100). See Pagination docs. | 
 **searchAfter** | **string** | Cursor for next page (from Link rel&#x3D;next). | 
 **searchBefore** | **string** | Cursor for previous page (from Link rel&#x3D;prev). | 
 **fields** | **string** | Comma-separated fields for collections (field selection). | 
 **sort** | **string** | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. | 
 **state** | **string** | Predefined filter name (alternative to path /state). See Filtering docs. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksFinished

> []map[string]interface{} GetTasksFinished(ctx).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()

List finished tasks

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
	xXurrentLanguage := "xXurrentLanguage_example" // string | Override response language for enums/errors (e.g. nl, fr). See API introduction. (optional)
	perPage := int32(56) // int32 | Page size (max 100). See Pagination docs. (optional)
	searchAfter := "searchAfter_example" // string | Cursor for next page (from Link rel=next). (optional)
	searchBefore := "searchBefore_example" // string | Cursor for previous page (from Link rel=prev). (optional)
	fields := "fields_example" // string | Comma-separated fields for collections (field selection). (optional)
	sort := "sort_example" // string | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. (optional)
	state := "state_example" // string | Predefined filter name (alternative to path /state). See Filtering docs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.GetTasksFinished(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.GetTasksFinished``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTasksFinished`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.GetTasksFinished`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksFinishedRequest struct via the builder pattern


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

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksId

> GetTasksId(ctx, id).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).IfNoneMatch(ifNoneMatch).Execute()

GetTasksProperties

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
	r, err := apiClient.TasksAPI.GetTasksId(context.Background(), id).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.GetTasksId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetTasksIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **xXurrentLanguage** | **string** | Override response language for enums/errors (e.g. nl, fr). See API introduction. | 
 **ifNoneMatch** | **string** | Conditional GET; returns 304 Not Modified when unchanged. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksManagedByMe

> GetTasksManagedByMe(ctx).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()

GetListOfTasks (managed_by_me)

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
	xXurrentLanguage := "xXurrentLanguage_example" // string | Override response language for enums/errors (e.g. nl, fr). See API introduction. (optional)
	perPage := int32(56) // int32 | Page size (max 100). See Pagination docs. (optional)
	searchAfter := "searchAfter_example" // string | Cursor for next page (from Link rel=next). (optional)
	searchBefore := "searchBefore_example" // string | Cursor for previous page (from Link rel=prev). (optional)
	fields := "fields_example" // string | Comma-separated fields for collections (field selection). (optional)
	sort := "sort_example" // string | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. (optional)
	state := "state_example" // string | Predefined filter name (alternative to path /state). See Filtering docs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TasksAPI.GetTasksManagedByMe(context.Background()).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.GetTasksManagedByMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksManagedByMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xXurrentLanguage** | **string** | Override response language for enums/errors (e.g. nl, fr). See API introduction. | 
 **perPage** | **int32** | Page size (max 100). See Pagination docs. | 
 **searchAfter** | **string** | Cursor for next page (from Link rel&#x3D;next). | 
 **searchBefore** | **string** | Cursor for previous page (from Link rel&#x3D;prev). | 
 **fields** | **string** | Comma-separated fields for collections (field selection). | 
 **sort** | **string** | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. | 
 **state** | **string** | Predefined filter name (alternative to path /state). See Filtering docs. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksOpen

> GetTasksOpen(ctx).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()

GetListOfTasks (opened)

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
	xXurrentLanguage := "xXurrentLanguage_example" // string | Override response language for enums/errors (e.g. nl, fr). See API introduction. (optional)
	perPage := int32(56) // int32 | Page size (max 100). See Pagination docs. (optional)
	searchAfter := "searchAfter_example" // string | Cursor for next page (from Link rel=next). (optional)
	searchBefore := "searchBefore_example" // string | Cursor for previous page (from Link rel=prev). (optional)
	fields := "fields_example" // string | Comma-separated fields for collections (field selection). (optional)
	sort := "sort_example" // string | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. (optional)
	state := "state_example" // string | Predefined filter name (alternative to path /state). See Filtering docs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TasksAPI.GetTasksOpen(context.Background()).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.GetTasksOpen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xXurrentLanguage** | **string** | Override response language for enums/errors (e.g. nl, fr). See API introduction. | 
 **perPage** | **int32** | Page size (max 100). See Pagination docs. | 
 **searchAfter** | **string** | Cursor for next page (from Link rel&#x3D;next). | 
 **searchBefore** | **string** | Cursor for previous page (from Link rel&#x3D;prev). | 
 **fields** | **string** | Comma-separated fields for collections (field selection). | 
 **sort** | **string** | Sort fields, comma-separated; prefix with - for descending. See Ordering docs. | 
 **state** | **string** | Predefined filter name (alternative to path /state). See Filtering docs. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTasksId

> string PatchTasksId(ctx, id).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()

SetNewStatus Copy

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.PatchTasksId(context.Background(), id).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.PatchTasksId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTasksId`: string
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.PatchTasksId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTasksIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTasksIdNotes

> map[string]interface{} PostTasksIdNotes(ctx, id).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()

AddCommentToTask

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.PostTasksIdNotes(context.Background(), id).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.PostTasksIdNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTasksIdNotes`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.PostTasksIdNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTasksIdNotesRequest struct via the builder pattern


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

