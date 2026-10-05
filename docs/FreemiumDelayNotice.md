# FreemiumDelayNotice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applied** | **bool** |  | 
**DelaySeconds** | **int64** |  | 
**DelayMinutes** | **int64** |  | 
**Cutoff** | **int64** |  | 
**CutoffIso** | Pointer to **NullableString** |  | [optional] 
**WithheldCount** | **int64** |  | 
**Message** | **string** |  | 
**SubscribeUrl** | **string** |  | 
**NextAvailableAt** | Pointer to **NullableInt64** |  | [optional] 
**NextAvailableAtIso** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFreemiumDelayNotice

`func NewFreemiumDelayNotice(applied bool, delaySeconds int64, delayMinutes int64, cutoff int64, withheldCount int64, message string, subscribeUrl string, ) *FreemiumDelayNotice`

NewFreemiumDelayNotice instantiates a new FreemiumDelayNotice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreemiumDelayNoticeWithDefaults

`func NewFreemiumDelayNoticeWithDefaults() *FreemiumDelayNotice`

NewFreemiumDelayNoticeWithDefaults instantiates a new FreemiumDelayNotice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplied

`func (o *FreemiumDelayNotice) GetApplied() bool`

GetApplied returns the Applied field if non-nil, zero value otherwise.

### GetAppliedOk

`func (o *FreemiumDelayNotice) GetAppliedOk() (*bool, bool)`

GetAppliedOk returns a tuple with the Applied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplied

`func (o *FreemiumDelayNotice) SetApplied(v bool)`

SetApplied sets Applied field to given value.


### GetDelaySeconds

`func (o *FreemiumDelayNotice) GetDelaySeconds() int64`

GetDelaySeconds returns the DelaySeconds field if non-nil, zero value otherwise.

### GetDelaySecondsOk

`func (o *FreemiumDelayNotice) GetDelaySecondsOk() (*int64, bool)`

GetDelaySecondsOk returns a tuple with the DelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelaySeconds

`func (o *FreemiumDelayNotice) SetDelaySeconds(v int64)`

SetDelaySeconds sets DelaySeconds field to given value.


### GetDelayMinutes

`func (o *FreemiumDelayNotice) GetDelayMinutes() int64`

GetDelayMinutes returns the DelayMinutes field if non-nil, zero value otherwise.

### GetDelayMinutesOk

`func (o *FreemiumDelayNotice) GetDelayMinutesOk() (*int64, bool)`

GetDelayMinutesOk returns a tuple with the DelayMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMinutes

`func (o *FreemiumDelayNotice) SetDelayMinutes(v int64)`

SetDelayMinutes sets DelayMinutes field to given value.


### GetCutoff

`func (o *FreemiumDelayNotice) GetCutoff() int64`

GetCutoff returns the Cutoff field if non-nil, zero value otherwise.

### GetCutoffOk

`func (o *FreemiumDelayNotice) GetCutoffOk() (*int64, bool)`

GetCutoffOk returns a tuple with the Cutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoff

`func (o *FreemiumDelayNotice) SetCutoff(v int64)`

SetCutoff sets Cutoff field to given value.


### GetCutoffIso

`func (o *FreemiumDelayNotice) GetCutoffIso() string`

GetCutoffIso returns the CutoffIso field if non-nil, zero value otherwise.

### GetCutoffIsoOk

`func (o *FreemiumDelayNotice) GetCutoffIsoOk() (*string, bool)`

GetCutoffIsoOk returns a tuple with the CutoffIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoffIso

`func (o *FreemiumDelayNotice) SetCutoffIso(v string)`

SetCutoffIso sets CutoffIso field to given value.

### HasCutoffIso

`func (o *FreemiumDelayNotice) HasCutoffIso() bool`

HasCutoffIso returns a boolean if a field has been set.

### SetCutoffIsoNil

`func (o *FreemiumDelayNotice) SetCutoffIsoNil(b bool)`

 SetCutoffIsoNil sets the value for CutoffIso to be an explicit nil

### UnsetCutoffIso
`func (o *FreemiumDelayNotice) UnsetCutoffIso()`

UnsetCutoffIso ensures that no value is present for CutoffIso, not even an explicit nil
### GetWithheldCount

`func (o *FreemiumDelayNotice) GetWithheldCount() int64`

GetWithheldCount returns the WithheldCount field if non-nil, zero value otherwise.

### GetWithheldCountOk

`func (o *FreemiumDelayNotice) GetWithheldCountOk() (*int64, bool)`

GetWithheldCountOk returns a tuple with the WithheldCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheldCount

`func (o *FreemiumDelayNotice) SetWithheldCount(v int64)`

SetWithheldCount sets WithheldCount field to given value.


### GetMessage

`func (o *FreemiumDelayNotice) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FreemiumDelayNotice) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FreemiumDelayNotice) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSubscribeUrl

`func (o *FreemiumDelayNotice) GetSubscribeUrl() string`

GetSubscribeUrl returns the SubscribeUrl field if non-nil, zero value otherwise.

### GetSubscribeUrlOk

`func (o *FreemiumDelayNotice) GetSubscribeUrlOk() (*string, bool)`

GetSubscribeUrlOk returns a tuple with the SubscribeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeUrl

`func (o *FreemiumDelayNotice) SetSubscribeUrl(v string)`

SetSubscribeUrl sets SubscribeUrl field to given value.


### GetNextAvailableAt

`func (o *FreemiumDelayNotice) GetNextAvailableAt() int64`

GetNextAvailableAt returns the NextAvailableAt field if non-nil, zero value otherwise.

### GetNextAvailableAtOk

`func (o *FreemiumDelayNotice) GetNextAvailableAtOk() (*int64, bool)`

GetNextAvailableAtOk returns a tuple with the NextAvailableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAvailableAt

`func (o *FreemiumDelayNotice) SetNextAvailableAt(v int64)`

SetNextAvailableAt sets NextAvailableAt field to given value.

### HasNextAvailableAt

`func (o *FreemiumDelayNotice) HasNextAvailableAt() bool`

HasNextAvailableAt returns a boolean if a field has been set.

### SetNextAvailableAtNil

`func (o *FreemiumDelayNotice) SetNextAvailableAtNil(b bool)`

 SetNextAvailableAtNil sets the value for NextAvailableAt to be an explicit nil

### UnsetNextAvailableAt
`func (o *FreemiumDelayNotice) UnsetNextAvailableAt()`

UnsetNextAvailableAt ensures that no value is present for NextAvailableAt, not even an explicit nil
### GetNextAvailableAtIso

`func (o *FreemiumDelayNotice) GetNextAvailableAtIso() string`

GetNextAvailableAtIso returns the NextAvailableAtIso field if non-nil, zero value otherwise.

### GetNextAvailableAtIsoOk

`func (o *FreemiumDelayNotice) GetNextAvailableAtIsoOk() (*string, bool)`

GetNextAvailableAtIsoOk returns a tuple with the NextAvailableAtIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAvailableAtIso

`func (o *FreemiumDelayNotice) SetNextAvailableAtIso(v string)`

SetNextAvailableAtIso sets NextAvailableAtIso field to given value.

### HasNextAvailableAtIso

`func (o *FreemiumDelayNotice) HasNextAvailableAtIso() bool`

HasNextAvailableAtIso returns a boolean if a field has been set.

### SetNextAvailableAtIsoNil

`func (o *FreemiumDelayNotice) SetNextAvailableAtIsoNil(b bool)`

 SetNextAvailableAtIsoNil sets the value for NextAvailableAtIso to be an explicit nil

### UnsetNextAvailableAtIso
`func (o *FreemiumDelayNotice) UnsetNextAvailableAtIso()`

UnsetNextAvailableAtIso ensures that no value is present for NextAvailableAtIso, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


