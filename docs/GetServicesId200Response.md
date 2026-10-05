# GetServicesId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**Attachments** | Pointer to [**[]GetServicesId200ResponseAttachmentsInner**](GetServicesId200ResponseAttachmentsInner.md) |  | [optional] 
**AvailabilityManager** | Pointer to **NullableString** |  | [optional] 
**CapacityManager** | Pointer to **NullableString** |  | [optional] 
**ChangeManager** | Pointer to **NullableString** |  | [optional] 
**ContinuityManager** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**FirstLineTeam** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Impact** | Pointer to **NullableString** |  | [optional] 
**Keywords** | Pointer to **NullableString** |  | [optional] 
**KnowledgeManager** | Pointer to [**GetPeopleDisabled200ResponseInnerManager**](GetPeopleDisabled200ResponseInnerManager.md) |  | [optional] 
**LocalizedDescription** | Pointer to **string** |  | [optional] 
**LocalizedKeywords** | Pointer to **NullableString** |  | [optional] 
**LocalizedName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**PictureUri** | Pointer to **string** |  | [optional] 
**ProblemManager** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**Provider** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**ReleaseManager** | Pointer to **NullableString** |  | [optional] 
**ServiceCategory** | Pointer to **NullableString** |  | [optional] 
**ServiceOwner** | Pointer to [**GetPeopleDisabled200ResponseInnerManager**](GetPeopleDisabled200ResponseInnerManager.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**SupportTeam** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**Survey** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**UiExtension** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewGetServicesId200Response

`func NewGetServicesId200Response() *GetServicesId200Response`

NewGetServicesId200Response instantiates a new GetServicesId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServicesId200ResponseWithDefaults

`func NewGetServicesId200ResponseWithDefaults() *GetServicesId200Response`

NewGetServicesId200ResponseWithDefaults instantiates a new GetServicesId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetServicesId200Response) GetAccount() GetRequestsId200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetServicesId200Response) GetAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetServicesId200Response) SetAccount(v GetRequestsId200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetServicesId200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAttachments

`func (o *GetServicesId200Response) GetAttachments() []GetServicesId200ResponseAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetServicesId200Response) GetAttachmentsOk() (*[]GetServicesId200ResponseAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetServicesId200Response) SetAttachments(v []GetServicesId200ResponseAttachmentsInner)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetServicesId200Response) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetAvailabilityManager

`func (o *GetServicesId200Response) GetAvailabilityManager() string`

GetAvailabilityManager returns the AvailabilityManager field if non-nil, zero value otherwise.

### GetAvailabilityManagerOk

`func (o *GetServicesId200Response) GetAvailabilityManagerOk() (*string, bool)`

GetAvailabilityManagerOk returns a tuple with the AvailabilityManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityManager

`func (o *GetServicesId200Response) SetAvailabilityManager(v string)`

SetAvailabilityManager sets AvailabilityManager field to given value.

### HasAvailabilityManager

`func (o *GetServicesId200Response) HasAvailabilityManager() bool`

HasAvailabilityManager returns a boolean if a field has been set.

### SetAvailabilityManagerNil

`func (o *GetServicesId200Response) SetAvailabilityManagerNil(b bool)`

 SetAvailabilityManagerNil sets the value for AvailabilityManager to be an explicit nil

### UnsetAvailabilityManager
`func (o *GetServicesId200Response) UnsetAvailabilityManager()`

UnsetAvailabilityManager ensures that no value is present for AvailabilityManager, not even an explicit nil
### GetCapacityManager

`func (o *GetServicesId200Response) GetCapacityManager() string`

GetCapacityManager returns the CapacityManager field if non-nil, zero value otherwise.

### GetCapacityManagerOk

`func (o *GetServicesId200Response) GetCapacityManagerOk() (*string, bool)`

GetCapacityManagerOk returns a tuple with the CapacityManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityManager

`func (o *GetServicesId200Response) SetCapacityManager(v string)`

SetCapacityManager sets CapacityManager field to given value.

### HasCapacityManager

`func (o *GetServicesId200Response) HasCapacityManager() bool`

HasCapacityManager returns a boolean if a field has been set.

### SetCapacityManagerNil

`func (o *GetServicesId200Response) SetCapacityManagerNil(b bool)`

 SetCapacityManagerNil sets the value for CapacityManager to be an explicit nil

### UnsetCapacityManager
`func (o *GetServicesId200Response) UnsetCapacityManager()`

UnsetCapacityManager ensures that no value is present for CapacityManager, not even an explicit nil
### GetChangeManager

`func (o *GetServicesId200Response) GetChangeManager() string`

GetChangeManager returns the ChangeManager field if non-nil, zero value otherwise.

### GetChangeManagerOk

`func (o *GetServicesId200Response) GetChangeManagerOk() (*string, bool)`

GetChangeManagerOk returns a tuple with the ChangeManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeManager

`func (o *GetServicesId200Response) SetChangeManager(v string)`

SetChangeManager sets ChangeManager field to given value.

### HasChangeManager

`func (o *GetServicesId200Response) HasChangeManager() bool`

HasChangeManager returns a boolean if a field has been set.

### SetChangeManagerNil

`func (o *GetServicesId200Response) SetChangeManagerNil(b bool)`

 SetChangeManagerNil sets the value for ChangeManager to be an explicit nil

### UnsetChangeManager
`func (o *GetServicesId200Response) UnsetChangeManager()`

UnsetChangeManager ensures that no value is present for ChangeManager, not even an explicit nil
### GetContinuityManager

`func (o *GetServicesId200Response) GetContinuityManager() string`

GetContinuityManager returns the ContinuityManager field if non-nil, zero value otherwise.

### GetContinuityManagerOk

`func (o *GetServicesId200Response) GetContinuityManagerOk() (*string, bool)`

GetContinuityManagerOk returns a tuple with the ContinuityManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuityManager

`func (o *GetServicesId200Response) SetContinuityManager(v string)`

SetContinuityManager sets ContinuityManager field to given value.

### HasContinuityManager

`func (o *GetServicesId200Response) HasContinuityManager() bool`

HasContinuityManager returns a boolean if a field has been set.

### SetContinuityManagerNil

`func (o *GetServicesId200Response) SetContinuityManagerNil(b bool)`

 SetContinuityManagerNil sets the value for ContinuityManager to be an explicit nil

### UnsetContinuityManager
`func (o *GetServicesId200Response) UnsetContinuityManager()`

UnsetContinuityManager ensures that no value is present for ContinuityManager, not even an explicit nil
### GetCreatedAt

`func (o *GetServicesId200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetServicesId200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetServicesId200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetServicesId200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomFields

`func (o *GetServicesId200Response) GetCustomFields() string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GetServicesId200Response) GetCustomFieldsOk() (*string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GetServicesId200Response) SetCustomFields(v string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GetServicesId200Response) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *GetServicesId200Response) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *GetServicesId200Response) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDescription

`func (o *GetServicesId200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetServicesId200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetServicesId200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetServicesId200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisabled

`func (o *GetServicesId200Response) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetServicesId200Response) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetServicesId200Response) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetServicesId200Response) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetFirstLineTeam

`func (o *GetServicesId200Response) GetFirstLineTeam() GetRequestsId200ResponseCreatedBy`

GetFirstLineTeam returns the FirstLineTeam field if non-nil, zero value otherwise.

### GetFirstLineTeamOk

`func (o *GetServicesId200Response) GetFirstLineTeamOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetFirstLineTeamOk returns a tuple with the FirstLineTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLineTeam

`func (o *GetServicesId200Response) SetFirstLineTeam(v GetRequestsId200ResponseCreatedBy)`

SetFirstLineTeam sets FirstLineTeam field to given value.

### HasFirstLineTeam

`func (o *GetServicesId200Response) HasFirstLineTeam() bool`

HasFirstLineTeam returns a boolean if a field has been set.

### GetId

`func (o *GetServicesId200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetServicesId200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetServicesId200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetServicesId200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImpact

`func (o *GetServicesId200Response) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *GetServicesId200Response) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *GetServicesId200Response) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *GetServicesId200Response) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### SetImpactNil

`func (o *GetServicesId200Response) SetImpactNil(b bool)`

 SetImpactNil sets the value for Impact to be an explicit nil

### UnsetImpact
`func (o *GetServicesId200Response) UnsetImpact()`

UnsetImpact ensures that no value is present for Impact, not even an explicit nil
### GetKeywords

`func (o *GetServicesId200Response) GetKeywords() string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *GetServicesId200Response) GetKeywordsOk() (*string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *GetServicesId200Response) SetKeywords(v string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *GetServicesId200Response) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### SetKeywordsNil

`func (o *GetServicesId200Response) SetKeywordsNil(b bool)`

 SetKeywordsNil sets the value for Keywords to be an explicit nil

### UnsetKeywords
`func (o *GetServicesId200Response) UnsetKeywords()`

UnsetKeywords ensures that no value is present for Keywords, not even an explicit nil
### GetKnowledgeManager

`func (o *GetServicesId200Response) GetKnowledgeManager() GetPeopleDisabled200ResponseInnerManager`

GetKnowledgeManager returns the KnowledgeManager field if non-nil, zero value otherwise.

### GetKnowledgeManagerOk

`func (o *GetServicesId200Response) GetKnowledgeManagerOk() (*GetPeopleDisabled200ResponseInnerManager, bool)`

GetKnowledgeManagerOk returns a tuple with the KnowledgeManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeManager

`func (o *GetServicesId200Response) SetKnowledgeManager(v GetPeopleDisabled200ResponseInnerManager)`

SetKnowledgeManager sets KnowledgeManager field to given value.

### HasKnowledgeManager

`func (o *GetServicesId200Response) HasKnowledgeManager() bool`

HasKnowledgeManager returns a boolean if a field has been set.

### GetLocalizedDescription

`func (o *GetServicesId200Response) GetLocalizedDescription() string`

GetLocalizedDescription returns the LocalizedDescription field if non-nil, zero value otherwise.

### GetLocalizedDescriptionOk

`func (o *GetServicesId200Response) GetLocalizedDescriptionOk() (*string, bool)`

GetLocalizedDescriptionOk returns a tuple with the LocalizedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedDescription

`func (o *GetServicesId200Response) SetLocalizedDescription(v string)`

SetLocalizedDescription sets LocalizedDescription field to given value.

### HasLocalizedDescription

`func (o *GetServicesId200Response) HasLocalizedDescription() bool`

HasLocalizedDescription returns a boolean if a field has been set.

### GetLocalizedKeywords

`func (o *GetServicesId200Response) GetLocalizedKeywords() string`

GetLocalizedKeywords returns the LocalizedKeywords field if non-nil, zero value otherwise.

### GetLocalizedKeywordsOk

`func (o *GetServicesId200Response) GetLocalizedKeywordsOk() (*string, bool)`

GetLocalizedKeywordsOk returns a tuple with the LocalizedKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedKeywords

`func (o *GetServicesId200Response) SetLocalizedKeywords(v string)`

SetLocalizedKeywords sets LocalizedKeywords field to given value.

### HasLocalizedKeywords

`func (o *GetServicesId200Response) HasLocalizedKeywords() bool`

HasLocalizedKeywords returns a boolean if a field has been set.

### SetLocalizedKeywordsNil

`func (o *GetServicesId200Response) SetLocalizedKeywordsNil(b bool)`

 SetLocalizedKeywordsNil sets the value for LocalizedKeywords to be an explicit nil

### UnsetLocalizedKeywords
`func (o *GetServicesId200Response) UnsetLocalizedKeywords()`

UnsetLocalizedKeywords ensures that no value is present for LocalizedKeywords, not even an explicit nil
### GetLocalizedName

`func (o *GetServicesId200Response) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *GetServicesId200Response) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *GetServicesId200Response) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *GetServicesId200Response) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### GetName

`func (o *GetServicesId200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetServicesId200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetServicesId200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetServicesId200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeID

`func (o *GetServicesId200Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *GetServicesId200Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *GetServicesId200Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *GetServicesId200Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetPictureUri

`func (o *GetServicesId200Response) GetPictureUri() string`

GetPictureUri returns the PictureUri field if non-nil, zero value otherwise.

### GetPictureUriOk

`func (o *GetServicesId200Response) GetPictureUriOk() (*string, bool)`

GetPictureUriOk returns a tuple with the PictureUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureUri

`func (o *GetServicesId200Response) SetPictureUri(v string)`

SetPictureUri sets PictureUri field to given value.

### HasPictureUri

`func (o *GetServicesId200Response) HasPictureUri() bool`

HasPictureUri returns a boolean if a field has been set.

### GetProblemManager

`func (o *GetServicesId200Response) GetProblemManager() GetRequestsId200ResponseCreatedBy`

GetProblemManager returns the ProblemManager field if non-nil, zero value otherwise.

### GetProblemManagerOk

`func (o *GetServicesId200Response) GetProblemManagerOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetProblemManagerOk returns a tuple with the ProblemManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemManager

`func (o *GetServicesId200Response) SetProblemManager(v GetRequestsId200ResponseCreatedBy)`

SetProblemManager sets ProblemManager field to given value.

### HasProblemManager

`func (o *GetServicesId200Response) HasProblemManager() bool`

HasProblemManager returns a boolean if a field has been set.

### GetProvider

`func (o *GetServicesId200Response) GetProvider() GetRequestsId200ResponseCreatedBy`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GetServicesId200Response) GetProviderOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GetServicesId200Response) SetProvider(v GetRequestsId200ResponseCreatedBy)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GetServicesId200Response) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReleaseManager

`func (o *GetServicesId200Response) GetReleaseManager() string`

GetReleaseManager returns the ReleaseManager field if non-nil, zero value otherwise.

### GetReleaseManagerOk

`func (o *GetServicesId200Response) GetReleaseManagerOk() (*string, bool)`

GetReleaseManagerOk returns a tuple with the ReleaseManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseManager

`func (o *GetServicesId200Response) SetReleaseManager(v string)`

SetReleaseManager sets ReleaseManager field to given value.

### HasReleaseManager

`func (o *GetServicesId200Response) HasReleaseManager() bool`

HasReleaseManager returns a boolean if a field has been set.

### SetReleaseManagerNil

`func (o *GetServicesId200Response) SetReleaseManagerNil(b bool)`

 SetReleaseManagerNil sets the value for ReleaseManager to be an explicit nil

### UnsetReleaseManager
`func (o *GetServicesId200Response) UnsetReleaseManager()`

UnsetReleaseManager ensures that no value is present for ReleaseManager, not even an explicit nil
### GetServiceCategory

`func (o *GetServicesId200Response) GetServiceCategory() string`

GetServiceCategory returns the ServiceCategory field if non-nil, zero value otherwise.

### GetServiceCategoryOk

`func (o *GetServicesId200Response) GetServiceCategoryOk() (*string, bool)`

GetServiceCategoryOk returns a tuple with the ServiceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCategory

`func (o *GetServicesId200Response) SetServiceCategory(v string)`

SetServiceCategory sets ServiceCategory field to given value.

### HasServiceCategory

`func (o *GetServicesId200Response) HasServiceCategory() bool`

HasServiceCategory returns a boolean if a field has been set.

### SetServiceCategoryNil

`func (o *GetServicesId200Response) SetServiceCategoryNil(b bool)`

 SetServiceCategoryNil sets the value for ServiceCategory to be an explicit nil

### UnsetServiceCategory
`func (o *GetServicesId200Response) UnsetServiceCategory()`

UnsetServiceCategory ensures that no value is present for ServiceCategory, not even an explicit nil
### GetServiceOwner

`func (o *GetServicesId200Response) GetServiceOwner() GetPeopleDisabled200ResponseInnerManager`

GetServiceOwner returns the ServiceOwner field if non-nil, zero value otherwise.

### GetServiceOwnerOk

`func (o *GetServicesId200Response) GetServiceOwnerOk() (*GetPeopleDisabled200ResponseInnerManager, bool)`

GetServiceOwnerOk returns a tuple with the ServiceOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOwner

`func (o *GetServicesId200Response) SetServiceOwner(v GetPeopleDisabled200ResponseInnerManager)`

SetServiceOwner sets ServiceOwner field to given value.

### HasServiceOwner

`func (o *GetServicesId200Response) HasServiceOwner() bool`

HasServiceOwner returns a boolean if a field has been set.

### GetSource

`func (o *GetServicesId200Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetServicesId200Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetServicesId200Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetServicesId200Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *GetServicesId200Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *GetServicesId200Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *GetServicesId200Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *GetServicesId200Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *GetServicesId200Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *GetServicesId200Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetSupportTeam

`func (o *GetServicesId200Response) GetSupportTeam() GetRequestsId200ResponseCreatedBy`

GetSupportTeam returns the SupportTeam field if non-nil, zero value otherwise.

### GetSupportTeamOk

`func (o *GetServicesId200Response) GetSupportTeamOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetSupportTeamOk returns a tuple with the SupportTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTeam

`func (o *GetServicesId200Response) SetSupportTeam(v GetRequestsId200ResponseCreatedBy)`

SetSupportTeam sets SupportTeam field to given value.

### HasSupportTeam

`func (o *GetServicesId200Response) HasSupportTeam() bool`

HasSupportTeam returns a boolean if a field has been set.

### GetSurvey

`func (o *GetServicesId200Response) GetSurvey() GetRequestsId200ResponseCreatedBy`

GetSurvey returns the Survey field if non-nil, zero value otherwise.

### GetSurveyOk

`func (o *GetServicesId200Response) GetSurveyOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetSurveyOk returns a tuple with the Survey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurvey

`func (o *GetServicesId200Response) SetSurvey(v GetRequestsId200ResponseCreatedBy)`

SetSurvey sets Survey field to given value.

### HasSurvey

`func (o *GetServicesId200Response) HasSurvey() bool`

HasSurvey returns a boolean if a field has been set.

### GetUiExtension

`func (o *GetServicesId200Response) GetUiExtension() string`

GetUiExtension returns the UiExtension field if non-nil, zero value otherwise.

### GetUiExtensionOk

`func (o *GetServicesId200Response) GetUiExtensionOk() (*string, bool)`

GetUiExtensionOk returns a tuple with the UiExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiExtension

`func (o *GetServicesId200Response) SetUiExtension(v string)`

SetUiExtension sets UiExtension field to given value.

### HasUiExtension

`func (o *GetServicesId200Response) HasUiExtension() bool`

HasUiExtension returns a boolean if a field has been set.

### SetUiExtensionNil

`func (o *GetServicesId200Response) SetUiExtensionNil(b bool)`

 SetUiExtensionNil sets the value for UiExtension to be an explicit nil

### UnsetUiExtension
`func (o *GetServicesId200Response) UnsetUiExtension()`

UnsetUiExtension ensures that no value is present for UiExtension, not even an explicit nil
### GetUpdatedAt

`func (o *GetServicesId200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetServicesId200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetServicesId200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetServicesId200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


