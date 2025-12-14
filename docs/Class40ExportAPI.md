# \Class40ExportAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ExportPost**](Class40ExportAPI.md#V1ExportPost) | **Post** /v1/export | RunNewExport (sites, teams)
[**V1Import75715a414338bb1516ee782a470e5372Get**](Class40ExportAPI.md#V1Import75715a414338bb1516ee782a470e5372Get) | **Get** /v1/import/75715a414338bb1516ee782a470e5372 | GetExportJobProperties
[**V1ImportGet**](Class40ExportAPI.md#V1ImportGet) | **Get** /v1/import | GetExportJobsList (all)



## V1ExportPost

> V1ExportPost(ctx).Authorization(authorization).X4meAccount(x4meAccount).Type_(type_).From(from).ExportFormat(exportFormat).LineSeparator(lineSeparator).Execute()

RunNewExport (sites, teams)

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
	authorization := "Bearer <oauth-token>" // string |  (optional)
	x4meAccount := "wdc" // string |  (optional)
	type_ := "type__example" // string |  (optional)
	from := int32(56) // int32 |  (optional)
	exportFormat := "exportFormat_example" // string |  (optional)
	lineSeparator := "lineSeparator_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class40ExportAPI.V1ExportPost(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Type_(type_).From(from).ExportFormat(exportFormat).LineSeparator(lineSeparator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class40ExportAPI.V1ExportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **type_** | **string** |  | 
 **from** | **int32** |  | 
 **exportFormat** | **string** |  | 
 **lineSeparator** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1Import75715a414338bb1516ee782a470e5372Get

> V1Import75715a414338bb1516ee782a470e5372Get(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetExportJobProperties

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
	r, err := apiClient.Class40ExportAPI.V1Import75715a414338bb1516ee782a470e5372Get(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class40ExportAPI.V1Import75715a414338bb1516ee782a470e5372Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Import75715a414338bb1516ee782a470e5372GetRequest struct via the builder pattern


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


## V1ImportGet

> V1ImportGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetExportJobsList (all)

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
	r, err := apiClient.Class40ExportAPI.V1ImportGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class40ExportAPI.V1ImportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ImportGetRequest struct via the builder pattern


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

