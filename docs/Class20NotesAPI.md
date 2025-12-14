# \Class20NotesAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1Requests1NotesGet**](Class20NotesAPI.md#V1Requests1NotesGet) | **Get** /v1/requests/1/notes | GetIssuesNotesList
[**V1Requests1NotesGet_0**](Class20NotesAPI.md#V1Requests1NotesGet_0) | **Get** /v1/requests/1/notes/ | GetIssueNotesList (query - created at)



## V1Requests1NotesGet

> map[string]interface{} V1Requests1NotesGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetIssuesNotesList

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
	resp, r, err := apiClient.Class20NotesAPI.V1Requests1NotesGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class20NotesAPI.V1Requests1NotesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1Requests1NotesGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class20NotesAPI.V1Requests1NotesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Requests1NotesGetRequest struct via the builder pattern


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


## V1Requests1NotesGet_0

> map[string]interface{} V1Requests1NotesGet_0(ctx).Authorization(authorization).X4meAccount(x4meAccount).CreatedAt(createdAt).Execute()

GetIssueNotesList (query - created at)

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
	createdAt := ">={{some time}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class20NotesAPI.V1Requests1NotesGet_0(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).CreatedAt(createdAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class20NotesAPI.V1Requests1NotesGet_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1Requests1NotesGet_0`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class20NotesAPI.V1Requests1NotesGet_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Requests1NotesGet_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **createdAt** | **string** |  | 

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

