# GetRequestsIdNotes200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Internal** | Pointer to **bool** |  | [optional] 
**Medium** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**Person** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 

## Methods

### NewGetRequestsIdNotes200ResponseInner

`func NewGetRequestsIdNotes200ResponseInner() *GetRequestsIdNotes200ResponseInner`

NewGetRequestsIdNotes200ResponseInner instantiates a new GetRequestsIdNotes200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRequestsIdNotes200ResponseInnerWithDefaults

`func NewGetRequestsIdNotes200ResponseInnerWithDefaults() *GetRequestsIdNotes200ResponseInner`

NewGetRequestsIdNotes200ResponseInnerWithDefaults instantiates a new GetRequestsIdNotes200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetRequestsIdNotes200ResponseInner) GetAccount() GetRequestsId200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetRequestsIdNotes200ResponseInner) GetAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetRequestsIdNotes200ResponseInner) SetAccount(v GetRequestsId200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetRequestsIdNotes200ResponseInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAttachments

`func (o *GetRequestsIdNotes200ResponseInner) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetRequestsIdNotes200ResponseInner) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetRequestsIdNotes200ResponseInner) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetRequestsIdNotes200ResponseInner) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetRequestsIdNotes200ResponseInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetRequestsIdNotes200ResponseInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetRequestsIdNotes200ResponseInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetRequestsIdNotes200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *GetRequestsIdNotes200ResponseInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRequestsIdNotes200ResponseInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRequestsIdNotes200ResponseInner) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetRequestsIdNotes200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternal

`func (o *GetRequestsIdNotes200ResponseInner) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *GetRequestsIdNotes200ResponseInner) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *GetRequestsIdNotes200ResponseInner) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *GetRequestsIdNotes200ResponseInner) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetMedium

`func (o *GetRequestsIdNotes200ResponseInner) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *GetRequestsIdNotes200ResponseInner) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *GetRequestsIdNotes200ResponseInner) SetMedium(v string)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *GetRequestsIdNotes200ResponseInner) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetNodeID

`func (o *GetRequestsIdNotes200ResponseInner) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *GetRequestsIdNotes200ResponseInner) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *GetRequestsIdNotes200ResponseInner) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *GetRequestsIdNotes200ResponseInner) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetPerson

`func (o *GetRequestsIdNotes200ResponseInner) GetPerson() GetRequestsId200ResponseCreatedBy`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *GetRequestsIdNotes200ResponseInner) GetPersonOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *GetRequestsIdNotes200ResponseInner) SetPerson(v GetRequestsId200ResponseCreatedBy)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *GetRequestsIdNotes200ResponseInner) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetText

`func (o *GetRequestsIdNotes200ResponseInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GetRequestsIdNotes200ResponseInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GetRequestsIdNotes200ResponseInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *GetRequestsIdNotes200ResponseInner) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


