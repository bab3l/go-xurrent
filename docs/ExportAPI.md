# \ExportAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImport**](ExportAPI.md#GetImport) | **Get** /v1/import | GetExportJobsList (all)
[**GetImportToken**](ExportAPI.md#GetImportToken) | **Get** /v1/import/{token} | GetExportJobProperties
[**PostExport**](ExportAPI.md#PostExport) | **Post** /v1/export | RunNewExport (sites, teams)



## GetImport

> GetImport(ctx).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()

GetExportJobsList (all)

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
	r, err := apiClient.ExportAPI.GetImport(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).PerPage(perPage).SearchAfter(searchAfter).SearchBefore(searchBefore).Fields(fields).Sort(sort).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.GetImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImportRequest struct via the builder pattern


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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImportToken

> GetImportToken(ctx, token).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).IfNoneMatch(ifNoneMatch).Execute()

GetExportJobProperties

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
	token := "token_example" // string | 
	authorization := "Bearer {{system_user acess token}}" // string |  (optional)
	x4meAccount := "{{X-4me-Account}}" // string |  (optional)
	xXurrentLanguage := "xXurrentLanguage_example" // string | Override response language for enums/errors (e.g. nl, fr). See API introduction. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Conditional GET; returns 304 Not Modified when unchanged. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExportAPI.GetImportToken(context.Background(), token).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.GetImportToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportTokenRequest struct via the builder pattern


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


## PostExport

> PostExport(ctx).Authorization(authorization).X4meAccount(x4meAccount).Type_(type_).From(from).ExportFormat(exportFormat).LineSeparator(lineSeparator).Execute()

RunNewExport (sites, teams)

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
	authorization := "Bearer <oauth-token>" // string |  (optional)
	x4meAccount := "wdc" // string |  (optional)
	type_ := "type__example" // string |  (optional)
	from := int32(56) // int32 |  (optional)
	exportFormat := "exportFormat_example" // string |  (optional)
	lineSeparator := "lineSeparator_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExportAPI.PostExport(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Type_(type_).From(from).ExportFormat(exportFormat).LineSeparator(lineSeparator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.PostExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostExportRequest struct via the builder pattern


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

