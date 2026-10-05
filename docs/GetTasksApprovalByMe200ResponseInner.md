# GetTasksApprovalByMe200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CompletionTargetAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**FinishedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Impact** | Pointer to **NullableString** |  | [optional] 
**Member** | Pointer to **NullableString** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to [**GetTasksId200ResponsePhase**](GetTasksId200ResponsePhase.md) |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Team** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewGetTasksApprovalByMe200ResponseInner

`func NewGetTasksApprovalByMe200ResponseInner() *GetTasksApprovalByMe200ResponseInner`

NewGetTasksApprovalByMe200ResponseInner instantiates a new GetTasksApprovalByMe200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTasksApprovalByMe200ResponseInnerWithDefaults

`func NewGetTasksApprovalByMe200ResponseInnerWithDefaults() *GetTasksApprovalByMe200ResponseInner`

NewGetTasksApprovalByMe200ResponseInnerWithDefaults instantiates a new GetTasksApprovalByMe200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetTasksApprovalByMe200ResponseInner) GetAccount() GetRequestsId200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetTasksApprovalByMe200ResponseInner) SetAccount(v GetRequestsId200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetTasksApprovalByMe200ResponseInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCategory

`func (o *GetTasksApprovalByMe200ResponseInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetTasksApprovalByMe200ResponseInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetTasksApprovalByMe200ResponseInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCompletionTargetAt

`func (o *GetTasksApprovalByMe200ResponseInner) GetCompletionTargetAt() string`

GetCompletionTargetAt returns the CompletionTargetAt field if non-nil, zero value otherwise.

### GetCompletionTargetAtOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetCompletionTargetAtOk() (*string, bool)`

GetCompletionTargetAtOk returns a tuple with the CompletionTargetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTargetAt

`func (o *GetTasksApprovalByMe200ResponseInner) SetCompletionTargetAt(v string)`

SetCompletionTargetAt sets CompletionTargetAt field to given value.

### HasCompletionTargetAt

`func (o *GetTasksApprovalByMe200ResponseInner) HasCompletionTargetAt() bool`

HasCompletionTargetAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetTasksApprovalByMe200ResponseInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetTasksApprovalByMe200ResponseInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetTasksApprovalByMe200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *GetTasksApprovalByMe200ResponseInner) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *GetTasksApprovalByMe200ResponseInner) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *GetTasksApprovalByMe200ResponseInner) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *GetTasksApprovalByMe200ResponseInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTasksApprovalByMe200ResponseInner) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetTasksApprovalByMe200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImpact

`func (o *GetTasksApprovalByMe200ResponseInner) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *GetTasksApprovalByMe200ResponseInner) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *GetTasksApprovalByMe200ResponseInner) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### SetImpactNil

`func (o *GetTasksApprovalByMe200ResponseInner) SetImpactNil(b bool)`

 SetImpactNil sets the value for Impact to be an explicit nil

### UnsetImpact
`func (o *GetTasksApprovalByMe200ResponseInner) UnsetImpact()`

UnsetImpact ensures that no value is present for Impact, not even an explicit nil
### GetMember

`func (o *GetTasksApprovalByMe200ResponseInner) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetTasksApprovalByMe200ResponseInner) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetTasksApprovalByMe200ResponseInner) HasMember() bool`

HasMember returns a boolean if a field has been set.

### SetMemberNil

`func (o *GetTasksApprovalByMe200ResponseInner) SetMemberNil(b bool)`

 SetMemberNil sets the value for Member to be an explicit nil

### UnsetMember
`func (o *GetTasksApprovalByMe200ResponseInner) UnsetMember()`

UnsetMember ensures that no value is present for Member, not even an explicit nil
### GetNodeID

`func (o *GetTasksApprovalByMe200ResponseInner) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *GetTasksApprovalByMe200ResponseInner) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *GetTasksApprovalByMe200ResponseInner) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetPhase

`func (o *GetTasksApprovalByMe200ResponseInner) GetPhase() GetTasksId200ResponsePhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetPhaseOk() (*GetTasksId200ResponsePhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *GetTasksApprovalByMe200ResponseInner) SetPhase(v GetTasksId200ResponsePhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *GetTasksApprovalByMe200ResponseInner) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetSourceID

`func (o *GetTasksApprovalByMe200ResponseInner) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *GetTasksApprovalByMe200ResponseInner) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *GetTasksApprovalByMe200ResponseInner) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *GetTasksApprovalByMe200ResponseInner) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *GetTasksApprovalByMe200ResponseInner) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetStatus

`func (o *GetTasksApprovalByMe200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTasksApprovalByMe200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTasksApprovalByMe200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *GetTasksApprovalByMe200ResponseInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetTasksApprovalByMe200ResponseInner) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetTasksApprovalByMe200ResponseInner) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTeam

`func (o *GetTasksApprovalByMe200ResponseInner) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GetTasksApprovalByMe200ResponseInner) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GetTasksApprovalByMe200ResponseInner) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *GetTasksApprovalByMe200ResponseInner) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *GetTasksApprovalByMe200ResponseInner) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetUpdatedAt

`func (o *GetTasksApprovalByMe200ResponseInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetTasksApprovalByMe200ResponseInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetTasksApprovalByMe200ResponseInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetTasksApprovalByMe200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


