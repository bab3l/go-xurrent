# GetPeople200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Manager** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**PrimaryEmail** | Pointer to **string** |  | [optional] 
**Site** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewGetPeople200ResponseInner

`func NewGetPeople200ResponseInner() *GetPeople200ResponseInner`

NewGetPeople200ResponseInner instantiates a new GetPeople200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPeople200ResponseInnerWithDefaults

`func NewGetPeople200ResponseInnerWithDefaults() *GetPeople200ResponseInner`

NewGetPeople200ResponseInnerWithDefaults instantiates a new GetPeople200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetPeople200ResponseInner) GetAccount() GetRequestsId200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetPeople200ResponseInner) GetAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetPeople200ResponseInner) SetAccount(v GetRequestsId200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetPeople200ResponseInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetPeople200ResponseInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetPeople200ResponseInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetPeople200ResponseInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetPeople200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *GetPeople200ResponseInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPeople200ResponseInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPeople200ResponseInner) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetPeople200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManager

`func (o *GetPeople200ResponseInner) GetManager() GetRequestsId200ResponseCreatedBy`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *GetPeople200ResponseInner) GetManagerOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *GetPeople200ResponseInner) SetManager(v GetRequestsId200ResponseCreatedBy)`

SetManager sets Manager field to given value.

### HasManager

`func (o *GetPeople200ResponseInner) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetName

`func (o *GetPeople200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetPeople200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetPeople200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetPeople200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeID

`func (o *GetPeople200ResponseInner) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *GetPeople200ResponseInner) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *GetPeople200ResponseInner) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *GetPeople200ResponseInner) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetOrganization

`func (o *GetPeople200ResponseInner) GetOrganization() GetRequestsId200ResponseCreatedBy`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GetPeople200ResponseInner) GetOrganizationOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GetPeople200ResponseInner) SetOrganization(v GetRequestsId200ResponseCreatedBy)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *GetPeople200ResponseInner) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPrimaryEmail

`func (o *GetPeople200ResponseInner) GetPrimaryEmail() string`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *GetPeople200ResponseInner) GetPrimaryEmailOk() (*string, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *GetPeople200ResponseInner) SetPrimaryEmail(v string)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *GetPeople200ResponseInner) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### GetSite

`func (o *GetPeople200ResponseInner) GetSite() GetRequestsId200ResponseCreatedBy`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GetPeople200ResponseInner) GetSiteOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GetPeople200ResponseInner) SetSite(v GetRequestsId200ResponseCreatedBy)`

SetSite sets Site field to given value.

### HasSite

`func (o *GetPeople200ResponseInner) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSourceID

`func (o *GetPeople200ResponseInner) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *GetPeople200ResponseInner) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *GetPeople200ResponseInner) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *GetPeople200ResponseInner) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *GetPeople200ResponseInner) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *GetPeople200ResponseInner) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetUpdatedAt

`func (o *GetPeople200ResponseInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetPeople200ResponseInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetPeople200ResponseInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetPeople200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


