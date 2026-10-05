# ForexPairMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourcePair** | Pointer to **NullableString** |  | [optional] 
**IsDerived** | Pointer to **NullableBool** |  | [optional] 
**DirectAvailable** | Pointer to **NullableBool** |  | [optional] 
**InverseAvailable** | Pointer to **NullableBool** |  | [optional] 
**OriginalPairAvailable** | Pointer to **NullableBool** |  | [optional] 
**OriginalInversePairAvailable** | Pointer to **NullableBool** |  | [optional] 
**DerivedFromInverse** | Pointer to **NullableBool** |  | [optional] 
**DerivedFromCross** | Pointer to **NullableBool** |  | [optional] 
**DerivedFromSourceLegs** | Pointer to **NullableBool** |  | [optional] 
**DerivationMethod** | Pointer to **NullableString** |  | [optional] 
**AnchorCurrency** | Pointer to **NullableString** |  | [optional] 
**DerivedValue** | Pointer to **NullableFloat64** |  | [optional] 
**PairConsistencyCheck** | Pointer to **NullableString** |  | [optional] 
**SourceNote** | Pointer to **NullableString** |  | [optional] 
**SourceLegs** | Pointer to [**[]FxSourceLeg**](FxSourceLeg.md) |  | [optional] 

## Methods

### NewForexPairMetadata

`func NewForexPairMetadata() *ForexPairMetadata`

NewForexPairMetadata instantiates a new ForexPairMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForexPairMetadataWithDefaults

`func NewForexPairMetadataWithDefaults() *ForexPairMetadata`

NewForexPairMetadataWithDefaults instantiates a new ForexPairMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourcePair

`func (o *ForexPairMetadata) GetSourcePair() string`

GetSourcePair returns the SourcePair field if non-nil, zero value otherwise.

### GetSourcePairOk

`func (o *ForexPairMetadata) GetSourcePairOk() (*string, bool)`

GetSourcePairOk returns a tuple with the SourcePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePair

`func (o *ForexPairMetadata) SetSourcePair(v string)`

SetSourcePair sets SourcePair field to given value.

### HasSourcePair

`func (o *ForexPairMetadata) HasSourcePair() bool`

HasSourcePair returns a boolean if a field has been set.

### SetSourcePairNil

`func (o *ForexPairMetadata) SetSourcePairNil(b bool)`

 SetSourcePairNil sets the value for SourcePair to be an explicit nil

### UnsetSourcePair
`func (o *ForexPairMetadata) UnsetSourcePair()`

UnsetSourcePair ensures that no value is present for SourcePair, not even an explicit nil
### GetIsDerived

`func (o *ForexPairMetadata) GetIsDerived() bool`

GetIsDerived returns the IsDerived field if non-nil, zero value otherwise.

### GetIsDerivedOk

`func (o *ForexPairMetadata) GetIsDerivedOk() (*bool, bool)`

GetIsDerivedOk returns a tuple with the IsDerived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDerived

`func (o *ForexPairMetadata) SetIsDerived(v bool)`

SetIsDerived sets IsDerived field to given value.

### HasIsDerived

`func (o *ForexPairMetadata) HasIsDerived() bool`

HasIsDerived returns a boolean if a field has been set.

### SetIsDerivedNil

`func (o *ForexPairMetadata) SetIsDerivedNil(b bool)`

 SetIsDerivedNil sets the value for IsDerived to be an explicit nil

### UnsetIsDerived
`func (o *ForexPairMetadata) UnsetIsDerived()`

UnsetIsDerived ensures that no value is present for IsDerived, not even an explicit nil
### GetDirectAvailable

`func (o *ForexPairMetadata) GetDirectAvailable() bool`

GetDirectAvailable returns the DirectAvailable field if non-nil, zero value otherwise.

### GetDirectAvailableOk

`func (o *ForexPairMetadata) GetDirectAvailableOk() (*bool, bool)`

GetDirectAvailableOk returns a tuple with the DirectAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectAvailable

`func (o *ForexPairMetadata) SetDirectAvailable(v bool)`

SetDirectAvailable sets DirectAvailable field to given value.

### HasDirectAvailable

`func (o *ForexPairMetadata) HasDirectAvailable() bool`

HasDirectAvailable returns a boolean if a field has been set.

### SetDirectAvailableNil

`func (o *ForexPairMetadata) SetDirectAvailableNil(b bool)`

 SetDirectAvailableNil sets the value for DirectAvailable to be an explicit nil

### UnsetDirectAvailable
`func (o *ForexPairMetadata) UnsetDirectAvailable()`

UnsetDirectAvailable ensures that no value is present for DirectAvailable, not even an explicit nil
### GetInverseAvailable

`func (o *ForexPairMetadata) GetInverseAvailable() bool`

GetInverseAvailable returns the InverseAvailable field if non-nil, zero value otherwise.

### GetInverseAvailableOk

`func (o *ForexPairMetadata) GetInverseAvailableOk() (*bool, bool)`

GetInverseAvailableOk returns a tuple with the InverseAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseAvailable

`func (o *ForexPairMetadata) SetInverseAvailable(v bool)`

SetInverseAvailable sets InverseAvailable field to given value.

### HasInverseAvailable

`func (o *ForexPairMetadata) HasInverseAvailable() bool`

HasInverseAvailable returns a boolean if a field has been set.

### SetInverseAvailableNil

`func (o *ForexPairMetadata) SetInverseAvailableNil(b bool)`

 SetInverseAvailableNil sets the value for InverseAvailable to be an explicit nil

### UnsetInverseAvailable
`func (o *ForexPairMetadata) UnsetInverseAvailable()`

UnsetInverseAvailable ensures that no value is present for InverseAvailable, not even an explicit nil
### GetOriginalPairAvailable

`func (o *ForexPairMetadata) GetOriginalPairAvailable() bool`

GetOriginalPairAvailable returns the OriginalPairAvailable field if non-nil, zero value otherwise.

### GetOriginalPairAvailableOk

`func (o *ForexPairMetadata) GetOriginalPairAvailableOk() (*bool, bool)`

GetOriginalPairAvailableOk returns a tuple with the OriginalPairAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPairAvailable

`func (o *ForexPairMetadata) SetOriginalPairAvailable(v bool)`

SetOriginalPairAvailable sets OriginalPairAvailable field to given value.

### HasOriginalPairAvailable

`func (o *ForexPairMetadata) HasOriginalPairAvailable() bool`

HasOriginalPairAvailable returns a boolean if a field has been set.

### SetOriginalPairAvailableNil

`func (o *ForexPairMetadata) SetOriginalPairAvailableNil(b bool)`

 SetOriginalPairAvailableNil sets the value for OriginalPairAvailable to be an explicit nil

### UnsetOriginalPairAvailable
`func (o *ForexPairMetadata) UnsetOriginalPairAvailable()`

UnsetOriginalPairAvailable ensures that no value is present for OriginalPairAvailable, not even an explicit nil
### GetOriginalInversePairAvailable

`func (o *ForexPairMetadata) GetOriginalInversePairAvailable() bool`

GetOriginalInversePairAvailable returns the OriginalInversePairAvailable field if non-nil, zero value otherwise.

### GetOriginalInversePairAvailableOk

`func (o *ForexPairMetadata) GetOriginalInversePairAvailableOk() (*bool, bool)`

GetOriginalInversePairAvailableOk returns a tuple with the OriginalInversePairAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalInversePairAvailable

`func (o *ForexPairMetadata) SetOriginalInversePairAvailable(v bool)`

SetOriginalInversePairAvailable sets OriginalInversePairAvailable field to given value.

### HasOriginalInversePairAvailable

`func (o *ForexPairMetadata) HasOriginalInversePairAvailable() bool`

HasOriginalInversePairAvailable returns a boolean if a field has been set.

### SetOriginalInversePairAvailableNil

`func (o *ForexPairMetadata) SetOriginalInversePairAvailableNil(b bool)`

 SetOriginalInversePairAvailableNil sets the value for OriginalInversePairAvailable to be an explicit nil

### UnsetOriginalInversePairAvailable
`func (o *ForexPairMetadata) UnsetOriginalInversePairAvailable()`

UnsetOriginalInversePairAvailable ensures that no value is present for OriginalInversePairAvailable, not even an explicit nil
### GetDerivedFromInverse

`func (o *ForexPairMetadata) GetDerivedFromInverse() bool`

GetDerivedFromInverse returns the DerivedFromInverse field if non-nil, zero value otherwise.

### GetDerivedFromInverseOk

`func (o *ForexPairMetadata) GetDerivedFromInverseOk() (*bool, bool)`

GetDerivedFromInverseOk returns a tuple with the DerivedFromInverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedFromInverse

`func (o *ForexPairMetadata) SetDerivedFromInverse(v bool)`

SetDerivedFromInverse sets DerivedFromInverse field to given value.

### HasDerivedFromInverse

`func (o *ForexPairMetadata) HasDerivedFromInverse() bool`

HasDerivedFromInverse returns a boolean if a field has been set.

### SetDerivedFromInverseNil

`func (o *ForexPairMetadata) SetDerivedFromInverseNil(b bool)`

 SetDerivedFromInverseNil sets the value for DerivedFromInverse to be an explicit nil

### UnsetDerivedFromInverse
`func (o *ForexPairMetadata) UnsetDerivedFromInverse()`

UnsetDerivedFromInverse ensures that no value is present for DerivedFromInverse, not even an explicit nil
### GetDerivedFromCross

`func (o *ForexPairMetadata) GetDerivedFromCross() bool`

GetDerivedFromCross returns the DerivedFromCross field if non-nil, zero value otherwise.

### GetDerivedFromCrossOk

`func (o *ForexPairMetadata) GetDerivedFromCrossOk() (*bool, bool)`

GetDerivedFromCrossOk returns a tuple with the DerivedFromCross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedFromCross

`func (o *ForexPairMetadata) SetDerivedFromCross(v bool)`

SetDerivedFromCross sets DerivedFromCross field to given value.

### HasDerivedFromCross

`func (o *ForexPairMetadata) HasDerivedFromCross() bool`

HasDerivedFromCross returns a boolean if a field has been set.

### SetDerivedFromCrossNil

`func (o *ForexPairMetadata) SetDerivedFromCrossNil(b bool)`

 SetDerivedFromCrossNil sets the value for DerivedFromCross to be an explicit nil

### UnsetDerivedFromCross
`func (o *ForexPairMetadata) UnsetDerivedFromCross()`

UnsetDerivedFromCross ensures that no value is present for DerivedFromCross, not even an explicit nil
### GetDerivedFromSourceLegs

`func (o *ForexPairMetadata) GetDerivedFromSourceLegs() bool`

GetDerivedFromSourceLegs returns the DerivedFromSourceLegs field if non-nil, zero value otherwise.

### GetDerivedFromSourceLegsOk

`func (o *ForexPairMetadata) GetDerivedFromSourceLegsOk() (*bool, bool)`

GetDerivedFromSourceLegsOk returns a tuple with the DerivedFromSourceLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedFromSourceLegs

`func (o *ForexPairMetadata) SetDerivedFromSourceLegs(v bool)`

SetDerivedFromSourceLegs sets DerivedFromSourceLegs field to given value.

### HasDerivedFromSourceLegs

`func (o *ForexPairMetadata) HasDerivedFromSourceLegs() bool`

HasDerivedFromSourceLegs returns a boolean if a field has been set.

### SetDerivedFromSourceLegsNil

`func (o *ForexPairMetadata) SetDerivedFromSourceLegsNil(b bool)`

 SetDerivedFromSourceLegsNil sets the value for DerivedFromSourceLegs to be an explicit nil

### UnsetDerivedFromSourceLegs
`func (o *ForexPairMetadata) UnsetDerivedFromSourceLegs()`

UnsetDerivedFromSourceLegs ensures that no value is present for DerivedFromSourceLegs, not even an explicit nil
### GetDerivationMethod

`func (o *ForexPairMetadata) GetDerivationMethod() string`

GetDerivationMethod returns the DerivationMethod field if non-nil, zero value otherwise.

### GetDerivationMethodOk

`func (o *ForexPairMetadata) GetDerivationMethodOk() (*string, bool)`

GetDerivationMethodOk returns a tuple with the DerivationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivationMethod

`func (o *ForexPairMetadata) SetDerivationMethod(v string)`

SetDerivationMethod sets DerivationMethod field to given value.

### HasDerivationMethod

`func (o *ForexPairMetadata) HasDerivationMethod() bool`

HasDerivationMethod returns a boolean if a field has been set.

### SetDerivationMethodNil

`func (o *ForexPairMetadata) SetDerivationMethodNil(b bool)`

 SetDerivationMethodNil sets the value for DerivationMethod to be an explicit nil

### UnsetDerivationMethod
`func (o *ForexPairMetadata) UnsetDerivationMethod()`

UnsetDerivationMethod ensures that no value is present for DerivationMethod, not even an explicit nil
### GetAnchorCurrency

`func (o *ForexPairMetadata) GetAnchorCurrency() string`

GetAnchorCurrency returns the AnchorCurrency field if non-nil, zero value otherwise.

### GetAnchorCurrencyOk

`func (o *ForexPairMetadata) GetAnchorCurrencyOk() (*string, bool)`

GetAnchorCurrencyOk returns a tuple with the AnchorCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCurrency

`func (o *ForexPairMetadata) SetAnchorCurrency(v string)`

SetAnchorCurrency sets AnchorCurrency field to given value.

### HasAnchorCurrency

`func (o *ForexPairMetadata) HasAnchorCurrency() bool`

HasAnchorCurrency returns a boolean if a field has been set.

### SetAnchorCurrencyNil

`func (o *ForexPairMetadata) SetAnchorCurrencyNil(b bool)`

 SetAnchorCurrencyNil sets the value for AnchorCurrency to be an explicit nil

### UnsetAnchorCurrency
`func (o *ForexPairMetadata) UnsetAnchorCurrency()`

UnsetAnchorCurrency ensures that no value is present for AnchorCurrency, not even an explicit nil
### GetDerivedValue

`func (o *ForexPairMetadata) GetDerivedValue() float64`

GetDerivedValue returns the DerivedValue field if non-nil, zero value otherwise.

### GetDerivedValueOk

`func (o *ForexPairMetadata) GetDerivedValueOk() (*float64, bool)`

GetDerivedValueOk returns a tuple with the DerivedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedValue

`func (o *ForexPairMetadata) SetDerivedValue(v float64)`

SetDerivedValue sets DerivedValue field to given value.

### HasDerivedValue

`func (o *ForexPairMetadata) HasDerivedValue() bool`

HasDerivedValue returns a boolean if a field has been set.

### SetDerivedValueNil

`func (o *ForexPairMetadata) SetDerivedValueNil(b bool)`

 SetDerivedValueNil sets the value for DerivedValue to be an explicit nil

### UnsetDerivedValue
`func (o *ForexPairMetadata) UnsetDerivedValue()`

UnsetDerivedValue ensures that no value is present for DerivedValue, not even an explicit nil
### GetPairConsistencyCheck

`func (o *ForexPairMetadata) GetPairConsistencyCheck() string`

GetPairConsistencyCheck returns the PairConsistencyCheck field if non-nil, zero value otherwise.

### GetPairConsistencyCheckOk

`func (o *ForexPairMetadata) GetPairConsistencyCheckOk() (*string, bool)`

GetPairConsistencyCheckOk returns a tuple with the PairConsistencyCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairConsistencyCheck

`func (o *ForexPairMetadata) SetPairConsistencyCheck(v string)`

SetPairConsistencyCheck sets PairConsistencyCheck field to given value.

### HasPairConsistencyCheck

`func (o *ForexPairMetadata) HasPairConsistencyCheck() bool`

HasPairConsistencyCheck returns a boolean if a field has been set.

### SetPairConsistencyCheckNil

`func (o *ForexPairMetadata) SetPairConsistencyCheckNil(b bool)`

 SetPairConsistencyCheckNil sets the value for PairConsistencyCheck to be an explicit nil

### UnsetPairConsistencyCheck
`func (o *ForexPairMetadata) UnsetPairConsistencyCheck()`

UnsetPairConsistencyCheck ensures that no value is present for PairConsistencyCheck, not even an explicit nil
### GetSourceNote

`func (o *ForexPairMetadata) GetSourceNote() string`

GetSourceNote returns the SourceNote field if non-nil, zero value otherwise.

### GetSourceNoteOk

`func (o *ForexPairMetadata) GetSourceNoteOk() (*string, bool)`

GetSourceNoteOk returns a tuple with the SourceNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNote

`func (o *ForexPairMetadata) SetSourceNote(v string)`

SetSourceNote sets SourceNote field to given value.

### HasSourceNote

`func (o *ForexPairMetadata) HasSourceNote() bool`

HasSourceNote returns a boolean if a field has been set.

### SetSourceNoteNil

`func (o *ForexPairMetadata) SetSourceNoteNil(b bool)`

 SetSourceNoteNil sets the value for SourceNote to be an explicit nil

### UnsetSourceNote
`func (o *ForexPairMetadata) UnsetSourceNote()`

UnsetSourceNote ensures that no value is present for SourceNote, not even an explicit nil
### GetSourceLegs

`func (o *ForexPairMetadata) GetSourceLegs() []FxSourceLeg`

GetSourceLegs returns the SourceLegs field if non-nil, zero value otherwise.

### GetSourceLegsOk

`func (o *ForexPairMetadata) GetSourceLegsOk() (*[]FxSourceLeg, bool)`

GetSourceLegsOk returns a tuple with the SourceLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLegs

`func (o *ForexPairMetadata) SetSourceLegs(v []FxSourceLeg)`

SetSourceLegs sets SourceLegs field to given value.

### HasSourceLegs

`func (o *ForexPairMetadata) HasSourceLegs() bool`

HasSourceLegs returns a boolean if a field has been set.

### SetSourceLegsNil

`func (o *ForexPairMetadata) SetSourceLegsNil(b bool)`

 SetSourceLegsNil sets the value for SourceLegs to be an explicit nil

### UnsetSourceLegs
`func (o *ForexPairMetadata) UnsetSourceLegs()`

UnsetSourceLegs ensures that no value is present for SourceLegs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


