# RateDifferentialPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**BaseVal** | Pointer to **NullableFloat64** |  | [optional] 
**QuoteVal** | Pointer to **NullableFloat64** |  | [optional] 
**Spread** | Pointer to **NullableFloat64** |  | [optional] 
**SpreadBps** | Pointer to **NullableFloat64** |  | [optional] 
**Differential** | Pointer to **NullableFloat64** |  | [optional] 
**DifferentialBps** | Pointer to **NullableFloat64** |  | [optional] 
**BaseForwardVal** | Pointer to **NullableFloat64** |  | [optional] 
**QuoteForwardVal** | Pointer to **NullableFloat64** |  | [optional] 
**BaseStartVal** | Pointer to **NullableFloat64** |  | [optional] 
**BaseEndVal** | Pointer to **NullableFloat64** |  | [optional] 
**QuoteStartVal** | Pointer to **NullableFloat64** |  | [optional] 
**QuoteEndVal** | Pointer to **NullableFloat64** |  | [optional] 
**BaseAnnouncementDatetime** | Pointer to **NullableInt64** |  | [optional] 
**QuoteAnnouncementDatetime** | Pointer to **NullableInt64** |  | [optional] 
**BaseStartAnnouncementDatetime** | Pointer to **NullableInt64** |  | [optional] 
**BaseEndAnnouncementDatetime** | Pointer to **NullableInt64** |  | [optional] 
**QuoteStartAnnouncementDatetime** | Pointer to **NullableInt64** |  | [optional] 
**QuoteEndAnnouncementDatetime** | Pointer to **NullableInt64** |  | [optional] 
**BaseSelectedVintageKnownAtNs** | Pointer to **NullableString** |  | [optional] 
**QuoteSelectedVintageKnownAtNs** | Pointer to **NullableString** |  | [optional] 
**BaseStartSelectedVintageKnownAtNs** | Pointer to **NullableString** |  | [optional] 
**BaseEndSelectedVintageKnownAtNs** | Pointer to **NullableString** |  | [optional] 
**QuoteStartSelectedVintageKnownAtNs** | Pointer to **NullableString** |  | [optional] 
**QuoteEndSelectedVintageKnownAtNs** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRateDifferentialPoint

`func NewRateDifferentialPoint(date string, ) *RateDifferentialPoint`

NewRateDifferentialPoint instantiates a new RateDifferentialPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateDifferentialPointWithDefaults

`func NewRateDifferentialPointWithDefaults() *RateDifferentialPoint`

NewRateDifferentialPointWithDefaults instantiates a new RateDifferentialPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *RateDifferentialPoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RateDifferentialPoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RateDifferentialPoint) SetDate(v string)`

SetDate sets Date field to given value.


### GetBaseVal

`func (o *RateDifferentialPoint) GetBaseVal() float64`

GetBaseVal returns the BaseVal field if non-nil, zero value otherwise.

### GetBaseValOk

`func (o *RateDifferentialPoint) GetBaseValOk() (*float64, bool)`

GetBaseValOk returns a tuple with the BaseVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVal

`func (o *RateDifferentialPoint) SetBaseVal(v float64)`

SetBaseVal sets BaseVal field to given value.

### HasBaseVal

`func (o *RateDifferentialPoint) HasBaseVal() bool`

HasBaseVal returns a boolean if a field has been set.

### SetBaseValNil

`func (o *RateDifferentialPoint) SetBaseValNil(b bool)`

 SetBaseValNil sets the value for BaseVal to be an explicit nil

### UnsetBaseVal
`func (o *RateDifferentialPoint) UnsetBaseVal()`

UnsetBaseVal ensures that no value is present for BaseVal, not even an explicit nil
### GetQuoteVal

`func (o *RateDifferentialPoint) GetQuoteVal() float64`

GetQuoteVal returns the QuoteVal field if non-nil, zero value otherwise.

### GetQuoteValOk

`func (o *RateDifferentialPoint) GetQuoteValOk() (*float64, bool)`

GetQuoteValOk returns a tuple with the QuoteVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVal

`func (o *RateDifferentialPoint) SetQuoteVal(v float64)`

SetQuoteVal sets QuoteVal field to given value.

### HasQuoteVal

`func (o *RateDifferentialPoint) HasQuoteVal() bool`

HasQuoteVal returns a boolean if a field has been set.

### SetQuoteValNil

`func (o *RateDifferentialPoint) SetQuoteValNil(b bool)`

 SetQuoteValNil sets the value for QuoteVal to be an explicit nil

### UnsetQuoteVal
`func (o *RateDifferentialPoint) UnsetQuoteVal()`

UnsetQuoteVal ensures that no value is present for QuoteVal, not even an explicit nil
### GetSpread

`func (o *RateDifferentialPoint) GetSpread() float64`

GetSpread returns the Spread field if non-nil, zero value otherwise.

### GetSpreadOk

`func (o *RateDifferentialPoint) GetSpreadOk() (*float64, bool)`

GetSpreadOk returns a tuple with the Spread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpread

`func (o *RateDifferentialPoint) SetSpread(v float64)`

SetSpread sets Spread field to given value.

### HasSpread

`func (o *RateDifferentialPoint) HasSpread() bool`

HasSpread returns a boolean if a field has been set.

### SetSpreadNil

`func (o *RateDifferentialPoint) SetSpreadNil(b bool)`

 SetSpreadNil sets the value for Spread to be an explicit nil

### UnsetSpread
`func (o *RateDifferentialPoint) UnsetSpread()`

UnsetSpread ensures that no value is present for Spread, not even an explicit nil
### GetSpreadBps

`func (o *RateDifferentialPoint) GetSpreadBps() float64`

GetSpreadBps returns the SpreadBps field if non-nil, zero value otherwise.

### GetSpreadBpsOk

`func (o *RateDifferentialPoint) GetSpreadBpsOk() (*float64, bool)`

GetSpreadBpsOk returns a tuple with the SpreadBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadBps

`func (o *RateDifferentialPoint) SetSpreadBps(v float64)`

SetSpreadBps sets SpreadBps field to given value.

### HasSpreadBps

`func (o *RateDifferentialPoint) HasSpreadBps() bool`

HasSpreadBps returns a boolean if a field has been set.

### SetSpreadBpsNil

`func (o *RateDifferentialPoint) SetSpreadBpsNil(b bool)`

 SetSpreadBpsNil sets the value for SpreadBps to be an explicit nil

### UnsetSpreadBps
`func (o *RateDifferentialPoint) UnsetSpreadBps()`

UnsetSpreadBps ensures that no value is present for SpreadBps, not even an explicit nil
### GetDifferential

`func (o *RateDifferentialPoint) GetDifferential() float64`

GetDifferential returns the Differential field if non-nil, zero value otherwise.

### GetDifferentialOk

`func (o *RateDifferentialPoint) GetDifferentialOk() (*float64, bool)`

GetDifferentialOk returns a tuple with the Differential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferential

`func (o *RateDifferentialPoint) SetDifferential(v float64)`

SetDifferential sets Differential field to given value.

### HasDifferential

`func (o *RateDifferentialPoint) HasDifferential() bool`

HasDifferential returns a boolean if a field has been set.

### SetDifferentialNil

`func (o *RateDifferentialPoint) SetDifferentialNil(b bool)`

 SetDifferentialNil sets the value for Differential to be an explicit nil

### UnsetDifferential
`func (o *RateDifferentialPoint) UnsetDifferential()`

UnsetDifferential ensures that no value is present for Differential, not even an explicit nil
### GetDifferentialBps

`func (o *RateDifferentialPoint) GetDifferentialBps() float64`

GetDifferentialBps returns the DifferentialBps field if non-nil, zero value otherwise.

### GetDifferentialBpsOk

`func (o *RateDifferentialPoint) GetDifferentialBpsOk() (*float64, bool)`

GetDifferentialBpsOk returns a tuple with the DifferentialBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferentialBps

`func (o *RateDifferentialPoint) SetDifferentialBps(v float64)`

SetDifferentialBps sets DifferentialBps field to given value.

### HasDifferentialBps

`func (o *RateDifferentialPoint) HasDifferentialBps() bool`

HasDifferentialBps returns a boolean if a field has been set.

### SetDifferentialBpsNil

`func (o *RateDifferentialPoint) SetDifferentialBpsNil(b bool)`

 SetDifferentialBpsNil sets the value for DifferentialBps to be an explicit nil

### UnsetDifferentialBps
`func (o *RateDifferentialPoint) UnsetDifferentialBps()`

UnsetDifferentialBps ensures that no value is present for DifferentialBps, not even an explicit nil
### GetBaseForwardVal

`func (o *RateDifferentialPoint) GetBaseForwardVal() float64`

GetBaseForwardVal returns the BaseForwardVal field if non-nil, zero value otherwise.

### GetBaseForwardValOk

`func (o *RateDifferentialPoint) GetBaseForwardValOk() (*float64, bool)`

GetBaseForwardValOk returns a tuple with the BaseForwardVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseForwardVal

`func (o *RateDifferentialPoint) SetBaseForwardVal(v float64)`

SetBaseForwardVal sets BaseForwardVal field to given value.

### HasBaseForwardVal

`func (o *RateDifferentialPoint) HasBaseForwardVal() bool`

HasBaseForwardVal returns a boolean if a field has been set.

### SetBaseForwardValNil

`func (o *RateDifferentialPoint) SetBaseForwardValNil(b bool)`

 SetBaseForwardValNil sets the value for BaseForwardVal to be an explicit nil

### UnsetBaseForwardVal
`func (o *RateDifferentialPoint) UnsetBaseForwardVal()`

UnsetBaseForwardVal ensures that no value is present for BaseForwardVal, not even an explicit nil
### GetQuoteForwardVal

`func (o *RateDifferentialPoint) GetQuoteForwardVal() float64`

GetQuoteForwardVal returns the QuoteForwardVal field if non-nil, zero value otherwise.

### GetQuoteForwardValOk

`func (o *RateDifferentialPoint) GetQuoteForwardValOk() (*float64, bool)`

GetQuoteForwardValOk returns a tuple with the QuoteForwardVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteForwardVal

`func (o *RateDifferentialPoint) SetQuoteForwardVal(v float64)`

SetQuoteForwardVal sets QuoteForwardVal field to given value.

### HasQuoteForwardVal

`func (o *RateDifferentialPoint) HasQuoteForwardVal() bool`

HasQuoteForwardVal returns a boolean if a field has been set.

### SetQuoteForwardValNil

`func (o *RateDifferentialPoint) SetQuoteForwardValNil(b bool)`

 SetQuoteForwardValNil sets the value for QuoteForwardVal to be an explicit nil

### UnsetQuoteForwardVal
`func (o *RateDifferentialPoint) UnsetQuoteForwardVal()`

UnsetQuoteForwardVal ensures that no value is present for QuoteForwardVal, not even an explicit nil
### GetBaseStartVal

`func (o *RateDifferentialPoint) GetBaseStartVal() float64`

GetBaseStartVal returns the BaseStartVal field if non-nil, zero value otherwise.

### GetBaseStartValOk

`func (o *RateDifferentialPoint) GetBaseStartValOk() (*float64, bool)`

GetBaseStartValOk returns a tuple with the BaseStartVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseStartVal

`func (o *RateDifferentialPoint) SetBaseStartVal(v float64)`

SetBaseStartVal sets BaseStartVal field to given value.

### HasBaseStartVal

`func (o *RateDifferentialPoint) HasBaseStartVal() bool`

HasBaseStartVal returns a boolean if a field has been set.

### SetBaseStartValNil

`func (o *RateDifferentialPoint) SetBaseStartValNil(b bool)`

 SetBaseStartValNil sets the value for BaseStartVal to be an explicit nil

### UnsetBaseStartVal
`func (o *RateDifferentialPoint) UnsetBaseStartVal()`

UnsetBaseStartVal ensures that no value is present for BaseStartVal, not even an explicit nil
### GetBaseEndVal

`func (o *RateDifferentialPoint) GetBaseEndVal() float64`

GetBaseEndVal returns the BaseEndVal field if non-nil, zero value otherwise.

### GetBaseEndValOk

`func (o *RateDifferentialPoint) GetBaseEndValOk() (*float64, bool)`

GetBaseEndValOk returns a tuple with the BaseEndVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseEndVal

`func (o *RateDifferentialPoint) SetBaseEndVal(v float64)`

SetBaseEndVal sets BaseEndVal field to given value.

### HasBaseEndVal

`func (o *RateDifferentialPoint) HasBaseEndVal() bool`

HasBaseEndVal returns a boolean if a field has been set.

### SetBaseEndValNil

`func (o *RateDifferentialPoint) SetBaseEndValNil(b bool)`

 SetBaseEndValNil sets the value for BaseEndVal to be an explicit nil

### UnsetBaseEndVal
`func (o *RateDifferentialPoint) UnsetBaseEndVal()`

UnsetBaseEndVal ensures that no value is present for BaseEndVal, not even an explicit nil
### GetQuoteStartVal

`func (o *RateDifferentialPoint) GetQuoteStartVal() float64`

GetQuoteStartVal returns the QuoteStartVal field if non-nil, zero value otherwise.

### GetQuoteStartValOk

`func (o *RateDifferentialPoint) GetQuoteStartValOk() (*float64, bool)`

GetQuoteStartValOk returns a tuple with the QuoteStartVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteStartVal

`func (o *RateDifferentialPoint) SetQuoteStartVal(v float64)`

SetQuoteStartVal sets QuoteStartVal field to given value.

### HasQuoteStartVal

`func (o *RateDifferentialPoint) HasQuoteStartVal() bool`

HasQuoteStartVal returns a boolean if a field has been set.

### SetQuoteStartValNil

`func (o *RateDifferentialPoint) SetQuoteStartValNil(b bool)`

 SetQuoteStartValNil sets the value for QuoteStartVal to be an explicit nil

### UnsetQuoteStartVal
`func (o *RateDifferentialPoint) UnsetQuoteStartVal()`

UnsetQuoteStartVal ensures that no value is present for QuoteStartVal, not even an explicit nil
### GetQuoteEndVal

`func (o *RateDifferentialPoint) GetQuoteEndVal() float64`

GetQuoteEndVal returns the QuoteEndVal field if non-nil, zero value otherwise.

### GetQuoteEndValOk

`func (o *RateDifferentialPoint) GetQuoteEndValOk() (*float64, bool)`

GetQuoteEndValOk returns a tuple with the QuoteEndVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteEndVal

`func (o *RateDifferentialPoint) SetQuoteEndVal(v float64)`

SetQuoteEndVal sets QuoteEndVal field to given value.

### HasQuoteEndVal

`func (o *RateDifferentialPoint) HasQuoteEndVal() bool`

HasQuoteEndVal returns a boolean if a field has been set.

### SetQuoteEndValNil

`func (o *RateDifferentialPoint) SetQuoteEndValNil(b bool)`

 SetQuoteEndValNil sets the value for QuoteEndVal to be an explicit nil

### UnsetQuoteEndVal
`func (o *RateDifferentialPoint) UnsetQuoteEndVal()`

UnsetQuoteEndVal ensures that no value is present for QuoteEndVal, not even an explicit nil
### GetBaseAnnouncementDatetime

`func (o *RateDifferentialPoint) GetBaseAnnouncementDatetime() int64`

GetBaseAnnouncementDatetime returns the BaseAnnouncementDatetime field if non-nil, zero value otherwise.

### GetBaseAnnouncementDatetimeOk

`func (o *RateDifferentialPoint) GetBaseAnnouncementDatetimeOk() (*int64, bool)`

GetBaseAnnouncementDatetimeOk returns a tuple with the BaseAnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAnnouncementDatetime

`func (o *RateDifferentialPoint) SetBaseAnnouncementDatetime(v int64)`

SetBaseAnnouncementDatetime sets BaseAnnouncementDatetime field to given value.

### HasBaseAnnouncementDatetime

`func (o *RateDifferentialPoint) HasBaseAnnouncementDatetime() bool`

HasBaseAnnouncementDatetime returns a boolean if a field has been set.

### SetBaseAnnouncementDatetimeNil

`func (o *RateDifferentialPoint) SetBaseAnnouncementDatetimeNil(b bool)`

 SetBaseAnnouncementDatetimeNil sets the value for BaseAnnouncementDatetime to be an explicit nil

### UnsetBaseAnnouncementDatetime
`func (o *RateDifferentialPoint) UnsetBaseAnnouncementDatetime()`

UnsetBaseAnnouncementDatetime ensures that no value is present for BaseAnnouncementDatetime, not even an explicit nil
### GetQuoteAnnouncementDatetime

`func (o *RateDifferentialPoint) GetQuoteAnnouncementDatetime() int64`

GetQuoteAnnouncementDatetime returns the QuoteAnnouncementDatetime field if non-nil, zero value otherwise.

### GetQuoteAnnouncementDatetimeOk

`func (o *RateDifferentialPoint) GetQuoteAnnouncementDatetimeOk() (*int64, bool)`

GetQuoteAnnouncementDatetimeOk returns a tuple with the QuoteAnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAnnouncementDatetime

`func (o *RateDifferentialPoint) SetQuoteAnnouncementDatetime(v int64)`

SetQuoteAnnouncementDatetime sets QuoteAnnouncementDatetime field to given value.

### HasQuoteAnnouncementDatetime

`func (o *RateDifferentialPoint) HasQuoteAnnouncementDatetime() bool`

HasQuoteAnnouncementDatetime returns a boolean if a field has been set.

### SetQuoteAnnouncementDatetimeNil

`func (o *RateDifferentialPoint) SetQuoteAnnouncementDatetimeNil(b bool)`

 SetQuoteAnnouncementDatetimeNil sets the value for QuoteAnnouncementDatetime to be an explicit nil

### UnsetQuoteAnnouncementDatetime
`func (o *RateDifferentialPoint) UnsetQuoteAnnouncementDatetime()`

UnsetQuoteAnnouncementDatetime ensures that no value is present for QuoteAnnouncementDatetime, not even an explicit nil
### GetBaseStartAnnouncementDatetime

`func (o *RateDifferentialPoint) GetBaseStartAnnouncementDatetime() int64`

GetBaseStartAnnouncementDatetime returns the BaseStartAnnouncementDatetime field if non-nil, zero value otherwise.

### GetBaseStartAnnouncementDatetimeOk

`func (o *RateDifferentialPoint) GetBaseStartAnnouncementDatetimeOk() (*int64, bool)`

GetBaseStartAnnouncementDatetimeOk returns a tuple with the BaseStartAnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseStartAnnouncementDatetime

`func (o *RateDifferentialPoint) SetBaseStartAnnouncementDatetime(v int64)`

SetBaseStartAnnouncementDatetime sets BaseStartAnnouncementDatetime field to given value.

### HasBaseStartAnnouncementDatetime

`func (o *RateDifferentialPoint) HasBaseStartAnnouncementDatetime() bool`

HasBaseStartAnnouncementDatetime returns a boolean if a field has been set.

### SetBaseStartAnnouncementDatetimeNil

`func (o *RateDifferentialPoint) SetBaseStartAnnouncementDatetimeNil(b bool)`

 SetBaseStartAnnouncementDatetimeNil sets the value for BaseStartAnnouncementDatetime to be an explicit nil

### UnsetBaseStartAnnouncementDatetime
`func (o *RateDifferentialPoint) UnsetBaseStartAnnouncementDatetime()`

UnsetBaseStartAnnouncementDatetime ensures that no value is present for BaseStartAnnouncementDatetime, not even an explicit nil
### GetBaseEndAnnouncementDatetime

`func (o *RateDifferentialPoint) GetBaseEndAnnouncementDatetime() int64`

GetBaseEndAnnouncementDatetime returns the BaseEndAnnouncementDatetime field if non-nil, zero value otherwise.

### GetBaseEndAnnouncementDatetimeOk

`func (o *RateDifferentialPoint) GetBaseEndAnnouncementDatetimeOk() (*int64, bool)`

GetBaseEndAnnouncementDatetimeOk returns a tuple with the BaseEndAnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseEndAnnouncementDatetime

`func (o *RateDifferentialPoint) SetBaseEndAnnouncementDatetime(v int64)`

SetBaseEndAnnouncementDatetime sets BaseEndAnnouncementDatetime field to given value.

### HasBaseEndAnnouncementDatetime

`func (o *RateDifferentialPoint) HasBaseEndAnnouncementDatetime() bool`

HasBaseEndAnnouncementDatetime returns a boolean if a field has been set.

### SetBaseEndAnnouncementDatetimeNil

`func (o *RateDifferentialPoint) SetBaseEndAnnouncementDatetimeNil(b bool)`

 SetBaseEndAnnouncementDatetimeNil sets the value for BaseEndAnnouncementDatetime to be an explicit nil

### UnsetBaseEndAnnouncementDatetime
`func (o *RateDifferentialPoint) UnsetBaseEndAnnouncementDatetime()`

UnsetBaseEndAnnouncementDatetime ensures that no value is present for BaseEndAnnouncementDatetime, not even an explicit nil
### GetQuoteStartAnnouncementDatetime

`func (o *RateDifferentialPoint) GetQuoteStartAnnouncementDatetime() int64`

GetQuoteStartAnnouncementDatetime returns the QuoteStartAnnouncementDatetime field if non-nil, zero value otherwise.

### GetQuoteStartAnnouncementDatetimeOk

`func (o *RateDifferentialPoint) GetQuoteStartAnnouncementDatetimeOk() (*int64, bool)`

GetQuoteStartAnnouncementDatetimeOk returns a tuple with the QuoteStartAnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteStartAnnouncementDatetime

`func (o *RateDifferentialPoint) SetQuoteStartAnnouncementDatetime(v int64)`

SetQuoteStartAnnouncementDatetime sets QuoteStartAnnouncementDatetime field to given value.

### HasQuoteStartAnnouncementDatetime

`func (o *RateDifferentialPoint) HasQuoteStartAnnouncementDatetime() bool`

HasQuoteStartAnnouncementDatetime returns a boolean if a field has been set.

### SetQuoteStartAnnouncementDatetimeNil

`func (o *RateDifferentialPoint) SetQuoteStartAnnouncementDatetimeNil(b bool)`

 SetQuoteStartAnnouncementDatetimeNil sets the value for QuoteStartAnnouncementDatetime to be an explicit nil

### UnsetQuoteStartAnnouncementDatetime
`func (o *RateDifferentialPoint) UnsetQuoteStartAnnouncementDatetime()`

UnsetQuoteStartAnnouncementDatetime ensures that no value is present for QuoteStartAnnouncementDatetime, not even an explicit nil
### GetQuoteEndAnnouncementDatetime

`func (o *RateDifferentialPoint) GetQuoteEndAnnouncementDatetime() int64`

GetQuoteEndAnnouncementDatetime returns the QuoteEndAnnouncementDatetime field if non-nil, zero value otherwise.

### GetQuoteEndAnnouncementDatetimeOk

`func (o *RateDifferentialPoint) GetQuoteEndAnnouncementDatetimeOk() (*int64, bool)`

GetQuoteEndAnnouncementDatetimeOk returns a tuple with the QuoteEndAnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteEndAnnouncementDatetime

`func (o *RateDifferentialPoint) SetQuoteEndAnnouncementDatetime(v int64)`

SetQuoteEndAnnouncementDatetime sets QuoteEndAnnouncementDatetime field to given value.

### HasQuoteEndAnnouncementDatetime

`func (o *RateDifferentialPoint) HasQuoteEndAnnouncementDatetime() bool`

HasQuoteEndAnnouncementDatetime returns a boolean if a field has been set.

### SetQuoteEndAnnouncementDatetimeNil

`func (o *RateDifferentialPoint) SetQuoteEndAnnouncementDatetimeNil(b bool)`

 SetQuoteEndAnnouncementDatetimeNil sets the value for QuoteEndAnnouncementDatetime to be an explicit nil

### UnsetQuoteEndAnnouncementDatetime
`func (o *RateDifferentialPoint) UnsetQuoteEndAnnouncementDatetime()`

UnsetQuoteEndAnnouncementDatetime ensures that no value is present for QuoteEndAnnouncementDatetime, not even an explicit nil
### GetBaseSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) GetBaseSelectedVintageKnownAtNs() string`

GetBaseSelectedVintageKnownAtNs returns the BaseSelectedVintageKnownAtNs field if non-nil, zero value otherwise.

### GetBaseSelectedVintageKnownAtNsOk

`func (o *RateDifferentialPoint) GetBaseSelectedVintageKnownAtNsOk() (*string, bool)`

GetBaseSelectedVintageKnownAtNsOk returns a tuple with the BaseSelectedVintageKnownAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) SetBaseSelectedVintageKnownAtNs(v string)`

SetBaseSelectedVintageKnownAtNs sets BaseSelectedVintageKnownAtNs field to given value.

### HasBaseSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) HasBaseSelectedVintageKnownAtNs() bool`

HasBaseSelectedVintageKnownAtNs returns a boolean if a field has been set.

### SetBaseSelectedVintageKnownAtNsNil

`func (o *RateDifferentialPoint) SetBaseSelectedVintageKnownAtNsNil(b bool)`

 SetBaseSelectedVintageKnownAtNsNil sets the value for BaseSelectedVintageKnownAtNs to be an explicit nil

### UnsetBaseSelectedVintageKnownAtNs
`func (o *RateDifferentialPoint) UnsetBaseSelectedVintageKnownAtNs()`

UnsetBaseSelectedVintageKnownAtNs ensures that no value is present for BaseSelectedVintageKnownAtNs, not even an explicit nil
### GetQuoteSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) GetQuoteSelectedVintageKnownAtNs() string`

GetQuoteSelectedVintageKnownAtNs returns the QuoteSelectedVintageKnownAtNs field if non-nil, zero value otherwise.

### GetQuoteSelectedVintageKnownAtNsOk

`func (o *RateDifferentialPoint) GetQuoteSelectedVintageKnownAtNsOk() (*string, bool)`

GetQuoteSelectedVintageKnownAtNsOk returns a tuple with the QuoteSelectedVintageKnownAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) SetQuoteSelectedVintageKnownAtNs(v string)`

SetQuoteSelectedVintageKnownAtNs sets QuoteSelectedVintageKnownAtNs field to given value.

### HasQuoteSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) HasQuoteSelectedVintageKnownAtNs() bool`

HasQuoteSelectedVintageKnownAtNs returns a boolean if a field has been set.

### SetQuoteSelectedVintageKnownAtNsNil

`func (o *RateDifferentialPoint) SetQuoteSelectedVintageKnownAtNsNil(b bool)`

 SetQuoteSelectedVintageKnownAtNsNil sets the value for QuoteSelectedVintageKnownAtNs to be an explicit nil

### UnsetQuoteSelectedVintageKnownAtNs
`func (o *RateDifferentialPoint) UnsetQuoteSelectedVintageKnownAtNs()`

UnsetQuoteSelectedVintageKnownAtNs ensures that no value is present for QuoteSelectedVintageKnownAtNs, not even an explicit nil
### GetBaseStartSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) GetBaseStartSelectedVintageKnownAtNs() string`

GetBaseStartSelectedVintageKnownAtNs returns the BaseStartSelectedVintageKnownAtNs field if non-nil, zero value otherwise.

### GetBaseStartSelectedVintageKnownAtNsOk

`func (o *RateDifferentialPoint) GetBaseStartSelectedVintageKnownAtNsOk() (*string, bool)`

GetBaseStartSelectedVintageKnownAtNsOk returns a tuple with the BaseStartSelectedVintageKnownAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseStartSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) SetBaseStartSelectedVintageKnownAtNs(v string)`

SetBaseStartSelectedVintageKnownAtNs sets BaseStartSelectedVintageKnownAtNs field to given value.

### HasBaseStartSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) HasBaseStartSelectedVintageKnownAtNs() bool`

HasBaseStartSelectedVintageKnownAtNs returns a boolean if a field has been set.

### SetBaseStartSelectedVintageKnownAtNsNil

`func (o *RateDifferentialPoint) SetBaseStartSelectedVintageKnownAtNsNil(b bool)`

 SetBaseStartSelectedVintageKnownAtNsNil sets the value for BaseStartSelectedVintageKnownAtNs to be an explicit nil

### UnsetBaseStartSelectedVintageKnownAtNs
`func (o *RateDifferentialPoint) UnsetBaseStartSelectedVintageKnownAtNs()`

UnsetBaseStartSelectedVintageKnownAtNs ensures that no value is present for BaseStartSelectedVintageKnownAtNs, not even an explicit nil
### GetBaseEndSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) GetBaseEndSelectedVintageKnownAtNs() string`

GetBaseEndSelectedVintageKnownAtNs returns the BaseEndSelectedVintageKnownAtNs field if non-nil, zero value otherwise.

### GetBaseEndSelectedVintageKnownAtNsOk

`func (o *RateDifferentialPoint) GetBaseEndSelectedVintageKnownAtNsOk() (*string, bool)`

GetBaseEndSelectedVintageKnownAtNsOk returns a tuple with the BaseEndSelectedVintageKnownAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseEndSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) SetBaseEndSelectedVintageKnownAtNs(v string)`

SetBaseEndSelectedVintageKnownAtNs sets BaseEndSelectedVintageKnownAtNs field to given value.

### HasBaseEndSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) HasBaseEndSelectedVintageKnownAtNs() bool`

HasBaseEndSelectedVintageKnownAtNs returns a boolean if a field has been set.

### SetBaseEndSelectedVintageKnownAtNsNil

`func (o *RateDifferentialPoint) SetBaseEndSelectedVintageKnownAtNsNil(b bool)`

 SetBaseEndSelectedVintageKnownAtNsNil sets the value for BaseEndSelectedVintageKnownAtNs to be an explicit nil

### UnsetBaseEndSelectedVintageKnownAtNs
`func (o *RateDifferentialPoint) UnsetBaseEndSelectedVintageKnownAtNs()`

UnsetBaseEndSelectedVintageKnownAtNs ensures that no value is present for BaseEndSelectedVintageKnownAtNs, not even an explicit nil
### GetQuoteStartSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) GetQuoteStartSelectedVintageKnownAtNs() string`

GetQuoteStartSelectedVintageKnownAtNs returns the QuoteStartSelectedVintageKnownAtNs field if non-nil, zero value otherwise.

### GetQuoteStartSelectedVintageKnownAtNsOk

`func (o *RateDifferentialPoint) GetQuoteStartSelectedVintageKnownAtNsOk() (*string, bool)`

GetQuoteStartSelectedVintageKnownAtNsOk returns a tuple with the QuoteStartSelectedVintageKnownAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteStartSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) SetQuoteStartSelectedVintageKnownAtNs(v string)`

SetQuoteStartSelectedVintageKnownAtNs sets QuoteStartSelectedVintageKnownAtNs field to given value.

### HasQuoteStartSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) HasQuoteStartSelectedVintageKnownAtNs() bool`

HasQuoteStartSelectedVintageKnownAtNs returns a boolean if a field has been set.

### SetQuoteStartSelectedVintageKnownAtNsNil

`func (o *RateDifferentialPoint) SetQuoteStartSelectedVintageKnownAtNsNil(b bool)`

 SetQuoteStartSelectedVintageKnownAtNsNil sets the value for QuoteStartSelectedVintageKnownAtNs to be an explicit nil

### UnsetQuoteStartSelectedVintageKnownAtNs
`func (o *RateDifferentialPoint) UnsetQuoteStartSelectedVintageKnownAtNs()`

UnsetQuoteStartSelectedVintageKnownAtNs ensures that no value is present for QuoteStartSelectedVintageKnownAtNs, not even an explicit nil
### GetQuoteEndSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) GetQuoteEndSelectedVintageKnownAtNs() string`

GetQuoteEndSelectedVintageKnownAtNs returns the QuoteEndSelectedVintageKnownAtNs field if non-nil, zero value otherwise.

### GetQuoteEndSelectedVintageKnownAtNsOk

`func (o *RateDifferentialPoint) GetQuoteEndSelectedVintageKnownAtNsOk() (*string, bool)`

GetQuoteEndSelectedVintageKnownAtNsOk returns a tuple with the QuoteEndSelectedVintageKnownAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteEndSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) SetQuoteEndSelectedVintageKnownAtNs(v string)`

SetQuoteEndSelectedVintageKnownAtNs sets QuoteEndSelectedVintageKnownAtNs field to given value.

### HasQuoteEndSelectedVintageKnownAtNs

`func (o *RateDifferentialPoint) HasQuoteEndSelectedVintageKnownAtNs() bool`

HasQuoteEndSelectedVintageKnownAtNs returns a boolean if a field has been set.

### SetQuoteEndSelectedVintageKnownAtNsNil

`func (o *RateDifferentialPoint) SetQuoteEndSelectedVintageKnownAtNsNil(b bool)`

 SetQuoteEndSelectedVintageKnownAtNsNil sets the value for QuoteEndSelectedVintageKnownAtNs to be an explicit nil

### UnsetQuoteEndSelectedVintageKnownAtNs
`func (o *RateDifferentialPoint) UnsetQuoteEndSelectedVintageKnownAtNs()`

UnsetQuoteEndSelectedVintageKnownAtNs ensures that no value is present for QuoteEndSelectedVintageKnownAtNs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


