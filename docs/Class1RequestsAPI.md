# \Class1RequestsAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1Requests1Cis1Delete**](Class1RequestsAPI.md#V1Requests1Cis1Delete) | **Delete** /v1/requests/1/cis/1 | RemoveLinkWithCi
[**V1Requests1Cis1Post**](Class1RequestsAPI.md#V1Requests1Cis1Post) | **Post** /v1/requests/1/cis/1 | SetLinkWithOneCi
[**V1Requests1Cis460359Post**](Class1RequestsAPI.md#V1Requests1Cis460359Post) | **Post** /v1/requests/1/cis/460359 | SetCIToRequest
[**V1Requests1CisGet**](Class1RequestsAPI.md#V1Requests1CisGet) | **Get** /v1/requests/1/cis | GetIssueAllRelatedCiList
[**V1Requests1DissatisfiedPost**](Class1RequestsAPI.md#V1Requests1DissatisfiedPost) | **Post** /v1/requests/1/dissatisfied | ReturnToWork - Decline Execution
[**V1Requests1Get**](Class1RequestsAPI.md#V1Requests1Get) | **Get** /v1/requests/1 | GetIssuesProperties
[**V1Requests1GroupedRequestsGet**](Class1RequestsAPI.md#V1Requests1GroupedRequestsGet) | **Get** /v1/requests/1/grouped_requests | GetGroupedRequestLinkedRecordsList
[**V1Requests1NotesPost**](Class1RequestsAPI.md#V1Requests1NotesPost) | **Post** /v1/requests/1/notes | AddInternalComment
[**V1Requests1Put**](Class1RequestsAPI.md#V1Requests1Put) | **Put** /v1/requests/1 | SetNewStatus (no_reply)
[**V1Requests1SatisfiedPost**](Class1RequestsAPI.md#V1Requests1SatisfiedPost) | **Post** /v1/requests/1/satisfied | CloseRequest - Satisfied
[**V1Requests2048741AuditGet**](Class1RequestsAPI.md#V1Requests2048741AuditGet) | **Get** /v1/requests/2048741/audit | GetIssuesAuditEntries
[**V1RequestsAssignedToMeGet**](Class1RequestsAPI.md#V1RequestsAssignedToMeGet) | **Get** /v1/requests/assigned_to_me | GetIssuesList (assigned to ...)
[**V1RequestsAssignedToMyTeamGet**](Class1RequestsAPI.md#V1RequestsAssignedToMyTeamGet) | **Get** /v1/requests/assigned_to_my_team | GetIssuesList (assigned to my team)
[**V1RequestsCompletedGet**](Class1RequestsAPI.md#V1RequestsCompletedGet) | **Get** /v1/requests/completed | GetIssuesList (in status \&quot;completed\&quot;)
[**V1RequestsGet**](Class1RequestsAPI.md#V1RequestsGet) | **Get** /v1/requests/ | GetIssuesList (by subject) Copy
[**V1RequestsOpenGet**](Class1RequestsAPI.md#V1RequestsOpenGet) | **Get** /v1/requests/open | GetIssuesList (in status \&quot;open\&quot;)
[**V1RequestsRequestedByOrForMeGet**](Class1RequestsAPI.md#V1RequestsRequestedByOrForMeGet) | **Get** /v1/requests/requested_by_or_for_me | GetIssuesList (requested by or for current user)
[**V1RequestsRequestsOfMyOrganizationGet**](Class1RequestsAPI.md#V1RequestsRequestsOfMyOrganizationGet) | **Get** /v1/requests/requests_of_my_organization | GetIssuesList (requests_of_my_organization)



## V1Requests1Cis1Delete

> map[string]interface{} V1Requests1Cis1Delete(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

RemoveLinkWithCi

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1RequestsAPI.V1Requests1Cis1Delete(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1Requests1Cis1Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1Requests1Cis1Delete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class1RequestsAPI.V1Requests1Cis1Delete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Requests1Cis1DeleteRequest struct via the builder pattern


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


## V1Requests1Cis1Post

> map[string]interface{} V1Requests1Cis1Post(ctx).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()

SetLinkWithOneCi

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1RequestsAPI.V1Requests1Cis1Post(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1Requests1Cis1Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1Requests1Cis1Post`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class1RequestsAPI.V1Requests1Cis1Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Requests1Cis1PostRequest struct via the builder pattern


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


## V1Requests1Cis460359Post

> V1Requests1Cis460359Post(ctx).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()

SetCIToRequest

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class1RequestsAPI.V1Requests1Cis460359Post(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1Requests1Cis460359Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Requests1Cis460359PostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1Requests1CisGet

> map[string]interface{} V1Requests1CisGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetIssueAllRelatedCiList

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1RequestsAPI.V1Requests1CisGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1Requests1CisGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1Requests1CisGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class1RequestsAPI.V1Requests1CisGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Requests1CisGetRequest struct via the builder pattern


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


## V1Requests1DissatisfiedPost

> string V1Requests1DissatisfiedPost(ctx).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()

ReturnToWork - Decline Execution

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1RequestsAPI.V1Requests1DissatisfiedPost(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1Requests1DissatisfiedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1Requests1DissatisfiedPost`: string
	fmt.Fprintf(os.Stdout, "Response from `Class1RequestsAPI.V1Requests1DissatisfiedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Requests1DissatisfiedPostRequest struct via the builder pattern


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


## V1Requests1Get

> map[string]interface{} V1Requests1Get(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetIssuesProperties

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1RequestsAPI.V1Requests1Get(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1Requests1Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1Requests1Get`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class1RequestsAPI.V1Requests1Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Requests1GetRequest struct via the builder pattern


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


## V1Requests1GroupedRequestsGet

> map[string]interface{} V1Requests1GroupedRequestsGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetGroupedRequestLinkedRecordsList

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1RequestsAPI.V1Requests1GroupedRequestsGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1Requests1GroupedRequestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1Requests1GroupedRequestsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class1RequestsAPI.V1Requests1GroupedRequestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Requests1GroupedRequestsGetRequest struct via the builder pattern


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


## V1Requests1NotesPost

> map[string]interface{} V1Requests1NotesPost(ctx).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()

AddInternalComment

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1RequestsAPI.V1Requests1NotesPost(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1Requests1NotesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1Requests1NotesPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class1RequestsAPI.V1Requests1NotesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Requests1NotesPostRequest struct via the builder pattern


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


## V1Requests1Put

> V1Requests1Put(ctx).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()

SetNewStatus (no_reply)

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class1RequestsAPI.V1Requests1Put(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1Requests1Put``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Requests1PutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1Requests1SatisfiedPost

> V1Requests1SatisfiedPost(ctx).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()

CloseRequest - Satisfied

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class1RequestsAPI.V1Requests1SatisfiedPost(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1Requests1SatisfiedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Requests1SatisfiedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1Requests2048741AuditGet

> V1Requests2048741AuditGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetIssuesAuditEntries

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class1RequestsAPI.V1Requests2048741AuditGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1Requests2048741AuditGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Requests2048741AuditGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 

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


## V1RequestsAssignedToMeGet

> map[string]interface{} V1RequestsAssignedToMeGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Subject(subject).Status(status).Execute()

GetIssuesList (assigned to ...)

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)
	subject := "{{subject}}" // string |  (optional)
	status := "{{status}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1RequestsAPI.V1RequestsAssignedToMeGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Subject(subject).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1RequestsAssignedToMeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1RequestsAssignedToMeGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class1RequestsAPI.V1RequestsAssignedToMeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1RequestsAssignedToMeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **subject** | **string** |  | 
 **status** | **string** |  | 

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


## V1RequestsAssignedToMyTeamGet

> map[string]interface{} V1RequestsAssignedToMyTeamGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetIssuesList (assigned to my team)

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1RequestsAPI.V1RequestsAssignedToMyTeamGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1RequestsAssignedToMyTeamGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1RequestsAssignedToMyTeamGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class1RequestsAPI.V1RequestsAssignedToMyTeamGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1RequestsAssignedToMyTeamGetRequest struct via the builder pattern


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


## V1RequestsCompletedGet

> map[string]interface{} V1RequestsCompletedGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetIssuesList (in status \"completed\")

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1RequestsAPI.V1RequestsCompletedGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1RequestsCompletedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1RequestsCompletedGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class1RequestsAPI.V1RequestsCompletedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1RequestsCompletedGetRequest struct via the builder pattern


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


## V1RequestsGet

> V1RequestsGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Subject(subject).Execute()

GetIssuesList (by subject) Copy

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)
	subject := "{{subject}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class1RequestsAPI.V1RequestsGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Subject(subject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1RequestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1RequestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **subject** | **string** |  | 

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


## V1RequestsOpenGet

> map[string]interface{} V1RequestsOpenGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetIssuesList (in status \"open\")

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1RequestsAPI.V1RequestsOpenGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1RequestsOpenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1RequestsOpenGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class1RequestsAPI.V1RequestsOpenGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1RequestsOpenGetRequest struct via the builder pattern


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


## V1RequestsRequestedByOrForMeGet

> map[string]interface{} V1RequestsRequestedByOrForMeGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetIssuesList (requested by or for current user)

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1RequestsAPI.V1RequestsRequestedByOrForMeGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1RequestsRequestedByOrForMeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1RequestsRequestedByOrForMeGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class1RequestsAPI.V1RequestsRequestedByOrForMeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1RequestsRequestedByOrForMeGetRequest struct via the builder pattern


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


## V1RequestsRequestsOfMyOrganizationGet

> map[string]interface{} V1RequestsRequestsOfMyOrganizationGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetIssuesList (requests_of_my_organization)

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class1RequestsAPI.V1RequestsRequestsOfMyOrganizationGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class1RequestsAPI.V1RequestsRequestsOfMyOrganizationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1RequestsRequestsOfMyOrganizationGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class1RequestsAPI.V1RequestsRequestsOfMyOrganizationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1RequestsRequestsOfMyOrganizationGetRequest struct via the builder pattern


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

