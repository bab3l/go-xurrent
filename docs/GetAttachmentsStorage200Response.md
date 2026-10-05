# GetAttachmentsStorage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedExtensions** | Pointer to **[]string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**S3** | Pointer to [**GetAttachmentsStorage200ResponseS3**](GetAttachmentsStorage200ResponseS3.md) |  | [optional] 
**SizeLimit** | Pointer to **float32** |  | [optional] 
**UploadUri** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAttachmentsStorage200Response

`func NewGetAttachmentsStorage200Response() *GetAttachmentsStorage200Response`

NewGetAttachmentsStorage200Response instantiates a new GetAttachmentsStorage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAttachmentsStorage200ResponseWithDefaults

`func NewGetAttachmentsStorage200ResponseWithDefaults() *GetAttachmentsStorage200Response`

NewGetAttachmentsStorage200ResponseWithDefaults instantiates a new GetAttachmentsStorage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedExtensions

`func (o *GetAttachmentsStorage200Response) GetAllowedExtensions() []string`

GetAllowedExtensions returns the AllowedExtensions field if non-nil, zero value otherwise.

### GetAllowedExtensionsOk

`func (o *GetAttachmentsStorage200Response) GetAllowedExtensionsOk() (*[]string, bool)`

GetAllowedExtensionsOk returns a tuple with the AllowedExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedExtensions

`func (o *GetAttachmentsStorage200Response) SetAllowedExtensions(v []string)`

SetAllowedExtensions sets AllowedExtensions field to given value.

### HasAllowedExtensions

`func (o *GetAttachmentsStorage200Response) HasAllowedExtensions() bool`

HasAllowedExtensions returns a boolean if a field has been set.

### GetProvider

`func (o *GetAttachmentsStorage200Response) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GetAttachmentsStorage200Response) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GetAttachmentsStorage200Response) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GetAttachmentsStorage200Response) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetS3

`func (o *GetAttachmentsStorage200Response) GetS3() GetAttachmentsStorage200ResponseS3`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *GetAttachmentsStorage200Response) GetS3Ok() (*GetAttachmentsStorage200ResponseS3, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *GetAttachmentsStorage200Response) SetS3(v GetAttachmentsStorage200ResponseS3)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *GetAttachmentsStorage200Response) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetSizeLimit

`func (o *GetAttachmentsStorage200Response) GetSizeLimit() float32`

GetSizeLimit returns the SizeLimit field if non-nil, zero value otherwise.

### GetSizeLimitOk

`func (o *GetAttachmentsStorage200Response) GetSizeLimitOk() (*float32, bool)`

GetSizeLimitOk returns a tuple with the SizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimit

`func (o *GetAttachmentsStorage200Response) SetSizeLimit(v float32)`

SetSizeLimit sets SizeLimit field to given value.

### HasSizeLimit

`func (o *GetAttachmentsStorage200Response) HasSizeLimit() bool`

HasSizeLimit returns a boolean if a field has been set.

### GetUploadUri

`func (o *GetAttachmentsStorage200Response) GetUploadUri() string`

GetUploadUri returns the UploadUri field if non-nil, zero value otherwise.

### GetUploadUriOk

`func (o *GetAttachmentsStorage200Response) GetUploadUriOk() (*string, bool)`

GetUploadUriOk returns a tuple with the UploadUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUri

`func (o *GetAttachmentsStorage200Response) SetUploadUri(v string)`

SetUploadUri sets UploadUri field to given value.

### HasUploadUri

`func (o *GetAttachmentsStorage200Response) HasUploadUri() bool`

HasUploadUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


