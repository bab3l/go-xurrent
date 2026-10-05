# GetOrganizationsId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**Addresses** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**BusinessUnit** | Pointer to **bool** |  | [optional] 
**BusinessUnitOrganization** | Pointer to **NullableString** |  | [optional] 
**Contacts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**EndUserPrivacy** | Pointer to **bool** |  | [optional] 
**FinancialID** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Manager** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**OrderTemplate** | Pointer to **NullableString** |  | [optional] 
**Parent** | Pointer to **NullableString** |  | [optional] 
**PermittedCustomers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**PictureUri** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**Substitute** | Pointer to **NullableString** |  | [optional] 
**UiExtension** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewGetOrganizationsId200Response

`func NewGetOrganizationsId200Response() *GetOrganizationsId200Response`

NewGetOrganizationsId200Response instantiates a new GetOrganizationsId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationsId200ResponseWithDefaults

`func NewGetOrganizationsId200ResponseWithDefaults() *GetOrganizationsId200Response`

NewGetOrganizationsId200ResponseWithDefaults instantiates a new GetOrganizationsId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetOrganizationsId200Response) GetAccount() GetRequestsId200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetOrganizationsId200Response) GetAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetOrganizationsId200Response) SetAccount(v GetRequestsId200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetOrganizationsId200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAddresses

`func (o *GetOrganizationsId200Response) GetAddresses() []map[string]interface{}`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetOrganizationsId200Response) GetAddressesOk() (*[]map[string]interface{}, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetOrganizationsId200Response) SetAddresses(v []map[string]interface{})`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *GetOrganizationsId200Response) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAttachments

`func (o *GetOrganizationsId200Response) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetOrganizationsId200Response) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetOrganizationsId200Response) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetOrganizationsId200Response) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBusinessUnit

`func (o *GetOrganizationsId200Response) GetBusinessUnit() bool`

GetBusinessUnit returns the BusinessUnit field if non-nil, zero value otherwise.

### GetBusinessUnitOk

`func (o *GetOrganizationsId200Response) GetBusinessUnitOk() (*bool, bool)`

GetBusinessUnitOk returns a tuple with the BusinessUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnit

`func (o *GetOrganizationsId200Response) SetBusinessUnit(v bool)`

SetBusinessUnit sets BusinessUnit field to given value.

### HasBusinessUnit

`func (o *GetOrganizationsId200Response) HasBusinessUnit() bool`

HasBusinessUnit returns a boolean if a field has been set.

### GetBusinessUnitOrganization

`func (o *GetOrganizationsId200Response) GetBusinessUnitOrganization() string`

GetBusinessUnitOrganization returns the BusinessUnitOrganization field if non-nil, zero value otherwise.

### GetBusinessUnitOrganizationOk

`func (o *GetOrganizationsId200Response) GetBusinessUnitOrganizationOk() (*string, bool)`

GetBusinessUnitOrganizationOk returns a tuple with the BusinessUnitOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitOrganization

`func (o *GetOrganizationsId200Response) SetBusinessUnitOrganization(v string)`

SetBusinessUnitOrganization sets BusinessUnitOrganization field to given value.

### HasBusinessUnitOrganization

`func (o *GetOrganizationsId200Response) HasBusinessUnitOrganization() bool`

HasBusinessUnitOrganization returns a boolean if a field has been set.

### SetBusinessUnitOrganizationNil

`func (o *GetOrganizationsId200Response) SetBusinessUnitOrganizationNil(b bool)`

 SetBusinessUnitOrganizationNil sets the value for BusinessUnitOrganization to be an explicit nil

### UnsetBusinessUnitOrganization
`func (o *GetOrganizationsId200Response) UnsetBusinessUnitOrganization()`

UnsetBusinessUnitOrganization ensures that no value is present for BusinessUnitOrganization, not even an explicit nil
### GetContacts

`func (o *GetOrganizationsId200Response) GetContacts() []map[string]interface{}`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *GetOrganizationsId200Response) GetContactsOk() (*[]map[string]interface{}, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *GetOrganizationsId200Response) SetContacts(v []map[string]interface{})`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *GetOrganizationsId200Response) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetOrganizationsId200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOrganizationsId200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOrganizationsId200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetOrganizationsId200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomFields

`func (o *GetOrganizationsId200Response) GetCustomFields() string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GetOrganizationsId200Response) GetCustomFieldsOk() (*string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GetOrganizationsId200Response) SetCustomFields(v string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GetOrganizationsId200Response) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *GetOrganizationsId200Response) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *GetOrganizationsId200Response) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDisabled

`func (o *GetOrganizationsId200Response) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetOrganizationsId200Response) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetOrganizationsId200Response) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetOrganizationsId200Response) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEndUserPrivacy

`func (o *GetOrganizationsId200Response) GetEndUserPrivacy() bool`

GetEndUserPrivacy returns the EndUserPrivacy field if non-nil, zero value otherwise.

### GetEndUserPrivacyOk

`func (o *GetOrganizationsId200Response) GetEndUserPrivacyOk() (*bool, bool)`

GetEndUserPrivacyOk returns a tuple with the EndUserPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserPrivacy

`func (o *GetOrganizationsId200Response) SetEndUserPrivacy(v bool)`

SetEndUserPrivacy sets EndUserPrivacy field to given value.

### HasEndUserPrivacy

`func (o *GetOrganizationsId200Response) HasEndUserPrivacy() bool`

HasEndUserPrivacy returns a boolean if a field has been set.

### GetFinancialID

`func (o *GetOrganizationsId200Response) GetFinancialID() string`

GetFinancialID returns the FinancialID field if non-nil, zero value otherwise.

### GetFinancialIDOk

`func (o *GetOrganizationsId200Response) GetFinancialIDOk() (*string, bool)`

GetFinancialIDOk returns a tuple with the FinancialID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialID

`func (o *GetOrganizationsId200Response) SetFinancialID(v string)`

SetFinancialID sets FinancialID field to given value.

### HasFinancialID

`func (o *GetOrganizationsId200Response) HasFinancialID() bool`

HasFinancialID returns a boolean if a field has been set.

### SetFinancialIDNil

`func (o *GetOrganizationsId200Response) SetFinancialIDNil(b bool)`

 SetFinancialIDNil sets the value for FinancialID to be an explicit nil

### UnsetFinancialID
`func (o *GetOrganizationsId200Response) UnsetFinancialID()`

UnsetFinancialID ensures that no value is present for FinancialID, not even an explicit nil
### GetId

`func (o *GetOrganizationsId200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationsId200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationsId200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationsId200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManager

`func (o *GetOrganizationsId200Response) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *GetOrganizationsId200Response) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *GetOrganizationsId200Response) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *GetOrganizationsId200Response) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *GetOrganizationsId200Response) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *GetOrganizationsId200Response) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetName

`func (o *GetOrganizationsId200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationsId200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationsId200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationsId200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeID

`func (o *GetOrganizationsId200Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *GetOrganizationsId200Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *GetOrganizationsId200Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *GetOrganizationsId200Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetOrderTemplate

`func (o *GetOrganizationsId200Response) GetOrderTemplate() string`

GetOrderTemplate returns the OrderTemplate field if non-nil, zero value otherwise.

### GetOrderTemplateOk

`func (o *GetOrganizationsId200Response) GetOrderTemplateOk() (*string, bool)`

GetOrderTemplateOk returns a tuple with the OrderTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTemplate

`func (o *GetOrganizationsId200Response) SetOrderTemplate(v string)`

SetOrderTemplate sets OrderTemplate field to given value.

### HasOrderTemplate

`func (o *GetOrganizationsId200Response) HasOrderTemplate() bool`

HasOrderTemplate returns a boolean if a field has been set.

### SetOrderTemplateNil

`func (o *GetOrganizationsId200Response) SetOrderTemplateNil(b bool)`

 SetOrderTemplateNil sets the value for OrderTemplate to be an explicit nil

### UnsetOrderTemplate
`func (o *GetOrganizationsId200Response) UnsetOrderTemplate()`

UnsetOrderTemplate ensures that no value is present for OrderTemplate, not even an explicit nil
### GetParent

`func (o *GetOrganizationsId200Response) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetOrganizationsId200Response) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetOrganizationsId200Response) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetOrganizationsId200Response) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *GetOrganizationsId200Response) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *GetOrganizationsId200Response) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetPermittedCustomers

`func (o *GetOrganizationsId200Response) GetPermittedCustomers() []map[string]interface{}`

GetPermittedCustomers returns the PermittedCustomers field if non-nil, zero value otherwise.

### GetPermittedCustomersOk

`func (o *GetOrganizationsId200Response) GetPermittedCustomersOk() (*[]map[string]interface{}, bool)`

GetPermittedCustomersOk returns a tuple with the PermittedCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermittedCustomers

`func (o *GetOrganizationsId200Response) SetPermittedCustomers(v []map[string]interface{})`

SetPermittedCustomers sets PermittedCustomers field to given value.

### HasPermittedCustomers

`func (o *GetOrganizationsId200Response) HasPermittedCustomers() bool`

HasPermittedCustomers returns a boolean if a field has been set.

### GetPictureUri

`func (o *GetOrganizationsId200Response) GetPictureUri() string`

GetPictureUri returns the PictureUri field if non-nil, zero value otherwise.

### GetPictureUriOk

`func (o *GetOrganizationsId200Response) GetPictureUriOk() (*string, bool)`

GetPictureUriOk returns a tuple with the PictureUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureUri

`func (o *GetOrganizationsId200Response) SetPictureUri(v string)`

SetPictureUri sets PictureUri field to given value.

### HasPictureUri

`func (o *GetOrganizationsId200Response) HasPictureUri() bool`

HasPictureUri returns a boolean if a field has been set.

### SetPictureUriNil

`func (o *GetOrganizationsId200Response) SetPictureUriNil(b bool)`

 SetPictureUriNil sets the value for PictureUri to be an explicit nil

### UnsetPictureUri
`func (o *GetOrganizationsId200Response) UnsetPictureUri()`

UnsetPictureUri ensures that no value is present for PictureUri, not even an explicit nil
### GetRegion

`func (o *GetOrganizationsId200Response) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetOrganizationsId200Response) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetOrganizationsId200Response) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetOrganizationsId200Response) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *GetOrganizationsId200Response) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *GetOrganizationsId200Response) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetRemarks

`func (o *GetOrganizationsId200Response) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *GetOrganizationsId200Response) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *GetOrganizationsId200Response) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *GetOrganizationsId200Response) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSource

`func (o *GetOrganizationsId200Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetOrganizationsId200Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetOrganizationsId200Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetOrganizationsId200Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *GetOrganizationsId200Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *GetOrganizationsId200Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *GetOrganizationsId200Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *GetOrganizationsId200Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *GetOrganizationsId200Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *GetOrganizationsId200Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetSubstitute

`func (o *GetOrganizationsId200Response) GetSubstitute() string`

GetSubstitute returns the Substitute field if non-nil, zero value otherwise.

### GetSubstituteOk

`func (o *GetOrganizationsId200Response) GetSubstituteOk() (*string, bool)`

GetSubstituteOk returns a tuple with the Substitute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitute

`func (o *GetOrganizationsId200Response) SetSubstitute(v string)`

SetSubstitute sets Substitute field to given value.

### HasSubstitute

`func (o *GetOrganizationsId200Response) HasSubstitute() bool`

HasSubstitute returns a boolean if a field has been set.

### SetSubstituteNil

`func (o *GetOrganizationsId200Response) SetSubstituteNil(b bool)`

 SetSubstituteNil sets the value for Substitute to be an explicit nil

### UnsetSubstitute
`func (o *GetOrganizationsId200Response) UnsetSubstitute()`

UnsetSubstitute ensures that no value is present for Substitute, not even an explicit nil
### GetUiExtension

`func (o *GetOrganizationsId200Response) GetUiExtension() string`

GetUiExtension returns the UiExtension field if non-nil, zero value otherwise.

### GetUiExtensionOk

`func (o *GetOrganizationsId200Response) GetUiExtensionOk() (*string, bool)`

GetUiExtensionOk returns a tuple with the UiExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiExtension

`func (o *GetOrganizationsId200Response) SetUiExtension(v string)`

SetUiExtension sets UiExtension field to given value.

### HasUiExtension

`func (o *GetOrganizationsId200Response) HasUiExtension() bool`

HasUiExtension returns a boolean if a field has been set.

### SetUiExtensionNil

`func (o *GetOrganizationsId200Response) SetUiExtensionNil(b bool)`

 SetUiExtensionNil sets the value for UiExtension to be an explicit nil

### UnsetUiExtension
`func (o *GetOrganizationsId200Response) UnsetUiExtension()`

UnsetUiExtension ensures that no value is present for UiExtension, not even an explicit nil
### GetUpdatedAt

`func (o *GetOrganizationsId200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetOrganizationsId200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetOrganizationsId200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetOrganizationsId200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


