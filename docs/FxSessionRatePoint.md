# FxSessionRatePoint

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
**Session** | Pointer to **NullableString** |  | [optional] 
**SessionCloseUtc** | Pointer to **NullableString** |  | [optional] 
**SelectedTimestampUtc** | Pointer to **NullableString** |  | [optional] 
**SelectedSourceTimestampUtc** | Pointer to **NullableString** |  | [optional] 
**SelectionRule** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFxSessionRatePoint

`func NewFxSessionRatePoint(date string, referenceDate string, base string, quote string, pair string, rate float64, ) *FxSessionRatePoint`

NewFxSessionRatePoint instantiates a new FxSessionRatePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxSessionRatePointWithDefaults

`func NewFxSessionRatePointWithDefaults() *FxSessionRatePoint`

NewFxSessionRatePointWithDefaults instantiates a new FxSessionRatePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *FxSessionRatePoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FxSessionRatePoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FxSessionRatePoint) SetDate(v string)`

SetDate sets Date field to given value.


### GetReferenceDate

`func (o *FxSessionRatePoint) GetReferenceDate() string`

GetReferenceDate returns the ReferenceDate field if non-nil, zero value otherwise.

### GetReferenceDateOk

`func (o *FxSessionRatePoint) GetReferenceDateOk() (*string, bool)`

GetReferenceDateOk returns a tuple with the ReferenceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceDate

`func (o *FxSessionRatePoint) SetReferenceDate(v string)`

SetReferenceDate sets ReferenceDate field to given value.


### GetBase

`func (o *FxSessionRatePoint) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *FxSessionRatePoint) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *FxSessionRatePoint) SetBase(v string)`

SetBase sets Base field to given value.


### GetQuote

`func (o *FxSessionRatePoint) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *FxSessionRatePoint) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *FxSessionRatePoint) SetQuote(v string)`

SetQuote sets Quote field to given value.


### GetPair

`func (o *FxSessionRatePoint) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *FxSessionRatePoint) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *FxSessionRatePoint) SetPair(v string)`

SetPair sets Pair field to given value.


### GetRate

`func (o *FxSessionRatePoint) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *FxSessionRatePoint) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *FxSessionRatePoint) SetRate(v float64)`

SetRate sets Rate field to given value.


### GetTimestampUtc

`func (o *FxSessionRatePoint) GetTimestampUtc() string`

GetTimestampUtc returns the TimestampUtc field if non-nil, zero value otherwise.

### GetTimestampUtcOk

`func (o *FxSessionRatePoint) GetTimestampUtcOk() (*string, bool)`

GetTimestampUtcOk returns a tuple with the TimestampUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampUtc

`func (o *FxSessionRatePoint) SetTimestampUtc(v string)`

SetTimestampUtc sets TimestampUtc field to given value.

### HasTimestampUtc

`func (o *FxSessionRatePoint) HasTimestampUtc() bool`

HasTimestampUtc returns a boolean if a field has been set.

### SetTimestampUtcNil

`func (o *FxSessionRatePoint) SetTimestampUtcNil(b bool)`

 SetTimestampUtcNil sets the value for TimestampUtc to be an explicit nil

### UnsetTimestampUtc
`func (o *FxSessionRatePoint) UnsetTimestampUtc()`

UnsetTimestampUtc ensures that no value is present for TimestampUtc, not even an explicit nil
### GetTimestampType

`func (o *FxSessionRatePoint) GetTimestampType() string`

GetTimestampType returns the TimestampType field if non-nil, zero value otherwise.

### GetTimestampTypeOk

`func (o *FxSessionRatePoint) GetTimestampTypeOk() (*string, bool)`

GetTimestampTypeOk returns a tuple with the TimestampType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampType

`func (o *FxSessionRatePoint) SetTimestampType(v string)`

SetTimestampType sets TimestampType field to given value.

### HasTimestampType

`func (o *FxSessionRatePoint) HasTimestampType() bool`

HasTimestampType returns a boolean if a field has been set.

### SetTimestampTypeNil

`func (o *FxSessionRatePoint) SetTimestampTypeNil(b bool)`

 SetTimestampTypeNil sets the value for TimestampType to be an explicit nil

### UnsetTimestampType
`func (o *FxSessionRatePoint) UnsetTimestampType()`

UnsetTimestampType ensures that no value is present for TimestampType, not even an explicit nil
### GetFixingTimeUtc

`func (o *FxSessionRatePoint) GetFixingTimeUtc() string`

GetFixingTimeUtc returns the FixingTimeUtc field if non-nil, zero value otherwise.

### GetFixingTimeUtcOk

`func (o *FxSessionRatePoint) GetFixingTimeUtcOk() (*string, bool)`

GetFixingTimeUtcOk returns a tuple with the FixingTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixingTimeUtc

`func (o *FxSessionRatePoint) SetFixingTimeUtc(v string)`

SetFixingTimeUtc sets FixingTimeUtc field to given value.

### HasFixingTimeUtc

`func (o *FxSessionRatePoint) HasFixingTimeUtc() bool`

HasFixingTimeUtc returns a boolean if a field has been set.

### SetFixingTimeUtcNil

`func (o *FxSessionRatePoint) SetFixingTimeUtcNil(b bool)`

 SetFixingTimeUtcNil sets the value for FixingTimeUtc to be an explicit nil

### UnsetFixingTimeUtc
`func (o *FxSessionRatePoint) UnsetFixingTimeUtc()`

UnsetFixingTimeUtc ensures that no value is present for FixingTimeUtc, not even an explicit nil
### GetPublishedAtUtc

`func (o *FxSessionRatePoint) GetPublishedAtUtc() string`

GetPublishedAtUtc returns the PublishedAtUtc field if non-nil, zero value otherwise.

### GetPublishedAtUtcOk

`func (o *FxSessionRatePoint) GetPublishedAtUtcOk() (*string, bool)`

GetPublishedAtUtcOk returns a tuple with the PublishedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAtUtc

`func (o *FxSessionRatePoint) SetPublishedAtUtc(v string)`

SetPublishedAtUtc sets PublishedAtUtc field to given value.

### HasPublishedAtUtc

`func (o *FxSessionRatePoint) HasPublishedAtUtc() bool`

HasPublishedAtUtc returns a boolean if a field has been set.

### SetPublishedAtUtcNil

`func (o *FxSessionRatePoint) SetPublishedAtUtcNil(b bool)`

 SetPublishedAtUtcNil sets the value for PublishedAtUtc to be an explicit nil

### UnsetPublishedAtUtc
`func (o *FxSessionRatePoint) UnsetPublishedAtUtc()`

UnsetPublishedAtUtc ensures that no value is present for PublishedAtUtc, not even an explicit nil
### GetTimePrecision

`func (o *FxSessionRatePoint) GetTimePrecision() string`

GetTimePrecision returns the TimePrecision field if non-nil, zero value otherwise.

### GetTimePrecisionOk

`func (o *FxSessionRatePoint) GetTimePrecisionOk() (*string, bool)`

GetTimePrecisionOk returns a tuple with the TimePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePrecision

`func (o *FxSessionRatePoint) SetTimePrecision(v string)`

SetTimePrecision sets TimePrecision field to given value.

### HasTimePrecision

`func (o *FxSessionRatePoint) HasTimePrecision() bool`

HasTimePrecision returns a boolean if a field has been set.

### SetTimePrecisionNil

`func (o *FxSessionRatePoint) SetTimePrecisionNil(b bool)`

 SetTimePrecisionNil sets the value for TimePrecision to be an explicit nil

### UnsetTimePrecision
`func (o *FxSessionRatePoint) UnsetTimePrecision()`

UnsetTimePrecision ensures that no value is present for TimePrecision, not even an explicit nil
### GetSource

`func (o *FxSessionRatePoint) GetSource() ForexPointSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FxSessionRatePoint) GetSourceOk() (*ForexPointSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FxSessionRatePoint) SetSource(v ForexPointSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *FxSessionRatePoint) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetNativeOrDerived

`func (o *FxSessionRatePoint) GetNativeOrDerived() string`

GetNativeOrDerived returns the NativeOrDerived field if non-nil, zero value otherwise.

### GetNativeOrDerivedOk

`func (o *FxSessionRatePoint) GetNativeOrDerivedOk() (*string, bool)`

GetNativeOrDerivedOk returns a tuple with the NativeOrDerived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeOrDerived

`func (o *FxSessionRatePoint) SetNativeOrDerived(v string)`

SetNativeOrDerived sets NativeOrDerived field to given value.

### HasNativeOrDerived

`func (o *FxSessionRatePoint) HasNativeOrDerived() bool`

HasNativeOrDerived returns a boolean if a field has been set.

### SetNativeOrDerivedNil

`func (o *FxSessionRatePoint) SetNativeOrDerivedNil(b bool)`

 SetNativeOrDerivedNil sets the value for NativeOrDerived to be an explicit nil

### UnsetNativeOrDerived
`func (o *FxSessionRatePoint) UnsetNativeOrDerived()`

UnsetNativeOrDerived ensures that no value is present for NativeOrDerived, not even an explicit nil
### GetDerivationMethod

`func (o *FxSessionRatePoint) GetDerivationMethod() string`

GetDerivationMethod returns the DerivationMethod field if non-nil, zero value otherwise.

### GetDerivationMethodOk

`func (o *FxSessionRatePoint) GetDerivationMethodOk() (*string, bool)`

GetDerivationMethodOk returns a tuple with the DerivationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivationMethod

`func (o *FxSessionRatePoint) SetDerivationMethod(v string)`

SetDerivationMethod sets DerivationMethod field to given value.

### HasDerivationMethod

`func (o *FxSessionRatePoint) HasDerivationMethod() bool`

HasDerivationMethod returns a boolean if a field has been set.

### SetDerivationMethodNil

`func (o *FxSessionRatePoint) SetDerivationMethodNil(b bool)`

 SetDerivationMethodNil sets the value for DerivationMethod to be an explicit nil

### UnsetDerivationMethod
`func (o *FxSessionRatePoint) UnsetDerivationMethod()`

UnsetDerivationMethod ensures that no value is present for DerivationMethod, not even an explicit nil
### GetSourcePair

`func (o *FxSessionRatePoint) GetSourcePair() string`

GetSourcePair returns the SourcePair field if non-nil, zero value otherwise.

### GetSourcePairOk

`func (o *FxSessionRatePoint) GetSourcePairOk() (*string, bool)`

GetSourcePairOk returns a tuple with the SourcePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePair

`func (o *FxSessionRatePoint) SetSourcePair(v string)`

SetSourcePair sets SourcePair field to given value.

### HasSourcePair

`func (o *FxSessionRatePoint) HasSourcePair() bool`

HasSourcePair returns a boolean if a field has been set.

### SetSourcePairNil

`func (o *FxSessionRatePoint) SetSourcePairNil(b bool)`

 SetSourcePairNil sets the value for SourcePair to be an explicit nil

### UnsetSourcePair
`func (o *FxSessionRatePoint) UnsetSourcePair()`

UnsetSourcePair ensures that no value is present for SourcePair, not even an explicit nil
### GetSourceLegs

`func (o *FxSessionRatePoint) GetSourceLegs() []FxSourceLeg`

GetSourceLegs returns the SourceLegs field if non-nil, zero value otherwise.

### GetSourceLegsOk

`func (o *FxSessionRatePoint) GetSourceLegsOk() (*[]FxSourceLeg, bool)`

GetSourceLegsOk returns a tuple with the SourceLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLegs

`func (o *FxSessionRatePoint) SetSourceLegs(v []FxSourceLeg)`

SetSourceLegs sets SourceLegs field to given value.

### HasSourceLegs

`func (o *FxSessionRatePoint) HasSourceLegs() bool`

HasSourceLegs returns a boolean if a field has been set.

### GetStalenessMinutes

`func (o *FxSessionRatePoint) GetStalenessMinutes() int64`

GetStalenessMinutes returns the StalenessMinutes field if non-nil, zero value otherwise.

### GetStalenessMinutesOk

`func (o *FxSessionRatePoint) GetStalenessMinutesOk() (*int64, bool)`

GetStalenessMinutesOk returns a tuple with the StalenessMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStalenessMinutes

`func (o *FxSessionRatePoint) SetStalenessMinutes(v int64)`

SetStalenessMinutes sets StalenessMinutes field to given value.

### HasStalenessMinutes

`func (o *FxSessionRatePoint) HasStalenessMinutes() bool`

HasStalenessMinutes returns a boolean if a field has been set.

### SetStalenessMinutesNil

`func (o *FxSessionRatePoint) SetStalenessMinutesNil(b bool)`

 SetStalenessMinutesNil sets the value for StalenessMinutes to be an explicit nil

### UnsetStalenessMinutes
`func (o *FxSessionRatePoint) UnsetStalenessMinutes()`

UnsetStalenessMinutes ensures that no value is present for StalenessMinutes, not even an explicit nil
### GetSession

`func (o *FxSessionRatePoint) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *FxSessionRatePoint) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *FxSessionRatePoint) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *FxSessionRatePoint) HasSession() bool`

HasSession returns a boolean if a field has been set.

### SetSessionNil

`func (o *FxSessionRatePoint) SetSessionNil(b bool)`

 SetSessionNil sets the value for Session to be an explicit nil

### UnsetSession
`func (o *FxSessionRatePoint) UnsetSession()`

UnsetSession ensures that no value is present for Session, not even an explicit nil
### GetSessionCloseUtc

`func (o *FxSessionRatePoint) GetSessionCloseUtc() string`

GetSessionCloseUtc returns the SessionCloseUtc field if non-nil, zero value otherwise.

### GetSessionCloseUtcOk

`func (o *FxSessionRatePoint) GetSessionCloseUtcOk() (*string, bool)`

GetSessionCloseUtcOk returns a tuple with the SessionCloseUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCloseUtc

`func (o *FxSessionRatePoint) SetSessionCloseUtc(v string)`

SetSessionCloseUtc sets SessionCloseUtc field to given value.

### HasSessionCloseUtc

`func (o *FxSessionRatePoint) HasSessionCloseUtc() bool`

HasSessionCloseUtc returns a boolean if a field has been set.

### SetSessionCloseUtcNil

`func (o *FxSessionRatePoint) SetSessionCloseUtcNil(b bool)`

 SetSessionCloseUtcNil sets the value for SessionCloseUtc to be an explicit nil

### UnsetSessionCloseUtc
`func (o *FxSessionRatePoint) UnsetSessionCloseUtc()`

UnsetSessionCloseUtc ensures that no value is present for SessionCloseUtc, not even an explicit nil
### GetSelectedTimestampUtc

`func (o *FxSessionRatePoint) GetSelectedTimestampUtc() string`

GetSelectedTimestampUtc returns the SelectedTimestampUtc field if non-nil, zero value otherwise.

### GetSelectedTimestampUtcOk

`func (o *FxSessionRatePoint) GetSelectedTimestampUtcOk() (*string, bool)`

GetSelectedTimestampUtcOk returns a tuple with the SelectedTimestampUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedTimestampUtc

`func (o *FxSessionRatePoint) SetSelectedTimestampUtc(v string)`

SetSelectedTimestampUtc sets SelectedTimestampUtc field to given value.

### HasSelectedTimestampUtc

`func (o *FxSessionRatePoint) HasSelectedTimestampUtc() bool`

HasSelectedTimestampUtc returns a boolean if a field has been set.

### SetSelectedTimestampUtcNil

`func (o *FxSessionRatePoint) SetSelectedTimestampUtcNil(b bool)`

 SetSelectedTimestampUtcNil sets the value for SelectedTimestampUtc to be an explicit nil

### UnsetSelectedTimestampUtc
`func (o *FxSessionRatePoint) UnsetSelectedTimestampUtc()`

UnsetSelectedTimestampUtc ensures that no value is present for SelectedTimestampUtc, not even an explicit nil
### GetSelectedSourceTimestampUtc

`func (o *FxSessionRatePoint) GetSelectedSourceTimestampUtc() string`

GetSelectedSourceTimestampUtc returns the SelectedSourceTimestampUtc field if non-nil, zero value otherwise.

### GetSelectedSourceTimestampUtcOk

`func (o *FxSessionRatePoint) GetSelectedSourceTimestampUtcOk() (*string, bool)`

GetSelectedSourceTimestampUtcOk returns a tuple with the SelectedSourceTimestampUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSourceTimestampUtc

`func (o *FxSessionRatePoint) SetSelectedSourceTimestampUtc(v string)`

SetSelectedSourceTimestampUtc sets SelectedSourceTimestampUtc field to given value.

### HasSelectedSourceTimestampUtc

`func (o *FxSessionRatePoint) HasSelectedSourceTimestampUtc() bool`

HasSelectedSourceTimestampUtc returns a boolean if a field has been set.

### SetSelectedSourceTimestampUtcNil

`func (o *FxSessionRatePoint) SetSelectedSourceTimestampUtcNil(b bool)`

 SetSelectedSourceTimestampUtcNil sets the value for SelectedSourceTimestampUtc to be an explicit nil

### UnsetSelectedSourceTimestampUtc
`func (o *FxSessionRatePoint) UnsetSelectedSourceTimestampUtc()`

UnsetSelectedSourceTimestampUtc ensures that no value is present for SelectedSourceTimestampUtc, not even an explicit nil
### GetSelectionRule

`func (o *FxSessionRatePoint) GetSelectionRule() string`

GetSelectionRule returns the SelectionRule field if non-nil, zero value otherwise.

### GetSelectionRuleOk

`func (o *FxSessionRatePoint) GetSelectionRuleOk() (*string, bool)`

GetSelectionRuleOk returns a tuple with the SelectionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionRule

`func (o *FxSessionRatePoint) SetSelectionRule(v string)`

SetSelectionRule sets SelectionRule field to given value.

### HasSelectionRule

`func (o *FxSessionRatePoint) HasSelectionRule() bool`

HasSelectionRule returns a boolean if a field has been set.

### SetSelectionRuleNil

`func (o *FxSessionRatePoint) SetSelectionRuleNil(b bool)`

 SetSelectionRuleNil sets the value for SelectionRule to be an explicit nil

### UnsetSelectionRule
`func (o *FxSessionRatePoint) UnsetSelectionRule()`

UnsetSelectionRule ensures that no value is present for SelectionRule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


