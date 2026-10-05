# COTResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**Instrument** | **string** |  | 
**Source** | **string** |  | 
**SourceUrl** | **string** |  | 
**Provenance** | [**COTProvenance**](COTProvenance.md) |  | 
**FxOverlay** | [**COTFXOverlay**](COTFXOverlay.md) |  | 
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**LatestAvailableDate** | Pointer to **NullableString** |  | [optional] 
**LatestAvailableAnnouncementDatetime** | Pointer to **NullableInt64** |  | [optional] 
**ExpectedNextRelease** | Pointer to **NullableString** |  | [optional] 
**ExpectedNextReleaseEpoch** | Pointer to **NullableInt64** |  | [optional] 
**LastSyncStatus** | **string** |  | 
**LastSyncDatetime** | Pointer to **NullableString** |  | [optional] 
**DataLagDays** | Pointer to **NullableInt64** |  | [optional] 
**NextExpectedCftcReportDate** | Pointer to **NullableString** |  | [optional] 
**NextExpectedCftcReleaseDate** | Pointer to **NullableString** |  | [optional] 
**ExpectedNextReleaseHolidayAdjusted** | Pointer to **NullableBool** |  | [optional] 
**ExpectedNextReleaseSource** | Pointer to **NullableString** |  | [optional] 
**ExpectedNextReleaseSourceUrl** | Pointer to **NullableString** |  | [optional] 
**ExpectedNextReleaseScheduleStorage** | Pointer to **NullableString** |  | [optional] 
**ExpectedNextReleaseScheduleLastUpdated** | Pointer to **NullableString** |  | [optional] 
**RequestedWindowHasData** | Pointer to **NullableBool** |  | [optional] 
**RequestedWindowLatestDate** | Pointer to **NullableString** |  | [optional] 
**RequestedWindowIncludesLatestAvailable** | Pointer to **NullableBool** |  | [optional] 
**PageIncludesLatestAvailable** | Pointer to **NullableBool** |  | [optional] 
**LastUpdated** | Pointer to **NullableString** |  | [optional] 
**DataQuality** | [**DataQuality**](DataQuality.md) |  | 
**Pagination** | [**COTPaginationInfo**](COTPaginationInfo.md) |  | 
**Data** | [**[]COTDataPoint**](COTDataPoint.md) |  | 

## Methods

### NewCOTResponse

`func NewCOTResponse(currency string, instrument string, source string, sourceUrl string, provenance COTProvenance, fxOverlay COTFXOverlay, startDate string, endDate string, lastSyncStatus string, dataQuality DataQuality, pagination COTPaginationInfo, data []COTDataPoint, ) *COTResponse`

NewCOTResponse instantiates a new COTResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCOTResponseWithDefaults

`func NewCOTResponseWithDefaults() *COTResponse`

NewCOTResponseWithDefaults instantiates a new COTResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *COTResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *COTResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *COTResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetInstrument

`func (o *COTResponse) GetInstrument() string`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *COTResponse) GetInstrumentOk() (*string, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *COTResponse) SetInstrument(v string)`

SetInstrument sets Instrument field to given value.


### GetSource

`func (o *COTResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *COTResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *COTResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceUrl

`func (o *COTResponse) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *COTResponse) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *COTResponse) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.


### GetProvenance

`func (o *COTResponse) GetProvenance() COTProvenance`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *COTResponse) GetProvenanceOk() (*COTProvenance, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *COTResponse) SetProvenance(v COTProvenance)`

SetProvenance sets Provenance field to given value.


### GetFxOverlay

`func (o *COTResponse) GetFxOverlay() COTFXOverlay`

GetFxOverlay returns the FxOverlay field if non-nil, zero value otherwise.

### GetFxOverlayOk

`func (o *COTResponse) GetFxOverlayOk() (*COTFXOverlay, bool)`

GetFxOverlayOk returns a tuple with the FxOverlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxOverlay

`func (o *COTResponse) SetFxOverlay(v COTFXOverlay)`

SetFxOverlay sets FxOverlay field to given value.


### GetStartDate

`func (o *COTResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *COTResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *COTResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *COTResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *COTResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *COTResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetLatestAvailableDate

`func (o *COTResponse) GetLatestAvailableDate() string`

GetLatestAvailableDate returns the LatestAvailableDate field if non-nil, zero value otherwise.

### GetLatestAvailableDateOk

`func (o *COTResponse) GetLatestAvailableDateOk() (*string, bool)`

GetLatestAvailableDateOk returns a tuple with the LatestAvailableDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAvailableDate

`func (o *COTResponse) SetLatestAvailableDate(v string)`

SetLatestAvailableDate sets LatestAvailableDate field to given value.

### HasLatestAvailableDate

`func (o *COTResponse) HasLatestAvailableDate() bool`

HasLatestAvailableDate returns a boolean if a field has been set.

### SetLatestAvailableDateNil

`func (o *COTResponse) SetLatestAvailableDateNil(b bool)`

 SetLatestAvailableDateNil sets the value for LatestAvailableDate to be an explicit nil

### UnsetLatestAvailableDate
`func (o *COTResponse) UnsetLatestAvailableDate()`

UnsetLatestAvailableDate ensures that no value is present for LatestAvailableDate, not even an explicit nil
### GetLatestAvailableAnnouncementDatetime

`func (o *COTResponse) GetLatestAvailableAnnouncementDatetime() int64`

GetLatestAvailableAnnouncementDatetime returns the LatestAvailableAnnouncementDatetime field if non-nil, zero value otherwise.

### GetLatestAvailableAnnouncementDatetimeOk

`func (o *COTResponse) GetLatestAvailableAnnouncementDatetimeOk() (*int64, bool)`

GetLatestAvailableAnnouncementDatetimeOk returns a tuple with the LatestAvailableAnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAvailableAnnouncementDatetime

`func (o *COTResponse) SetLatestAvailableAnnouncementDatetime(v int64)`

SetLatestAvailableAnnouncementDatetime sets LatestAvailableAnnouncementDatetime field to given value.

### HasLatestAvailableAnnouncementDatetime

`func (o *COTResponse) HasLatestAvailableAnnouncementDatetime() bool`

HasLatestAvailableAnnouncementDatetime returns a boolean if a field has been set.

### SetLatestAvailableAnnouncementDatetimeNil

`func (o *COTResponse) SetLatestAvailableAnnouncementDatetimeNil(b bool)`

 SetLatestAvailableAnnouncementDatetimeNil sets the value for LatestAvailableAnnouncementDatetime to be an explicit nil

### UnsetLatestAvailableAnnouncementDatetime
`func (o *COTResponse) UnsetLatestAvailableAnnouncementDatetime()`

UnsetLatestAvailableAnnouncementDatetime ensures that no value is present for LatestAvailableAnnouncementDatetime, not even an explicit nil
### GetExpectedNextRelease

`func (o *COTResponse) GetExpectedNextRelease() string`

GetExpectedNextRelease returns the ExpectedNextRelease field if non-nil, zero value otherwise.

### GetExpectedNextReleaseOk

`func (o *COTResponse) GetExpectedNextReleaseOk() (*string, bool)`

GetExpectedNextReleaseOk returns a tuple with the ExpectedNextRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedNextRelease

`func (o *COTResponse) SetExpectedNextRelease(v string)`

SetExpectedNextRelease sets ExpectedNextRelease field to given value.

### HasExpectedNextRelease

`func (o *COTResponse) HasExpectedNextRelease() bool`

HasExpectedNextRelease returns a boolean if a field has been set.

### SetExpectedNextReleaseNil

`func (o *COTResponse) SetExpectedNextReleaseNil(b bool)`

 SetExpectedNextReleaseNil sets the value for ExpectedNextRelease to be an explicit nil

### UnsetExpectedNextRelease
`func (o *COTResponse) UnsetExpectedNextRelease()`

UnsetExpectedNextRelease ensures that no value is present for ExpectedNextRelease, not even an explicit nil
### GetExpectedNextReleaseEpoch

`func (o *COTResponse) GetExpectedNextReleaseEpoch() int64`

GetExpectedNextReleaseEpoch returns the ExpectedNextReleaseEpoch field if non-nil, zero value otherwise.

### GetExpectedNextReleaseEpochOk

`func (o *COTResponse) GetExpectedNextReleaseEpochOk() (*int64, bool)`

GetExpectedNextReleaseEpochOk returns a tuple with the ExpectedNextReleaseEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedNextReleaseEpoch

`func (o *COTResponse) SetExpectedNextReleaseEpoch(v int64)`

SetExpectedNextReleaseEpoch sets ExpectedNextReleaseEpoch field to given value.

### HasExpectedNextReleaseEpoch

`func (o *COTResponse) HasExpectedNextReleaseEpoch() bool`

HasExpectedNextReleaseEpoch returns a boolean if a field has been set.

### SetExpectedNextReleaseEpochNil

`func (o *COTResponse) SetExpectedNextReleaseEpochNil(b bool)`

 SetExpectedNextReleaseEpochNil sets the value for ExpectedNextReleaseEpoch to be an explicit nil

### UnsetExpectedNextReleaseEpoch
`func (o *COTResponse) UnsetExpectedNextReleaseEpoch()`

UnsetExpectedNextReleaseEpoch ensures that no value is present for ExpectedNextReleaseEpoch, not even an explicit nil
### GetLastSyncStatus

`func (o *COTResponse) GetLastSyncStatus() string`

GetLastSyncStatus returns the LastSyncStatus field if non-nil, zero value otherwise.

### GetLastSyncStatusOk

`func (o *COTResponse) GetLastSyncStatusOk() (*string, bool)`

GetLastSyncStatusOk returns a tuple with the LastSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncStatus

`func (o *COTResponse) SetLastSyncStatus(v string)`

SetLastSyncStatus sets LastSyncStatus field to given value.


### GetLastSyncDatetime

`func (o *COTResponse) GetLastSyncDatetime() string`

GetLastSyncDatetime returns the LastSyncDatetime field if non-nil, zero value otherwise.

### GetLastSyncDatetimeOk

`func (o *COTResponse) GetLastSyncDatetimeOk() (*string, bool)`

GetLastSyncDatetimeOk returns a tuple with the LastSyncDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDatetime

`func (o *COTResponse) SetLastSyncDatetime(v string)`

SetLastSyncDatetime sets LastSyncDatetime field to given value.

### HasLastSyncDatetime

`func (o *COTResponse) HasLastSyncDatetime() bool`

HasLastSyncDatetime returns a boolean if a field has been set.

### SetLastSyncDatetimeNil

`func (o *COTResponse) SetLastSyncDatetimeNil(b bool)`

 SetLastSyncDatetimeNil sets the value for LastSyncDatetime to be an explicit nil

### UnsetLastSyncDatetime
`func (o *COTResponse) UnsetLastSyncDatetime()`

UnsetLastSyncDatetime ensures that no value is present for LastSyncDatetime, not even an explicit nil
### GetDataLagDays

`func (o *COTResponse) GetDataLagDays() int64`

GetDataLagDays returns the DataLagDays field if non-nil, zero value otherwise.

### GetDataLagDaysOk

`func (o *COTResponse) GetDataLagDaysOk() (*int64, bool)`

GetDataLagDaysOk returns a tuple with the DataLagDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLagDays

`func (o *COTResponse) SetDataLagDays(v int64)`

SetDataLagDays sets DataLagDays field to given value.

### HasDataLagDays

`func (o *COTResponse) HasDataLagDays() bool`

HasDataLagDays returns a boolean if a field has been set.

### SetDataLagDaysNil

`func (o *COTResponse) SetDataLagDaysNil(b bool)`

 SetDataLagDaysNil sets the value for DataLagDays to be an explicit nil

### UnsetDataLagDays
`func (o *COTResponse) UnsetDataLagDays()`

UnsetDataLagDays ensures that no value is present for DataLagDays, not even an explicit nil
### GetNextExpectedCftcReportDate

`func (o *COTResponse) GetNextExpectedCftcReportDate() string`

GetNextExpectedCftcReportDate returns the NextExpectedCftcReportDate field if non-nil, zero value otherwise.

### GetNextExpectedCftcReportDateOk

`func (o *COTResponse) GetNextExpectedCftcReportDateOk() (*string, bool)`

GetNextExpectedCftcReportDateOk returns a tuple with the NextExpectedCftcReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExpectedCftcReportDate

`func (o *COTResponse) SetNextExpectedCftcReportDate(v string)`

SetNextExpectedCftcReportDate sets NextExpectedCftcReportDate field to given value.

### HasNextExpectedCftcReportDate

`func (o *COTResponse) HasNextExpectedCftcReportDate() bool`

HasNextExpectedCftcReportDate returns a boolean if a field has been set.

### SetNextExpectedCftcReportDateNil

`func (o *COTResponse) SetNextExpectedCftcReportDateNil(b bool)`

 SetNextExpectedCftcReportDateNil sets the value for NextExpectedCftcReportDate to be an explicit nil

### UnsetNextExpectedCftcReportDate
`func (o *COTResponse) UnsetNextExpectedCftcReportDate()`

UnsetNextExpectedCftcReportDate ensures that no value is present for NextExpectedCftcReportDate, not even an explicit nil
### GetNextExpectedCftcReleaseDate

`func (o *COTResponse) GetNextExpectedCftcReleaseDate() string`

GetNextExpectedCftcReleaseDate returns the NextExpectedCftcReleaseDate field if non-nil, zero value otherwise.

### GetNextExpectedCftcReleaseDateOk

`func (o *COTResponse) GetNextExpectedCftcReleaseDateOk() (*string, bool)`

GetNextExpectedCftcReleaseDateOk returns a tuple with the NextExpectedCftcReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExpectedCftcReleaseDate

`func (o *COTResponse) SetNextExpectedCftcReleaseDate(v string)`

SetNextExpectedCftcReleaseDate sets NextExpectedCftcReleaseDate field to given value.

### HasNextExpectedCftcReleaseDate

`func (o *COTResponse) HasNextExpectedCftcReleaseDate() bool`

HasNextExpectedCftcReleaseDate returns a boolean if a field has been set.

### SetNextExpectedCftcReleaseDateNil

`func (o *COTResponse) SetNextExpectedCftcReleaseDateNil(b bool)`

 SetNextExpectedCftcReleaseDateNil sets the value for NextExpectedCftcReleaseDate to be an explicit nil

### UnsetNextExpectedCftcReleaseDate
`func (o *COTResponse) UnsetNextExpectedCftcReleaseDate()`

UnsetNextExpectedCftcReleaseDate ensures that no value is present for NextExpectedCftcReleaseDate, not even an explicit nil
### GetExpectedNextReleaseHolidayAdjusted

`func (o *COTResponse) GetExpectedNextReleaseHolidayAdjusted() bool`

GetExpectedNextReleaseHolidayAdjusted returns the ExpectedNextReleaseHolidayAdjusted field if non-nil, zero value otherwise.

### GetExpectedNextReleaseHolidayAdjustedOk

`func (o *COTResponse) GetExpectedNextReleaseHolidayAdjustedOk() (*bool, bool)`

GetExpectedNextReleaseHolidayAdjustedOk returns a tuple with the ExpectedNextReleaseHolidayAdjusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedNextReleaseHolidayAdjusted

`func (o *COTResponse) SetExpectedNextReleaseHolidayAdjusted(v bool)`

SetExpectedNextReleaseHolidayAdjusted sets ExpectedNextReleaseHolidayAdjusted field to given value.

### HasExpectedNextReleaseHolidayAdjusted

`func (o *COTResponse) HasExpectedNextReleaseHolidayAdjusted() bool`

HasExpectedNextReleaseHolidayAdjusted returns a boolean if a field has been set.

### SetExpectedNextReleaseHolidayAdjustedNil

`func (o *COTResponse) SetExpectedNextReleaseHolidayAdjustedNil(b bool)`

 SetExpectedNextReleaseHolidayAdjustedNil sets the value for ExpectedNextReleaseHolidayAdjusted to be an explicit nil

### UnsetExpectedNextReleaseHolidayAdjusted
`func (o *COTResponse) UnsetExpectedNextReleaseHolidayAdjusted()`

UnsetExpectedNextReleaseHolidayAdjusted ensures that no value is present for ExpectedNextReleaseHolidayAdjusted, not even an explicit nil
### GetExpectedNextReleaseSource

`func (o *COTResponse) GetExpectedNextReleaseSource() string`

GetExpectedNextReleaseSource returns the ExpectedNextReleaseSource field if non-nil, zero value otherwise.

### GetExpectedNextReleaseSourceOk

`func (o *COTResponse) GetExpectedNextReleaseSourceOk() (*string, bool)`

GetExpectedNextReleaseSourceOk returns a tuple with the ExpectedNextReleaseSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedNextReleaseSource

`func (o *COTResponse) SetExpectedNextReleaseSource(v string)`

SetExpectedNextReleaseSource sets ExpectedNextReleaseSource field to given value.

### HasExpectedNextReleaseSource

`func (o *COTResponse) HasExpectedNextReleaseSource() bool`

HasExpectedNextReleaseSource returns a boolean if a field has been set.

### SetExpectedNextReleaseSourceNil

`func (o *COTResponse) SetExpectedNextReleaseSourceNil(b bool)`

 SetExpectedNextReleaseSourceNil sets the value for ExpectedNextReleaseSource to be an explicit nil

### UnsetExpectedNextReleaseSource
`func (o *COTResponse) UnsetExpectedNextReleaseSource()`

UnsetExpectedNextReleaseSource ensures that no value is present for ExpectedNextReleaseSource, not even an explicit nil
### GetExpectedNextReleaseSourceUrl

`func (o *COTResponse) GetExpectedNextReleaseSourceUrl() string`

GetExpectedNextReleaseSourceUrl returns the ExpectedNextReleaseSourceUrl field if non-nil, zero value otherwise.

### GetExpectedNextReleaseSourceUrlOk

`func (o *COTResponse) GetExpectedNextReleaseSourceUrlOk() (*string, bool)`

GetExpectedNextReleaseSourceUrlOk returns a tuple with the ExpectedNextReleaseSourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedNextReleaseSourceUrl

`func (o *COTResponse) SetExpectedNextReleaseSourceUrl(v string)`

SetExpectedNextReleaseSourceUrl sets ExpectedNextReleaseSourceUrl field to given value.

### HasExpectedNextReleaseSourceUrl

`func (o *COTResponse) HasExpectedNextReleaseSourceUrl() bool`

HasExpectedNextReleaseSourceUrl returns a boolean if a field has been set.

### SetExpectedNextReleaseSourceUrlNil

`func (o *COTResponse) SetExpectedNextReleaseSourceUrlNil(b bool)`

 SetExpectedNextReleaseSourceUrlNil sets the value for ExpectedNextReleaseSourceUrl to be an explicit nil

### UnsetExpectedNextReleaseSourceUrl
`func (o *COTResponse) UnsetExpectedNextReleaseSourceUrl()`

UnsetExpectedNextReleaseSourceUrl ensures that no value is present for ExpectedNextReleaseSourceUrl, not even an explicit nil
### GetExpectedNextReleaseScheduleStorage

`func (o *COTResponse) GetExpectedNextReleaseScheduleStorage() string`

GetExpectedNextReleaseScheduleStorage returns the ExpectedNextReleaseScheduleStorage field if non-nil, zero value otherwise.

### GetExpectedNextReleaseScheduleStorageOk

`func (o *COTResponse) GetExpectedNextReleaseScheduleStorageOk() (*string, bool)`

GetExpectedNextReleaseScheduleStorageOk returns a tuple with the ExpectedNextReleaseScheduleStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedNextReleaseScheduleStorage

`func (o *COTResponse) SetExpectedNextReleaseScheduleStorage(v string)`

SetExpectedNextReleaseScheduleStorage sets ExpectedNextReleaseScheduleStorage field to given value.

### HasExpectedNextReleaseScheduleStorage

`func (o *COTResponse) HasExpectedNextReleaseScheduleStorage() bool`

HasExpectedNextReleaseScheduleStorage returns a boolean if a field has been set.

### SetExpectedNextReleaseScheduleStorageNil

`func (o *COTResponse) SetExpectedNextReleaseScheduleStorageNil(b bool)`

 SetExpectedNextReleaseScheduleStorageNil sets the value for ExpectedNextReleaseScheduleStorage to be an explicit nil

### UnsetExpectedNextReleaseScheduleStorage
`func (o *COTResponse) UnsetExpectedNextReleaseScheduleStorage()`

UnsetExpectedNextReleaseScheduleStorage ensures that no value is present for ExpectedNextReleaseScheduleStorage, not even an explicit nil
### GetExpectedNextReleaseScheduleLastUpdated

`func (o *COTResponse) GetExpectedNextReleaseScheduleLastUpdated() string`

GetExpectedNextReleaseScheduleLastUpdated returns the ExpectedNextReleaseScheduleLastUpdated field if non-nil, zero value otherwise.

### GetExpectedNextReleaseScheduleLastUpdatedOk

`func (o *COTResponse) GetExpectedNextReleaseScheduleLastUpdatedOk() (*string, bool)`

GetExpectedNextReleaseScheduleLastUpdatedOk returns a tuple with the ExpectedNextReleaseScheduleLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedNextReleaseScheduleLastUpdated

`func (o *COTResponse) SetExpectedNextReleaseScheduleLastUpdated(v string)`

SetExpectedNextReleaseScheduleLastUpdated sets ExpectedNextReleaseScheduleLastUpdated field to given value.

### HasExpectedNextReleaseScheduleLastUpdated

`func (o *COTResponse) HasExpectedNextReleaseScheduleLastUpdated() bool`

HasExpectedNextReleaseScheduleLastUpdated returns a boolean if a field has been set.

### SetExpectedNextReleaseScheduleLastUpdatedNil

`func (o *COTResponse) SetExpectedNextReleaseScheduleLastUpdatedNil(b bool)`

 SetExpectedNextReleaseScheduleLastUpdatedNil sets the value for ExpectedNextReleaseScheduleLastUpdated to be an explicit nil

### UnsetExpectedNextReleaseScheduleLastUpdated
`func (o *COTResponse) UnsetExpectedNextReleaseScheduleLastUpdated()`

UnsetExpectedNextReleaseScheduleLastUpdated ensures that no value is present for ExpectedNextReleaseScheduleLastUpdated, not even an explicit nil
### GetRequestedWindowHasData

`func (o *COTResponse) GetRequestedWindowHasData() bool`

GetRequestedWindowHasData returns the RequestedWindowHasData field if non-nil, zero value otherwise.

### GetRequestedWindowHasDataOk

`func (o *COTResponse) GetRequestedWindowHasDataOk() (*bool, bool)`

GetRequestedWindowHasDataOk returns a tuple with the RequestedWindowHasData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedWindowHasData

`func (o *COTResponse) SetRequestedWindowHasData(v bool)`

SetRequestedWindowHasData sets RequestedWindowHasData field to given value.

### HasRequestedWindowHasData

`func (o *COTResponse) HasRequestedWindowHasData() bool`

HasRequestedWindowHasData returns a boolean if a field has been set.

### SetRequestedWindowHasDataNil

`func (o *COTResponse) SetRequestedWindowHasDataNil(b bool)`

 SetRequestedWindowHasDataNil sets the value for RequestedWindowHasData to be an explicit nil

### UnsetRequestedWindowHasData
`func (o *COTResponse) UnsetRequestedWindowHasData()`

UnsetRequestedWindowHasData ensures that no value is present for RequestedWindowHasData, not even an explicit nil
### GetRequestedWindowLatestDate

`func (o *COTResponse) GetRequestedWindowLatestDate() string`

GetRequestedWindowLatestDate returns the RequestedWindowLatestDate field if non-nil, zero value otherwise.

### GetRequestedWindowLatestDateOk

`func (o *COTResponse) GetRequestedWindowLatestDateOk() (*string, bool)`

GetRequestedWindowLatestDateOk returns a tuple with the RequestedWindowLatestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedWindowLatestDate

`func (o *COTResponse) SetRequestedWindowLatestDate(v string)`

SetRequestedWindowLatestDate sets RequestedWindowLatestDate field to given value.

### HasRequestedWindowLatestDate

`func (o *COTResponse) HasRequestedWindowLatestDate() bool`

HasRequestedWindowLatestDate returns a boolean if a field has been set.

### SetRequestedWindowLatestDateNil

`func (o *COTResponse) SetRequestedWindowLatestDateNil(b bool)`

 SetRequestedWindowLatestDateNil sets the value for RequestedWindowLatestDate to be an explicit nil

### UnsetRequestedWindowLatestDate
`func (o *COTResponse) UnsetRequestedWindowLatestDate()`

UnsetRequestedWindowLatestDate ensures that no value is present for RequestedWindowLatestDate, not even an explicit nil
### GetRequestedWindowIncludesLatestAvailable

`func (o *COTResponse) GetRequestedWindowIncludesLatestAvailable() bool`

GetRequestedWindowIncludesLatestAvailable returns the RequestedWindowIncludesLatestAvailable field if non-nil, zero value otherwise.

### GetRequestedWindowIncludesLatestAvailableOk

`func (o *COTResponse) GetRequestedWindowIncludesLatestAvailableOk() (*bool, bool)`

GetRequestedWindowIncludesLatestAvailableOk returns a tuple with the RequestedWindowIncludesLatestAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedWindowIncludesLatestAvailable

`func (o *COTResponse) SetRequestedWindowIncludesLatestAvailable(v bool)`

SetRequestedWindowIncludesLatestAvailable sets RequestedWindowIncludesLatestAvailable field to given value.

### HasRequestedWindowIncludesLatestAvailable

`func (o *COTResponse) HasRequestedWindowIncludesLatestAvailable() bool`

HasRequestedWindowIncludesLatestAvailable returns a boolean if a field has been set.

### SetRequestedWindowIncludesLatestAvailableNil

`func (o *COTResponse) SetRequestedWindowIncludesLatestAvailableNil(b bool)`

 SetRequestedWindowIncludesLatestAvailableNil sets the value for RequestedWindowIncludesLatestAvailable to be an explicit nil

### UnsetRequestedWindowIncludesLatestAvailable
`func (o *COTResponse) UnsetRequestedWindowIncludesLatestAvailable()`

UnsetRequestedWindowIncludesLatestAvailable ensures that no value is present for RequestedWindowIncludesLatestAvailable, not even an explicit nil
### GetPageIncludesLatestAvailable

`func (o *COTResponse) GetPageIncludesLatestAvailable() bool`

GetPageIncludesLatestAvailable returns the PageIncludesLatestAvailable field if non-nil, zero value otherwise.

### GetPageIncludesLatestAvailableOk

`func (o *COTResponse) GetPageIncludesLatestAvailableOk() (*bool, bool)`

GetPageIncludesLatestAvailableOk returns a tuple with the PageIncludesLatestAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIncludesLatestAvailable

`func (o *COTResponse) SetPageIncludesLatestAvailable(v bool)`

SetPageIncludesLatestAvailable sets PageIncludesLatestAvailable field to given value.

### HasPageIncludesLatestAvailable

`func (o *COTResponse) HasPageIncludesLatestAvailable() bool`

HasPageIncludesLatestAvailable returns a boolean if a field has been set.

### SetPageIncludesLatestAvailableNil

`func (o *COTResponse) SetPageIncludesLatestAvailableNil(b bool)`

 SetPageIncludesLatestAvailableNil sets the value for PageIncludesLatestAvailable to be an explicit nil

### UnsetPageIncludesLatestAvailable
`func (o *COTResponse) UnsetPageIncludesLatestAvailable()`

UnsetPageIncludesLatestAvailable ensures that no value is present for PageIncludesLatestAvailable, not even an explicit nil
### GetLastUpdated

`func (o *COTResponse) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *COTResponse) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *COTResponse) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *COTResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *COTResponse) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *COTResponse) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetDataQuality

`func (o *COTResponse) GetDataQuality() DataQuality`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *COTResponse) GetDataQualityOk() (*DataQuality, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *COTResponse) SetDataQuality(v DataQuality)`

SetDataQuality sets DataQuality field to given value.


### GetPagination

`func (o *COTResponse) GetPagination() COTPaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *COTResponse) GetPaginationOk() (*COTPaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *COTResponse) SetPagination(v COTPaginationInfo)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *COTResponse) GetData() []COTDataPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *COTResponse) GetDataOk() (*[]COTDataPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *COTResponse) SetData(v []COTDataPoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


