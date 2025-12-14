# \Class3AttachmentsAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AttachmentsPost**](Class3AttachmentsAPI.md#V1AttachmentsPost) | **Post** /v1/attachments | Add Attach - PutAttachToStorageAndGetURL
[**V1AttachmentsStorageGet**](Class3AttachmentsAPI.md#V1AttachmentsStorageGet) | **Get** /v1/attachments/storage | Add Attach - ReservePlaceForAttachment
[**V1AttachmentsUploadGet**](Class3AttachmentsAPI.md#V1AttachmentsUploadGet) | **Get** /v1/attachments/upload | Get Attach
[**V1Requests1Patch**](Class3AttachmentsAPI.md#V1Requests1Patch) | **Patch** /v1/requests/1 | Add Attach (multi) - AddAttachmentsToRequest



## V1AttachmentsPost

> map[string]interface{} V1AttachmentsPost(ctx).Authorization(authorization).X4meAccount(x4meAccount).Key(key).X4meExpiration(x4meExpiration).X4meSignature(x4meSignature).File(file).Execute()

Add Attach - PutAttachToStorageAndGetURL

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
	key := "key_example" // string |  (optional)
	x4meExpiration := "x4meExpiration_example" // string |  (optional)
	x4meSignature := "x4meSignature_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class3AttachmentsAPI.V1AttachmentsPost(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Key(key).X4meExpiration(x4meExpiration).X4meSignature(x4meSignature).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class3AttachmentsAPI.V1AttachmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AttachmentsPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class3AttachmentsAPI.V1AttachmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AttachmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **key** | **string** |  | 
 **x4meExpiration** | **string** |  | 
 **x4meSignature** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AttachmentsStorageGet

> map[string]interface{} V1AttachmentsStorageGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

Add Attach - ReservePlaceForAttachment

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
	resp, r, err := apiClient.Class3AttachmentsAPI.V1AttachmentsStorageGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class3AttachmentsAPI.V1AttachmentsStorageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AttachmentsStorageGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class3AttachmentsAPI.V1AttachmentsStorageGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AttachmentsStorageGetRequest struct via the builder pattern


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


## V1AttachmentsUploadGet

> V1AttachmentsUploadGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

Get Attach

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
	r, err := apiClient.Class3AttachmentsAPI.V1AttachmentsUploadGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class3AttachmentsAPI.V1AttachmentsUploadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AttachmentsUploadGetRequest struct via the builder pattern


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


## V1Requests1Patch

> map[string]interface{} V1Requests1Patch(ctx).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()

Add Attach (multi) - AddAttachmentsToRequest

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
	resp, r, err := apiClient.Class3AttachmentsAPI.V1Requests1Patch(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class3AttachmentsAPI.V1Requests1Patch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1Requests1Patch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class3AttachmentsAPI.V1Requests1Patch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Requests1PatchRequest struct via the builder pattern


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

