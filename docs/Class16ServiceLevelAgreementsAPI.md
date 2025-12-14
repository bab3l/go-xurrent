# \Class16ServiceLevelAgreementsAPI

All URIs are relative to *https://api.xurrent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1Slas3809Get**](Class16ServiceLevelAgreementsAPI.md#V1Slas3809Get) | **Get** /v1/slas/3809 | GetSLAProperties
[**V1SlasActiveGet**](Class16ServiceLevelAgreementsAPI.md#V1SlasActiveGet) | **Get** /v1/slas/active | GetSLAListActive
[**V1SlasGet**](Class16ServiceLevelAgreementsAPI.md#V1SlasGet) | **Get** /v1/slas | GetSLAListAllByService
[**V1SlasInactiveGet**](Class16ServiceLevelAgreementsAPI.md#V1SlasInactiveGet) | **Get** /v1/slas/inactive | GetSLAListInactive



## V1Slas3809Get

> V1Slas3809Get(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetSLAProperties

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
	r, err := apiClient.Class16ServiceLevelAgreementsAPI.V1Slas3809Get(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class16ServiceLevelAgreementsAPI.V1Slas3809Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1Slas3809GetRequest struct via the builder pattern


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


## V1SlasActiveGet

> map[string]interface{} V1SlasActiveGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetSLAListActive

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
	resp, r, err := apiClient.Class16ServiceLevelAgreementsAPI.V1SlasActiveGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class16ServiceLevelAgreementsAPI.V1SlasActiveGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SlasActiveGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class16ServiceLevelAgreementsAPI.V1SlasActiveGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SlasActiveGetRequest struct via the builder pattern


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


## V1SlasGet

> map[string]interface{} V1SlasGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).ServiceInstance(serviceInstance).Execute()

GetSLAListAllByService

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
	serviceInstance := int32(2439) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class16ServiceLevelAgreementsAPI.V1SlasGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).ServiceInstance(serviceInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class16ServiceLevelAgreementsAPI.V1SlasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SlasGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class16ServiceLevelAgreementsAPI.V1SlasGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SlasGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **x4meAccount** | **string** |  | 
 **serviceInstance** | **int32** |  | 

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


## V1SlasInactiveGet

> map[string]interface{} V1SlasInactiveGet(ctx).Authorization(authorization).X4meAccount(x4meAccount).Execute()

GetSLAListInactive

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
	resp, r, err := apiClient.Class16ServiceLevelAgreementsAPI.V1SlasInactiveGet(context.Background()).Authorization(authorization).X4meAccount(x4meAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class16ServiceLevelAgreementsAPI.V1SlasInactiveGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SlasInactiveGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Class16ServiceLevelAgreementsAPI.V1SlasInactiveGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SlasInactiveGetRequest struct via the builder pattern


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

