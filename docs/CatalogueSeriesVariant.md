# CatalogueSeriesVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeriesId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**StorageIndicator** | Pointer to **NullableString** | Indicator slug whose observations make up this series variant. | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceSeriesId** | Pointer to **NullableString** |  | [optional] 
**SourceSeriesName** | Pointer to **NullableString** |  | [optional] 
**SourceUrl** | Pointer to **NullableString** |  | [optional] 
**SourceUrlScope** | Pointer to **NullableString** |  | [optional] 
**SeasonalAdjustment** | Pointer to **NullableString** |  | [optional] 
**PriceBasis** | Pointer to **NullableString** |  | [optional] 
**Annualization** | Pointer to **NullableString** |  | [optional] 
**PeriodAggregation** | Pointer to **NullableString** |  | [optional] 
**FrequencySelector** | Pointer to **NullableString** |  | [optional] 
**Unit** | Pointer to **NullableString** |  | [optional] 
**Frequency** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**HasOfficialForecast** | Pointer to **NullableBool** |  | [optional] 
**StandardizationNote** | Pointer to **NullableString** |  | [optional] 
**Coverage** | Pointer to [**NullableCatalogueCoverage**](CatalogueCoverage.md) |  | [optional] 

## Methods

### NewCatalogueSeriesVariant

`func NewCatalogueSeriesVariant() *CatalogueSeriesVariant`

NewCatalogueSeriesVariant instantiates a new CatalogueSeriesVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogueSeriesVariantWithDefaults

`func NewCatalogueSeriesVariantWithDefaults() *CatalogueSeriesVariant`

NewCatalogueSeriesVariantWithDefaults instantiates a new CatalogueSeriesVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeriesId

`func (o *CatalogueSeriesVariant) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *CatalogueSeriesVariant) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *CatalogueSeriesVariant) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *CatalogueSeriesVariant) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### SetSeriesIdNil

`func (o *CatalogueSeriesVariant) SetSeriesIdNil(b bool)`

 SetSeriesIdNil sets the value for SeriesId to be an explicit nil

### UnsetSeriesId
`func (o *CatalogueSeriesVariant) UnsetSeriesId()`

UnsetSeriesId ensures that no value is present for SeriesId, not even an explicit nil
### GetName

`func (o *CatalogueSeriesVariant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogueSeriesVariant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogueSeriesVariant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogueSeriesVariant) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CatalogueSeriesVariant) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CatalogueSeriesVariant) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStorageIndicator

`func (o *CatalogueSeriesVariant) GetStorageIndicator() string`

GetStorageIndicator returns the StorageIndicator field if non-nil, zero value otherwise.

### GetStorageIndicatorOk

`func (o *CatalogueSeriesVariant) GetStorageIndicatorOk() (*string, bool)`

GetStorageIndicatorOk returns a tuple with the StorageIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageIndicator

`func (o *CatalogueSeriesVariant) SetStorageIndicator(v string)`

SetStorageIndicator sets StorageIndicator field to given value.

### HasStorageIndicator

`func (o *CatalogueSeriesVariant) HasStorageIndicator() bool`

HasStorageIndicator returns a boolean if a field has been set.

### SetStorageIndicatorNil

`func (o *CatalogueSeriesVariant) SetStorageIndicatorNil(b bool)`

 SetStorageIndicatorNil sets the value for StorageIndicator to be an explicit nil

### UnsetStorageIndicator
`func (o *CatalogueSeriesVariant) UnsetStorageIndicator()`

UnsetStorageIndicator ensures that no value is present for StorageIndicator, not even an explicit nil
### GetSource

`func (o *CatalogueSeriesVariant) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CatalogueSeriesVariant) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CatalogueSeriesVariant) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CatalogueSeriesVariant) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *CatalogueSeriesVariant) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *CatalogueSeriesVariant) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceSeriesId

`func (o *CatalogueSeriesVariant) GetSourceSeriesId() string`

GetSourceSeriesId returns the SourceSeriesId field if non-nil, zero value otherwise.

### GetSourceSeriesIdOk

`func (o *CatalogueSeriesVariant) GetSourceSeriesIdOk() (*string, bool)`

GetSourceSeriesIdOk returns a tuple with the SourceSeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSeriesId

`func (o *CatalogueSeriesVariant) SetSourceSeriesId(v string)`

SetSourceSeriesId sets SourceSeriesId field to given value.

### HasSourceSeriesId

`func (o *CatalogueSeriesVariant) HasSourceSeriesId() bool`

HasSourceSeriesId returns a boolean if a field has been set.

### SetSourceSeriesIdNil

`func (o *CatalogueSeriesVariant) SetSourceSeriesIdNil(b bool)`

 SetSourceSeriesIdNil sets the value for SourceSeriesId to be an explicit nil

### UnsetSourceSeriesId
`func (o *CatalogueSeriesVariant) UnsetSourceSeriesId()`

UnsetSourceSeriesId ensures that no value is present for SourceSeriesId, not even an explicit nil
### GetSourceSeriesName

`func (o *CatalogueSeriesVariant) GetSourceSeriesName() string`

GetSourceSeriesName returns the SourceSeriesName field if non-nil, zero value otherwise.

### GetSourceSeriesNameOk

`func (o *CatalogueSeriesVariant) GetSourceSeriesNameOk() (*string, bool)`

GetSourceSeriesNameOk returns a tuple with the SourceSeriesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSeriesName

`func (o *CatalogueSeriesVariant) SetSourceSeriesName(v string)`

SetSourceSeriesName sets SourceSeriesName field to given value.

### HasSourceSeriesName

`func (o *CatalogueSeriesVariant) HasSourceSeriesName() bool`

HasSourceSeriesName returns a boolean if a field has been set.

### SetSourceSeriesNameNil

`func (o *CatalogueSeriesVariant) SetSourceSeriesNameNil(b bool)`

 SetSourceSeriesNameNil sets the value for SourceSeriesName to be an explicit nil

### UnsetSourceSeriesName
`func (o *CatalogueSeriesVariant) UnsetSourceSeriesName()`

UnsetSourceSeriesName ensures that no value is present for SourceSeriesName, not even an explicit nil
### GetSourceUrl

`func (o *CatalogueSeriesVariant) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *CatalogueSeriesVariant) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *CatalogueSeriesVariant) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *CatalogueSeriesVariant) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *CatalogueSeriesVariant) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *CatalogueSeriesVariant) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSourceUrlScope

`func (o *CatalogueSeriesVariant) GetSourceUrlScope() string`

GetSourceUrlScope returns the SourceUrlScope field if non-nil, zero value otherwise.

### GetSourceUrlScopeOk

`func (o *CatalogueSeriesVariant) GetSourceUrlScopeOk() (*string, bool)`

GetSourceUrlScopeOk returns a tuple with the SourceUrlScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrlScope

`func (o *CatalogueSeriesVariant) SetSourceUrlScope(v string)`

SetSourceUrlScope sets SourceUrlScope field to given value.

### HasSourceUrlScope

`func (o *CatalogueSeriesVariant) HasSourceUrlScope() bool`

HasSourceUrlScope returns a boolean if a field has been set.

### SetSourceUrlScopeNil

`func (o *CatalogueSeriesVariant) SetSourceUrlScopeNil(b bool)`

 SetSourceUrlScopeNil sets the value for SourceUrlScope to be an explicit nil

### UnsetSourceUrlScope
`func (o *CatalogueSeriesVariant) UnsetSourceUrlScope()`

UnsetSourceUrlScope ensures that no value is present for SourceUrlScope, not even an explicit nil
### GetSeasonalAdjustment

`func (o *CatalogueSeriesVariant) GetSeasonalAdjustment() string`

GetSeasonalAdjustment returns the SeasonalAdjustment field if non-nil, zero value otherwise.

### GetSeasonalAdjustmentOk

`func (o *CatalogueSeriesVariant) GetSeasonalAdjustmentOk() (*string, bool)`

GetSeasonalAdjustmentOk returns a tuple with the SeasonalAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonalAdjustment

`func (o *CatalogueSeriesVariant) SetSeasonalAdjustment(v string)`

SetSeasonalAdjustment sets SeasonalAdjustment field to given value.

### HasSeasonalAdjustment

`func (o *CatalogueSeriesVariant) HasSeasonalAdjustment() bool`

HasSeasonalAdjustment returns a boolean if a field has been set.

### SetSeasonalAdjustmentNil

`func (o *CatalogueSeriesVariant) SetSeasonalAdjustmentNil(b bool)`

 SetSeasonalAdjustmentNil sets the value for SeasonalAdjustment to be an explicit nil

### UnsetSeasonalAdjustment
`func (o *CatalogueSeriesVariant) UnsetSeasonalAdjustment()`

UnsetSeasonalAdjustment ensures that no value is present for SeasonalAdjustment, not even an explicit nil
### GetPriceBasis

`func (o *CatalogueSeriesVariant) GetPriceBasis() string`

GetPriceBasis returns the PriceBasis field if non-nil, zero value otherwise.

### GetPriceBasisOk

`func (o *CatalogueSeriesVariant) GetPriceBasisOk() (*string, bool)`

GetPriceBasisOk returns a tuple with the PriceBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceBasis

`func (o *CatalogueSeriesVariant) SetPriceBasis(v string)`

SetPriceBasis sets PriceBasis field to given value.

### HasPriceBasis

`func (o *CatalogueSeriesVariant) HasPriceBasis() bool`

HasPriceBasis returns a boolean if a field has been set.

### SetPriceBasisNil

`func (o *CatalogueSeriesVariant) SetPriceBasisNil(b bool)`

 SetPriceBasisNil sets the value for PriceBasis to be an explicit nil

### UnsetPriceBasis
`func (o *CatalogueSeriesVariant) UnsetPriceBasis()`

UnsetPriceBasis ensures that no value is present for PriceBasis, not even an explicit nil
### GetAnnualization

`func (o *CatalogueSeriesVariant) GetAnnualization() string`

GetAnnualization returns the Annualization field if non-nil, zero value otherwise.

### GetAnnualizationOk

`func (o *CatalogueSeriesVariant) GetAnnualizationOk() (*string, bool)`

GetAnnualizationOk returns a tuple with the Annualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualization

`func (o *CatalogueSeriesVariant) SetAnnualization(v string)`

SetAnnualization sets Annualization field to given value.

### HasAnnualization

`func (o *CatalogueSeriesVariant) HasAnnualization() bool`

HasAnnualization returns a boolean if a field has been set.

### SetAnnualizationNil

`func (o *CatalogueSeriesVariant) SetAnnualizationNil(b bool)`

 SetAnnualizationNil sets the value for Annualization to be an explicit nil

### UnsetAnnualization
`func (o *CatalogueSeriesVariant) UnsetAnnualization()`

UnsetAnnualization ensures that no value is present for Annualization, not even an explicit nil
### GetPeriodAggregation

`func (o *CatalogueSeriesVariant) GetPeriodAggregation() string`

GetPeriodAggregation returns the PeriodAggregation field if non-nil, zero value otherwise.

### GetPeriodAggregationOk

`func (o *CatalogueSeriesVariant) GetPeriodAggregationOk() (*string, bool)`

GetPeriodAggregationOk returns a tuple with the PeriodAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodAggregation

`func (o *CatalogueSeriesVariant) SetPeriodAggregation(v string)`

SetPeriodAggregation sets PeriodAggregation field to given value.

### HasPeriodAggregation

`func (o *CatalogueSeriesVariant) HasPeriodAggregation() bool`

HasPeriodAggregation returns a boolean if a field has been set.

### SetPeriodAggregationNil

`func (o *CatalogueSeriesVariant) SetPeriodAggregationNil(b bool)`

 SetPeriodAggregationNil sets the value for PeriodAggregation to be an explicit nil

### UnsetPeriodAggregation
`func (o *CatalogueSeriesVariant) UnsetPeriodAggregation()`

UnsetPeriodAggregation ensures that no value is present for PeriodAggregation, not even an explicit nil
### GetFrequencySelector

`func (o *CatalogueSeriesVariant) GetFrequencySelector() string`

GetFrequencySelector returns the FrequencySelector field if non-nil, zero value otherwise.

### GetFrequencySelectorOk

`func (o *CatalogueSeriesVariant) GetFrequencySelectorOk() (*string, bool)`

GetFrequencySelectorOk returns a tuple with the FrequencySelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencySelector

`func (o *CatalogueSeriesVariant) SetFrequencySelector(v string)`

SetFrequencySelector sets FrequencySelector field to given value.

### HasFrequencySelector

`func (o *CatalogueSeriesVariant) HasFrequencySelector() bool`

HasFrequencySelector returns a boolean if a field has been set.

### SetFrequencySelectorNil

`func (o *CatalogueSeriesVariant) SetFrequencySelectorNil(b bool)`

 SetFrequencySelectorNil sets the value for FrequencySelector to be an explicit nil

### UnsetFrequencySelector
`func (o *CatalogueSeriesVariant) UnsetFrequencySelector()`

UnsetFrequencySelector ensures that no value is present for FrequencySelector, not even an explicit nil
### GetUnit

`func (o *CatalogueSeriesVariant) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CatalogueSeriesVariant) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CatalogueSeriesVariant) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CatalogueSeriesVariant) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *CatalogueSeriesVariant) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *CatalogueSeriesVariant) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetFrequency

`func (o *CatalogueSeriesVariant) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *CatalogueSeriesVariant) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *CatalogueSeriesVariant) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *CatalogueSeriesVariant) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *CatalogueSeriesVariant) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *CatalogueSeriesVariant) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetIsDefault

`func (o *CatalogueSeriesVariant) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CatalogueSeriesVariant) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CatalogueSeriesVariant) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CatalogueSeriesVariant) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *CatalogueSeriesVariant) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *CatalogueSeriesVariant) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetHasOfficialForecast

`func (o *CatalogueSeriesVariant) GetHasOfficialForecast() bool`

GetHasOfficialForecast returns the HasOfficialForecast field if non-nil, zero value otherwise.

### GetHasOfficialForecastOk

`func (o *CatalogueSeriesVariant) GetHasOfficialForecastOk() (*bool, bool)`

GetHasOfficialForecastOk returns a tuple with the HasOfficialForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOfficialForecast

`func (o *CatalogueSeriesVariant) SetHasOfficialForecast(v bool)`

SetHasOfficialForecast sets HasOfficialForecast field to given value.

### HasHasOfficialForecast

`func (o *CatalogueSeriesVariant) HasHasOfficialForecast() bool`

HasHasOfficialForecast returns a boolean if a field has been set.

### SetHasOfficialForecastNil

`func (o *CatalogueSeriesVariant) SetHasOfficialForecastNil(b bool)`

 SetHasOfficialForecastNil sets the value for HasOfficialForecast to be an explicit nil

### UnsetHasOfficialForecast
`func (o *CatalogueSeriesVariant) UnsetHasOfficialForecast()`

UnsetHasOfficialForecast ensures that no value is present for HasOfficialForecast, not even an explicit nil
### GetStandardizationNote

`func (o *CatalogueSeriesVariant) GetStandardizationNote() string`

GetStandardizationNote returns the StandardizationNote field if non-nil, zero value otherwise.

### GetStandardizationNoteOk

`func (o *CatalogueSeriesVariant) GetStandardizationNoteOk() (*string, bool)`

GetStandardizationNoteOk returns a tuple with the StandardizationNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardizationNote

`func (o *CatalogueSeriesVariant) SetStandardizationNote(v string)`

SetStandardizationNote sets StandardizationNote field to given value.

### HasStandardizationNote

`func (o *CatalogueSeriesVariant) HasStandardizationNote() bool`

HasStandardizationNote returns a boolean if a field has been set.

### SetStandardizationNoteNil

`func (o *CatalogueSeriesVariant) SetStandardizationNoteNil(b bool)`

 SetStandardizationNoteNil sets the value for StandardizationNote to be an explicit nil

### UnsetStandardizationNote
`func (o *CatalogueSeriesVariant) UnsetStandardizationNote()`

UnsetStandardizationNote ensures that no value is present for StandardizationNote, not even an explicit nil
### GetCoverage

`func (o *CatalogueSeriesVariant) GetCoverage() CatalogueCoverage`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *CatalogueSeriesVariant) GetCoverageOk() (*CatalogueCoverage, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *CatalogueSeriesVariant) SetCoverage(v CatalogueCoverage)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *CatalogueSeriesVariant) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### SetCoverageNil

`func (o *CatalogueSeriesVariant) SetCoverageNil(b bool)`

 SetCoverageNil sets the value for Coverage to be an explicit nil

### UnsetCoverage
`func (o *CatalogueSeriesVariant) UnsetCoverage()`

UnsetCoverage ensures that no value is present for Coverage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


