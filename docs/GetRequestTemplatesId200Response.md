# GetRequestTemplatesId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**AssetSelection** | Pointer to **bool** |  | [optional] 
**AssignAfterWorkflowCompletion** | Pointer to **bool** |  | [optional] 
**AssignToSelf** | Pointer to **bool** |  | [optional] 
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Ci** | Pointer to **NullableString** |  | [optional] 
**CompletionReason** | Pointer to **NullableString** |  | [optional] 
**CopySubjectToRequests** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DesiredCompletion** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**EffortClass** | Pointer to **NullableString** |  | [optional] 
**EndUsers** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**Keywords** | Pointer to **NullableString** |  | [optional] 
**LocalizedInstructions** | Pointer to **NullableString** |  | [optional] 
**LocalizedKeywords** | Pointer to **NullableString** |  | [optional] 
**LocalizedNote** | Pointer to **NullableString** |  | [optional] 
**LocalizedRegistrationHints** | Pointer to **NullableString** |  | [optional] 
**LocalizedSubject** | Pointer to **string** |  | [optional] 
**Member** | Pointer to **NullableString** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **NullableString** |  | [optional] 
**PlannedEffort** | Pointer to **NullableString** |  | [optional] 
**RegistrationHints** | Pointer to **NullableString** |  | [optional] 
**ResolutionTarget** | Pointer to **NullableString** |  | [optional] 
**RfcType** | Pointer to **NullableString** |  | [optional] 
**Service** | Pointer to [**GetRequestsIdCis200ResponseInnerService**](GetRequestsIdCis200ResponseInnerService.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**Specialists** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Supplier** | Pointer to **NullableString** |  | [optional] 
**SupportHours** | Pointer to **NullableString** |  | [optional] 
**Team** | Pointer to **NullableString** |  | [optional] 
**TimeZone** | Pointer to **NullableString** |  | [optional] 
**TimesApplied** | Pointer to **float32** |  | [optional] 
**UiExtension** | Pointer to [**GetRequestTemplatesId200ResponseUiExtension**](GetRequestTemplatesId200ResponseUiExtension.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Urgent** | Pointer to **bool** |  | [optional] 
**WorkflowManager** | Pointer to **NullableString** |  | [optional] 
**WorkflowTemplate** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetRequestTemplatesId200Response

`func NewGetRequestTemplatesId200Response() *GetRequestTemplatesId200Response`

NewGetRequestTemplatesId200Response instantiates a new GetRequestTemplatesId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRequestTemplatesId200ResponseWithDefaults

`func NewGetRequestTemplatesId200ResponseWithDefaults() *GetRequestTemplatesId200Response`

NewGetRequestTemplatesId200ResponseWithDefaults instantiates a new GetRequestTemplatesId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetRequestTemplatesId200Response) GetAccount() GetRequestsId200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetRequestTemplatesId200Response) GetAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetRequestTemplatesId200Response) SetAccount(v GetRequestsId200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetRequestTemplatesId200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAssetSelection

`func (o *GetRequestTemplatesId200Response) GetAssetSelection() bool`

GetAssetSelection returns the AssetSelection field if non-nil, zero value otherwise.

### GetAssetSelectionOk

`func (o *GetRequestTemplatesId200Response) GetAssetSelectionOk() (*bool, bool)`

GetAssetSelectionOk returns a tuple with the AssetSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSelection

`func (o *GetRequestTemplatesId200Response) SetAssetSelection(v bool)`

SetAssetSelection sets AssetSelection field to given value.

### HasAssetSelection

`func (o *GetRequestTemplatesId200Response) HasAssetSelection() bool`

HasAssetSelection returns a boolean if a field has been set.

### GetAssignAfterWorkflowCompletion

`func (o *GetRequestTemplatesId200Response) GetAssignAfterWorkflowCompletion() bool`

GetAssignAfterWorkflowCompletion returns the AssignAfterWorkflowCompletion field if non-nil, zero value otherwise.

### GetAssignAfterWorkflowCompletionOk

`func (o *GetRequestTemplatesId200Response) GetAssignAfterWorkflowCompletionOk() (*bool, bool)`

GetAssignAfterWorkflowCompletionOk returns a tuple with the AssignAfterWorkflowCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignAfterWorkflowCompletion

`func (o *GetRequestTemplatesId200Response) SetAssignAfterWorkflowCompletion(v bool)`

SetAssignAfterWorkflowCompletion sets AssignAfterWorkflowCompletion field to given value.

### HasAssignAfterWorkflowCompletion

`func (o *GetRequestTemplatesId200Response) HasAssignAfterWorkflowCompletion() bool`

HasAssignAfterWorkflowCompletion returns a boolean if a field has been set.

### GetAssignToSelf

`func (o *GetRequestTemplatesId200Response) GetAssignToSelf() bool`

GetAssignToSelf returns the AssignToSelf field if non-nil, zero value otherwise.

### GetAssignToSelfOk

`func (o *GetRequestTemplatesId200Response) GetAssignToSelfOk() (*bool, bool)`

GetAssignToSelfOk returns a tuple with the AssignToSelf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignToSelf

`func (o *GetRequestTemplatesId200Response) SetAssignToSelf(v bool)`

SetAssignToSelf sets AssignToSelf field to given value.

### HasAssignToSelf

`func (o *GetRequestTemplatesId200Response) HasAssignToSelf() bool`

HasAssignToSelf returns a boolean if a field has been set.

### GetAttachments

`func (o *GetRequestTemplatesId200Response) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetRequestTemplatesId200Response) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetRequestTemplatesId200Response) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetRequestTemplatesId200Response) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCategory

`func (o *GetRequestTemplatesId200Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetRequestTemplatesId200Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetRequestTemplatesId200Response) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetRequestTemplatesId200Response) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCi

`func (o *GetRequestTemplatesId200Response) GetCi() string`

GetCi returns the Ci field if non-nil, zero value otherwise.

### GetCiOk

`func (o *GetRequestTemplatesId200Response) GetCiOk() (*string, bool)`

GetCiOk returns a tuple with the Ci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCi

`func (o *GetRequestTemplatesId200Response) SetCi(v string)`

SetCi sets Ci field to given value.

### HasCi

`func (o *GetRequestTemplatesId200Response) HasCi() bool`

HasCi returns a boolean if a field has been set.

### SetCiNil

`func (o *GetRequestTemplatesId200Response) SetCiNil(b bool)`

 SetCiNil sets the value for Ci to be an explicit nil

### UnsetCi
`func (o *GetRequestTemplatesId200Response) UnsetCi()`

UnsetCi ensures that no value is present for Ci, not even an explicit nil
### GetCompletionReason

`func (o *GetRequestTemplatesId200Response) GetCompletionReason() string`

GetCompletionReason returns the CompletionReason field if non-nil, zero value otherwise.

### GetCompletionReasonOk

`func (o *GetRequestTemplatesId200Response) GetCompletionReasonOk() (*string, bool)`

GetCompletionReasonOk returns a tuple with the CompletionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionReason

`func (o *GetRequestTemplatesId200Response) SetCompletionReason(v string)`

SetCompletionReason sets CompletionReason field to given value.

### HasCompletionReason

`func (o *GetRequestTemplatesId200Response) HasCompletionReason() bool`

HasCompletionReason returns a boolean if a field has been set.

### SetCompletionReasonNil

`func (o *GetRequestTemplatesId200Response) SetCompletionReasonNil(b bool)`

 SetCompletionReasonNil sets the value for CompletionReason to be an explicit nil

### UnsetCompletionReason
`func (o *GetRequestTemplatesId200Response) UnsetCompletionReason()`

UnsetCompletionReason ensures that no value is present for CompletionReason, not even an explicit nil
### GetCopySubjectToRequests

`func (o *GetRequestTemplatesId200Response) GetCopySubjectToRequests() bool`

GetCopySubjectToRequests returns the CopySubjectToRequests field if non-nil, zero value otherwise.

### GetCopySubjectToRequestsOk

`func (o *GetRequestTemplatesId200Response) GetCopySubjectToRequestsOk() (*bool, bool)`

GetCopySubjectToRequestsOk returns a tuple with the CopySubjectToRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySubjectToRequests

`func (o *GetRequestTemplatesId200Response) SetCopySubjectToRequests(v bool)`

SetCopySubjectToRequests sets CopySubjectToRequests field to given value.

### HasCopySubjectToRequests

`func (o *GetRequestTemplatesId200Response) HasCopySubjectToRequests() bool`

HasCopySubjectToRequests returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetRequestTemplatesId200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetRequestTemplatesId200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetRequestTemplatesId200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetRequestTemplatesId200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDesiredCompletion

`func (o *GetRequestTemplatesId200Response) GetDesiredCompletion() string`

GetDesiredCompletion returns the DesiredCompletion field if non-nil, zero value otherwise.

### GetDesiredCompletionOk

`func (o *GetRequestTemplatesId200Response) GetDesiredCompletionOk() (*string, bool)`

GetDesiredCompletionOk returns a tuple with the DesiredCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredCompletion

`func (o *GetRequestTemplatesId200Response) SetDesiredCompletion(v string)`

SetDesiredCompletion sets DesiredCompletion field to given value.

### HasDesiredCompletion

`func (o *GetRequestTemplatesId200Response) HasDesiredCompletion() bool`

HasDesiredCompletion returns a boolean if a field has been set.

### SetDesiredCompletionNil

`func (o *GetRequestTemplatesId200Response) SetDesiredCompletionNil(b bool)`

 SetDesiredCompletionNil sets the value for DesiredCompletion to be an explicit nil

### UnsetDesiredCompletion
`func (o *GetRequestTemplatesId200Response) UnsetDesiredCompletion()`

UnsetDesiredCompletion ensures that no value is present for DesiredCompletion, not even an explicit nil
### GetDisabled

`func (o *GetRequestTemplatesId200Response) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetRequestTemplatesId200Response) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetRequestTemplatesId200Response) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetRequestTemplatesId200Response) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEffortClass

`func (o *GetRequestTemplatesId200Response) GetEffortClass() string`

GetEffortClass returns the EffortClass field if non-nil, zero value otherwise.

### GetEffortClassOk

`func (o *GetRequestTemplatesId200Response) GetEffortClassOk() (*string, bool)`

GetEffortClassOk returns a tuple with the EffortClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffortClass

`func (o *GetRequestTemplatesId200Response) SetEffortClass(v string)`

SetEffortClass sets EffortClass field to given value.

### HasEffortClass

`func (o *GetRequestTemplatesId200Response) HasEffortClass() bool`

HasEffortClass returns a boolean if a field has been set.

### SetEffortClassNil

`func (o *GetRequestTemplatesId200Response) SetEffortClassNil(b bool)`

 SetEffortClassNil sets the value for EffortClass to be an explicit nil

### UnsetEffortClass
`func (o *GetRequestTemplatesId200Response) UnsetEffortClass()`

UnsetEffortClass ensures that no value is present for EffortClass, not even an explicit nil
### GetEndUsers

`func (o *GetRequestTemplatesId200Response) GetEndUsers() bool`

GetEndUsers returns the EndUsers field if non-nil, zero value otherwise.

### GetEndUsersOk

`func (o *GetRequestTemplatesId200Response) GetEndUsersOk() (*bool, bool)`

GetEndUsersOk returns a tuple with the EndUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUsers

`func (o *GetRequestTemplatesId200Response) SetEndUsers(v bool)`

SetEndUsers sets EndUsers field to given value.

### HasEndUsers

`func (o *GetRequestTemplatesId200Response) HasEndUsers() bool`

HasEndUsers returns a boolean if a field has been set.

### GetId

`func (o *GetRequestTemplatesId200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRequestTemplatesId200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRequestTemplatesId200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetRequestTemplatesId200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImpact

`func (o *GetRequestTemplatesId200Response) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *GetRequestTemplatesId200Response) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *GetRequestTemplatesId200Response) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *GetRequestTemplatesId200Response) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetInstructions

`func (o *GetRequestTemplatesId200Response) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *GetRequestTemplatesId200Response) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *GetRequestTemplatesId200Response) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *GetRequestTemplatesId200Response) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *GetRequestTemplatesId200Response) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *GetRequestTemplatesId200Response) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetKeywords

`func (o *GetRequestTemplatesId200Response) GetKeywords() string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *GetRequestTemplatesId200Response) GetKeywordsOk() (*string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *GetRequestTemplatesId200Response) SetKeywords(v string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *GetRequestTemplatesId200Response) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### SetKeywordsNil

`func (o *GetRequestTemplatesId200Response) SetKeywordsNil(b bool)`

 SetKeywordsNil sets the value for Keywords to be an explicit nil

### UnsetKeywords
`func (o *GetRequestTemplatesId200Response) UnsetKeywords()`

UnsetKeywords ensures that no value is present for Keywords, not even an explicit nil
### GetLocalizedInstructions

`func (o *GetRequestTemplatesId200Response) GetLocalizedInstructions() string`

GetLocalizedInstructions returns the LocalizedInstructions field if non-nil, zero value otherwise.

### GetLocalizedInstructionsOk

`func (o *GetRequestTemplatesId200Response) GetLocalizedInstructionsOk() (*string, bool)`

GetLocalizedInstructionsOk returns a tuple with the LocalizedInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedInstructions

`func (o *GetRequestTemplatesId200Response) SetLocalizedInstructions(v string)`

SetLocalizedInstructions sets LocalizedInstructions field to given value.

### HasLocalizedInstructions

`func (o *GetRequestTemplatesId200Response) HasLocalizedInstructions() bool`

HasLocalizedInstructions returns a boolean if a field has been set.

### SetLocalizedInstructionsNil

`func (o *GetRequestTemplatesId200Response) SetLocalizedInstructionsNil(b bool)`

 SetLocalizedInstructionsNil sets the value for LocalizedInstructions to be an explicit nil

### UnsetLocalizedInstructions
`func (o *GetRequestTemplatesId200Response) UnsetLocalizedInstructions()`

UnsetLocalizedInstructions ensures that no value is present for LocalizedInstructions, not even an explicit nil
### GetLocalizedKeywords

`func (o *GetRequestTemplatesId200Response) GetLocalizedKeywords() string`

GetLocalizedKeywords returns the LocalizedKeywords field if non-nil, zero value otherwise.

### GetLocalizedKeywordsOk

`func (o *GetRequestTemplatesId200Response) GetLocalizedKeywordsOk() (*string, bool)`

GetLocalizedKeywordsOk returns a tuple with the LocalizedKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedKeywords

`func (o *GetRequestTemplatesId200Response) SetLocalizedKeywords(v string)`

SetLocalizedKeywords sets LocalizedKeywords field to given value.

### HasLocalizedKeywords

`func (o *GetRequestTemplatesId200Response) HasLocalizedKeywords() bool`

HasLocalizedKeywords returns a boolean if a field has been set.

### SetLocalizedKeywordsNil

`func (o *GetRequestTemplatesId200Response) SetLocalizedKeywordsNil(b bool)`

 SetLocalizedKeywordsNil sets the value for LocalizedKeywords to be an explicit nil

### UnsetLocalizedKeywords
`func (o *GetRequestTemplatesId200Response) UnsetLocalizedKeywords()`

UnsetLocalizedKeywords ensures that no value is present for LocalizedKeywords, not even an explicit nil
### GetLocalizedNote

`func (o *GetRequestTemplatesId200Response) GetLocalizedNote() string`

GetLocalizedNote returns the LocalizedNote field if non-nil, zero value otherwise.

### GetLocalizedNoteOk

`func (o *GetRequestTemplatesId200Response) GetLocalizedNoteOk() (*string, bool)`

GetLocalizedNoteOk returns a tuple with the LocalizedNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedNote

`func (o *GetRequestTemplatesId200Response) SetLocalizedNote(v string)`

SetLocalizedNote sets LocalizedNote field to given value.

### HasLocalizedNote

`func (o *GetRequestTemplatesId200Response) HasLocalizedNote() bool`

HasLocalizedNote returns a boolean if a field has been set.

### SetLocalizedNoteNil

`func (o *GetRequestTemplatesId200Response) SetLocalizedNoteNil(b bool)`

 SetLocalizedNoteNil sets the value for LocalizedNote to be an explicit nil

### UnsetLocalizedNote
`func (o *GetRequestTemplatesId200Response) UnsetLocalizedNote()`

UnsetLocalizedNote ensures that no value is present for LocalizedNote, not even an explicit nil
### GetLocalizedRegistrationHints

`func (o *GetRequestTemplatesId200Response) GetLocalizedRegistrationHints() string`

GetLocalizedRegistrationHints returns the LocalizedRegistrationHints field if non-nil, zero value otherwise.

### GetLocalizedRegistrationHintsOk

`func (o *GetRequestTemplatesId200Response) GetLocalizedRegistrationHintsOk() (*string, bool)`

GetLocalizedRegistrationHintsOk returns a tuple with the LocalizedRegistrationHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedRegistrationHints

`func (o *GetRequestTemplatesId200Response) SetLocalizedRegistrationHints(v string)`

SetLocalizedRegistrationHints sets LocalizedRegistrationHints field to given value.

### HasLocalizedRegistrationHints

`func (o *GetRequestTemplatesId200Response) HasLocalizedRegistrationHints() bool`

HasLocalizedRegistrationHints returns a boolean if a field has been set.

### SetLocalizedRegistrationHintsNil

`func (o *GetRequestTemplatesId200Response) SetLocalizedRegistrationHintsNil(b bool)`

 SetLocalizedRegistrationHintsNil sets the value for LocalizedRegistrationHints to be an explicit nil

### UnsetLocalizedRegistrationHints
`func (o *GetRequestTemplatesId200Response) UnsetLocalizedRegistrationHints()`

UnsetLocalizedRegistrationHints ensures that no value is present for LocalizedRegistrationHints, not even an explicit nil
### GetLocalizedSubject

`func (o *GetRequestTemplatesId200Response) GetLocalizedSubject() string`

GetLocalizedSubject returns the LocalizedSubject field if non-nil, zero value otherwise.

### GetLocalizedSubjectOk

`func (o *GetRequestTemplatesId200Response) GetLocalizedSubjectOk() (*string, bool)`

GetLocalizedSubjectOk returns a tuple with the LocalizedSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedSubject

`func (o *GetRequestTemplatesId200Response) SetLocalizedSubject(v string)`

SetLocalizedSubject sets LocalizedSubject field to given value.

### HasLocalizedSubject

`func (o *GetRequestTemplatesId200Response) HasLocalizedSubject() bool`

HasLocalizedSubject returns a boolean if a field has been set.

### GetMember

`func (o *GetRequestTemplatesId200Response) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetRequestTemplatesId200Response) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetRequestTemplatesId200Response) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetRequestTemplatesId200Response) HasMember() bool`

HasMember returns a boolean if a field has been set.

### SetMemberNil

`func (o *GetRequestTemplatesId200Response) SetMemberNil(b bool)`

 SetMemberNil sets the value for Member to be an explicit nil

### UnsetMember
`func (o *GetRequestTemplatesId200Response) UnsetMember()`

UnsetMember ensures that no value is present for Member, not even an explicit nil
### GetNodeID

`func (o *GetRequestTemplatesId200Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *GetRequestTemplatesId200Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *GetRequestTemplatesId200Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *GetRequestTemplatesId200Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetNote

`func (o *GetRequestTemplatesId200Response) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *GetRequestTemplatesId200Response) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *GetRequestTemplatesId200Response) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *GetRequestTemplatesId200Response) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *GetRequestTemplatesId200Response) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *GetRequestTemplatesId200Response) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetPlannedEffort

`func (o *GetRequestTemplatesId200Response) GetPlannedEffort() string`

GetPlannedEffort returns the PlannedEffort field if non-nil, zero value otherwise.

### GetPlannedEffortOk

`func (o *GetRequestTemplatesId200Response) GetPlannedEffortOk() (*string, bool)`

GetPlannedEffortOk returns a tuple with the PlannedEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedEffort

`func (o *GetRequestTemplatesId200Response) SetPlannedEffort(v string)`

SetPlannedEffort sets PlannedEffort field to given value.

### HasPlannedEffort

`func (o *GetRequestTemplatesId200Response) HasPlannedEffort() bool`

HasPlannedEffort returns a boolean if a field has been set.

### SetPlannedEffortNil

`func (o *GetRequestTemplatesId200Response) SetPlannedEffortNil(b bool)`

 SetPlannedEffortNil sets the value for PlannedEffort to be an explicit nil

### UnsetPlannedEffort
`func (o *GetRequestTemplatesId200Response) UnsetPlannedEffort()`

UnsetPlannedEffort ensures that no value is present for PlannedEffort, not even an explicit nil
### GetRegistrationHints

`func (o *GetRequestTemplatesId200Response) GetRegistrationHints() string`

GetRegistrationHints returns the RegistrationHints field if non-nil, zero value otherwise.

### GetRegistrationHintsOk

`func (o *GetRequestTemplatesId200Response) GetRegistrationHintsOk() (*string, bool)`

GetRegistrationHintsOk returns a tuple with the RegistrationHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationHints

`func (o *GetRequestTemplatesId200Response) SetRegistrationHints(v string)`

SetRegistrationHints sets RegistrationHints field to given value.

### HasRegistrationHints

`func (o *GetRequestTemplatesId200Response) HasRegistrationHints() bool`

HasRegistrationHints returns a boolean if a field has been set.

### SetRegistrationHintsNil

`func (o *GetRequestTemplatesId200Response) SetRegistrationHintsNil(b bool)`

 SetRegistrationHintsNil sets the value for RegistrationHints to be an explicit nil

### UnsetRegistrationHints
`func (o *GetRequestTemplatesId200Response) UnsetRegistrationHints()`

UnsetRegistrationHints ensures that no value is present for RegistrationHints, not even an explicit nil
### GetResolutionTarget

`func (o *GetRequestTemplatesId200Response) GetResolutionTarget() string`

GetResolutionTarget returns the ResolutionTarget field if non-nil, zero value otherwise.

### GetResolutionTargetOk

`func (o *GetRequestTemplatesId200Response) GetResolutionTargetOk() (*string, bool)`

GetResolutionTargetOk returns a tuple with the ResolutionTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionTarget

`func (o *GetRequestTemplatesId200Response) SetResolutionTarget(v string)`

SetResolutionTarget sets ResolutionTarget field to given value.

### HasResolutionTarget

`func (o *GetRequestTemplatesId200Response) HasResolutionTarget() bool`

HasResolutionTarget returns a boolean if a field has been set.

### SetResolutionTargetNil

`func (o *GetRequestTemplatesId200Response) SetResolutionTargetNil(b bool)`

 SetResolutionTargetNil sets the value for ResolutionTarget to be an explicit nil

### UnsetResolutionTarget
`func (o *GetRequestTemplatesId200Response) UnsetResolutionTarget()`

UnsetResolutionTarget ensures that no value is present for ResolutionTarget, not even an explicit nil
### GetRfcType

`func (o *GetRequestTemplatesId200Response) GetRfcType() string`

GetRfcType returns the RfcType field if non-nil, zero value otherwise.

### GetRfcTypeOk

`func (o *GetRequestTemplatesId200Response) GetRfcTypeOk() (*string, bool)`

GetRfcTypeOk returns a tuple with the RfcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfcType

`func (o *GetRequestTemplatesId200Response) SetRfcType(v string)`

SetRfcType sets RfcType field to given value.

### HasRfcType

`func (o *GetRequestTemplatesId200Response) HasRfcType() bool`

HasRfcType returns a boolean if a field has been set.

### SetRfcTypeNil

`func (o *GetRequestTemplatesId200Response) SetRfcTypeNil(b bool)`

 SetRfcTypeNil sets the value for RfcType to be an explicit nil

### UnsetRfcType
`func (o *GetRequestTemplatesId200Response) UnsetRfcType()`

UnsetRfcType ensures that no value is present for RfcType, not even an explicit nil
### GetService

`func (o *GetRequestTemplatesId200Response) GetService() GetRequestsIdCis200ResponseInnerService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetRequestTemplatesId200Response) GetServiceOk() (*GetRequestsIdCis200ResponseInnerService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetRequestTemplatesId200Response) SetService(v GetRequestsIdCis200ResponseInnerService)`

SetService sets Service field to given value.

### HasService

`func (o *GetRequestTemplatesId200Response) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSource

`func (o *GetRequestTemplatesId200Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetRequestTemplatesId200Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetRequestTemplatesId200Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetRequestTemplatesId200Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *GetRequestTemplatesId200Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *GetRequestTemplatesId200Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *GetRequestTemplatesId200Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *GetRequestTemplatesId200Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *GetRequestTemplatesId200Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *GetRequestTemplatesId200Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetSpecialists

`func (o *GetRequestTemplatesId200Response) GetSpecialists() bool`

GetSpecialists returns the Specialists field if non-nil, zero value otherwise.

### GetSpecialistsOk

`func (o *GetRequestTemplatesId200Response) GetSpecialistsOk() (*bool, bool)`

GetSpecialistsOk returns a tuple with the Specialists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialists

`func (o *GetRequestTemplatesId200Response) SetSpecialists(v bool)`

SetSpecialists sets Specialists field to given value.

### HasSpecialists

`func (o *GetRequestTemplatesId200Response) HasSpecialists() bool`

HasSpecialists returns a boolean if a field has been set.

### GetStatus

`func (o *GetRequestTemplatesId200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRequestTemplatesId200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRequestTemplatesId200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetRequestTemplatesId200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *GetRequestTemplatesId200Response) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetRequestTemplatesId200Response) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetRequestTemplatesId200Response) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetRequestTemplatesId200Response) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSupplier

`func (o *GetRequestTemplatesId200Response) GetSupplier() string`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *GetRequestTemplatesId200Response) GetSupplierOk() (*string, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *GetRequestTemplatesId200Response) SetSupplier(v string)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *GetRequestTemplatesId200Response) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### SetSupplierNil

`func (o *GetRequestTemplatesId200Response) SetSupplierNil(b bool)`

 SetSupplierNil sets the value for Supplier to be an explicit nil

### UnsetSupplier
`func (o *GetRequestTemplatesId200Response) UnsetSupplier()`

UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
### GetSupportHours

`func (o *GetRequestTemplatesId200Response) GetSupportHours() string`

GetSupportHours returns the SupportHours field if non-nil, zero value otherwise.

### GetSupportHoursOk

`func (o *GetRequestTemplatesId200Response) GetSupportHoursOk() (*string, bool)`

GetSupportHoursOk returns a tuple with the SupportHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportHours

`func (o *GetRequestTemplatesId200Response) SetSupportHours(v string)`

SetSupportHours sets SupportHours field to given value.

### HasSupportHours

`func (o *GetRequestTemplatesId200Response) HasSupportHours() bool`

HasSupportHours returns a boolean if a field has been set.

### SetSupportHoursNil

`func (o *GetRequestTemplatesId200Response) SetSupportHoursNil(b bool)`

 SetSupportHoursNil sets the value for SupportHours to be an explicit nil

### UnsetSupportHours
`func (o *GetRequestTemplatesId200Response) UnsetSupportHours()`

UnsetSupportHours ensures that no value is present for SupportHours, not even an explicit nil
### GetTeam

`func (o *GetRequestTemplatesId200Response) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GetRequestTemplatesId200Response) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GetRequestTemplatesId200Response) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GetRequestTemplatesId200Response) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *GetRequestTemplatesId200Response) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *GetRequestTemplatesId200Response) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetTimeZone

`func (o *GetRequestTemplatesId200Response) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetRequestTemplatesId200Response) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetRequestTemplatesId200Response) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetRequestTemplatesId200Response) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *GetRequestTemplatesId200Response) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *GetRequestTemplatesId200Response) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetTimesApplied

`func (o *GetRequestTemplatesId200Response) GetTimesApplied() float32`

GetTimesApplied returns the TimesApplied field if non-nil, zero value otherwise.

### GetTimesAppliedOk

`func (o *GetRequestTemplatesId200Response) GetTimesAppliedOk() (*float32, bool)`

GetTimesAppliedOk returns a tuple with the TimesApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesApplied

`func (o *GetRequestTemplatesId200Response) SetTimesApplied(v float32)`

SetTimesApplied sets TimesApplied field to given value.

### HasTimesApplied

`func (o *GetRequestTemplatesId200Response) HasTimesApplied() bool`

HasTimesApplied returns a boolean if a field has been set.

### GetUiExtension

`func (o *GetRequestTemplatesId200Response) GetUiExtension() GetRequestTemplatesId200ResponseUiExtension`

GetUiExtension returns the UiExtension field if non-nil, zero value otherwise.

### GetUiExtensionOk

`func (o *GetRequestTemplatesId200Response) GetUiExtensionOk() (*GetRequestTemplatesId200ResponseUiExtension, bool)`

GetUiExtensionOk returns a tuple with the UiExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiExtension

`func (o *GetRequestTemplatesId200Response) SetUiExtension(v GetRequestTemplatesId200ResponseUiExtension)`

SetUiExtension sets UiExtension field to given value.

### HasUiExtension

`func (o *GetRequestTemplatesId200Response) HasUiExtension() bool`

HasUiExtension returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetRequestTemplatesId200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetRequestTemplatesId200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetRequestTemplatesId200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetRequestTemplatesId200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrgent

`func (o *GetRequestTemplatesId200Response) GetUrgent() bool`

GetUrgent returns the Urgent field if non-nil, zero value otherwise.

### GetUrgentOk

`func (o *GetRequestTemplatesId200Response) GetUrgentOk() (*bool, bool)`

GetUrgentOk returns a tuple with the Urgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgent

`func (o *GetRequestTemplatesId200Response) SetUrgent(v bool)`

SetUrgent sets Urgent field to given value.

### HasUrgent

`func (o *GetRequestTemplatesId200Response) HasUrgent() bool`

HasUrgent returns a boolean if a field has been set.

### GetWorkflowManager

`func (o *GetRequestTemplatesId200Response) GetWorkflowManager() string`

GetWorkflowManager returns the WorkflowManager field if non-nil, zero value otherwise.

### GetWorkflowManagerOk

`func (o *GetRequestTemplatesId200Response) GetWorkflowManagerOk() (*string, bool)`

GetWorkflowManagerOk returns a tuple with the WorkflowManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowManager

`func (o *GetRequestTemplatesId200Response) SetWorkflowManager(v string)`

SetWorkflowManager sets WorkflowManager field to given value.

### HasWorkflowManager

`func (o *GetRequestTemplatesId200Response) HasWorkflowManager() bool`

HasWorkflowManager returns a boolean if a field has been set.

### SetWorkflowManagerNil

`func (o *GetRequestTemplatesId200Response) SetWorkflowManagerNil(b bool)`

 SetWorkflowManagerNil sets the value for WorkflowManager to be an explicit nil

### UnsetWorkflowManager
`func (o *GetRequestTemplatesId200Response) UnsetWorkflowManager()`

UnsetWorkflowManager ensures that no value is present for WorkflowManager, not even an explicit nil
### GetWorkflowTemplate

`func (o *GetRequestTemplatesId200Response) GetWorkflowTemplate() string`

GetWorkflowTemplate returns the WorkflowTemplate field if non-nil, zero value otherwise.

### GetWorkflowTemplateOk

`func (o *GetRequestTemplatesId200Response) GetWorkflowTemplateOk() (*string, bool)`

GetWorkflowTemplateOk returns a tuple with the WorkflowTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTemplate

`func (o *GetRequestTemplatesId200Response) SetWorkflowTemplate(v string)`

SetWorkflowTemplate sets WorkflowTemplate field to given value.

### HasWorkflowTemplate

`func (o *GetRequestTemplatesId200Response) HasWorkflowTemplate() bool`

HasWorkflowTemplate returns a boolean if a field has been set.

### SetWorkflowTemplateNil

`func (o *GetRequestTemplatesId200Response) SetWorkflowTemplateNil(b bool)`

 SetWorkflowTemplateNil sets the value for WorkflowTemplate to be an explicit nil

### UnsetWorkflowTemplate
`func (o *GetRequestTemplatesId200Response) UnsetWorkflowTemplate()`

UnsetWorkflowTemplate ensures that no value is present for WorkflowTemplate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


