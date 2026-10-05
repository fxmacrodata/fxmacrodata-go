# FreemiumWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applied** | **bool** |  | 
**MaxDays** | **int64** |  | 
**CutoffDate** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFreemiumWindow

`func NewFreemiumWindow(applied bool, maxDays int64, ) *FreemiumWindow`

NewFreemiumWindow instantiates a new FreemiumWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreemiumWindowWithDefaults

`func NewFreemiumWindowWithDefaults() *FreemiumWindow`

NewFreemiumWindowWithDefaults instantiates a new FreemiumWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplied

`func (o *FreemiumWindow) GetApplied() bool`

GetApplied returns the Applied field if non-nil, zero value otherwise.

### GetAppliedOk

`func (o *FreemiumWindow) GetAppliedOk() (*bool, bool)`

GetAppliedOk returns a tuple with the Applied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplied

`func (o *FreemiumWindow) SetApplied(v bool)`

SetApplied sets Applied field to given value.


### GetMaxDays

`func (o *FreemiumWindow) GetMaxDays() int64`

GetMaxDays returns the MaxDays field if non-nil, zero value otherwise.

### GetMaxDaysOk

`func (o *FreemiumWindow) GetMaxDaysOk() (*int64, bool)`

GetMaxDaysOk returns a tuple with the MaxDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDays

`func (o *FreemiumWindow) SetMaxDays(v int64)`

SetMaxDays sets MaxDays field to given value.


### GetCutoffDate

`func (o *FreemiumWindow) GetCutoffDate() string`

GetCutoffDate returns the CutoffDate field if non-nil, zero value otherwise.

### GetCutoffDateOk

`func (o *FreemiumWindow) GetCutoffDateOk() (*string, bool)`

GetCutoffDateOk returns a tuple with the CutoffDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoffDate

`func (o *FreemiumWindow) SetCutoffDate(v string)`

SetCutoffDate sets CutoffDate field to given value.

### HasCutoffDate

`func (o *FreemiumWindow) HasCutoffDate() bool`

HasCutoffDate returns a boolean if a field has been set.

### SetCutoffDateNil

`func (o *FreemiumWindow) SetCutoffDateNil(b bool)`

 SetCutoffDateNil sets the value for CutoffDate to be an explicit nil

### UnsetCutoffDate
`func (o *FreemiumWindow) UnsetCutoffDate()`

UnsetCutoffDate ensures that no value is present for CutoffDate, not even an explicit nil
### GetMessage

`func (o *FreemiumWindow) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FreemiumWindow) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FreemiumWindow) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FreemiumWindow) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *FreemiumWindow) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *FreemiumWindow) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


