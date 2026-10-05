# SelectedSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeriesId** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**SourceSeriesId** | Pointer to **NullableString** |  | [optional] 
**Seasonality** | Pointer to **NullableString** |  | [optional] 
**Basis** | Pointer to **NullableString** |  | [optional] 
**Annualization** | Pointer to **NullableString** |  | [optional] 
**PeriodAggregation** | Pointer to **NullableString** |  | [optional] 
**Frequency** | Pointer to **NullableString** |  | [optional] 
**Revisions** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSelectedSeries

`func NewSelectedSeries() *SelectedSeries`

NewSelectedSeries instantiates a new SelectedSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectedSeriesWithDefaults

`func NewSelectedSeriesWithDefaults() *SelectedSeries`

NewSelectedSeriesWithDefaults instantiates a new SelectedSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeriesId

`func (o *SelectedSeries) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *SelectedSeries) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *SelectedSeries) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *SelectedSeries) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### SetSeriesIdNil

`func (o *SelectedSeries) SetSeriesIdNil(b bool)`

 SetSeriesIdNil sets the value for SeriesId to be an explicit nil

### UnsetSeriesId
`func (o *SelectedSeries) UnsetSeriesId()`

UnsetSeriesId ensures that no value is present for SeriesId, not even an explicit nil
### GetLabel

`func (o *SelectedSeries) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SelectedSeries) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SelectedSeries) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SelectedSeries) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *SelectedSeries) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *SelectedSeries) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetSourceSeriesId

`func (o *SelectedSeries) GetSourceSeriesId() string`

GetSourceSeriesId returns the SourceSeriesId field if non-nil, zero value otherwise.

### GetSourceSeriesIdOk

`func (o *SelectedSeries) GetSourceSeriesIdOk() (*string, bool)`

GetSourceSeriesIdOk returns a tuple with the SourceSeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSeriesId

`func (o *SelectedSeries) SetSourceSeriesId(v string)`

SetSourceSeriesId sets SourceSeriesId field to given value.

### HasSourceSeriesId

`func (o *SelectedSeries) HasSourceSeriesId() bool`

HasSourceSeriesId returns a boolean if a field has been set.

### SetSourceSeriesIdNil

`func (o *SelectedSeries) SetSourceSeriesIdNil(b bool)`

 SetSourceSeriesIdNil sets the value for SourceSeriesId to be an explicit nil

### UnsetSourceSeriesId
`func (o *SelectedSeries) UnsetSourceSeriesId()`

UnsetSourceSeriesId ensures that no value is present for SourceSeriesId, not even an explicit nil
### GetSeasonality

`func (o *SelectedSeries) GetSeasonality() string`

GetSeasonality returns the Seasonality field if non-nil, zero value otherwise.

### GetSeasonalityOk

`func (o *SelectedSeries) GetSeasonalityOk() (*string, bool)`

GetSeasonalityOk returns a tuple with the Seasonality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonality

`func (o *SelectedSeries) SetSeasonality(v string)`

SetSeasonality sets Seasonality field to given value.

### HasSeasonality

`func (o *SelectedSeries) HasSeasonality() bool`

HasSeasonality returns a boolean if a field has been set.

### SetSeasonalityNil

`func (o *SelectedSeries) SetSeasonalityNil(b bool)`

 SetSeasonalityNil sets the value for Seasonality to be an explicit nil

### UnsetSeasonality
`func (o *SelectedSeries) UnsetSeasonality()`

UnsetSeasonality ensures that no value is present for Seasonality, not even an explicit nil
### GetBasis

`func (o *SelectedSeries) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *SelectedSeries) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *SelectedSeries) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *SelectedSeries) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### SetBasisNil

`func (o *SelectedSeries) SetBasisNil(b bool)`

 SetBasisNil sets the value for Basis to be an explicit nil

### UnsetBasis
`func (o *SelectedSeries) UnsetBasis()`

UnsetBasis ensures that no value is present for Basis, not even an explicit nil
### GetAnnualization

`func (o *SelectedSeries) GetAnnualization() string`

GetAnnualization returns the Annualization field if non-nil, zero value otherwise.

### GetAnnualizationOk

`func (o *SelectedSeries) GetAnnualizationOk() (*string, bool)`

GetAnnualizationOk returns a tuple with the Annualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualization

`func (o *SelectedSeries) SetAnnualization(v string)`

SetAnnualization sets Annualization field to given value.

### HasAnnualization

`func (o *SelectedSeries) HasAnnualization() bool`

HasAnnualization returns a boolean if a field has been set.

### SetAnnualizationNil

`func (o *SelectedSeries) SetAnnualizationNil(b bool)`

 SetAnnualizationNil sets the value for Annualization to be an explicit nil

### UnsetAnnualization
`func (o *SelectedSeries) UnsetAnnualization()`

UnsetAnnualization ensures that no value is present for Annualization, not even an explicit nil
### GetPeriodAggregation

`func (o *SelectedSeries) GetPeriodAggregation() string`

GetPeriodAggregation returns the PeriodAggregation field if non-nil, zero value otherwise.

### GetPeriodAggregationOk

`func (o *SelectedSeries) GetPeriodAggregationOk() (*string, bool)`

GetPeriodAggregationOk returns a tuple with the PeriodAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodAggregation

`func (o *SelectedSeries) SetPeriodAggregation(v string)`

SetPeriodAggregation sets PeriodAggregation field to given value.

### HasPeriodAggregation

`func (o *SelectedSeries) HasPeriodAggregation() bool`

HasPeriodAggregation returns a boolean if a field has been set.

### SetPeriodAggregationNil

`func (o *SelectedSeries) SetPeriodAggregationNil(b bool)`

 SetPeriodAggregationNil sets the value for PeriodAggregation to be an explicit nil

### UnsetPeriodAggregation
`func (o *SelectedSeries) UnsetPeriodAggregation()`

UnsetPeriodAggregation ensures that no value is present for PeriodAggregation, not even an explicit nil
### GetFrequency

`func (o *SelectedSeries) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *SelectedSeries) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *SelectedSeries) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *SelectedSeries) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *SelectedSeries) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *SelectedSeries) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetRevisions

`func (o *SelectedSeries) GetRevisions() string`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *SelectedSeries) GetRevisionsOk() (*string, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *SelectedSeries) SetRevisions(v string)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *SelectedSeries) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### SetRevisionsNil

`func (o *SelectedSeries) SetRevisionsNil(b bool)`

 SetRevisionsNil sets the value for Revisions to be an explicit nil

### UnsetRevisions
`func (o *SelectedSeries) UnsetRevisions()`

UnsetRevisions ensures that no value is present for Revisions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


