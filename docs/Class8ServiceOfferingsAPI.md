# \Class8ServiceOfferingsAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ServiceOfferings1AuditGet**](Class8ServiceOfferingsAPI.md#V1ServiceOfferings1AuditGet) | **Get** /v1/service_offerings/1/audit | GetServiceOfferAuditEntries
[**V1ServiceOfferings1AuditGet_0**](Class8ServiceOfferingsAPI.md#V1ServiceOfferings1AuditGet_0) | **Get** /v1/service_offerings/1/audit/ | GetServiceOfferAuditEntries_UpdatedAfterDDMMYY
[**V1ServiceOfferings1Get**](Class8ServiceOfferingsAPI.md#V1ServiceOfferings1Get) | **Get** /v1/service_offerings/1 | GetServiceOfferById
[**V1ServiceOfferingsGet**](Class8ServiceOfferingsAPI.md#V1ServiceOfferingsGet) | **Get** /v1/service_offerings | GetServiceOfferList



## V1ServiceOfferings1AuditGet

> V1ServiceOfferings1AuditGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetServiceOfferAuditEntries

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
	r, err := apiClient.Class8ServiceOfferingsAPI.V1ServiceOfferings1AuditGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class8ServiceOfferingsAPI.V1ServiceOfferings1AuditGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ServiceOfferings1AuditGetRequest struct via the builder pattern


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


## V1ServiceOfferings1AuditGet_0

> V1ServiceOfferings1AuditGet_0(ctx).Authorization(authorization).X4meAccount(x4meAccount).Action(action).CreatedAt(createdAt).Execute()

GetServiceOfferAuditEntries_UpdatedAfterDDMMYY

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
	action := "update" // string |  (optional)
	createdAt := ">2022-11-11T13:25:27" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Class8ServiceOfferingsAPI.V1ServiceOfferings1AuditGet_0(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Action(action).CreatedAt(createdAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class8ServiceOfferingsAPI.V1ServiceOfferings1AuditGet_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ServiceOfferings1AuditGet_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **action** | **string** |  | 
 **createdAt** | **string** |  | 

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


## V1ServiceOfferings1Get

> V1ServiceOfferings1Get(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetServiceOfferById

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
	r, err := apiClient.Class8ServiceOfferingsAPI.V1ServiceOfferings1Get(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class8ServiceOfferingsAPI.V1ServiceOfferings1Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ServiceOfferings1GetRequest struct via the builder pattern


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


## V1ServiceOfferingsGet

> V1ServiceOfferingsGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetServiceOfferList

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
	r, err := apiClient.Class8ServiceOfferingsAPI.V1ServiceOfferingsGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class8ServiceOfferingsAPI.V1ServiceOfferingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ServiceOfferingsGetRequest struct via the builder pattern


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

