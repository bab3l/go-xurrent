# \Class23CustomCollectionsAndElementsAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CustomCollectionElements31328Get**](Class23CustomCollectionsAndElementsAPI.md#V1CustomCollectionElements31328Get) | **Get** /v1/custom_collection_elements/31328 | GetCustomCollectionElementsProperties
[**V1CustomCollectionElementsGet**](Class23CustomCollectionsAndElementsAPI.md#V1CustomCollectionElementsGet) | **Get** /v1/custom_collection_elements | GetCustomCollectionElementsList (by custom_collection_id)
[**V1CustomCollections894Get**](Class23CustomCollectionsAndElementsAPI.md#V1CustomCollections894Get) | **Get** /v1/custom_collections/894 | GetCustomCollectionProperties
[**V1CustomCollectionsGet**](Class23CustomCollectionsAndElementsAPI.md#V1CustomCollectionsGet) | **Get** /v1/custom_collections | GetCustomCollectionsList



## V1CustomCollectionElements31328Get

> map[string]interface{} V1CustomCollectionElements31328Get(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetCustomCollectionElementsProperties

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
	resp, r, err := apiClient.Class23CustomCollectionsAndElementsAPI.V1CustomCollectionElements31328Get(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class23CustomCollectionsAndElementsAPI.V1CustomCollectionElements31328Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CustomCollectionElements31328Get`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class23CustomCollectionsAndElementsAPI.V1CustomCollectionElements31328Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CustomCollectionElements31328GetRequest struct via the builder pattern


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


## V1CustomCollectionElementsGet

> map[string]interface{} V1CustomCollectionElementsGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).CustomCollection(customCollection).Execute()

GetCustomCollectionElementsList (by custom_collection_id)

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
	customCollection := "product_backlogs" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class23CustomCollectionsAndElementsAPI.V1CustomCollectionElementsGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).CustomCollection(customCollection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class23CustomCollectionsAndElementsAPI.V1CustomCollectionElementsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CustomCollectionElementsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class23CustomCollectionsAndElementsAPI.V1CustomCollectionElementsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CustomCollectionElementsGetRequest struct via the builder pattern


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


## V1CustomCollections894Get

> map[string]interface{} V1CustomCollections894Get(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetCustomCollectionProperties

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
	resp, r, err := apiClient.Class23CustomCollectionsAndElementsAPI.V1CustomCollections894Get(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class23CustomCollectionsAndElementsAPI.V1CustomCollections894Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CustomCollections894Get`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class23CustomCollectionsAndElementsAPI.V1CustomCollections894Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CustomCollections894GetRequest struct via the builder pattern


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


## V1CustomCollectionsGet

> map[string]interface{} V1CustomCollectionsGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetCustomCollectionsList

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
	resp, r, err := apiClient.Class23CustomCollectionsAndElementsAPI.V1CustomCollectionsGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class23CustomCollectionsAndElementsAPI.V1CustomCollectionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CustomCollectionsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class23CustomCollectionsAndElementsAPI.V1CustomCollectionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CustomCollectionsGetRequest struct via the builder pattern


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

