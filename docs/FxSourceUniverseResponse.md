# FxSourceUniverseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourcePolicy** | [**FxSourcePolicy**](FxSourcePolicy.md) |  | 
**Currency** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**Data** | [**[]FxSourceUniverseEntry**](FxSourceUniverseEntry.md) |  | 

## Methods

### NewFxSourceUniverseResponse

`func NewFxSourceUniverseResponse(sourcePolicy FxSourcePolicy, data []FxSourceUniverseEntry, ) *FxSourceUniverseResponse`

NewFxSourceUniverseResponse instantiates a new FxSourceUniverseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxSourceUniverseResponseWithDefaults

`func NewFxSourceUniverseResponseWithDefaults() *FxSourceUniverseResponse`

NewFxSourceUniverseResponseWithDefaults instantiates a new FxSourceUniverseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourcePolicy

`func (o *FxSourceUniverseResponse) GetSourcePolicy() FxSourcePolicy`

GetSourcePolicy returns the SourcePolicy field if non-nil, zero value otherwise.

### GetSourcePolicyOk

`func (o *FxSourceUniverseResponse) GetSourcePolicyOk() (*FxSourcePolicy, bool)`

GetSourcePolicyOk returns a tuple with the SourcePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePolicy

`func (o *FxSourceUniverseResponse) SetSourcePolicy(v FxSourcePolicy)`

SetSourcePolicy sets SourcePolicy field to given value.


### GetCurrency

`func (o *FxSourceUniverseResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FxSourceUniverseResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FxSourceUniverseResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FxSourceUniverseResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *FxSourceUniverseResponse) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *FxSourceUniverseResponse) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetSource

`func (o *FxSourceUniverseResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FxSourceUniverseResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FxSourceUniverseResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FxSourceUniverseResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *FxSourceUniverseResponse) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *FxSourceUniverseResponse) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetData

`func (o *FxSourceUniverseResponse) GetData() []FxSourceUniverseEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FxSourceUniverseResponse) GetDataOk() (*[]FxSourceUniverseEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FxSourceUniverseResponse) SetData(v []FxSourceUniverseEntry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


