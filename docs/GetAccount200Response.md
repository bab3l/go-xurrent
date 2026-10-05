# GetAccount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**DirectoryAccount** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**Owner** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**StartOfWeek** | Pointer to **string** |  | [optional] 
**TimeFormat24h** | Pointer to **bool** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAccount200Response

`func NewGetAccount200Response() *GetAccount200Response`

NewGetAccount200Response instantiates a new GetAccount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccount200ResponseWithDefaults

`func NewGetAccount200ResponseWithDefaults() *GetAccount200Response`

NewGetAccount200ResponseWithDefaults instantiates a new GetAccount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *GetAccount200Response) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetAccount200Response) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetAccount200Response) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetAccount200Response) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDirectoryAccount

`func (o *GetAccount200Response) GetDirectoryAccount() GetRequestsId200ResponseAccount`

GetDirectoryAccount returns the DirectoryAccount field if non-nil, zero value otherwise.

### GetDirectoryAccountOk

`func (o *GetAccount200Response) GetDirectoryAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetDirectoryAccountOk returns a tuple with the DirectoryAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryAccount

`func (o *GetAccount200Response) SetDirectoryAccount(v GetRequestsId200ResponseAccount)`

SetDirectoryAccount sets DirectoryAccount field to given value.

### HasDirectoryAccount

`func (o *GetAccount200Response) HasDirectoryAccount() bool`

HasDirectoryAccount returns a boolean if a field has been set.

### GetId

`func (o *GetAccount200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAccount200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAccount200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetAccount200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocale

`func (o *GetAccount200Response) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *GetAccount200Response) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *GetAccount200Response) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *GetAccount200Response) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetName

`func (o *GetAccount200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAccount200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAccount200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAccount200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *GetAccount200Response) GetOrganization() GetRequestsId200ResponseCreatedBy`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GetAccount200Response) GetOrganizationOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GetAccount200Response) SetOrganization(v GetRequestsId200ResponseCreatedBy)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *GetAccount200Response) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *GetAccount200Response) GetOwner() GetRequestsId200ResponseCreatedBy`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetAccount200Response) GetOwnerOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetAccount200Response) SetOwner(v GetRequestsId200ResponseCreatedBy)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetAccount200Response) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPlan

`func (o *GetAccount200Response) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetAccount200Response) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetAccount200Response) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *GetAccount200Response) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetStartOfWeek

`func (o *GetAccount200Response) GetStartOfWeek() string`

GetStartOfWeek returns the StartOfWeek field if non-nil, zero value otherwise.

### GetStartOfWeekOk

`func (o *GetAccount200Response) GetStartOfWeekOk() (*string, bool)`

GetStartOfWeekOk returns a tuple with the StartOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOfWeek

`func (o *GetAccount200Response) SetStartOfWeek(v string)`

SetStartOfWeek sets StartOfWeek field to given value.

### HasStartOfWeek

`func (o *GetAccount200Response) HasStartOfWeek() bool`

HasStartOfWeek returns a boolean if a field has been set.

### GetTimeFormat24h

`func (o *GetAccount200Response) GetTimeFormat24h() bool`

GetTimeFormat24h returns the TimeFormat24h field if non-nil, zero value otherwise.

### GetTimeFormat24hOk

`func (o *GetAccount200Response) GetTimeFormat24hOk() (*bool, bool)`

GetTimeFormat24hOk returns a tuple with the TimeFormat24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFormat24h

`func (o *GetAccount200Response) SetTimeFormat24h(v bool)`

SetTimeFormat24h sets TimeFormat24h field to given value.

### HasTimeFormat24h

`func (o *GetAccount200Response) HasTimeFormat24h() bool`

HasTimeFormat24h returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetAccount200Response) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetAccount200Response) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetAccount200Response) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetAccount200Response) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUrl

`func (o *GetAccount200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetAccount200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetAccount200Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetAccount200Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


