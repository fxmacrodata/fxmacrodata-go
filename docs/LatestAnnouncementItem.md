# LatestAnnouncementItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indicator** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceUrl** | Pointer to **NullableString** |  | [optional] 
**SourceSeriesId** | Pointer to **NullableString** |  | [optional] 
**SourceSeriesName** | Pointer to **NullableString** |  | [optional] 
**SourceLocalName** | Pointer to **NullableString** |  | [optional] 
**SeasonalAdjustment** | Pointer to **NullableString** |  | [optional] 
**PriceBasis** | Pointer to **NullableString** |  | [optional] 
**IsProxy** | Pointer to **NullableBool** |  | [optional] 
**ProxyNote** | Pointer to **NullableString** |  | [optional] 
**Provenance** | Pointer to [**NullableLatestAnnouncementProvenance**](LatestAnnouncementProvenance.md) |  | [optional] 
**Unit** | Pointer to **NullableString** |  | [optional] 
**Frequency** | Pointer to **NullableString** |  | [optional] 
**HasOfficialForecast** | Pointer to **NullableBool** |  | [optional] 
**Latest** | Pointer to [**NullableLatestAnnouncementValue**](LatestAnnouncementValue.md) |  | [optional] 
**Previous** | Pointer to [**NullableLatestAnnouncementValue**](LatestAnnouncementValue.md) |  | [optional] 
**PctDiffPrev** | Pointer to **NullableFloat64** |  | [optional] 
**PctChangeYoy** | Pointer to **NullableFloat64** |  | [optional] 
**PctChangeQoq** | Pointer to **NullableFloat64** |  | [optional] 
**PctChangeMom** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewLatestAnnouncementItem

`func NewLatestAnnouncementItem() *LatestAnnouncementItem`

NewLatestAnnouncementItem instantiates a new LatestAnnouncementItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLatestAnnouncementItemWithDefaults

`func NewLatestAnnouncementItemWithDefaults() *LatestAnnouncementItem`

NewLatestAnnouncementItemWithDefaults instantiates a new LatestAnnouncementItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndicator

`func (o *LatestAnnouncementItem) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *LatestAnnouncementItem) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *LatestAnnouncementItem) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.

### HasIndicator

`func (o *LatestAnnouncementItem) HasIndicator() bool`

HasIndicator returns a boolean if a field has been set.

### SetIndicatorNil

`func (o *LatestAnnouncementItem) SetIndicatorNil(b bool)`

 SetIndicatorNil sets the value for Indicator to be an explicit nil

### UnsetIndicator
`func (o *LatestAnnouncementItem) UnsetIndicator()`

UnsetIndicator ensures that no value is present for Indicator, not even an explicit nil
### GetName

`func (o *LatestAnnouncementItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LatestAnnouncementItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LatestAnnouncementItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LatestAnnouncementItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LatestAnnouncementItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LatestAnnouncementItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSource

`func (o *LatestAnnouncementItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LatestAnnouncementItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LatestAnnouncementItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LatestAnnouncementItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *LatestAnnouncementItem) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *LatestAnnouncementItem) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceUrl

`func (o *LatestAnnouncementItem) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *LatestAnnouncementItem) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *LatestAnnouncementItem) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *LatestAnnouncementItem) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *LatestAnnouncementItem) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *LatestAnnouncementItem) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSourceSeriesId

`func (o *LatestAnnouncementItem) GetSourceSeriesId() string`

GetSourceSeriesId returns the SourceSeriesId field if non-nil, zero value otherwise.

### GetSourceSeriesIdOk

`func (o *LatestAnnouncementItem) GetSourceSeriesIdOk() (*string, bool)`

GetSourceSeriesIdOk returns a tuple with the SourceSeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSeriesId

`func (o *LatestAnnouncementItem) SetSourceSeriesId(v string)`

SetSourceSeriesId sets SourceSeriesId field to given value.

### HasSourceSeriesId

`func (o *LatestAnnouncementItem) HasSourceSeriesId() bool`

HasSourceSeriesId returns a boolean if a field has been set.

### SetSourceSeriesIdNil

`func (o *LatestAnnouncementItem) SetSourceSeriesIdNil(b bool)`

 SetSourceSeriesIdNil sets the value for SourceSeriesId to be an explicit nil

### UnsetSourceSeriesId
`func (o *LatestAnnouncementItem) UnsetSourceSeriesId()`

UnsetSourceSeriesId ensures that no value is present for SourceSeriesId, not even an explicit nil
### GetSourceSeriesName

`func (o *LatestAnnouncementItem) GetSourceSeriesName() string`

GetSourceSeriesName returns the SourceSeriesName field if non-nil, zero value otherwise.

### GetSourceSeriesNameOk

`func (o *LatestAnnouncementItem) GetSourceSeriesNameOk() (*string, bool)`

GetSourceSeriesNameOk returns a tuple with the SourceSeriesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSeriesName

`func (o *LatestAnnouncementItem) SetSourceSeriesName(v string)`

SetSourceSeriesName sets SourceSeriesName field to given value.

### HasSourceSeriesName

`func (o *LatestAnnouncementItem) HasSourceSeriesName() bool`

HasSourceSeriesName returns a boolean if a field has been set.

### SetSourceSeriesNameNil

`func (o *LatestAnnouncementItem) SetSourceSeriesNameNil(b bool)`

 SetSourceSeriesNameNil sets the value for SourceSeriesName to be an explicit nil

### UnsetSourceSeriesName
`func (o *LatestAnnouncementItem) UnsetSourceSeriesName()`

UnsetSourceSeriesName ensures that no value is present for SourceSeriesName, not even an explicit nil
### GetSourceLocalName

`func (o *LatestAnnouncementItem) GetSourceLocalName() string`

GetSourceLocalName returns the SourceLocalName field if non-nil, zero value otherwise.

### GetSourceLocalNameOk

`func (o *LatestAnnouncementItem) GetSourceLocalNameOk() (*string, bool)`

GetSourceLocalNameOk returns a tuple with the SourceLocalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLocalName

`func (o *LatestAnnouncementItem) SetSourceLocalName(v string)`

SetSourceLocalName sets SourceLocalName field to given value.

### HasSourceLocalName

`func (o *LatestAnnouncementItem) HasSourceLocalName() bool`

HasSourceLocalName returns a boolean if a field has been set.

### SetSourceLocalNameNil

`func (o *LatestAnnouncementItem) SetSourceLocalNameNil(b bool)`

 SetSourceLocalNameNil sets the value for SourceLocalName to be an explicit nil

### UnsetSourceLocalName
`func (o *LatestAnnouncementItem) UnsetSourceLocalName()`

UnsetSourceLocalName ensures that no value is present for SourceLocalName, not even an explicit nil
### GetSeasonalAdjustment

`func (o *LatestAnnouncementItem) GetSeasonalAdjustment() string`

GetSeasonalAdjustment returns the SeasonalAdjustment field if non-nil, zero value otherwise.

### GetSeasonalAdjustmentOk

`func (o *LatestAnnouncementItem) GetSeasonalAdjustmentOk() (*string, bool)`

GetSeasonalAdjustmentOk returns a tuple with the SeasonalAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonalAdjustment

`func (o *LatestAnnouncementItem) SetSeasonalAdjustment(v string)`

SetSeasonalAdjustment sets SeasonalAdjustment field to given value.

### HasSeasonalAdjustment

`func (o *LatestAnnouncementItem) HasSeasonalAdjustment() bool`

HasSeasonalAdjustment returns a boolean if a field has been set.

### SetSeasonalAdjustmentNil

`func (o *LatestAnnouncementItem) SetSeasonalAdjustmentNil(b bool)`

 SetSeasonalAdjustmentNil sets the value for SeasonalAdjustment to be an explicit nil

### UnsetSeasonalAdjustment
`func (o *LatestAnnouncementItem) UnsetSeasonalAdjustment()`

UnsetSeasonalAdjustment ensures that no value is present for SeasonalAdjustment, not even an explicit nil
### GetPriceBasis

`func (o *LatestAnnouncementItem) GetPriceBasis() string`

GetPriceBasis returns the PriceBasis field if non-nil, zero value otherwise.

### GetPriceBasisOk

`func (o *LatestAnnouncementItem) GetPriceBasisOk() (*string, bool)`

GetPriceBasisOk returns a tuple with the PriceBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceBasis

`func (o *LatestAnnouncementItem) SetPriceBasis(v string)`

SetPriceBasis sets PriceBasis field to given value.

### HasPriceBasis

`func (o *LatestAnnouncementItem) HasPriceBasis() bool`

HasPriceBasis returns a boolean if a field has been set.

### SetPriceBasisNil

`func (o *LatestAnnouncementItem) SetPriceBasisNil(b bool)`

 SetPriceBasisNil sets the value for PriceBasis to be an explicit nil

### UnsetPriceBasis
`func (o *LatestAnnouncementItem) UnsetPriceBasis()`

UnsetPriceBasis ensures that no value is present for PriceBasis, not even an explicit nil
### GetIsProxy

`func (o *LatestAnnouncementItem) GetIsProxy() bool`

GetIsProxy returns the IsProxy field if non-nil, zero value otherwise.

### GetIsProxyOk

`func (o *LatestAnnouncementItem) GetIsProxyOk() (*bool, bool)`

GetIsProxyOk returns a tuple with the IsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxy

`func (o *LatestAnnouncementItem) SetIsProxy(v bool)`

SetIsProxy sets IsProxy field to given value.

### HasIsProxy

`func (o *LatestAnnouncementItem) HasIsProxy() bool`

HasIsProxy returns a boolean if a field has been set.

### SetIsProxyNil

`func (o *LatestAnnouncementItem) SetIsProxyNil(b bool)`

 SetIsProxyNil sets the value for IsProxy to be an explicit nil

### UnsetIsProxy
`func (o *LatestAnnouncementItem) UnsetIsProxy()`

UnsetIsProxy ensures that no value is present for IsProxy, not even an explicit nil
### GetProxyNote

`func (o *LatestAnnouncementItem) GetProxyNote() string`

GetProxyNote returns the ProxyNote field if non-nil, zero value otherwise.

### GetProxyNoteOk

`func (o *LatestAnnouncementItem) GetProxyNoteOk() (*string, bool)`

GetProxyNoteOk returns a tuple with the ProxyNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyNote

`func (o *LatestAnnouncementItem) SetProxyNote(v string)`

SetProxyNote sets ProxyNote field to given value.

### HasProxyNote

`func (o *LatestAnnouncementItem) HasProxyNote() bool`

HasProxyNote returns a boolean if a field has been set.

### SetProxyNoteNil

`func (o *LatestAnnouncementItem) SetProxyNoteNil(b bool)`

 SetProxyNoteNil sets the value for ProxyNote to be an explicit nil

### UnsetProxyNote
`func (o *LatestAnnouncementItem) UnsetProxyNote()`

UnsetProxyNote ensures that no value is present for ProxyNote, not even an explicit nil
### GetProvenance

`func (o *LatestAnnouncementItem) GetProvenance() LatestAnnouncementProvenance`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *LatestAnnouncementItem) GetProvenanceOk() (*LatestAnnouncementProvenance, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *LatestAnnouncementItem) SetProvenance(v LatestAnnouncementProvenance)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *LatestAnnouncementItem) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### SetProvenanceNil

`func (o *LatestAnnouncementItem) SetProvenanceNil(b bool)`

 SetProvenanceNil sets the value for Provenance to be an explicit nil

### UnsetProvenance
`func (o *LatestAnnouncementItem) UnsetProvenance()`

UnsetProvenance ensures that no value is present for Provenance, not even an explicit nil
### GetUnit

`func (o *LatestAnnouncementItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *LatestAnnouncementItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *LatestAnnouncementItem) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *LatestAnnouncementItem) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *LatestAnnouncementItem) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *LatestAnnouncementItem) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetFrequency

`func (o *LatestAnnouncementItem) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *LatestAnnouncementItem) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *LatestAnnouncementItem) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *LatestAnnouncementItem) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *LatestAnnouncementItem) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *LatestAnnouncementItem) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetHasOfficialForecast

`func (o *LatestAnnouncementItem) GetHasOfficialForecast() bool`

GetHasOfficialForecast returns the HasOfficialForecast field if non-nil, zero value otherwise.

### GetHasOfficialForecastOk

`func (o *LatestAnnouncementItem) GetHasOfficialForecastOk() (*bool, bool)`

GetHasOfficialForecastOk returns a tuple with the HasOfficialForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOfficialForecast

`func (o *LatestAnnouncementItem) SetHasOfficialForecast(v bool)`

SetHasOfficialForecast sets HasOfficialForecast field to given value.

### HasHasOfficialForecast

`func (o *LatestAnnouncementItem) HasHasOfficialForecast() bool`

HasHasOfficialForecast returns a boolean if a field has been set.

### SetHasOfficialForecastNil

`func (o *LatestAnnouncementItem) SetHasOfficialForecastNil(b bool)`

 SetHasOfficialForecastNil sets the value for HasOfficialForecast to be an explicit nil

### UnsetHasOfficialForecast
`func (o *LatestAnnouncementItem) UnsetHasOfficialForecast()`

UnsetHasOfficialForecast ensures that no value is present for HasOfficialForecast, not even an explicit nil
### GetLatest

`func (o *LatestAnnouncementItem) GetLatest() LatestAnnouncementValue`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *LatestAnnouncementItem) GetLatestOk() (*LatestAnnouncementValue, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *LatestAnnouncementItem) SetLatest(v LatestAnnouncementValue)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *LatestAnnouncementItem) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### SetLatestNil

`func (o *LatestAnnouncementItem) SetLatestNil(b bool)`

 SetLatestNil sets the value for Latest to be an explicit nil

### UnsetLatest
`func (o *LatestAnnouncementItem) UnsetLatest()`

UnsetLatest ensures that no value is present for Latest, not even an explicit nil
### GetPrevious

`func (o *LatestAnnouncementItem) GetPrevious() LatestAnnouncementValue`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *LatestAnnouncementItem) GetPreviousOk() (*LatestAnnouncementValue, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *LatestAnnouncementItem) SetPrevious(v LatestAnnouncementValue)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *LatestAnnouncementItem) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *LatestAnnouncementItem) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *LatestAnnouncementItem) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetPctDiffPrev

`func (o *LatestAnnouncementItem) GetPctDiffPrev() float64`

GetPctDiffPrev returns the PctDiffPrev field if non-nil, zero value otherwise.

### GetPctDiffPrevOk

`func (o *LatestAnnouncementItem) GetPctDiffPrevOk() (*float64, bool)`

GetPctDiffPrevOk returns a tuple with the PctDiffPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctDiffPrev

`func (o *LatestAnnouncementItem) SetPctDiffPrev(v float64)`

SetPctDiffPrev sets PctDiffPrev field to given value.

### HasPctDiffPrev

`func (o *LatestAnnouncementItem) HasPctDiffPrev() bool`

HasPctDiffPrev returns a boolean if a field has been set.

### SetPctDiffPrevNil

`func (o *LatestAnnouncementItem) SetPctDiffPrevNil(b bool)`

 SetPctDiffPrevNil sets the value for PctDiffPrev to be an explicit nil

### UnsetPctDiffPrev
`func (o *LatestAnnouncementItem) UnsetPctDiffPrev()`

UnsetPctDiffPrev ensures that no value is present for PctDiffPrev, not even an explicit nil
### GetPctChangeYoy

`func (o *LatestAnnouncementItem) GetPctChangeYoy() float64`

GetPctChangeYoy returns the PctChangeYoy field if non-nil, zero value otherwise.

### GetPctChangeYoyOk

`func (o *LatestAnnouncementItem) GetPctChangeYoyOk() (*float64, bool)`

GetPctChangeYoyOk returns a tuple with the PctChangeYoy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctChangeYoy

`func (o *LatestAnnouncementItem) SetPctChangeYoy(v float64)`

SetPctChangeYoy sets PctChangeYoy field to given value.

### HasPctChangeYoy

`func (o *LatestAnnouncementItem) HasPctChangeYoy() bool`

HasPctChangeYoy returns a boolean if a field has been set.

### SetPctChangeYoyNil

`func (o *LatestAnnouncementItem) SetPctChangeYoyNil(b bool)`

 SetPctChangeYoyNil sets the value for PctChangeYoy to be an explicit nil

### UnsetPctChangeYoy
`func (o *LatestAnnouncementItem) UnsetPctChangeYoy()`

UnsetPctChangeYoy ensures that no value is present for PctChangeYoy, not even an explicit nil
### GetPctChangeQoq

`func (o *LatestAnnouncementItem) GetPctChangeQoq() float64`

GetPctChangeQoq returns the PctChangeQoq field if non-nil, zero value otherwise.

### GetPctChangeQoqOk

`func (o *LatestAnnouncementItem) GetPctChangeQoqOk() (*float64, bool)`

GetPctChangeQoqOk returns a tuple with the PctChangeQoq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctChangeQoq

`func (o *LatestAnnouncementItem) SetPctChangeQoq(v float64)`

SetPctChangeQoq sets PctChangeQoq field to given value.

### HasPctChangeQoq

`func (o *LatestAnnouncementItem) HasPctChangeQoq() bool`

HasPctChangeQoq returns a boolean if a field has been set.

### SetPctChangeQoqNil

`func (o *LatestAnnouncementItem) SetPctChangeQoqNil(b bool)`

 SetPctChangeQoqNil sets the value for PctChangeQoq to be an explicit nil

### UnsetPctChangeQoq
`func (o *LatestAnnouncementItem) UnsetPctChangeQoq()`

UnsetPctChangeQoq ensures that no value is present for PctChangeQoq, not even an explicit nil
### GetPctChangeMom

`func (o *LatestAnnouncementItem) GetPctChangeMom() float64`

GetPctChangeMom returns the PctChangeMom field if non-nil, zero value otherwise.

### GetPctChangeMomOk

`func (o *LatestAnnouncementItem) GetPctChangeMomOk() (*float64, bool)`

GetPctChangeMomOk returns a tuple with the PctChangeMom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctChangeMom

`func (o *LatestAnnouncementItem) SetPctChangeMom(v float64)`

SetPctChangeMom sets PctChangeMom field to given value.

### HasPctChangeMom

`func (o *LatestAnnouncementItem) HasPctChangeMom() bool`

HasPctChangeMom returns a boolean if a field has been set.

### SetPctChangeMomNil

`func (o *LatestAnnouncementItem) SetPctChangeMomNil(b bool)`

 SetPctChangeMomNil sets the value for PctChangeMom to be an explicit nil

### UnsetPctChangeMom
`func (o *LatestAnnouncementItem) UnsetPctChangeMom()`

UnsetPctChangeMom ensures that no value is present for PctChangeMom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


