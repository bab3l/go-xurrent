# \CustomCollectionsAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomCollectionElements**](CustomCollectionsAPI.md#GetCustomCollectionElements) | **Get** /v1/custom_collection_elements | GetCustomCollectionElementsList (by custom_collection_id)
[**GetCustomCollectionElementsId**](CustomCollectionsAPI.md#GetCustomCollectionElementsId) | **Get** /v1/custom_collection_elements/{id} | GetCustomCollectionElementsProperties
[**GetCustomCollections**](CustomCollectionsAPI.md#GetCustomCollections) | **Get** /v1/custom_collections | GetCustomCollectionsList
[**GetCustomCollectionsId**](CustomCollectionsAPI.md#GetCustomCollectionsId) | **Get** /v1/custom_collections/{id} | GetCustomCollectionProperties



## GetCustomCollectionElements

> map[string]interface{} GetCustomCollectionElements(ctx).Authorization(authorization).X4meAccount(x4meAccount).CustomCollection(customCollection).Execute()

GetCustomCollectionElementsList (by custom_collection_id)

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
	customCollection := "product_backlogs" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomCollectionsAPI.GetCustomCollectionElements(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).CustomCollection(customCollection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomCollectionsAPI.GetCustomCollectionElements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomCollectionElements`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomCollectionsAPI.GetCustomCollectionElements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomCollectionElementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **customCollection** | **string** |  | 

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


## GetCustomCollectionElementsId

> map[string]interface{} GetCustomCollectionElementsId(ctx, id).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetCustomCollectionElementsProperties

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
	resp, r, err := apiClient.CustomCollectionsAPI.GetCustomCollectionElementsId(context.Background(), id).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomCollectionsAPI.GetCustomCollectionElementsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomCollectionElementsId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomCollectionsAPI.GetCustomCollectionElementsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomCollectionElementsIdRequest struct via the builder pattern


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


## GetCustomCollections

> map[string]interface{} GetCustomCollections(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetCustomCollectionsList

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomCollectionsAPI.GetCustomCollections(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomCollectionsAPI.GetCustomCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomCollections`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomCollectionsAPI.GetCustomCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomCollectionsRequest struct via the builder pattern


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


## GetCustomCollectionsId

> map[string]interface{} GetCustomCollectionsId(ctx, id).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetCustomCollectionProperties

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
	resp, r, err := apiClient.CustomCollectionsAPI.GetCustomCollectionsId(context.Background(), id).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomCollectionsAPI.GetCustomCollectionsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomCollectionsId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomCollectionsAPI.GetCustomCollectionsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomCollectionsIdRequest struct via the builder pattern


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

