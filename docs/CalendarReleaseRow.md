# CalendarReleaseRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnouncementDatetime** | **int64** |  | 
**Release** | **string** |  | 
**CalendarEventId** | Pointer to **NullableString** |  | [optional] 
**AnnouncementDatetimeUtc** | Pointer to **NullableString** |  | [optional] 
**AnnouncementDatetimeLocal** | Pointer to **NullableString** |  | [optional] 
**AnnouncementDatetimeRequestedTimezone** | Pointer to **NullableString** |  | [optional] 
**ReleaseDateConfirmed** | Pointer to **NullableBool** |  | [optional] 
**ReleaseTimeAssumed** | Pointer to **NullableBool** |  | [optional] 
**ReleaseTimeStatus** | Pointer to **NullableString** |  | [optional] 
**ReleaseTimeAssumption** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceUrl** | Pointer to **NullableString** |  | [optional] 
**SourceRelease** | Pointer to **NullableString** |  | [optional] 
**ScheduleUrl** | Pointer to **NullableString** |  | [optional] 
**ReleaseStage** | Pointer to **NullableString** |  | [optional] 
**ReferencePeriod** | Pointer to **NullableString** |  | [optional] 
**DatasetCodes** | Pointer to **interface{}** |  | [optional] 
**Date** | Pointer to **NullableString** |  | [optional] 
**Domain** | Pointer to **NullableString** |  | [optional] 
**DataCurrency** | Pointer to **NullableString** |  | [optional] 
**EndpointFamily** | Pointer to **NullableString** |  | [optional] 
**EndpointPath** | Pointer to **NullableString** |  | [optional] 
**RequiresApiKey** | Pointer to **NullableBool** |  | [optional] 
**EventImportance** | Pointer to **NullableString** |  | [optional] 
**MarketTier** | Pointer to **NullableInt64** |  | [optional] 
**TopTierForCurrency** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCalendarReleaseRow

`func NewCalendarReleaseRow(announcementDatetime int64, release string, ) *CalendarReleaseRow`

NewCalendarReleaseRow instantiates a new CalendarReleaseRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarReleaseRowWithDefaults

`func NewCalendarReleaseRowWithDefaults() *CalendarReleaseRow`

NewCalendarReleaseRowWithDefaults instantiates a new CalendarReleaseRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnouncementDatetime

`func (o *CalendarReleaseRow) GetAnnouncementDatetime() int64`

GetAnnouncementDatetime returns the AnnouncementDatetime field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeOk

`func (o *CalendarReleaseRow) GetAnnouncementDatetimeOk() (*int64, bool)`

GetAnnouncementDatetimeOk returns a tuple with the AnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetime

`func (o *CalendarReleaseRow) SetAnnouncementDatetime(v int64)`

SetAnnouncementDatetime sets AnnouncementDatetime field to given value.


### GetRelease

`func (o *CalendarReleaseRow) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *CalendarReleaseRow) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *CalendarReleaseRow) SetRelease(v string)`

SetRelease sets Release field to given value.


### GetCalendarEventId

`func (o *CalendarReleaseRow) GetCalendarEventId() string`

GetCalendarEventId returns the CalendarEventId field if non-nil, zero value otherwise.

### GetCalendarEventIdOk

`func (o *CalendarReleaseRow) GetCalendarEventIdOk() (*string, bool)`

GetCalendarEventIdOk returns a tuple with the CalendarEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarEventId

`func (o *CalendarReleaseRow) SetCalendarEventId(v string)`

SetCalendarEventId sets CalendarEventId field to given value.

### HasCalendarEventId

`func (o *CalendarReleaseRow) HasCalendarEventId() bool`

HasCalendarEventId returns a boolean if a field has been set.

### SetCalendarEventIdNil

`func (o *CalendarReleaseRow) SetCalendarEventIdNil(b bool)`

 SetCalendarEventIdNil sets the value for CalendarEventId to be an explicit nil

### UnsetCalendarEventId
`func (o *CalendarReleaseRow) UnsetCalendarEventId()`

UnsetCalendarEventId ensures that no value is present for CalendarEventId, not even an explicit nil
### GetAnnouncementDatetimeUtc

`func (o *CalendarReleaseRow) GetAnnouncementDatetimeUtc() string`

GetAnnouncementDatetimeUtc returns the AnnouncementDatetimeUtc field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeUtcOk

`func (o *CalendarReleaseRow) GetAnnouncementDatetimeUtcOk() (*string, bool)`

GetAnnouncementDatetimeUtcOk returns a tuple with the AnnouncementDatetimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetimeUtc

`func (o *CalendarReleaseRow) SetAnnouncementDatetimeUtc(v string)`

SetAnnouncementDatetimeUtc sets AnnouncementDatetimeUtc field to given value.

### HasAnnouncementDatetimeUtc

`func (o *CalendarReleaseRow) HasAnnouncementDatetimeUtc() bool`

HasAnnouncementDatetimeUtc returns a boolean if a field has been set.

### SetAnnouncementDatetimeUtcNil

`func (o *CalendarReleaseRow) SetAnnouncementDatetimeUtcNil(b bool)`

 SetAnnouncementDatetimeUtcNil sets the value for AnnouncementDatetimeUtc to be an explicit nil

### UnsetAnnouncementDatetimeUtc
`func (o *CalendarReleaseRow) UnsetAnnouncementDatetimeUtc()`

UnsetAnnouncementDatetimeUtc ensures that no value is present for AnnouncementDatetimeUtc, not even an explicit nil
### GetAnnouncementDatetimeLocal

`func (o *CalendarReleaseRow) GetAnnouncementDatetimeLocal() string`

GetAnnouncementDatetimeLocal returns the AnnouncementDatetimeLocal field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeLocalOk

`func (o *CalendarReleaseRow) GetAnnouncementDatetimeLocalOk() (*string, bool)`

GetAnnouncementDatetimeLocalOk returns a tuple with the AnnouncementDatetimeLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetimeLocal

`func (o *CalendarReleaseRow) SetAnnouncementDatetimeLocal(v string)`

SetAnnouncementDatetimeLocal sets AnnouncementDatetimeLocal field to given value.

### HasAnnouncementDatetimeLocal

`func (o *CalendarReleaseRow) HasAnnouncementDatetimeLocal() bool`

HasAnnouncementDatetimeLocal returns a boolean if a field has been set.

### SetAnnouncementDatetimeLocalNil

`func (o *CalendarReleaseRow) SetAnnouncementDatetimeLocalNil(b bool)`

 SetAnnouncementDatetimeLocalNil sets the value for AnnouncementDatetimeLocal to be an explicit nil

### UnsetAnnouncementDatetimeLocal
`func (o *CalendarReleaseRow) UnsetAnnouncementDatetimeLocal()`

UnsetAnnouncementDatetimeLocal ensures that no value is present for AnnouncementDatetimeLocal, not even an explicit nil
### GetAnnouncementDatetimeRequestedTimezone

`func (o *CalendarReleaseRow) GetAnnouncementDatetimeRequestedTimezone() string`

GetAnnouncementDatetimeRequestedTimezone returns the AnnouncementDatetimeRequestedTimezone field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeRequestedTimezoneOk

`func (o *CalendarReleaseRow) GetAnnouncementDatetimeRequestedTimezoneOk() (*string, bool)`

GetAnnouncementDatetimeRequestedTimezoneOk returns a tuple with the AnnouncementDatetimeRequestedTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetimeRequestedTimezone

`func (o *CalendarReleaseRow) SetAnnouncementDatetimeRequestedTimezone(v string)`

SetAnnouncementDatetimeRequestedTimezone sets AnnouncementDatetimeRequestedTimezone field to given value.

### HasAnnouncementDatetimeRequestedTimezone

`func (o *CalendarReleaseRow) HasAnnouncementDatetimeRequestedTimezone() bool`

HasAnnouncementDatetimeRequestedTimezone returns a boolean if a field has been set.

### SetAnnouncementDatetimeRequestedTimezoneNil

`func (o *CalendarReleaseRow) SetAnnouncementDatetimeRequestedTimezoneNil(b bool)`

 SetAnnouncementDatetimeRequestedTimezoneNil sets the value for AnnouncementDatetimeRequestedTimezone to be an explicit nil

### UnsetAnnouncementDatetimeRequestedTimezone
`func (o *CalendarReleaseRow) UnsetAnnouncementDatetimeRequestedTimezone()`

UnsetAnnouncementDatetimeRequestedTimezone ensures that no value is present for AnnouncementDatetimeRequestedTimezone, not even an explicit nil
### GetReleaseDateConfirmed

`func (o *CalendarReleaseRow) GetReleaseDateConfirmed() bool`

GetReleaseDateConfirmed returns the ReleaseDateConfirmed field if non-nil, zero value otherwise.

### GetReleaseDateConfirmedOk

`func (o *CalendarReleaseRow) GetReleaseDateConfirmedOk() (*bool, bool)`

GetReleaseDateConfirmedOk returns a tuple with the ReleaseDateConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDateConfirmed

`func (o *CalendarReleaseRow) SetReleaseDateConfirmed(v bool)`

SetReleaseDateConfirmed sets ReleaseDateConfirmed field to given value.

### HasReleaseDateConfirmed

`func (o *CalendarReleaseRow) HasReleaseDateConfirmed() bool`

HasReleaseDateConfirmed returns a boolean if a field has been set.

### SetReleaseDateConfirmedNil

`func (o *CalendarReleaseRow) SetReleaseDateConfirmedNil(b bool)`

 SetReleaseDateConfirmedNil sets the value for ReleaseDateConfirmed to be an explicit nil

### UnsetReleaseDateConfirmed
`func (o *CalendarReleaseRow) UnsetReleaseDateConfirmed()`

UnsetReleaseDateConfirmed ensures that no value is present for ReleaseDateConfirmed, not even an explicit nil
### GetReleaseTimeAssumed

`func (o *CalendarReleaseRow) GetReleaseTimeAssumed() bool`

GetReleaseTimeAssumed returns the ReleaseTimeAssumed field if non-nil, zero value otherwise.

### GetReleaseTimeAssumedOk

`func (o *CalendarReleaseRow) GetReleaseTimeAssumedOk() (*bool, bool)`

GetReleaseTimeAssumedOk returns a tuple with the ReleaseTimeAssumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTimeAssumed

`func (o *CalendarReleaseRow) SetReleaseTimeAssumed(v bool)`

SetReleaseTimeAssumed sets ReleaseTimeAssumed field to given value.

### HasReleaseTimeAssumed

`func (o *CalendarReleaseRow) HasReleaseTimeAssumed() bool`

HasReleaseTimeAssumed returns a boolean if a field has been set.

### SetReleaseTimeAssumedNil

`func (o *CalendarReleaseRow) SetReleaseTimeAssumedNil(b bool)`

 SetReleaseTimeAssumedNil sets the value for ReleaseTimeAssumed to be an explicit nil

### UnsetReleaseTimeAssumed
`func (o *CalendarReleaseRow) UnsetReleaseTimeAssumed()`

UnsetReleaseTimeAssumed ensures that no value is present for ReleaseTimeAssumed, not even an explicit nil
### GetReleaseTimeStatus

`func (o *CalendarReleaseRow) GetReleaseTimeStatus() string`

GetReleaseTimeStatus returns the ReleaseTimeStatus field if non-nil, zero value otherwise.

### GetReleaseTimeStatusOk

`func (o *CalendarReleaseRow) GetReleaseTimeStatusOk() (*string, bool)`

GetReleaseTimeStatusOk returns a tuple with the ReleaseTimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTimeStatus

`func (o *CalendarReleaseRow) SetReleaseTimeStatus(v string)`

SetReleaseTimeStatus sets ReleaseTimeStatus field to given value.

### HasReleaseTimeStatus

`func (o *CalendarReleaseRow) HasReleaseTimeStatus() bool`

HasReleaseTimeStatus returns a boolean if a field has been set.

### SetReleaseTimeStatusNil

`func (o *CalendarReleaseRow) SetReleaseTimeStatusNil(b bool)`

 SetReleaseTimeStatusNil sets the value for ReleaseTimeStatus to be an explicit nil

### UnsetReleaseTimeStatus
`func (o *CalendarReleaseRow) UnsetReleaseTimeStatus()`

UnsetReleaseTimeStatus ensures that no value is present for ReleaseTimeStatus, not even an explicit nil
### GetReleaseTimeAssumption

`func (o *CalendarReleaseRow) GetReleaseTimeAssumption() string`

GetReleaseTimeAssumption returns the ReleaseTimeAssumption field if non-nil, zero value otherwise.

### GetReleaseTimeAssumptionOk

`func (o *CalendarReleaseRow) GetReleaseTimeAssumptionOk() (*string, bool)`

GetReleaseTimeAssumptionOk returns a tuple with the ReleaseTimeAssumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTimeAssumption

`func (o *CalendarReleaseRow) SetReleaseTimeAssumption(v string)`

SetReleaseTimeAssumption sets ReleaseTimeAssumption field to given value.

### HasReleaseTimeAssumption

`func (o *CalendarReleaseRow) HasReleaseTimeAssumption() bool`

HasReleaseTimeAssumption returns a boolean if a field has been set.

### SetReleaseTimeAssumptionNil

`func (o *CalendarReleaseRow) SetReleaseTimeAssumptionNil(b bool)`

 SetReleaseTimeAssumptionNil sets the value for ReleaseTimeAssumption to be an explicit nil

### UnsetReleaseTimeAssumption
`func (o *CalendarReleaseRow) UnsetReleaseTimeAssumption()`

UnsetReleaseTimeAssumption ensures that no value is present for ReleaseTimeAssumption, not even an explicit nil
### GetName

`func (o *CalendarReleaseRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CalendarReleaseRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CalendarReleaseRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CalendarReleaseRow) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CalendarReleaseRow) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CalendarReleaseRow) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSource

`func (o *CalendarReleaseRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CalendarReleaseRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CalendarReleaseRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CalendarReleaseRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *CalendarReleaseRow) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *CalendarReleaseRow) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceUrl

`func (o *CalendarReleaseRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *CalendarReleaseRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *CalendarReleaseRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *CalendarReleaseRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *CalendarReleaseRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *CalendarReleaseRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSourceRelease

`func (o *CalendarReleaseRow) GetSourceRelease() string`

GetSourceRelease returns the SourceRelease field if non-nil, zero value otherwise.

### GetSourceReleaseOk

`func (o *CalendarReleaseRow) GetSourceReleaseOk() (*string, bool)`

GetSourceReleaseOk returns a tuple with the SourceRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRelease

`func (o *CalendarReleaseRow) SetSourceRelease(v string)`

SetSourceRelease sets SourceRelease field to given value.

### HasSourceRelease

`func (o *CalendarReleaseRow) HasSourceRelease() bool`

HasSourceRelease returns a boolean if a field has been set.

### SetSourceReleaseNil

`func (o *CalendarReleaseRow) SetSourceReleaseNil(b bool)`

 SetSourceReleaseNil sets the value for SourceRelease to be an explicit nil

### UnsetSourceRelease
`func (o *CalendarReleaseRow) UnsetSourceRelease()`

UnsetSourceRelease ensures that no value is present for SourceRelease, not even an explicit nil
### GetScheduleUrl

`func (o *CalendarReleaseRow) GetScheduleUrl() string`

GetScheduleUrl returns the ScheduleUrl field if non-nil, zero value otherwise.

### GetScheduleUrlOk

`func (o *CalendarReleaseRow) GetScheduleUrlOk() (*string, bool)`

GetScheduleUrlOk returns a tuple with the ScheduleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUrl

`func (o *CalendarReleaseRow) SetScheduleUrl(v string)`

SetScheduleUrl sets ScheduleUrl field to given value.

### HasScheduleUrl

`func (o *CalendarReleaseRow) HasScheduleUrl() bool`

HasScheduleUrl returns a boolean if a field has been set.

### SetScheduleUrlNil

`func (o *CalendarReleaseRow) SetScheduleUrlNil(b bool)`

 SetScheduleUrlNil sets the value for ScheduleUrl to be an explicit nil

### UnsetScheduleUrl
`func (o *CalendarReleaseRow) UnsetScheduleUrl()`

UnsetScheduleUrl ensures that no value is present for ScheduleUrl, not even an explicit nil
### GetReleaseStage

`func (o *CalendarReleaseRow) GetReleaseStage() string`

GetReleaseStage returns the ReleaseStage field if non-nil, zero value otherwise.

### GetReleaseStageOk

`func (o *CalendarReleaseRow) GetReleaseStageOk() (*string, bool)`

GetReleaseStageOk returns a tuple with the ReleaseStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStage

`func (o *CalendarReleaseRow) SetReleaseStage(v string)`

SetReleaseStage sets ReleaseStage field to given value.

### HasReleaseStage

`func (o *CalendarReleaseRow) HasReleaseStage() bool`

HasReleaseStage returns a boolean if a field has been set.

### SetReleaseStageNil

`func (o *CalendarReleaseRow) SetReleaseStageNil(b bool)`

 SetReleaseStageNil sets the value for ReleaseStage to be an explicit nil

### UnsetReleaseStage
`func (o *CalendarReleaseRow) UnsetReleaseStage()`

UnsetReleaseStage ensures that no value is present for ReleaseStage, not even an explicit nil
### GetReferencePeriod

`func (o *CalendarReleaseRow) GetReferencePeriod() string`

GetReferencePeriod returns the ReferencePeriod field if non-nil, zero value otherwise.

### GetReferencePeriodOk

`func (o *CalendarReleaseRow) GetReferencePeriodOk() (*string, bool)`

GetReferencePeriodOk returns a tuple with the ReferencePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencePeriod

`func (o *CalendarReleaseRow) SetReferencePeriod(v string)`

SetReferencePeriod sets ReferencePeriod field to given value.

### HasReferencePeriod

`func (o *CalendarReleaseRow) HasReferencePeriod() bool`

HasReferencePeriod returns a boolean if a field has been set.

### SetReferencePeriodNil

`func (o *CalendarReleaseRow) SetReferencePeriodNil(b bool)`

 SetReferencePeriodNil sets the value for ReferencePeriod to be an explicit nil

### UnsetReferencePeriod
`func (o *CalendarReleaseRow) UnsetReferencePeriod()`

UnsetReferencePeriod ensures that no value is present for ReferencePeriod, not even an explicit nil
### GetDatasetCodes

`func (o *CalendarReleaseRow) GetDatasetCodes() interface{}`

GetDatasetCodes returns the DatasetCodes field if non-nil, zero value otherwise.

### GetDatasetCodesOk

`func (o *CalendarReleaseRow) GetDatasetCodesOk() (*interface{}, bool)`

GetDatasetCodesOk returns a tuple with the DatasetCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetCodes

`func (o *CalendarReleaseRow) SetDatasetCodes(v interface{})`

SetDatasetCodes sets DatasetCodes field to given value.

### HasDatasetCodes

`func (o *CalendarReleaseRow) HasDatasetCodes() bool`

HasDatasetCodes returns a boolean if a field has been set.

### SetDatasetCodesNil

`func (o *CalendarReleaseRow) SetDatasetCodesNil(b bool)`

 SetDatasetCodesNil sets the value for DatasetCodes to be an explicit nil

### UnsetDatasetCodes
`func (o *CalendarReleaseRow) UnsetDatasetCodes()`

UnsetDatasetCodes ensures that no value is present for DatasetCodes, not even an explicit nil
### GetDate

`func (o *CalendarReleaseRow) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalendarReleaseRow) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalendarReleaseRow) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CalendarReleaseRow) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *CalendarReleaseRow) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *CalendarReleaseRow) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDomain

`func (o *CalendarReleaseRow) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CalendarReleaseRow) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CalendarReleaseRow) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CalendarReleaseRow) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *CalendarReleaseRow) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *CalendarReleaseRow) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetDataCurrency

`func (o *CalendarReleaseRow) GetDataCurrency() string`

GetDataCurrency returns the DataCurrency field if non-nil, zero value otherwise.

### GetDataCurrencyOk

`func (o *CalendarReleaseRow) GetDataCurrencyOk() (*string, bool)`

GetDataCurrencyOk returns a tuple with the DataCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCurrency

`func (o *CalendarReleaseRow) SetDataCurrency(v string)`

SetDataCurrency sets DataCurrency field to given value.

### HasDataCurrency

`func (o *CalendarReleaseRow) HasDataCurrency() bool`

HasDataCurrency returns a boolean if a field has been set.

### SetDataCurrencyNil

`func (o *CalendarReleaseRow) SetDataCurrencyNil(b bool)`

 SetDataCurrencyNil sets the value for DataCurrency to be an explicit nil

### UnsetDataCurrency
`func (o *CalendarReleaseRow) UnsetDataCurrency()`

UnsetDataCurrency ensures that no value is present for DataCurrency, not even an explicit nil
### GetEndpointFamily

`func (o *CalendarReleaseRow) GetEndpointFamily() string`

GetEndpointFamily returns the EndpointFamily field if non-nil, zero value otherwise.

### GetEndpointFamilyOk

`func (o *CalendarReleaseRow) GetEndpointFamilyOk() (*string, bool)`

GetEndpointFamilyOk returns a tuple with the EndpointFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointFamily

`func (o *CalendarReleaseRow) SetEndpointFamily(v string)`

SetEndpointFamily sets EndpointFamily field to given value.

### HasEndpointFamily

`func (o *CalendarReleaseRow) HasEndpointFamily() bool`

HasEndpointFamily returns a boolean if a field has been set.

### SetEndpointFamilyNil

`func (o *CalendarReleaseRow) SetEndpointFamilyNil(b bool)`

 SetEndpointFamilyNil sets the value for EndpointFamily to be an explicit nil

### UnsetEndpointFamily
`func (o *CalendarReleaseRow) UnsetEndpointFamily()`

UnsetEndpointFamily ensures that no value is present for EndpointFamily, not even an explicit nil
### GetEndpointPath

`func (o *CalendarReleaseRow) GetEndpointPath() string`

GetEndpointPath returns the EndpointPath field if non-nil, zero value otherwise.

### GetEndpointPathOk

`func (o *CalendarReleaseRow) GetEndpointPathOk() (*string, bool)`

GetEndpointPathOk returns a tuple with the EndpointPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointPath

`func (o *CalendarReleaseRow) SetEndpointPath(v string)`

SetEndpointPath sets EndpointPath field to given value.

### HasEndpointPath

`func (o *CalendarReleaseRow) HasEndpointPath() bool`

HasEndpointPath returns a boolean if a field has been set.

### SetEndpointPathNil

`func (o *CalendarReleaseRow) SetEndpointPathNil(b bool)`

 SetEndpointPathNil sets the value for EndpointPath to be an explicit nil

### UnsetEndpointPath
`func (o *CalendarReleaseRow) UnsetEndpointPath()`

UnsetEndpointPath ensures that no value is present for EndpointPath, not even an explicit nil
### GetRequiresApiKey

`func (o *CalendarReleaseRow) GetRequiresApiKey() bool`

GetRequiresApiKey returns the RequiresApiKey field if non-nil, zero value otherwise.

### GetRequiresApiKeyOk

`func (o *CalendarReleaseRow) GetRequiresApiKeyOk() (*bool, bool)`

GetRequiresApiKeyOk returns a tuple with the RequiresApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApiKey

`func (o *CalendarReleaseRow) SetRequiresApiKey(v bool)`

SetRequiresApiKey sets RequiresApiKey field to given value.

### HasRequiresApiKey

`func (o *CalendarReleaseRow) HasRequiresApiKey() bool`

HasRequiresApiKey returns a boolean if a field has been set.

### SetRequiresApiKeyNil

`func (o *CalendarReleaseRow) SetRequiresApiKeyNil(b bool)`

 SetRequiresApiKeyNil sets the value for RequiresApiKey to be an explicit nil

### UnsetRequiresApiKey
`func (o *CalendarReleaseRow) UnsetRequiresApiKey()`

UnsetRequiresApiKey ensures that no value is present for RequiresApiKey, not even an explicit nil
### GetEventImportance

`func (o *CalendarReleaseRow) GetEventImportance() string`

GetEventImportance returns the EventImportance field if non-nil, zero value otherwise.

### GetEventImportanceOk

`func (o *CalendarReleaseRow) GetEventImportanceOk() (*string, bool)`

GetEventImportanceOk returns a tuple with the EventImportance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventImportance

`func (o *CalendarReleaseRow) SetEventImportance(v string)`

SetEventImportance sets EventImportance field to given value.

### HasEventImportance

`func (o *CalendarReleaseRow) HasEventImportance() bool`

HasEventImportance returns a boolean if a field has been set.

### SetEventImportanceNil

`func (o *CalendarReleaseRow) SetEventImportanceNil(b bool)`

 SetEventImportanceNil sets the value for EventImportance to be an explicit nil

### UnsetEventImportance
`func (o *CalendarReleaseRow) UnsetEventImportance()`

UnsetEventImportance ensures that no value is present for EventImportance, not even an explicit nil
### GetMarketTier

`func (o *CalendarReleaseRow) GetMarketTier() int64`

GetMarketTier returns the MarketTier field if non-nil, zero value otherwise.

### GetMarketTierOk

`func (o *CalendarReleaseRow) GetMarketTierOk() (*int64, bool)`

GetMarketTierOk returns a tuple with the MarketTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTier

`func (o *CalendarReleaseRow) SetMarketTier(v int64)`

SetMarketTier sets MarketTier field to given value.

### HasMarketTier

`func (o *CalendarReleaseRow) HasMarketTier() bool`

HasMarketTier returns a boolean if a field has been set.

### SetMarketTierNil

`func (o *CalendarReleaseRow) SetMarketTierNil(b bool)`

 SetMarketTierNil sets the value for MarketTier to be an explicit nil

### UnsetMarketTier
`func (o *CalendarReleaseRow) UnsetMarketTier()`

UnsetMarketTier ensures that no value is present for MarketTier, not even an explicit nil
### GetTopTierForCurrency

`func (o *CalendarReleaseRow) GetTopTierForCurrency() bool`

GetTopTierForCurrency returns the TopTierForCurrency field if non-nil, zero value otherwise.

### GetTopTierForCurrencyOk

`func (o *CalendarReleaseRow) GetTopTierForCurrencyOk() (*bool, bool)`

GetTopTierForCurrencyOk returns a tuple with the TopTierForCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopTierForCurrency

`func (o *CalendarReleaseRow) SetTopTierForCurrency(v bool)`

SetTopTierForCurrency sets TopTierForCurrency field to given value.

### HasTopTierForCurrency

`func (o *CalendarReleaseRow) HasTopTierForCurrency() bool`

HasTopTierForCurrency returns a boolean if a field has been set.

### SetTopTierForCurrencyNil

`func (o *CalendarReleaseRow) SetTopTierForCurrencyNil(b bool)`

 SetTopTierForCurrencyNil sets the value for TopTierForCurrency to be an explicit nil

### UnsetTopTierForCurrency
`func (o *CalendarReleaseRow) UnsetTopTierForCurrency()`

UnsetTopTierForCurrency ensures that no value is present for TopTierForCurrency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


