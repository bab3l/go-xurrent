# PostPeople201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GetRequestsId200ResponseAccount**](GetRequestsId200ResponseAccount.md) |  | [optional] 
**Addresses** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**AuthenticationID** | Pointer to **NullableString** |  | [optional] 
**AutoTranslation** | Pointer to **bool** |  | [optional] 
**ColorMode** | Pointer to **string** |  | [optional] 
**Contacts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CostPerHour** | Pointer to **NullableString** |  | [optional] 
**CostPerHourCurrency** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**DoNotTranslateLanguages** | Pointer to **NullableString** |  | [optional] 
**EmployeeID** | Pointer to **NullableString** |  | [optional] 
**ExcludeTeamNotifications** | Pointer to **bool** |  | [optional] 
**Formats** | Pointer to [**PostPeople201ResponseFormats**](PostPeople201ResponseFormats.md) |  | [optional] 
**Guest** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Information** | Pointer to **NullableString** |  | [optional] 
**JobTitle** | Pointer to **NullableString** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**Manager** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**OauthPersonEnablement** | Pointer to **bool** |  | [optional] 
**Organization** | Pointer to [**GetRequestsId200ResponseCreatedBy**](GetRequestsId200ResponseCreatedBy.md) |  | [optional] 
**PictureUri** | Pointer to **NullableString** |  | [optional] 
**PlayPrivateChatSound** | Pointer to **string** |  | [optional] 
**PlaySupportChatSound** | Pointer to **string** |  | [optional] 
**PrimaryEmail** | Pointer to **string** |  | [optional] 
**SendEmailNotifications** | Pointer to **string** |  | [optional] 
**ShowNotificationPopup** | Pointer to **string** |  | [optional] 
**ShowPrivateChatPopup** | Pointer to **string** |  | [optional] 
**ShowSupportChatPopup** | Pointer to **string** |  | [optional] 
**Site** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**SupportID** | Pointer to **NullableString** |  | [optional] 
**TimeFormat24h** | Pointer to **bool** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**UiExtension** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Vip** | Pointer to **bool** |  | [optional] 
**WorkHours** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPostPeople201Response

`func NewPostPeople201Response() *PostPeople201Response`

NewPostPeople201Response instantiates a new PostPeople201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPeople201ResponseWithDefaults

`func NewPostPeople201ResponseWithDefaults() *PostPeople201Response`

NewPostPeople201ResponseWithDefaults instantiates a new PostPeople201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *PostPeople201Response) GetAccount() GetRequestsId200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PostPeople201Response) GetAccountOk() (*GetRequestsId200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PostPeople201Response) SetAccount(v GetRequestsId200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PostPeople201Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAddresses

`func (o *PostPeople201Response) GetAddresses() []map[string]interface{}`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PostPeople201Response) GetAddressesOk() (*[]map[string]interface{}, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PostPeople201Response) SetAddresses(v []map[string]interface{})`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *PostPeople201Response) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAttachments

`func (o *PostPeople201Response) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PostPeople201Response) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PostPeople201Response) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *PostPeople201Response) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetAuthenticationID

`func (o *PostPeople201Response) GetAuthenticationID() string`

GetAuthenticationID returns the AuthenticationID field if non-nil, zero value otherwise.

### GetAuthenticationIDOk

`func (o *PostPeople201Response) GetAuthenticationIDOk() (*string, bool)`

GetAuthenticationIDOk returns a tuple with the AuthenticationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationID

`func (o *PostPeople201Response) SetAuthenticationID(v string)`

SetAuthenticationID sets AuthenticationID field to given value.

### HasAuthenticationID

`func (o *PostPeople201Response) HasAuthenticationID() bool`

HasAuthenticationID returns a boolean if a field has been set.

### SetAuthenticationIDNil

`func (o *PostPeople201Response) SetAuthenticationIDNil(b bool)`

 SetAuthenticationIDNil sets the value for AuthenticationID to be an explicit nil

### UnsetAuthenticationID
`func (o *PostPeople201Response) UnsetAuthenticationID()`

UnsetAuthenticationID ensures that no value is present for AuthenticationID, not even an explicit nil
### GetAutoTranslation

`func (o *PostPeople201Response) GetAutoTranslation() bool`

GetAutoTranslation returns the AutoTranslation field if non-nil, zero value otherwise.

### GetAutoTranslationOk

`func (o *PostPeople201Response) GetAutoTranslationOk() (*bool, bool)`

GetAutoTranslationOk returns a tuple with the AutoTranslation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTranslation

`func (o *PostPeople201Response) SetAutoTranslation(v bool)`

SetAutoTranslation sets AutoTranslation field to given value.

### HasAutoTranslation

`func (o *PostPeople201Response) HasAutoTranslation() bool`

HasAutoTranslation returns a boolean if a field has been set.

### GetColorMode

`func (o *PostPeople201Response) GetColorMode() string`

GetColorMode returns the ColorMode field if non-nil, zero value otherwise.

### GetColorModeOk

`func (o *PostPeople201Response) GetColorModeOk() (*string, bool)`

GetColorModeOk returns a tuple with the ColorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorMode

`func (o *PostPeople201Response) SetColorMode(v string)`

SetColorMode sets ColorMode field to given value.

### HasColorMode

`func (o *PostPeople201Response) HasColorMode() bool`

HasColorMode returns a boolean if a field has been set.

### GetContacts

`func (o *PostPeople201Response) GetContacts() []map[string]interface{}`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *PostPeople201Response) GetContactsOk() (*[]map[string]interface{}, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *PostPeople201Response) SetContacts(v []map[string]interface{})`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *PostPeople201Response) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetCostPerHour

`func (o *PostPeople201Response) GetCostPerHour() string`

GetCostPerHour returns the CostPerHour field if non-nil, zero value otherwise.

### GetCostPerHourOk

`func (o *PostPeople201Response) GetCostPerHourOk() (*string, bool)`

GetCostPerHourOk returns a tuple with the CostPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerHour

`func (o *PostPeople201Response) SetCostPerHour(v string)`

SetCostPerHour sets CostPerHour field to given value.

### HasCostPerHour

`func (o *PostPeople201Response) HasCostPerHour() bool`

HasCostPerHour returns a boolean if a field has been set.

### SetCostPerHourNil

`func (o *PostPeople201Response) SetCostPerHourNil(b bool)`

 SetCostPerHourNil sets the value for CostPerHour to be an explicit nil

### UnsetCostPerHour
`func (o *PostPeople201Response) UnsetCostPerHour()`

UnsetCostPerHour ensures that no value is present for CostPerHour, not even an explicit nil
### GetCostPerHourCurrency

`func (o *PostPeople201Response) GetCostPerHourCurrency() string`

GetCostPerHourCurrency returns the CostPerHourCurrency field if non-nil, zero value otherwise.

### GetCostPerHourCurrencyOk

`func (o *PostPeople201Response) GetCostPerHourCurrencyOk() (*string, bool)`

GetCostPerHourCurrencyOk returns a tuple with the CostPerHourCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerHourCurrency

`func (o *PostPeople201Response) SetCostPerHourCurrency(v string)`

SetCostPerHourCurrency sets CostPerHourCurrency field to given value.

### HasCostPerHourCurrency

`func (o *PostPeople201Response) HasCostPerHourCurrency() bool`

HasCostPerHourCurrency returns a boolean if a field has been set.

### SetCostPerHourCurrencyNil

`func (o *PostPeople201Response) SetCostPerHourCurrencyNil(b bool)`

 SetCostPerHourCurrencyNil sets the value for CostPerHourCurrency to be an explicit nil

### UnsetCostPerHourCurrency
`func (o *PostPeople201Response) UnsetCostPerHourCurrency()`

UnsetCostPerHourCurrency ensures that no value is present for CostPerHourCurrency, not even an explicit nil
### GetCreatedAt

`func (o *PostPeople201Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostPeople201Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostPeople201Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PostPeople201Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomFields

`func (o *PostPeople201Response) GetCustomFields() string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PostPeople201Response) GetCustomFieldsOk() (*string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PostPeople201Response) SetCustomFields(v string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PostPeople201Response) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *PostPeople201Response) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *PostPeople201Response) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDisabled

`func (o *PostPeople201Response) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PostPeople201Response) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PostPeople201Response) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PostPeople201Response) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDoNotTranslateLanguages

`func (o *PostPeople201Response) GetDoNotTranslateLanguages() string`

GetDoNotTranslateLanguages returns the DoNotTranslateLanguages field if non-nil, zero value otherwise.

### GetDoNotTranslateLanguagesOk

`func (o *PostPeople201Response) GetDoNotTranslateLanguagesOk() (*string, bool)`

GetDoNotTranslateLanguagesOk returns a tuple with the DoNotTranslateLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTranslateLanguages

`func (o *PostPeople201Response) SetDoNotTranslateLanguages(v string)`

SetDoNotTranslateLanguages sets DoNotTranslateLanguages field to given value.

### HasDoNotTranslateLanguages

`func (o *PostPeople201Response) HasDoNotTranslateLanguages() bool`

HasDoNotTranslateLanguages returns a boolean if a field has been set.

### SetDoNotTranslateLanguagesNil

`func (o *PostPeople201Response) SetDoNotTranslateLanguagesNil(b bool)`

 SetDoNotTranslateLanguagesNil sets the value for DoNotTranslateLanguages to be an explicit nil

### UnsetDoNotTranslateLanguages
`func (o *PostPeople201Response) UnsetDoNotTranslateLanguages()`

UnsetDoNotTranslateLanguages ensures that no value is present for DoNotTranslateLanguages, not even an explicit nil
### GetEmployeeID

`func (o *PostPeople201Response) GetEmployeeID() string`

GetEmployeeID returns the EmployeeID field if non-nil, zero value otherwise.

### GetEmployeeIDOk

`func (o *PostPeople201Response) GetEmployeeIDOk() (*string, bool)`

GetEmployeeIDOk returns a tuple with the EmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeID

`func (o *PostPeople201Response) SetEmployeeID(v string)`

SetEmployeeID sets EmployeeID field to given value.

### HasEmployeeID

`func (o *PostPeople201Response) HasEmployeeID() bool`

HasEmployeeID returns a boolean if a field has been set.

### SetEmployeeIDNil

`func (o *PostPeople201Response) SetEmployeeIDNil(b bool)`

 SetEmployeeIDNil sets the value for EmployeeID to be an explicit nil

### UnsetEmployeeID
`func (o *PostPeople201Response) UnsetEmployeeID()`

UnsetEmployeeID ensures that no value is present for EmployeeID, not even an explicit nil
### GetExcludeTeamNotifications

`func (o *PostPeople201Response) GetExcludeTeamNotifications() bool`

GetExcludeTeamNotifications returns the ExcludeTeamNotifications field if non-nil, zero value otherwise.

### GetExcludeTeamNotificationsOk

`func (o *PostPeople201Response) GetExcludeTeamNotificationsOk() (*bool, bool)`

GetExcludeTeamNotificationsOk returns a tuple with the ExcludeTeamNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeTeamNotifications

`func (o *PostPeople201Response) SetExcludeTeamNotifications(v bool)`

SetExcludeTeamNotifications sets ExcludeTeamNotifications field to given value.

### HasExcludeTeamNotifications

`func (o *PostPeople201Response) HasExcludeTeamNotifications() bool`

HasExcludeTeamNotifications returns a boolean if a field has been set.

### GetFormats

`func (o *PostPeople201Response) GetFormats() PostPeople201ResponseFormats`

GetFormats returns the Formats field if non-nil, zero value otherwise.

### GetFormatsOk

`func (o *PostPeople201Response) GetFormatsOk() (*PostPeople201ResponseFormats, bool)`

GetFormatsOk returns a tuple with the Formats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormats

`func (o *PostPeople201Response) SetFormats(v PostPeople201ResponseFormats)`

SetFormats sets Formats field to given value.

### HasFormats

`func (o *PostPeople201Response) HasFormats() bool`

HasFormats returns a boolean if a field has been set.

### GetGuest

`func (o *PostPeople201Response) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *PostPeople201Response) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *PostPeople201Response) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *PostPeople201Response) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetId

`func (o *PostPeople201Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostPeople201Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostPeople201Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *PostPeople201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInformation

`func (o *PostPeople201Response) GetInformation() string`

GetInformation returns the Information field if non-nil, zero value otherwise.

### GetInformationOk

`func (o *PostPeople201Response) GetInformationOk() (*string, bool)`

GetInformationOk returns a tuple with the Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformation

`func (o *PostPeople201Response) SetInformation(v string)`

SetInformation sets Information field to given value.

### HasInformation

`func (o *PostPeople201Response) HasInformation() bool`

HasInformation returns a boolean if a field has been set.

### SetInformationNil

`func (o *PostPeople201Response) SetInformationNil(b bool)`

 SetInformationNil sets the value for Information to be an explicit nil

### UnsetInformation
`func (o *PostPeople201Response) UnsetInformation()`

UnsetInformation ensures that no value is present for Information, not even an explicit nil
### GetJobTitle

`func (o *PostPeople201Response) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *PostPeople201Response) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *PostPeople201Response) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *PostPeople201Response) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *PostPeople201Response) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *PostPeople201Response) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetLocale

`func (o *PostPeople201Response) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PostPeople201Response) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PostPeople201Response) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *PostPeople201Response) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetLocation

`func (o *PostPeople201Response) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PostPeople201Response) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PostPeople201Response) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PostPeople201Response) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PostPeople201Response) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PostPeople201Response) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetManager

`func (o *PostPeople201Response) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *PostPeople201Response) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *PostPeople201Response) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *PostPeople201Response) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *PostPeople201Response) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *PostPeople201Response) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetName

`func (o *PostPeople201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostPeople201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostPeople201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostPeople201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeID

`func (o *PostPeople201Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *PostPeople201Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *PostPeople201Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *PostPeople201Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetOauthPersonEnablement

`func (o *PostPeople201Response) GetOauthPersonEnablement() bool`

GetOauthPersonEnablement returns the OauthPersonEnablement field if non-nil, zero value otherwise.

### GetOauthPersonEnablementOk

`func (o *PostPeople201Response) GetOauthPersonEnablementOk() (*bool, bool)`

GetOauthPersonEnablementOk returns a tuple with the OauthPersonEnablement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthPersonEnablement

`func (o *PostPeople201Response) SetOauthPersonEnablement(v bool)`

SetOauthPersonEnablement sets OauthPersonEnablement field to given value.

### HasOauthPersonEnablement

`func (o *PostPeople201Response) HasOauthPersonEnablement() bool`

HasOauthPersonEnablement returns a boolean if a field has been set.

### GetOrganization

`func (o *PostPeople201Response) GetOrganization() GetRequestsId200ResponseCreatedBy`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PostPeople201Response) GetOrganizationOk() (*GetRequestsId200ResponseCreatedBy, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PostPeople201Response) SetOrganization(v GetRequestsId200ResponseCreatedBy)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PostPeople201Response) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPictureUri

`func (o *PostPeople201Response) GetPictureUri() string`

GetPictureUri returns the PictureUri field if non-nil, zero value otherwise.

### GetPictureUriOk

`func (o *PostPeople201Response) GetPictureUriOk() (*string, bool)`

GetPictureUriOk returns a tuple with the PictureUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureUri

`func (o *PostPeople201Response) SetPictureUri(v string)`

SetPictureUri sets PictureUri field to given value.

### HasPictureUri

`func (o *PostPeople201Response) HasPictureUri() bool`

HasPictureUri returns a boolean if a field has been set.

### SetPictureUriNil

`func (o *PostPeople201Response) SetPictureUriNil(b bool)`

 SetPictureUriNil sets the value for PictureUri to be an explicit nil

### UnsetPictureUri
`func (o *PostPeople201Response) UnsetPictureUri()`

UnsetPictureUri ensures that no value is present for PictureUri, not even an explicit nil
### GetPlayPrivateChatSound

`func (o *PostPeople201Response) GetPlayPrivateChatSound() string`

GetPlayPrivateChatSound returns the PlayPrivateChatSound field if non-nil, zero value otherwise.

### GetPlayPrivateChatSoundOk

`func (o *PostPeople201Response) GetPlayPrivateChatSoundOk() (*string, bool)`

GetPlayPrivateChatSoundOk returns a tuple with the PlayPrivateChatSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayPrivateChatSound

`func (o *PostPeople201Response) SetPlayPrivateChatSound(v string)`

SetPlayPrivateChatSound sets PlayPrivateChatSound field to given value.

### HasPlayPrivateChatSound

`func (o *PostPeople201Response) HasPlayPrivateChatSound() bool`

HasPlayPrivateChatSound returns a boolean if a field has been set.

### GetPlaySupportChatSound

`func (o *PostPeople201Response) GetPlaySupportChatSound() string`

GetPlaySupportChatSound returns the PlaySupportChatSound field if non-nil, zero value otherwise.

### GetPlaySupportChatSoundOk

`func (o *PostPeople201Response) GetPlaySupportChatSoundOk() (*string, bool)`

GetPlaySupportChatSoundOk returns a tuple with the PlaySupportChatSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySupportChatSound

`func (o *PostPeople201Response) SetPlaySupportChatSound(v string)`

SetPlaySupportChatSound sets PlaySupportChatSound field to given value.

### HasPlaySupportChatSound

`func (o *PostPeople201Response) HasPlaySupportChatSound() bool`

HasPlaySupportChatSound returns a boolean if a field has been set.

### GetPrimaryEmail

`func (o *PostPeople201Response) GetPrimaryEmail() string`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *PostPeople201Response) GetPrimaryEmailOk() (*string, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *PostPeople201Response) SetPrimaryEmail(v string)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *PostPeople201Response) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### GetSendEmailNotifications

`func (o *PostPeople201Response) GetSendEmailNotifications() string`

GetSendEmailNotifications returns the SendEmailNotifications field if non-nil, zero value otherwise.

### GetSendEmailNotificationsOk

`func (o *PostPeople201Response) GetSendEmailNotificationsOk() (*string, bool)`

GetSendEmailNotificationsOk returns a tuple with the SendEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmailNotifications

`func (o *PostPeople201Response) SetSendEmailNotifications(v string)`

SetSendEmailNotifications sets SendEmailNotifications field to given value.

### HasSendEmailNotifications

`func (o *PostPeople201Response) HasSendEmailNotifications() bool`

HasSendEmailNotifications returns a boolean if a field has been set.

### GetShowNotificationPopup

`func (o *PostPeople201Response) GetShowNotificationPopup() string`

GetShowNotificationPopup returns the ShowNotificationPopup field if non-nil, zero value otherwise.

### GetShowNotificationPopupOk

`func (o *PostPeople201Response) GetShowNotificationPopupOk() (*string, bool)`

GetShowNotificationPopupOk returns a tuple with the ShowNotificationPopup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNotificationPopup

`func (o *PostPeople201Response) SetShowNotificationPopup(v string)`

SetShowNotificationPopup sets ShowNotificationPopup field to given value.

### HasShowNotificationPopup

`func (o *PostPeople201Response) HasShowNotificationPopup() bool`

HasShowNotificationPopup returns a boolean if a field has been set.

### GetShowPrivateChatPopup

`func (o *PostPeople201Response) GetShowPrivateChatPopup() string`

GetShowPrivateChatPopup returns the ShowPrivateChatPopup field if non-nil, zero value otherwise.

### GetShowPrivateChatPopupOk

`func (o *PostPeople201Response) GetShowPrivateChatPopupOk() (*string, bool)`

GetShowPrivateChatPopupOk returns a tuple with the ShowPrivateChatPopup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPrivateChatPopup

`func (o *PostPeople201Response) SetShowPrivateChatPopup(v string)`

SetShowPrivateChatPopup sets ShowPrivateChatPopup field to given value.

### HasShowPrivateChatPopup

`func (o *PostPeople201Response) HasShowPrivateChatPopup() bool`

HasShowPrivateChatPopup returns a boolean if a field has been set.

### GetShowSupportChatPopup

`func (o *PostPeople201Response) GetShowSupportChatPopup() string`

GetShowSupportChatPopup returns the ShowSupportChatPopup field if non-nil, zero value otherwise.

### GetShowSupportChatPopupOk

`func (o *PostPeople201Response) GetShowSupportChatPopupOk() (*string, bool)`

GetShowSupportChatPopupOk returns a tuple with the ShowSupportChatPopup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSupportChatPopup

`func (o *PostPeople201Response) SetShowSupportChatPopup(v string)`

SetShowSupportChatPopup sets ShowSupportChatPopup field to given value.

### HasShowSupportChatPopup

`func (o *PostPeople201Response) HasShowSupportChatPopup() bool`

HasShowSupportChatPopup returns a boolean if a field has been set.

### GetSite

`func (o *PostPeople201Response) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PostPeople201Response) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PostPeople201Response) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *PostPeople201Response) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *PostPeople201Response) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *PostPeople201Response) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetSource

`func (o *PostPeople201Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PostPeople201Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PostPeople201Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PostPeople201Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *PostPeople201Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *PostPeople201Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *PostPeople201Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *PostPeople201Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *PostPeople201Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *PostPeople201Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetSupportID

`func (o *PostPeople201Response) GetSupportID() string`

GetSupportID returns the SupportID field if non-nil, zero value otherwise.

### GetSupportIDOk

`func (o *PostPeople201Response) GetSupportIDOk() (*string, bool)`

GetSupportIDOk returns a tuple with the SupportID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportID

`func (o *PostPeople201Response) SetSupportID(v string)`

SetSupportID sets SupportID field to given value.

### HasSupportID

`func (o *PostPeople201Response) HasSupportID() bool`

HasSupportID returns a boolean if a field has been set.

### SetSupportIDNil

`func (o *PostPeople201Response) SetSupportIDNil(b bool)`

 SetSupportIDNil sets the value for SupportID to be an explicit nil

### UnsetSupportID
`func (o *PostPeople201Response) UnsetSupportID()`

UnsetSupportID ensures that no value is present for SupportID, not even an explicit nil
### GetTimeFormat24h

`func (o *PostPeople201Response) GetTimeFormat24h() bool`

GetTimeFormat24h returns the TimeFormat24h field if non-nil, zero value otherwise.

### GetTimeFormat24hOk

`func (o *PostPeople201Response) GetTimeFormat24hOk() (*bool, bool)`

GetTimeFormat24hOk returns a tuple with the TimeFormat24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFormat24h

`func (o *PostPeople201Response) SetTimeFormat24h(v bool)`

SetTimeFormat24h sets TimeFormat24h field to given value.

### HasTimeFormat24h

`func (o *PostPeople201Response) HasTimeFormat24h() bool`

HasTimeFormat24h returns a boolean if a field has been set.

### GetTimeZone

`func (o *PostPeople201Response) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *PostPeople201Response) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *PostPeople201Response) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *PostPeople201Response) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUiExtension

`func (o *PostPeople201Response) GetUiExtension() string`

GetUiExtension returns the UiExtension field if non-nil, zero value otherwise.

### GetUiExtensionOk

`func (o *PostPeople201Response) GetUiExtensionOk() (*string, bool)`

GetUiExtensionOk returns a tuple with the UiExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiExtension

`func (o *PostPeople201Response) SetUiExtension(v string)`

SetUiExtension sets UiExtension field to given value.

### HasUiExtension

`func (o *PostPeople201Response) HasUiExtension() bool`

HasUiExtension returns a boolean if a field has been set.

### SetUiExtensionNil

`func (o *PostPeople201Response) SetUiExtensionNil(b bool)`

 SetUiExtensionNil sets the value for UiExtension to be an explicit nil

### UnsetUiExtension
`func (o *PostPeople201Response) UnsetUiExtension()`

UnsetUiExtension ensures that no value is present for UiExtension, not even an explicit nil
### GetUpdatedAt

`func (o *PostPeople201Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PostPeople201Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PostPeople201Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PostPeople201Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVip

`func (o *PostPeople201Response) GetVip() bool`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *PostPeople201Response) GetVipOk() (*bool, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *PostPeople201Response) SetVip(v bool)`

SetVip sets Vip field to given value.

### HasVip

`func (o *PostPeople201Response) HasVip() bool`

HasVip returns a boolean if a field has been set.

### GetWorkHours

`func (o *PostPeople201Response) GetWorkHours() string`

GetWorkHours returns the WorkHours field if non-nil, zero value otherwise.

### GetWorkHoursOk

`func (o *PostPeople201Response) GetWorkHoursOk() (*string, bool)`

GetWorkHoursOk returns a tuple with the WorkHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkHours

`func (o *PostPeople201Response) SetWorkHours(v string)`

SetWorkHours sets WorkHours field to given value.

### HasWorkHours

`func (o *PostPeople201Response) HasWorkHours() bool`

HasWorkHours returns a boolean if a field has been set.

### SetWorkHoursNil

`func (o *PostPeople201Response) SetWorkHoursNil(b bool)`

 SetWorkHoursNil sets the value for WorkHours to be an explicit nil

### UnsetWorkHours
`func (o *PostPeople201Response) UnsetWorkHours()`

UnsetWorkHours ensures that no value is present for WorkHours, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


