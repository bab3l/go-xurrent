# GetWorkflowsId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**ActualEffort** | Pointer to **float32** |  | [optional] 
**ActualVsPlannedEffortPercentage** | Pointer to **float32** |  | [optional] 
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CompletedAt** | Pointer to **NullableString** |  | [optional] 
**CompletionReason** | Pointer to **NullableString** |  | [optional] 
**CompletionTargetAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**Justification** | Pointer to **string** |  | [optional] 
**Manager** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**PlannedEffort** | Pointer to **float32** |  | [optional] 
**PreventRequestCompletion** | Pointer to **bool** |  | [optional] 
**Project** | Pointer to **NullableString** |  | [optional] 
**Release** | Pointer to **NullableString** |  | [optional] 
**ResolutionDuration** | Pointer to **NullableString** |  | [optional] 
**Service** | Pointer to [**GetRequestsIdCis200ResponseInnerService**](GetRequestsIdCis200ResponseInnerService.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**StartAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Template** | Pointer to [**GetTasksId200ResponseTemplate**](GetTasksId200ResponseTemplate.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**WorkflowType** | Pointer to **string** |  | [optional] 

## Methods

### NewGetWorkflowsId200Response

`func NewGetWorkflowsId200Response() *GetWorkflowsId200Response`

NewGetWorkflowsId200Response instantiates a new GetWorkflowsId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWorkflowsId200ResponseWithDefaults

`func NewGetWorkflowsId200ResponseWithDefaults() *GetWorkflowsId200Response`

NewGetWorkflowsId200ResponseWithDefaults instantiates a new GetWorkflowsId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetWorkflowsId200Response) GetAccount() GetRequestsId200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetWorkflowsId200Response) GetAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetWorkflowsId200Response) SetAccount(v GetRequestsId200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetWorkflowsId200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActualEffort

`func (o *GetWorkflowsId200Response) GetActualEffort() float32`

GetActualEffort returns the ActualEffort field if non-nil, zero value otherwise.

### GetActualEffortOk

`func (o *GetWorkflowsId200Response) GetActualEffortOk() (*float32, bool)`

GetActualEffortOk returns a tuple with the ActualEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualEffort

`func (o *GetWorkflowsId200Response) SetActualEffort(v float32)`

SetActualEffort sets ActualEffort field to given value.

### HasActualEffort

`func (o *GetWorkflowsId200Response) HasActualEffort() bool`

HasActualEffort returns a boolean if a field has been set.

### GetActualVsPlannedEffortPercentage

`func (o *GetWorkflowsId200Response) GetActualVsPlannedEffortPercentage() float32`

GetActualVsPlannedEffortPercentage returns the ActualVsPlannedEffortPercentage field if non-nil, zero value otherwise.

### GetActualVsPlannedEffortPercentageOk

`func (o *GetWorkflowsId200Response) GetActualVsPlannedEffortPercentageOk() (*float32, bool)`

GetActualVsPlannedEffortPercentageOk returns a tuple with the ActualVsPlannedEffortPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualVsPlannedEffortPercentage

`func (o *GetWorkflowsId200Response) SetActualVsPlannedEffortPercentage(v float32)`

SetActualVsPlannedEffortPercentage sets ActualVsPlannedEffortPercentage field to given value.

### HasActualVsPlannedEffortPercentage

`func (o *GetWorkflowsId200Response) HasActualVsPlannedEffortPercentage() bool`

HasActualVsPlannedEffortPercentage returns a boolean if a field has been set.

### GetAttachments

`func (o *GetWorkflowsId200Response) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetWorkflowsId200Response) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetWorkflowsId200Response) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetWorkflowsId200Response) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCategory

`func (o *GetWorkflowsId200Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetWorkflowsId200Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetWorkflowsId200Response) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetWorkflowsId200Response) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCompletedAt

`func (o *GetWorkflowsId200Response) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetWorkflowsId200Response) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetWorkflowsId200Response) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetWorkflowsId200Response) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GetWorkflowsId200Response) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GetWorkflowsId200Response) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCompletionReason

`func (o *GetWorkflowsId200Response) GetCompletionReason() string`

GetCompletionReason returns the CompletionReason field if non-nil, zero value otherwise.

### GetCompletionReasonOk

`func (o *GetWorkflowsId200Response) GetCompletionReasonOk() (*string, bool)`

GetCompletionReasonOk returns a tuple with the CompletionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionReason

`func (o *GetWorkflowsId200Response) SetCompletionReason(v string)`

SetCompletionReason sets CompletionReason field to given value.

### HasCompletionReason

`func (o *GetWorkflowsId200Response) HasCompletionReason() bool`

HasCompletionReason returns a boolean if a field has been set.

### SetCompletionReasonNil

`func (o *GetWorkflowsId200Response) SetCompletionReasonNil(b bool)`

 SetCompletionReasonNil sets the value for CompletionReason to be an explicit nil

### UnsetCompletionReason
`func (o *GetWorkflowsId200Response) UnsetCompletionReason()`

UnsetCompletionReason ensures that no value is present for CompletionReason, not even an explicit nil
### GetCompletionTargetAt

`func (o *GetWorkflowsId200Response) GetCompletionTargetAt() string`

GetCompletionTargetAt returns the CompletionTargetAt field if non-nil, zero value otherwise.

### GetCompletionTargetAtOk

`func (o *GetWorkflowsId200Response) GetCompletionTargetAtOk() (*string, bool)`

GetCompletionTargetAtOk returns a tuple with the CompletionTargetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTargetAt

`func (o *GetWorkflowsId200Response) SetCompletionTargetAt(v string)`

SetCompletionTargetAt sets CompletionTargetAt field to given value.

### HasCompletionTargetAt

`func (o *GetWorkflowsId200Response) HasCompletionTargetAt() bool`

HasCompletionTargetAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetWorkflowsId200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetWorkflowsId200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetWorkflowsId200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetWorkflowsId200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomFields

`func (o *GetWorkflowsId200Response) GetCustomFields() string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GetWorkflowsId200Response) GetCustomFieldsOk() (*string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GetWorkflowsId200Response) SetCustomFields(v string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GetWorkflowsId200Response) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *GetWorkflowsId200Response) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *GetWorkflowsId200Response) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetId

`func (o *GetWorkflowsId200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetWorkflowsId200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetWorkflowsId200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetWorkflowsId200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImpact

`func (o *GetWorkflowsId200Response) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *GetWorkflowsId200Response) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *GetWorkflowsId200Response) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *GetWorkflowsId200Response) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetJustification

`func (o *GetWorkflowsId200Response) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *GetWorkflowsId200Response) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *GetWorkflowsId200Response) SetJustification(v string)`

SetJustification sets Justification field to given value.

### HasJustification

`func (o *GetWorkflowsId200Response) HasJustification() bool`

HasJustification returns a boolean if a field has been set.

### GetManager

`func (o *GetWorkflowsId200Response) GetManager() GetRequestsId200ResponseCreatedBy`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *GetWorkflowsId200Response) GetManagerOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *GetWorkflowsId200Response) SetManager(v GetRequestsId200ResponseCreatedBy)`

SetManager sets Manager field to given value.

### HasManager

`func (o *GetWorkflowsId200Response) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetNodeID

`func (o *GetWorkflowsId200Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *GetWorkflowsId200Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *GetWorkflowsId200Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *GetWorkflowsId200Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetPlannedEffort

`func (o *GetWorkflowsId200Response) GetPlannedEffort() float32`

GetPlannedEffort returns the PlannedEffort field if non-nil, zero value otherwise.

### GetPlannedEffortOk

`func (o *GetWorkflowsId200Response) GetPlannedEffortOk() (*float32, bool)`

GetPlannedEffortOk returns a tuple with the PlannedEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedEffort

`func (o *GetWorkflowsId200Response) SetPlannedEffort(v float32)`

SetPlannedEffort sets PlannedEffort field to given value.

### HasPlannedEffort

`func (o *GetWorkflowsId200Response) HasPlannedEffort() bool`

HasPlannedEffort returns a boolean if a field has been set.

### GetPreventRequestCompletion

`func (o *GetWorkflowsId200Response) GetPreventRequestCompletion() bool`

GetPreventRequestCompletion returns the PreventRequestCompletion field if non-nil, zero value otherwise.

### GetPreventRequestCompletionOk

`func (o *GetWorkflowsId200Response) GetPreventRequestCompletionOk() (*bool, bool)`

GetPreventRequestCompletionOk returns a tuple with the PreventRequestCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventRequestCompletion

`func (o *GetWorkflowsId200Response) SetPreventRequestCompletion(v bool)`

SetPreventRequestCompletion sets PreventRequestCompletion field to given value.

### HasPreventRequestCompletion

`func (o *GetWorkflowsId200Response) HasPreventRequestCompletion() bool`

HasPreventRequestCompletion returns a boolean if a field has been set.

### GetProject

`func (o *GetWorkflowsId200Response) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GetWorkflowsId200Response) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GetWorkflowsId200Response) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *GetWorkflowsId200Response) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *GetWorkflowsId200Response) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *GetWorkflowsId200Response) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetRelease

`func (o *GetWorkflowsId200Response) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *GetWorkflowsId200Response) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *GetWorkflowsId200Response) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *GetWorkflowsId200Response) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### SetReleaseNil

`func (o *GetWorkflowsId200Response) SetReleaseNil(b bool)`

 SetReleaseNil sets the value for Release to be an explicit nil

### UnsetRelease
`func (o *GetWorkflowsId200Response) UnsetRelease()`

UnsetRelease ensures that no value is present for Release, not even an explicit nil
### GetResolutionDuration

`func (o *GetWorkflowsId200Response) GetResolutionDuration() string`

GetResolutionDuration returns the ResolutionDuration field if non-nil, zero value otherwise.

### GetResolutionDurationOk

`func (o *GetWorkflowsId200Response) GetResolutionDurationOk() (*string, bool)`

GetResolutionDurationOk returns a tuple with the ResolutionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionDuration

`func (o *GetWorkflowsId200Response) SetResolutionDuration(v string)`

SetResolutionDuration sets ResolutionDuration field to given value.

### HasResolutionDuration

`func (o *GetWorkflowsId200Response) HasResolutionDuration() bool`

HasResolutionDuration returns a boolean if a field has been set.

### SetResolutionDurationNil

`func (o *GetWorkflowsId200Response) SetResolutionDurationNil(b bool)`

 SetResolutionDurationNil sets the value for ResolutionDuration to be an explicit nil

### UnsetResolutionDuration
`func (o *GetWorkflowsId200Response) UnsetResolutionDuration()`

UnsetResolutionDuration ensures that no value is present for ResolutionDuration, not even an explicit nil
### GetService

`func (o *GetWorkflowsId200Response) GetService() GetRequestsIdCis200ResponseInnerService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetWorkflowsId200Response) GetServiceOk() (*GetRequestsIdCis200ResponseInnerService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetWorkflowsId200Response) SetService(v GetRequestsIdCis200ResponseInnerService)`

SetService sets Service field to given value.

### HasService

`func (o *GetWorkflowsId200Response) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSource

`func (o *GetWorkflowsId200Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetWorkflowsId200Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetWorkflowsId200Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetWorkflowsId200Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *GetWorkflowsId200Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *GetWorkflowsId200Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *GetWorkflowsId200Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *GetWorkflowsId200Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *GetWorkflowsId200Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *GetWorkflowsId200Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetStartAt

`func (o *GetWorkflowsId200Response) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *GetWorkflowsId200Response) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *GetWorkflowsId200Response) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *GetWorkflowsId200Response) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetStatus

`func (o *GetWorkflowsId200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetWorkflowsId200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetWorkflowsId200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetWorkflowsId200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *GetWorkflowsId200Response) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetWorkflowsId200Response) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetWorkflowsId200Response) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetWorkflowsId200Response) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTemplate

`func (o *GetWorkflowsId200Response) GetTemplate() GetTasksId200ResponseTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetWorkflowsId200Response) GetTemplateOk() (*GetTasksId200ResponseTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetWorkflowsId200Response) SetTemplate(v GetTasksId200ResponseTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetWorkflowsId200Response) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetWorkflowsId200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetWorkflowsId200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetWorkflowsId200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetWorkflowsId200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWorkflowType

`func (o *GetWorkflowsId200Response) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *GetWorkflowsId200Response) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *GetWorkflowsId200Response) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *GetWorkflowsId200Response) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


