# MarketSessionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NowUtc** | Pointer to **NullableString** |  | [optional] 
**NowUnix** | Pointer to **NullableInt64** |  | [optional] 
**IsMarketDay** | Pointer to **NullableBool** |  | [optional] 
**Sessions** | Pointer to [**[]MarketSession**](MarketSession.md) |  | [optional] 
**Overlaps** | Pointer to [**[]MarketSessionOverlap**](MarketSessionOverlap.md) |  | [optional] 

## Methods

### NewMarketSessionsResponse

`func NewMarketSessionsResponse() *MarketSessionsResponse`

NewMarketSessionsResponse instantiates a new MarketSessionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketSessionsResponseWithDefaults

`func NewMarketSessionsResponseWithDefaults() *MarketSessionsResponse`

NewMarketSessionsResponseWithDefaults instantiates a new MarketSessionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNowUtc

`func (o *MarketSessionsResponse) GetNowUtc() string`

GetNowUtc returns the NowUtc field if non-nil, zero value otherwise.

### GetNowUtcOk

`func (o *MarketSessionsResponse) GetNowUtcOk() (*string, bool)`

GetNowUtcOk returns a tuple with the NowUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowUtc

`func (o *MarketSessionsResponse) SetNowUtc(v string)`

SetNowUtc sets NowUtc field to given value.

### HasNowUtc

`func (o *MarketSessionsResponse) HasNowUtc() bool`

HasNowUtc returns a boolean if a field has been set.

### SetNowUtcNil

`func (o *MarketSessionsResponse) SetNowUtcNil(b bool)`

 SetNowUtcNil sets the value for NowUtc to be an explicit nil

### UnsetNowUtc
`func (o *MarketSessionsResponse) UnsetNowUtc()`

UnsetNowUtc ensures that no value is present for NowUtc, not even an explicit nil
### GetNowUnix

`func (o *MarketSessionsResponse) GetNowUnix() int64`

GetNowUnix returns the NowUnix field if non-nil, zero value otherwise.

### GetNowUnixOk

`func (o *MarketSessionsResponse) GetNowUnixOk() (*int64, bool)`

GetNowUnixOk returns a tuple with the NowUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowUnix

`func (o *MarketSessionsResponse) SetNowUnix(v int64)`

SetNowUnix sets NowUnix field to given value.

### HasNowUnix

`func (o *MarketSessionsResponse) HasNowUnix() bool`

HasNowUnix returns a boolean if a field has been set.

### SetNowUnixNil

`func (o *MarketSessionsResponse) SetNowUnixNil(b bool)`

 SetNowUnixNil sets the value for NowUnix to be an explicit nil

### UnsetNowUnix
`func (o *MarketSessionsResponse) UnsetNowUnix()`

UnsetNowUnix ensures that no value is present for NowUnix, not even an explicit nil
### GetIsMarketDay

`func (o *MarketSessionsResponse) GetIsMarketDay() bool`

GetIsMarketDay returns the IsMarketDay field if non-nil, zero value otherwise.

### GetIsMarketDayOk

`func (o *MarketSessionsResponse) GetIsMarketDayOk() (*bool, bool)`

GetIsMarketDayOk returns a tuple with the IsMarketDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarketDay

`func (o *MarketSessionsResponse) SetIsMarketDay(v bool)`

SetIsMarketDay sets IsMarketDay field to given value.

### HasIsMarketDay

`func (o *MarketSessionsResponse) HasIsMarketDay() bool`

HasIsMarketDay returns a boolean if a field has been set.

### SetIsMarketDayNil

`func (o *MarketSessionsResponse) SetIsMarketDayNil(b bool)`

 SetIsMarketDayNil sets the value for IsMarketDay to be an explicit nil

### UnsetIsMarketDay
`func (o *MarketSessionsResponse) UnsetIsMarketDay()`

UnsetIsMarketDay ensures that no value is present for IsMarketDay, not even an explicit nil
### GetSessions

`func (o *MarketSessionsResponse) GetSessions() []MarketSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *MarketSessionsResponse) GetSessionsOk() (*[]MarketSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *MarketSessionsResponse) SetSessions(v []MarketSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *MarketSessionsResponse) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### SetSessionsNil

`func (o *MarketSessionsResponse) SetSessionsNil(b bool)`

 SetSessionsNil sets the value for Sessions to be an explicit nil

### UnsetSessions
`func (o *MarketSessionsResponse) UnsetSessions()`

UnsetSessions ensures that no value is present for Sessions, not even an explicit nil
### GetOverlaps

`func (o *MarketSessionsResponse) GetOverlaps() []MarketSessionOverlap`

GetOverlaps returns the Overlaps field if non-nil, zero value otherwise.

### GetOverlapsOk

`func (o *MarketSessionsResponse) GetOverlapsOk() (*[]MarketSessionOverlap, bool)`

GetOverlapsOk returns a tuple with the Overlaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlaps

`func (o *MarketSessionsResponse) SetOverlaps(v []MarketSessionOverlap)`

SetOverlaps sets Overlaps field to given value.

### HasOverlaps

`func (o *MarketSessionsResponse) HasOverlaps() bool`

HasOverlaps returns a boolean if a field has been set.

### SetOverlapsNil

`func (o *MarketSessionsResponse) SetOverlapsNil(b bool)`

 SetOverlapsNil sets the value for Overlaps to be an explicit nil

### UnsetOverlaps
`func (o *MarketSessionsResponse) UnsetOverlaps()`

UnsetOverlaps ensures that no value is present for Overlaps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


