# PostProducts201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Brand** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **NullableString** |  | [optional] 
**DepreciationMethod** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**FinancialOwner** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Model** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**PictureUri** | Pointer to **NullableString** |  | [optional] 
**ProductID** | Pointer to **NullableString** |  | [optional] 
**Rate** | Pointer to **NullableString** |  | [optional] 
**Recurrence** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**RuleSet** | Pointer to **string** |  | [optional] 
**SalvageValue** | Pointer to **NullableString** |  | [optional] 
**SalvageValueCurrency** | Pointer to **NullableString** |  | [optional] 
**Service** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**Supplier** | Pointer to **NullableString** |  | [optional] 
**SupportTeam** | Pointer to **NullableString** |  | [optional] 
**UiExtension** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UsefulLife** | Pointer to **NullableString** |  | [optional] 
**WorkflowManager** | Pointer to **NullableString** |  | [optional] 
**WorkflowTemplate** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPostProducts201Response

`func NewPostProducts201Response() *PostProducts201Response`

NewPostProducts201Response instantiates a new PostProducts201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostProducts201ResponseWithDefaults

`func NewPostProducts201ResponseWithDefaults() *PostProducts201Response`

NewPostProducts201ResponseWithDefaults instantiates a new PostProducts201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *PostProducts201Response) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PostProducts201Response) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PostProducts201Response) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *PostProducts201Response) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBrand

`func (o *PostProducts201Response) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *PostProducts201Response) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *PostProducts201Response) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *PostProducts201Response) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetCategory

`func (o *PostProducts201Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PostProducts201Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PostProducts201Response) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PostProducts201Response) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PostProducts201Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostProducts201Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostProducts201Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PostProducts201Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomFields

`func (o *PostProducts201Response) GetCustomFields() string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PostProducts201Response) GetCustomFieldsOk() (*string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PostProducts201Response) SetCustomFields(v string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PostProducts201Response) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *PostProducts201Response) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *PostProducts201Response) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDepreciationMethod

`func (o *PostProducts201Response) GetDepreciationMethod() string`

GetDepreciationMethod returns the DepreciationMethod field if non-nil, zero value otherwise.

### GetDepreciationMethodOk

`func (o *PostProducts201Response) GetDepreciationMethodOk() (*string, bool)`

GetDepreciationMethodOk returns a tuple with the DepreciationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationMethod

`func (o *PostProducts201Response) SetDepreciationMethod(v string)`

SetDepreciationMethod sets DepreciationMethod field to given value.

### HasDepreciationMethod

`func (o *PostProducts201Response) HasDepreciationMethod() bool`

HasDepreciationMethod returns a boolean if a field has been set.

### SetDepreciationMethodNil

`func (o *PostProducts201Response) SetDepreciationMethodNil(b bool)`

 SetDepreciationMethodNil sets the value for DepreciationMethod to be an explicit nil

### UnsetDepreciationMethod
`func (o *PostProducts201Response) UnsetDepreciationMethod()`

UnsetDepreciationMethod ensures that no value is present for DepreciationMethod, not even an explicit nil
### GetDisabled

`func (o *PostProducts201Response) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PostProducts201Response) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PostProducts201Response) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PostProducts201Response) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetFinancialOwner

`func (o *PostProducts201Response) GetFinancialOwner() string`

GetFinancialOwner returns the FinancialOwner field if non-nil, zero value otherwise.

### GetFinancialOwnerOk

`func (o *PostProducts201Response) GetFinancialOwnerOk() (*string, bool)`

GetFinancialOwnerOk returns a tuple with the FinancialOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialOwner

`func (o *PostProducts201Response) SetFinancialOwner(v string)`

SetFinancialOwner sets FinancialOwner field to given value.

### HasFinancialOwner

`func (o *PostProducts201Response) HasFinancialOwner() bool`

HasFinancialOwner returns a boolean if a field has been set.

### SetFinancialOwnerNil

`func (o *PostProducts201Response) SetFinancialOwnerNil(b bool)`

 SetFinancialOwnerNil sets the value for FinancialOwner to be an explicit nil

### UnsetFinancialOwner
`func (o *PostProducts201Response) UnsetFinancialOwner()`

UnsetFinancialOwner ensures that no value is present for FinancialOwner, not even an explicit nil
### GetId

`func (o *PostProducts201Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostProducts201Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostProducts201Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *PostProducts201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *PostProducts201Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PostProducts201Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PostProducts201Response) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PostProducts201Response) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *PostProducts201Response) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *PostProducts201Response) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetName

`func (o *PostProducts201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostProducts201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostProducts201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostProducts201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeID

`func (o *PostProducts201Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *PostProducts201Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *PostProducts201Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *PostProducts201Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetPictureUri

`func (o *PostProducts201Response) GetPictureUri() string`

GetPictureUri returns the PictureUri field if non-nil, zero value otherwise.

### GetPictureUriOk

`func (o *PostProducts201Response) GetPictureUriOk() (*string, bool)`

GetPictureUriOk returns a tuple with the PictureUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureUri

`func (o *PostProducts201Response) SetPictureUri(v string)`

SetPictureUri sets PictureUri field to given value.

### HasPictureUri

`func (o *PostProducts201Response) HasPictureUri() bool`

HasPictureUri returns a boolean if a field has been set.

### SetPictureUriNil

`func (o *PostProducts201Response) SetPictureUriNil(b bool)`

 SetPictureUriNil sets the value for PictureUri to be an explicit nil

### UnsetPictureUri
`func (o *PostProducts201Response) UnsetPictureUri()`

UnsetPictureUri ensures that no value is present for PictureUri, not even an explicit nil
### GetProductID

`func (o *PostProducts201Response) GetProductID() string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *PostProducts201Response) GetProductIDOk() (*string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *PostProducts201Response) SetProductID(v string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *PostProducts201Response) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### SetProductIDNil

`func (o *PostProducts201Response) SetProductIDNil(b bool)`

 SetProductIDNil sets the value for ProductID to be an explicit nil

### UnsetProductID
`func (o *PostProducts201Response) UnsetProductID()`

UnsetProductID ensures that no value is present for ProductID, not even an explicit nil
### GetRate

`func (o *PostProducts201Response) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *PostProducts201Response) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *PostProducts201Response) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *PostProducts201Response) HasRate() bool`

HasRate returns a boolean if a field has been set.

### SetRateNil

`func (o *PostProducts201Response) SetRateNil(b bool)`

 SetRateNil sets the value for Rate to be an explicit nil

### UnsetRate
`func (o *PostProducts201Response) UnsetRate()`

UnsetRate ensures that no value is present for Rate, not even an explicit nil
### GetRecurrence

`func (o *PostProducts201Response) GetRecurrence() string`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *PostProducts201Response) GetRecurrenceOk() (*string, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *PostProducts201Response) SetRecurrence(v string)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *PostProducts201Response) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### SetRecurrenceNil

`func (o *PostProducts201Response) SetRecurrenceNil(b bool)`

 SetRecurrenceNil sets the value for Recurrence to be an explicit nil

### UnsetRecurrence
`func (o *PostProducts201Response) UnsetRecurrence()`

UnsetRecurrence ensures that no value is present for Recurrence, not even an explicit nil
### GetRemarks

`func (o *PostProducts201Response) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *PostProducts201Response) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *PostProducts201Response) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *PostProducts201Response) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetRuleSet

`func (o *PostProducts201Response) GetRuleSet() string`

GetRuleSet returns the RuleSet field if non-nil, zero value otherwise.

### GetRuleSetOk

`func (o *PostProducts201Response) GetRuleSetOk() (*string, bool)`

GetRuleSetOk returns a tuple with the RuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSet

`func (o *PostProducts201Response) SetRuleSet(v string)`

SetRuleSet sets RuleSet field to given value.

### HasRuleSet

`func (o *PostProducts201Response) HasRuleSet() bool`

HasRuleSet returns a boolean if a field has been set.

### GetSalvageValue

`func (o *PostProducts201Response) GetSalvageValue() string`

GetSalvageValue returns the SalvageValue field if non-nil, zero value otherwise.

### GetSalvageValueOk

`func (o *PostProducts201Response) GetSalvageValueOk() (*string, bool)`

GetSalvageValueOk returns a tuple with the SalvageValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalvageValue

`func (o *PostProducts201Response) SetSalvageValue(v string)`

SetSalvageValue sets SalvageValue field to given value.

### HasSalvageValue

`func (o *PostProducts201Response) HasSalvageValue() bool`

HasSalvageValue returns a boolean if a field has been set.

### SetSalvageValueNil

`func (o *PostProducts201Response) SetSalvageValueNil(b bool)`

 SetSalvageValueNil sets the value for SalvageValue to be an explicit nil

### UnsetSalvageValue
`func (o *PostProducts201Response) UnsetSalvageValue()`

UnsetSalvageValue ensures that no value is present for SalvageValue, not even an explicit nil
### GetSalvageValueCurrency

`func (o *PostProducts201Response) GetSalvageValueCurrency() string`

GetSalvageValueCurrency returns the SalvageValueCurrency field if non-nil, zero value otherwise.

### GetSalvageValueCurrencyOk

`func (o *PostProducts201Response) GetSalvageValueCurrencyOk() (*string, bool)`

GetSalvageValueCurrencyOk returns a tuple with the SalvageValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalvageValueCurrency

`func (o *PostProducts201Response) SetSalvageValueCurrency(v string)`

SetSalvageValueCurrency sets SalvageValueCurrency field to given value.

### HasSalvageValueCurrency

`func (o *PostProducts201Response) HasSalvageValueCurrency() bool`

HasSalvageValueCurrency returns a boolean if a field has been set.

### SetSalvageValueCurrencyNil

`func (o *PostProducts201Response) SetSalvageValueCurrencyNil(b bool)`

 SetSalvageValueCurrencyNil sets the value for SalvageValueCurrency to be an explicit nil

### UnsetSalvageValueCurrency
`func (o *PostProducts201Response) UnsetSalvageValueCurrency()`

UnsetSalvageValueCurrency ensures that no value is present for SalvageValueCurrency, not even an explicit nil
### GetService

`func (o *PostProducts201Response) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PostProducts201Response) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PostProducts201Response) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *PostProducts201Response) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *PostProducts201Response) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *PostProducts201Response) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetSource

`func (o *PostProducts201Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PostProducts201Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PostProducts201Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PostProducts201Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *PostProducts201Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *PostProducts201Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *PostProducts201Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *PostProducts201Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *PostProducts201Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *PostProducts201Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetSupplier

`func (o *PostProducts201Response) GetSupplier() string`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *PostProducts201Response) GetSupplierOk() (*string, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *PostProducts201Response) SetSupplier(v string)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *PostProducts201Response) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### SetSupplierNil

`func (o *PostProducts201Response) SetSupplierNil(b bool)`

 SetSupplierNil sets the value for Supplier to be an explicit nil

### UnsetSupplier
`func (o *PostProducts201Response) UnsetSupplier()`

UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
### GetSupportTeam

`func (o *PostProducts201Response) GetSupportTeam() string`

GetSupportTeam returns the SupportTeam field if non-nil, zero value otherwise.

### GetSupportTeamOk

`func (o *PostProducts201Response) GetSupportTeamOk() (*string, bool)`

GetSupportTeamOk returns a tuple with the SupportTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTeam

`func (o *PostProducts201Response) SetSupportTeam(v string)`

SetSupportTeam sets SupportTeam field to given value.

### HasSupportTeam

`func (o *PostProducts201Response) HasSupportTeam() bool`

HasSupportTeam returns a boolean if a field has been set.

### SetSupportTeamNil

`func (o *PostProducts201Response) SetSupportTeamNil(b bool)`

 SetSupportTeamNil sets the value for SupportTeam to be an explicit nil

### UnsetSupportTeam
`func (o *PostProducts201Response) UnsetSupportTeam()`

UnsetSupportTeam ensures that no value is present for SupportTeam, not even an explicit nil
### GetUiExtension

`func (o *PostProducts201Response) GetUiExtension() string`

GetUiExtension returns the UiExtension field if non-nil, zero value otherwise.

### GetUiExtensionOk

`func (o *PostProducts201Response) GetUiExtensionOk() (*string, bool)`

GetUiExtensionOk returns a tuple with the UiExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiExtension

`func (o *PostProducts201Response) SetUiExtension(v string)`

SetUiExtension sets UiExtension field to given value.

### HasUiExtension

`func (o *PostProducts201Response) HasUiExtension() bool`

HasUiExtension returns a boolean if a field has been set.

### SetUiExtensionNil

`func (o *PostProducts201Response) SetUiExtensionNil(b bool)`

 SetUiExtensionNil sets the value for UiExtension to be an explicit nil

### UnsetUiExtension
`func (o *PostProducts201Response) UnsetUiExtension()`

UnsetUiExtension ensures that no value is present for UiExtension, not even an explicit nil
### GetUpdatedAt

`func (o *PostProducts201Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PostProducts201Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PostProducts201Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PostProducts201Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsefulLife

`func (o *PostProducts201Response) GetUsefulLife() string`

GetUsefulLife returns the UsefulLife field if non-nil, zero value otherwise.

### GetUsefulLifeOk

`func (o *PostProducts201Response) GetUsefulLifeOk() (*string, bool)`

GetUsefulLifeOk returns a tuple with the UsefulLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsefulLife

`func (o *PostProducts201Response) SetUsefulLife(v string)`

SetUsefulLife sets UsefulLife field to given value.

### HasUsefulLife

`func (o *PostProducts201Response) HasUsefulLife() bool`

HasUsefulLife returns a boolean if a field has been set.

### SetUsefulLifeNil

`func (o *PostProducts201Response) SetUsefulLifeNil(b bool)`

 SetUsefulLifeNil sets the value for UsefulLife to be an explicit nil

### UnsetUsefulLife
`func (o *PostProducts201Response) UnsetUsefulLife()`

UnsetUsefulLife ensures that no value is present for UsefulLife, not even an explicit nil
### GetWorkflowManager

`func (o *PostProducts201Response) GetWorkflowManager() string`

GetWorkflowManager returns the WorkflowManager field if non-nil, zero value otherwise.

### GetWorkflowManagerOk

`func (o *PostProducts201Response) GetWorkflowManagerOk() (*string, bool)`

GetWorkflowManagerOk returns a tuple with the WorkflowManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowManager

`func (o *PostProducts201Response) SetWorkflowManager(v string)`

SetWorkflowManager sets WorkflowManager field to given value.

### HasWorkflowManager

`func (o *PostProducts201Response) HasWorkflowManager() bool`

HasWorkflowManager returns a boolean if a field has been set.

### SetWorkflowManagerNil

`func (o *PostProducts201Response) SetWorkflowManagerNil(b bool)`

 SetWorkflowManagerNil sets the value for WorkflowManager to be an explicit nil

### UnsetWorkflowManager
`func (o *PostProducts201Response) UnsetWorkflowManager()`

UnsetWorkflowManager ensures that no value is present for WorkflowManager, not even an explicit nil
### GetWorkflowTemplate

`func (o *PostProducts201Response) GetWorkflowTemplate() string`

GetWorkflowTemplate returns the WorkflowTemplate field if non-nil, zero value otherwise.

### GetWorkflowTemplateOk

`func (o *PostProducts201Response) GetWorkflowTemplateOk() (*string, bool)`

GetWorkflowTemplateOk returns a tuple with the WorkflowTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTemplate

`func (o *PostProducts201Response) SetWorkflowTemplate(v string)`

SetWorkflowTemplate sets WorkflowTemplate field to given value.

### HasWorkflowTemplate

`func (o *PostProducts201Response) HasWorkflowTemplate() bool`

HasWorkflowTemplate returns a boolean if a field has been set.

### SetWorkflowTemplateNil

`func (o *PostProducts201Response) SetWorkflowTemplateNil(b bool)`

 SetWorkflowTemplateNil sets the value for WorkflowTemplate to be an explicit nil

### UnsetWorkflowTemplate
`func (o *PostProducts201Response) UnsetWorkflowTemplate()`

UnsetWorkflowTemplate ensures that no value is present for WorkflowTemplate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


