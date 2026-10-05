# GetSitesId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Integration** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**PictureUri** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**UiExtension** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Zip** | Pointer to **string** |  | [optional] 

## Methods

### NewGetSitesId200Response

`func NewGetSitesId200Response() *GetSitesId200Response`

NewGetSitesId200Response instantiates a new GetSitesId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSitesId200ResponseWithDefaults

`func NewGetSitesId200ResponseWithDefaults() *GetSitesId200Response`

NewGetSitesId200ResponseWithDefaults instantiates a new GetSitesId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetSitesId200Response) GetAccount() GetRequestsId200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetSitesId200Response) GetAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetSitesId200Response) SetAccount(v GetRequestsId200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetSitesId200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAddress

`func (o *GetSitesId200Response) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetSitesId200Response) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetSitesId200Response) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetSitesId200Response) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAttachments

`func (o *GetSitesId200Response) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetSitesId200Response) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetSitesId200Response) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetSitesId200Response) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCity

`func (o *GetSitesId200Response) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GetSitesId200Response) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GetSitesId200Response) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GetSitesId200Response) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *GetSitesId200Response) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetSitesId200Response) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetSitesId200Response) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetSitesId200Response) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetSitesId200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetSitesId200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetSitesId200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetSitesId200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomFields

`func (o *GetSitesId200Response) GetCustomFields() string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GetSitesId200Response) GetCustomFieldsOk() (*string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GetSitesId200Response) SetCustomFields(v string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GetSitesId200Response) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *GetSitesId200Response) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *GetSitesId200Response) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDisabled

`func (o *GetSitesId200Response) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetSitesId200Response) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetSitesId200Response) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetSitesId200Response) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetId

`func (o *GetSitesId200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSitesId200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSitesId200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetSitesId200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegration

`func (o *GetSitesId200Response) GetIntegration() bool`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *GetSitesId200Response) GetIntegrationOk() (*bool, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *GetSitesId200Response) SetIntegration(v bool)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *GetSitesId200Response) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetName

`func (o *GetSitesId200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSitesId200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSitesId200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSitesId200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeID

`func (o *GetSitesId200Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *GetSitesId200Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *GetSitesId200Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *GetSitesId200Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetPictureUri

`func (o *GetSitesId200Response) GetPictureUri() string`

GetPictureUri returns the PictureUri field if non-nil, zero value otherwise.

### GetPictureUriOk

`func (o *GetSitesId200Response) GetPictureUriOk() (*string, bool)`

GetPictureUriOk returns a tuple with the PictureUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureUri

`func (o *GetSitesId200Response) SetPictureUri(v string)`

SetPictureUri sets PictureUri field to given value.

### HasPictureUri

`func (o *GetSitesId200Response) HasPictureUri() bool`

HasPictureUri returns a boolean if a field has been set.

### SetPictureUriNil

`func (o *GetSitesId200Response) SetPictureUriNil(b bool)`

 SetPictureUriNil sets the value for PictureUri to be an explicit nil

### UnsetPictureUri
`func (o *GetSitesId200Response) UnsetPictureUri()`

UnsetPictureUri ensures that no value is present for PictureUri, not even an explicit nil
### GetRemarks

`func (o *GetSitesId200Response) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *GetSitesId200Response) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *GetSitesId200Response) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *GetSitesId200Response) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSource

`func (o *GetSitesId200Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetSitesId200Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetSitesId200Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetSitesId200Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *GetSitesId200Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *GetSitesId200Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *GetSitesId200Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *GetSitesId200Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *GetSitesId200Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *GetSitesId200Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetState

`func (o *GetSitesId200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetSitesId200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetSitesId200Response) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetSitesId200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetSitesId200Response) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetSitesId200Response) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetSitesId200Response) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetSitesId200Response) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUiExtension

`func (o *GetSitesId200Response) GetUiExtension() string`

GetUiExtension returns the UiExtension field if non-nil, zero value otherwise.

### GetUiExtensionOk

`func (o *GetSitesId200Response) GetUiExtensionOk() (*string, bool)`

GetUiExtensionOk returns a tuple with the UiExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiExtension

`func (o *GetSitesId200Response) SetUiExtension(v string)`

SetUiExtension sets UiExtension field to given value.

### HasUiExtension

`func (o *GetSitesId200Response) HasUiExtension() bool`

HasUiExtension returns a boolean if a field has been set.

### SetUiExtensionNil

`func (o *GetSitesId200Response) SetUiExtensionNil(b bool)`

 SetUiExtensionNil sets the value for UiExtension to be an explicit nil

### UnsetUiExtension
`func (o *GetSitesId200Response) UnsetUiExtension()`

UnsetUiExtension ensures that no value is present for UiExtension, not even an explicit nil
### GetUpdatedAt

`func (o *GetSitesId200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetSitesId200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetSitesId200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetSitesId200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetZip

`func (o *GetSitesId200Response) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *GetSitesId200Response) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *GetSitesId200Response) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *GetSitesId200Response) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


