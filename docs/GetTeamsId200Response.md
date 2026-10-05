# GetTeamsId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgileBoard** | Pointer to **NullableString** |  | [optional] 
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**AutoAssign** | Pointer to **bool** |  | [optional] 
**ConfigurationManager** | Pointer to **NullableString** |  | [optional] 
**Coordinator** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**InboundEmailLocalPart** | Pointer to **NullableString** |  | [optional] 
**Manager** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**PictureUri** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**ScrumWorkspace** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**TimeZone** | Pointer to **NullableString** |  | [optional] 
**UiExtension** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**WorkHours** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetTeamsId200Response

`func NewGetTeamsId200Response() *GetTeamsId200Response`

NewGetTeamsId200Response instantiates a new GetTeamsId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTeamsId200ResponseWithDefaults

`func NewGetTeamsId200ResponseWithDefaults() *GetTeamsId200Response`

NewGetTeamsId200ResponseWithDefaults instantiates a new GetTeamsId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgileBoard

`func (o *GetTeamsId200Response) GetAgileBoard() string`

GetAgileBoard returns the AgileBoard field if non-nil, zero value otherwise.

### GetAgileBoardOk

`func (o *GetTeamsId200Response) GetAgileBoardOk() (*string, bool)`

GetAgileBoardOk returns a tuple with the AgileBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgileBoard

`func (o *GetTeamsId200Response) SetAgileBoard(v string)`

SetAgileBoard sets AgileBoard field to given value.

### HasAgileBoard

`func (o *GetTeamsId200Response) HasAgileBoard() bool`

HasAgileBoard returns a boolean if a field has been set.

### SetAgileBoardNil

`func (o *GetTeamsId200Response) SetAgileBoardNil(b bool)`

 SetAgileBoardNil sets the value for AgileBoard to be an explicit nil

### UnsetAgileBoard
`func (o *GetTeamsId200Response) UnsetAgileBoard()`

UnsetAgileBoard ensures that no value is present for AgileBoard, not even an explicit nil
### GetAttachments

`func (o *GetTeamsId200Response) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetTeamsId200Response) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetTeamsId200Response) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetTeamsId200Response) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetAutoAssign

`func (o *GetTeamsId200Response) GetAutoAssign() bool`

GetAutoAssign returns the AutoAssign field if non-nil, zero value otherwise.

### GetAutoAssignOk

`func (o *GetTeamsId200Response) GetAutoAssignOk() (*bool, bool)`

GetAutoAssignOk returns a tuple with the AutoAssign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAssign

`func (o *GetTeamsId200Response) SetAutoAssign(v bool)`

SetAutoAssign sets AutoAssign field to given value.

### HasAutoAssign

`func (o *GetTeamsId200Response) HasAutoAssign() bool`

HasAutoAssign returns a boolean if a field has been set.

### GetConfigurationManager

`func (o *GetTeamsId200Response) GetConfigurationManager() string`

GetConfigurationManager returns the ConfigurationManager field if non-nil, zero value otherwise.

### GetConfigurationManagerOk

`func (o *GetTeamsId200Response) GetConfigurationManagerOk() (*string, bool)`

GetConfigurationManagerOk returns a tuple with the ConfigurationManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationManager

`func (o *GetTeamsId200Response) SetConfigurationManager(v string)`

SetConfigurationManager sets ConfigurationManager field to given value.

### HasConfigurationManager

`func (o *GetTeamsId200Response) HasConfigurationManager() bool`

HasConfigurationManager returns a boolean if a field has been set.

### SetConfigurationManagerNil

`func (o *GetTeamsId200Response) SetConfigurationManagerNil(b bool)`

 SetConfigurationManagerNil sets the value for ConfigurationManager to be an explicit nil

### UnsetConfigurationManager
`func (o *GetTeamsId200Response) UnsetConfigurationManager()`

UnsetConfigurationManager ensures that no value is present for ConfigurationManager, not even an explicit nil
### GetCoordinator

`func (o *GetTeamsId200Response) GetCoordinator() string`

GetCoordinator returns the Coordinator field if non-nil, zero value otherwise.

### GetCoordinatorOk

`func (o *GetTeamsId200Response) GetCoordinatorOk() (*string, bool)`

GetCoordinatorOk returns a tuple with the Coordinator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinator

`func (o *GetTeamsId200Response) SetCoordinator(v string)`

SetCoordinator sets Coordinator field to given value.

### HasCoordinator

`func (o *GetTeamsId200Response) HasCoordinator() bool`

HasCoordinator returns a boolean if a field has been set.

### SetCoordinatorNil

`func (o *GetTeamsId200Response) SetCoordinatorNil(b bool)`

 SetCoordinatorNil sets the value for Coordinator to be an explicit nil

### UnsetCoordinator
`func (o *GetTeamsId200Response) UnsetCoordinator()`

UnsetCoordinator ensures that no value is present for Coordinator, not even an explicit nil
### GetCreatedAt

`func (o *GetTeamsId200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetTeamsId200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetTeamsId200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetTeamsId200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomFields

`func (o *GetTeamsId200Response) GetCustomFields() string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GetTeamsId200Response) GetCustomFieldsOk() (*string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GetTeamsId200Response) SetCustomFields(v string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GetTeamsId200Response) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *GetTeamsId200Response) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *GetTeamsId200Response) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDisabled

`func (o *GetTeamsId200Response) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetTeamsId200Response) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetTeamsId200Response) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetTeamsId200Response) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetId

`func (o *GetTeamsId200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTeamsId200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTeamsId200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetTeamsId200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInboundEmailLocalPart

`func (o *GetTeamsId200Response) GetInboundEmailLocalPart() string`

GetInboundEmailLocalPart returns the InboundEmailLocalPart field if non-nil, zero value otherwise.

### GetInboundEmailLocalPartOk

`func (o *GetTeamsId200Response) GetInboundEmailLocalPartOk() (*string, bool)`

GetInboundEmailLocalPartOk returns a tuple with the InboundEmailLocalPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundEmailLocalPart

`func (o *GetTeamsId200Response) SetInboundEmailLocalPart(v string)`

SetInboundEmailLocalPart sets InboundEmailLocalPart field to given value.

### HasInboundEmailLocalPart

`func (o *GetTeamsId200Response) HasInboundEmailLocalPart() bool`

HasInboundEmailLocalPart returns a boolean if a field has been set.

### SetInboundEmailLocalPartNil

`func (o *GetTeamsId200Response) SetInboundEmailLocalPartNil(b bool)`

 SetInboundEmailLocalPartNil sets the value for InboundEmailLocalPart to be an explicit nil

### UnsetInboundEmailLocalPart
`func (o *GetTeamsId200Response) UnsetInboundEmailLocalPart()`

UnsetInboundEmailLocalPart ensures that no value is present for InboundEmailLocalPart, not even an explicit nil
### GetManager

`func (o *GetTeamsId200Response) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *GetTeamsId200Response) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *GetTeamsId200Response) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *GetTeamsId200Response) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *GetTeamsId200Response) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *GetTeamsId200Response) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetName

`func (o *GetTeamsId200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTeamsId200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTeamsId200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTeamsId200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeID

`func (o *GetTeamsId200Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *GetTeamsId200Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *GetTeamsId200Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *GetTeamsId200Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetPictureUri

`func (o *GetTeamsId200Response) GetPictureUri() string`

GetPictureUri returns the PictureUri field if non-nil, zero value otherwise.

### GetPictureUriOk

`func (o *GetTeamsId200Response) GetPictureUriOk() (*string, bool)`

GetPictureUriOk returns a tuple with the PictureUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureUri

`func (o *GetTeamsId200Response) SetPictureUri(v string)`

SetPictureUri sets PictureUri field to given value.

### HasPictureUri

`func (o *GetTeamsId200Response) HasPictureUri() bool`

HasPictureUri returns a boolean if a field has been set.

### SetPictureUriNil

`func (o *GetTeamsId200Response) SetPictureUriNil(b bool)`

 SetPictureUriNil sets the value for PictureUri to be an explicit nil

### UnsetPictureUri
`func (o *GetTeamsId200Response) UnsetPictureUri()`

UnsetPictureUri ensures that no value is present for PictureUri, not even an explicit nil
### GetRemarks

`func (o *GetTeamsId200Response) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *GetTeamsId200Response) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *GetTeamsId200Response) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *GetTeamsId200Response) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetScrumWorkspace

`func (o *GetTeamsId200Response) GetScrumWorkspace() string`

GetScrumWorkspace returns the ScrumWorkspace field if non-nil, zero value otherwise.

### GetScrumWorkspaceOk

`func (o *GetTeamsId200Response) GetScrumWorkspaceOk() (*string, bool)`

GetScrumWorkspaceOk returns a tuple with the ScrumWorkspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrumWorkspace

`func (o *GetTeamsId200Response) SetScrumWorkspace(v string)`

SetScrumWorkspace sets ScrumWorkspace field to given value.

### HasScrumWorkspace

`func (o *GetTeamsId200Response) HasScrumWorkspace() bool`

HasScrumWorkspace returns a boolean if a field has been set.

### SetScrumWorkspaceNil

`func (o *GetTeamsId200Response) SetScrumWorkspaceNil(b bool)`

 SetScrumWorkspaceNil sets the value for ScrumWorkspace to be an explicit nil

### UnsetScrumWorkspace
`func (o *GetTeamsId200Response) UnsetScrumWorkspace()`

UnsetScrumWorkspace ensures that no value is present for ScrumWorkspace, not even an explicit nil
### GetSource

`func (o *GetTeamsId200Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetTeamsId200Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetTeamsId200Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetTeamsId200Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *GetTeamsId200Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *GetTeamsId200Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *GetTeamsId200Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *GetTeamsId200Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *GetTeamsId200Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *GetTeamsId200Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetTimeZone

`func (o *GetTeamsId200Response) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetTeamsId200Response) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetTeamsId200Response) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetTeamsId200Response) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *GetTeamsId200Response) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *GetTeamsId200Response) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetUiExtension

`func (o *GetTeamsId200Response) GetUiExtension() string`

GetUiExtension returns the UiExtension field if non-nil, zero value otherwise.

### GetUiExtensionOk

`func (o *GetTeamsId200Response) GetUiExtensionOk() (*string, bool)`

GetUiExtensionOk returns a tuple with the UiExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiExtension

`func (o *GetTeamsId200Response) SetUiExtension(v string)`

SetUiExtension sets UiExtension field to given value.

### HasUiExtension

`func (o *GetTeamsId200Response) HasUiExtension() bool`

HasUiExtension returns a boolean if a field has been set.

### SetUiExtensionNil

`func (o *GetTeamsId200Response) SetUiExtensionNil(b bool)`

 SetUiExtensionNil sets the value for UiExtension to be an explicit nil

### UnsetUiExtension
`func (o *GetTeamsId200Response) UnsetUiExtension()`

UnsetUiExtension ensures that no value is present for UiExtension, not even an explicit nil
### GetUpdatedAt

`func (o *GetTeamsId200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetTeamsId200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetTeamsId200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetTeamsId200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWorkHours

`func (o *GetTeamsId200Response) GetWorkHours() string`

GetWorkHours returns the WorkHours field if non-nil, zero value otherwise.

### GetWorkHoursOk

`func (o *GetTeamsId200Response) GetWorkHoursOk() (*string, bool)`

GetWorkHoursOk returns a tuple with the WorkHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkHours

`func (o *GetTeamsId200Response) SetWorkHours(v string)`

SetWorkHours sets WorkHours field to given value.

### HasWorkHours

`func (o *GetTeamsId200Response) HasWorkHours() bool`

HasWorkHours returns a boolean if a field has been set.

### SetWorkHoursNil

`func (o *GetTeamsId200Response) SetWorkHoursNil(b bool)`

 SetWorkHoursNil sets the value for WorkHours to be an explicit nil

### UnsetWorkHours
`func (o *GetTeamsId200Response) UnsetWorkHours()`

UnsetWorkHours ensures that no value is present for WorkHours, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


