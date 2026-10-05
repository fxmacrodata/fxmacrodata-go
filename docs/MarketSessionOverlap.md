# MarketSessionOverlap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Sessions** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Priority** | Pointer to **NullableString** |  | [optional] 
**NotablePairs** | Pointer to **[]string** |  | [optional] 
**StartUtc** | Pointer to **NullableString** |  | [optional] 
**EndUtc** | Pointer to **NullableString** |  | [optional] 
**StartUnix** | Pointer to **NullableInt64** |  | [optional] 
**EndUnix** | Pointer to **NullableInt64** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**SecondsToStart** | Pointer to **NullableInt64** |  | [optional] 
**SecondsToEnd** | Pointer to **NullableInt64** |  | [optional] 
**DurationHours** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewMarketSessionOverlap

`func NewMarketSessionOverlap() *MarketSessionOverlap`

NewMarketSessionOverlap instantiates a new MarketSessionOverlap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketSessionOverlapWithDefaults

`func NewMarketSessionOverlapWithDefaults() *MarketSessionOverlap`

NewMarketSessionOverlapWithDefaults instantiates a new MarketSessionOverlap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MarketSessionOverlap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketSessionOverlap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketSessionOverlap) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketSessionOverlap) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MarketSessionOverlap) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MarketSessionOverlap) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSessions

`func (o *MarketSessionOverlap) GetSessions() []string`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *MarketSessionOverlap) GetSessionsOk() (*[]string, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *MarketSessionOverlap) SetSessions(v []string)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *MarketSessionOverlap) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### SetSessionsNil

`func (o *MarketSessionOverlap) SetSessionsNil(b bool)`

 SetSessionsNil sets the value for Sessions to be an explicit nil

### UnsetSessions
`func (o *MarketSessionOverlap) UnsetSessions()`

UnsetSessions ensures that no value is present for Sessions, not even an explicit nil
### GetDescription

`func (o *MarketSessionOverlap) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketSessionOverlap) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketSessionOverlap) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketSessionOverlap) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MarketSessionOverlap) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MarketSessionOverlap) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriority

`func (o *MarketSessionOverlap) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MarketSessionOverlap) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MarketSessionOverlap) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MarketSessionOverlap) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *MarketSessionOverlap) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *MarketSessionOverlap) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetNotablePairs

`func (o *MarketSessionOverlap) GetNotablePairs() []string`

GetNotablePairs returns the NotablePairs field if non-nil, zero value otherwise.

### GetNotablePairsOk

`func (o *MarketSessionOverlap) GetNotablePairsOk() (*[]string, bool)`

GetNotablePairsOk returns a tuple with the NotablePairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotablePairs

`func (o *MarketSessionOverlap) SetNotablePairs(v []string)`

SetNotablePairs sets NotablePairs field to given value.

### HasNotablePairs

`func (o *MarketSessionOverlap) HasNotablePairs() bool`

HasNotablePairs returns a boolean if a field has been set.

### SetNotablePairsNil

`func (o *MarketSessionOverlap) SetNotablePairsNil(b bool)`

 SetNotablePairsNil sets the value for NotablePairs to be an explicit nil

### UnsetNotablePairs
`func (o *MarketSessionOverlap) UnsetNotablePairs()`

UnsetNotablePairs ensures that no value is present for NotablePairs, not even an explicit nil
### GetStartUtc

`func (o *MarketSessionOverlap) GetStartUtc() string`

GetStartUtc returns the StartUtc field if non-nil, zero value otherwise.

### GetStartUtcOk

`func (o *MarketSessionOverlap) GetStartUtcOk() (*string, bool)`

GetStartUtcOk returns a tuple with the StartUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartUtc

`func (o *MarketSessionOverlap) SetStartUtc(v string)`

SetStartUtc sets StartUtc field to given value.

### HasStartUtc

`func (o *MarketSessionOverlap) HasStartUtc() bool`

HasStartUtc returns a boolean if a field has been set.

### SetStartUtcNil

`func (o *MarketSessionOverlap) SetStartUtcNil(b bool)`

 SetStartUtcNil sets the value for StartUtc to be an explicit nil

### UnsetStartUtc
`func (o *MarketSessionOverlap) UnsetStartUtc()`

UnsetStartUtc ensures that no value is present for StartUtc, not even an explicit nil
### GetEndUtc

`func (o *MarketSessionOverlap) GetEndUtc() string`

GetEndUtc returns the EndUtc field if non-nil, zero value otherwise.

### GetEndUtcOk

`func (o *MarketSessionOverlap) GetEndUtcOk() (*string, bool)`

GetEndUtcOk returns a tuple with the EndUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUtc

`func (o *MarketSessionOverlap) SetEndUtc(v string)`

SetEndUtc sets EndUtc field to given value.

### HasEndUtc

`func (o *MarketSessionOverlap) HasEndUtc() bool`

HasEndUtc returns a boolean if a field has been set.

### SetEndUtcNil

`func (o *MarketSessionOverlap) SetEndUtcNil(b bool)`

 SetEndUtcNil sets the value for EndUtc to be an explicit nil

### UnsetEndUtc
`func (o *MarketSessionOverlap) UnsetEndUtc()`

UnsetEndUtc ensures that no value is present for EndUtc, not even an explicit nil
### GetStartUnix

`func (o *MarketSessionOverlap) GetStartUnix() int64`

GetStartUnix returns the StartUnix field if non-nil, zero value otherwise.

### GetStartUnixOk

`func (o *MarketSessionOverlap) GetStartUnixOk() (*int64, bool)`

GetStartUnixOk returns a tuple with the StartUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartUnix

`func (o *MarketSessionOverlap) SetStartUnix(v int64)`

SetStartUnix sets StartUnix field to given value.

### HasStartUnix

`func (o *MarketSessionOverlap) HasStartUnix() bool`

HasStartUnix returns a boolean if a field has been set.

### SetStartUnixNil

`func (o *MarketSessionOverlap) SetStartUnixNil(b bool)`

 SetStartUnixNil sets the value for StartUnix to be an explicit nil

### UnsetStartUnix
`func (o *MarketSessionOverlap) UnsetStartUnix()`

UnsetStartUnix ensures that no value is present for StartUnix, not even an explicit nil
### GetEndUnix

`func (o *MarketSessionOverlap) GetEndUnix() int64`

GetEndUnix returns the EndUnix field if non-nil, zero value otherwise.

### GetEndUnixOk

`func (o *MarketSessionOverlap) GetEndUnixOk() (*int64, bool)`

GetEndUnixOk returns a tuple with the EndUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUnix

`func (o *MarketSessionOverlap) SetEndUnix(v int64)`

SetEndUnix sets EndUnix field to given value.

### HasEndUnix

`func (o *MarketSessionOverlap) HasEndUnix() bool`

HasEndUnix returns a boolean if a field has been set.

### SetEndUnixNil

`func (o *MarketSessionOverlap) SetEndUnixNil(b bool)`

 SetEndUnixNil sets the value for EndUnix to be an explicit nil

### UnsetEndUnix
`func (o *MarketSessionOverlap) UnsetEndUnix()`

UnsetEndUnix ensures that no value is present for EndUnix, not even an explicit nil
### GetIsActive

`func (o *MarketSessionOverlap) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *MarketSessionOverlap) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *MarketSessionOverlap) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *MarketSessionOverlap) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *MarketSessionOverlap) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *MarketSessionOverlap) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetSecondsToStart

`func (o *MarketSessionOverlap) GetSecondsToStart() int64`

GetSecondsToStart returns the SecondsToStart field if non-nil, zero value otherwise.

### GetSecondsToStartOk

`func (o *MarketSessionOverlap) GetSecondsToStartOk() (*int64, bool)`

GetSecondsToStartOk returns a tuple with the SecondsToStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsToStart

`func (o *MarketSessionOverlap) SetSecondsToStart(v int64)`

SetSecondsToStart sets SecondsToStart field to given value.

### HasSecondsToStart

`func (o *MarketSessionOverlap) HasSecondsToStart() bool`

HasSecondsToStart returns a boolean if a field has been set.

### SetSecondsToStartNil

`func (o *MarketSessionOverlap) SetSecondsToStartNil(b bool)`

 SetSecondsToStartNil sets the value for SecondsToStart to be an explicit nil

### UnsetSecondsToStart
`func (o *MarketSessionOverlap) UnsetSecondsToStart()`

UnsetSecondsToStart ensures that no value is present for SecondsToStart, not even an explicit nil
### GetSecondsToEnd

`func (o *MarketSessionOverlap) GetSecondsToEnd() int64`

GetSecondsToEnd returns the SecondsToEnd field if non-nil, zero value otherwise.

### GetSecondsToEndOk

`func (o *MarketSessionOverlap) GetSecondsToEndOk() (*int64, bool)`

GetSecondsToEndOk returns a tuple with the SecondsToEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsToEnd

`func (o *MarketSessionOverlap) SetSecondsToEnd(v int64)`

SetSecondsToEnd sets SecondsToEnd field to given value.

### HasSecondsToEnd

`func (o *MarketSessionOverlap) HasSecondsToEnd() bool`

HasSecondsToEnd returns a boolean if a field has been set.

### SetSecondsToEndNil

`func (o *MarketSessionOverlap) SetSecondsToEndNil(b bool)`

 SetSecondsToEndNil sets the value for SecondsToEnd to be an explicit nil

### UnsetSecondsToEnd
`func (o *MarketSessionOverlap) UnsetSecondsToEnd()`

UnsetSecondsToEnd ensures that no value is present for SecondsToEnd, not even an explicit nil
### GetDurationHours

`func (o *MarketSessionOverlap) GetDurationHours() float64`

GetDurationHours returns the DurationHours field if non-nil, zero value otherwise.

### GetDurationHoursOk

`func (o *MarketSessionOverlap) GetDurationHoursOk() (*float64, bool)`

GetDurationHoursOk returns a tuple with the DurationHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationHours

`func (o *MarketSessionOverlap) SetDurationHours(v float64)`

SetDurationHours sets DurationHours field to given value.

### HasDurationHours

`func (o *MarketSessionOverlap) HasDurationHours() bool`

HasDurationHours returns a boolean if a field has been set.

### SetDurationHoursNil

`func (o *MarketSessionOverlap) SetDurationHoursNil(b bool)`

 SetDurationHoursNil sets the value for DurationHours to be an explicit nil

### UnsetDurationHours
`func (o *MarketSessionOverlap) UnsetDurationHours()`

UnsetDurationHours ensures that no value is present for DurationHours, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


