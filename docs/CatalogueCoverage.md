# CatalogueCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **NullableBool** |  | [optional] 
**RequiresApiKey** | Pointer to **NullableBool** |  | [optional] 
**EarliestAvailableDate** | Pointer to **NullableString** |  | [optional] 
**LatestAvailableDate** | Pointer to **NullableString** |  | [optional] 
**LatestReleaseDate** | Pointer to **NullableString** |  | [optional] 
**RowCount** | Pointer to **NullableInt64** |  | [optional] 
**ValueAvailable** | Pointer to **NullableBool** |  | [optional] 
**HasRecentData** | Pointer to **NullableBool** |  | [optional] 
**RecentObservationCount** | Pointer to **NullableInt64** |  | [optional] 
**CoverageQuality** | Pointer to **NullableString** |  | [optional] 
**HistoryCoverageQuality** | Pointer to **NullableString** |  | [optional] 
**FreshnessQuality** | Pointer to **NullableString** |  | [optional] 
**UsableForContext** | Pointer to **NullableBool** |  | [optional] 
**UsableForSignal** | Pointer to **NullableBool** |  | [optional] 
**DataLagDays** | Pointer to **NullableInt64** |  | [optional] 
**StaleAfterDays** | Pointer to **NullableInt64** |  | [optional] 
**HasYoyTransform** | Pointer to **NullableBool** |  | [optional] 
**HasQoqTransform** | Pointer to **NullableBool** |  | [optional] 
**HasMomTransform** | Pointer to **NullableBool** |  | [optional] 
**SourceFreshness** | Pointer to **NullableString** |  | [optional] 
**SourceMonitoringComplete** | Pointer to **NullableBool** |  | [optional] 
**KnownSourceGap** | Pointer to **NullableBool** |  | [optional] 
**SourceCheckedAt** | Pointer to **NullableInt64** |  | [optional] 
**SourceCheckExpiresAt** | Pointer to **NullableInt64** |  | [optional] 
**LatestPublishedPeriod** | Pointer to **NullableString** |  | [optional] 
**IsStale** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCatalogueCoverage

`func NewCatalogueCoverage() *CatalogueCoverage`

NewCatalogueCoverage instantiates a new CatalogueCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogueCoverageWithDefaults

`func NewCatalogueCoverageWithDefaults() *CatalogueCoverage`

NewCatalogueCoverageWithDefaults instantiates a new CatalogueCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *CatalogueCoverage) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CatalogueCoverage) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CatalogueCoverage) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CatalogueCoverage) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *CatalogueCoverage) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *CatalogueCoverage) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetRequiresApiKey

`func (o *CatalogueCoverage) GetRequiresApiKey() bool`

GetRequiresApiKey returns the RequiresApiKey field if non-nil, zero value otherwise.

### GetRequiresApiKeyOk

`func (o *CatalogueCoverage) GetRequiresApiKeyOk() (*bool, bool)`

GetRequiresApiKeyOk returns a tuple with the RequiresApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApiKey

`func (o *CatalogueCoverage) SetRequiresApiKey(v bool)`

SetRequiresApiKey sets RequiresApiKey field to given value.

### HasRequiresApiKey

`func (o *CatalogueCoverage) HasRequiresApiKey() bool`

HasRequiresApiKey returns a boolean if a field has been set.

### SetRequiresApiKeyNil

`func (o *CatalogueCoverage) SetRequiresApiKeyNil(b bool)`

 SetRequiresApiKeyNil sets the value for RequiresApiKey to be an explicit nil

### UnsetRequiresApiKey
`func (o *CatalogueCoverage) UnsetRequiresApiKey()`

UnsetRequiresApiKey ensures that no value is present for RequiresApiKey, not even an explicit nil
### GetEarliestAvailableDate

`func (o *CatalogueCoverage) GetEarliestAvailableDate() string`

GetEarliestAvailableDate returns the EarliestAvailableDate field if non-nil, zero value otherwise.

### GetEarliestAvailableDateOk

`func (o *CatalogueCoverage) GetEarliestAvailableDateOk() (*string, bool)`

GetEarliestAvailableDateOk returns a tuple with the EarliestAvailableDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestAvailableDate

`func (o *CatalogueCoverage) SetEarliestAvailableDate(v string)`

SetEarliestAvailableDate sets EarliestAvailableDate field to given value.

### HasEarliestAvailableDate

`func (o *CatalogueCoverage) HasEarliestAvailableDate() bool`

HasEarliestAvailableDate returns a boolean if a field has been set.

### SetEarliestAvailableDateNil

`func (o *CatalogueCoverage) SetEarliestAvailableDateNil(b bool)`

 SetEarliestAvailableDateNil sets the value for EarliestAvailableDate to be an explicit nil

### UnsetEarliestAvailableDate
`func (o *CatalogueCoverage) UnsetEarliestAvailableDate()`

UnsetEarliestAvailableDate ensures that no value is present for EarliestAvailableDate, not even an explicit nil
### GetLatestAvailableDate

`func (o *CatalogueCoverage) GetLatestAvailableDate() string`

GetLatestAvailableDate returns the LatestAvailableDate field if non-nil, zero value otherwise.

### GetLatestAvailableDateOk

`func (o *CatalogueCoverage) GetLatestAvailableDateOk() (*string, bool)`

GetLatestAvailableDateOk returns a tuple with the LatestAvailableDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAvailableDate

`func (o *CatalogueCoverage) SetLatestAvailableDate(v string)`

SetLatestAvailableDate sets LatestAvailableDate field to given value.

### HasLatestAvailableDate

`func (o *CatalogueCoverage) HasLatestAvailableDate() bool`

HasLatestAvailableDate returns a boolean if a field has been set.

### SetLatestAvailableDateNil

`func (o *CatalogueCoverage) SetLatestAvailableDateNil(b bool)`

 SetLatestAvailableDateNil sets the value for LatestAvailableDate to be an explicit nil

### UnsetLatestAvailableDate
`func (o *CatalogueCoverage) UnsetLatestAvailableDate()`

UnsetLatestAvailableDate ensures that no value is present for LatestAvailableDate, not even an explicit nil
### GetLatestReleaseDate

`func (o *CatalogueCoverage) GetLatestReleaseDate() string`

GetLatestReleaseDate returns the LatestReleaseDate field if non-nil, zero value otherwise.

### GetLatestReleaseDateOk

`func (o *CatalogueCoverage) GetLatestReleaseDateOk() (*string, bool)`

GetLatestReleaseDateOk returns a tuple with the LatestReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReleaseDate

`func (o *CatalogueCoverage) SetLatestReleaseDate(v string)`

SetLatestReleaseDate sets LatestReleaseDate field to given value.

### HasLatestReleaseDate

`func (o *CatalogueCoverage) HasLatestReleaseDate() bool`

HasLatestReleaseDate returns a boolean if a field has been set.

### SetLatestReleaseDateNil

`func (o *CatalogueCoverage) SetLatestReleaseDateNil(b bool)`

 SetLatestReleaseDateNil sets the value for LatestReleaseDate to be an explicit nil

### UnsetLatestReleaseDate
`func (o *CatalogueCoverage) UnsetLatestReleaseDate()`

UnsetLatestReleaseDate ensures that no value is present for LatestReleaseDate, not even an explicit nil
### GetRowCount

`func (o *CatalogueCoverage) GetRowCount() int64`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *CatalogueCoverage) GetRowCountOk() (*int64, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *CatalogueCoverage) SetRowCount(v int64)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *CatalogueCoverage) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### SetRowCountNil

`func (o *CatalogueCoverage) SetRowCountNil(b bool)`

 SetRowCountNil sets the value for RowCount to be an explicit nil

### UnsetRowCount
`func (o *CatalogueCoverage) UnsetRowCount()`

UnsetRowCount ensures that no value is present for RowCount, not even an explicit nil
### GetValueAvailable

`func (o *CatalogueCoverage) GetValueAvailable() bool`

GetValueAvailable returns the ValueAvailable field if non-nil, zero value otherwise.

### GetValueAvailableOk

`func (o *CatalogueCoverage) GetValueAvailableOk() (*bool, bool)`

GetValueAvailableOk returns a tuple with the ValueAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAvailable

`func (o *CatalogueCoverage) SetValueAvailable(v bool)`

SetValueAvailable sets ValueAvailable field to given value.

### HasValueAvailable

`func (o *CatalogueCoverage) HasValueAvailable() bool`

HasValueAvailable returns a boolean if a field has been set.

### SetValueAvailableNil

`func (o *CatalogueCoverage) SetValueAvailableNil(b bool)`

 SetValueAvailableNil sets the value for ValueAvailable to be an explicit nil

### UnsetValueAvailable
`func (o *CatalogueCoverage) UnsetValueAvailable()`

UnsetValueAvailable ensures that no value is present for ValueAvailable, not even an explicit nil
### GetHasRecentData

`func (o *CatalogueCoverage) GetHasRecentData() bool`

GetHasRecentData returns the HasRecentData field if non-nil, zero value otherwise.

### GetHasRecentDataOk

`func (o *CatalogueCoverage) GetHasRecentDataOk() (*bool, bool)`

GetHasRecentDataOk returns a tuple with the HasRecentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRecentData

`func (o *CatalogueCoverage) SetHasRecentData(v bool)`

SetHasRecentData sets HasRecentData field to given value.

### HasHasRecentData

`func (o *CatalogueCoverage) HasHasRecentData() bool`

HasHasRecentData returns a boolean if a field has been set.

### SetHasRecentDataNil

`func (o *CatalogueCoverage) SetHasRecentDataNil(b bool)`

 SetHasRecentDataNil sets the value for HasRecentData to be an explicit nil

### UnsetHasRecentData
`func (o *CatalogueCoverage) UnsetHasRecentData()`

UnsetHasRecentData ensures that no value is present for HasRecentData, not even an explicit nil
### GetRecentObservationCount

`func (o *CatalogueCoverage) GetRecentObservationCount() int64`

GetRecentObservationCount returns the RecentObservationCount field if non-nil, zero value otherwise.

### GetRecentObservationCountOk

`func (o *CatalogueCoverage) GetRecentObservationCountOk() (*int64, bool)`

GetRecentObservationCountOk returns a tuple with the RecentObservationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentObservationCount

`func (o *CatalogueCoverage) SetRecentObservationCount(v int64)`

SetRecentObservationCount sets RecentObservationCount field to given value.

### HasRecentObservationCount

`func (o *CatalogueCoverage) HasRecentObservationCount() bool`

HasRecentObservationCount returns a boolean if a field has been set.

### SetRecentObservationCountNil

`func (o *CatalogueCoverage) SetRecentObservationCountNil(b bool)`

 SetRecentObservationCountNil sets the value for RecentObservationCount to be an explicit nil

### UnsetRecentObservationCount
`func (o *CatalogueCoverage) UnsetRecentObservationCount()`

UnsetRecentObservationCount ensures that no value is present for RecentObservationCount, not even an explicit nil
### GetCoverageQuality

`func (o *CatalogueCoverage) GetCoverageQuality() string`

GetCoverageQuality returns the CoverageQuality field if non-nil, zero value otherwise.

### GetCoverageQualityOk

`func (o *CatalogueCoverage) GetCoverageQualityOk() (*string, bool)`

GetCoverageQualityOk returns a tuple with the CoverageQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageQuality

`func (o *CatalogueCoverage) SetCoverageQuality(v string)`

SetCoverageQuality sets CoverageQuality field to given value.

### HasCoverageQuality

`func (o *CatalogueCoverage) HasCoverageQuality() bool`

HasCoverageQuality returns a boolean if a field has been set.

### SetCoverageQualityNil

`func (o *CatalogueCoverage) SetCoverageQualityNil(b bool)`

 SetCoverageQualityNil sets the value for CoverageQuality to be an explicit nil

### UnsetCoverageQuality
`func (o *CatalogueCoverage) UnsetCoverageQuality()`

UnsetCoverageQuality ensures that no value is present for CoverageQuality, not even an explicit nil
### GetHistoryCoverageQuality

`func (o *CatalogueCoverage) GetHistoryCoverageQuality() string`

GetHistoryCoverageQuality returns the HistoryCoverageQuality field if non-nil, zero value otherwise.

### GetHistoryCoverageQualityOk

`func (o *CatalogueCoverage) GetHistoryCoverageQualityOk() (*string, bool)`

GetHistoryCoverageQualityOk returns a tuple with the HistoryCoverageQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryCoverageQuality

`func (o *CatalogueCoverage) SetHistoryCoverageQuality(v string)`

SetHistoryCoverageQuality sets HistoryCoverageQuality field to given value.

### HasHistoryCoverageQuality

`func (o *CatalogueCoverage) HasHistoryCoverageQuality() bool`

HasHistoryCoverageQuality returns a boolean if a field has been set.

### SetHistoryCoverageQualityNil

`func (o *CatalogueCoverage) SetHistoryCoverageQualityNil(b bool)`

 SetHistoryCoverageQualityNil sets the value for HistoryCoverageQuality to be an explicit nil

### UnsetHistoryCoverageQuality
`func (o *CatalogueCoverage) UnsetHistoryCoverageQuality()`

UnsetHistoryCoverageQuality ensures that no value is present for HistoryCoverageQuality, not even an explicit nil
### GetFreshnessQuality

`func (o *CatalogueCoverage) GetFreshnessQuality() string`

GetFreshnessQuality returns the FreshnessQuality field if non-nil, zero value otherwise.

### GetFreshnessQualityOk

`func (o *CatalogueCoverage) GetFreshnessQualityOk() (*string, bool)`

GetFreshnessQualityOk returns a tuple with the FreshnessQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshnessQuality

`func (o *CatalogueCoverage) SetFreshnessQuality(v string)`

SetFreshnessQuality sets FreshnessQuality field to given value.

### HasFreshnessQuality

`func (o *CatalogueCoverage) HasFreshnessQuality() bool`

HasFreshnessQuality returns a boolean if a field has been set.

### SetFreshnessQualityNil

`func (o *CatalogueCoverage) SetFreshnessQualityNil(b bool)`

 SetFreshnessQualityNil sets the value for FreshnessQuality to be an explicit nil

### UnsetFreshnessQuality
`func (o *CatalogueCoverage) UnsetFreshnessQuality()`

UnsetFreshnessQuality ensures that no value is present for FreshnessQuality, not even an explicit nil
### GetUsableForContext

`func (o *CatalogueCoverage) GetUsableForContext() bool`

GetUsableForContext returns the UsableForContext field if non-nil, zero value otherwise.

### GetUsableForContextOk

`func (o *CatalogueCoverage) GetUsableForContextOk() (*bool, bool)`

GetUsableForContextOk returns a tuple with the UsableForContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsableForContext

`func (o *CatalogueCoverage) SetUsableForContext(v bool)`

SetUsableForContext sets UsableForContext field to given value.

### HasUsableForContext

`func (o *CatalogueCoverage) HasUsableForContext() bool`

HasUsableForContext returns a boolean if a field has been set.

### SetUsableForContextNil

`func (o *CatalogueCoverage) SetUsableForContextNil(b bool)`

 SetUsableForContextNil sets the value for UsableForContext to be an explicit nil

### UnsetUsableForContext
`func (o *CatalogueCoverage) UnsetUsableForContext()`

UnsetUsableForContext ensures that no value is present for UsableForContext, not even an explicit nil
### GetUsableForSignal

`func (o *CatalogueCoverage) GetUsableForSignal() bool`

GetUsableForSignal returns the UsableForSignal field if non-nil, zero value otherwise.

### GetUsableForSignalOk

`func (o *CatalogueCoverage) GetUsableForSignalOk() (*bool, bool)`

GetUsableForSignalOk returns a tuple with the UsableForSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsableForSignal

`func (o *CatalogueCoverage) SetUsableForSignal(v bool)`

SetUsableForSignal sets UsableForSignal field to given value.

### HasUsableForSignal

`func (o *CatalogueCoverage) HasUsableForSignal() bool`

HasUsableForSignal returns a boolean if a field has been set.

### SetUsableForSignalNil

`func (o *CatalogueCoverage) SetUsableForSignalNil(b bool)`

 SetUsableForSignalNil sets the value for UsableForSignal to be an explicit nil

### UnsetUsableForSignal
`func (o *CatalogueCoverage) UnsetUsableForSignal()`

UnsetUsableForSignal ensures that no value is present for UsableForSignal, not even an explicit nil
### GetDataLagDays

`func (o *CatalogueCoverage) GetDataLagDays() int64`

GetDataLagDays returns the DataLagDays field if non-nil, zero value otherwise.

### GetDataLagDaysOk

`func (o *CatalogueCoverage) GetDataLagDaysOk() (*int64, bool)`

GetDataLagDaysOk returns a tuple with the DataLagDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLagDays

`func (o *CatalogueCoverage) SetDataLagDays(v int64)`

SetDataLagDays sets DataLagDays field to given value.

### HasDataLagDays

`func (o *CatalogueCoverage) HasDataLagDays() bool`

HasDataLagDays returns a boolean if a field has been set.

### SetDataLagDaysNil

`func (o *CatalogueCoverage) SetDataLagDaysNil(b bool)`

 SetDataLagDaysNil sets the value for DataLagDays to be an explicit nil

### UnsetDataLagDays
`func (o *CatalogueCoverage) UnsetDataLagDays()`

UnsetDataLagDays ensures that no value is present for DataLagDays, not even an explicit nil
### GetStaleAfterDays

`func (o *CatalogueCoverage) GetStaleAfterDays() int64`

GetStaleAfterDays returns the StaleAfterDays field if non-nil, zero value otherwise.

### GetStaleAfterDaysOk

`func (o *CatalogueCoverage) GetStaleAfterDaysOk() (*int64, bool)`

GetStaleAfterDaysOk returns a tuple with the StaleAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleAfterDays

`func (o *CatalogueCoverage) SetStaleAfterDays(v int64)`

SetStaleAfterDays sets StaleAfterDays field to given value.

### HasStaleAfterDays

`func (o *CatalogueCoverage) HasStaleAfterDays() bool`

HasStaleAfterDays returns a boolean if a field has been set.

### SetStaleAfterDaysNil

`func (o *CatalogueCoverage) SetStaleAfterDaysNil(b bool)`

 SetStaleAfterDaysNil sets the value for StaleAfterDays to be an explicit nil

### UnsetStaleAfterDays
`func (o *CatalogueCoverage) UnsetStaleAfterDays()`

UnsetStaleAfterDays ensures that no value is present for StaleAfterDays, not even an explicit nil
### GetHasYoyTransform

`func (o *CatalogueCoverage) GetHasYoyTransform() bool`

GetHasYoyTransform returns the HasYoyTransform field if non-nil, zero value otherwise.

### GetHasYoyTransformOk

`func (o *CatalogueCoverage) GetHasYoyTransformOk() (*bool, bool)`

GetHasYoyTransformOk returns a tuple with the HasYoyTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasYoyTransform

`func (o *CatalogueCoverage) SetHasYoyTransform(v bool)`

SetHasYoyTransform sets HasYoyTransform field to given value.

### HasHasYoyTransform

`func (o *CatalogueCoverage) HasHasYoyTransform() bool`

HasHasYoyTransform returns a boolean if a field has been set.

### SetHasYoyTransformNil

`func (o *CatalogueCoverage) SetHasYoyTransformNil(b bool)`

 SetHasYoyTransformNil sets the value for HasYoyTransform to be an explicit nil

### UnsetHasYoyTransform
`func (o *CatalogueCoverage) UnsetHasYoyTransform()`

UnsetHasYoyTransform ensures that no value is present for HasYoyTransform, not even an explicit nil
### GetHasQoqTransform

`func (o *CatalogueCoverage) GetHasQoqTransform() bool`

GetHasQoqTransform returns the HasQoqTransform field if non-nil, zero value otherwise.

### GetHasQoqTransformOk

`func (o *CatalogueCoverage) GetHasQoqTransformOk() (*bool, bool)`

GetHasQoqTransformOk returns a tuple with the HasQoqTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasQoqTransform

`func (o *CatalogueCoverage) SetHasQoqTransform(v bool)`

SetHasQoqTransform sets HasQoqTransform field to given value.

### HasHasQoqTransform

`func (o *CatalogueCoverage) HasHasQoqTransform() bool`

HasHasQoqTransform returns a boolean if a field has been set.

### SetHasQoqTransformNil

`func (o *CatalogueCoverage) SetHasQoqTransformNil(b bool)`

 SetHasQoqTransformNil sets the value for HasQoqTransform to be an explicit nil

### UnsetHasQoqTransform
`func (o *CatalogueCoverage) UnsetHasQoqTransform()`

UnsetHasQoqTransform ensures that no value is present for HasQoqTransform, not even an explicit nil
### GetHasMomTransform

`func (o *CatalogueCoverage) GetHasMomTransform() bool`

GetHasMomTransform returns the HasMomTransform field if non-nil, zero value otherwise.

### GetHasMomTransformOk

`func (o *CatalogueCoverage) GetHasMomTransformOk() (*bool, bool)`

GetHasMomTransformOk returns a tuple with the HasMomTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMomTransform

`func (o *CatalogueCoverage) SetHasMomTransform(v bool)`

SetHasMomTransform sets HasMomTransform field to given value.

### HasHasMomTransform

`func (o *CatalogueCoverage) HasHasMomTransform() bool`

HasHasMomTransform returns a boolean if a field has been set.

### SetHasMomTransformNil

`func (o *CatalogueCoverage) SetHasMomTransformNil(b bool)`

 SetHasMomTransformNil sets the value for HasMomTransform to be an explicit nil

### UnsetHasMomTransform
`func (o *CatalogueCoverage) UnsetHasMomTransform()`

UnsetHasMomTransform ensures that no value is present for HasMomTransform, not even an explicit nil
### GetSourceFreshness

`func (o *CatalogueCoverage) GetSourceFreshness() string`

GetSourceFreshness returns the SourceFreshness field if non-nil, zero value otherwise.

### GetSourceFreshnessOk

`func (o *CatalogueCoverage) GetSourceFreshnessOk() (*string, bool)`

GetSourceFreshnessOk returns a tuple with the SourceFreshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFreshness

`func (o *CatalogueCoverage) SetSourceFreshness(v string)`

SetSourceFreshness sets SourceFreshness field to given value.

### HasSourceFreshness

`func (o *CatalogueCoverage) HasSourceFreshness() bool`

HasSourceFreshness returns a boolean if a field has been set.

### SetSourceFreshnessNil

`func (o *CatalogueCoverage) SetSourceFreshnessNil(b bool)`

 SetSourceFreshnessNil sets the value for SourceFreshness to be an explicit nil

### UnsetSourceFreshness
`func (o *CatalogueCoverage) UnsetSourceFreshness()`

UnsetSourceFreshness ensures that no value is present for SourceFreshness, not even an explicit nil
### GetSourceMonitoringComplete

`func (o *CatalogueCoverage) GetSourceMonitoringComplete() bool`

GetSourceMonitoringComplete returns the SourceMonitoringComplete field if non-nil, zero value otherwise.

### GetSourceMonitoringCompleteOk

`func (o *CatalogueCoverage) GetSourceMonitoringCompleteOk() (*bool, bool)`

GetSourceMonitoringCompleteOk returns a tuple with the SourceMonitoringComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMonitoringComplete

`func (o *CatalogueCoverage) SetSourceMonitoringComplete(v bool)`

SetSourceMonitoringComplete sets SourceMonitoringComplete field to given value.

### HasSourceMonitoringComplete

`func (o *CatalogueCoverage) HasSourceMonitoringComplete() bool`

HasSourceMonitoringComplete returns a boolean if a field has been set.

### SetSourceMonitoringCompleteNil

`func (o *CatalogueCoverage) SetSourceMonitoringCompleteNil(b bool)`

 SetSourceMonitoringCompleteNil sets the value for SourceMonitoringComplete to be an explicit nil

### UnsetSourceMonitoringComplete
`func (o *CatalogueCoverage) UnsetSourceMonitoringComplete()`

UnsetSourceMonitoringComplete ensures that no value is present for SourceMonitoringComplete, not even an explicit nil
### GetKnownSourceGap

`func (o *CatalogueCoverage) GetKnownSourceGap() bool`

GetKnownSourceGap returns the KnownSourceGap field if non-nil, zero value otherwise.

### GetKnownSourceGapOk

`func (o *CatalogueCoverage) GetKnownSourceGapOk() (*bool, bool)`

GetKnownSourceGapOk returns a tuple with the KnownSourceGap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownSourceGap

`func (o *CatalogueCoverage) SetKnownSourceGap(v bool)`

SetKnownSourceGap sets KnownSourceGap field to given value.

### HasKnownSourceGap

`func (o *CatalogueCoverage) HasKnownSourceGap() bool`

HasKnownSourceGap returns a boolean if a field has been set.

### SetKnownSourceGapNil

`func (o *CatalogueCoverage) SetKnownSourceGapNil(b bool)`

 SetKnownSourceGapNil sets the value for KnownSourceGap to be an explicit nil

### UnsetKnownSourceGap
`func (o *CatalogueCoverage) UnsetKnownSourceGap()`

UnsetKnownSourceGap ensures that no value is present for KnownSourceGap, not even an explicit nil
### GetSourceCheckedAt

`func (o *CatalogueCoverage) GetSourceCheckedAt() int64`

GetSourceCheckedAt returns the SourceCheckedAt field if non-nil, zero value otherwise.

### GetSourceCheckedAtOk

`func (o *CatalogueCoverage) GetSourceCheckedAtOk() (*int64, bool)`

GetSourceCheckedAtOk returns a tuple with the SourceCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCheckedAt

`func (o *CatalogueCoverage) SetSourceCheckedAt(v int64)`

SetSourceCheckedAt sets SourceCheckedAt field to given value.

### HasSourceCheckedAt

`func (o *CatalogueCoverage) HasSourceCheckedAt() bool`

HasSourceCheckedAt returns a boolean if a field has been set.

### SetSourceCheckedAtNil

`func (o *CatalogueCoverage) SetSourceCheckedAtNil(b bool)`

 SetSourceCheckedAtNil sets the value for SourceCheckedAt to be an explicit nil

### UnsetSourceCheckedAt
`func (o *CatalogueCoverage) UnsetSourceCheckedAt()`

UnsetSourceCheckedAt ensures that no value is present for SourceCheckedAt, not even an explicit nil
### GetSourceCheckExpiresAt

`func (o *CatalogueCoverage) GetSourceCheckExpiresAt() int64`

GetSourceCheckExpiresAt returns the SourceCheckExpiresAt field if non-nil, zero value otherwise.

### GetSourceCheckExpiresAtOk

`func (o *CatalogueCoverage) GetSourceCheckExpiresAtOk() (*int64, bool)`

GetSourceCheckExpiresAtOk returns a tuple with the SourceCheckExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCheckExpiresAt

`func (o *CatalogueCoverage) SetSourceCheckExpiresAt(v int64)`

SetSourceCheckExpiresAt sets SourceCheckExpiresAt field to given value.

### HasSourceCheckExpiresAt

`func (o *CatalogueCoverage) HasSourceCheckExpiresAt() bool`

HasSourceCheckExpiresAt returns a boolean if a field has been set.

### SetSourceCheckExpiresAtNil

`func (o *CatalogueCoverage) SetSourceCheckExpiresAtNil(b bool)`

 SetSourceCheckExpiresAtNil sets the value for SourceCheckExpiresAt to be an explicit nil

### UnsetSourceCheckExpiresAt
`func (o *CatalogueCoverage) UnsetSourceCheckExpiresAt()`

UnsetSourceCheckExpiresAt ensures that no value is present for SourceCheckExpiresAt, not even an explicit nil
### GetLatestPublishedPeriod

`func (o *CatalogueCoverage) GetLatestPublishedPeriod() string`

GetLatestPublishedPeriod returns the LatestPublishedPeriod field if non-nil, zero value otherwise.

### GetLatestPublishedPeriodOk

`func (o *CatalogueCoverage) GetLatestPublishedPeriodOk() (*string, bool)`

GetLatestPublishedPeriodOk returns a tuple with the LatestPublishedPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestPublishedPeriod

`func (o *CatalogueCoverage) SetLatestPublishedPeriod(v string)`

SetLatestPublishedPeriod sets LatestPublishedPeriod field to given value.

### HasLatestPublishedPeriod

`func (o *CatalogueCoverage) HasLatestPublishedPeriod() bool`

HasLatestPublishedPeriod returns a boolean if a field has been set.

### SetLatestPublishedPeriodNil

`func (o *CatalogueCoverage) SetLatestPublishedPeriodNil(b bool)`

 SetLatestPublishedPeriodNil sets the value for LatestPublishedPeriod to be an explicit nil

### UnsetLatestPublishedPeriod
`func (o *CatalogueCoverage) UnsetLatestPublishedPeriod()`

UnsetLatestPublishedPeriod ensures that no value is present for LatestPublishedPeriod, not even an explicit nil
### GetIsStale

`func (o *CatalogueCoverage) GetIsStale() bool`

GetIsStale returns the IsStale field if non-nil, zero value otherwise.

### GetIsStaleOk

`func (o *CatalogueCoverage) GetIsStaleOk() (*bool, bool)`

GetIsStaleOk returns a tuple with the IsStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStale

`func (o *CatalogueCoverage) SetIsStale(v bool)`

SetIsStale sets IsStale field to given value.

### HasIsStale

`func (o *CatalogueCoverage) HasIsStale() bool`

HasIsStale returns a boolean if a field has been set.

### SetIsStaleNil

`func (o *CatalogueCoverage) SetIsStaleNil(b bool)`

 SetIsStaleNil sets the value for IsStale to be an explicit nil

### UnsetIsStale
`func (o *CatalogueCoverage) UnsetIsStale()`

UnsetIsStale ensures that no value is present for IsStale, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


