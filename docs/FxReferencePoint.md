# FxReferencePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**ReferenceDate** | **string** |  | 
**Base** | **string** |  | 
**Quote** | **string** |  | 
**Pair** | **string** |  | 
**Rate** | **float64** |  | 
**TimestampUtc** | Pointer to **NullableString** |  | [optional] 
**TimestampType** | Pointer to **NullableString** |  | [optional] 
**FixingTimeUtc** | Pointer to **NullableString** |  | [optional] 
**PublishedAtUtc** | Pointer to **NullableString** |  | [optional] 
**TimePrecision** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to [**ForexPointSource**](ForexPointSource.md) |  | [optional] 
**NativeOrDerived** | Pointer to **NullableString** |  | [optional] 
**DerivationMethod** | Pointer to **NullableString** |  | [optional] 
**SourcePair** | Pointer to **NullableString** |  | [optional] 
**SourceLegs** | Pointer to [**[]FxSourceLeg**](FxSourceLeg.md) |  | [optional] 
**StalenessMinutes** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewFxReferencePoint

`func NewFxReferencePoint(date string, referenceDate string, base string, quote string, pair string, rate float64, ) *FxReferencePoint`

NewFxReferencePoint instantiates a new FxReferencePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxReferencePointWithDefaults

`func NewFxReferencePointWithDefaults() *FxReferencePoint`

NewFxReferencePointWithDefaults instantiates a new FxReferencePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *FxReferencePoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FxReferencePoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FxReferencePoint) SetDate(v string)`

SetDate sets Date field to given value.


### GetReferenceDate

`func (o *FxReferencePoint) GetReferenceDate() string`

GetReferenceDate returns the ReferenceDate field if non-nil, zero value otherwise.

### GetReferenceDateOk

`func (o *FxReferencePoint) GetReferenceDateOk() (*string, bool)`

GetReferenceDateOk returns a tuple with the ReferenceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceDate

`func (o *FxReferencePoint) SetReferenceDate(v string)`

SetReferenceDate sets ReferenceDate field to given value.


### GetBase

`func (o *FxReferencePoint) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *FxReferencePoint) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *FxReferencePoint) SetBase(v string)`

SetBase sets Base field to given value.


### GetQuote

`func (o *FxReferencePoint) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *FxReferencePoint) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *FxReferencePoint) SetQuote(v string)`

SetQuote sets Quote field to given value.


### GetPair

`func (o *FxReferencePoint) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *FxReferencePoint) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *FxReferencePoint) SetPair(v string)`

SetPair sets Pair field to given value.


### GetRate

`func (o *FxReferencePoint) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *FxReferencePoint) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *FxReferencePoint) SetRate(v float64)`

SetRate sets Rate field to given value.


### GetTimestampUtc

`func (o *FxReferencePoint) GetTimestampUtc() string`

GetTimestampUtc returns the TimestampUtc field if non-nil, zero value otherwise.

### GetTimestampUtcOk

`func (o *FxReferencePoint) GetTimestampUtcOk() (*string, bool)`

GetTimestampUtcOk returns a tuple with the TimestampUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampUtc

`func (o *FxReferencePoint) SetTimestampUtc(v string)`

SetTimestampUtc sets TimestampUtc field to given value.

### HasTimestampUtc

`func (o *FxReferencePoint) HasTimestampUtc() bool`

HasTimestampUtc returns a boolean if a field has been set.

### SetTimestampUtcNil

`func (o *FxReferencePoint) SetTimestampUtcNil(b bool)`

 SetTimestampUtcNil sets the value for TimestampUtc to be an explicit nil

### UnsetTimestampUtc
`func (o *FxReferencePoint) UnsetTimestampUtc()`

UnsetTimestampUtc ensures that no value is present for TimestampUtc, not even an explicit nil
### GetTimestampType

`func (o *FxReferencePoint) GetTimestampType() string`

GetTimestampType returns the TimestampType field if non-nil, zero value otherwise.

### GetTimestampTypeOk

`func (o *FxReferencePoint) GetTimestampTypeOk() (*string, bool)`

GetTimestampTypeOk returns a tuple with the TimestampType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampType

`func (o *FxReferencePoint) SetTimestampType(v string)`

SetTimestampType sets TimestampType field to given value.

### HasTimestampType

`func (o *FxReferencePoint) HasTimestampType() bool`

HasTimestampType returns a boolean if a field has been set.

### SetTimestampTypeNil

`func (o *FxReferencePoint) SetTimestampTypeNil(b bool)`

 SetTimestampTypeNil sets the value for TimestampType to be an explicit nil

### UnsetTimestampType
`func (o *FxReferencePoint) UnsetTimestampType()`

UnsetTimestampType ensures that no value is present for TimestampType, not even an explicit nil
### GetFixingTimeUtc

`func (o *FxReferencePoint) GetFixingTimeUtc() string`

GetFixingTimeUtc returns the FixingTimeUtc field if non-nil, zero value otherwise.

### GetFixingTimeUtcOk

`func (o *FxReferencePoint) GetFixingTimeUtcOk() (*string, bool)`

GetFixingTimeUtcOk returns a tuple with the FixingTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixingTimeUtc

`func (o *FxReferencePoint) SetFixingTimeUtc(v string)`

SetFixingTimeUtc sets FixingTimeUtc field to given value.

### HasFixingTimeUtc

`func (o *FxReferencePoint) HasFixingTimeUtc() bool`

HasFixingTimeUtc returns a boolean if a field has been set.

### SetFixingTimeUtcNil

`func (o *FxReferencePoint) SetFixingTimeUtcNil(b bool)`

 SetFixingTimeUtcNil sets the value for FixingTimeUtc to be an explicit nil

### UnsetFixingTimeUtc
`func (o *FxReferencePoint) UnsetFixingTimeUtc()`

UnsetFixingTimeUtc ensures that no value is present for FixingTimeUtc, not even an explicit nil
### GetPublishedAtUtc

`func (o *FxReferencePoint) GetPublishedAtUtc() string`

GetPublishedAtUtc returns the PublishedAtUtc field if non-nil, zero value otherwise.

### GetPublishedAtUtcOk

`func (o *FxReferencePoint) GetPublishedAtUtcOk() (*string, bool)`

GetPublishedAtUtcOk returns a tuple with the PublishedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAtUtc

`func (o *FxReferencePoint) SetPublishedAtUtc(v string)`

SetPublishedAtUtc sets PublishedAtUtc field to given value.

### HasPublishedAtUtc

`func (o *FxReferencePoint) HasPublishedAtUtc() bool`

HasPublishedAtUtc returns a boolean if a field has been set.

### SetPublishedAtUtcNil

`func (o *FxReferencePoint) SetPublishedAtUtcNil(b bool)`

 SetPublishedAtUtcNil sets the value for PublishedAtUtc to be an explicit nil

### UnsetPublishedAtUtc
`func (o *FxReferencePoint) UnsetPublishedAtUtc()`

UnsetPublishedAtUtc ensures that no value is present for PublishedAtUtc, not even an explicit nil
### GetTimePrecision

`func (o *FxReferencePoint) GetTimePrecision() string`

GetTimePrecision returns the TimePrecision field if non-nil, zero value otherwise.

### GetTimePrecisionOk

`func (o *FxReferencePoint) GetTimePrecisionOk() (*string, bool)`

GetTimePrecisionOk returns a tuple with the TimePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePrecision

`func (o *FxReferencePoint) SetTimePrecision(v string)`

SetTimePrecision sets TimePrecision field to given value.

### HasTimePrecision

`func (o *FxReferencePoint) HasTimePrecision() bool`

HasTimePrecision returns a boolean if a field has been set.

### SetTimePrecisionNil

`func (o *FxReferencePoint) SetTimePrecisionNil(b bool)`

 SetTimePrecisionNil sets the value for TimePrecision to be an explicit nil

### UnsetTimePrecision
`func (o *FxReferencePoint) UnsetTimePrecision()`

UnsetTimePrecision ensures that no value is present for TimePrecision, not even an explicit nil
### GetSource

`func (o *FxReferencePoint) GetSource() ForexPointSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FxReferencePoint) GetSourceOk() (*ForexPointSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FxReferencePoint) SetSource(v ForexPointSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *FxReferencePoint) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetNativeOrDerived

`func (o *FxReferencePoint) GetNativeOrDerived() string`

GetNativeOrDerived returns the NativeOrDerived field if non-nil, zero value otherwise.

### GetNativeOrDerivedOk

`func (o *FxReferencePoint) GetNativeOrDerivedOk() (*string, bool)`

GetNativeOrDerivedOk returns a tuple with the NativeOrDerived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeOrDerived

`func (o *FxReferencePoint) SetNativeOrDerived(v string)`

SetNativeOrDerived sets NativeOrDerived field to given value.

### HasNativeOrDerived

`func (o *FxReferencePoint) HasNativeOrDerived() bool`

HasNativeOrDerived returns a boolean if a field has been set.

### SetNativeOrDerivedNil

`func (o *FxReferencePoint) SetNativeOrDerivedNil(b bool)`

 SetNativeOrDerivedNil sets the value for NativeOrDerived to be an explicit nil

### UnsetNativeOrDerived
`func (o *FxReferencePoint) UnsetNativeOrDerived()`

UnsetNativeOrDerived ensures that no value is present for NativeOrDerived, not even an explicit nil
### GetDerivationMethod

`func (o *FxReferencePoint) GetDerivationMethod() string`

GetDerivationMethod returns the DerivationMethod field if non-nil, zero value otherwise.

### GetDerivationMethodOk

`func (o *FxReferencePoint) GetDerivationMethodOk() (*string, bool)`

GetDerivationMethodOk returns a tuple with the DerivationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivationMethod

`func (o *FxReferencePoint) SetDerivationMethod(v string)`

SetDerivationMethod sets DerivationMethod field to given value.

### HasDerivationMethod

`func (o *FxReferencePoint) HasDerivationMethod() bool`

HasDerivationMethod returns a boolean if a field has been set.

### SetDerivationMethodNil

`func (o *FxReferencePoint) SetDerivationMethodNil(b bool)`

 SetDerivationMethodNil sets the value for DerivationMethod to be an explicit nil

### UnsetDerivationMethod
`func (o *FxReferencePoint) UnsetDerivationMethod()`

UnsetDerivationMethod ensures that no value is present for DerivationMethod, not even an explicit nil
### GetSourcePair

`func (o *FxReferencePoint) GetSourcePair() string`

GetSourcePair returns the SourcePair field if non-nil, zero value otherwise.

### GetSourcePairOk

`func (o *FxReferencePoint) GetSourcePairOk() (*string, bool)`

GetSourcePairOk returns a tuple with the SourcePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePair

`func (o *FxReferencePoint) SetSourcePair(v string)`

SetSourcePair sets SourcePair field to given value.

### HasSourcePair

`func (o *FxReferencePoint) HasSourcePair() bool`

HasSourcePair returns a boolean if a field has been set.

### SetSourcePairNil

`func (o *FxReferencePoint) SetSourcePairNil(b bool)`

 SetSourcePairNil sets the value for SourcePair to be an explicit nil

### UnsetSourcePair
`func (o *FxReferencePoint) UnsetSourcePair()`

UnsetSourcePair ensures that no value is present for SourcePair, not even an explicit nil
### GetSourceLegs

`func (o *FxReferencePoint) GetSourceLegs() []FxSourceLeg`

GetSourceLegs returns the SourceLegs field if non-nil, zero value otherwise.

### GetSourceLegsOk

`func (o *FxReferencePoint) GetSourceLegsOk() (*[]FxSourceLeg, bool)`

GetSourceLegsOk returns a tuple with the SourceLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLegs

`func (o *FxReferencePoint) SetSourceLegs(v []FxSourceLeg)`

SetSourceLegs sets SourceLegs field to given value.

### HasSourceLegs

`func (o *FxReferencePoint) HasSourceLegs() bool`

HasSourceLegs returns a boolean if a field has been set.

### GetStalenessMinutes

`func (o *FxReferencePoint) GetStalenessMinutes() int64`

GetStalenessMinutes returns the StalenessMinutes field if non-nil, zero value otherwise.

### GetStalenessMinutesOk

`func (o *FxReferencePoint) GetStalenessMinutesOk() (*int64, bool)`

GetStalenessMinutesOk returns a tuple with the StalenessMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStalenessMinutes

`func (o *FxReferencePoint) SetStalenessMinutes(v int64)`

SetStalenessMinutes sets StalenessMinutes field to given value.

### HasStalenessMinutes

`func (o *FxReferencePoint) HasStalenessMinutes() bool`

HasStalenessMinutes returns a boolean if a field has been set.

### SetStalenessMinutesNil

`func (o *FxReferencePoint) SetStalenessMinutesNil(b bool)`

 SetStalenessMinutesNil sets the value for StalenessMinutes to be an explicit nil

### UnsetStalenessMinutes
`func (o *FxReferencePoint) UnsetStalenessMinutes()`

UnsetStalenessMinutes ensures that no value is present for StalenessMinutes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


