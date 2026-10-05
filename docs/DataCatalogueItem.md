# DataCatalogueItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Unit** | Pointer to **NullableString** |  | [optional] 
**Frequency** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceUrl** | Pointer to **NullableString** |  | [optional] 
**SourceUrlScope** | Pointer to **NullableString** |  | [optional] 
**SourceSeriesId** | Pointer to **NullableString** |  | [optional] 
**SourceSeriesName** | Pointer to **NullableString** |  | [optional] 
**SourceHistoryStart** | Pointer to **NullableString** |  | [optional] 
**SeasonalAdjustment** | Pointer to **NullableString** |  | [optional] 
**PriceBasis** | Pointer to **NullableString** |  | [optional] 
**Annualization** | Pointer to **NullableString** |  | [optional] 
**PeriodAggregation** | Pointer to **NullableString** |  | [optional] 
**HasOfficialForecast** | Pointer to **NullableBool** |  | [optional] 
**IsProxy** | Pointer to **NullableBool** |  | [optional] 
**ProxyNote** | Pointer to **NullableString** |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] 
**RelatedIndicators** | Pointer to **[]string** |  | [optional] 
**StandardizationNote** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**LocalName** | Pointer to **NullableString** |  | [optional] 
**MaturityMonths** | Pointer to **NullableFloat64** |  | [optional] 
**YieldType** | Pointer to **NullableString** |  | [optional] 
**Geography** | Pointer to **NullableString** |  | [optional] 
**PeriodBasis** | Pointer to **NullableString** |  | [optional] 
**MethodologyId** | Pointer to **NullableString** |  | [optional] 
**FreshnessThresholdDays** | Pointer to **NullableInt64** |  | [optional] 
**ReleaseCalendarUrl** | Pointer to **NullableString** |  | [optional] 
**SourcePermissionStatus** | Pointer to **NullableString** |  | [optional] 
**SourceTermsUrl** | Pointer to **NullableString** |  | [optional] 
**SupportedFrequencyOptions** | Pointer to **[]string** |  | [optional] 
**SeriesVariants** | Pointer to [**[]CatalogueSeriesVariant**](CatalogueSeriesVariant.md) |  | [optional] 
**Coverage** | Pointer to [**NullableCatalogueCoverage**](CatalogueCoverage.md) |  | [optional] 
**SupportedOptions** | Pointer to [**NullableCatalogueSupportedOptions**](CatalogueSupportedOptions.md) |  | [optional] 

## Methods

### NewDataCatalogueItem

`func NewDataCatalogueItem() *DataCatalogueItem`

NewDataCatalogueItem instantiates a new DataCatalogueItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataCatalogueItemWithDefaults

`func NewDataCatalogueItemWithDefaults() *DataCatalogueItem`

NewDataCatalogueItemWithDefaults instantiates a new DataCatalogueItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataCatalogueItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataCatalogueItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataCatalogueItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataCatalogueItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DataCatalogueItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DataCatalogueItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUnit

`func (o *DataCatalogueItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *DataCatalogueItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *DataCatalogueItem) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *DataCatalogueItem) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *DataCatalogueItem) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *DataCatalogueItem) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetFrequency

`func (o *DataCatalogueItem) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *DataCatalogueItem) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *DataCatalogueItem) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *DataCatalogueItem) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *DataCatalogueItem) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *DataCatalogueItem) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetSource

`func (o *DataCatalogueItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DataCatalogueItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DataCatalogueItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DataCatalogueItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *DataCatalogueItem) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *DataCatalogueItem) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceUrl

`func (o *DataCatalogueItem) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *DataCatalogueItem) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *DataCatalogueItem) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *DataCatalogueItem) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *DataCatalogueItem) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *DataCatalogueItem) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSourceUrlScope

`func (o *DataCatalogueItem) GetSourceUrlScope() string`

GetSourceUrlScope returns the SourceUrlScope field if non-nil, zero value otherwise.

### GetSourceUrlScopeOk

`func (o *DataCatalogueItem) GetSourceUrlScopeOk() (*string, bool)`

GetSourceUrlScopeOk returns a tuple with the SourceUrlScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrlScope

`func (o *DataCatalogueItem) SetSourceUrlScope(v string)`

SetSourceUrlScope sets SourceUrlScope field to given value.

### HasSourceUrlScope

`func (o *DataCatalogueItem) HasSourceUrlScope() bool`

HasSourceUrlScope returns a boolean if a field has been set.

### SetSourceUrlScopeNil

`func (o *DataCatalogueItem) SetSourceUrlScopeNil(b bool)`

 SetSourceUrlScopeNil sets the value for SourceUrlScope to be an explicit nil

### UnsetSourceUrlScope
`func (o *DataCatalogueItem) UnsetSourceUrlScope()`

UnsetSourceUrlScope ensures that no value is present for SourceUrlScope, not even an explicit nil
### GetSourceSeriesId

`func (o *DataCatalogueItem) GetSourceSeriesId() string`

GetSourceSeriesId returns the SourceSeriesId field if non-nil, zero value otherwise.

### GetSourceSeriesIdOk

`func (o *DataCatalogueItem) GetSourceSeriesIdOk() (*string, bool)`

GetSourceSeriesIdOk returns a tuple with the SourceSeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSeriesId

`func (o *DataCatalogueItem) SetSourceSeriesId(v string)`

SetSourceSeriesId sets SourceSeriesId field to given value.

### HasSourceSeriesId

`func (o *DataCatalogueItem) HasSourceSeriesId() bool`

HasSourceSeriesId returns a boolean if a field has been set.

### SetSourceSeriesIdNil

`func (o *DataCatalogueItem) SetSourceSeriesIdNil(b bool)`

 SetSourceSeriesIdNil sets the value for SourceSeriesId to be an explicit nil

### UnsetSourceSeriesId
`func (o *DataCatalogueItem) UnsetSourceSeriesId()`

UnsetSourceSeriesId ensures that no value is present for SourceSeriesId, not even an explicit nil
### GetSourceSeriesName

`func (o *DataCatalogueItem) GetSourceSeriesName() string`

GetSourceSeriesName returns the SourceSeriesName field if non-nil, zero value otherwise.

### GetSourceSeriesNameOk

`func (o *DataCatalogueItem) GetSourceSeriesNameOk() (*string, bool)`

GetSourceSeriesNameOk returns a tuple with the SourceSeriesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSeriesName

`func (o *DataCatalogueItem) SetSourceSeriesName(v string)`

SetSourceSeriesName sets SourceSeriesName field to given value.

### HasSourceSeriesName

`func (o *DataCatalogueItem) HasSourceSeriesName() bool`

HasSourceSeriesName returns a boolean if a field has been set.

### SetSourceSeriesNameNil

`func (o *DataCatalogueItem) SetSourceSeriesNameNil(b bool)`

 SetSourceSeriesNameNil sets the value for SourceSeriesName to be an explicit nil

### UnsetSourceSeriesName
`func (o *DataCatalogueItem) UnsetSourceSeriesName()`

UnsetSourceSeriesName ensures that no value is present for SourceSeriesName, not even an explicit nil
### GetSourceHistoryStart

`func (o *DataCatalogueItem) GetSourceHistoryStart() string`

GetSourceHistoryStart returns the SourceHistoryStart field if non-nil, zero value otherwise.

### GetSourceHistoryStartOk

`func (o *DataCatalogueItem) GetSourceHistoryStartOk() (*string, bool)`

GetSourceHistoryStartOk returns a tuple with the SourceHistoryStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHistoryStart

`func (o *DataCatalogueItem) SetSourceHistoryStart(v string)`

SetSourceHistoryStart sets SourceHistoryStart field to given value.

### HasSourceHistoryStart

`func (o *DataCatalogueItem) HasSourceHistoryStart() bool`

HasSourceHistoryStart returns a boolean if a field has been set.

### SetSourceHistoryStartNil

`func (o *DataCatalogueItem) SetSourceHistoryStartNil(b bool)`

 SetSourceHistoryStartNil sets the value for SourceHistoryStart to be an explicit nil

### UnsetSourceHistoryStart
`func (o *DataCatalogueItem) UnsetSourceHistoryStart()`

UnsetSourceHistoryStart ensures that no value is present for SourceHistoryStart, not even an explicit nil
### GetSeasonalAdjustment

`func (o *DataCatalogueItem) GetSeasonalAdjustment() string`

GetSeasonalAdjustment returns the SeasonalAdjustment field if non-nil, zero value otherwise.

### GetSeasonalAdjustmentOk

`func (o *DataCatalogueItem) GetSeasonalAdjustmentOk() (*string, bool)`

GetSeasonalAdjustmentOk returns a tuple with the SeasonalAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonalAdjustment

`func (o *DataCatalogueItem) SetSeasonalAdjustment(v string)`

SetSeasonalAdjustment sets SeasonalAdjustment field to given value.

### HasSeasonalAdjustment

`func (o *DataCatalogueItem) HasSeasonalAdjustment() bool`

HasSeasonalAdjustment returns a boolean if a field has been set.

### SetSeasonalAdjustmentNil

`func (o *DataCatalogueItem) SetSeasonalAdjustmentNil(b bool)`

 SetSeasonalAdjustmentNil sets the value for SeasonalAdjustment to be an explicit nil

### UnsetSeasonalAdjustment
`func (o *DataCatalogueItem) UnsetSeasonalAdjustment()`

UnsetSeasonalAdjustment ensures that no value is present for SeasonalAdjustment, not even an explicit nil
### GetPriceBasis

`func (o *DataCatalogueItem) GetPriceBasis() string`

GetPriceBasis returns the PriceBasis field if non-nil, zero value otherwise.

### GetPriceBasisOk

`func (o *DataCatalogueItem) GetPriceBasisOk() (*string, bool)`

GetPriceBasisOk returns a tuple with the PriceBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceBasis

`func (o *DataCatalogueItem) SetPriceBasis(v string)`

SetPriceBasis sets PriceBasis field to given value.

### HasPriceBasis

`func (o *DataCatalogueItem) HasPriceBasis() bool`

HasPriceBasis returns a boolean if a field has been set.

### SetPriceBasisNil

`func (o *DataCatalogueItem) SetPriceBasisNil(b bool)`

 SetPriceBasisNil sets the value for PriceBasis to be an explicit nil

### UnsetPriceBasis
`func (o *DataCatalogueItem) UnsetPriceBasis()`

UnsetPriceBasis ensures that no value is present for PriceBasis, not even an explicit nil
### GetAnnualization

`func (o *DataCatalogueItem) GetAnnualization() string`

GetAnnualization returns the Annualization field if non-nil, zero value otherwise.

### GetAnnualizationOk

`func (o *DataCatalogueItem) GetAnnualizationOk() (*string, bool)`

GetAnnualizationOk returns a tuple with the Annualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualization

`func (o *DataCatalogueItem) SetAnnualization(v string)`

SetAnnualization sets Annualization field to given value.

### HasAnnualization

`func (o *DataCatalogueItem) HasAnnualization() bool`

HasAnnualization returns a boolean if a field has been set.

### SetAnnualizationNil

`func (o *DataCatalogueItem) SetAnnualizationNil(b bool)`

 SetAnnualizationNil sets the value for Annualization to be an explicit nil

### UnsetAnnualization
`func (o *DataCatalogueItem) UnsetAnnualization()`

UnsetAnnualization ensures that no value is present for Annualization, not even an explicit nil
### GetPeriodAggregation

`func (o *DataCatalogueItem) GetPeriodAggregation() string`

GetPeriodAggregation returns the PeriodAggregation field if non-nil, zero value otherwise.

### GetPeriodAggregationOk

`func (o *DataCatalogueItem) GetPeriodAggregationOk() (*string, bool)`

GetPeriodAggregationOk returns a tuple with the PeriodAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodAggregation

`func (o *DataCatalogueItem) SetPeriodAggregation(v string)`

SetPeriodAggregation sets PeriodAggregation field to given value.

### HasPeriodAggregation

`func (o *DataCatalogueItem) HasPeriodAggregation() bool`

HasPeriodAggregation returns a boolean if a field has been set.

### SetPeriodAggregationNil

`func (o *DataCatalogueItem) SetPeriodAggregationNil(b bool)`

 SetPeriodAggregationNil sets the value for PeriodAggregation to be an explicit nil

### UnsetPeriodAggregation
`func (o *DataCatalogueItem) UnsetPeriodAggregation()`

UnsetPeriodAggregation ensures that no value is present for PeriodAggregation, not even an explicit nil
### GetHasOfficialForecast

`func (o *DataCatalogueItem) GetHasOfficialForecast() bool`

GetHasOfficialForecast returns the HasOfficialForecast field if non-nil, zero value otherwise.

### GetHasOfficialForecastOk

`func (o *DataCatalogueItem) GetHasOfficialForecastOk() (*bool, bool)`

GetHasOfficialForecastOk returns a tuple with the HasOfficialForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOfficialForecast

`func (o *DataCatalogueItem) SetHasOfficialForecast(v bool)`

SetHasOfficialForecast sets HasOfficialForecast field to given value.

### HasHasOfficialForecast

`func (o *DataCatalogueItem) HasHasOfficialForecast() bool`

HasHasOfficialForecast returns a boolean if a field has been set.

### SetHasOfficialForecastNil

`func (o *DataCatalogueItem) SetHasOfficialForecastNil(b bool)`

 SetHasOfficialForecastNil sets the value for HasOfficialForecast to be an explicit nil

### UnsetHasOfficialForecast
`func (o *DataCatalogueItem) UnsetHasOfficialForecast()`

UnsetHasOfficialForecast ensures that no value is present for HasOfficialForecast, not even an explicit nil
### GetIsProxy

`func (o *DataCatalogueItem) GetIsProxy() bool`

GetIsProxy returns the IsProxy field if non-nil, zero value otherwise.

### GetIsProxyOk

`func (o *DataCatalogueItem) GetIsProxyOk() (*bool, bool)`

GetIsProxyOk returns a tuple with the IsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxy

`func (o *DataCatalogueItem) SetIsProxy(v bool)`

SetIsProxy sets IsProxy field to given value.

### HasIsProxy

`func (o *DataCatalogueItem) HasIsProxy() bool`

HasIsProxy returns a boolean if a field has been set.

### SetIsProxyNil

`func (o *DataCatalogueItem) SetIsProxyNil(b bool)`

 SetIsProxyNil sets the value for IsProxy to be an explicit nil

### UnsetIsProxy
`func (o *DataCatalogueItem) UnsetIsProxy()`

UnsetIsProxy ensures that no value is present for IsProxy, not even an explicit nil
### GetProxyNote

`func (o *DataCatalogueItem) GetProxyNote() string`

GetProxyNote returns the ProxyNote field if non-nil, zero value otherwise.

### GetProxyNoteOk

`func (o *DataCatalogueItem) GetProxyNoteOk() (*string, bool)`

GetProxyNoteOk returns a tuple with the ProxyNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyNote

`func (o *DataCatalogueItem) SetProxyNote(v string)`

SetProxyNote sets ProxyNote field to given value.

### HasProxyNote

`func (o *DataCatalogueItem) HasProxyNote() bool`

HasProxyNote returns a boolean if a field has been set.

### SetProxyNoteNil

`func (o *DataCatalogueItem) SetProxyNoteNil(b bool)`

 SetProxyNoteNil sets the value for ProxyNote to be an explicit nil

### UnsetProxyNote
`func (o *DataCatalogueItem) UnsetProxyNote()`

UnsetProxyNote ensures that no value is present for ProxyNote, not even an explicit nil
### GetAliases

`func (o *DataCatalogueItem) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *DataCatalogueItem) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *DataCatalogueItem) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *DataCatalogueItem) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### SetAliasesNil

`func (o *DataCatalogueItem) SetAliasesNil(b bool)`

 SetAliasesNil sets the value for Aliases to be an explicit nil

### UnsetAliases
`func (o *DataCatalogueItem) UnsetAliases()`

UnsetAliases ensures that no value is present for Aliases, not even an explicit nil
### GetRelatedIndicators

`func (o *DataCatalogueItem) GetRelatedIndicators() []string`

GetRelatedIndicators returns the RelatedIndicators field if non-nil, zero value otherwise.

### GetRelatedIndicatorsOk

`func (o *DataCatalogueItem) GetRelatedIndicatorsOk() (*[]string, bool)`

GetRelatedIndicatorsOk returns a tuple with the RelatedIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIndicators

`func (o *DataCatalogueItem) SetRelatedIndicators(v []string)`

SetRelatedIndicators sets RelatedIndicators field to given value.

### HasRelatedIndicators

`func (o *DataCatalogueItem) HasRelatedIndicators() bool`

HasRelatedIndicators returns a boolean if a field has been set.

### SetRelatedIndicatorsNil

`func (o *DataCatalogueItem) SetRelatedIndicatorsNil(b bool)`

 SetRelatedIndicatorsNil sets the value for RelatedIndicators to be an explicit nil

### UnsetRelatedIndicators
`func (o *DataCatalogueItem) UnsetRelatedIndicators()`

UnsetRelatedIndicators ensures that no value is present for RelatedIndicators, not even an explicit nil
### GetStandardizationNote

`func (o *DataCatalogueItem) GetStandardizationNote() string`

GetStandardizationNote returns the StandardizationNote field if non-nil, zero value otherwise.

### GetStandardizationNoteOk

`func (o *DataCatalogueItem) GetStandardizationNoteOk() (*string, bool)`

GetStandardizationNoteOk returns a tuple with the StandardizationNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardizationNote

`func (o *DataCatalogueItem) SetStandardizationNote(v string)`

SetStandardizationNote sets StandardizationNote field to given value.

### HasStandardizationNote

`func (o *DataCatalogueItem) HasStandardizationNote() bool`

HasStandardizationNote returns a boolean if a field has been set.

### SetStandardizationNoteNil

`func (o *DataCatalogueItem) SetStandardizationNoteNil(b bool)`

 SetStandardizationNoteNil sets the value for StandardizationNote to be an explicit nil

### UnsetStandardizationNote
`func (o *DataCatalogueItem) UnsetStandardizationNote()`

UnsetStandardizationNote ensures that no value is present for StandardizationNote, not even an explicit nil
### GetDescription

`func (o *DataCatalogueItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataCatalogueItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataCatalogueItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataCatalogueItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DataCatalogueItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DataCatalogueItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLocalName

`func (o *DataCatalogueItem) GetLocalName() string`

GetLocalName returns the LocalName field if non-nil, zero value otherwise.

### GetLocalNameOk

`func (o *DataCatalogueItem) GetLocalNameOk() (*string, bool)`

GetLocalNameOk returns a tuple with the LocalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalName

`func (o *DataCatalogueItem) SetLocalName(v string)`

SetLocalName sets LocalName field to given value.

### HasLocalName

`func (o *DataCatalogueItem) HasLocalName() bool`

HasLocalName returns a boolean if a field has been set.

### SetLocalNameNil

`func (o *DataCatalogueItem) SetLocalNameNil(b bool)`

 SetLocalNameNil sets the value for LocalName to be an explicit nil

### UnsetLocalName
`func (o *DataCatalogueItem) UnsetLocalName()`

UnsetLocalName ensures that no value is present for LocalName, not even an explicit nil
### GetMaturityMonths

`func (o *DataCatalogueItem) GetMaturityMonths() float64`

GetMaturityMonths returns the MaturityMonths field if non-nil, zero value otherwise.

### GetMaturityMonthsOk

`func (o *DataCatalogueItem) GetMaturityMonthsOk() (*float64, bool)`

GetMaturityMonthsOk returns a tuple with the MaturityMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityMonths

`func (o *DataCatalogueItem) SetMaturityMonths(v float64)`

SetMaturityMonths sets MaturityMonths field to given value.

### HasMaturityMonths

`func (o *DataCatalogueItem) HasMaturityMonths() bool`

HasMaturityMonths returns a boolean if a field has been set.

### SetMaturityMonthsNil

`func (o *DataCatalogueItem) SetMaturityMonthsNil(b bool)`

 SetMaturityMonthsNil sets the value for MaturityMonths to be an explicit nil

### UnsetMaturityMonths
`func (o *DataCatalogueItem) UnsetMaturityMonths()`

UnsetMaturityMonths ensures that no value is present for MaturityMonths, not even an explicit nil
### GetYieldType

`func (o *DataCatalogueItem) GetYieldType() string`

GetYieldType returns the YieldType field if non-nil, zero value otherwise.

### GetYieldTypeOk

`func (o *DataCatalogueItem) GetYieldTypeOk() (*string, bool)`

GetYieldTypeOk returns a tuple with the YieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYieldType

`func (o *DataCatalogueItem) SetYieldType(v string)`

SetYieldType sets YieldType field to given value.

### HasYieldType

`func (o *DataCatalogueItem) HasYieldType() bool`

HasYieldType returns a boolean if a field has been set.

### SetYieldTypeNil

`func (o *DataCatalogueItem) SetYieldTypeNil(b bool)`

 SetYieldTypeNil sets the value for YieldType to be an explicit nil

### UnsetYieldType
`func (o *DataCatalogueItem) UnsetYieldType()`

UnsetYieldType ensures that no value is present for YieldType, not even an explicit nil
### GetGeography

`func (o *DataCatalogueItem) GetGeography() string`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *DataCatalogueItem) GetGeographyOk() (*string, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *DataCatalogueItem) SetGeography(v string)`

SetGeography sets Geography field to given value.

### HasGeography

`func (o *DataCatalogueItem) HasGeography() bool`

HasGeography returns a boolean if a field has been set.

### SetGeographyNil

`func (o *DataCatalogueItem) SetGeographyNil(b bool)`

 SetGeographyNil sets the value for Geography to be an explicit nil

### UnsetGeography
`func (o *DataCatalogueItem) UnsetGeography()`

UnsetGeography ensures that no value is present for Geography, not even an explicit nil
### GetPeriodBasis

`func (o *DataCatalogueItem) GetPeriodBasis() string`

GetPeriodBasis returns the PeriodBasis field if non-nil, zero value otherwise.

### GetPeriodBasisOk

`func (o *DataCatalogueItem) GetPeriodBasisOk() (*string, bool)`

GetPeriodBasisOk returns a tuple with the PeriodBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBasis

`func (o *DataCatalogueItem) SetPeriodBasis(v string)`

SetPeriodBasis sets PeriodBasis field to given value.

### HasPeriodBasis

`func (o *DataCatalogueItem) HasPeriodBasis() bool`

HasPeriodBasis returns a boolean if a field has been set.

### SetPeriodBasisNil

`func (o *DataCatalogueItem) SetPeriodBasisNil(b bool)`

 SetPeriodBasisNil sets the value for PeriodBasis to be an explicit nil

### UnsetPeriodBasis
`func (o *DataCatalogueItem) UnsetPeriodBasis()`

UnsetPeriodBasis ensures that no value is present for PeriodBasis, not even an explicit nil
### GetMethodologyId

`func (o *DataCatalogueItem) GetMethodologyId() string`

GetMethodologyId returns the MethodologyId field if non-nil, zero value otherwise.

### GetMethodologyIdOk

`func (o *DataCatalogueItem) GetMethodologyIdOk() (*string, bool)`

GetMethodologyIdOk returns a tuple with the MethodologyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodologyId

`func (o *DataCatalogueItem) SetMethodologyId(v string)`

SetMethodologyId sets MethodologyId field to given value.

### HasMethodologyId

`func (o *DataCatalogueItem) HasMethodologyId() bool`

HasMethodologyId returns a boolean if a field has been set.

### SetMethodologyIdNil

`func (o *DataCatalogueItem) SetMethodologyIdNil(b bool)`

 SetMethodologyIdNil sets the value for MethodologyId to be an explicit nil

### UnsetMethodologyId
`func (o *DataCatalogueItem) UnsetMethodologyId()`

UnsetMethodologyId ensures that no value is present for MethodologyId, not even an explicit nil
### GetFreshnessThresholdDays

`func (o *DataCatalogueItem) GetFreshnessThresholdDays() int64`

GetFreshnessThresholdDays returns the FreshnessThresholdDays field if non-nil, zero value otherwise.

### GetFreshnessThresholdDaysOk

`func (o *DataCatalogueItem) GetFreshnessThresholdDaysOk() (*int64, bool)`

GetFreshnessThresholdDaysOk returns a tuple with the FreshnessThresholdDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshnessThresholdDays

`func (o *DataCatalogueItem) SetFreshnessThresholdDays(v int64)`

SetFreshnessThresholdDays sets FreshnessThresholdDays field to given value.

### HasFreshnessThresholdDays

`func (o *DataCatalogueItem) HasFreshnessThresholdDays() bool`

HasFreshnessThresholdDays returns a boolean if a field has been set.

### SetFreshnessThresholdDaysNil

`func (o *DataCatalogueItem) SetFreshnessThresholdDaysNil(b bool)`

 SetFreshnessThresholdDaysNil sets the value for FreshnessThresholdDays to be an explicit nil

### UnsetFreshnessThresholdDays
`func (o *DataCatalogueItem) UnsetFreshnessThresholdDays()`

UnsetFreshnessThresholdDays ensures that no value is present for FreshnessThresholdDays, not even an explicit nil
### GetReleaseCalendarUrl

`func (o *DataCatalogueItem) GetReleaseCalendarUrl() string`

GetReleaseCalendarUrl returns the ReleaseCalendarUrl field if non-nil, zero value otherwise.

### GetReleaseCalendarUrlOk

`func (o *DataCatalogueItem) GetReleaseCalendarUrlOk() (*string, bool)`

GetReleaseCalendarUrlOk returns a tuple with the ReleaseCalendarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseCalendarUrl

`func (o *DataCatalogueItem) SetReleaseCalendarUrl(v string)`

SetReleaseCalendarUrl sets ReleaseCalendarUrl field to given value.

### HasReleaseCalendarUrl

`func (o *DataCatalogueItem) HasReleaseCalendarUrl() bool`

HasReleaseCalendarUrl returns a boolean if a field has been set.

### SetReleaseCalendarUrlNil

`func (o *DataCatalogueItem) SetReleaseCalendarUrlNil(b bool)`

 SetReleaseCalendarUrlNil sets the value for ReleaseCalendarUrl to be an explicit nil

### UnsetReleaseCalendarUrl
`func (o *DataCatalogueItem) UnsetReleaseCalendarUrl()`

UnsetReleaseCalendarUrl ensures that no value is present for ReleaseCalendarUrl, not even an explicit nil
### GetSourcePermissionStatus

`func (o *DataCatalogueItem) GetSourcePermissionStatus() string`

GetSourcePermissionStatus returns the SourcePermissionStatus field if non-nil, zero value otherwise.

### GetSourcePermissionStatusOk

`func (o *DataCatalogueItem) GetSourcePermissionStatusOk() (*string, bool)`

GetSourcePermissionStatusOk returns a tuple with the SourcePermissionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePermissionStatus

`func (o *DataCatalogueItem) SetSourcePermissionStatus(v string)`

SetSourcePermissionStatus sets SourcePermissionStatus field to given value.

### HasSourcePermissionStatus

`func (o *DataCatalogueItem) HasSourcePermissionStatus() bool`

HasSourcePermissionStatus returns a boolean if a field has been set.

### SetSourcePermissionStatusNil

`func (o *DataCatalogueItem) SetSourcePermissionStatusNil(b bool)`

 SetSourcePermissionStatusNil sets the value for SourcePermissionStatus to be an explicit nil

### UnsetSourcePermissionStatus
`func (o *DataCatalogueItem) UnsetSourcePermissionStatus()`

UnsetSourcePermissionStatus ensures that no value is present for SourcePermissionStatus, not even an explicit nil
### GetSourceTermsUrl

`func (o *DataCatalogueItem) GetSourceTermsUrl() string`

GetSourceTermsUrl returns the SourceTermsUrl field if non-nil, zero value otherwise.

### GetSourceTermsUrlOk

`func (o *DataCatalogueItem) GetSourceTermsUrlOk() (*string, bool)`

GetSourceTermsUrlOk returns a tuple with the SourceTermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTermsUrl

`func (o *DataCatalogueItem) SetSourceTermsUrl(v string)`

SetSourceTermsUrl sets SourceTermsUrl field to given value.

### HasSourceTermsUrl

`func (o *DataCatalogueItem) HasSourceTermsUrl() bool`

HasSourceTermsUrl returns a boolean if a field has been set.

### SetSourceTermsUrlNil

`func (o *DataCatalogueItem) SetSourceTermsUrlNil(b bool)`

 SetSourceTermsUrlNil sets the value for SourceTermsUrl to be an explicit nil

### UnsetSourceTermsUrl
`func (o *DataCatalogueItem) UnsetSourceTermsUrl()`

UnsetSourceTermsUrl ensures that no value is present for SourceTermsUrl, not even an explicit nil
### GetSupportedFrequencyOptions

`func (o *DataCatalogueItem) GetSupportedFrequencyOptions() []string`

GetSupportedFrequencyOptions returns the SupportedFrequencyOptions field if non-nil, zero value otherwise.

### GetSupportedFrequencyOptionsOk

`func (o *DataCatalogueItem) GetSupportedFrequencyOptionsOk() (*[]string, bool)`

GetSupportedFrequencyOptionsOk returns a tuple with the SupportedFrequencyOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFrequencyOptions

`func (o *DataCatalogueItem) SetSupportedFrequencyOptions(v []string)`

SetSupportedFrequencyOptions sets SupportedFrequencyOptions field to given value.

### HasSupportedFrequencyOptions

`func (o *DataCatalogueItem) HasSupportedFrequencyOptions() bool`

HasSupportedFrequencyOptions returns a boolean if a field has been set.

### SetSupportedFrequencyOptionsNil

`func (o *DataCatalogueItem) SetSupportedFrequencyOptionsNil(b bool)`

 SetSupportedFrequencyOptionsNil sets the value for SupportedFrequencyOptions to be an explicit nil

### UnsetSupportedFrequencyOptions
`func (o *DataCatalogueItem) UnsetSupportedFrequencyOptions()`

UnsetSupportedFrequencyOptions ensures that no value is present for SupportedFrequencyOptions, not even an explicit nil
### GetSeriesVariants

`func (o *DataCatalogueItem) GetSeriesVariants() []CatalogueSeriesVariant`

GetSeriesVariants returns the SeriesVariants field if non-nil, zero value otherwise.

### GetSeriesVariantsOk

`func (o *DataCatalogueItem) GetSeriesVariantsOk() (*[]CatalogueSeriesVariant, bool)`

GetSeriesVariantsOk returns a tuple with the SeriesVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesVariants

`func (o *DataCatalogueItem) SetSeriesVariants(v []CatalogueSeriesVariant)`

SetSeriesVariants sets SeriesVariants field to given value.

### HasSeriesVariants

`func (o *DataCatalogueItem) HasSeriesVariants() bool`

HasSeriesVariants returns a boolean if a field has been set.

### SetSeriesVariantsNil

`func (o *DataCatalogueItem) SetSeriesVariantsNil(b bool)`

 SetSeriesVariantsNil sets the value for SeriesVariants to be an explicit nil

### UnsetSeriesVariants
`func (o *DataCatalogueItem) UnsetSeriesVariants()`

UnsetSeriesVariants ensures that no value is present for SeriesVariants, not even an explicit nil
### GetCoverage

`func (o *DataCatalogueItem) GetCoverage() CatalogueCoverage`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *DataCatalogueItem) GetCoverageOk() (*CatalogueCoverage, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *DataCatalogueItem) SetCoverage(v CatalogueCoverage)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *DataCatalogueItem) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### SetCoverageNil

`func (o *DataCatalogueItem) SetCoverageNil(b bool)`

 SetCoverageNil sets the value for Coverage to be an explicit nil

### UnsetCoverage
`func (o *DataCatalogueItem) UnsetCoverage()`

UnsetCoverage ensures that no value is present for Coverage, not even an explicit nil
### GetSupportedOptions

`func (o *DataCatalogueItem) GetSupportedOptions() CatalogueSupportedOptions`

GetSupportedOptions returns the SupportedOptions field if non-nil, zero value otherwise.

### GetSupportedOptionsOk

`func (o *DataCatalogueItem) GetSupportedOptionsOk() (*CatalogueSupportedOptions, bool)`

GetSupportedOptionsOk returns a tuple with the SupportedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOptions

`func (o *DataCatalogueItem) SetSupportedOptions(v CatalogueSupportedOptions)`

SetSupportedOptions sets SupportedOptions field to given value.

### HasSupportedOptions

`func (o *DataCatalogueItem) HasSupportedOptions() bool`

HasSupportedOptions returns a boolean if a field has been set.

### SetSupportedOptionsNil

`func (o *DataCatalogueItem) SetSupportedOptionsNil(b bool)`

 SetSupportedOptionsNil sets the value for SupportedOptions to be an explicit nil

### UnsetSupportedOptions
`func (o *DataCatalogueItem) UnsetSupportedOptions()`

UnsetSupportedOptions ensures that no value is present for SupportedOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


