# GetUiExtensionsId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**Activate** | Pointer to **NullableString** |  | [optional] 
**ActiveVersion** | Pointer to [**GetUiExtensionsId200ResponseActiveVersion**](GetUiExtensionsId200ResponseActiveVersion.md) |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CompiledCss** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**GetPeopleDisabled200ResponseInnerManager**](GetPeopleDisabled200ResponseInnerManager.md) |  | [optional] 
**Css** | Pointer to **NullableString** |  | [optional] 
**DarkModeSafe** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**FormDefinitionJson** | Pointer to **NullableString** |  | [optional] 
**HideNote** | Pointer to **bool** |  | [optional] 
**Html** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Javascript** | Pointer to **string** |  | [optional] 
**LocalizedHtml** | Pointer to **string** |  | [optional] 
**LocalizedTitle** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**Phrases** | Pointer to **[]string** |  | [optional] 
**PreparedVersion** | Pointer to [**GetUiExtensionsId200ResponseActiveVersion**](GetUiExtensionsId200ResponseActiveVersion.md) |  | [optional] 
**ShowOnComplete** | Pointer to **bool** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 

## Methods

### NewGetUiExtensionsId200Response

`func NewGetUiExtensionsId200Response() *GetUiExtensionsId200Response`

NewGetUiExtensionsId200Response instantiates a new GetUiExtensionsId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUiExtensionsId200ResponseWithDefaults

`func NewGetUiExtensionsId200ResponseWithDefaults() *GetUiExtensionsId200Response`

NewGetUiExtensionsId200ResponseWithDefaults instantiates a new GetUiExtensionsId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetUiExtensionsId200Response) GetAccount() GetRequestsId200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetUiExtensionsId200Response) GetAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetUiExtensionsId200Response) SetAccount(v GetRequestsId200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetUiExtensionsId200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActivate

`func (o *GetUiExtensionsId200Response) GetActivate() string`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *GetUiExtensionsId200Response) GetActivateOk() (*string, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *GetUiExtensionsId200Response) SetActivate(v string)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *GetUiExtensionsId200Response) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### SetActivateNil

`func (o *GetUiExtensionsId200Response) SetActivateNil(b bool)`

 SetActivateNil sets the value for Activate to be an explicit nil

### UnsetActivate
`func (o *GetUiExtensionsId200Response) UnsetActivate()`

UnsetActivate ensures that no value is present for Activate, not even an explicit nil
### GetActiveVersion

`func (o *GetUiExtensionsId200Response) GetActiveVersion() GetUiExtensionsId200ResponseActiveVersion`

GetActiveVersion returns the ActiveVersion field if non-nil, zero value otherwise.

### GetActiveVersionOk

`func (o *GetUiExtensionsId200Response) GetActiveVersionOk() (*GetUiExtensionsId200ResponseActiveVersion, bool)`

GetActiveVersionOk returns a tuple with the ActiveVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveVersion

`func (o *GetUiExtensionsId200Response) SetActiveVersion(v GetUiExtensionsId200ResponseActiveVersion)`

SetActiveVersion sets ActiveVersion field to given value.

### HasActiveVersion

`func (o *GetUiExtensionsId200Response) HasActiveVersion() bool`

HasActiveVersion returns a boolean if a field has been set.

### GetCategory

`func (o *GetUiExtensionsId200Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetUiExtensionsId200Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetUiExtensionsId200Response) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetUiExtensionsId200Response) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCompiledCss

`func (o *GetUiExtensionsId200Response) GetCompiledCss() string`

GetCompiledCss returns the CompiledCss field if non-nil, zero value otherwise.

### GetCompiledCssOk

`func (o *GetUiExtensionsId200Response) GetCompiledCssOk() (*string, bool)`

GetCompiledCssOk returns a tuple with the CompiledCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompiledCss

`func (o *GetUiExtensionsId200Response) SetCompiledCss(v string)`

SetCompiledCss sets CompiledCss field to given value.

### HasCompiledCss

`func (o *GetUiExtensionsId200Response) HasCompiledCss() bool`

HasCompiledCss returns a boolean if a field has been set.

### SetCompiledCssNil

`func (o *GetUiExtensionsId200Response) SetCompiledCssNil(b bool)`

 SetCompiledCssNil sets the value for CompiledCss to be an explicit nil

### UnsetCompiledCss
`func (o *GetUiExtensionsId200Response) UnsetCompiledCss()`

UnsetCompiledCss ensures that no value is present for CompiledCss, not even an explicit nil
### GetCreatedAt

`func (o *GetUiExtensionsId200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetUiExtensionsId200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetUiExtensionsId200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetUiExtensionsId200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetUiExtensionsId200Response) GetCreatedBy() GetPeopleDisabled200ResponseInnerManager`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetUiExtensionsId200Response) GetCreatedByOk() (*GetPeopleDisabled200ResponseInnerManager, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetUiExtensionsId200Response) SetCreatedBy(v GetPeopleDisabled200ResponseInnerManager)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetUiExtensionsId200Response) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCss

`func (o *GetUiExtensionsId200Response) GetCss() string`

GetCss returns the Css field if non-nil, zero value otherwise.

### GetCssOk

`func (o *GetUiExtensionsId200Response) GetCssOk() (*string, bool)`

GetCssOk returns a tuple with the Css field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCss

`func (o *GetUiExtensionsId200Response) SetCss(v string)`

SetCss sets Css field to given value.

### HasCss

`func (o *GetUiExtensionsId200Response) HasCss() bool`

HasCss returns a boolean if a field has been set.

### SetCssNil

`func (o *GetUiExtensionsId200Response) SetCssNil(b bool)`

 SetCssNil sets the value for Css to be an explicit nil

### UnsetCss
`func (o *GetUiExtensionsId200Response) UnsetCss()`

UnsetCss ensures that no value is present for Css, not even an explicit nil
### GetDarkModeSafe

`func (o *GetUiExtensionsId200Response) GetDarkModeSafe() bool`

GetDarkModeSafe returns the DarkModeSafe field if non-nil, zero value otherwise.

### GetDarkModeSafeOk

`func (o *GetUiExtensionsId200Response) GetDarkModeSafeOk() (*bool, bool)`

GetDarkModeSafeOk returns a tuple with the DarkModeSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkModeSafe

`func (o *GetUiExtensionsId200Response) SetDarkModeSafe(v bool)`

SetDarkModeSafe sets DarkModeSafe field to given value.

### HasDarkModeSafe

`func (o *GetUiExtensionsId200Response) HasDarkModeSafe() bool`

HasDarkModeSafe returns a boolean if a field has been set.

### GetDescription

`func (o *GetUiExtensionsId200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetUiExtensionsId200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetUiExtensionsId200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetUiExtensionsId200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetUiExtensionsId200Response) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetUiExtensionsId200Response) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *GetUiExtensionsId200Response) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetUiExtensionsId200Response) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetUiExtensionsId200Response) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetUiExtensionsId200Response) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetFormDefinitionJson

`func (o *GetUiExtensionsId200Response) GetFormDefinitionJson() string`

GetFormDefinitionJson returns the FormDefinitionJson field if non-nil, zero value otherwise.

### GetFormDefinitionJsonOk

`func (o *GetUiExtensionsId200Response) GetFormDefinitionJsonOk() (*string, bool)`

GetFormDefinitionJsonOk returns a tuple with the FormDefinitionJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormDefinitionJson

`func (o *GetUiExtensionsId200Response) SetFormDefinitionJson(v string)`

SetFormDefinitionJson sets FormDefinitionJson field to given value.

### HasFormDefinitionJson

`func (o *GetUiExtensionsId200Response) HasFormDefinitionJson() bool`

HasFormDefinitionJson returns a boolean if a field has been set.

### SetFormDefinitionJsonNil

`func (o *GetUiExtensionsId200Response) SetFormDefinitionJsonNil(b bool)`

 SetFormDefinitionJsonNil sets the value for FormDefinitionJson to be an explicit nil

### UnsetFormDefinitionJson
`func (o *GetUiExtensionsId200Response) UnsetFormDefinitionJson()`

UnsetFormDefinitionJson ensures that no value is present for FormDefinitionJson, not even an explicit nil
### GetHideNote

`func (o *GetUiExtensionsId200Response) GetHideNote() bool`

GetHideNote returns the HideNote field if non-nil, zero value otherwise.

### GetHideNoteOk

`func (o *GetUiExtensionsId200Response) GetHideNoteOk() (*bool, bool)`

GetHideNoteOk returns a tuple with the HideNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideNote

`func (o *GetUiExtensionsId200Response) SetHideNote(v bool)`

SetHideNote sets HideNote field to given value.

### HasHideNote

`func (o *GetUiExtensionsId200Response) HasHideNote() bool`

HasHideNote returns a boolean if a field has been set.

### GetHtml

`func (o *GetUiExtensionsId200Response) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *GetUiExtensionsId200Response) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *GetUiExtensionsId200Response) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *GetUiExtensionsId200Response) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetId

`func (o *GetUiExtensionsId200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUiExtensionsId200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUiExtensionsId200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetUiExtensionsId200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJavascript

`func (o *GetUiExtensionsId200Response) GetJavascript() string`

GetJavascript returns the Javascript field if non-nil, zero value otherwise.

### GetJavascriptOk

`func (o *GetUiExtensionsId200Response) GetJavascriptOk() (*string, bool)`

GetJavascriptOk returns a tuple with the Javascript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavascript

`func (o *GetUiExtensionsId200Response) SetJavascript(v string)`

SetJavascript sets Javascript field to given value.

### HasJavascript

`func (o *GetUiExtensionsId200Response) HasJavascript() bool`

HasJavascript returns a boolean if a field has been set.

### GetLocalizedHtml

`func (o *GetUiExtensionsId200Response) GetLocalizedHtml() string`

GetLocalizedHtml returns the LocalizedHtml field if non-nil, zero value otherwise.

### GetLocalizedHtmlOk

`func (o *GetUiExtensionsId200Response) GetLocalizedHtmlOk() (*string, bool)`

GetLocalizedHtmlOk returns a tuple with the LocalizedHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedHtml

`func (o *GetUiExtensionsId200Response) SetLocalizedHtml(v string)`

SetLocalizedHtml sets LocalizedHtml field to given value.

### HasLocalizedHtml

`func (o *GetUiExtensionsId200Response) HasLocalizedHtml() bool`

HasLocalizedHtml returns a boolean if a field has been set.

### GetLocalizedTitle

`func (o *GetUiExtensionsId200Response) GetLocalizedTitle() string`

GetLocalizedTitle returns the LocalizedTitle field if non-nil, zero value otherwise.

### GetLocalizedTitleOk

`func (o *GetUiExtensionsId200Response) GetLocalizedTitleOk() (*string, bool)`

GetLocalizedTitleOk returns a tuple with the LocalizedTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedTitle

`func (o *GetUiExtensionsId200Response) SetLocalizedTitle(v string)`

SetLocalizedTitle sets LocalizedTitle field to given value.

### HasLocalizedTitle

`func (o *GetUiExtensionsId200Response) HasLocalizedTitle() bool`

HasLocalizedTitle returns a boolean if a field has been set.

### GetName

`func (o *GetUiExtensionsId200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetUiExtensionsId200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetUiExtensionsId200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetUiExtensionsId200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeID

`func (o *GetUiExtensionsId200Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *GetUiExtensionsId200Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *GetUiExtensionsId200Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *GetUiExtensionsId200Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetPhrases

`func (o *GetUiExtensionsId200Response) GetPhrases() []string`

GetPhrases returns the Phrases field if non-nil, zero value otherwise.

### GetPhrasesOk

`func (o *GetUiExtensionsId200Response) GetPhrasesOk() (*[]string, bool)`

GetPhrasesOk returns a tuple with the Phrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhrases

`func (o *GetUiExtensionsId200Response) SetPhrases(v []string)`

SetPhrases sets Phrases field to given value.

### HasPhrases

`func (o *GetUiExtensionsId200Response) HasPhrases() bool`

HasPhrases returns a boolean if a field has been set.

### GetPreparedVersion

`func (o *GetUiExtensionsId200Response) GetPreparedVersion() GetUiExtensionsId200ResponseActiveVersion`

GetPreparedVersion returns the PreparedVersion field if non-nil, zero value otherwise.

### GetPreparedVersionOk

`func (o *GetUiExtensionsId200Response) GetPreparedVersionOk() (*GetUiExtensionsId200ResponseActiveVersion, bool)`

GetPreparedVersionOk returns a tuple with the PreparedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparedVersion

`func (o *GetUiExtensionsId200Response) SetPreparedVersion(v GetUiExtensionsId200ResponseActiveVersion)`

SetPreparedVersion sets PreparedVersion field to given value.

### HasPreparedVersion

`func (o *GetUiExtensionsId200Response) HasPreparedVersion() bool`

HasPreparedVersion returns a boolean if a field has been set.

### GetShowOnComplete

`func (o *GetUiExtensionsId200Response) GetShowOnComplete() bool`

GetShowOnComplete returns the ShowOnComplete field if non-nil, zero value otherwise.

### GetShowOnCompleteOk

`func (o *GetUiExtensionsId200Response) GetShowOnCompleteOk() (*bool, bool)`

GetShowOnCompleteOk returns a tuple with the ShowOnComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnComplete

`func (o *GetUiExtensionsId200Response) SetShowOnComplete(v bool)`

SetShowOnComplete sets ShowOnComplete field to given value.

### HasShowOnComplete

`func (o *GetUiExtensionsId200Response) HasShowOnComplete() bool`

HasShowOnComplete returns a boolean if a field has been set.

### GetSource

`func (o *GetUiExtensionsId200Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetUiExtensionsId200Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetUiExtensionsId200Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetUiExtensionsId200Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *GetUiExtensionsId200Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *GetUiExtensionsId200Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *GetUiExtensionsId200Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *GetUiExtensionsId200Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *GetUiExtensionsId200Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *GetUiExtensionsId200Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetTitle

`func (o *GetUiExtensionsId200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetUiExtensionsId200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetUiExtensionsId200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetUiExtensionsId200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetUiExtensionsId200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetUiExtensionsId200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetUiExtensionsId200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetUiExtensionsId200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GetUiExtensionsId200Response) GetUpdatedBy() GetRequestsId200ResponseCreatedBy`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetUiExtensionsId200Response) GetUpdatedByOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetUiExtensionsId200Response) SetUpdatedBy(v GetRequestsId200ResponseCreatedBy)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GetUiExtensionsId200Response) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


