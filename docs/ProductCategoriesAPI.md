# \ProductCategoriesAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProductCategories**](ProductCategoriesAPI.md#GetProductCategories) | **Get** /v1/product_categories | GetProductCategoriesList
[**GetProductCategoriesId**](ProductCategoriesAPI.md#GetProductCategoriesId) | **Get** /v1/product_categories/{id} | GetProductCategoriesProperties



## GetProductCategories

> map[string]interface{} GetProductCategories(ctx).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()

GetProductCategoriesList

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
	resp, r, err := apiClient.ProductCategoriesAPI.GetProductCategories(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCategoriesAPI.GetProductCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductCategories`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProductCategoriesAPI.GetProductCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductCategoriesRequest struct via the builder pattern


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


## GetProductCategoriesId

> GetProductCategoriesId(ctx, id).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).IfNoneMatch(ifNoneMatch).Execute()

GetProductCategoriesProperties

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
	r, err := apiClient.ProductCategoriesAPI.GetProductCategoriesId(context.Background(), id).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCategoriesAPI.GetProductCategoriesId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetProductCategoriesIdRequest struct via the builder pattern


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

