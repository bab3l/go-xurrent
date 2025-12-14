# \Class18RequestTemplatesAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1RequestTemplates1Get**](Class18RequestTemplatesAPI.md#V1RequestTemplates1Get) | **Get** /v1/request_templates/1 | GetRequestTemplatesProperties
[**V1RequestTemplatesGet**](Class18RequestTemplatesAPI.md#V1RequestTemplatesGet) | **Get** /v1/request_templates | GetRequestTemplatesListByService



## V1RequestTemplates1Get

> map[string]interface{} V1RequestTemplates1Get(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetRequestTemplatesProperties

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
	resp, r, err := apiClient.Class18RequestTemplatesAPI.V1RequestTemplates1Get(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class18RequestTemplatesAPI.V1RequestTemplates1Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1RequestTemplates1Get`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class18RequestTemplatesAPI.V1RequestTemplates1Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1RequestTemplates1GetRequest struct via the builder pattern


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


## V1RequestTemplatesGet

> map[string]interface{} V1RequestTemplatesGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Service(service).Execute()

GetRequestTemplatesListByService

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
	service := "{{service id}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class18RequestTemplatesAPI.V1RequestTemplatesGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class18RequestTemplatesAPI.V1RequestTemplatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1RequestTemplatesGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class18RequestTemplatesAPI.V1RequestTemplatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1RequestTemplatesGetRequest struct via the builder pattern


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

