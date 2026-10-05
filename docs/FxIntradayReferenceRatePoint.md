# FxIntradayReferenceRatePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** |  | 
**Price** | **float64** |  | 
**ReferenceDate** | Pointer to **NullableString** |  | [optional] 
**TimestampType** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to [**NullableFxSource**](FxSource.md) |  | [optional] 
**SourcePair** | Pointer to **NullableString** |  | [optional] 
**DerivationMethod** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFxIntradayReferenceRatePoint

`func NewFxIntradayReferenceRatePoint(timestamp string, price float64, ) *FxIntradayReferenceRatePoint`

NewFxIntradayReferenceRatePoint instantiates a new FxIntradayReferenceRatePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxIntradayReferenceRatePointWithDefaults

`func NewFxIntradayReferenceRatePointWithDefaults() *FxIntradayReferenceRatePoint`

NewFxIntradayReferenceRatePointWithDefaults instantiates a new FxIntradayReferenceRatePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *FxIntradayReferenceRatePoint) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FxIntradayReferenceRatePoint) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FxIntradayReferenceRatePoint) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetPrice

`func (o *FxIntradayReferenceRatePoint) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *FxIntradayReferenceRatePoint) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *FxIntradayReferenceRatePoint) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetReferenceDate

`func (o *FxIntradayReferenceRatePoint) GetReferenceDate() string`

GetReferenceDate returns the ReferenceDate field if non-nil, zero value otherwise.

### GetReferenceDateOk

`func (o *FxIntradayReferenceRatePoint) GetReferenceDateOk() (*string, bool)`

GetReferenceDateOk returns a tuple with the ReferenceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceDate

`func (o *FxIntradayReferenceRatePoint) SetReferenceDate(v string)`

SetReferenceDate sets ReferenceDate field to given value.

### HasReferenceDate

`func (o *FxIntradayReferenceRatePoint) HasReferenceDate() bool`

HasReferenceDate returns a boolean if a field has been set.

### SetReferenceDateNil

`func (o *FxIntradayReferenceRatePoint) SetReferenceDateNil(b bool)`

 SetReferenceDateNil sets the value for ReferenceDate to be an explicit nil

### UnsetReferenceDate
`func (o *FxIntradayReferenceRatePoint) UnsetReferenceDate()`

UnsetReferenceDate ensures that no value is present for ReferenceDate, not even an explicit nil
### GetTimestampType

`func (o *FxIntradayReferenceRatePoint) GetTimestampType() string`

GetTimestampType returns the TimestampType field if non-nil, zero value otherwise.

### GetTimestampTypeOk

`func (o *FxIntradayReferenceRatePoint) GetTimestampTypeOk() (*string, bool)`

GetTimestampTypeOk returns a tuple with the TimestampType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampType

`func (o *FxIntradayReferenceRatePoint) SetTimestampType(v string)`

SetTimestampType sets TimestampType field to given value.

### HasTimestampType

`func (o *FxIntradayReferenceRatePoint) HasTimestampType() bool`

HasTimestampType returns a boolean if a field has been set.

### SetTimestampTypeNil

`func (o *FxIntradayReferenceRatePoint) SetTimestampTypeNil(b bool)`

 SetTimestampTypeNil sets the value for TimestampType to be an explicit nil

### UnsetTimestampType
`func (o *FxIntradayReferenceRatePoint) UnsetTimestampType()`

UnsetTimestampType ensures that no value is present for TimestampType, not even an explicit nil
### GetSource

`func (o *FxIntradayReferenceRatePoint) GetSource() FxSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FxIntradayReferenceRatePoint) GetSourceOk() (*FxSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FxIntradayReferenceRatePoint) SetSource(v FxSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *FxIntradayReferenceRatePoint) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *FxIntradayReferenceRatePoint) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *FxIntradayReferenceRatePoint) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourcePair

`func (o *FxIntradayReferenceRatePoint) GetSourcePair() string`

GetSourcePair returns the SourcePair field if non-nil, zero value otherwise.

### GetSourcePairOk

`func (o *FxIntradayReferenceRatePoint) GetSourcePairOk() (*string, bool)`

GetSourcePairOk returns a tuple with the SourcePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePair

`func (o *FxIntradayReferenceRatePoint) SetSourcePair(v string)`

SetSourcePair sets SourcePair field to given value.

### HasSourcePair

`func (o *FxIntradayReferenceRatePoint) HasSourcePair() bool`

HasSourcePair returns a boolean if a field has been set.

### SetSourcePairNil

`func (o *FxIntradayReferenceRatePoint) SetSourcePairNil(b bool)`

 SetSourcePairNil sets the value for SourcePair to be an explicit nil

### UnsetSourcePair
`func (o *FxIntradayReferenceRatePoint) UnsetSourcePair()`

UnsetSourcePair ensures that no value is present for SourcePair, not even an explicit nil
### GetDerivationMethod

`func (o *FxIntradayReferenceRatePoint) GetDerivationMethod() string`

GetDerivationMethod returns the DerivationMethod field if non-nil, zero value otherwise.

### GetDerivationMethodOk

`func (o *FxIntradayReferenceRatePoint) GetDerivationMethodOk() (*string, bool)`

GetDerivationMethodOk returns a tuple with the DerivationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivationMethod

`func (o *FxIntradayReferenceRatePoint) SetDerivationMethod(v string)`

SetDerivationMethod sets DerivationMethod field to given value.

### HasDerivationMethod

`func (o *FxIntradayReferenceRatePoint) HasDerivationMethod() bool`

HasDerivationMethod returns a boolean if a field has been set.

### SetDerivationMethodNil

`func (o *FxIntradayReferenceRatePoint) SetDerivationMethodNil(b bool)`

 SetDerivationMethodNil sets the value for DerivationMethod to be an explicit nil

### UnsetDerivationMethod
`func (o *FxIntradayReferenceRatePoint) UnsetDerivationMethod()`

UnsetDerivationMethod ensures that no value is present for DerivationMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


