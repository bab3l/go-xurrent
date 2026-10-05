# GetRequestsId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**Addressed** | Pointer to **bool** |  | [optional] 
**AgileBoard** | Pointer to **NullableString** |  | [optional] 
**AgileBoardColumn** | Pointer to **NullableString** |  | [optional] 
**AgileBoardColumnPosition** | Pointer to **NullableString** |  | [optional] 
**AssignmentCount** | Pointer to **float32** |  | [optional] 
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CheckedItems** | Pointer to **NullableString** |  | [optional] 
**Ci** | Pointer to [**GetRequestsId200ResponseCi**](GetRequestsId200ResponseCi.md) |  | [optional] 
**CiId** | Pointer to **float32** |  | [optional] 
**ClosureCode** | Pointer to **NullableString** |  | [optional] 
**CompletedAt** | Pointer to **string** |  | [optional] 
**CompletionReason** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**CustomFields** | Pointer to [**[]GetRequestsId200ResponseCustomFieldsInner**](GetRequestsId200ResponseCustomFieldsInner.md) |  | [optional] 
**DesiredCompletionAt** | Pointer to **string** |  | [optional] 
**DowntimeEndAt** | Pointer to **string** |  | [optional] 
**DowntimeStartAt** | Pointer to **string** |  | [optional] 
**Feedback** | Pointer to [**GetRequestsId200ResponseFeedback**](GetRequestsId200ResponseFeedback.md) |  | [optional] 
**FeedbackOnKnowledgeArticle** | Pointer to **NullableString** |  | [optional] 
**GroupedInto** | Pointer to **NullableString** |  | [optional] 
**Grouping** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**KnowledgeArticle** | Pointer to **NullableString** |  | [optional] 
**MajorIncidentStatus** | Pointer to **NullableString** |  | [optional] 
**Member** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**NewAssignment** | Pointer to **bool** |  | [optional] 
**NextTargetAt** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**PlannedEffort** | Pointer to **NullableString** |  | [optional] 
**Problem** | Pointer to **NullableString** |  | [optional] 
**ProductBacklog** | Pointer to **NullableString** |  | [optional] 
**ProductBacklogEstimate** | Pointer to **NullableString** |  | [optional] 
**ProductBacklogPosition** | Pointer to **NullableString** |  | [optional] 
**Project** | Pointer to **NullableString** |  | [optional] 
**ProviderNotAccountable** | Pointer to **bool** |  | [optional] 
**ProviderWasNotAccountable** | Pointer to **bool** |  | [optional] 
**ReopenCount** | Pointer to **float32** |  | [optional] 
**RequestedBy** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**RequestedFor** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**RequesterResolutionTargetAt** | Pointer to **string** |  | [optional] 
**Reservation** | Pointer to **NullableString** |  | [optional] 
**ResolutionDuration** | Pointer to **float32** |  | [optional] 
**ResolutionTargetAt** | Pointer to **string** |  | [optional] 
**ResponseTargetAt** | Pointer to **string** |  | [optional] 
**Reviewed** | Pointer to **bool** |  | [optional] 
**RfcType** | Pointer to **NullableString** |  | [optional] 
**Satisfaction** | Pointer to **NullableString** |  | [optional] 
**ServiceInstance** | Pointer to [**GetRequestsId200ResponseServiceInstance**](GetRequestsId200ResponseServiceInstance.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **NullableString** |  | [optional] 
**Supplier** | Pointer to **NullableString** |  | [optional] 
**SupplierRequestID** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Task** | Pointer to **NullableString** |  | [optional] 
**Team** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**Template** | Pointer to [**GetRequestsId200ResponseTemplate**](GetRequestsId200ResponseTemplate.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Urgent** | Pointer to **bool** |  | [optional] 
**WaitingUntil** | Pointer to **NullableString** |  | [optional] 
**Workflow** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetRequestsId200Response

`func NewGetRequestsId200Response() *GetRequestsId200Response`

NewGetRequestsId200Response instantiates a new GetRequestsId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRequestsId200ResponseWithDefaults

`func NewGetRequestsId200ResponseWithDefaults() *GetRequestsId200Response`

NewGetRequestsId200ResponseWithDefaults instantiates a new GetRequestsId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetRequestsId200Response) GetAccount() GetRequestsId200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetRequestsId200Response) GetAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetRequestsId200Response) SetAccount(v GetRequestsId200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetRequestsId200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAddressed

`func (o *GetRequestsId200Response) GetAddressed() bool`

GetAddressed returns the Addressed field if non-nil, zero value otherwise.

### GetAddressedOk

`func (o *GetRequestsId200Response) GetAddressedOk() (*bool, bool)`

GetAddressedOk returns a tuple with the Addressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressed

`func (o *GetRequestsId200Response) SetAddressed(v bool)`

SetAddressed sets Addressed field to given value.

### HasAddressed

`func (o *GetRequestsId200Response) HasAddressed() bool`

HasAddressed returns a boolean if a field has been set.

### GetAgileBoard

`func (o *GetRequestsId200Response) GetAgileBoard() string`

GetAgileBoard returns the AgileBoard field if non-nil, zero value otherwise.

### GetAgileBoardOk

`func (o *GetRequestsId200Response) GetAgileBoardOk() (*string, bool)`

GetAgileBoardOk returns a tuple with the AgileBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgileBoard

`func (o *GetRequestsId200Response) SetAgileBoard(v string)`

SetAgileBoard sets AgileBoard field to given value.

### HasAgileBoard

`func (o *GetRequestsId200Response) HasAgileBoard() bool`

HasAgileBoard returns a boolean if a field has been set.

### SetAgileBoardNil

`func (o *GetRequestsId200Response) SetAgileBoardNil(b bool)`

 SetAgileBoardNil sets the value for AgileBoard to be an explicit nil

### UnsetAgileBoard
`func (o *GetRequestsId200Response) UnsetAgileBoard()`

UnsetAgileBoard ensures that no value is present for AgileBoard, not even an explicit nil
### GetAgileBoardColumn

`func (o *GetRequestsId200Response) GetAgileBoardColumn() string`

GetAgileBoardColumn returns the AgileBoardColumn field if non-nil, zero value otherwise.

### GetAgileBoardColumnOk

`func (o *GetRequestsId200Response) GetAgileBoardColumnOk() (*string, bool)`

GetAgileBoardColumnOk returns a tuple with the AgileBoardColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgileBoardColumn

`func (o *GetRequestsId200Response) SetAgileBoardColumn(v string)`

SetAgileBoardColumn sets AgileBoardColumn field to given value.

### HasAgileBoardColumn

`func (o *GetRequestsId200Response) HasAgileBoardColumn() bool`

HasAgileBoardColumn returns a boolean if a field has been set.

### SetAgileBoardColumnNil

`func (o *GetRequestsId200Response) SetAgileBoardColumnNil(b bool)`

 SetAgileBoardColumnNil sets the value for AgileBoardColumn to be an explicit nil

### UnsetAgileBoardColumn
`func (o *GetRequestsId200Response) UnsetAgileBoardColumn()`

UnsetAgileBoardColumn ensures that no value is present for AgileBoardColumn, not even an explicit nil
### GetAgileBoardColumnPosition

`func (o *GetRequestsId200Response) GetAgileBoardColumnPosition() string`

GetAgileBoardColumnPosition returns the AgileBoardColumnPosition field if non-nil, zero value otherwise.

### GetAgileBoardColumnPositionOk

`func (o *GetRequestsId200Response) GetAgileBoardColumnPositionOk() (*string, bool)`

GetAgileBoardColumnPositionOk returns a tuple with the AgileBoardColumnPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgileBoardColumnPosition

`func (o *GetRequestsId200Response) SetAgileBoardColumnPosition(v string)`

SetAgileBoardColumnPosition sets AgileBoardColumnPosition field to given value.

### HasAgileBoardColumnPosition

`func (o *GetRequestsId200Response) HasAgileBoardColumnPosition() bool`

HasAgileBoardColumnPosition returns a boolean if a field has been set.

### SetAgileBoardColumnPositionNil

`func (o *GetRequestsId200Response) SetAgileBoardColumnPositionNil(b bool)`

 SetAgileBoardColumnPositionNil sets the value for AgileBoardColumnPosition to be an explicit nil

### UnsetAgileBoardColumnPosition
`func (o *GetRequestsId200Response) UnsetAgileBoardColumnPosition()`

UnsetAgileBoardColumnPosition ensures that no value is present for AgileBoardColumnPosition, not even an explicit nil
### GetAssignmentCount

`func (o *GetRequestsId200Response) GetAssignmentCount() float32`

GetAssignmentCount returns the AssignmentCount field if non-nil, zero value otherwise.

### GetAssignmentCountOk

`func (o *GetRequestsId200Response) GetAssignmentCountOk() (*float32, bool)`

GetAssignmentCountOk returns a tuple with the AssignmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentCount

`func (o *GetRequestsId200Response) SetAssignmentCount(v float32)`

SetAssignmentCount sets AssignmentCount field to given value.

### HasAssignmentCount

`func (o *GetRequestsId200Response) HasAssignmentCount() bool`

HasAssignmentCount returns a boolean if a field has been set.

### GetAttachments

`func (o *GetRequestsId200Response) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetRequestsId200Response) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetRequestsId200Response) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetRequestsId200Response) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCategory

`func (o *GetRequestsId200Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetRequestsId200Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetRequestsId200Response) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetRequestsId200Response) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCheckedItems

`func (o *GetRequestsId200Response) GetCheckedItems() string`

GetCheckedItems returns the CheckedItems field if non-nil, zero value otherwise.

### GetCheckedItemsOk

`func (o *GetRequestsId200Response) GetCheckedItemsOk() (*string, bool)`

GetCheckedItemsOk returns a tuple with the CheckedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedItems

`func (o *GetRequestsId200Response) SetCheckedItems(v string)`

SetCheckedItems sets CheckedItems field to given value.

### HasCheckedItems

`func (o *GetRequestsId200Response) HasCheckedItems() bool`

HasCheckedItems returns a boolean if a field has been set.

### SetCheckedItemsNil

`func (o *GetRequestsId200Response) SetCheckedItemsNil(b bool)`

 SetCheckedItemsNil sets the value for CheckedItems to be an explicit nil

### UnsetCheckedItems
`func (o *GetRequestsId200Response) UnsetCheckedItems()`

UnsetCheckedItems ensures that no value is present for CheckedItems, not even an explicit nil
### GetCi

`func (o *GetRequestsId200Response) GetCi() GetRequestsId200ResponseCi`

GetCi returns the Ci field if non-nil, zero value otherwise.

### GetCiOk

`func (o *GetRequestsId200Response) GetCiOk() (*GetRequestsId200ResponseCi, bool)`

GetCiOk returns a tuple with the Ci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCi

`func (o *GetRequestsId200Response) SetCi(v GetRequestsId200ResponseCi)`

SetCi sets Ci field to given value.

### HasCi

`func (o *GetRequestsId200Response) HasCi() bool`

HasCi returns a boolean if a field has been set.

### GetCiId

`func (o *GetRequestsId200Response) GetCiId() float32`

GetCiId returns the CiId field if non-nil, zero value otherwise.

### GetCiIdOk

`func (o *GetRequestsId200Response) GetCiIdOk() (*float32, bool)`

GetCiIdOk returns a tuple with the CiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiId

`func (o *GetRequestsId200Response) SetCiId(v float32)`

SetCiId sets CiId field to given value.

### HasCiId

`func (o *GetRequestsId200Response) HasCiId() bool`

HasCiId returns a boolean if a field has been set.

### GetClosureCode

`func (o *GetRequestsId200Response) GetClosureCode() string`

GetClosureCode returns the ClosureCode field if non-nil, zero value otherwise.

### GetClosureCodeOk

`func (o *GetRequestsId200Response) GetClosureCodeOk() (*string, bool)`

GetClosureCodeOk returns a tuple with the ClosureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosureCode

`func (o *GetRequestsId200Response) SetClosureCode(v string)`

SetClosureCode sets ClosureCode field to given value.

### HasClosureCode

`func (o *GetRequestsId200Response) HasClosureCode() bool`

HasClosureCode returns a boolean if a field has been set.

### SetClosureCodeNil

`func (o *GetRequestsId200Response) SetClosureCodeNil(b bool)`

 SetClosureCodeNil sets the value for ClosureCode to be an explicit nil

### UnsetClosureCode
`func (o *GetRequestsId200Response) UnsetClosureCode()`

UnsetClosureCode ensures that no value is present for ClosureCode, not even an explicit nil
### GetCompletedAt

`func (o *GetRequestsId200Response) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetRequestsId200Response) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetRequestsId200Response) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetRequestsId200Response) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCompletionReason

`func (o *GetRequestsId200Response) GetCompletionReason() string`

GetCompletionReason returns the CompletionReason field if non-nil, zero value otherwise.

### GetCompletionReasonOk

`func (o *GetRequestsId200Response) GetCompletionReasonOk() (*string, bool)`

GetCompletionReasonOk returns a tuple with the CompletionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionReason

`func (o *GetRequestsId200Response) SetCompletionReason(v string)`

SetCompletionReason sets CompletionReason field to given value.

### HasCompletionReason

`func (o *GetRequestsId200Response) HasCompletionReason() bool`

HasCompletionReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetRequestsId200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetRequestsId200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetRequestsId200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetRequestsId200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetRequestsId200Response) GetCreatedBy() GetRequestsId200ResponseCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetRequestsId200Response) GetCreatedByOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetRequestsId200Response) SetCreatedBy(v GetRequestsId200ResponseCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetRequestsId200Response) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCustomFields

`func (o *GetRequestsId200Response) GetCustomFields() []GetRequestsId200ResponseCustomFieldsInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GetRequestsId200Response) GetCustomFieldsOk() (*[]GetRequestsId200ResponseCustomFieldsInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GetRequestsId200Response) SetCustomFields(v []GetRequestsId200ResponseCustomFieldsInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GetRequestsId200Response) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDesiredCompletionAt

`func (o *GetRequestsId200Response) GetDesiredCompletionAt() string`

GetDesiredCompletionAt returns the DesiredCompletionAt field if non-nil, zero value otherwise.

### GetDesiredCompletionAtOk

`func (o *GetRequestsId200Response) GetDesiredCompletionAtOk() (*string, bool)`

GetDesiredCompletionAtOk returns a tuple with the DesiredCompletionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredCompletionAt

`func (o *GetRequestsId200Response) SetDesiredCompletionAt(v string)`

SetDesiredCompletionAt sets DesiredCompletionAt field to given value.

### HasDesiredCompletionAt

`func (o *GetRequestsId200Response) HasDesiredCompletionAt() bool`

HasDesiredCompletionAt returns a boolean if a field has been set.

### GetDowntimeEndAt

`func (o *GetRequestsId200Response) GetDowntimeEndAt() string`

GetDowntimeEndAt returns the DowntimeEndAt field if non-nil, zero value otherwise.

### GetDowntimeEndAtOk

`func (o *GetRequestsId200Response) GetDowntimeEndAtOk() (*string, bool)`

GetDowntimeEndAtOk returns a tuple with the DowntimeEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntimeEndAt

`func (o *GetRequestsId200Response) SetDowntimeEndAt(v string)`

SetDowntimeEndAt sets DowntimeEndAt field to given value.

### HasDowntimeEndAt

`func (o *GetRequestsId200Response) HasDowntimeEndAt() bool`

HasDowntimeEndAt returns a boolean if a field has been set.

### GetDowntimeStartAt

`func (o *GetRequestsId200Response) GetDowntimeStartAt() string`

GetDowntimeStartAt returns the DowntimeStartAt field if non-nil, zero value otherwise.

### GetDowntimeStartAtOk

`func (o *GetRequestsId200Response) GetDowntimeStartAtOk() (*string, bool)`

GetDowntimeStartAtOk returns a tuple with the DowntimeStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntimeStartAt

`func (o *GetRequestsId200Response) SetDowntimeStartAt(v string)`

SetDowntimeStartAt sets DowntimeStartAt field to given value.

### HasDowntimeStartAt

`func (o *GetRequestsId200Response) HasDowntimeStartAt() bool`

HasDowntimeStartAt returns a boolean if a field has been set.

### GetFeedback

`func (o *GetRequestsId200Response) GetFeedback() GetRequestsId200ResponseFeedback`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *GetRequestsId200Response) GetFeedbackOk() (*GetRequestsId200ResponseFeedback, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *GetRequestsId200Response) SetFeedback(v GetRequestsId200ResponseFeedback)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *GetRequestsId200Response) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetFeedbackOnKnowledgeArticle

`func (o *GetRequestsId200Response) GetFeedbackOnKnowledgeArticle() string`

GetFeedbackOnKnowledgeArticle returns the FeedbackOnKnowledgeArticle field if non-nil, zero value otherwise.

### GetFeedbackOnKnowledgeArticleOk

`func (o *GetRequestsId200Response) GetFeedbackOnKnowledgeArticleOk() (*string, bool)`

GetFeedbackOnKnowledgeArticleOk returns a tuple with the FeedbackOnKnowledgeArticle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackOnKnowledgeArticle

`func (o *GetRequestsId200Response) SetFeedbackOnKnowledgeArticle(v string)`

SetFeedbackOnKnowledgeArticle sets FeedbackOnKnowledgeArticle field to given value.

### HasFeedbackOnKnowledgeArticle

`func (o *GetRequestsId200Response) HasFeedbackOnKnowledgeArticle() bool`

HasFeedbackOnKnowledgeArticle returns a boolean if a field has been set.

### SetFeedbackOnKnowledgeArticleNil

`func (o *GetRequestsId200Response) SetFeedbackOnKnowledgeArticleNil(b bool)`

 SetFeedbackOnKnowledgeArticleNil sets the value for FeedbackOnKnowledgeArticle to be an explicit nil

### UnsetFeedbackOnKnowledgeArticle
`func (o *GetRequestsId200Response) UnsetFeedbackOnKnowledgeArticle()`

UnsetFeedbackOnKnowledgeArticle ensures that no value is present for FeedbackOnKnowledgeArticle, not even an explicit nil
### GetGroupedInto

`func (o *GetRequestsId200Response) GetGroupedInto() string`

GetGroupedInto returns the GroupedInto field if non-nil, zero value otherwise.

### GetGroupedIntoOk

`func (o *GetRequestsId200Response) GetGroupedIntoOk() (*string, bool)`

GetGroupedIntoOk returns a tuple with the GroupedInto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedInto

`func (o *GetRequestsId200Response) SetGroupedInto(v string)`

SetGroupedInto sets GroupedInto field to given value.

### HasGroupedInto

`func (o *GetRequestsId200Response) HasGroupedInto() bool`

HasGroupedInto returns a boolean if a field has been set.

### SetGroupedIntoNil

`func (o *GetRequestsId200Response) SetGroupedIntoNil(b bool)`

 SetGroupedIntoNil sets the value for GroupedInto to be an explicit nil

### UnsetGroupedInto
`func (o *GetRequestsId200Response) UnsetGroupedInto()`

UnsetGroupedInto ensures that no value is present for GroupedInto, not even an explicit nil
### GetGrouping

`func (o *GetRequestsId200Response) GetGrouping() string`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *GetRequestsId200Response) GetGroupingOk() (*string, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *GetRequestsId200Response) SetGrouping(v string)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *GetRequestsId200Response) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetId

`func (o *GetRequestsId200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRequestsId200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRequestsId200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetRequestsId200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImpact

`func (o *GetRequestsId200Response) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *GetRequestsId200Response) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *GetRequestsId200Response) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *GetRequestsId200Response) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetKnowledgeArticle

`func (o *GetRequestsId200Response) GetKnowledgeArticle() string`

GetKnowledgeArticle returns the KnowledgeArticle field if non-nil, zero value otherwise.

### GetKnowledgeArticleOk

`func (o *GetRequestsId200Response) GetKnowledgeArticleOk() (*string, bool)`

GetKnowledgeArticleOk returns a tuple with the KnowledgeArticle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeArticle

`func (o *GetRequestsId200Response) SetKnowledgeArticle(v string)`

SetKnowledgeArticle sets KnowledgeArticle field to given value.

### HasKnowledgeArticle

`func (o *GetRequestsId200Response) HasKnowledgeArticle() bool`

HasKnowledgeArticle returns a boolean if a field has been set.

### SetKnowledgeArticleNil

`func (o *GetRequestsId200Response) SetKnowledgeArticleNil(b bool)`

 SetKnowledgeArticleNil sets the value for KnowledgeArticle to be an explicit nil

### UnsetKnowledgeArticle
`func (o *GetRequestsId200Response) UnsetKnowledgeArticle()`

UnsetKnowledgeArticle ensures that no value is present for KnowledgeArticle, not even an explicit nil
### GetMajorIncidentStatus

`func (o *GetRequestsId200Response) GetMajorIncidentStatus() string`

GetMajorIncidentStatus returns the MajorIncidentStatus field if non-nil, zero value otherwise.

### GetMajorIncidentStatusOk

`func (o *GetRequestsId200Response) GetMajorIncidentStatusOk() (*string, bool)`

GetMajorIncidentStatusOk returns a tuple with the MajorIncidentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorIncidentStatus

`func (o *GetRequestsId200Response) SetMajorIncidentStatus(v string)`

SetMajorIncidentStatus sets MajorIncidentStatus field to given value.

### HasMajorIncidentStatus

`func (o *GetRequestsId200Response) HasMajorIncidentStatus() bool`

HasMajorIncidentStatus returns a boolean if a field has been set.

### SetMajorIncidentStatusNil

`func (o *GetRequestsId200Response) SetMajorIncidentStatusNil(b bool)`

 SetMajorIncidentStatusNil sets the value for MajorIncidentStatus to be an explicit nil

### UnsetMajorIncidentStatus
`func (o *GetRequestsId200Response) UnsetMajorIncidentStatus()`

UnsetMajorIncidentStatus ensures that no value is present for MajorIncidentStatus, not even an explicit nil
### GetMember

`func (o *GetRequestsId200Response) GetMember() GetRequestsId200ResponseCreatedBy`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetRequestsId200Response) GetMemberOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetRequestsId200Response) SetMember(v GetRequestsId200ResponseCreatedBy)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetRequestsId200Response) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetNewAssignment

`func (o *GetRequestsId200Response) GetNewAssignment() bool`

GetNewAssignment returns the NewAssignment field if non-nil, zero value otherwise.

### GetNewAssignmentOk

`func (o *GetRequestsId200Response) GetNewAssignmentOk() (*bool, bool)`

GetNewAssignmentOk returns a tuple with the NewAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAssignment

`func (o *GetRequestsId200Response) SetNewAssignment(v bool)`

SetNewAssignment sets NewAssignment field to given value.

### HasNewAssignment

`func (o *GetRequestsId200Response) HasNewAssignment() bool`

HasNewAssignment returns a boolean if a field has been set.

### GetNextTargetAt

`func (o *GetRequestsId200Response) GetNextTargetAt() string`

GetNextTargetAt returns the NextTargetAt field if non-nil, zero value otherwise.

### GetNextTargetAtOk

`func (o *GetRequestsId200Response) GetNextTargetAtOk() (*string, bool)`

GetNextTargetAtOk returns a tuple with the NextTargetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTargetAt

`func (o *GetRequestsId200Response) SetNextTargetAt(v string)`

SetNextTargetAt sets NextTargetAt field to given value.

### HasNextTargetAt

`func (o *GetRequestsId200Response) HasNextTargetAt() bool`

HasNextTargetAt returns a boolean if a field has been set.

### GetNodeID

`func (o *GetRequestsId200Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *GetRequestsId200Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *GetRequestsId200Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *GetRequestsId200Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetOrganization

`func (o *GetRequestsId200Response) GetOrganization() GetRequestsId200ResponseCreatedBy`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GetRequestsId200Response) GetOrganizationOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GetRequestsId200Response) SetOrganization(v GetRequestsId200ResponseCreatedBy)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *GetRequestsId200Response) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPlannedEffort

`func (o *GetRequestsId200Response) GetPlannedEffort() string`

GetPlannedEffort returns the PlannedEffort field if non-nil, zero value otherwise.

### GetPlannedEffortOk

`func (o *GetRequestsId200Response) GetPlannedEffortOk() (*string, bool)`

GetPlannedEffortOk returns a tuple with the PlannedEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedEffort

`func (o *GetRequestsId200Response) SetPlannedEffort(v string)`

SetPlannedEffort sets PlannedEffort field to given value.

### HasPlannedEffort

`func (o *GetRequestsId200Response) HasPlannedEffort() bool`

HasPlannedEffort returns a boolean if a field has been set.

### SetPlannedEffortNil

`func (o *GetRequestsId200Response) SetPlannedEffortNil(b bool)`

 SetPlannedEffortNil sets the value for PlannedEffort to be an explicit nil

### UnsetPlannedEffort
`func (o *GetRequestsId200Response) UnsetPlannedEffort()`

UnsetPlannedEffort ensures that no value is present for PlannedEffort, not even an explicit nil
### GetProblem

`func (o *GetRequestsId200Response) GetProblem() string`

GetProblem returns the Problem field if non-nil, zero value otherwise.

### GetProblemOk

`func (o *GetRequestsId200Response) GetProblemOk() (*string, bool)`

GetProblemOk returns a tuple with the Problem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblem

`func (o *GetRequestsId200Response) SetProblem(v string)`

SetProblem sets Problem field to given value.

### HasProblem

`func (o *GetRequestsId200Response) HasProblem() bool`

HasProblem returns a boolean if a field has been set.

### SetProblemNil

`func (o *GetRequestsId200Response) SetProblemNil(b bool)`

 SetProblemNil sets the value for Problem to be an explicit nil

### UnsetProblem
`func (o *GetRequestsId200Response) UnsetProblem()`

UnsetProblem ensures that no value is present for Problem, not even an explicit nil
### GetProductBacklog

`func (o *GetRequestsId200Response) GetProductBacklog() string`

GetProductBacklog returns the ProductBacklog field if non-nil, zero value otherwise.

### GetProductBacklogOk

`func (o *GetRequestsId200Response) GetProductBacklogOk() (*string, bool)`

GetProductBacklogOk returns a tuple with the ProductBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBacklog

`func (o *GetRequestsId200Response) SetProductBacklog(v string)`

SetProductBacklog sets ProductBacklog field to given value.

### HasProductBacklog

`func (o *GetRequestsId200Response) HasProductBacklog() bool`

HasProductBacklog returns a boolean if a field has been set.

### SetProductBacklogNil

`func (o *GetRequestsId200Response) SetProductBacklogNil(b bool)`

 SetProductBacklogNil sets the value for ProductBacklog to be an explicit nil

### UnsetProductBacklog
`func (o *GetRequestsId200Response) UnsetProductBacklog()`

UnsetProductBacklog ensures that no value is present for ProductBacklog, not even an explicit nil
### GetProductBacklogEstimate

`func (o *GetRequestsId200Response) GetProductBacklogEstimate() string`

GetProductBacklogEstimate returns the ProductBacklogEstimate field if non-nil, zero value otherwise.

### GetProductBacklogEstimateOk

`func (o *GetRequestsId200Response) GetProductBacklogEstimateOk() (*string, bool)`

GetProductBacklogEstimateOk returns a tuple with the ProductBacklogEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBacklogEstimate

`func (o *GetRequestsId200Response) SetProductBacklogEstimate(v string)`

SetProductBacklogEstimate sets ProductBacklogEstimate field to given value.

### HasProductBacklogEstimate

`func (o *GetRequestsId200Response) HasProductBacklogEstimate() bool`

HasProductBacklogEstimate returns a boolean if a field has been set.

### SetProductBacklogEstimateNil

`func (o *GetRequestsId200Response) SetProductBacklogEstimateNil(b bool)`

 SetProductBacklogEstimateNil sets the value for ProductBacklogEstimate to be an explicit nil

### UnsetProductBacklogEstimate
`func (o *GetRequestsId200Response) UnsetProductBacklogEstimate()`

UnsetProductBacklogEstimate ensures that no value is present for ProductBacklogEstimate, not even an explicit nil
### GetProductBacklogPosition

`func (o *GetRequestsId200Response) GetProductBacklogPosition() string`

GetProductBacklogPosition returns the ProductBacklogPosition field if non-nil, zero value otherwise.

### GetProductBacklogPositionOk

`func (o *GetRequestsId200Response) GetProductBacklogPositionOk() (*string, bool)`

GetProductBacklogPositionOk returns a tuple with the ProductBacklogPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBacklogPosition

`func (o *GetRequestsId200Response) SetProductBacklogPosition(v string)`

SetProductBacklogPosition sets ProductBacklogPosition field to given value.

### HasProductBacklogPosition

`func (o *GetRequestsId200Response) HasProductBacklogPosition() bool`

HasProductBacklogPosition returns a boolean if a field has been set.

### SetProductBacklogPositionNil

`func (o *GetRequestsId200Response) SetProductBacklogPositionNil(b bool)`

 SetProductBacklogPositionNil sets the value for ProductBacklogPosition to be an explicit nil

### UnsetProductBacklogPosition
`func (o *GetRequestsId200Response) UnsetProductBacklogPosition()`

UnsetProductBacklogPosition ensures that no value is present for ProductBacklogPosition, not even an explicit nil
### GetProject

`func (o *GetRequestsId200Response) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GetRequestsId200Response) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GetRequestsId200Response) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *GetRequestsId200Response) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *GetRequestsId200Response) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *GetRequestsId200Response) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetProviderNotAccountable

`func (o *GetRequestsId200Response) GetProviderNotAccountable() bool`

GetProviderNotAccountable returns the ProviderNotAccountable field if non-nil, zero value otherwise.

### GetProviderNotAccountableOk

`func (o *GetRequestsId200Response) GetProviderNotAccountableOk() (*bool, bool)`

GetProviderNotAccountableOk returns a tuple with the ProviderNotAccountable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNotAccountable

`func (o *GetRequestsId200Response) SetProviderNotAccountable(v bool)`

SetProviderNotAccountable sets ProviderNotAccountable field to given value.

### HasProviderNotAccountable

`func (o *GetRequestsId200Response) HasProviderNotAccountable() bool`

HasProviderNotAccountable returns a boolean if a field has been set.

### GetProviderWasNotAccountable

`func (o *GetRequestsId200Response) GetProviderWasNotAccountable() bool`

GetProviderWasNotAccountable returns the ProviderWasNotAccountable field if non-nil, zero value otherwise.

### GetProviderWasNotAccountableOk

`func (o *GetRequestsId200Response) GetProviderWasNotAccountableOk() (*bool, bool)`

GetProviderWasNotAccountableOk returns a tuple with the ProviderWasNotAccountable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderWasNotAccountable

`func (o *GetRequestsId200Response) SetProviderWasNotAccountable(v bool)`

SetProviderWasNotAccountable sets ProviderWasNotAccountable field to given value.

### HasProviderWasNotAccountable

`func (o *GetRequestsId200Response) HasProviderWasNotAccountable() bool`

HasProviderWasNotAccountable returns a boolean if a field has been set.

### GetReopenCount

`func (o *GetRequestsId200Response) GetReopenCount() float32`

GetReopenCount returns the ReopenCount field if non-nil, zero value otherwise.

### GetReopenCountOk

`func (o *GetRequestsId200Response) GetReopenCountOk() (*float32, bool)`

GetReopenCountOk returns a tuple with the ReopenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenCount

`func (o *GetRequestsId200Response) SetReopenCount(v float32)`

SetReopenCount sets ReopenCount field to given value.

### HasReopenCount

`func (o *GetRequestsId200Response) HasReopenCount() bool`

HasReopenCount returns a boolean if a field has been set.

### GetRequestedBy

`func (o *GetRequestsId200Response) GetRequestedBy() GetRequestsId200ResponseCreatedBy`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *GetRequestsId200Response) GetRequestedByOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *GetRequestsId200Response) SetRequestedBy(v GetRequestsId200ResponseCreatedBy)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *GetRequestsId200Response) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetRequestedFor

`func (o *GetRequestsId200Response) GetRequestedFor() GetRequestsId200ResponseCreatedBy`

GetRequestedFor returns the RequestedFor field if non-nil, zero value otherwise.

### GetRequestedForOk

`func (o *GetRequestsId200Response) GetRequestedForOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetRequestedForOk returns a tuple with the RequestedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFor

`func (o *GetRequestsId200Response) SetRequestedFor(v GetRequestsId200ResponseCreatedBy)`

SetRequestedFor sets RequestedFor field to given value.

### HasRequestedFor

`func (o *GetRequestsId200Response) HasRequestedFor() bool`

HasRequestedFor returns a boolean if a field has been set.

### GetRequesterResolutionTargetAt

`func (o *GetRequestsId200Response) GetRequesterResolutionTargetAt() string`

GetRequesterResolutionTargetAt returns the RequesterResolutionTargetAt field if non-nil, zero value otherwise.

### GetRequesterResolutionTargetAtOk

`func (o *GetRequestsId200Response) GetRequesterResolutionTargetAtOk() (*string, bool)`

GetRequesterResolutionTargetAtOk returns a tuple with the RequesterResolutionTargetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterResolutionTargetAt

`func (o *GetRequestsId200Response) SetRequesterResolutionTargetAt(v string)`

SetRequesterResolutionTargetAt sets RequesterResolutionTargetAt field to given value.

### HasRequesterResolutionTargetAt

`func (o *GetRequestsId200Response) HasRequesterResolutionTargetAt() bool`

HasRequesterResolutionTargetAt returns a boolean if a field has been set.

### GetReservation

`func (o *GetRequestsId200Response) GetReservation() string`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *GetRequestsId200Response) GetReservationOk() (*string, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *GetRequestsId200Response) SetReservation(v string)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *GetRequestsId200Response) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### SetReservationNil

`func (o *GetRequestsId200Response) SetReservationNil(b bool)`

 SetReservationNil sets the value for Reservation to be an explicit nil

### UnsetReservation
`func (o *GetRequestsId200Response) UnsetReservation()`

UnsetReservation ensures that no value is present for Reservation, not even an explicit nil
### GetResolutionDuration

`func (o *GetRequestsId200Response) GetResolutionDuration() float32`

GetResolutionDuration returns the ResolutionDuration field if non-nil, zero value otherwise.

### GetResolutionDurationOk

`func (o *GetRequestsId200Response) GetResolutionDurationOk() (*float32, bool)`

GetResolutionDurationOk returns a tuple with the ResolutionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionDuration

`func (o *GetRequestsId200Response) SetResolutionDuration(v float32)`

SetResolutionDuration sets ResolutionDuration field to given value.

### HasResolutionDuration

`func (o *GetRequestsId200Response) HasResolutionDuration() bool`

HasResolutionDuration returns a boolean if a field has been set.

### GetResolutionTargetAt

`func (o *GetRequestsId200Response) GetResolutionTargetAt() string`

GetResolutionTargetAt returns the ResolutionTargetAt field if non-nil, zero value otherwise.

### GetResolutionTargetAtOk

`func (o *GetRequestsId200Response) GetResolutionTargetAtOk() (*string, bool)`

GetResolutionTargetAtOk returns a tuple with the ResolutionTargetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionTargetAt

`func (o *GetRequestsId200Response) SetResolutionTargetAt(v string)`

SetResolutionTargetAt sets ResolutionTargetAt field to given value.

### HasResolutionTargetAt

`func (o *GetRequestsId200Response) HasResolutionTargetAt() bool`

HasResolutionTargetAt returns a boolean if a field has been set.

### GetResponseTargetAt

`func (o *GetRequestsId200Response) GetResponseTargetAt() string`

GetResponseTargetAt returns the ResponseTargetAt field if non-nil, zero value otherwise.

### GetResponseTargetAtOk

`func (o *GetRequestsId200Response) GetResponseTargetAtOk() (*string, bool)`

GetResponseTargetAtOk returns a tuple with the ResponseTargetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTargetAt

`func (o *GetRequestsId200Response) SetResponseTargetAt(v string)`

SetResponseTargetAt sets ResponseTargetAt field to given value.

### HasResponseTargetAt

`func (o *GetRequestsId200Response) HasResponseTargetAt() bool`

HasResponseTargetAt returns a boolean if a field has been set.

### GetReviewed

`func (o *GetRequestsId200Response) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *GetRequestsId200Response) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *GetRequestsId200Response) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *GetRequestsId200Response) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetRfcType

`func (o *GetRequestsId200Response) GetRfcType() string`

GetRfcType returns the RfcType field if non-nil, zero value otherwise.

### GetRfcTypeOk

`func (o *GetRequestsId200Response) GetRfcTypeOk() (*string, bool)`

GetRfcTypeOk returns a tuple with the RfcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfcType

`func (o *GetRequestsId200Response) SetRfcType(v string)`

SetRfcType sets RfcType field to given value.

### HasRfcType

`func (o *GetRequestsId200Response) HasRfcType() bool`

HasRfcType returns a boolean if a field has been set.

### SetRfcTypeNil

`func (o *GetRequestsId200Response) SetRfcTypeNil(b bool)`

 SetRfcTypeNil sets the value for RfcType to be an explicit nil

### UnsetRfcType
`func (o *GetRequestsId200Response) UnsetRfcType()`

UnsetRfcType ensures that no value is present for RfcType, not even an explicit nil
### GetSatisfaction

`func (o *GetRequestsId200Response) GetSatisfaction() string`

GetSatisfaction returns the Satisfaction field if non-nil, zero value otherwise.

### GetSatisfactionOk

`func (o *GetRequestsId200Response) GetSatisfactionOk() (*string, bool)`

GetSatisfactionOk returns a tuple with the Satisfaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatisfaction

`func (o *GetRequestsId200Response) SetSatisfaction(v string)`

SetSatisfaction sets Satisfaction field to given value.

### HasSatisfaction

`func (o *GetRequestsId200Response) HasSatisfaction() bool`

HasSatisfaction returns a boolean if a field has been set.

### SetSatisfactionNil

`func (o *GetRequestsId200Response) SetSatisfactionNil(b bool)`

 SetSatisfactionNil sets the value for Satisfaction to be an explicit nil

### UnsetSatisfaction
`func (o *GetRequestsId200Response) UnsetSatisfaction()`

UnsetSatisfaction ensures that no value is present for Satisfaction, not even an explicit nil
### GetServiceInstance

`func (o *GetRequestsId200Response) GetServiceInstance() GetRequestsId200ResponseServiceInstance`

GetServiceInstance returns the ServiceInstance field if non-nil, zero value otherwise.

### GetServiceInstanceOk

`func (o *GetRequestsId200Response) GetServiceInstanceOk() (*GetRequestsId200ResponseServiceInstance, bool)`

GetServiceInstanceOk returns a tuple with the ServiceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstance

`func (o *GetRequestsId200Response) SetServiceInstance(v GetRequestsId200ResponseServiceInstance)`

SetServiceInstance sets ServiceInstance field to given value.

### HasServiceInstance

`func (o *GetRequestsId200Response) HasServiceInstance() bool`

HasServiceInstance returns a boolean if a field has been set.

### GetSource

`func (o *GetRequestsId200Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetRequestsId200Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetRequestsId200Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetRequestsId200Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *GetRequestsId200Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *GetRequestsId200Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *GetRequestsId200Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *GetRequestsId200Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *GetRequestsId200Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *GetRequestsId200Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetStatus

`func (o *GetRequestsId200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRequestsId200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRequestsId200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetRequestsId200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *GetRequestsId200Response) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetRequestsId200Response) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetRequestsId200Response) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetRequestsId200Response) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSummary

`func (o *GetRequestsId200Response) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetRequestsId200Response) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetRequestsId200Response) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetRequestsId200Response) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *GetRequestsId200Response) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *GetRequestsId200Response) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetSupplier

`func (o *GetRequestsId200Response) GetSupplier() string`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *GetRequestsId200Response) GetSupplierOk() (*string, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *GetRequestsId200Response) SetSupplier(v string)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *GetRequestsId200Response) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### SetSupplierNil

`func (o *GetRequestsId200Response) SetSupplierNil(b bool)`

 SetSupplierNil sets the value for Supplier to be an explicit nil

### UnsetSupplier
`func (o *GetRequestsId200Response) UnsetSupplier()`

UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
### GetSupplierRequestID

`func (o *GetRequestsId200Response) GetSupplierRequestID() string`

GetSupplierRequestID returns the SupplierRequestID field if non-nil, zero value otherwise.

### GetSupplierRequestIDOk

`func (o *GetRequestsId200Response) GetSupplierRequestIDOk() (*string, bool)`

GetSupplierRequestIDOk returns a tuple with the SupplierRequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierRequestID

`func (o *GetRequestsId200Response) SetSupplierRequestID(v string)`

SetSupplierRequestID sets SupplierRequestID field to given value.

### HasSupplierRequestID

`func (o *GetRequestsId200Response) HasSupplierRequestID() bool`

HasSupplierRequestID returns a boolean if a field has been set.

### SetSupplierRequestIDNil

`func (o *GetRequestsId200Response) SetSupplierRequestIDNil(b bool)`

 SetSupplierRequestIDNil sets the value for SupplierRequestID to be an explicit nil

### UnsetSupplierRequestID
`func (o *GetRequestsId200Response) UnsetSupplierRequestID()`

UnsetSupplierRequestID ensures that no value is present for SupplierRequestID, not even an explicit nil
### GetTags

`func (o *GetRequestsId200Response) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetRequestsId200Response) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetRequestsId200Response) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetRequestsId200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTask

`func (o *GetRequestsId200Response) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *GetRequestsId200Response) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *GetRequestsId200Response) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *GetRequestsId200Response) HasTask() bool`

HasTask returns a boolean if a field has been set.

### SetTaskNil

`func (o *GetRequestsId200Response) SetTaskNil(b bool)`

 SetTaskNil sets the value for Task to be an explicit nil

### UnsetTask
`func (o *GetRequestsId200Response) UnsetTask()`

UnsetTask ensures that no value is present for Task, not even an explicit nil
### GetTeam

`func (o *GetRequestsId200Response) GetTeam() GetRequestsId200ResponseCreatedBy`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GetRequestsId200Response) GetTeamOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GetRequestsId200Response) SetTeam(v GetRequestsId200ResponseCreatedBy)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GetRequestsId200Response) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetTemplate

`func (o *GetRequestsId200Response) GetTemplate() GetRequestsId200ResponseTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetRequestsId200Response) GetTemplateOk() (*GetRequestsId200ResponseTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetRequestsId200Response) SetTemplate(v GetRequestsId200ResponseTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetRequestsId200Response) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetRequestsId200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetRequestsId200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetRequestsId200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetRequestsId200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrgent

`func (o *GetRequestsId200Response) GetUrgent() bool`

GetUrgent returns the Urgent field if non-nil, zero value otherwise.

### GetUrgentOk

`func (o *GetRequestsId200Response) GetUrgentOk() (*bool, bool)`

GetUrgentOk returns a tuple with the Urgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgent

`func (o *GetRequestsId200Response) SetUrgent(v bool)`

SetUrgent sets Urgent field to given value.

### HasUrgent

`func (o *GetRequestsId200Response) HasUrgent() bool`

HasUrgent returns a boolean if a field has been set.

### GetWaitingUntil

`func (o *GetRequestsId200Response) GetWaitingUntil() string`

GetWaitingUntil returns the WaitingUntil field if non-nil, zero value otherwise.

### GetWaitingUntilOk

`func (o *GetRequestsId200Response) GetWaitingUntilOk() (*string, bool)`

GetWaitingUntilOk returns a tuple with the WaitingUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingUntil

`func (o *GetRequestsId200Response) SetWaitingUntil(v string)`

SetWaitingUntil sets WaitingUntil field to given value.

### HasWaitingUntil

`func (o *GetRequestsId200Response) HasWaitingUntil() bool`

HasWaitingUntil returns a boolean if a field has been set.

### SetWaitingUntilNil

`func (o *GetRequestsId200Response) SetWaitingUntilNil(b bool)`

 SetWaitingUntilNil sets the value for WaitingUntil to be an explicit nil

### UnsetWaitingUntil
`func (o *GetRequestsId200Response) UnsetWaitingUntil()`

UnsetWaitingUntil ensures that no value is present for WaitingUntil, not even an explicit nil
### GetWorkflow

`func (o *GetRequestsId200Response) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *GetRequestsId200Response) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *GetRequestsId200Response) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *GetRequestsId200Response) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### SetWorkflowNil

`func (o *GetRequestsId200Response) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *GetRequestsId200Response) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


