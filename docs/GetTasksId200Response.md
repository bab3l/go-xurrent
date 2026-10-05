# GetTasksId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**AgileBoard** | Pointer to **NullableString** |  | [optional] 
**AgileBoardColumn** | Pointer to **NullableString** |  | [optional] 
**AgileBoardColumnPosition** | Pointer to **NullableString** |  | [optional] 
**AnticipatedAssignmentAt** | Pointer to **NullableString** |  | [optional] 
**AssignedAt** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CheckedItems** | Pointer to **NullableString** |  | [optional] 
**CompletionTargetAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **NullableString** |  | [optional] 
**FailureTask** | Pointer to **NullableString** |  | [optional] 
**FinishedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**Manager** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**Member** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**NewAssignment** | Pointer to **bool** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**PdfDesign** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**Phase** | Pointer to [**GetTasksId200ResponsePhase**](GetTasksId200ResponsePhase.md) |  | [optional] 
**PlannedDuration** | Pointer to **float32** |  | [optional] 
**PlannedEffort** | Pointer to **NullableString** |  | [optional] 
**ProviderNotAccountable** | Pointer to **bool** |  | [optional] 
**RejectionCount** | Pointer to **float32** |  | [optional] 
**Request** | Pointer to **NullableString** |  | [optional] 
**RequestServiceInstance** | Pointer to **NullableString** |  | [optional] 
**RequestTemplate** | Pointer to **NullableString** |  | [optional] 
**RequiredApprovals** | Pointer to **float32** |  | [optional] 
**ResolutionDuration** | Pointer to **float32** |  | [optional] 
**SkillPool** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**StartAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Supplier** | Pointer to **NullableString** |  | [optional] 
**SupplierRequestID** | Pointer to **NullableString** |  | [optional] 
**Team** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**Template** | Pointer to [**GetTasksId200ResponseTemplate**](GetTasksId200ResponseTemplate.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Urgent** | Pointer to **bool** |  | [optional] 
**WaitingUntil** | Pointer to **NullableString** |  | [optional] 
**WorkHoursAre24x7** | Pointer to **bool** |  | [optional] 
**Workflow** | Pointer to [**GetTasksId200ResponseTemplate**](GetTasksId200ResponseTemplate.md) |  | [optional] 

## Methods

### NewGetTasksId200Response

`func NewGetTasksId200Response() *GetTasksId200Response`

NewGetTasksId200Response instantiates a new GetTasksId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTasksId200ResponseWithDefaults

`func NewGetTasksId200ResponseWithDefaults() *GetTasksId200Response`

NewGetTasksId200ResponseWithDefaults instantiates a new GetTasksId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetTasksId200Response) GetAccount() GetRequestsId200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetTasksId200Response) GetAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetTasksId200Response) SetAccount(v GetRequestsId200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetTasksId200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAgileBoard

`func (o *GetTasksId200Response) GetAgileBoard() string`

GetAgileBoard returns the AgileBoard field if non-nil, zero value otherwise.

### GetAgileBoardOk

`func (o *GetTasksId200Response) GetAgileBoardOk() (*string, bool)`

GetAgileBoardOk returns a tuple with the AgileBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgileBoard

`func (o *GetTasksId200Response) SetAgileBoard(v string)`

SetAgileBoard sets AgileBoard field to given value.

### HasAgileBoard

`func (o *GetTasksId200Response) HasAgileBoard() bool`

HasAgileBoard returns a boolean if a field has been set.

### SetAgileBoardNil

`func (o *GetTasksId200Response) SetAgileBoardNil(b bool)`

 SetAgileBoardNil sets the value for AgileBoard to be an explicit nil

### UnsetAgileBoard
`func (o *GetTasksId200Response) UnsetAgileBoard()`

UnsetAgileBoard ensures that no value is present for AgileBoard, not even an explicit nil
### GetAgileBoardColumn

`func (o *GetTasksId200Response) GetAgileBoardColumn() string`

GetAgileBoardColumn returns the AgileBoardColumn field if non-nil, zero value otherwise.

### GetAgileBoardColumnOk

`func (o *GetTasksId200Response) GetAgileBoardColumnOk() (*string, bool)`

GetAgileBoardColumnOk returns a tuple with the AgileBoardColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgileBoardColumn

`func (o *GetTasksId200Response) SetAgileBoardColumn(v string)`

SetAgileBoardColumn sets AgileBoardColumn field to given value.

### HasAgileBoardColumn

`func (o *GetTasksId200Response) HasAgileBoardColumn() bool`

HasAgileBoardColumn returns a boolean if a field has been set.

### SetAgileBoardColumnNil

`func (o *GetTasksId200Response) SetAgileBoardColumnNil(b bool)`

 SetAgileBoardColumnNil sets the value for AgileBoardColumn to be an explicit nil

### UnsetAgileBoardColumn
`func (o *GetTasksId200Response) UnsetAgileBoardColumn()`

UnsetAgileBoardColumn ensures that no value is present for AgileBoardColumn, not even an explicit nil
### GetAgileBoardColumnPosition

`func (o *GetTasksId200Response) GetAgileBoardColumnPosition() string`

GetAgileBoardColumnPosition returns the AgileBoardColumnPosition field if non-nil, zero value otherwise.

### GetAgileBoardColumnPositionOk

`func (o *GetTasksId200Response) GetAgileBoardColumnPositionOk() (*string, bool)`

GetAgileBoardColumnPositionOk returns a tuple with the AgileBoardColumnPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgileBoardColumnPosition

`func (o *GetTasksId200Response) SetAgileBoardColumnPosition(v string)`

SetAgileBoardColumnPosition sets AgileBoardColumnPosition field to given value.

### HasAgileBoardColumnPosition

`func (o *GetTasksId200Response) HasAgileBoardColumnPosition() bool`

HasAgileBoardColumnPosition returns a boolean if a field has been set.

### SetAgileBoardColumnPositionNil

`func (o *GetTasksId200Response) SetAgileBoardColumnPositionNil(b bool)`

 SetAgileBoardColumnPositionNil sets the value for AgileBoardColumnPosition to be an explicit nil

### UnsetAgileBoardColumnPosition
`func (o *GetTasksId200Response) UnsetAgileBoardColumnPosition()`

UnsetAgileBoardColumnPosition ensures that no value is present for AgileBoardColumnPosition, not even an explicit nil
### GetAnticipatedAssignmentAt

`func (o *GetTasksId200Response) GetAnticipatedAssignmentAt() string`

GetAnticipatedAssignmentAt returns the AnticipatedAssignmentAt field if non-nil, zero value otherwise.

### GetAnticipatedAssignmentAtOk

`func (o *GetTasksId200Response) GetAnticipatedAssignmentAtOk() (*string, bool)`

GetAnticipatedAssignmentAtOk returns a tuple with the AnticipatedAssignmentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnticipatedAssignmentAt

`func (o *GetTasksId200Response) SetAnticipatedAssignmentAt(v string)`

SetAnticipatedAssignmentAt sets AnticipatedAssignmentAt field to given value.

### HasAnticipatedAssignmentAt

`func (o *GetTasksId200Response) HasAnticipatedAssignmentAt() bool`

HasAnticipatedAssignmentAt returns a boolean if a field has been set.

### SetAnticipatedAssignmentAtNil

`func (o *GetTasksId200Response) SetAnticipatedAssignmentAtNil(b bool)`

 SetAnticipatedAssignmentAtNil sets the value for AnticipatedAssignmentAt to be an explicit nil

### UnsetAnticipatedAssignmentAt
`func (o *GetTasksId200Response) UnsetAnticipatedAssignmentAt()`

UnsetAnticipatedAssignmentAt ensures that no value is present for AnticipatedAssignmentAt, not even an explicit nil
### GetAssignedAt

`func (o *GetTasksId200Response) GetAssignedAt() string`

GetAssignedAt returns the AssignedAt field if non-nil, zero value otherwise.

### GetAssignedAtOk

`func (o *GetTasksId200Response) GetAssignedAtOk() (*string, bool)`

GetAssignedAtOk returns a tuple with the AssignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedAt

`func (o *GetTasksId200Response) SetAssignedAt(v string)`

SetAssignedAt sets AssignedAt field to given value.

### HasAssignedAt

`func (o *GetTasksId200Response) HasAssignedAt() bool`

HasAssignedAt returns a boolean if a field has been set.

### GetAttachments

`func (o *GetTasksId200Response) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetTasksId200Response) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetTasksId200Response) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetTasksId200Response) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCategory

`func (o *GetTasksId200Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetTasksId200Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetTasksId200Response) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetTasksId200Response) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCheckedItems

`func (o *GetTasksId200Response) GetCheckedItems() string`

GetCheckedItems returns the CheckedItems field if non-nil, zero value otherwise.

### GetCheckedItemsOk

`func (o *GetTasksId200Response) GetCheckedItemsOk() (*string, bool)`

GetCheckedItemsOk returns a tuple with the CheckedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedItems

`func (o *GetTasksId200Response) SetCheckedItems(v string)`

SetCheckedItems sets CheckedItems field to given value.

### HasCheckedItems

`func (o *GetTasksId200Response) HasCheckedItems() bool`

HasCheckedItems returns a boolean if a field has been set.

### SetCheckedItemsNil

`func (o *GetTasksId200Response) SetCheckedItemsNil(b bool)`

 SetCheckedItemsNil sets the value for CheckedItems to be an explicit nil

### UnsetCheckedItems
`func (o *GetTasksId200Response) UnsetCheckedItems()`

UnsetCheckedItems ensures that no value is present for CheckedItems, not even an explicit nil
### GetCompletionTargetAt

`func (o *GetTasksId200Response) GetCompletionTargetAt() string`

GetCompletionTargetAt returns the CompletionTargetAt field if non-nil, zero value otherwise.

### GetCompletionTargetAtOk

`func (o *GetTasksId200Response) GetCompletionTargetAtOk() (*string, bool)`

GetCompletionTargetAtOk returns a tuple with the CompletionTargetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTargetAt

`func (o *GetTasksId200Response) SetCompletionTargetAt(v string)`

SetCompletionTargetAt sets CompletionTargetAt field to given value.

### HasCompletionTargetAt

`func (o *GetTasksId200Response) HasCompletionTargetAt() bool`

HasCompletionTargetAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetTasksId200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetTasksId200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetTasksId200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetTasksId200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomFields

`func (o *GetTasksId200Response) GetCustomFields() string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GetTasksId200Response) GetCustomFieldsOk() (*string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GetTasksId200Response) SetCustomFields(v string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GetTasksId200Response) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *GetTasksId200Response) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *GetTasksId200Response) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetFailureTask

`func (o *GetTasksId200Response) GetFailureTask() string`

GetFailureTask returns the FailureTask field if non-nil, zero value otherwise.

### GetFailureTaskOk

`func (o *GetTasksId200Response) GetFailureTaskOk() (*string, bool)`

GetFailureTaskOk returns a tuple with the FailureTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureTask

`func (o *GetTasksId200Response) SetFailureTask(v string)`

SetFailureTask sets FailureTask field to given value.

### HasFailureTask

`func (o *GetTasksId200Response) HasFailureTask() bool`

HasFailureTask returns a boolean if a field has been set.

### SetFailureTaskNil

`func (o *GetTasksId200Response) SetFailureTaskNil(b bool)`

 SetFailureTaskNil sets the value for FailureTask to be an explicit nil

### UnsetFailureTask
`func (o *GetTasksId200Response) UnsetFailureTask()`

UnsetFailureTask ensures that no value is present for FailureTask, not even an explicit nil
### GetFinishedAt

`func (o *GetTasksId200Response) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *GetTasksId200Response) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *GetTasksId200Response) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *GetTasksId200Response) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *GetTasksId200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTasksId200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTasksId200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetTasksId200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImpact

`func (o *GetTasksId200Response) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *GetTasksId200Response) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *GetTasksId200Response) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *GetTasksId200Response) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetInstructions

`func (o *GetTasksId200Response) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *GetTasksId200Response) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *GetTasksId200Response) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *GetTasksId200Response) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetManager

`func (o *GetTasksId200Response) GetManager() GetRequestsId200ResponseCreatedBy`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *GetTasksId200Response) GetManagerOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *GetTasksId200Response) SetManager(v GetRequestsId200ResponseCreatedBy)`

SetManager sets Manager field to given value.

### HasManager

`func (o *GetTasksId200Response) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetMember

`func (o *GetTasksId200Response) GetMember() GetRequestsId200ResponseCreatedBy`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetTasksId200Response) GetMemberOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetTasksId200Response) SetMember(v GetRequestsId200ResponseCreatedBy)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetTasksId200Response) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetNewAssignment

`func (o *GetTasksId200Response) GetNewAssignment() bool`

GetNewAssignment returns the NewAssignment field if non-nil, zero value otherwise.

### GetNewAssignmentOk

`func (o *GetTasksId200Response) GetNewAssignmentOk() (*bool, bool)`

GetNewAssignmentOk returns a tuple with the NewAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAssignment

`func (o *GetTasksId200Response) SetNewAssignment(v bool)`

SetNewAssignment sets NewAssignment field to given value.

### HasNewAssignment

`func (o *GetTasksId200Response) HasNewAssignment() bool`

HasNewAssignment returns a boolean if a field has been set.

### GetNodeID

`func (o *GetTasksId200Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *GetTasksId200Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *GetTasksId200Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *GetTasksId200Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetPdfDesign

`func (o *GetTasksId200Response) GetPdfDesign() GetRequestsId200ResponseCreatedBy`

GetPdfDesign returns the PdfDesign field if non-nil, zero value otherwise.

### GetPdfDesignOk

`func (o *GetTasksId200Response) GetPdfDesignOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetPdfDesignOk returns a tuple with the PdfDesign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfDesign

`func (o *GetTasksId200Response) SetPdfDesign(v GetRequestsId200ResponseCreatedBy)`

SetPdfDesign sets PdfDesign field to given value.

### HasPdfDesign

`func (o *GetTasksId200Response) HasPdfDesign() bool`

HasPdfDesign returns a boolean if a field has been set.

### GetPhase

`func (o *GetTasksId200Response) GetPhase() GetTasksId200ResponsePhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *GetTasksId200Response) GetPhaseOk() (*GetTasksId200ResponsePhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *GetTasksId200Response) SetPhase(v GetTasksId200ResponsePhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *GetTasksId200Response) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPlannedDuration

`func (o *GetTasksId200Response) GetPlannedDuration() float32`

GetPlannedDuration returns the PlannedDuration field if non-nil, zero value otherwise.

### GetPlannedDurationOk

`func (o *GetTasksId200Response) GetPlannedDurationOk() (*float32, bool)`

GetPlannedDurationOk returns a tuple with the PlannedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDuration

`func (o *GetTasksId200Response) SetPlannedDuration(v float32)`

SetPlannedDuration sets PlannedDuration field to given value.

### HasPlannedDuration

`func (o *GetTasksId200Response) HasPlannedDuration() bool`

HasPlannedDuration returns a boolean if a field has been set.

### GetPlannedEffort

`func (o *GetTasksId200Response) GetPlannedEffort() string`

GetPlannedEffort returns the PlannedEffort field if non-nil, zero value otherwise.

### GetPlannedEffortOk

`func (o *GetTasksId200Response) GetPlannedEffortOk() (*string, bool)`

GetPlannedEffortOk returns a tuple with the PlannedEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedEffort

`func (o *GetTasksId200Response) SetPlannedEffort(v string)`

SetPlannedEffort sets PlannedEffort field to given value.

### HasPlannedEffort

`func (o *GetTasksId200Response) HasPlannedEffort() bool`

HasPlannedEffort returns a boolean if a field has been set.

### SetPlannedEffortNil

`func (o *GetTasksId200Response) SetPlannedEffortNil(b bool)`

 SetPlannedEffortNil sets the value for PlannedEffort to be an explicit nil

### UnsetPlannedEffort
`func (o *GetTasksId200Response) UnsetPlannedEffort()`

UnsetPlannedEffort ensures that no value is present for PlannedEffort, not even an explicit nil
### GetProviderNotAccountable

`func (o *GetTasksId200Response) GetProviderNotAccountable() bool`

GetProviderNotAccountable returns the ProviderNotAccountable field if non-nil, zero value otherwise.

### GetProviderNotAccountableOk

`func (o *GetTasksId200Response) GetProviderNotAccountableOk() (*bool, bool)`

GetProviderNotAccountableOk returns a tuple with the ProviderNotAccountable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNotAccountable

`func (o *GetTasksId200Response) SetProviderNotAccountable(v bool)`

SetProviderNotAccountable sets ProviderNotAccountable field to given value.

### HasProviderNotAccountable

`func (o *GetTasksId200Response) HasProviderNotAccountable() bool`

HasProviderNotAccountable returns a boolean if a field has been set.

### GetRejectionCount

`func (o *GetTasksId200Response) GetRejectionCount() float32`

GetRejectionCount returns the RejectionCount field if non-nil, zero value otherwise.

### GetRejectionCountOk

`func (o *GetTasksId200Response) GetRejectionCountOk() (*float32, bool)`

GetRejectionCountOk returns a tuple with the RejectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionCount

`func (o *GetTasksId200Response) SetRejectionCount(v float32)`

SetRejectionCount sets RejectionCount field to given value.

### HasRejectionCount

`func (o *GetTasksId200Response) HasRejectionCount() bool`

HasRejectionCount returns a boolean if a field has been set.

### GetRequest

`func (o *GetTasksId200Response) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *GetTasksId200Response) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *GetTasksId200Response) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *GetTasksId200Response) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequestNil

`func (o *GetTasksId200Response) SetRequestNil(b bool)`

 SetRequestNil sets the value for Request to be an explicit nil

### UnsetRequest
`func (o *GetTasksId200Response) UnsetRequest()`

UnsetRequest ensures that no value is present for Request, not even an explicit nil
### GetRequestServiceInstance

`func (o *GetTasksId200Response) GetRequestServiceInstance() string`

GetRequestServiceInstance returns the RequestServiceInstance field if non-nil, zero value otherwise.

### GetRequestServiceInstanceOk

`func (o *GetTasksId200Response) GetRequestServiceInstanceOk() (*string, bool)`

GetRequestServiceInstanceOk returns a tuple with the RequestServiceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestServiceInstance

`func (o *GetTasksId200Response) SetRequestServiceInstance(v string)`

SetRequestServiceInstance sets RequestServiceInstance field to given value.

### HasRequestServiceInstance

`func (o *GetTasksId200Response) HasRequestServiceInstance() bool`

HasRequestServiceInstance returns a boolean if a field has been set.

### SetRequestServiceInstanceNil

`func (o *GetTasksId200Response) SetRequestServiceInstanceNil(b bool)`

 SetRequestServiceInstanceNil sets the value for RequestServiceInstance to be an explicit nil

### UnsetRequestServiceInstance
`func (o *GetTasksId200Response) UnsetRequestServiceInstance()`

UnsetRequestServiceInstance ensures that no value is present for RequestServiceInstance, not even an explicit nil
### GetRequestTemplate

`func (o *GetTasksId200Response) GetRequestTemplate() string`

GetRequestTemplate returns the RequestTemplate field if non-nil, zero value otherwise.

### GetRequestTemplateOk

`func (o *GetTasksId200Response) GetRequestTemplateOk() (*string, bool)`

GetRequestTemplateOk returns a tuple with the RequestTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTemplate

`func (o *GetTasksId200Response) SetRequestTemplate(v string)`

SetRequestTemplate sets RequestTemplate field to given value.

### HasRequestTemplate

`func (o *GetTasksId200Response) HasRequestTemplate() bool`

HasRequestTemplate returns a boolean if a field has been set.

### SetRequestTemplateNil

`func (o *GetTasksId200Response) SetRequestTemplateNil(b bool)`

 SetRequestTemplateNil sets the value for RequestTemplate to be an explicit nil

### UnsetRequestTemplate
`func (o *GetTasksId200Response) UnsetRequestTemplate()`

UnsetRequestTemplate ensures that no value is present for RequestTemplate, not even an explicit nil
### GetRequiredApprovals

`func (o *GetTasksId200Response) GetRequiredApprovals() float32`

GetRequiredApprovals returns the RequiredApprovals field if non-nil, zero value otherwise.

### GetRequiredApprovalsOk

`func (o *GetTasksId200Response) GetRequiredApprovalsOk() (*float32, bool)`

GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovals

`func (o *GetTasksId200Response) SetRequiredApprovals(v float32)`

SetRequiredApprovals sets RequiredApprovals field to given value.

### HasRequiredApprovals

`func (o *GetTasksId200Response) HasRequiredApprovals() bool`

HasRequiredApprovals returns a boolean if a field has been set.

### GetResolutionDuration

`func (o *GetTasksId200Response) GetResolutionDuration() float32`

GetResolutionDuration returns the ResolutionDuration field if non-nil, zero value otherwise.

### GetResolutionDurationOk

`func (o *GetTasksId200Response) GetResolutionDurationOk() (*float32, bool)`

GetResolutionDurationOk returns a tuple with the ResolutionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionDuration

`func (o *GetTasksId200Response) SetResolutionDuration(v float32)`

SetResolutionDuration sets ResolutionDuration field to given value.

### HasResolutionDuration

`func (o *GetTasksId200Response) HasResolutionDuration() bool`

HasResolutionDuration returns a boolean if a field has been set.

### GetSkillPool

`func (o *GetTasksId200Response) GetSkillPool() string`

GetSkillPool returns the SkillPool field if non-nil, zero value otherwise.

### GetSkillPoolOk

`func (o *GetTasksId200Response) GetSkillPoolOk() (*string, bool)`

GetSkillPoolOk returns a tuple with the SkillPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillPool

`func (o *GetTasksId200Response) SetSkillPool(v string)`

SetSkillPool sets SkillPool field to given value.

### HasSkillPool

`func (o *GetTasksId200Response) HasSkillPool() bool`

HasSkillPool returns a boolean if a field has been set.

### SetSkillPoolNil

`func (o *GetTasksId200Response) SetSkillPoolNil(b bool)`

 SetSkillPoolNil sets the value for SkillPool to be an explicit nil

### UnsetSkillPool
`func (o *GetTasksId200Response) UnsetSkillPool()`

UnsetSkillPool ensures that no value is present for SkillPool, not even an explicit nil
### GetSource

`func (o *GetTasksId200Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetTasksId200Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetTasksId200Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetTasksId200Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *GetTasksId200Response) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *GetTasksId200Response) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceID

`func (o *GetTasksId200Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *GetTasksId200Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *GetTasksId200Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *GetTasksId200Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *GetTasksId200Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *GetTasksId200Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetStartAt

`func (o *GetTasksId200Response) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *GetTasksId200Response) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *GetTasksId200Response) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *GetTasksId200Response) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetStatus

`func (o *GetTasksId200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTasksId200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTasksId200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTasksId200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *GetTasksId200Response) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetTasksId200Response) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetTasksId200Response) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetTasksId200Response) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSupplier

`func (o *GetTasksId200Response) GetSupplier() string`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *GetTasksId200Response) GetSupplierOk() (*string, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *GetTasksId200Response) SetSupplier(v string)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *GetTasksId200Response) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### SetSupplierNil

`func (o *GetTasksId200Response) SetSupplierNil(b bool)`

 SetSupplierNil sets the value for Supplier to be an explicit nil

### UnsetSupplier
`func (o *GetTasksId200Response) UnsetSupplier()`

UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
### GetSupplierRequestID

`func (o *GetTasksId200Response) GetSupplierRequestID() string`

GetSupplierRequestID returns the SupplierRequestID field if non-nil, zero value otherwise.

### GetSupplierRequestIDOk

`func (o *GetTasksId200Response) GetSupplierRequestIDOk() (*string, bool)`

GetSupplierRequestIDOk returns a tuple with the SupplierRequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierRequestID

`func (o *GetTasksId200Response) SetSupplierRequestID(v string)`

SetSupplierRequestID sets SupplierRequestID field to given value.

### HasSupplierRequestID

`func (o *GetTasksId200Response) HasSupplierRequestID() bool`

HasSupplierRequestID returns a boolean if a field has been set.

### SetSupplierRequestIDNil

`func (o *GetTasksId200Response) SetSupplierRequestIDNil(b bool)`

 SetSupplierRequestIDNil sets the value for SupplierRequestID to be an explicit nil

### UnsetSupplierRequestID
`func (o *GetTasksId200Response) UnsetSupplierRequestID()`

UnsetSupplierRequestID ensures that no value is present for SupplierRequestID, not even an explicit nil
### GetTeam

`func (o *GetTasksId200Response) GetTeam() GetRequestsId200ResponseCreatedBy`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GetTasksId200Response) GetTeamOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GetTasksId200Response) SetTeam(v GetRequestsId200ResponseCreatedBy)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GetTasksId200Response) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetTemplate

`func (o *GetTasksId200Response) GetTemplate() GetTasksId200ResponseTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetTasksId200Response) GetTemplateOk() (*GetTasksId200ResponseTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetTasksId200Response) SetTemplate(v GetTasksId200ResponseTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetTasksId200Response) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetTasksId200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetTasksId200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetTasksId200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetTasksId200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrgent

`func (o *GetTasksId200Response) GetUrgent() bool`

GetUrgent returns the Urgent field if non-nil, zero value otherwise.

### GetUrgentOk

`func (o *GetTasksId200Response) GetUrgentOk() (*bool, bool)`

GetUrgentOk returns a tuple with the Urgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgent

`func (o *GetTasksId200Response) SetUrgent(v bool)`

SetUrgent sets Urgent field to given value.

### HasUrgent

`func (o *GetTasksId200Response) HasUrgent() bool`

HasUrgent returns a boolean if a field has been set.

### GetWaitingUntil

`func (o *GetTasksId200Response) GetWaitingUntil() string`

GetWaitingUntil returns the WaitingUntil field if non-nil, zero value otherwise.

### GetWaitingUntilOk

`func (o *GetTasksId200Response) GetWaitingUntilOk() (*string, bool)`

GetWaitingUntilOk returns a tuple with the WaitingUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingUntil

`func (o *GetTasksId200Response) SetWaitingUntil(v string)`

SetWaitingUntil sets WaitingUntil field to given value.

### HasWaitingUntil

`func (o *GetTasksId200Response) HasWaitingUntil() bool`

HasWaitingUntil returns a boolean if a field has been set.

### SetWaitingUntilNil

`func (o *GetTasksId200Response) SetWaitingUntilNil(b bool)`

 SetWaitingUntilNil sets the value for WaitingUntil to be an explicit nil

### UnsetWaitingUntil
`func (o *GetTasksId200Response) UnsetWaitingUntil()`

UnsetWaitingUntil ensures that no value is present for WaitingUntil, not even an explicit nil
### GetWorkHoursAre24x7

`func (o *GetTasksId200Response) GetWorkHoursAre24x7() bool`

GetWorkHoursAre24x7 returns the WorkHoursAre24x7 field if non-nil, zero value otherwise.

### GetWorkHoursAre24x7Ok

`func (o *GetTasksId200Response) GetWorkHoursAre24x7Ok() (*bool, bool)`

GetWorkHoursAre24x7Ok returns a tuple with the WorkHoursAre24x7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkHoursAre24x7

`func (o *GetTasksId200Response) SetWorkHoursAre24x7(v bool)`

SetWorkHoursAre24x7 sets WorkHoursAre24x7 field to given value.

### HasWorkHoursAre24x7

`func (o *GetTasksId200Response) HasWorkHoursAre24x7() bool`

HasWorkHoursAre24x7 returns a boolean if a field has been set.

### GetWorkflow

`func (o *GetTasksId200Response) GetWorkflow() GetTasksId200ResponseTemplate`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *GetTasksId200Response) GetWorkflowOk() (*GetTasksId200ResponseTemplate, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *GetTasksId200Response) SetWorkflow(v GetTasksId200ResponseTemplate)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *GetTasksId200Response) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


