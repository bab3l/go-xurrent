# PostCalendars201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalendarHours** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceID** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewPostCalendars201Response

`func NewPostCalendars201Response() *PostCalendars201Response`

NewPostCalendars201Response instantiates a new PostCalendars201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCalendars201ResponseWithDefaults

`func NewPostCalendars201ResponseWithDefaults() *PostCalendars201Response`

NewPostCalendars201ResponseWithDefaults instantiates a new PostCalendars201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalendarHours

`func (o *PostCalendars201Response) GetCalendarHours() []map[string]interface{}`

GetCalendarHours returns the CalendarHours field if non-nil, zero value otherwise.

### GetCalendarHoursOk

`func (o *PostCalendars201Response) GetCalendarHoursOk() (*[]map[string]interface{}, bool)`

GetCalendarHoursOk returns a tuple with the CalendarHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarHours

`func (o *PostCalendars201Response) SetCalendarHours(v []map[string]interface{})`

SetCalendarHours sets CalendarHours field to given value.

### HasCalendarHours

`func (o *PostCalendars201Response) HasCalendarHours() bool`

HasCalendarHours returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PostCalendars201Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostCalendars201Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostCalendars201Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PostCalendars201Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisabled

`func (o *PostCalendars201Response) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PostCalendars201Response) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PostCalendars201Response) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PostCalendars201Response) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetId

`func (o *PostCalendars201Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCalendars201Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCalendars201Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *PostCalendars201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PostCalendars201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostCalendars201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostCalendars201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostCalendars201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeID

`func (o *PostCalendars201Response) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *PostCalendars201Response) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *PostCalendars201Response) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *PostCalendars201Response) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetSource

`func (o *PostCalendars201Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PostCalendars201Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PostCalendars201Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PostCalendars201Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *PostCalendars201Response) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *PostCalendars201Response) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *PostCalendars201Response) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *PostCalendars201Response) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### SetSourceIDNil

`func (o *PostCalendars201Response) SetSourceIDNil(b bool)`

 SetSourceIDNil sets the value for SourceID to be an explicit nil

### UnsetSourceID
`func (o *PostCalendars201Response) UnsetSourceID()`

UnsetSourceID ensures that no value is present for SourceID, not even an explicit nil
### GetUpdatedAt

`func (o *PostCalendars201Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PostCalendars201Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PostCalendars201Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PostCalendars201Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


