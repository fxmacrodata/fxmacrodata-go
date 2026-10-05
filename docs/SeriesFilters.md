# SeriesFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seasonality** | Pointer to **NullableString** |  | [optional] 
**Basis** | Pointer to **NullableString** |  | [optional] 
**Annualization** | Pointer to **NullableString** |  | [optional] 
**PeriodAggregation** | Pointer to **NullableString** |  | [optional] 
**Frequency** | Pointer to **NullableString** |  | [optional] 
**Revisions** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSeriesFilters

`func NewSeriesFilters() *SeriesFilters`

NewSeriesFilters instantiates a new SeriesFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesFiltersWithDefaults

`func NewSeriesFiltersWithDefaults() *SeriesFilters`

NewSeriesFiltersWithDefaults instantiates a new SeriesFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeasonality

`func (o *SeriesFilters) GetSeasonality() string`

GetSeasonality returns the Seasonality field if non-nil, zero value otherwise.

### GetSeasonalityOk

`func (o *SeriesFilters) GetSeasonalityOk() (*string, bool)`

GetSeasonalityOk returns a tuple with the Seasonality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonality

`func (o *SeriesFilters) SetSeasonality(v string)`

SetSeasonality sets Seasonality field to given value.

### HasSeasonality

`func (o *SeriesFilters) HasSeasonality() bool`

HasSeasonality returns a boolean if a field has been set.

### SetSeasonalityNil

`func (o *SeriesFilters) SetSeasonalityNil(b bool)`

 SetSeasonalityNil sets the value for Seasonality to be an explicit nil

### UnsetSeasonality
`func (o *SeriesFilters) UnsetSeasonality()`

UnsetSeasonality ensures that no value is present for Seasonality, not even an explicit nil
### GetBasis

`func (o *SeriesFilters) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *SeriesFilters) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *SeriesFilters) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *SeriesFilters) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### SetBasisNil

`func (o *SeriesFilters) SetBasisNil(b bool)`

 SetBasisNil sets the value for Basis to be an explicit nil

### UnsetBasis
`func (o *SeriesFilters) UnsetBasis()`

UnsetBasis ensures that no value is present for Basis, not even an explicit nil
### GetAnnualization

`func (o *SeriesFilters) GetAnnualization() string`

GetAnnualization returns the Annualization field if non-nil, zero value otherwise.

### GetAnnualizationOk

`func (o *SeriesFilters) GetAnnualizationOk() (*string, bool)`

GetAnnualizationOk returns a tuple with the Annualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualization

`func (o *SeriesFilters) SetAnnualization(v string)`

SetAnnualization sets Annualization field to given value.

### HasAnnualization

`func (o *SeriesFilters) HasAnnualization() bool`

HasAnnualization returns a boolean if a field has been set.

### SetAnnualizationNil

`func (o *SeriesFilters) SetAnnualizationNil(b bool)`

 SetAnnualizationNil sets the value for Annualization to be an explicit nil

### UnsetAnnualization
`func (o *SeriesFilters) UnsetAnnualization()`

UnsetAnnualization ensures that no value is present for Annualization, not even an explicit nil
### GetPeriodAggregation

`func (o *SeriesFilters) GetPeriodAggregation() string`

GetPeriodAggregation returns the PeriodAggregation field if non-nil, zero value otherwise.

### GetPeriodAggregationOk

`func (o *SeriesFilters) GetPeriodAggregationOk() (*string, bool)`

GetPeriodAggregationOk returns a tuple with the PeriodAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodAggregation

`func (o *SeriesFilters) SetPeriodAggregation(v string)`

SetPeriodAggregation sets PeriodAggregation field to given value.

### HasPeriodAggregation

`func (o *SeriesFilters) HasPeriodAggregation() bool`

HasPeriodAggregation returns a boolean if a field has been set.

### SetPeriodAggregationNil

`func (o *SeriesFilters) SetPeriodAggregationNil(b bool)`

 SetPeriodAggregationNil sets the value for PeriodAggregation to be an explicit nil

### UnsetPeriodAggregation
`func (o *SeriesFilters) UnsetPeriodAggregation()`

UnsetPeriodAggregation ensures that no value is present for PeriodAggregation, not even an explicit nil
### GetFrequency

`func (o *SeriesFilters) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *SeriesFilters) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *SeriesFilters) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *SeriesFilters) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *SeriesFilters) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *SeriesFilters) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetRevisions

`func (o *SeriesFilters) GetRevisions() string`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *SeriesFilters) GetRevisionsOk() (*string, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *SeriesFilters) SetRevisions(v string)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *SeriesFilters) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### SetRevisionsNil

`func (o *SeriesFilters) SetRevisionsNil(b bool)`

 SetRevisionsNil sets the value for Revisions to be an explicit nil

### UnsetRevisions
`func (o *SeriesFilters) UnsetRevisions()`

UnsetRevisions ensures that no value is present for Revisions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


