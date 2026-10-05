# GetSlasId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**ActivityIDCase** | Pointer to **NullableString** |  | [optional] 
**ActivityIDHigh** | Pointer to **NullableString** |  | [optional] 
**ActivityIDLow** | Pointer to **NullableString** |  | [optional] 
**ActivityIDMedium** | Pointer to **NullableString** |  | [optional] 
**ActivityIDRfc** | Pointer to **NullableString** |  | [optional] 
**ActivityIDRfi** | Pointer to **NullableString** |  | [optional] 
**ActivityIDTop** | Pointer to **NullableString** |  | [optional] 
**AgreementID** | Pointer to **NullableString** |  | [optional] 
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Coverage** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Customer** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**CustomerAccount** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**CustomerRep** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**ExpiryDate** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**NoticeDate** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**ServiceInstance** | Pointer to [**GetRequestsId200ResponseServiceInstance**](GetRequestsId200ResponseServiceInstance.md) |  | [optional] 
**ServiceLevelManager** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**ServiceOffering** | Pointer to [**GetSlasId200ResponseServiceOffering**](GetSlasId200ResponseServiceOffering.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UseKnowledgeFromServiceProvider** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetSlasId200Response

`func NewGetSlasId200Response() *GetSlasId200Response`

NewGetSlasId200Response instantiates a new GetSlasId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSlasId200ResponseWithDefaults

`func NewGetSlasId200ResponseWithDefaults() *GetSlasId200Response`

NewGetSlasId200ResponseWithDefaults instantiates a new GetSlasId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetSlasId200Response) GetAccount() GetRequestsId200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetSlasId200Response) GetAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetSlasId200Response) SetAccount(v GetRequestsId200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetSlasId200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActivityIDCase

`func (o *GetSlasId200Response) GetActivityIDCase() string`

GetActivityIDCase returns the ActivityIDCase field if non-nil, zero value otherwise.

### GetActivityIDCaseOk

`func (o *GetSlasId200Response) GetActivityIDCaseOk() (*string, bool)`

GetActivityIDCaseOk returns a tuple with the ActivityIDCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIDCase

`func (o *GetSlasId200Response) SetActivityIDCase(v string)`

SetActivityIDCase sets ActivityIDCase field to given value.

### HasActivityIDCase

`func (o *GetSlasId200Response) HasActivityIDCase() bool`

HasActivityIDCase returns a boolean if a field has been set.

### SetActivityIDCaseNil

`func (o *GetSlasId200Response) SetActivityIDCaseNil(b bool)`

 SetActivityIDCaseNil sets the value for ActivityIDCase to be an explicit nil

### UnsetActivityIDCase
`func (o *GetSlasId200Response) UnsetActivityIDCase()`

UnsetActivityIDCase ensures that no value is present for ActivityIDCase, not even an explicit nil
### GetActivityIDHigh

`func (o *GetSlasId200Response) GetActivityIDHigh() string`

GetActivityIDHigh returns the ActivityIDHigh field if non-nil, zero value otherwise.

### GetActivityIDHighOk

`func (o *GetSlasId200Response) GetActivityIDHighOk() (*string, bool)`

GetActivityIDHighOk returns a tuple with the ActivityIDHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIDHigh

`func (o *GetSlasId200Response) SetActivityIDHigh(v string)`

SetActivityIDHigh sets ActivityIDHigh field to given value.

### HasActivityIDHigh

`func (o *GetSlasId200Response) HasActivityIDHigh() bool`

HasActivityIDHigh returns a boolean if a field has been set.

### SetActivityIDHighNil

`func (o *GetSlasId200Response) SetActivityIDHighNil(b bool)`

 SetActivityIDHighNil sets the value for ActivityIDHigh to be an explicit nil

### UnsetActivityIDHigh
`func (o *GetSlasId200Response) UnsetActivityIDHigh()`

UnsetActivityIDHigh ensures that no value is present for ActivityIDHigh, not even an explicit nil
### GetActivityIDLow

`func (o *GetSlasId200Response) GetActivityIDLow() string`

GetActivityIDLow returns the ActivityIDLow field if non-nil, zero value otherwise.

### GetActivityIDLowOk

`func (o *GetSlasId200Response) GetActivityIDLowOk() (*string, bool)`

GetActivityIDLowOk returns a tuple with the ActivityIDLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIDLow

`func (o *GetSlasId200Response) SetActivityIDLow(v string)`

SetActivityIDLow sets ActivityIDLow field to given value.

### HasActivityIDLow

`func (o *GetSlasId200Response) HasActivityIDLow() bool`

HasActivityIDLow returns a boolean if a field has been set.

### SetActivityIDLowNil

`func (o *GetSlasId200Response) SetActivityIDLowNil(b bool)`

 SetActivityIDLowNil sets the value for ActivityIDLow to be an explicit nil

### UnsetActivityIDLow
`func (o *GetSlasId200Response) UnsetActivityIDLow()`

UnsetActivityIDLow ensures that no value is present for ActivityIDLow, not even an explicit nil
### GetActivityIDMedium

`func (o *GetSlasId200Response) GetActivityIDMedium() string`

GetActivityIDMedium returns the ActivityIDMedium field if non-nil, zero value otherwise.

### GetActivityIDMediumOk

`func (o *GetSlasId200Response) GetActivityIDMediumOk() (*string, bool)`

GetActivityIDMediumOk returns a tuple with the ActivityIDMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIDMedium

`func (o *GetSlasId200Response) SetActivityIDMedium(v string)`

SetActivityIDMedium sets ActivityIDMedium field to given value.

### HasActivityIDMedium

`func (o *GetSlasId200Response) HasActivityIDMedium() bool`

HasActivityIDMedium returns a boolean if a field has been set.

### SetActivityIDMediumNil

`func (o *GetSlasId200Response) SetActivityIDMediumNil(b bool)`

 SetActivityIDMediumNil sets the value for ActivityIDMedium to be an explicit nil

### UnsetActivityIDMedium
`func (o *GetSlasId200Response) UnsetActivityIDMedium()`

UnsetActivityIDMedium ensures that no value is present for ActivityIDMedium, not even an explicit nil
### GetActivityIDRfc

`func (o *GetSlasId200Response) GetActivityIDRfc() string`

GetActivityIDRfc returns the ActivityIDRfc field if non-nil, zero value otherwise.

### GetActivityIDRfcOk

`func (o *GetSlasId200Response) GetActivityIDRfcOk() (*string, bool)`

GetActivityIDRfcOk returns a tuple with the ActivityIDRfc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIDRfc

`func (o *GetSlasId200Response) SetActivityIDRfc(v string)`

SetActivityIDRfc sets ActivityIDRfc field to given value.

### HasActivityIDRfc

`func (o *GetSlasId200Response) HasActivityIDRfc() bool`

HasActivityIDRfc returns a boolean if a field has been set.

### SetActivityIDRfcNil

`func (o *GetSlasId200Response) SetActivityIDRfcNil(b bool)`

 SetActivityIDRfcNil sets the value for ActivityIDRfc to be an explicit nil

### UnsetActivityIDRfc
`func (o *GetSlasId200Response) UnsetActivityIDRfc()`

UnsetActivityIDRfc ensures that no value is present for ActivityIDRfc, not even an explicit nil
### GetActivityIDRfi

`func (o *GetSlasId200Response) GetActivityIDRfi() string`

GetActivityIDRfi returns the ActivityIDRfi field if non-nil, zero value otherwise.

### GetActivityIDRfiOk

`func (o *GetSlasId200Response) GetActivityIDRfiOk() (*string, bool)`

GetActivityIDRfiOk returns a tuple with the ActivityIDRfi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIDRfi

`func (o *GetSlasId200Response) SetActivityIDRfi(v string)`

SetActivityIDRfi sets ActivityIDRfi field to given value.

### HasActivityIDRfi

`func (o *GetSlasId200Response) HasActivityIDRfi() bool`

HasActivityIDRfi returns a boolean if a field has been set.

### SetActivityIDRfiNil

`func (o *GetSlasId200Response) SetActivityIDRfiNil(b bool)`

 SetActivityIDRfiNil sets the value for ActivityIDRfi to be an explicit nil

### UnsetActivityIDRfi
`func (o *GetSlasId200Response) UnsetActivityIDRfi()`

UnsetActivityIDRfi ensures that no value is present for ActivityIDRfi, not even an explicit nil
### GetActivityIDTop

`func (o *GetSlasId200Response) GetActivityIDTop() string`

GetActivityIDTop returns the ActivityIDTop field if non-nil, zero value otherwise.

### GetActivityIDTopOk

`func (o *GetSlasId200Response) GetActivityIDTopOk() (*string, bool)`

GetActivityIDTopOk returns a tuple with the ActivityIDTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIDTop

`func (o *GetSlasId200Response) SetActivityIDTop(v string)`

SetActivityIDTop sets ActivityIDTop field to given value.

### HasActivityIDTop

`func (o *GetSlasId200Response) HasActivityIDTop() bool`

HasActivityIDTop returns a boolean if a field has been set.

### SetActivityIDTopNil

`func (o *GetSlasId200Response) SetActivityIDTopNil(b bool)`

 SetActivityIDTopNil sets the value for ActivityIDTop to be an explicit nil

### UnsetActivityIDTop
`func (o *GetSlasId200Response) UnsetActivityIDTop()`

UnsetActivityIDTop ensures that no value is present for ActivityIDTop, not even an explicit nil
### GetAgreementID

`func (o *GetSlasId200Response) GetAgreementID() string`

GetAgreementID returns the AgreementID field if non-nil, zero value otherwise.

### GetAgreementIDOk

`func (o *GetSlasId200Response) GetAgreementIDOk() (*string, bool)`

GetAgreementIDOk returns a tuple with the AgreementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementID

`func (o *GetSlasId200Response) SetAgreementID(v string)`

SetAgreementID sets AgreementID field to given value.

### HasAgreementID

`func (o *GetSlasId200Response) HasAgreementID() bool`

HasAgreementID returns a boolean if a field has been set.

### SetAgreementIDNil

`func (o *GetSlasId200Response) SetAgreementIDNil(b bool)`

 SetAgreementIDNil sets the value for AgreementID to be an explicit nil

### UnsetAgreementID
`func (o *GetSlasId200Response) UnsetAgreementID()`

UnsetAgreementID ensures that no value is present for AgreementID, not even an explicit nil
### GetAttachments

`func (o *GetSlasId200Response) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetSlasId200Response) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetSlasId200Response) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetSlasId200Response) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCoverage

`func (o *GetSlasId200Response) GetCoverage() string`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *GetSlasId200Response) GetCoverageOk() (*string, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *GetSlasId200Response) SetCoverage(v string)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *GetSlasId200Response) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetSlasId200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetSlasId200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetSlasId200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetSlasId200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomer

`func (o *GetSlasId200Response) GetCustomer() GetRequestsId200ResponseCreatedBy`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *GetSlasId200Response) GetCustomerOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *GetSlasId200Response) SetCustomer(v GetRequestsId200ResponseCreatedBy)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *GetSlasId200Response) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerAccount

`func (o *GetSlasId200Response) GetCustomerAccount() GetRequestsId200ResponseAccount`

GetCustomerAccount returns the CustomerAccount field if non-nil, zero value otherwise.

### GetCustomerAccountOk

`func (o *GetSlasId200Response) GetCustomerAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetCustomerAccountOk returns a tuple with the CustomerAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAccount

`func (o *GetSlasId200Response) SetCustomerAccount(v GetRequestsId200ResponseAccount)`

SetCustomerAccount sets CustomerAccount field to given value.

### HasCustomerAccount

`func (o *GetSlasId200Response) HasCustomerAccount() bool`

HasCustomerAccount returns a boolean if a field has been set.

### GetCustomerRep

`func (o *GetSlasId200Response) GetCustomerRep() GetRequestsId200ResponseCreatedBy`

GetCustomerRep returns the CustomerRep field if non-nil, zero value otherwise.

### GetCustomerRepOk

`func (o *GetSlasId200Response) GetCustomerRepOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetCustomerRepOk returns a tuple with the CustomerRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRep

`func (o *GetSlasId200Response) SetCustomerRep(v GetRequestsId200ResponseCreatedBy)`

SetCustomerRep sets CustomerRep field to given value.

### HasCustomerRep

`func (o *GetSlasId200Response) HasCustomerRep() bool`

HasCustomerRep returns a boolean if a field has been set.

### GetExpiryDate

`func (o *GetSlasId200Response) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *GetSlasId200Response) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *GetSlasId200Response) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *GetSlasId200Response) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *GetSlasId200Response) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *GetSlasId200Response) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetId

`func (o *GetSlasId200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSlasId200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSlasId200Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetSlasId200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetSlasId200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSlasId200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSlasId200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSlasId200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeID

`func (o *GetSlasId200Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *GetSlasId200Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *GetSlasId200Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *GetSlasId200Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetNoticeDate

`func (o *GetSlasId200Response) GetNoticeDate() string`

GetNoticeDate returns the NoticeDate field if non-nil, zero value otherwise.

### GetNoticeDateOk

`func (o *GetSlasId200Response) GetNoticeDateOk() (*string, bool)`

GetNoticeDateOk returns a tuple with the NoticeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeDate

`func (o *GetSlasId200Response) SetNoticeDate(v string)`

SetNoticeDate sets NoticeDate field to given value.

### HasNoticeDate

`func (o *GetSlasId200Response) HasNoticeDate() bool`

HasNoticeDate returns a boolean if a field has been set.

### SetNoticeDateNil

`func (o *GetSlasId200Response) SetNoticeDateNil(b bool)`

 SetNoticeDateNil sets the value for NoticeDate to be an explicit nil

### UnsetNoticeDate
`func (o *GetSlasId200Response) UnsetNoticeDate()`

UnsetNoticeDate ensures that no value is present for NoticeDate, not even an explicit nil
### GetRemarks

`func (o *GetSlasId200Response) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *GetSlasId200Response) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *GetSlasId200Response) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *GetSlasId200Response) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *GetSlasId200Response) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *GetSlasId200Response) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetServiceInstance

`func (o *GetSlasId200Response) GetServiceInstance() GetRequestsId200ResponseServiceInstance`

GetServiceInstance returns the ServiceInstance field if non-nil, zero value otherwise.

### GetServiceInstanceOk

`func (o *GetSlasId200Response) GetServiceInstanceOk() (*GetRequestsId200ResponseServiceInstance, bool)`

GetServiceInstanceOk returns a tuple with the ServiceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstance

`func (o *GetSlasId200Response) SetServiceInstance(v GetRequestsId200ResponseServiceInstance)`

SetServiceInstance sets ServiceInstance field to given value.

### HasServiceInstance

`func (o *GetSlasId200Response) HasServiceInstance() bool`

HasServiceInstance returns a boolean if a field has been set.

### GetServiceLevelManager

`func (o *GetSlasId200Response) GetServiceLevelManager() GetRequestsId200ResponseCreatedBy`

GetServiceLevelManager returns the ServiceLevelManager field if non-nil, zero value otherwise.

### GetServiceLevelManagerOk

`func (o *GetSlasId200Response) GetServiceLevelManagerOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetServiceLevelManagerOk returns a tuple with the ServiceLevelManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelManager

`func (o *GetSlasId200Response) SetServiceLevelManager(v GetRequestsId200ResponseCreatedBy)`

SetServiceLevelManager sets ServiceLevelManager field to given value.

### HasServiceLevelManager

`func (o *GetSlasId200Response) HasServiceLevelManager() bool`

HasServiceLevelManager returns a boolean if a field has been set.

### GetServiceOffering

`func (o *GetSlasId200Response) GetServiceOffering() GetSlasId200ResponseServiceOffering`

GetServiceOffering returns the ServiceOffering field if non-nil, zero value otherwise.

### GetServiceOfferingOk

`func (o *GetSlasId200Response) GetServiceOfferingOk() (*GetSlasId200ResponseServiceOffering, bool)`

GetServiceOfferingOk returns a tuple with the ServiceOffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOffering

`func (o *GetSlasId200Response) SetServiceOffering(v GetSlasId200ResponseServiceOffering)`

SetServiceOffering sets ServiceOffering field to given value.

### HasServiceOffering

`func (o *GetSlasId200Response) HasServiceOffering() bool`

HasServiceOffering returns a boolean if a field has been set.

### GetSource

`func (o *GetSlasId200Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetSlasId200Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetSlasId200Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetSlasId200Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *GetSlasId200Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *GetSlasId200Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *GetSlasId200Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *GetSlasId200Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *GetSlasId200Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *GetSlasId200Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetStartDate

`func (o *GetSlasId200Response) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetSlasId200Response) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetSlasId200Response) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetSlasId200Response) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *GetSlasId200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSlasId200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSlasId200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSlasId200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetSlasId200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetSlasId200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetSlasId200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetSlasId200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUseKnowledgeFromServiceProvider

`func (o *GetSlasId200Response) GetUseKnowledgeFromServiceProvider() bool`

GetUseKnowledgeFromServiceProvider returns the UseKnowledgeFromServiceProvider field if non-nil, zero value otherwise.

### GetUseKnowledgeFromServiceProviderOk

`func (o *GetSlasId200Response) GetUseKnowledgeFromServiceProviderOk() (*bool, bool)`

GetUseKnowledgeFromServiceProviderOk returns a tuple with the UseKnowledgeFromServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKnowledgeFromServiceProvider

`func (o *GetSlasId200Response) SetUseKnowledgeFromServiceProvider(v bool)`

SetUseKnowledgeFromServiceProvider sets UseKnowledgeFromServiceProvider field to given value.

### HasUseKnowledgeFromServiceProvider

`func (o *GetSlasId200Response) HasUseKnowledgeFromServiceProvider() bool`

HasUseKnowledgeFromServiceProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


