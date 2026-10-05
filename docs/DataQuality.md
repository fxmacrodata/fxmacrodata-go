# DataQuality

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceFreshness** | Pointer to **NullableString** |  | [optional] 
**SourceMonitoringComplete** | Pointer to **NullableBool** |  | [optional] 
**KnownSourceGap** | Pointer to **NullableBool** |  | [optional] 
**SourceCheckedAt** | Pointer to **NullableInt64** |  | [optional] 
**SourceCheckExpiresAt** | Pointer to **NullableInt64** |  | [optional] 
**LatestPublishedPeriod** | Pointer to **NullableString** |  | [optional] 
**SourcePermissionStatus** | Pointer to **NullableString** |  | [optional] 
**CommodityCoverage** | Pointer to [**NullableCommodityCoverage**](CommodityCoverage.md) |  | [optional] 
**StoredHistoryCoverage** | Pointer to [**NullableCommodityCoverage**](CommodityCoverage.md) |  | [optional] 
**PointInTimeBasis** | Pointer to **NullableString** |  | [optional] 
**IsOfficial** | **bool** |  | 
**IsProxy** | **bool** |  | 
**IsFallback** | **bool** |  | 
**IsStale** | **bool** |  | 
**HasAnnouncementDatetime** | **bool** |  | 
**PointInTimeSafe** | **bool** | Whether the returned values have verified availability and vintage evidence; estimated release times do not qualify. | 
**UnverifiedVintageCount** | Pointer to **int64** |  | [optional] [default to 0]
**UnknownPublicationTimeCount** | Pointer to **int64** |  | [optional] [default to 0]
**LatestAvailableDate** | Pointer to **NullableString** |  | [optional] 
**LastUpdated** | Pointer to **NullableString** |  | [optional] 
**DataLagDays** | Pointer to **NullableInt64** |  | [optional] 
**SourceName** | Pointer to **NullableString** |  | [optional] 
**SourceType** | **string** |  | 
**IsDerived** | Pointer to **bool** |  | [optional] [default to false]
**RowCount** | Pointer to **int64** |  | [optional] [default to 0]
**AnnouncementDatetimeCount** | Pointer to **int64** |  | [optional] [default to 0]
**MissingAnnouncementDatetimeCount** | Pointer to **int64** |  | [optional] [default to 0]
**HasAssumedReleaseTimes** | Pointer to **bool** | True when at least one row in the quality scope carries a publication timestamp derived from the publication lag measured on that series&#39; own confirmed releases, because no official release calendar reaches the period. | [optional] [default to false]
**AssumedReleaseTimeCount** | Pointer to **int64** | Number of rows in the quality scope whose publication timestamp was derived rather than captured. Compare against row_count to size the derived portion of the requested window. | [optional] [default to 0]
**BoundedReleaseTimeCount** | Pointer to **int64** | Legacy rows carrying an upper-bound estimate flag; this does not certify replay safety. | [optional] [default to 0]
**UnboundedReleaseTimeCount** | Pointer to **int64** | Assumed timestamps without proven availability bounds. | [optional] [default to 0]
**QualityScope** | Pointer to **string** |  | [optional] [default to "full_result"]
**StaleAfterDays** | Pointer to **NullableInt64** |  | [optional] 
**RequestedWindowHasData** | Pointer to **NullableBool** |  | [optional] 
**RequestedWindowLatestDate** | Pointer to **NullableString** |  | [optional] 
**RequestedWindowIncludesLatestAvailable** | Pointer to **NullableBool** |  | [optional] 
**PageIncludesLatestAvailable** | Pointer to **NullableBool** |  | [optional] 
**ReturnedLatestAvailableBeforeWindow** | Pointer to **NullableBool** |  | [optional] 
**StalenessDays** | Pointer to **NullableInt64** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**DatetimeField** | Pointer to **NullableString** |  | [optional] 
**DatetimePrecision** | Pointer to **NullableString** |  | [optional] 
**DatetimeSemantics** | Pointer to **NullableString** |  | [optional] 
**SourceLegs** | Pointer to **[]map[string]interface{}** |  | [optional] 
**HistoricalPointInTime** | Pointer to [**NullablePointInTimeCompleteness**](PointInTimeCompleteness.md) |  | [optional] 

## Methods

### NewDataQuality

`func NewDataQuality(isOfficial bool, isProxy bool, isFallback bool, isStale bool, hasAnnouncementDatetime bool, pointInTimeSafe bool, sourceType string, ) *DataQuality`

NewDataQuality instantiates a new DataQuality object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataQualityWithDefaults

`func NewDataQualityWithDefaults() *DataQuality`

NewDataQualityWithDefaults instantiates a new DataQuality object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceFreshness

`func (o *DataQuality) GetSourceFreshness() string`

GetSourceFreshness returns the SourceFreshness field if non-nil, zero value otherwise.

### GetSourceFreshnessOk

`func (o *DataQuality) GetSourceFreshnessOk() (*string, bool)`

GetSourceFreshnessOk returns a tuple with the SourceFreshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFreshness

`func (o *DataQuality) SetSourceFreshness(v string)`

SetSourceFreshness sets SourceFreshness field to given value.

### HasSourceFreshness

`func (o *DataQuality) HasSourceFreshness() bool`

HasSourceFreshness returns a boolean if a field has been set.

### SetSourceFreshnessNil

`func (o *DataQuality) SetSourceFreshnessNil(b bool)`

 SetSourceFreshnessNil sets the value for SourceFreshness to be an explicit nil

### UnsetSourceFreshness
`func (o *DataQuality) UnsetSourceFreshness()`

UnsetSourceFreshness ensures that no value is present for SourceFreshness, not even an explicit nil
### GetSourceMonitoringComplete

`func (o *DataQuality) GetSourceMonitoringComplete() bool`

GetSourceMonitoringComplete returns the SourceMonitoringComplete field if non-nil, zero value otherwise.

### GetSourceMonitoringCompleteOk

`func (o *DataQuality) GetSourceMonitoringCompleteOk() (*bool, bool)`

GetSourceMonitoringCompleteOk returns a tuple with the SourceMonitoringComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMonitoringComplete

`func (o *DataQuality) SetSourceMonitoringComplete(v bool)`

SetSourceMonitoringComplete sets SourceMonitoringComplete field to given value.

### HasSourceMonitoringComplete

`func (o *DataQuality) HasSourceMonitoringComplete() bool`

HasSourceMonitoringComplete returns a boolean if a field has been set.

### SetSourceMonitoringCompleteNil

`func (o *DataQuality) SetSourceMonitoringCompleteNil(b bool)`

 SetSourceMonitoringCompleteNil sets the value for SourceMonitoringComplete to be an explicit nil

### UnsetSourceMonitoringComplete
`func (o *DataQuality) UnsetSourceMonitoringComplete()`

UnsetSourceMonitoringComplete ensures that no value is present for SourceMonitoringComplete, not even an explicit nil
### GetKnownSourceGap

`func (o *DataQuality) GetKnownSourceGap() bool`

GetKnownSourceGap returns the KnownSourceGap field if non-nil, zero value otherwise.

### GetKnownSourceGapOk

`func (o *DataQuality) GetKnownSourceGapOk() (*bool, bool)`

GetKnownSourceGapOk returns a tuple with the KnownSourceGap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownSourceGap

`func (o *DataQuality) SetKnownSourceGap(v bool)`

SetKnownSourceGap sets KnownSourceGap field to given value.

### HasKnownSourceGap

`func (o *DataQuality) HasKnownSourceGap() bool`

HasKnownSourceGap returns a boolean if a field has been set.

### SetKnownSourceGapNil

`func (o *DataQuality) SetKnownSourceGapNil(b bool)`

 SetKnownSourceGapNil sets the value for KnownSourceGap to be an explicit nil

### UnsetKnownSourceGap
`func (o *DataQuality) UnsetKnownSourceGap()`

UnsetKnownSourceGap ensures that no value is present for KnownSourceGap, not even an explicit nil
### GetSourceCheckedAt

`func (o *DataQuality) GetSourceCheckedAt() int64`

GetSourceCheckedAt returns the SourceCheckedAt field if non-nil, zero value otherwise.

### GetSourceCheckedAtOk

`func (o *DataQuality) GetSourceCheckedAtOk() (*int64, bool)`

GetSourceCheckedAtOk returns a tuple with the SourceCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCheckedAt

`func (o *DataQuality) SetSourceCheckedAt(v int64)`

SetSourceCheckedAt sets SourceCheckedAt field to given value.

### HasSourceCheckedAt

`func (o *DataQuality) HasSourceCheckedAt() bool`

HasSourceCheckedAt returns a boolean if a field has been set.

### SetSourceCheckedAtNil

`func (o *DataQuality) SetSourceCheckedAtNil(b bool)`

 SetSourceCheckedAtNil sets the value for SourceCheckedAt to be an explicit nil

### UnsetSourceCheckedAt
`func (o *DataQuality) UnsetSourceCheckedAt()`

UnsetSourceCheckedAt ensures that no value is present for SourceCheckedAt, not even an explicit nil
### GetSourceCheckExpiresAt

`func (o *DataQuality) GetSourceCheckExpiresAt() int64`

GetSourceCheckExpiresAt returns the SourceCheckExpiresAt field if non-nil, zero value otherwise.

### GetSourceCheckExpiresAtOk

`func (o *DataQuality) GetSourceCheckExpiresAtOk() (*int64, bool)`

GetSourceCheckExpiresAtOk returns a tuple with the SourceCheckExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCheckExpiresAt

`func (o *DataQuality) SetSourceCheckExpiresAt(v int64)`

SetSourceCheckExpiresAt sets SourceCheckExpiresAt field to given value.

### HasSourceCheckExpiresAt

`func (o *DataQuality) HasSourceCheckExpiresAt() bool`

HasSourceCheckExpiresAt returns a boolean if a field has been set.

### SetSourceCheckExpiresAtNil

`func (o *DataQuality) SetSourceCheckExpiresAtNil(b bool)`

 SetSourceCheckExpiresAtNil sets the value for SourceCheckExpiresAt to be an explicit nil

### UnsetSourceCheckExpiresAt
`func (o *DataQuality) UnsetSourceCheckExpiresAt()`

UnsetSourceCheckExpiresAt ensures that no value is present for SourceCheckExpiresAt, not even an explicit nil
### GetLatestPublishedPeriod

`func (o *DataQuality) GetLatestPublishedPeriod() string`

GetLatestPublishedPeriod returns the LatestPublishedPeriod field if non-nil, zero value otherwise.

### GetLatestPublishedPeriodOk

`func (o *DataQuality) GetLatestPublishedPeriodOk() (*string, bool)`

GetLatestPublishedPeriodOk returns a tuple with the LatestPublishedPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestPublishedPeriod

`func (o *DataQuality) SetLatestPublishedPeriod(v string)`

SetLatestPublishedPeriod sets LatestPublishedPeriod field to given value.

### HasLatestPublishedPeriod

`func (o *DataQuality) HasLatestPublishedPeriod() bool`

HasLatestPublishedPeriod returns a boolean if a field has been set.

### SetLatestPublishedPeriodNil

`func (o *DataQuality) SetLatestPublishedPeriodNil(b bool)`

 SetLatestPublishedPeriodNil sets the value for LatestPublishedPeriod to be an explicit nil

### UnsetLatestPublishedPeriod
`func (o *DataQuality) UnsetLatestPublishedPeriod()`

UnsetLatestPublishedPeriod ensures that no value is present for LatestPublishedPeriod, not even an explicit nil
### GetSourcePermissionStatus

`func (o *DataQuality) GetSourcePermissionStatus() string`

GetSourcePermissionStatus returns the SourcePermissionStatus field if non-nil, zero value otherwise.

### GetSourcePermissionStatusOk

`func (o *DataQuality) GetSourcePermissionStatusOk() (*string, bool)`

GetSourcePermissionStatusOk returns a tuple with the SourcePermissionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePermissionStatus

`func (o *DataQuality) SetSourcePermissionStatus(v string)`

SetSourcePermissionStatus sets SourcePermissionStatus field to given value.

### HasSourcePermissionStatus

`func (o *DataQuality) HasSourcePermissionStatus() bool`

HasSourcePermissionStatus returns a boolean if a field has been set.

### SetSourcePermissionStatusNil

`func (o *DataQuality) SetSourcePermissionStatusNil(b bool)`

 SetSourcePermissionStatusNil sets the value for SourcePermissionStatus to be an explicit nil

### UnsetSourcePermissionStatus
`func (o *DataQuality) UnsetSourcePermissionStatus()`

UnsetSourcePermissionStatus ensures that no value is present for SourcePermissionStatus, not even an explicit nil
### GetCommodityCoverage

`func (o *DataQuality) GetCommodityCoverage() CommodityCoverage`

GetCommodityCoverage returns the CommodityCoverage field if non-nil, zero value otherwise.

### GetCommodityCoverageOk

`func (o *DataQuality) GetCommodityCoverageOk() (*CommodityCoverage, bool)`

GetCommodityCoverageOk returns a tuple with the CommodityCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommodityCoverage

`func (o *DataQuality) SetCommodityCoverage(v CommodityCoverage)`

SetCommodityCoverage sets CommodityCoverage field to given value.

### HasCommodityCoverage

`func (o *DataQuality) HasCommodityCoverage() bool`

HasCommodityCoverage returns a boolean if a field has been set.

### SetCommodityCoverageNil

`func (o *DataQuality) SetCommodityCoverageNil(b bool)`

 SetCommodityCoverageNil sets the value for CommodityCoverage to be an explicit nil

### UnsetCommodityCoverage
`func (o *DataQuality) UnsetCommodityCoverage()`

UnsetCommodityCoverage ensures that no value is present for CommodityCoverage, not even an explicit nil
### GetStoredHistoryCoverage

`func (o *DataQuality) GetStoredHistoryCoverage() CommodityCoverage`

GetStoredHistoryCoverage returns the StoredHistoryCoverage field if non-nil, zero value otherwise.

### GetStoredHistoryCoverageOk

`func (o *DataQuality) GetStoredHistoryCoverageOk() (*CommodityCoverage, bool)`

GetStoredHistoryCoverageOk returns a tuple with the StoredHistoryCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredHistoryCoverage

`func (o *DataQuality) SetStoredHistoryCoverage(v CommodityCoverage)`

SetStoredHistoryCoverage sets StoredHistoryCoverage field to given value.

### HasStoredHistoryCoverage

`func (o *DataQuality) HasStoredHistoryCoverage() bool`

HasStoredHistoryCoverage returns a boolean if a field has been set.

### SetStoredHistoryCoverageNil

`func (o *DataQuality) SetStoredHistoryCoverageNil(b bool)`

 SetStoredHistoryCoverageNil sets the value for StoredHistoryCoverage to be an explicit nil

### UnsetStoredHistoryCoverage
`func (o *DataQuality) UnsetStoredHistoryCoverage()`

UnsetStoredHistoryCoverage ensures that no value is present for StoredHistoryCoverage, not even an explicit nil
### GetPointInTimeBasis

`func (o *DataQuality) GetPointInTimeBasis() string`

GetPointInTimeBasis returns the PointInTimeBasis field if non-nil, zero value otherwise.

### GetPointInTimeBasisOk

`func (o *DataQuality) GetPointInTimeBasisOk() (*string, bool)`

GetPointInTimeBasisOk returns a tuple with the PointInTimeBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeBasis

`func (o *DataQuality) SetPointInTimeBasis(v string)`

SetPointInTimeBasis sets PointInTimeBasis field to given value.

### HasPointInTimeBasis

`func (o *DataQuality) HasPointInTimeBasis() bool`

HasPointInTimeBasis returns a boolean if a field has been set.

### SetPointInTimeBasisNil

`func (o *DataQuality) SetPointInTimeBasisNil(b bool)`

 SetPointInTimeBasisNil sets the value for PointInTimeBasis to be an explicit nil

### UnsetPointInTimeBasis
`func (o *DataQuality) UnsetPointInTimeBasis()`

UnsetPointInTimeBasis ensures that no value is present for PointInTimeBasis, not even an explicit nil
### GetIsOfficial

`func (o *DataQuality) GetIsOfficial() bool`

GetIsOfficial returns the IsOfficial field if non-nil, zero value otherwise.

### GetIsOfficialOk

`func (o *DataQuality) GetIsOfficialOk() (*bool, bool)`

GetIsOfficialOk returns a tuple with the IsOfficial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOfficial

`func (o *DataQuality) SetIsOfficial(v bool)`

SetIsOfficial sets IsOfficial field to given value.


### GetIsProxy

`func (o *DataQuality) GetIsProxy() bool`

GetIsProxy returns the IsProxy field if non-nil, zero value otherwise.

### GetIsProxyOk

`func (o *DataQuality) GetIsProxyOk() (*bool, bool)`

GetIsProxyOk returns a tuple with the IsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxy

`func (o *DataQuality) SetIsProxy(v bool)`

SetIsProxy sets IsProxy field to given value.


### GetIsFallback

`func (o *DataQuality) GetIsFallback() bool`

GetIsFallback returns the IsFallback field if non-nil, zero value otherwise.

### GetIsFallbackOk

`func (o *DataQuality) GetIsFallbackOk() (*bool, bool)`

GetIsFallbackOk returns a tuple with the IsFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFallback

`func (o *DataQuality) SetIsFallback(v bool)`

SetIsFallback sets IsFallback field to given value.


### GetIsStale

`func (o *DataQuality) GetIsStale() bool`

GetIsStale returns the IsStale field if non-nil, zero value otherwise.

### GetIsStaleOk

`func (o *DataQuality) GetIsStaleOk() (*bool, bool)`

GetIsStaleOk returns a tuple with the IsStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStale

`func (o *DataQuality) SetIsStale(v bool)`

SetIsStale sets IsStale field to given value.


### GetHasAnnouncementDatetime

`func (o *DataQuality) GetHasAnnouncementDatetime() bool`

GetHasAnnouncementDatetime returns the HasAnnouncementDatetime field if non-nil, zero value otherwise.

### GetHasAnnouncementDatetimeOk

`func (o *DataQuality) GetHasAnnouncementDatetimeOk() (*bool, bool)`

GetHasAnnouncementDatetimeOk returns a tuple with the HasAnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAnnouncementDatetime

`func (o *DataQuality) SetHasAnnouncementDatetime(v bool)`

SetHasAnnouncementDatetime sets HasAnnouncementDatetime field to given value.


### GetPointInTimeSafe

`func (o *DataQuality) GetPointInTimeSafe() bool`

GetPointInTimeSafe returns the PointInTimeSafe field if non-nil, zero value otherwise.

### GetPointInTimeSafeOk

`func (o *DataQuality) GetPointInTimeSafeOk() (*bool, bool)`

GetPointInTimeSafeOk returns a tuple with the PointInTimeSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeSafe

`func (o *DataQuality) SetPointInTimeSafe(v bool)`

SetPointInTimeSafe sets PointInTimeSafe field to given value.


### GetUnverifiedVintageCount

`func (o *DataQuality) GetUnverifiedVintageCount() int64`

GetUnverifiedVintageCount returns the UnverifiedVintageCount field if non-nil, zero value otherwise.

### GetUnverifiedVintageCountOk

`func (o *DataQuality) GetUnverifiedVintageCountOk() (*int64, bool)`

GetUnverifiedVintageCountOk returns a tuple with the UnverifiedVintageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnverifiedVintageCount

`func (o *DataQuality) SetUnverifiedVintageCount(v int64)`

SetUnverifiedVintageCount sets UnverifiedVintageCount field to given value.

### HasUnverifiedVintageCount

`func (o *DataQuality) HasUnverifiedVintageCount() bool`

HasUnverifiedVintageCount returns a boolean if a field has been set.

### GetUnknownPublicationTimeCount

`func (o *DataQuality) GetUnknownPublicationTimeCount() int64`

GetUnknownPublicationTimeCount returns the UnknownPublicationTimeCount field if non-nil, zero value otherwise.

### GetUnknownPublicationTimeCountOk

`func (o *DataQuality) GetUnknownPublicationTimeCountOk() (*int64, bool)`

GetUnknownPublicationTimeCountOk returns a tuple with the UnknownPublicationTimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownPublicationTimeCount

`func (o *DataQuality) SetUnknownPublicationTimeCount(v int64)`

SetUnknownPublicationTimeCount sets UnknownPublicationTimeCount field to given value.

### HasUnknownPublicationTimeCount

`func (o *DataQuality) HasUnknownPublicationTimeCount() bool`

HasUnknownPublicationTimeCount returns a boolean if a field has been set.

### GetLatestAvailableDate

`func (o *DataQuality) GetLatestAvailableDate() string`

GetLatestAvailableDate returns the LatestAvailableDate field if non-nil, zero value otherwise.

### GetLatestAvailableDateOk

`func (o *DataQuality) GetLatestAvailableDateOk() (*string, bool)`

GetLatestAvailableDateOk returns a tuple with the LatestAvailableDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAvailableDate

`func (o *DataQuality) SetLatestAvailableDate(v string)`

SetLatestAvailableDate sets LatestAvailableDate field to given value.

### HasLatestAvailableDate

`func (o *DataQuality) HasLatestAvailableDate() bool`

HasLatestAvailableDate returns a boolean if a field has been set.

### SetLatestAvailableDateNil

`func (o *DataQuality) SetLatestAvailableDateNil(b bool)`

 SetLatestAvailableDateNil sets the value for LatestAvailableDate to be an explicit nil

### UnsetLatestAvailableDate
`func (o *DataQuality) UnsetLatestAvailableDate()`

UnsetLatestAvailableDate ensures that no value is present for LatestAvailableDate, not even an explicit nil
### GetLastUpdated

`func (o *DataQuality) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DataQuality) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DataQuality) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *DataQuality) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *DataQuality) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *DataQuality) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetDataLagDays

`func (o *DataQuality) GetDataLagDays() int64`

GetDataLagDays returns the DataLagDays field if non-nil, zero value otherwise.

### GetDataLagDaysOk

`func (o *DataQuality) GetDataLagDaysOk() (*int64, bool)`

GetDataLagDaysOk returns a tuple with the DataLagDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLagDays

`func (o *DataQuality) SetDataLagDays(v int64)`

SetDataLagDays sets DataLagDays field to given value.

### HasDataLagDays

`func (o *DataQuality) HasDataLagDays() bool`

HasDataLagDays returns a boolean if a field has been set.

### SetDataLagDaysNil

`func (o *DataQuality) SetDataLagDaysNil(b bool)`

 SetDataLagDaysNil sets the value for DataLagDays to be an explicit nil

### UnsetDataLagDays
`func (o *DataQuality) UnsetDataLagDays()`

UnsetDataLagDays ensures that no value is present for DataLagDays, not even an explicit nil
### GetSourceName

`func (o *DataQuality) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *DataQuality) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *DataQuality) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *DataQuality) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *DataQuality) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *DataQuality) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetSourceType

`func (o *DataQuality) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *DataQuality) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *DataQuality) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetIsDerived

`func (o *DataQuality) GetIsDerived() bool`

GetIsDerived returns the IsDerived field if non-nil, zero value otherwise.

### GetIsDerivedOk

`func (o *DataQuality) GetIsDerivedOk() (*bool, bool)`

GetIsDerivedOk returns a tuple with the IsDerived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDerived

`func (o *DataQuality) SetIsDerived(v bool)`

SetIsDerived sets IsDerived field to given value.

### HasIsDerived

`func (o *DataQuality) HasIsDerived() bool`

HasIsDerived returns a boolean if a field has been set.

### GetRowCount

`func (o *DataQuality) GetRowCount() int64`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *DataQuality) GetRowCountOk() (*int64, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *DataQuality) SetRowCount(v int64)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *DataQuality) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### GetAnnouncementDatetimeCount

`func (o *DataQuality) GetAnnouncementDatetimeCount() int64`

GetAnnouncementDatetimeCount returns the AnnouncementDatetimeCount field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeCountOk

`func (o *DataQuality) GetAnnouncementDatetimeCountOk() (*int64, bool)`

GetAnnouncementDatetimeCountOk returns a tuple with the AnnouncementDatetimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetimeCount

`func (o *DataQuality) SetAnnouncementDatetimeCount(v int64)`

SetAnnouncementDatetimeCount sets AnnouncementDatetimeCount field to given value.

### HasAnnouncementDatetimeCount

`func (o *DataQuality) HasAnnouncementDatetimeCount() bool`

HasAnnouncementDatetimeCount returns a boolean if a field has been set.

### GetMissingAnnouncementDatetimeCount

`func (o *DataQuality) GetMissingAnnouncementDatetimeCount() int64`

GetMissingAnnouncementDatetimeCount returns the MissingAnnouncementDatetimeCount field if non-nil, zero value otherwise.

### GetMissingAnnouncementDatetimeCountOk

`func (o *DataQuality) GetMissingAnnouncementDatetimeCountOk() (*int64, bool)`

GetMissingAnnouncementDatetimeCountOk returns a tuple with the MissingAnnouncementDatetimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingAnnouncementDatetimeCount

`func (o *DataQuality) SetMissingAnnouncementDatetimeCount(v int64)`

SetMissingAnnouncementDatetimeCount sets MissingAnnouncementDatetimeCount field to given value.

### HasMissingAnnouncementDatetimeCount

`func (o *DataQuality) HasMissingAnnouncementDatetimeCount() bool`

HasMissingAnnouncementDatetimeCount returns a boolean if a field has been set.

### GetHasAssumedReleaseTimes

`func (o *DataQuality) GetHasAssumedReleaseTimes() bool`

GetHasAssumedReleaseTimes returns the HasAssumedReleaseTimes field if non-nil, zero value otherwise.

### GetHasAssumedReleaseTimesOk

`func (o *DataQuality) GetHasAssumedReleaseTimesOk() (*bool, bool)`

GetHasAssumedReleaseTimesOk returns a tuple with the HasAssumedReleaseTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAssumedReleaseTimes

`func (o *DataQuality) SetHasAssumedReleaseTimes(v bool)`

SetHasAssumedReleaseTimes sets HasAssumedReleaseTimes field to given value.

### HasHasAssumedReleaseTimes

`func (o *DataQuality) HasHasAssumedReleaseTimes() bool`

HasHasAssumedReleaseTimes returns a boolean if a field has been set.

### GetAssumedReleaseTimeCount

`func (o *DataQuality) GetAssumedReleaseTimeCount() int64`

GetAssumedReleaseTimeCount returns the AssumedReleaseTimeCount field if non-nil, zero value otherwise.

### GetAssumedReleaseTimeCountOk

`func (o *DataQuality) GetAssumedReleaseTimeCountOk() (*int64, bool)`

GetAssumedReleaseTimeCountOk returns a tuple with the AssumedReleaseTimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumedReleaseTimeCount

`func (o *DataQuality) SetAssumedReleaseTimeCount(v int64)`

SetAssumedReleaseTimeCount sets AssumedReleaseTimeCount field to given value.

### HasAssumedReleaseTimeCount

`func (o *DataQuality) HasAssumedReleaseTimeCount() bool`

HasAssumedReleaseTimeCount returns a boolean if a field has been set.

### GetBoundedReleaseTimeCount

`func (o *DataQuality) GetBoundedReleaseTimeCount() int64`

GetBoundedReleaseTimeCount returns the BoundedReleaseTimeCount field if non-nil, zero value otherwise.

### GetBoundedReleaseTimeCountOk

`func (o *DataQuality) GetBoundedReleaseTimeCountOk() (*int64, bool)`

GetBoundedReleaseTimeCountOk returns a tuple with the BoundedReleaseTimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundedReleaseTimeCount

`func (o *DataQuality) SetBoundedReleaseTimeCount(v int64)`

SetBoundedReleaseTimeCount sets BoundedReleaseTimeCount field to given value.

### HasBoundedReleaseTimeCount

`func (o *DataQuality) HasBoundedReleaseTimeCount() bool`

HasBoundedReleaseTimeCount returns a boolean if a field has been set.

### GetUnboundedReleaseTimeCount

`func (o *DataQuality) GetUnboundedReleaseTimeCount() int64`

GetUnboundedReleaseTimeCount returns the UnboundedReleaseTimeCount field if non-nil, zero value otherwise.

### GetUnboundedReleaseTimeCountOk

`func (o *DataQuality) GetUnboundedReleaseTimeCountOk() (*int64, bool)`

GetUnboundedReleaseTimeCountOk returns a tuple with the UnboundedReleaseTimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnboundedReleaseTimeCount

`func (o *DataQuality) SetUnboundedReleaseTimeCount(v int64)`

SetUnboundedReleaseTimeCount sets UnboundedReleaseTimeCount field to given value.

### HasUnboundedReleaseTimeCount

`func (o *DataQuality) HasUnboundedReleaseTimeCount() bool`

HasUnboundedReleaseTimeCount returns a boolean if a field has been set.

### GetQualityScope

`func (o *DataQuality) GetQualityScope() string`

GetQualityScope returns the QualityScope field if non-nil, zero value otherwise.

### GetQualityScopeOk

`func (o *DataQuality) GetQualityScopeOk() (*string, bool)`

GetQualityScopeOk returns a tuple with the QualityScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityScope

`func (o *DataQuality) SetQualityScope(v string)`

SetQualityScope sets QualityScope field to given value.

### HasQualityScope

`func (o *DataQuality) HasQualityScope() bool`

HasQualityScope returns a boolean if a field has been set.

### GetStaleAfterDays

`func (o *DataQuality) GetStaleAfterDays() int64`

GetStaleAfterDays returns the StaleAfterDays field if non-nil, zero value otherwise.

### GetStaleAfterDaysOk

`func (o *DataQuality) GetStaleAfterDaysOk() (*int64, bool)`

GetStaleAfterDaysOk returns a tuple with the StaleAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleAfterDays

`func (o *DataQuality) SetStaleAfterDays(v int64)`

SetStaleAfterDays sets StaleAfterDays field to given value.

### HasStaleAfterDays

`func (o *DataQuality) HasStaleAfterDays() bool`

HasStaleAfterDays returns a boolean if a field has been set.

### SetStaleAfterDaysNil

`func (o *DataQuality) SetStaleAfterDaysNil(b bool)`

 SetStaleAfterDaysNil sets the value for StaleAfterDays to be an explicit nil

### UnsetStaleAfterDays
`func (o *DataQuality) UnsetStaleAfterDays()`

UnsetStaleAfterDays ensures that no value is present for StaleAfterDays, not even an explicit nil
### GetRequestedWindowHasData

`func (o *DataQuality) GetRequestedWindowHasData() bool`

GetRequestedWindowHasData returns the RequestedWindowHasData field if non-nil, zero value otherwise.

### GetRequestedWindowHasDataOk

`func (o *DataQuality) GetRequestedWindowHasDataOk() (*bool, bool)`

GetRequestedWindowHasDataOk returns a tuple with the RequestedWindowHasData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedWindowHasData

`func (o *DataQuality) SetRequestedWindowHasData(v bool)`

SetRequestedWindowHasData sets RequestedWindowHasData field to given value.

### HasRequestedWindowHasData

`func (o *DataQuality) HasRequestedWindowHasData() bool`

HasRequestedWindowHasData returns a boolean if a field has been set.

### SetRequestedWindowHasDataNil

`func (o *DataQuality) SetRequestedWindowHasDataNil(b bool)`

 SetRequestedWindowHasDataNil sets the value for RequestedWindowHasData to be an explicit nil

### UnsetRequestedWindowHasData
`func (o *DataQuality) UnsetRequestedWindowHasData()`

UnsetRequestedWindowHasData ensures that no value is present for RequestedWindowHasData, not even an explicit nil
### GetRequestedWindowLatestDate

`func (o *DataQuality) GetRequestedWindowLatestDate() string`

GetRequestedWindowLatestDate returns the RequestedWindowLatestDate field if non-nil, zero value otherwise.

### GetRequestedWindowLatestDateOk

`func (o *DataQuality) GetRequestedWindowLatestDateOk() (*string, bool)`

GetRequestedWindowLatestDateOk returns a tuple with the RequestedWindowLatestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedWindowLatestDate

`func (o *DataQuality) SetRequestedWindowLatestDate(v string)`

SetRequestedWindowLatestDate sets RequestedWindowLatestDate field to given value.

### HasRequestedWindowLatestDate

`func (o *DataQuality) HasRequestedWindowLatestDate() bool`

HasRequestedWindowLatestDate returns a boolean if a field has been set.

### SetRequestedWindowLatestDateNil

`func (o *DataQuality) SetRequestedWindowLatestDateNil(b bool)`

 SetRequestedWindowLatestDateNil sets the value for RequestedWindowLatestDate to be an explicit nil

### UnsetRequestedWindowLatestDate
`func (o *DataQuality) UnsetRequestedWindowLatestDate()`

UnsetRequestedWindowLatestDate ensures that no value is present for RequestedWindowLatestDate, not even an explicit nil
### GetRequestedWindowIncludesLatestAvailable

`func (o *DataQuality) GetRequestedWindowIncludesLatestAvailable() bool`

GetRequestedWindowIncludesLatestAvailable returns the RequestedWindowIncludesLatestAvailable field if non-nil, zero value otherwise.

### GetRequestedWindowIncludesLatestAvailableOk

`func (o *DataQuality) GetRequestedWindowIncludesLatestAvailableOk() (*bool, bool)`

GetRequestedWindowIncludesLatestAvailableOk returns a tuple with the RequestedWindowIncludesLatestAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedWindowIncludesLatestAvailable

`func (o *DataQuality) SetRequestedWindowIncludesLatestAvailable(v bool)`

SetRequestedWindowIncludesLatestAvailable sets RequestedWindowIncludesLatestAvailable field to given value.

### HasRequestedWindowIncludesLatestAvailable

`func (o *DataQuality) HasRequestedWindowIncludesLatestAvailable() bool`

HasRequestedWindowIncludesLatestAvailable returns a boolean if a field has been set.

### SetRequestedWindowIncludesLatestAvailableNil

`func (o *DataQuality) SetRequestedWindowIncludesLatestAvailableNil(b bool)`

 SetRequestedWindowIncludesLatestAvailableNil sets the value for RequestedWindowIncludesLatestAvailable to be an explicit nil

### UnsetRequestedWindowIncludesLatestAvailable
`func (o *DataQuality) UnsetRequestedWindowIncludesLatestAvailable()`

UnsetRequestedWindowIncludesLatestAvailable ensures that no value is present for RequestedWindowIncludesLatestAvailable, not even an explicit nil
### GetPageIncludesLatestAvailable

`func (o *DataQuality) GetPageIncludesLatestAvailable() bool`

GetPageIncludesLatestAvailable returns the PageIncludesLatestAvailable field if non-nil, zero value otherwise.

### GetPageIncludesLatestAvailableOk

`func (o *DataQuality) GetPageIncludesLatestAvailableOk() (*bool, bool)`

GetPageIncludesLatestAvailableOk returns a tuple with the PageIncludesLatestAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIncludesLatestAvailable

`func (o *DataQuality) SetPageIncludesLatestAvailable(v bool)`

SetPageIncludesLatestAvailable sets PageIncludesLatestAvailable field to given value.

### HasPageIncludesLatestAvailable

`func (o *DataQuality) HasPageIncludesLatestAvailable() bool`

HasPageIncludesLatestAvailable returns a boolean if a field has been set.

### SetPageIncludesLatestAvailableNil

`func (o *DataQuality) SetPageIncludesLatestAvailableNil(b bool)`

 SetPageIncludesLatestAvailableNil sets the value for PageIncludesLatestAvailable to be an explicit nil

### UnsetPageIncludesLatestAvailable
`func (o *DataQuality) UnsetPageIncludesLatestAvailable()`

UnsetPageIncludesLatestAvailable ensures that no value is present for PageIncludesLatestAvailable, not even an explicit nil
### GetReturnedLatestAvailableBeforeWindow

`func (o *DataQuality) GetReturnedLatestAvailableBeforeWindow() bool`

GetReturnedLatestAvailableBeforeWindow returns the ReturnedLatestAvailableBeforeWindow field if non-nil, zero value otherwise.

### GetReturnedLatestAvailableBeforeWindowOk

`func (o *DataQuality) GetReturnedLatestAvailableBeforeWindowOk() (*bool, bool)`

GetReturnedLatestAvailableBeforeWindowOk returns a tuple with the ReturnedLatestAvailableBeforeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedLatestAvailableBeforeWindow

`func (o *DataQuality) SetReturnedLatestAvailableBeforeWindow(v bool)`

SetReturnedLatestAvailableBeforeWindow sets ReturnedLatestAvailableBeforeWindow field to given value.

### HasReturnedLatestAvailableBeforeWindow

`func (o *DataQuality) HasReturnedLatestAvailableBeforeWindow() bool`

HasReturnedLatestAvailableBeforeWindow returns a boolean if a field has been set.

### SetReturnedLatestAvailableBeforeWindowNil

`func (o *DataQuality) SetReturnedLatestAvailableBeforeWindowNil(b bool)`

 SetReturnedLatestAvailableBeforeWindowNil sets the value for ReturnedLatestAvailableBeforeWindow to be an explicit nil

### UnsetReturnedLatestAvailableBeforeWindow
`func (o *DataQuality) UnsetReturnedLatestAvailableBeforeWindow()`

UnsetReturnedLatestAvailableBeforeWindow ensures that no value is present for ReturnedLatestAvailableBeforeWindow, not even an explicit nil
### GetStalenessDays

`func (o *DataQuality) GetStalenessDays() int64`

GetStalenessDays returns the StalenessDays field if non-nil, zero value otherwise.

### GetStalenessDaysOk

`func (o *DataQuality) GetStalenessDaysOk() (*int64, bool)`

GetStalenessDaysOk returns a tuple with the StalenessDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStalenessDays

`func (o *DataQuality) SetStalenessDays(v int64)`

SetStalenessDays sets StalenessDays field to given value.

### HasStalenessDays

`func (o *DataQuality) HasStalenessDays() bool`

HasStalenessDays returns a boolean if a field has been set.

### SetStalenessDaysNil

`func (o *DataQuality) SetStalenessDaysNil(b bool)`

 SetStalenessDaysNil sets the value for StalenessDays to be an explicit nil

### UnsetStalenessDays
`func (o *DataQuality) UnsetStalenessDays()`

UnsetStalenessDays ensures that no value is present for StalenessDays, not even an explicit nil
### GetReason

`func (o *DataQuality) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DataQuality) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DataQuality) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DataQuality) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *DataQuality) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *DataQuality) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetDatetimeField

`func (o *DataQuality) GetDatetimeField() string`

GetDatetimeField returns the DatetimeField field if non-nil, zero value otherwise.

### GetDatetimeFieldOk

`func (o *DataQuality) GetDatetimeFieldOk() (*string, bool)`

GetDatetimeFieldOk returns a tuple with the DatetimeField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeField

`func (o *DataQuality) SetDatetimeField(v string)`

SetDatetimeField sets DatetimeField field to given value.

### HasDatetimeField

`func (o *DataQuality) HasDatetimeField() bool`

HasDatetimeField returns a boolean if a field has been set.

### SetDatetimeFieldNil

`func (o *DataQuality) SetDatetimeFieldNil(b bool)`

 SetDatetimeFieldNil sets the value for DatetimeField to be an explicit nil

### UnsetDatetimeField
`func (o *DataQuality) UnsetDatetimeField()`

UnsetDatetimeField ensures that no value is present for DatetimeField, not even an explicit nil
### GetDatetimePrecision

`func (o *DataQuality) GetDatetimePrecision() string`

GetDatetimePrecision returns the DatetimePrecision field if non-nil, zero value otherwise.

### GetDatetimePrecisionOk

`func (o *DataQuality) GetDatetimePrecisionOk() (*string, bool)`

GetDatetimePrecisionOk returns a tuple with the DatetimePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimePrecision

`func (o *DataQuality) SetDatetimePrecision(v string)`

SetDatetimePrecision sets DatetimePrecision field to given value.

### HasDatetimePrecision

`func (o *DataQuality) HasDatetimePrecision() bool`

HasDatetimePrecision returns a boolean if a field has been set.

### SetDatetimePrecisionNil

`func (o *DataQuality) SetDatetimePrecisionNil(b bool)`

 SetDatetimePrecisionNil sets the value for DatetimePrecision to be an explicit nil

### UnsetDatetimePrecision
`func (o *DataQuality) UnsetDatetimePrecision()`

UnsetDatetimePrecision ensures that no value is present for DatetimePrecision, not even an explicit nil
### GetDatetimeSemantics

`func (o *DataQuality) GetDatetimeSemantics() string`

GetDatetimeSemantics returns the DatetimeSemantics field if non-nil, zero value otherwise.

### GetDatetimeSemanticsOk

`func (o *DataQuality) GetDatetimeSemanticsOk() (*string, bool)`

GetDatetimeSemanticsOk returns a tuple with the DatetimeSemantics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeSemantics

`func (o *DataQuality) SetDatetimeSemantics(v string)`

SetDatetimeSemantics sets DatetimeSemantics field to given value.

### HasDatetimeSemantics

`func (o *DataQuality) HasDatetimeSemantics() bool`

HasDatetimeSemantics returns a boolean if a field has been set.

### SetDatetimeSemanticsNil

`func (o *DataQuality) SetDatetimeSemanticsNil(b bool)`

 SetDatetimeSemanticsNil sets the value for DatetimeSemantics to be an explicit nil

### UnsetDatetimeSemantics
`func (o *DataQuality) UnsetDatetimeSemantics()`

UnsetDatetimeSemantics ensures that no value is present for DatetimeSemantics, not even an explicit nil
### GetSourceLegs

`func (o *DataQuality) GetSourceLegs() []map[string]interface{}`

GetSourceLegs returns the SourceLegs field if non-nil, zero value otherwise.

### GetSourceLegsOk

`func (o *DataQuality) GetSourceLegsOk() (*[]map[string]interface{}, bool)`

GetSourceLegsOk returns a tuple with the SourceLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLegs

`func (o *DataQuality) SetSourceLegs(v []map[string]interface{})`

SetSourceLegs sets SourceLegs field to given value.

### HasSourceLegs

`func (o *DataQuality) HasSourceLegs() bool`

HasSourceLegs returns a boolean if a field has been set.

### SetSourceLegsNil

`func (o *DataQuality) SetSourceLegsNil(b bool)`

 SetSourceLegsNil sets the value for SourceLegs to be an explicit nil

### UnsetSourceLegs
`func (o *DataQuality) UnsetSourceLegs()`

UnsetSourceLegs ensures that no value is present for SourceLegs, not even an explicit nil
### GetHistoricalPointInTime

`func (o *DataQuality) GetHistoricalPointInTime() PointInTimeCompleteness`

GetHistoricalPointInTime returns the HistoricalPointInTime field if non-nil, zero value otherwise.

### GetHistoricalPointInTimeOk

`func (o *DataQuality) GetHistoricalPointInTimeOk() (*PointInTimeCompleteness, bool)`

GetHistoricalPointInTimeOk returns a tuple with the HistoricalPointInTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalPointInTime

`func (o *DataQuality) SetHistoricalPointInTime(v PointInTimeCompleteness)`

SetHistoricalPointInTime sets HistoricalPointInTime field to given value.

### HasHistoricalPointInTime

`func (o *DataQuality) HasHistoricalPointInTime() bool`

HasHistoricalPointInTime returns a boolean if a field has been set.

### SetHistoricalPointInTimeNil

`func (o *DataQuality) SetHistoricalPointInTimeNil(b bool)`

 SetHistoricalPointInTimeNil sets the value for HistoricalPointInTime to be an explicit nil

### UnsetHistoricalPointInTime
`func (o *DataQuality) UnsetHistoricalPointInTime()`

UnsetHistoricalPointInTime ensures that no value is present for HistoricalPointInTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


