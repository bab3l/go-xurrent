# \GeneralAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnums**](GeneralAPI.md#GetEnums) | **Get** /v1/enums | GetEnumerationsValues
[**GetMe**](GeneralAPI.md#GetMe) | **Get** /v1/me | GetMyData
[**GetRateLimit**](GeneralAPI.md#GetRateLimit) | **Get** /v1/rate_limit | Get current rate limit status (does not consume quota)



## GetEnums

> map[string]interface{} GetEnums(ctx).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).IfNoneMatch(ifNoneMatch).Execute()

GetEnumerationsValues

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
	ifNoneMatch := "ifNoneMatch_example" // string | Conditional GET; returns 304 Not Modified when unchanged. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeneralAPI.GetEnums(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.GetEnums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnums`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.GetEnums`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **xXurrentLanguage** | **string** | Override response language for enums/errors (e.g. nl, fr). See API introduction. | 
 **ifNoneMatch** | **string** | Conditional GET; returns 304 Not Modified when unchanged. | 

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


## GetMe

> GetMe(ctx).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).IfNoneMatch(ifNoneMatch).Execute()

GetMyData

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
	ifNoneMatch := "ifNoneMatch_example" // string | Conditional GET; returns 304 Not Modified when unchanged. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GeneralAPI.GetMe(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).XXurrentLanguage(xXurrentLanguage).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.GetMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMeRequest struct via the builder pattern


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


## GetRateLimit

> map[string]interface{} GetRateLimit(ctx).XXurrentLanguage(xXurrentLanguage).IfNoneMatch(ifNoneMatch).Execute()

Get current rate limit status (does not consume quota)

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
	xXurrentLanguage := "xXurrentLanguage_example" // string | Override response language for enums/errors (e.g. nl, fr). See API introduction. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Conditional GET; returns 304 Not Modified when unchanged. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeneralAPI.GetRateLimit(context.Background()).XXurrentLanguage(xXurrentLanguage).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.GetRateLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRateLimit`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.GetRateLimit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRateLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xXurrentLanguage** | **string** | Override response language for enums/errors (e.g. nl, fr). See API introduction. | 
 **ifNoneMatch** | **string** | Conditional GET; returns 304 Not Modified when unchanged. | 

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

