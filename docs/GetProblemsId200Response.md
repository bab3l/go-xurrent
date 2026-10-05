# GetProblemsId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**AgileBoard** | Pointer to **NullableString** |  | [optional] 
**AgileBoardColumn** | Pointer to **NullableString** |  | [optional] 
**AgileBoardColumnPosition** | Pointer to **NullableString** |  | [optional] 
**AnalysisTargetAt** | Pointer to **NullableString** |  | [optional] 
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to [**[]GetRequestsId200ResponseCustomFieldsInner**](GetRequestsId200ResponseCustomFieldsInner.md) |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**KnowledgeArticle** | Pointer to **NullableString** |  | [optional] 
**KnownError** | Pointer to **bool** |  | [optional] 
**Manager** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**Member** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**NewAssignment** | Pointer to **bool** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**PlannedEffort** | Pointer to **NullableString** |  | [optional] 
**ProductBacklog** | Pointer to **NullableString** |  | [optional] 
**ProductBacklogEstimate** | Pointer to **NullableString** |  | [optional] 
**ProductBacklogPosition** | Pointer to **NullableString** |  | [optional] 
**Project** | Pointer to **NullableString** |  | [optional] 
**ResolutionDuration** | Pointer to **NullableString** |  | [optional] 
**Service** | Pointer to [**GetRequestsIdCis200ResponseInnerService**](GetRequestsIdCis200ResponseInnerService.md) |  | [optional] 
**SolvedAt** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Supplier** | Pointer to **NullableString** |  | [optional] 
**SupplierRequestID** | Pointer to **NullableString** |  | [optional] 
**Team** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**UiExtension** | Pointer to [**GetPeopleId200ResponseUiExtension**](GetPeopleId200ResponseUiExtension.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Urgent** | Pointer to **bool** |  | [optional] 
**WaitingUntil** | Pointer to **NullableString** |  | [optional] 
**Workaround** | Pointer to **string** |  | [optional] 
**Workflow** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetProblemsId200Response

`func NewGetProblemsId200Response() *GetProblemsId200Response`

NewGetProblemsId200Response instantiates a new GetProblemsId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProblemsId200ResponseWithDefaults

`func NewGetProblemsId200ResponseWithDefaults() *GetProblemsId200Response`

NewGetProblemsId200ResponseWithDefaults instantiates a new GetProblemsId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetProblemsId200Response) GetAccount() GetRequestsId200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetProblemsId200Response) GetAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetProblemsId200Response) SetAccount(v GetRequestsId200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetProblemsId200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAgileBoard

`func (o *GetProblemsId200Response) GetAgileBoard() string`

GetAgileBoard returns the AgileBoard field if non-nil, zero value otherwise.

### GetAgileBoardOk

`func (o *GetProblemsId200Response) GetAgileBoardOk() (*string, bool)`

GetAgileBoardOk returns a tuple with the AgileBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgileBoard

`func (o *GetProblemsId200Response) SetAgileBoard(v string)`

SetAgileBoard sets AgileBoard field to given value.

### HasAgileBoard

`func (o *GetProblemsId200Response) HasAgileBoard() bool`

HasAgileBoard returns a boolean if a field has been set.

### SetAgileBoardNil

`func (o *GetProblemsId200Response) SetAgileBoardNil(b bool)`

 SetAgileBoardNil sets the value for AgileBoard to be an explicit nil

### UnsetAgileBoard
`func (o *GetProblemsId200Response) UnsetAgileBoard()`

UnsetAgileBoard ensures that no value is present for AgileBoard, not even an explicit nil
### GetAgileBoardColumn

`func (o *GetProblemsId200Response) GetAgileBoardColumn() string`

GetAgileBoardColumn returns the AgileBoardColumn field if non-nil, zero value otherwise.

### GetAgileBoardColumnOk

`func (o *GetProblemsId200Response) GetAgileBoardColumnOk() (*string, bool)`

GetAgileBoardColumnOk returns a tuple with the AgileBoardColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgileBoardColumn

`func (o *GetProblemsId200Response) SetAgileBoardColumn(v string)`

SetAgileBoardColumn sets AgileBoardColumn field to given value.

### HasAgileBoardColumn

`func (o *GetProblemsId200Response) HasAgileBoardColumn() bool`

HasAgileBoardColumn returns a boolean if a field has been set.

### SetAgileBoardColumnNil

`func (o *GetProblemsId200Response) SetAgileBoardColumnNil(b bool)`

 SetAgileBoardColumnNil sets the value for AgileBoardColumn to be an explicit nil

### UnsetAgileBoardColumn
`func (o *GetProblemsId200Response) UnsetAgileBoardColumn()`

UnsetAgileBoardColumn ensures that no value is present for AgileBoardColumn, not even an explicit nil
### GetAgileBoardColumnPosition

`func (o *GetProblemsId200Response) GetAgileBoardColumnPosition() string`

GetAgileBoardColumnPosition returns the AgileBoardColumnPosition field if non-nil, zero value otherwise.

### GetAgileBoardColumnPositionOk

`func (o *GetProblemsId200Response) GetAgileBoardColumnPositionOk() (*string, bool)`

GetAgileBoardColumnPositionOk returns a tuple with the AgileBoardColumnPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgileBoardColumnPosition

`func (o *GetProblemsId200Response) SetAgileBoardColumnPosition(v string)`

SetAgileBoardColumnPosition sets AgileBoardColumnPosition field to given value.

### HasAgileBoardColumnPosition

`func (o *GetProblemsId200Response) HasAgileBoardColumnPosition() bool`

HasAgileBoardColumnPosition returns a boolean if a field has been set.

### SetAgileBoardColumnPositionNil

`func (o *GetProblemsId200Response) SetAgileBoardColumnPositionNil(b bool)`

 SetAgileBoardColumnPositionNil sets the value for AgileBoardColumnPosition to be an explicit nil

### UnsetAgileBoardColumnPosition
`func (o *GetProblemsId200Response) UnsetAgileBoardColumnPosition()`

UnsetAgileBoardColumnPosition ensures that no value is present for AgileBoardColumnPosition, not even an explicit nil
### GetAnalysisTargetAt

`func (o *GetProblemsId200Response) GetAnalysisTargetAt() string`

GetAnalysisTargetAt returns the AnalysisTargetAt field if non-nil, zero value otherwise.

### GetAnalysisTargetAtOk

`func (o *GetProblemsId200Response) GetAnalysisTargetAtOk() (*string, bool)`

GetAnalysisTargetAtOk returns a tuple with the AnalysisTargetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisTargetAt

`func (o *GetProblemsId200Response) SetAnalysisTargetAt(v string)`

SetAnalysisTargetAt sets AnalysisTargetAt field to given value.

### HasAnalysisTargetAt

`func (o *GetProblemsId200Response) HasAnalysisTargetAt() bool`

HasAnalysisTargetAt returns a boolean if a field has been set.

### SetAnalysisTargetAtNil

`func (o *GetProblemsId200Response) SetAnalysisTargetAtNil(b bool)`

 SetAnalysisTargetAtNil sets the value for AnalysisTargetAt to be an explicit nil

### UnsetAnalysisTargetAt
`func (o *GetProblemsId200Response) UnsetAnalysisTargetAt()`

UnsetAnalysisTargetAt ensures that no value is present for AnalysisTargetAt, not even an explicit nil
### GetAttachments

`func (o *GetProblemsId200Response) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetProblemsId200Response) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetProblemsId200Response) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetProblemsId200Response) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCategory

`func (o *GetProblemsId200Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetProblemsId200Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetProblemsId200Response) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetProblemsId200Response) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetProblemsId200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetProblemsId200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetProblemsId200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetProblemsId200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomFields

`func (o *GetProblemsId200Response) GetCustomFields() []GetRequestsId200ResponseCustomFieldsInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GetProblemsId200Response) GetCustomFieldsOk() (*[]GetRequestsId200ResponseCustomFieldsInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GetProblemsId200Response) SetCustomFields(v []GetRequestsId200ResponseCustomFieldsInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GetProblemsId200Response) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetId

`func (o *GetProblemsId200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetProblemsId200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetProblemsId200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetProblemsId200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImpact

`func (o *GetProblemsId200Response) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *GetProblemsId200Response) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *GetProblemsId200Response) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *GetProblemsId200Response) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetKnowledgeArticle

`func (o *GetProblemsId200Response) GetKnowledgeArticle() string`

GetKnowledgeArticle returns the KnowledgeArticle field if non-nil, zero value otherwise.

### GetKnowledgeArticleOk

`func (o *GetProblemsId200Response) GetKnowledgeArticleOk() (*string, bool)`

GetKnowledgeArticleOk returns a tuple with the KnowledgeArticle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeArticle

`func (o *GetProblemsId200Response) SetKnowledgeArticle(v string)`

SetKnowledgeArticle sets KnowledgeArticle field to given value.

### HasKnowledgeArticle

`func (o *GetProblemsId200Response) HasKnowledgeArticle() bool`

HasKnowledgeArticle returns a boolean if a field has been set.

### SetKnowledgeArticleNil

`func (o *GetProblemsId200Response) SetKnowledgeArticleNil(b bool)`

 SetKnowledgeArticleNil sets the value for KnowledgeArticle to be an explicit nil

### UnsetKnowledgeArticle
`func (o *GetProblemsId200Response) UnsetKnowledgeArticle()`

UnsetKnowledgeArticle ensures that no value is present for KnowledgeArticle, not even an explicit nil
### GetKnownError

`func (o *GetProblemsId200Response) GetKnownError() bool`

GetKnownError returns the KnownError field if non-nil, zero value otherwise.

### GetKnownErrorOk

`func (o *GetProblemsId200Response) GetKnownErrorOk() (*bool, bool)`

GetKnownErrorOk returns a tuple with the KnownError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownError

`func (o *GetProblemsId200Response) SetKnownError(v bool)`

SetKnownError sets KnownError field to given value.

### HasKnownError

`func (o *GetProblemsId200Response) HasKnownError() bool`

HasKnownError returns a boolean if a field has been set.

### GetManager

`func (o *GetProblemsId200Response) GetManager() GetRequestsId200ResponseCreatedBy`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *GetProblemsId200Response) GetManagerOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *GetProblemsId200Response) SetManager(v GetRequestsId200ResponseCreatedBy)`

SetManager sets Manager field to given value.

### HasManager

`func (o *GetProblemsId200Response) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetMember

`func (o *GetProblemsId200Response) GetMember() GetRequestsId200ResponseCreatedBy`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetProblemsId200Response) GetMemberOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetProblemsId200Response) SetMember(v GetRequestsId200ResponseCreatedBy)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetProblemsId200Response) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetNewAssignment

`func (o *GetProblemsId200Response) GetNewAssignment() bool`

GetNewAssignment returns the NewAssignment field if non-nil, zero value otherwise.

### GetNewAssignmentOk

`func (o *GetProblemsId200Response) GetNewAssignmentOk() (*bool, bool)`

GetNewAssignmentOk returns a tuple with the NewAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAssignment

`func (o *GetProblemsId200Response) SetNewAssignment(v bool)`

SetNewAssignment sets NewAssignment field to given value.

### HasNewAssignment

`func (o *GetProblemsId200Response) HasNewAssignment() bool`

HasNewAssignment returns a boolean if a field has been set.

### GetNodeID

`func (o *GetProblemsId200Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *GetProblemsId200Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *GetProblemsId200Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *GetProblemsId200Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetPlannedEffort

`func (o *GetProblemsId200Response) GetPlannedEffort() string`

GetPlannedEffort returns the PlannedEffort field if non-nil, zero value otherwise.

### GetPlannedEffortOk

`func (o *GetProblemsId200Response) GetPlannedEffortOk() (*string, bool)`

GetPlannedEffortOk returns a tuple with the PlannedEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedEffort

`func (o *GetProblemsId200Response) SetPlannedEffort(v string)`

SetPlannedEffort sets PlannedEffort field to given value.

### HasPlannedEffort

`func (o *GetProblemsId200Response) HasPlannedEffort() bool`

HasPlannedEffort returns a boolean if a field has been set.

### SetPlannedEffortNil

`func (o *GetProblemsId200Response) SetPlannedEffortNil(b bool)`

 SetPlannedEffortNil sets the value for PlannedEffort to be an explicit nil

### UnsetPlannedEffort
`func (o *GetProblemsId200Response) UnsetPlannedEffort()`

UnsetPlannedEffort ensures that no value is present for PlannedEffort, not even an explicit nil
### GetProductBacklog

`func (o *GetProblemsId200Response) GetProductBacklog() string`

GetProductBacklog returns the ProductBacklog field if non-nil, zero value otherwise.

### GetProductBacklogOk

`func (o *GetProblemsId200Response) GetProductBacklogOk() (*string, bool)`

GetProductBacklogOk returns a tuple with the ProductBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBacklog

`func (o *GetProblemsId200Response) SetProductBacklog(v string)`

SetProductBacklog sets ProductBacklog field to given value.

### HasProductBacklog

`func (o *GetProblemsId200Response) HasProductBacklog() bool`

HasProductBacklog returns a boolean if a field has been set.

### SetProductBacklogNil

`func (o *GetProblemsId200Response) SetProductBacklogNil(b bool)`

 SetProductBacklogNil sets the value for ProductBacklog to be an explicit nil

### UnsetProductBacklog
`func (o *GetProblemsId200Response) UnsetProductBacklog()`

UnsetProductBacklog ensures that no value is present for ProductBacklog, not even an explicit nil
### GetProductBacklogEstimate

`func (o *GetProblemsId200Response) GetProductBacklogEstimate() string`

GetProductBacklogEstimate returns the ProductBacklogEstimate field if non-nil, zero value otherwise.

### GetProductBacklogEstimateOk

`func (o *GetProblemsId200Response) GetProductBacklogEstimateOk() (*string, bool)`

GetProductBacklogEstimateOk returns a tuple with the ProductBacklogEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBacklogEstimate

`func (o *GetProblemsId200Response) SetProductBacklogEstimate(v string)`

SetProductBacklogEstimate sets ProductBacklogEstimate field to given value.

### HasProductBacklogEstimate

`func (o *GetProblemsId200Response) HasProductBacklogEstimate() bool`

HasProductBacklogEstimate returns a boolean if a field has been set.

### SetProductBacklogEstimateNil

`func (o *GetProblemsId200Response) SetProductBacklogEstimateNil(b bool)`

 SetProductBacklogEstimateNil sets the value for ProductBacklogEstimate to be an explicit nil

### UnsetProductBacklogEstimate
`func (o *GetProblemsId200Response) UnsetProductBacklogEstimate()`

UnsetProductBacklogEstimate ensures that no value is present for ProductBacklogEstimate, not even an explicit nil
### GetProductBacklogPosition

`func (o *GetProblemsId200Response) GetProductBacklogPosition() string`

GetProductBacklogPosition returns the ProductBacklogPosition field if non-nil, zero value otherwise.

### GetProductBacklogPositionOk

`func (o *GetProblemsId200Response) GetProductBacklogPositionOk() (*string, bool)`

GetProductBacklogPositionOk returns a tuple with the ProductBacklogPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBacklogPosition

`func (o *GetProblemsId200Response) SetProductBacklogPosition(v string)`

SetProductBacklogPosition sets ProductBacklogPosition field to given value.

### HasProductBacklogPosition

`func (o *GetProblemsId200Response) HasProductBacklogPosition() bool`

HasProductBacklogPosition returns a boolean if a field has been set.

### SetProductBacklogPositionNil

`func (o *GetProblemsId200Response) SetProductBacklogPositionNil(b bool)`

 SetProductBacklogPositionNil sets the value for ProductBacklogPosition to be an explicit nil

### UnsetProductBacklogPosition
`func (o *GetProblemsId200Response) UnsetProductBacklogPosition()`

UnsetProductBacklogPosition ensures that no value is present for ProductBacklogPosition, not even an explicit nil
### GetProject

`func (o *GetProblemsId200Response) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GetProblemsId200Response) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GetProblemsId200Response) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *GetProblemsId200Response) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *GetProblemsId200Response) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *GetProblemsId200Response) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetResolutionDuration

`func (o *GetProblemsId200Response) GetResolutionDuration() string`

GetResolutionDuration returns the ResolutionDuration field if non-nil, zero value otherwise.

### GetResolutionDurationOk

`func (o *GetProblemsId200Response) GetResolutionDurationOk() (*string, bool)`

GetResolutionDurationOk returns a tuple with the ResolutionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionDuration

`func (o *GetProblemsId200Response) SetResolutionDuration(v string)`

SetResolutionDuration sets ResolutionDuration field to given value.

### HasResolutionDuration

`func (o *GetProblemsId200Response) HasResolutionDuration() bool`

HasResolutionDuration returns a boolean if a field has been set.

### SetResolutionDurationNil

`func (o *GetProblemsId200Response) SetResolutionDurationNil(b bool)`

 SetResolutionDurationNil sets the value for ResolutionDuration to be an explicit nil

### UnsetResolutionDuration
`func (o *GetProblemsId200Response) UnsetResolutionDuration()`

UnsetResolutionDuration ensures that no value is present for ResolutionDuration, not even an explicit nil
### GetService

`func (o *GetProblemsId200Response) GetService() GetRequestsIdCis200ResponseInnerService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetProblemsId200Response) GetServiceOk() (*GetRequestsIdCis200ResponseInnerService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetProblemsId200Response) SetService(v GetRequestsIdCis200ResponseInnerService)`

SetService sets Service field to given value.

### HasService

`func (o *GetProblemsId200Response) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSolvedAt

`func (o *GetProblemsId200Response) GetSolvedAt() string`

GetSolvedAt returns the SolvedAt field if non-nil, zero value otherwise.

### GetSolvedAtOk

`func (o *GetProblemsId200Response) GetSolvedAtOk() (*string, bool)`

GetSolvedAtOk returns a tuple with the SolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolvedAt

`func (o *GetProblemsId200Response) SetSolvedAt(v string)`

SetSolvedAt sets SolvedAt field to given value.

### HasSolvedAt

`func (o *GetProblemsId200Response) HasSolvedAt() bool`

HasSolvedAt returns a boolean if a field has been set.

### SetSolvedAtNil

`func (o *GetProblemsId200Response) SetSolvedAtNil(b bool)`

 SetSolvedAtNil sets the value for SolvedAt to be an explicit nil

### UnsetSolvedAt
`func (o *GetProblemsId200Response) UnsetSolvedAt()`

UnsetSolvedAt ensures that no value is present for SolvedAt, not even an explicit nil
### GetSource

`func (o *GetProblemsId200Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetProblemsId200Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetProblemsId200Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetProblemsId200Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *GetProblemsId200Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *GetProblemsId200Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *GetProblemsId200Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *GetProblemsId200Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *GetProblemsId200Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *GetProblemsId200Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetStatus

`func (o *GetProblemsId200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetProblemsId200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetProblemsId200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetProblemsId200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *GetProblemsId200Response) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetProblemsId200Response) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetProblemsId200Response) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetProblemsId200Response) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSupplier

`func (o *GetProblemsId200Response) GetSupplier() string`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *GetProblemsId200Response) GetSupplierOk() (*string, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *GetProblemsId200Response) SetSupplier(v string)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *GetProblemsId200Response) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### SetSupplierNil

`func (o *GetProblemsId200Response) SetSupplierNil(b bool)`

 SetSupplierNil sets the value for Supplier to be an explicit nil

### UnsetSupplier
`func (o *GetProblemsId200Response) UnsetSupplier()`

UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
### GetSupplierRequestID

`func (o *GetProblemsId200Response) GetSupplierRequestID() string`

GetSupplierRequestID returns the SupplierRequestID field if non-nil, zero value otherwise.

### GetSupplierRequestIDOk

`func (o *GetProblemsId200Response) GetSupplierRequestIDOk() (*string, bool)`

GetSupplierRequestIDOk returns a tuple with the SupplierRequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierRequestID

`func (o *GetProblemsId200Response) SetSupplierRequestID(v string)`

SetSupplierRequestID sets SupplierRequestID field to given value.

### HasSupplierRequestID

`func (o *GetProblemsId200Response) HasSupplierRequestID() bool`

HasSupplierRequestID returns a boolean if a field has been set.

### SetSupplierRequestIDNil

`func (o *GetProblemsId200Response) SetSupplierRequestIDNil(b bool)`

 SetSupplierRequestIDNil sets the value for SupplierRequestID to be an explicit nil

### UnsetSupplierRequestID
`func (o *GetProblemsId200Response) UnsetSupplierRequestID()`

UnsetSupplierRequestID ensures that no value is present for SupplierRequestID, not even an explicit nil
### GetTeam

`func (o *GetProblemsId200Response) GetTeam() GetRequestsId200ResponseCreatedBy`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GetProblemsId200Response) GetTeamOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GetProblemsId200Response) SetTeam(v GetRequestsId200ResponseCreatedBy)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GetProblemsId200Response) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetUiExtension

`func (o *GetProblemsId200Response) GetUiExtension() GetPeopleId200ResponseUiExtension`

GetUiExtension returns the UiExtension field if non-nil, zero value otherwise.

### GetUiExtensionOk

`func (o *GetProblemsId200Response) GetUiExtensionOk() (*GetPeopleId200ResponseUiExtension, bool)`

GetUiExtensionOk returns a tuple with the UiExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiExtension

`func (o *GetProblemsId200Response) SetUiExtension(v GetPeopleId200ResponseUiExtension)`

SetUiExtension sets UiExtension field to given value.

### HasUiExtension

`func (o *GetProblemsId200Response) HasUiExtension() bool`

HasUiExtension returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetProblemsId200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetProblemsId200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetProblemsId200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetProblemsId200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrgent

`func (o *GetProblemsId200Response) GetUrgent() bool`

GetUrgent returns the Urgent field if non-nil, zero value otherwise.

### GetUrgentOk

`func (o *GetProblemsId200Response) GetUrgentOk() (*bool, bool)`

GetUrgentOk returns a tuple with the Urgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgent

`func (o *GetProblemsId200Response) SetUrgent(v bool)`

SetUrgent sets Urgent field to given value.

### HasUrgent

`func (o *GetProblemsId200Response) HasUrgent() bool`

HasUrgent returns a boolean if a field has been set.

### GetWaitingUntil

`func (o *GetProblemsId200Response) GetWaitingUntil() string`

GetWaitingUntil returns the WaitingUntil field if non-nil, zero value otherwise.

### GetWaitingUntilOk

`func (o *GetProblemsId200Response) GetWaitingUntilOk() (*string, bool)`

GetWaitingUntilOk returns a tuple with the WaitingUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingUntil

`func (o *GetProblemsId200Response) SetWaitingUntil(v string)`

SetWaitingUntil sets WaitingUntil field to given value.

### HasWaitingUntil

`func (o *GetProblemsId200Response) HasWaitingUntil() bool`

HasWaitingUntil returns a boolean if a field has been set.

### SetWaitingUntilNil

`func (o *GetProblemsId200Response) SetWaitingUntilNil(b bool)`

 SetWaitingUntilNil sets the value for WaitingUntil to be an explicit nil

### UnsetWaitingUntil
`func (o *GetProblemsId200Response) UnsetWaitingUntil()`

UnsetWaitingUntil ensures that no value is present for WaitingUntil, not even an explicit nil
### GetWorkaround

`func (o *GetProblemsId200Response) GetWorkaround() string`

GetWorkaround returns the Workaround field if non-nil, zero value otherwise.

### GetWorkaroundOk

`func (o *GetProblemsId200Response) GetWorkaroundOk() (*string, bool)`

GetWorkaroundOk returns a tuple with the Workaround field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkaround

`func (o *GetProblemsId200Response) SetWorkaround(v string)`

SetWorkaround sets Workaround field to given value.

### HasWorkaround

`func (o *GetProblemsId200Response) HasWorkaround() bool`

HasWorkaround returns a boolean if a field has been set.

### GetWorkflow

`func (o *GetProblemsId200Response) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *GetProblemsId200Response) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *GetProblemsId200Response) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *GetProblemsId200Response) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### SetWorkflowNil

`func (o *GetProblemsId200Response) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *GetProblemsId200Response) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


