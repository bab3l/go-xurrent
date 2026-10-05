# PostTeams201Response

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
**Remarks** | Pointer to **NullableString** |  | [optional] 
**ScrumWorkspace** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**TimeZone** | Pointer to **NullableString** |  | [optional] 
**UiExtension** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**WorkHours** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPostTeams201Response

`func NewPostTeams201Response() *PostTeams201Response`

NewPostTeams201Response instantiates a new PostTeams201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTeams201ResponseWithDefaults

`func NewPostTeams201ResponseWithDefaults() *PostTeams201Response`

NewPostTeams201ResponseWithDefaults instantiates a new PostTeams201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgileBoard

`func (o *PostTeams201Response) GetAgileBoard() string`

GetAgileBoard returns the AgileBoard field if non-nil, zero value otherwise.

### GetAgileBoardOk

`func (o *PostTeams201Response) GetAgileBoardOk() (*string, bool)`

GetAgileBoardOk returns a tuple with the AgileBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgileBoard

`func (o *PostTeams201Response) SetAgileBoard(v string)`

SetAgileBoard sets AgileBoard field to given value.

### HasAgileBoard

`func (o *PostTeams201Response) HasAgileBoard() bool`

HasAgileBoard returns a boolean if a field has been set.

### SetAgileBoardNil

`func (o *PostTeams201Response) SetAgileBoardNil(b bool)`

 SetAgileBoardNil sets the value for AgileBoard to be an explicit nil

### UnsetAgileBoard
`func (o *PostTeams201Response) UnsetAgileBoard()`

UnsetAgileBoard ensures that no value is present for AgileBoard, not even an explicit nil
### GetAttachments

`func (o *PostTeams201Response) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PostTeams201Response) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PostTeams201Response) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *PostTeams201Response) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetAutoAssign

`func (o *PostTeams201Response) GetAutoAssign() bool`

GetAutoAssign returns the AutoAssign field if non-nil, zero value otherwise.

### GetAutoAssignOk

`func (o *PostTeams201Response) GetAutoAssignOk() (*bool, bool)`

GetAutoAssignOk returns a tuple with the AutoAssign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAssign

`func (o *PostTeams201Response) SetAutoAssign(v bool)`

SetAutoAssign sets AutoAssign field to given value.

### HasAutoAssign

`func (o *PostTeams201Response) HasAutoAssign() bool`

HasAutoAssign returns a boolean if a field has been set.

### GetConfigurationManager

`func (o *PostTeams201Response) GetConfigurationManager() string`

GetConfigurationManager returns the ConfigurationManager field if non-nil, zero value otherwise.

### GetConfigurationManagerOk

`func (o *PostTeams201Response) GetConfigurationManagerOk() (*string, bool)`

GetConfigurationManagerOk returns a tuple with the ConfigurationManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationManager

`func (o *PostTeams201Response) SetConfigurationManager(v string)`

SetConfigurationManager sets ConfigurationManager field to given value.

### HasConfigurationManager

`func (o *PostTeams201Response) HasConfigurationManager() bool`

HasConfigurationManager returns a boolean if a field has been set.

### SetConfigurationManagerNil

`func (o *PostTeams201Response) SetConfigurationManagerNil(b bool)`

 SetConfigurationManagerNil sets the value for ConfigurationManager to be an explicit nil

### UnsetConfigurationManager
`func (o *PostTeams201Response) UnsetConfigurationManager()`

UnsetConfigurationManager ensures that no value is present for ConfigurationManager, not even an explicit nil
### GetCoordinator

`func (o *PostTeams201Response) GetCoordinator() string`

GetCoordinator returns the Coordinator field if non-nil, zero value otherwise.

### GetCoordinatorOk

`func (o *PostTeams201Response) GetCoordinatorOk() (*string, bool)`

GetCoordinatorOk returns a tuple with the Coordinator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinator

`func (o *PostTeams201Response) SetCoordinator(v string)`

SetCoordinator sets Coordinator field to given value.

### HasCoordinator

`func (o *PostTeams201Response) HasCoordinator() bool`

HasCoordinator returns a boolean if a field has been set.

### SetCoordinatorNil

`func (o *PostTeams201Response) SetCoordinatorNil(b bool)`

 SetCoordinatorNil sets the value for Coordinator to be an explicit nil

### UnsetCoordinator
`func (o *PostTeams201Response) UnsetCoordinator()`

UnsetCoordinator ensures that no value is present for Coordinator, not even an explicit nil
### GetCreatedAt

`func (o *PostTeams201Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostTeams201Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostTeams201Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PostTeams201Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomFields

`func (o *PostTeams201Response) GetCustomFields() string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PostTeams201Response) GetCustomFieldsOk() (*string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PostTeams201Response) SetCustomFields(v string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PostTeams201Response) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *PostTeams201Response) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *PostTeams201Response) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDisabled

`func (o *PostTeams201Response) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PostTeams201Response) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PostTeams201Response) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PostTeams201Response) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetId

`func (o *PostTeams201Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostTeams201Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostTeams201Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *PostTeams201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInboundEmailLocalPart

`func (o *PostTeams201Response) GetInboundEmailLocalPart() string`

GetInboundEmailLocalPart returns the InboundEmailLocalPart field if non-nil, zero value otherwise.

### GetInboundEmailLocalPartOk

`func (o *PostTeams201Response) GetInboundEmailLocalPartOk() (*string, bool)`

GetInboundEmailLocalPartOk returns a tuple with the InboundEmailLocalPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundEmailLocalPart

`func (o *PostTeams201Response) SetInboundEmailLocalPart(v string)`

SetInboundEmailLocalPart sets InboundEmailLocalPart field to given value.

### HasInboundEmailLocalPart

`func (o *PostTeams201Response) HasInboundEmailLocalPart() bool`

HasInboundEmailLocalPart returns a boolean if a field has been set.

### SetInboundEmailLocalPartNil

`func (o *PostTeams201Response) SetInboundEmailLocalPartNil(b bool)`

 SetInboundEmailLocalPartNil sets the value for InboundEmailLocalPart to be an explicit nil

### UnsetInboundEmailLocalPart
`func (o *PostTeams201Response) UnsetInboundEmailLocalPart()`

UnsetInboundEmailLocalPart ensures that no value is present for InboundEmailLocalPart, not even an explicit nil
### GetManager

`func (o *PostTeams201Response) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *PostTeams201Response) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *PostTeams201Response) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *PostTeams201Response) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *PostTeams201Response) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *PostTeams201Response) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetName

`func (o *PostTeams201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostTeams201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostTeams201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostTeams201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeID

`func (o *PostTeams201Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *PostTeams201Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *PostTeams201Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *PostTeams201Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetPictureUri

`func (o *PostTeams201Response) GetPictureUri() string`

GetPictureUri returns the PictureUri field if non-nil, zero value otherwise.

### GetPictureUriOk

`func (o *PostTeams201Response) GetPictureUriOk() (*string, bool)`

GetPictureUriOk returns a tuple with the PictureUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureUri

`func (o *PostTeams201Response) SetPictureUri(v string)`

SetPictureUri sets PictureUri field to given value.

### HasPictureUri

`func (o *PostTeams201Response) HasPictureUri() bool`

HasPictureUri returns a boolean if a field has been set.

### SetPictureUriNil

`func (o *PostTeams201Response) SetPictureUriNil(b bool)`

 SetPictureUriNil sets the value for PictureUri to be an explicit nil

### UnsetPictureUri
`func (o *PostTeams201Response) UnsetPictureUri()`

UnsetPictureUri ensures that no value is present for PictureUri, not even an explicit nil
### GetRemarks

`func (o *PostTeams201Response) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *PostTeams201Response) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *PostTeams201Response) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *PostTeams201Response) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *PostTeams201Response) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *PostTeams201Response) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetScrumWorkspace

`func (o *PostTeams201Response) GetScrumWorkspace() string`

GetScrumWorkspace returns the ScrumWorkspace field if non-nil, zero value otherwise.

### GetScrumWorkspaceOk

`func (o *PostTeams201Response) GetScrumWorkspaceOk() (*string, bool)`

GetScrumWorkspaceOk returns a tuple with the ScrumWorkspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrumWorkspace

`func (o *PostTeams201Response) SetScrumWorkspace(v string)`

SetScrumWorkspace sets ScrumWorkspace field to given value.

### HasScrumWorkspace

`func (o *PostTeams201Response) HasScrumWorkspace() bool`

HasScrumWorkspace returns a boolean if a field has been set.

### SetScrumWorkspaceNil

`func (o *PostTeams201Response) SetScrumWorkspaceNil(b bool)`

 SetScrumWorkspaceNil sets the value for ScrumWorkspace to be an explicit nil

### UnsetScrumWorkspace
`func (o *PostTeams201Response) UnsetScrumWorkspace()`

UnsetScrumWorkspace ensures that no value is present for ScrumWorkspace, not even an explicit nil
### GetSource

`func (o *PostTeams201Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PostTeams201Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PostTeams201Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PostTeams201Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *PostTeams201Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *PostTeams201Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *PostTeams201Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *PostTeams201Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *PostTeams201Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *PostTeams201Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetTimeZone

`func (o *PostTeams201Response) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *PostTeams201Response) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *PostTeams201Response) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *PostTeams201Response) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *PostTeams201Response) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *PostTeams201Response) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetUiExtension

`func (o *PostTeams201Response) GetUiExtension() string`

GetUiExtension returns the UiExtension field if non-nil, zero value otherwise.

### GetUiExtensionOk

`func (o *PostTeams201Response) GetUiExtensionOk() (*string, bool)`

GetUiExtensionOk returns a tuple with the UiExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiExtension

`func (o *PostTeams201Response) SetUiExtension(v string)`

SetUiExtension sets UiExtension field to given value.

### HasUiExtension

`func (o *PostTeams201Response) HasUiExtension() bool`

HasUiExtension returns a boolean if a field has been set.

### SetUiExtensionNil

`func (o *PostTeams201Response) SetUiExtensionNil(b bool)`

 SetUiExtensionNil sets the value for UiExtension to be an explicit nil

### UnsetUiExtension
`func (o *PostTeams201Response) UnsetUiExtension()`

UnsetUiExtension ensures that no value is present for UiExtension, not even an explicit nil
### GetUpdatedAt

`func (o *PostTeams201Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PostTeams201Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PostTeams201Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PostTeams201Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWorkHours

`func (o *PostTeams201Response) GetWorkHours() string`

GetWorkHours returns the WorkHours field if non-nil, zero value otherwise.

### GetWorkHoursOk

`func (o *PostTeams201Response) GetWorkHoursOk() (*string, bool)`

GetWorkHoursOk returns a tuple with the WorkHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkHours

`func (o *PostTeams201Response) SetWorkHours(v string)`

SetWorkHours sets WorkHours field to given value.

### HasWorkHours

`func (o *PostTeams201Response) HasWorkHours() bool`

HasWorkHours returns a boolean if a field has been set.

### SetWorkHoursNil

`func (o *PostTeams201Response) SetWorkHoursNil(b bool)`

 SetWorkHoursNil sets the value for WorkHours to be an explicit nil

### UnsetWorkHours
`func (o *PostTeams201Response) UnsetWorkHours()`

UnsetWorkHours ensures that no value is present for WorkHours, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


