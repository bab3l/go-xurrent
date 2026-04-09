# \RequestTemplatesAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRequestTemplates**](RequestTemplatesAPI.md#GetRequestTemplates) | **Get** /v1/request_templates | GetRequestTemplatesListByService
[**GetRequestTemplatesId**](RequestTemplatesAPI.md#GetRequestTemplatesId) | **Get** /v1/request_templates/{id} | GetRequestTemplatesProperties



## GetRequestTemplates

> map[string]interface{} GetRequestTemplates(ctx).Authorization(authorization).X4meAccount(x4meAccount).Service(service).Execute()

GetRequestTemplatesListByService

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)
	service := "{{service id}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestTemplatesAPI.GetRequestTemplates(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestTemplatesAPI.GetRequestTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequestTemplates`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RequestTemplatesAPI.GetRequestTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **service** | **string** |  | 

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


## GetRequestTemplatesId

> map[string]interface{} GetRequestTemplatesId(ctx, id).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetRequestTemplatesProperties

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
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestTemplatesAPI.GetRequestTemplatesId(context.Background(), id).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestTemplatesAPI.GetRequestTemplatesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequestTemplatesId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RequestTemplatesAPI.GetRequestTemplatesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestTemplatesIdRequest struct via the builder pattern


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

