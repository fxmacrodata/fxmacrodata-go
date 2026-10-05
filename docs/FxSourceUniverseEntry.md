# FxSourceUniverseEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**NullableFxSource**](FxSource.md) |  | [optional] 
**PairCount** | Pointer to **NullableInt64** |  | [optional] 
**Pairs** | Pointer to [**[]FxSourceUniversePair**](FxSourceUniversePair.md) |  | [optional] 

## Methods

### NewFxSourceUniverseEntry

`func NewFxSourceUniverseEntry() *FxSourceUniverseEntry`

NewFxSourceUniverseEntry instantiates a new FxSourceUniverseEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxSourceUniverseEntryWithDefaults

`func NewFxSourceUniverseEntryWithDefaults() *FxSourceUniverseEntry`

NewFxSourceUniverseEntryWithDefaults instantiates a new FxSourceUniverseEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *FxSourceUniverseEntry) GetSource() FxSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FxSourceUniverseEntry) GetSourceOk() (*FxSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FxSourceUniverseEntry) SetSource(v FxSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *FxSourceUniverseEntry) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *FxSourceUniverseEntry) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *FxSourceUniverseEntry) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetPairCount

`func (o *FxSourceUniverseEntry) GetPairCount() int64`

GetPairCount returns the PairCount field if non-nil, zero value otherwise.

### GetPairCountOk

`func (o *FxSourceUniverseEntry) GetPairCountOk() (*int64, bool)`

GetPairCountOk returns a tuple with the PairCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairCount

`func (o *FxSourceUniverseEntry) SetPairCount(v int64)`

SetPairCount sets PairCount field to given value.

### HasPairCount

`func (o *FxSourceUniverseEntry) HasPairCount() bool`

HasPairCount returns a boolean if a field has been set.

### SetPairCountNil

`func (o *FxSourceUniverseEntry) SetPairCountNil(b bool)`

 SetPairCountNil sets the value for PairCount to be an explicit nil

### UnsetPairCount
`func (o *FxSourceUniverseEntry) UnsetPairCount()`

UnsetPairCount ensures that no value is present for PairCount, not even an explicit nil
### GetPairs

`func (o *FxSourceUniverseEntry) GetPairs() []FxSourceUniversePair`

GetPairs returns the Pairs field if non-nil, zero value otherwise.

### GetPairsOk

`func (o *FxSourceUniverseEntry) GetPairsOk() (*[]FxSourceUniversePair, bool)`

GetPairsOk returns a tuple with the Pairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairs

`func (o *FxSourceUniverseEntry) SetPairs(v []FxSourceUniversePair)`

SetPairs sets Pairs field to given value.

### HasPairs

`func (o *FxSourceUniverseEntry) HasPairs() bool`

HasPairs returns a boolean if a field has been set.

### SetPairsNil

`func (o *FxSourceUniverseEntry) SetPairsNil(b bool)`

 SetPairsNil sets the value for Pairs to be an explicit nil

### UnsetPairs
`func (o *FxSourceUniverseEntry) UnsetPairs()`

UnsetPairs ensures that no value is present for Pairs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


