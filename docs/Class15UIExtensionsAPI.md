# \Class15UIExtensionsAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1UiExtensions2405Get**](Class15UIExtensionsAPI.md#V1UiExtensions2405Get) | **Get** /v1/ui_extensions/2405 | GetUIExtensionProperties Copy
[**V1UiExtensionsGet**](Class15UIExtensionsAPI.md#V1UiExtensionsGet) | **Get** /v1/ui_extensions | GetUIExtensionPropertiesList



## V1UiExtensions2405Get

> map[string]interface{} V1UiExtensions2405Get(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetUIExtensionProperties Copy

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
	resp, r, err := apiClient.Class15UIExtensionsAPI.V1UiExtensions2405Get(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class15UIExtensionsAPI.V1UiExtensions2405Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UiExtensions2405Get`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class15UIExtensionsAPI.V1UiExtensions2405Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1UiExtensions2405GetRequest struct via the builder pattern


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


## V1UiExtensionsGet

> map[string]interface{} V1UiExtensionsGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetUIExtensionPropertiesList

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
	resp, r, err := apiClient.Class15UIExtensionsAPI.V1UiExtensionsGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class15UIExtensionsAPI.V1UiExtensionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UiExtensionsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class15UIExtensionsAPI.V1UiExtensionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1UiExtensionsGetRequest struct via the builder pattern


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

