# CommodityLatestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indicator** | Pointer to **NullableString** |  | [optional] 
**Unit** | Pointer to **NullableString** |  | [optional] 
**Frequency** | Pointer to **NullableString** |  | [optional] 
**HasOfficialForecast** | Pointer to **NullableBool** |  | [optional] 
**LastUpdated** | Pointer to **NullableString** |  | [optional] 
**DataQuality** | Pointer to [**NullableDataQuality**](DataQuality.md) |  | [optional] 
**Latest** | Pointer to [**NullableCommodityObservation**](CommodityObservation.md) |  | [optional] 
**Previous** | Pointer to [**NullableCommodityObservation**](CommodityObservation.md) |  | [optional] 
**PctDiffPrev** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewCommodityLatestItem

`func NewCommodityLatestItem() *CommodityLatestItem`

NewCommodityLatestItem instantiates a new CommodityLatestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommodityLatestItemWithDefaults

`func NewCommodityLatestItemWithDefaults() *CommodityLatestItem`

NewCommodityLatestItemWithDefaults instantiates a new CommodityLatestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndicator

`func (o *CommodityLatestItem) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *CommodityLatestItem) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *CommodityLatestItem) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.

### HasIndicator

`func (o *CommodityLatestItem) HasIndicator() bool`

HasIndicator returns a boolean if a field has been set.

### SetIndicatorNil

`func (o *CommodityLatestItem) SetIndicatorNil(b bool)`

 SetIndicatorNil sets the value for Indicator to be an explicit nil

### UnsetIndicator
`func (o *CommodityLatestItem) UnsetIndicator()`

UnsetIndicator ensures that no value is present for Indicator, not even an explicit nil
### GetUnit

`func (o *CommodityLatestItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CommodityLatestItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CommodityLatestItem) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CommodityLatestItem) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *CommodityLatestItem) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *CommodityLatestItem) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetFrequency

`func (o *CommodityLatestItem) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *CommodityLatestItem) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *CommodityLatestItem) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *CommodityLatestItem) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *CommodityLatestItem) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *CommodityLatestItem) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetHasOfficialForecast

`func (o *CommodityLatestItem) GetHasOfficialForecast() bool`

GetHasOfficialForecast returns the HasOfficialForecast field if non-nil, zero value otherwise.

### GetHasOfficialForecastOk

`func (o *CommodityLatestItem) GetHasOfficialForecastOk() (*bool, bool)`

GetHasOfficialForecastOk returns a tuple with the HasOfficialForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOfficialForecast

`func (o *CommodityLatestItem) SetHasOfficialForecast(v bool)`

SetHasOfficialForecast sets HasOfficialForecast field to given value.

### HasHasOfficialForecast

`func (o *CommodityLatestItem) HasHasOfficialForecast() bool`

HasHasOfficialForecast returns a boolean if a field has been set.

### SetHasOfficialForecastNil

`func (o *CommodityLatestItem) SetHasOfficialForecastNil(b bool)`

 SetHasOfficialForecastNil sets the value for HasOfficialForecast to be an explicit nil

### UnsetHasOfficialForecast
`func (o *CommodityLatestItem) UnsetHasOfficialForecast()`

UnsetHasOfficialForecast ensures that no value is present for HasOfficialForecast, not even an explicit nil
### GetLastUpdated

`func (o *CommodityLatestItem) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CommodityLatestItem) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CommodityLatestItem) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CommodityLatestItem) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *CommodityLatestItem) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *CommodityLatestItem) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetDataQuality

`func (o *CommodityLatestItem) GetDataQuality() DataQuality`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *CommodityLatestItem) GetDataQualityOk() (*DataQuality, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *CommodityLatestItem) SetDataQuality(v DataQuality)`

SetDataQuality sets DataQuality field to given value.

### HasDataQuality

`func (o *CommodityLatestItem) HasDataQuality() bool`

HasDataQuality returns a boolean if a field has been set.

### SetDataQualityNil

`func (o *CommodityLatestItem) SetDataQualityNil(b bool)`

 SetDataQualityNil sets the value for DataQuality to be an explicit nil

### UnsetDataQuality
`func (o *CommodityLatestItem) UnsetDataQuality()`

UnsetDataQuality ensures that no value is present for DataQuality, not even an explicit nil
### GetLatest

`func (o *CommodityLatestItem) GetLatest() CommodityObservation`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *CommodityLatestItem) GetLatestOk() (*CommodityObservation, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *CommodityLatestItem) SetLatest(v CommodityObservation)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *CommodityLatestItem) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### SetLatestNil

`func (o *CommodityLatestItem) SetLatestNil(b bool)`

 SetLatestNil sets the value for Latest to be an explicit nil

### UnsetLatest
`func (o *CommodityLatestItem) UnsetLatest()`

UnsetLatest ensures that no value is present for Latest, not even an explicit nil
### GetPrevious

`func (o *CommodityLatestItem) GetPrevious() CommodityObservation`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *CommodityLatestItem) GetPreviousOk() (*CommodityObservation, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *CommodityLatestItem) SetPrevious(v CommodityObservation)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *CommodityLatestItem) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *CommodityLatestItem) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *CommodityLatestItem) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetPctDiffPrev

`func (o *CommodityLatestItem) GetPctDiffPrev() float64`

GetPctDiffPrev returns the PctDiffPrev field if non-nil, zero value otherwise.

### GetPctDiffPrevOk

`func (o *CommodityLatestItem) GetPctDiffPrevOk() (*float64, bool)`

GetPctDiffPrevOk returns a tuple with the PctDiffPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctDiffPrev

`func (o *CommodityLatestItem) SetPctDiffPrev(v float64)`

SetPctDiffPrev sets PctDiffPrev field to given value.

### HasPctDiffPrev

`func (o *CommodityLatestItem) HasPctDiffPrev() bool`

HasPctDiffPrev returns a boolean if a field has been set.

### SetPctDiffPrevNil

`func (o *CommodityLatestItem) SetPctDiffPrevNil(b bool)`

 SetPctDiffPrevNil sets the value for PctDiffPrev to be an explicit nil

### UnsetPctDiffPrev
`func (o *CommodityLatestItem) UnsetPctDiffPrev()`

UnsetPctDiffPrev ensures that no value is present for PctDiffPrev, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


