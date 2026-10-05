# ForexPointSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourcePair** | Pointer to **NullableString** |  | [optional] 
**IsDerived** | Pointer to **NullableBool** |  | [optional] 
**DerivedFromSourceLegs** | Pointer to **NullableBool** |  | [optional] 
**DerivationMethod** | Pointer to **NullableString** |  | [optional] 
**PairConsistencyCheck** | Pointer to **NullableString** |  | [optional] 
**SourceLegs** | Pointer to [**[]FxSourceLeg**](FxSourceLeg.md) |  | [optional] 

## Methods

### NewForexPointSource

`func NewForexPointSource() *ForexPointSource`

NewForexPointSource instantiates a new ForexPointSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForexPointSourceWithDefaults

`func NewForexPointSourceWithDefaults() *ForexPointSource`

NewForexPointSourceWithDefaults instantiates a new ForexPointSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourcePair

`func (o *ForexPointSource) GetSourcePair() string`

GetSourcePair returns the SourcePair field if non-nil, zero value otherwise.

### GetSourcePairOk

`func (o *ForexPointSource) GetSourcePairOk() (*string, bool)`

GetSourcePairOk returns a tuple with the SourcePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePair

`func (o *ForexPointSource) SetSourcePair(v string)`

SetSourcePair sets SourcePair field to given value.

### HasSourcePair

`func (o *ForexPointSource) HasSourcePair() bool`

HasSourcePair returns a boolean if a field has been set.

### SetSourcePairNil

`func (o *ForexPointSource) SetSourcePairNil(b bool)`

 SetSourcePairNil sets the value for SourcePair to be an explicit nil

### UnsetSourcePair
`func (o *ForexPointSource) UnsetSourcePair()`

UnsetSourcePair ensures that no value is present for SourcePair, not even an explicit nil
### GetIsDerived

`func (o *ForexPointSource) GetIsDerived() bool`

GetIsDerived returns the IsDerived field if non-nil, zero value otherwise.

### GetIsDerivedOk

`func (o *ForexPointSource) GetIsDerivedOk() (*bool, bool)`

GetIsDerivedOk returns a tuple with the IsDerived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDerived

`func (o *ForexPointSource) SetIsDerived(v bool)`

SetIsDerived sets IsDerived field to given value.

### HasIsDerived

`func (o *ForexPointSource) HasIsDerived() bool`

HasIsDerived returns a boolean if a field has been set.

### SetIsDerivedNil

`func (o *ForexPointSource) SetIsDerivedNil(b bool)`

 SetIsDerivedNil sets the value for IsDerived to be an explicit nil

### UnsetIsDerived
`func (o *ForexPointSource) UnsetIsDerived()`

UnsetIsDerived ensures that no value is present for IsDerived, not even an explicit nil
### GetDerivedFromSourceLegs

`func (o *ForexPointSource) GetDerivedFromSourceLegs() bool`

GetDerivedFromSourceLegs returns the DerivedFromSourceLegs field if non-nil, zero value otherwise.

### GetDerivedFromSourceLegsOk

`func (o *ForexPointSource) GetDerivedFromSourceLegsOk() (*bool, bool)`

GetDerivedFromSourceLegsOk returns a tuple with the DerivedFromSourceLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedFromSourceLegs

`func (o *ForexPointSource) SetDerivedFromSourceLegs(v bool)`

SetDerivedFromSourceLegs sets DerivedFromSourceLegs field to given value.

### HasDerivedFromSourceLegs

`func (o *ForexPointSource) HasDerivedFromSourceLegs() bool`

HasDerivedFromSourceLegs returns a boolean if a field has been set.

### SetDerivedFromSourceLegsNil

`func (o *ForexPointSource) SetDerivedFromSourceLegsNil(b bool)`

 SetDerivedFromSourceLegsNil sets the value for DerivedFromSourceLegs to be an explicit nil

### UnsetDerivedFromSourceLegs
`func (o *ForexPointSource) UnsetDerivedFromSourceLegs()`

UnsetDerivedFromSourceLegs ensures that no value is present for DerivedFromSourceLegs, not even an explicit nil
### GetDerivationMethod

`func (o *ForexPointSource) GetDerivationMethod() string`

GetDerivationMethod returns the DerivationMethod field if non-nil, zero value otherwise.

### GetDerivationMethodOk

`func (o *ForexPointSource) GetDerivationMethodOk() (*string, bool)`

GetDerivationMethodOk returns a tuple with the DerivationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivationMethod

`func (o *ForexPointSource) SetDerivationMethod(v string)`

SetDerivationMethod sets DerivationMethod field to given value.

### HasDerivationMethod

`func (o *ForexPointSource) HasDerivationMethod() bool`

HasDerivationMethod returns a boolean if a field has been set.

### SetDerivationMethodNil

`func (o *ForexPointSource) SetDerivationMethodNil(b bool)`

 SetDerivationMethodNil sets the value for DerivationMethod to be an explicit nil

### UnsetDerivationMethod
`func (o *ForexPointSource) UnsetDerivationMethod()`

UnsetDerivationMethod ensures that no value is present for DerivationMethod, not even an explicit nil
### GetPairConsistencyCheck

`func (o *ForexPointSource) GetPairConsistencyCheck() string`

GetPairConsistencyCheck returns the PairConsistencyCheck field if non-nil, zero value otherwise.

### GetPairConsistencyCheckOk

`func (o *ForexPointSource) GetPairConsistencyCheckOk() (*string, bool)`

GetPairConsistencyCheckOk returns a tuple with the PairConsistencyCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairConsistencyCheck

`func (o *ForexPointSource) SetPairConsistencyCheck(v string)`

SetPairConsistencyCheck sets PairConsistencyCheck field to given value.

### HasPairConsistencyCheck

`func (o *ForexPointSource) HasPairConsistencyCheck() bool`

HasPairConsistencyCheck returns a boolean if a field has been set.

### SetPairConsistencyCheckNil

`func (o *ForexPointSource) SetPairConsistencyCheckNil(b bool)`

 SetPairConsistencyCheckNil sets the value for PairConsistencyCheck to be an explicit nil

### UnsetPairConsistencyCheck
`func (o *ForexPointSource) UnsetPairConsistencyCheck()`

UnsetPairConsistencyCheck ensures that no value is present for PairConsistencyCheck, not even an explicit nil
### GetSourceLegs

`func (o *ForexPointSource) GetSourceLegs() []FxSourceLeg`

GetSourceLegs returns the SourceLegs field if non-nil, zero value otherwise.

### GetSourceLegsOk

`func (o *ForexPointSource) GetSourceLegsOk() (*[]FxSourceLeg, bool)`

GetSourceLegsOk returns a tuple with the SourceLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLegs

`func (o *ForexPointSource) SetSourceLegs(v []FxSourceLeg)`

SetSourceLegs sets SourceLegs field to given value.

### HasSourceLegs

`func (o *ForexPointSource) HasSourceLegs() bool`

HasSourceLegs returns a boolean if a field has been set.

### SetSourceLegsNil

`func (o *ForexPointSource) SetSourceLegsNil(b bool)`

 SetSourceLegsNil sets the value for SourceLegs to be an explicit nil

### UnsetSourceLegs
`func (o *ForexPointSource) UnsetSourceLegs()`

UnsetSourceLegs ensures that no value is present for SourceLegs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


