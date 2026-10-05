# AnnouncementDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnouncementId** | Pointer to **NullableString** |  | [optional] 
**Date** | **string** |  | 
**Val** | Pointer to **NullableFloat64** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceUrl** | Pointer to **NullableString** |  | [optional] 
**SourceUrlScope** | Pointer to **NullableString** |  | [optional] 
**PreviousValue** | Pointer to **NullableFloat64** |  | [optional] 
**PreviousDate** | Pointer to **NullableString** |  | [optional] 
**PreviousAnnouncementDatetime** | Pointer to **NullableInt64** |  | [optional] 
**Change** | Pointer to **NullableFloat64** |  | [optional] 
**ChangeFromPrevious** | Pointer to **NullableFloat64** |  | [optional] 
**PctChangeFromPrevious** | Pointer to **NullableFloat64** |  | [optional] 
**OriginalVal** | Pointer to **NullableFloat64** |  | [optional] 
**OriginalUnit** | Pointer to **NullableString** |  | [optional] 
**ValMom** | Pointer to **NullableFloat64** |  | [optional] 
**ObservationId** | Pointer to **NullableString** |  | [optional] 
**AnnouncementDatetime** | Pointer to **NullableInt64** |  | [optional] 
**AnnouncementDatetimeLocal** | Pointer to **NullableString** |  | [optional] 
**ReleaseTimeAssumed** | Pointer to **NullableBool** | Present and true when announcement_datetime was derived from the publication lag measured on this series&#39; confirmed releases, because no official release calendar reaches the period. Omitted on rows whose publication time was captured from the release. | [optional] 
**ReleaseTimeUpperBound** | Pointer to **NullableBool** | Legacy estimate flag; does not establish point-in-time safety. | [optional] 
**PublicationTimeStatus** | Pointer to **NullableString** | What &#x60;announcement_datetime&#x60; on this row actually is. Present on every row that carries a value.  - &#x60;confirmed&#x60;: read from the publisher&#39;s own release - an embargo instant printed on the document, or a publication timestamp the source served with the data. - &#x60;unverified&#x60;: a real timestamp that this row does not evidence the origin of. Most historical rows are in this state. It is not an estimate, and it is not proof. - &#x60;assumed_historical&#x60;: derived by this service from the lag its own confirmed releases show. Only permitted where both the period and the estimated publication precede 2025-11-06; see &#x60;release_time_assumed&#x60;. - &#x60;unknown&#x60;: no publication time. &#x60;announcement_datetime&#x60; is null and the observation is served on its period date alone.  Only &#x60;confirmed&#x60; is evidence of when the value became public. Gate point-in-time work on this field per row, never on a series-level claim. | [optional] 
**PublicationTimePrecision** | Pointer to **NullableString** | Resolution of &#x60;announcement_datetime&#x60;: &#x60;minute&#x60; when the publisher gave a clock time, &#x60;day&#x60; when only the release date was sourced, &#x60;unknown&#x60; otherwise. | [optional] 
**VintageStatus** | Pointer to **NullableString** |  | [optional] 
**IsFirstRelease** | Pointer to **NullableBool** |  | [optional] 
**ObservedAtNs** | Pointer to **NullableInt64** |  | [optional] 
**ObservedAtNsString** | Pointer to **NullableString** |  | [optional] 
**CollectedAtNsString** | Pointer to **NullableString** |  | [optional] 
**PublicationAtNs** | Pointer to **NullableInt64** | Publisher-sourced publication instant, Unix epoch nanoseconds (UTC). Present only when the source states one; omitted otherwise. | [optional] 
**PublicationAtNsString** | Pointer to **NullableString** | &#x60;publication_at_ns&#x60; as a decimal string; present only alongside it. | [optional] 
**AnnouncementSourceUrl** | Pointer to **NullableString** | The specific release this value was read from - for USD CPI, the BLS archive page for that month&#39;s news release. Present only when the writer recorded a release-scoped citation; a dataset or landing-page URL is reported as &#x60;source_url&#x60; instead and is never promoted to this field. | [optional] 
**SourceArtifactSha256** | Pointer to **NullableString** |  | [optional] 
**CaptureTimeBasis** | Pointer to **NullableString** |  | [optional] 
**Geography** | Pointer to **NullableString** |  | [optional] 
**PeriodBasis** | Pointer to **NullableString** |  | [optional] 
**MethodologyId** | Pointer to **NullableString** |  | [optional] 
**SourcePeriodLabelConflict** | Pointer to **NullableBool** |  | [optional] 
**SourceMethodologyStatus** | Pointer to **NullableString** |  | [optional] 
**SelectedVintageKnownAtNs** | Pointer to **NullableString** |  | [optional] 
**ReplayVintageVerified** | Pointer to **NullableBool** |  | [optional] 
**OfficialPlannedReleaseDatetime** | Pointer to **NullableInt64** |  | [optional] 
**OfficialPlannedReleaseDatetimeLocal** | Pointer to **NullableString** |  | [optional] 
**OfficialActualReleaseDatetime** | Pointer to **NullableInt64** |  | [optional] 
**OfficialActualReleaseDatetimeLocal** | Pointer to **NullableString** |  | [optional] 
**CollectedAtNs** | Pointer to **NullableInt64** |  | [optional] 
**CollectedAtIso** | Pointer to **NullableString** |  | [optional] 
**ReleaseTiming** | Pointer to [**ReleaseTiming**](ReleaseTiming.md) |  | [optional] 
**IngestionLatencyMs** | Pointer to **NullableFloat64** |  | [optional] 
**IngestionLatencyReference** | Pointer to **NullableString** |  | [optional] 
**PctChange** | Pointer to **NullableFloat64** |  | [optional] 
**PctChangeYoy** | Pointer to **NullableFloat64** |  | [optional] 
**PctChangeQoq** | Pointer to **NullableFloat64** |  | [optional] 
**PctChangeMom** | Pointer to **NullableFloat64** |  | [optional] 
**PctChange12m** | Pointer to **NullableFloat64** |  | [optional] 
**Revisions** | Pointer to [**[]RevisionEntry**](RevisionEntry.md) |  | [optional] 
**OutsideRequestedWindow** | Pointer to **NullableBool** |  | [optional] 
**RequestedStartDate** | Pointer to **NullableString** |  | [optional] 
**RequestedEndDate** | Pointer to **NullableString** |  | [optional] 
**ReturnedReason** | Pointer to **NullableString** |  | [optional] 
**StalenessDays** | Pointer to **NullableInt64** |  | [optional] 
**CanonicalIndicator** | Pointer to **NullableString** |  | [optional] 
**RawIndicator** | Pointer to **NullableString** |  | [optional] 
**RemapApplied** | Pointer to **NullableBool** |  | [optional] 
**RemapRuleId** | Pointer to **NullableString** |  | [optional] 
**RemapSegmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAnnouncementDataPoint

`func NewAnnouncementDataPoint(date string, ) *AnnouncementDataPoint`

NewAnnouncementDataPoint instantiates a new AnnouncementDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementDataPointWithDefaults

`func NewAnnouncementDataPointWithDefaults() *AnnouncementDataPoint`

NewAnnouncementDataPointWithDefaults instantiates a new AnnouncementDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnouncementId

`func (o *AnnouncementDataPoint) GetAnnouncementId() string`

GetAnnouncementId returns the AnnouncementId field if non-nil, zero value otherwise.

### GetAnnouncementIdOk

`func (o *AnnouncementDataPoint) GetAnnouncementIdOk() (*string, bool)`

GetAnnouncementIdOk returns a tuple with the AnnouncementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementId

`func (o *AnnouncementDataPoint) SetAnnouncementId(v string)`

SetAnnouncementId sets AnnouncementId field to given value.

### HasAnnouncementId

`func (o *AnnouncementDataPoint) HasAnnouncementId() bool`

HasAnnouncementId returns a boolean if a field has been set.

### SetAnnouncementIdNil

`func (o *AnnouncementDataPoint) SetAnnouncementIdNil(b bool)`

 SetAnnouncementIdNil sets the value for AnnouncementId to be an explicit nil

### UnsetAnnouncementId
`func (o *AnnouncementDataPoint) UnsetAnnouncementId()`

UnsetAnnouncementId ensures that no value is present for AnnouncementId, not even an explicit nil
### GetDate

`func (o *AnnouncementDataPoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AnnouncementDataPoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AnnouncementDataPoint) SetDate(v string)`

SetDate sets Date field to given value.


### GetVal

`func (o *AnnouncementDataPoint) GetVal() float64`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *AnnouncementDataPoint) GetValOk() (*float64, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *AnnouncementDataPoint) SetVal(v float64)`

SetVal sets Val field to given value.

### HasVal

`func (o *AnnouncementDataPoint) HasVal() bool`

HasVal returns a boolean if a field has been set.

### SetValNil

`func (o *AnnouncementDataPoint) SetValNil(b bool)`

 SetValNil sets the value for Val to be an explicit nil

### UnsetVal
`func (o *AnnouncementDataPoint) UnsetVal()`

UnsetVal ensures that no value is present for Val, not even an explicit nil
### GetSource

`func (o *AnnouncementDataPoint) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AnnouncementDataPoint) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AnnouncementDataPoint) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AnnouncementDataPoint) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *AnnouncementDataPoint) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AnnouncementDataPoint) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceUrl

`func (o *AnnouncementDataPoint) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *AnnouncementDataPoint) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *AnnouncementDataPoint) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *AnnouncementDataPoint) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *AnnouncementDataPoint) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *AnnouncementDataPoint) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSourceUrlScope

`func (o *AnnouncementDataPoint) GetSourceUrlScope() string`

GetSourceUrlScope returns the SourceUrlScope field if non-nil, zero value otherwise.

### GetSourceUrlScopeOk

`func (o *AnnouncementDataPoint) GetSourceUrlScopeOk() (*string, bool)`

GetSourceUrlScopeOk returns a tuple with the SourceUrlScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrlScope

`func (o *AnnouncementDataPoint) SetSourceUrlScope(v string)`

SetSourceUrlScope sets SourceUrlScope field to given value.

### HasSourceUrlScope

`func (o *AnnouncementDataPoint) HasSourceUrlScope() bool`

HasSourceUrlScope returns a boolean if a field has been set.

### SetSourceUrlScopeNil

`func (o *AnnouncementDataPoint) SetSourceUrlScopeNil(b bool)`

 SetSourceUrlScopeNil sets the value for SourceUrlScope to be an explicit nil

### UnsetSourceUrlScope
`func (o *AnnouncementDataPoint) UnsetSourceUrlScope()`

UnsetSourceUrlScope ensures that no value is present for SourceUrlScope, not even an explicit nil
### GetPreviousValue

`func (o *AnnouncementDataPoint) GetPreviousValue() float64`

GetPreviousValue returns the PreviousValue field if non-nil, zero value otherwise.

### GetPreviousValueOk

`func (o *AnnouncementDataPoint) GetPreviousValueOk() (*float64, bool)`

GetPreviousValueOk returns a tuple with the PreviousValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValue

`func (o *AnnouncementDataPoint) SetPreviousValue(v float64)`

SetPreviousValue sets PreviousValue field to given value.

### HasPreviousValue

`func (o *AnnouncementDataPoint) HasPreviousValue() bool`

HasPreviousValue returns a boolean if a field has been set.

### SetPreviousValueNil

`func (o *AnnouncementDataPoint) SetPreviousValueNil(b bool)`

 SetPreviousValueNil sets the value for PreviousValue to be an explicit nil

### UnsetPreviousValue
`func (o *AnnouncementDataPoint) UnsetPreviousValue()`

UnsetPreviousValue ensures that no value is present for PreviousValue, not even an explicit nil
### GetPreviousDate

`func (o *AnnouncementDataPoint) GetPreviousDate() string`

GetPreviousDate returns the PreviousDate field if non-nil, zero value otherwise.

### GetPreviousDateOk

`func (o *AnnouncementDataPoint) GetPreviousDateOk() (*string, bool)`

GetPreviousDateOk returns a tuple with the PreviousDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousDate

`func (o *AnnouncementDataPoint) SetPreviousDate(v string)`

SetPreviousDate sets PreviousDate field to given value.

### HasPreviousDate

`func (o *AnnouncementDataPoint) HasPreviousDate() bool`

HasPreviousDate returns a boolean if a field has been set.

### SetPreviousDateNil

`func (o *AnnouncementDataPoint) SetPreviousDateNil(b bool)`

 SetPreviousDateNil sets the value for PreviousDate to be an explicit nil

### UnsetPreviousDate
`func (o *AnnouncementDataPoint) UnsetPreviousDate()`

UnsetPreviousDate ensures that no value is present for PreviousDate, not even an explicit nil
### GetPreviousAnnouncementDatetime

`func (o *AnnouncementDataPoint) GetPreviousAnnouncementDatetime() int64`

GetPreviousAnnouncementDatetime returns the PreviousAnnouncementDatetime field if non-nil, zero value otherwise.

### GetPreviousAnnouncementDatetimeOk

`func (o *AnnouncementDataPoint) GetPreviousAnnouncementDatetimeOk() (*int64, bool)`

GetPreviousAnnouncementDatetimeOk returns a tuple with the PreviousAnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAnnouncementDatetime

`func (o *AnnouncementDataPoint) SetPreviousAnnouncementDatetime(v int64)`

SetPreviousAnnouncementDatetime sets PreviousAnnouncementDatetime field to given value.

### HasPreviousAnnouncementDatetime

`func (o *AnnouncementDataPoint) HasPreviousAnnouncementDatetime() bool`

HasPreviousAnnouncementDatetime returns a boolean if a field has been set.

### SetPreviousAnnouncementDatetimeNil

`func (o *AnnouncementDataPoint) SetPreviousAnnouncementDatetimeNil(b bool)`

 SetPreviousAnnouncementDatetimeNil sets the value for PreviousAnnouncementDatetime to be an explicit nil

### UnsetPreviousAnnouncementDatetime
`func (o *AnnouncementDataPoint) UnsetPreviousAnnouncementDatetime()`

UnsetPreviousAnnouncementDatetime ensures that no value is present for PreviousAnnouncementDatetime, not even an explicit nil
### GetChange

`func (o *AnnouncementDataPoint) GetChange() float64`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *AnnouncementDataPoint) GetChangeOk() (*float64, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *AnnouncementDataPoint) SetChange(v float64)`

SetChange sets Change field to given value.

### HasChange

`func (o *AnnouncementDataPoint) HasChange() bool`

HasChange returns a boolean if a field has been set.

### SetChangeNil

`func (o *AnnouncementDataPoint) SetChangeNil(b bool)`

 SetChangeNil sets the value for Change to be an explicit nil

### UnsetChange
`func (o *AnnouncementDataPoint) UnsetChange()`

UnsetChange ensures that no value is present for Change, not even an explicit nil
### GetChangeFromPrevious

`func (o *AnnouncementDataPoint) GetChangeFromPrevious() float64`

GetChangeFromPrevious returns the ChangeFromPrevious field if non-nil, zero value otherwise.

### GetChangeFromPreviousOk

`func (o *AnnouncementDataPoint) GetChangeFromPreviousOk() (*float64, bool)`

GetChangeFromPreviousOk returns a tuple with the ChangeFromPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeFromPrevious

`func (o *AnnouncementDataPoint) SetChangeFromPrevious(v float64)`

SetChangeFromPrevious sets ChangeFromPrevious field to given value.

### HasChangeFromPrevious

`func (o *AnnouncementDataPoint) HasChangeFromPrevious() bool`

HasChangeFromPrevious returns a boolean if a field has been set.

### SetChangeFromPreviousNil

`func (o *AnnouncementDataPoint) SetChangeFromPreviousNil(b bool)`

 SetChangeFromPreviousNil sets the value for ChangeFromPrevious to be an explicit nil

### UnsetChangeFromPrevious
`func (o *AnnouncementDataPoint) UnsetChangeFromPrevious()`

UnsetChangeFromPrevious ensures that no value is present for ChangeFromPrevious, not even an explicit nil
### GetPctChangeFromPrevious

`func (o *AnnouncementDataPoint) GetPctChangeFromPrevious() float64`

GetPctChangeFromPrevious returns the PctChangeFromPrevious field if non-nil, zero value otherwise.

### GetPctChangeFromPreviousOk

`func (o *AnnouncementDataPoint) GetPctChangeFromPreviousOk() (*float64, bool)`

GetPctChangeFromPreviousOk returns a tuple with the PctChangeFromPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctChangeFromPrevious

`func (o *AnnouncementDataPoint) SetPctChangeFromPrevious(v float64)`

SetPctChangeFromPrevious sets PctChangeFromPrevious field to given value.

### HasPctChangeFromPrevious

`func (o *AnnouncementDataPoint) HasPctChangeFromPrevious() bool`

HasPctChangeFromPrevious returns a boolean if a field has been set.

### SetPctChangeFromPreviousNil

`func (o *AnnouncementDataPoint) SetPctChangeFromPreviousNil(b bool)`

 SetPctChangeFromPreviousNil sets the value for PctChangeFromPrevious to be an explicit nil

### UnsetPctChangeFromPrevious
`func (o *AnnouncementDataPoint) UnsetPctChangeFromPrevious()`

UnsetPctChangeFromPrevious ensures that no value is present for PctChangeFromPrevious, not even an explicit nil
### GetOriginalVal

`func (o *AnnouncementDataPoint) GetOriginalVal() float64`

GetOriginalVal returns the OriginalVal field if non-nil, zero value otherwise.

### GetOriginalValOk

`func (o *AnnouncementDataPoint) GetOriginalValOk() (*float64, bool)`

GetOriginalValOk returns a tuple with the OriginalVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalVal

`func (o *AnnouncementDataPoint) SetOriginalVal(v float64)`

SetOriginalVal sets OriginalVal field to given value.

### HasOriginalVal

`func (o *AnnouncementDataPoint) HasOriginalVal() bool`

HasOriginalVal returns a boolean if a field has been set.

### SetOriginalValNil

`func (o *AnnouncementDataPoint) SetOriginalValNil(b bool)`

 SetOriginalValNil sets the value for OriginalVal to be an explicit nil

### UnsetOriginalVal
`func (o *AnnouncementDataPoint) UnsetOriginalVal()`

UnsetOriginalVal ensures that no value is present for OriginalVal, not even an explicit nil
### GetOriginalUnit

`func (o *AnnouncementDataPoint) GetOriginalUnit() string`

GetOriginalUnit returns the OriginalUnit field if non-nil, zero value otherwise.

### GetOriginalUnitOk

`func (o *AnnouncementDataPoint) GetOriginalUnitOk() (*string, bool)`

GetOriginalUnitOk returns a tuple with the OriginalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalUnit

`func (o *AnnouncementDataPoint) SetOriginalUnit(v string)`

SetOriginalUnit sets OriginalUnit field to given value.

### HasOriginalUnit

`func (o *AnnouncementDataPoint) HasOriginalUnit() bool`

HasOriginalUnit returns a boolean if a field has been set.

### SetOriginalUnitNil

`func (o *AnnouncementDataPoint) SetOriginalUnitNil(b bool)`

 SetOriginalUnitNil sets the value for OriginalUnit to be an explicit nil

### UnsetOriginalUnit
`func (o *AnnouncementDataPoint) UnsetOriginalUnit()`

UnsetOriginalUnit ensures that no value is present for OriginalUnit, not even an explicit nil
### GetValMom

`func (o *AnnouncementDataPoint) GetValMom() float64`

GetValMom returns the ValMom field if non-nil, zero value otherwise.

### GetValMomOk

`func (o *AnnouncementDataPoint) GetValMomOk() (*float64, bool)`

GetValMomOk returns a tuple with the ValMom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValMom

`func (o *AnnouncementDataPoint) SetValMom(v float64)`

SetValMom sets ValMom field to given value.

### HasValMom

`func (o *AnnouncementDataPoint) HasValMom() bool`

HasValMom returns a boolean if a field has been set.

### SetValMomNil

`func (o *AnnouncementDataPoint) SetValMomNil(b bool)`

 SetValMomNil sets the value for ValMom to be an explicit nil

### UnsetValMom
`func (o *AnnouncementDataPoint) UnsetValMom()`

UnsetValMom ensures that no value is present for ValMom, not even an explicit nil
### GetObservationId

`func (o *AnnouncementDataPoint) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *AnnouncementDataPoint) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *AnnouncementDataPoint) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *AnnouncementDataPoint) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### SetObservationIdNil

`func (o *AnnouncementDataPoint) SetObservationIdNil(b bool)`

 SetObservationIdNil sets the value for ObservationId to be an explicit nil

### UnsetObservationId
`func (o *AnnouncementDataPoint) UnsetObservationId()`

UnsetObservationId ensures that no value is present for ObservationId, not even an explicit nil
### GetAnnouncementDatetime

`func (o *AnnouncementDataPoint) GetAnnouncementDatetime() int64`

GetAnnouncementDatetime returns the AnnouncementDatetime field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeOk

`func (o *AnnouncementDataPoint) GetAnnouncementDatetimeOk() (*int64, bool)`

GetAnnouncementDatetimeOk returns a tuple with the AnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetime

`func (o *AnnouncementDataPoint) SetAnnouncementDatetime(v int64)`

SetAnnouncementDatetime sets AnnouncementDatetime field to given value.

### HasAnnouncementDatetime

`func (o *AnnouncementDataPoint) HasAnnouncementDatetime() bool`

HasAnnouncementDatetime returns a boolean if a field has been set.

### SetAnnouncementDatetimeNil

`func (o *AnnouncementDataPoint) SetAnnouncementDatetimeNil(b bool)`

 SetAnnouncementDatetimeNil sets the value for AnnouncementDatetime to be an explicit nil

### UnsetAnnouncementDatetime
`func (o *AnnouncementDataPoint) UnsetAnnouncementDatetime()`

UnsetAnnouncementDatetime ensures that no value is present for AnnouncementDatetime, not even an explicit nil
### GetAnnouncementDatetimeLocal

`func (o *AnnouncementDataPoint) GetAnnouncementDatetimeLocal() string`

GetAnnouncementDatetimeLocal returns the AnnouncementDatetimeLocal field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeLocalOk

`func (o *AnnouncementDataPoint) GetAnnouncementDatetimeLocalOk() (*string, bool)`

GetAnnouncementDatetimeLocalOk returns a tuple with the AnnouncementDatetimeLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetimeLocal

`func (o *AnnouncementDataPoint) SetAnnouncementDatetimeLocal(v string)`

SetAnnouncementDatetimeLocal sets AnnouncementDatetimeLocal field to given value.

### HasAnnouncementDatetimeLocal

`func (o *AnnouncementDataPoint) HasAnnouncementDatetimeLocal() bool`

HasAnnouncementDatetimeLocal returns a boolean if a field has been set.

### SetAnnouncementDatetimeLocalNil

`func (o *AnnouncementDataPoint) SetAnnouncementDatetimeLocalNil(b bool)`

 SetAnnouncementDatetimeLocalNil sets the value for AnnouncementDatetimeLocal to be an explicit nil

### UnsetAnnouncementDatetimeLocal
`func (o *AnnouncementDataPoint) UnsetAnnouncementDatetimeLocal()`

UnsetAnnouncementDatetimeLocal ensures that no value is present for AnnouncementDatetimeLocal, not even an explicit nil
### GetReleaseTimeAssumed

`func (o *AnnouncementDataPoint) GetReleaseTimeAssumed() bool`

GetReleaseTimeAssumed returns the ReleaseTimeAssumed field if non-nil, zero value otherwise.

### GetReleaseTimeAssumedOk

`func (o *AnnouncementDataPoint) GetReleaseTimeAssumedOk() (*bool, bool)`

GetReleaseTimeAssumedOk returns a tuple with the ReleaseTimeAssumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTimeAssumed

`func (o *AnnouncementDataPoint) SetReleaseTimeAssumed(v bool)`

SetReleaseTimeAssumed sets ReleaseTimeAssumed field to given value.

### HasReleaseTimeAssumed

`func (o *AnnouncementDataPoint) HasReleaseTimeAssumed() bool`

HasReleaseTimeAssumed returns a boolean if a field has been set.

### SetReleaseTimeAssumedNil

`func (o *AnnouncementDataPoint) SetReleaseTimeAssumedNil(b bool)`

 SetReleaseTimeAssumedNil sets the value for ReleaseTimeAssumed to be an explicit nil

### UnsetReleaseTimeAssumed
`func (o *AnnouncementDataPoint) UnsetReleaseTimeAssumed()`

UnsetReleaseTimeAssumed ensures that no value is present for ReleaseTimeAssumed, not even an explicit nil
### GetReleaseTimeUpperBound

`func (o *AnnouncementDataPoint) GetReleaseTimeUpperBound() bool`

GetReleaseTimeUpperBound returns the ReleaseTimeUpperBound field if non-nil, zero value otherwise.

### GetReleaseTimeUpperBoundOk

`func (o *AnnouncementDataPoint) GetReleaseTimeUpperBoundOk() (*bool, bool)`

GetReleaseTimeUpperBoundOk returns a tuple with the ReleaseTimeUpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTimeUpperBound

`func (o *AnnouncementDataPoint) SetReleaseTimeUpperBound(v bool)`

SetReleaseTimeUpperBound sets ReleaseTimeUpperBound field to given value.

### HasReleaseTimeUpperBound

`func (o *AnnouncementDataPoint) HasReleaseTimeUpperBound() bool`

HasReleaseTimeUpperBound returns a boolean if a field has been set.

### SetReleaseTimeUpperBoundNil

`func (o *AnnouncementDataPoint) SetReleaseTimeUpperBoundNil(b bool)`

 SetReleaseTimeUpperBoundNil sets the value for ReleaseTimeUpperBound to be an explicit nil

### UnsetReleaseTimeUpperBound
`func (o *AnnouncementDataPoint) UnsetReleaseTimeUpperBound()`

UnsetReleaseTimeUpperBound ensures that no value is present for ReleaseTimeUpperBound, not even an explicit nil
### GetPublicationTimeStatus

`func (o *AnnouncementDataPoint) GetPublicationTimeStatus() string`

GetPublicationTimeStatus returns the PublicationTimeStatus field if non-nil, zero value otherwise.

### GetPublicationTimeStatusOk

`func (o *AnnouncementDataPoint) GetPublicationTimeStatusOk() (*string, bool)`

GetPublicationTimeStatusOk returns a tuple with the PublicationTimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationTimeStatus

`func (o *AnnouncementDataPoint) SetPublicationTimeStatus(v string)`

SetPublicationTimeStatus sets PublicationTimeStatus field to given value.

### HasPublicationTimeStatus

`func (o *AnnouncementDataPoint) HasPublicationTimeStatus() bool`

HasPublicationTimeStatus returns a boolean if a field has been set.

### SetPublicationTimeStatusNil

`func (o *AnnouncementDataPoint) SetPublicationTimeStatusNil(b bool)`

 SetPublicationTimeStatusNil sets the value for PublicationTimeStatus to be an explicit nil

### UnsetPublicationTimeStatus
`func (o *AnnouncementDataPoint) UnsetPublicationTimeStatus()`

UnsetPublicationTimeStatus ensures that no value is present for PublicationTimeStatus, not even an explicit nil
### GetPublicationTimePrecision

`func (o *AnnouncementDataPoint) GetPublicationTimePrecision() string`

GetPublicationTimePrecision returns the PublicationTimePrecision field if non-nil, zero value otherwise.

### GetPublicationTimePrecisionOk

`func (o *AnnouncementDataPoint) GetPublicationTimePrecisionOk() (*string, bool)`

GetPublicationTimePrecisionOk returns a tuple with the PublicationTimePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationTimePrecision

`func (o *AnnouncementDataPoint) SetPublicationTimePrecision(v string)`

SetPublicationTimePrecision sets PublicationTimePrecision field to given value.

### HasPublicationTimePrecision

`func (o *AnnouncementDataPoint) HasPublicationTimePrecision() bool`

HasPublicationTimePrecision returns a boolean if a field has been set.

### SetPublicationTimePrecisionNil

`func (o *AnnouncementDataPoint) SetPublicationTimePrecisionNil(b bool)`

 SetPublicationTimePrecisionNil sets the value for PublicationTimePrecision to be an explicit nil

### UnsetPublicationTimePrecision
`func (o *AnnouncementDataPoint) UnsetPublicationTimePrecision()`

UnsetPublicationTimePrecision ensures that no value is present for PublicationTimePrecision, not even an explicit nil
### GetVintageStatus

`func (o *AnnouncementDataPoint) GetVintageStatus() string`

GetVintageStatus returns the VintageStatus field if non-nil, zero value otherwise.

### GetVintageStatusOk

`func (o *AnnouncementDataPoint) GetVintageStatusOk() (*string, bool)`

GetVintageStatusOk returns a tuple with the VintageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVintageStatus

`func (o *AnnouncementDataPoint) SetVintageStatus(v string)`

SetVintageStatus sets VintageStatus field to given value.

### HasVintageStatus

`func (o *AnnouncementDataPoint) HasVintageStatus() bool`

HasVintageStatus returns a boolean if a field has been set.

### SetVintageStatusNil

`func (o *AnnouncementDataPoint) SetVintageStatusNil(b bool)`

 SetVintageStatusNil sets the value for VintageStatus to be an explicit nil

### UnsetVintageStatus
`func (o *AnnouncementDataPoint) UnsetVintageStatus()`

UnsetVintageStatus ensures that no value is present for VintageStatus, not even an explicit nil
### GetIsFirstRelease

`func (o *AnnouncementDataPoint) GetIsFirstRelease() bool`

GetIsFirstRelease returns the IsFirstRelease field if non-nil, zero value otherwise.

### GetIsFirstReleaseOk

`func (o *AnnouncementDataPoint) GetIsFirstReleaseOk() (*bool, bool)`

GetIsFirstReleaseOk returns a tuple with the IsFirstRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstRelease

`func (o *AnnouncementDataPoint) SetIsFirstRelease(v bool)`

SetIsFirstRelease sets IsFirstRelease field to given value.

### HasIsFirstRelease

`func (o *AnnouncementDataPoint) HasIsFirstRelease() bool`

HasIsFirstRelease returns a boolean if a field has been set.

### SetIsFirstReleaseNil

`func (o *AnnouncementDataPoint) SetIsFirstReleaseNil(b bool)`

 SetIsFirstReleaseNil sets the value for IsFirstRelease to be an explicit nil

### UnsetIsFirstRelease
`func (o *AnnouncementDataPoint) UnsetIsFirstRelease()`

UnsetIsFirstRelease ensures that no value is present for IsFirstRelease, not even an explicit nil
### GetObservedAtNs

`func (o *AnnouncementDataPoint) GetObservedAtNs() int64`

GetObservedAtNs returns the ObservedAtNs field if non-nil, zero value otherwise.

### GetObservedAtNsOk

`func (o *AnnouncementDataPoint) GetObservedAtNsOk() (*int64, bool)`

GetObservedAtNsOk returns a tuple with the ObservedAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAtNs

`func (o *AnnouncementDataPoint) SetObservedAtNs(v int64)`

SetObservedAtNs sets ObservedAtNs field to given value.

### HasObservedAtNs

`func (o *AnnouncementDataPoint) HasObservedAtNs() bool`

HasObservedAtNs returns a boolean if a field has been set.

### SetObservedAtNsNil

`func (o *AnnouncementDataPoint) SetObservedAtNsNil(b bool)`

 SetObservedAtNsNil sets the value for ObservedAtNs to be an explicit nil

### UnsetObservedAtNs
`func (o *AnnouncementDataPoint) UnsetObservedAtNs()`

UnsetObservedAtNs ensures that no value is present for ObservedAtNs, not even an explicit nil
### GetObservedAtNsString

`func (o *AnnouncementDataPoint) GetObservedAtNsString() string`

GetObservedAtNsString returns the ObservedAtNsString field if non-nil, zero value otherwise.

### GetObservedAtNsStringOk

`func (o *AnnouncementDataPoint) GetObservedAtNsStringOk() (*string, bool)`

GetObservedAtNsStringOk returns a tuple with the ObservedAtNsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAtNsString

`func (o *AnnouncementDataPoint) SetObservedAtNsString(v string)`

SetObservedAtNsString sets ObservedAtNsString field to given value.

### HasObservedAtNsString

`func (o *AnnouncementDataPoint) HasObservedAtNsString() bool`

HasObservedAtNsString returns a boolean if a field has been set.

### SetObservedAtNsStringNil

`func (o *AnnouncementDataPoint) SetObservedAtNsStringNil(b bool)`

 SetObservedAtNsStringNil sets the value for ObservedAtNsString to be an explicit nil

### UnsetObservedAtNsString
`func (o *AnnouncementDataPoint) UnsetObservedAtNsString()`

UnsetObservedAtNsString ensures that no value is present for ObservedAtNsString, not even an explicit nil
### GetCollectedAtNsString

`func (o *AnnouncementDataPoint) GetCollectedAtNsString() string`

GetCollectedAtNsString returns the CollectedAtNsString field if non-nil, zero value otherwise.

### GetCollectedAtNsStringOk

`func (o *AnnouncementDataPoint) GetCollectedAtNsStringOk() (*string, bool)`

GetCollectedAtNsStringOk returns a tuple with the CollectedAtNsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectedAtNsString

`func (o *AnnouncementDataPoint) SetCollectedAtNsString(v string)`

SetCollectedAtNsString sets CollectedAtNsString field to given value.

### HasCollectedAtNsString

`func (o *AnnouncementDataPoint) HasCollectedAtNsString() bool`

HasCollectedAtNsString returns a boolean if a field has been set.

### SetCollectedAtNsStringNil

`func (o *AnnouncementDataPoint) SetCollectedAtNsStringNil(b bool)`

 SetCollectedAtNsStringNil sets the value for CollectedAtNsString to be an explicit nil

### UnsetCollectedAtNsString
`func (o *AnnouncementDataPoint) UnsetCollectedAtNsString()`

UnsetCollectedAtNsString ensures that no value is present for CollectedAtNsString, not even an explicit nil
### GetPublicationAtNs

`func (o *AnnouncementDataPoint) GetPublicationAtNs() int64`

GetPublicationAtNs returns the PublicationAtNs field if non-nil, zero value otherwise.

### GetPublicationAtNsOk

`func (o *AnnouncementDataPoint) GetPublicationAtNsOk() (*int64, bool)`

GetPublicationAtNsOk returns a tuple with the PublicationAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationAtNs

`func (o *AnnouncementDataPoint) SetPublicationAtNs(v int64)`

SetPublicationAtNs sets PublicationAtNs field to given value.

### HasPublicationAtNs

`func (o *AnnouncementDataPoint) HasPublicationAtNs() bool`

HasPublicationAtNs returns a boolean if a field has been set.

### SetPublicationAtNsNil

`func (o *AnnouncementDataPoint) SetPublicationAtNsNil(b bool)`

 SetPublicationAtNsNil sets the value for PublicationAtNs to be an explicit nil

### UnsetPublicationAtNs
`func (o *AnnouncementDataPoint) UnsetPublicationAtNs()`

UnsetPublicationAtNs ensures that no value is present for PublicationAtNs, not even an explicit nil
### GetPublicationAtNsString

`func (o *AnnouncementDataPoint) GetPublicationAtNsString() string`

GetPublicationAtNsString returns the PublicationAtNsString field if non-nil, zero value otherwise.

### GetPublicationAtNsStringOk

`func (o *AnnouncementDataPoint) GetPublicationAtNsStringOk() (*string, bool)`

GetPublicationAtNsStringOk returns a tuple with the PublicationAtNsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationAtNsString

`func (o *AnnouncementDataPoint) SetPublicationAtNsString(v string)`

SetPublicationAtNsString sets PublicationAtNsString field to given value.

### HasPublicationAtNsString

`func (o *AnnouncementDataPoint) HasPublicationAtNsString() bool`

HasPublicationAtNsString returns a boolean if a field has been set.

### SetPublicationAtNsStringNil

`func (o *AnnouncementDataPoint) SetPublicationAtNsStringNil(b bool)`

 SetPublicationAtNsStringNil sets the value for PublicationAtNsString to be an explicit nil

### UnsetPublicationAtNsString
`func (o *AnnouncementDataPoint) UnsetPublicationAtNsString()`

UnsetPublicationAtNsString ensures that no value is present for PublicationAtNsString, not even an explicit nil
### GetAnnouncementSourceUrl

`func (o *AnnouncementDataPoint) GetAnnouncementSourceUrl() string`

GetAnnouncementSourceUrl returns the AnnouncementSourceUrl field if non-nil, zero value otherwise.

### GetAnnouncementSourceUrlOk

`func (o *AnnouncementDataPoint) GetAnnouncementSourceUrlOk() (*string, bool)`

GetAnnouncementSourceUrlOk returns a tuple with the AnnouncementSourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementSourceUrl

`func (o *AnnouncementDataPoint) SetAnnouncementSourceUrl(v string)`

SetAnnouncementSourceUrl sets AnnouncementSourceUrl field to given value.

### HasAnnouncementSourceUrl

`func (o *AnnouncementDataPoint) HasAnnouncementSourceUrl() bool`

HasAnnouncementSourceUrl returns a boolean if a field has been set.

### SetAnnouncementSourceUrlNil

`func (o *AnnouncementDataPoint) SetAnnouncementSourceUrlNil(b bool)`

 SetAnnouncementSourceUrlNil sets the value for AnnouncementSourceUrl to be an explicit nil

### UnsetAnnouncementSourceUrl
`func (o *AnnouncementDataPoint) UnsetAnnouncementSourceUrl()`

UnsetAnnouncementSourceUrl ensures that no value is present for AnnouncementSourceUrl, not even an explicit nil
### GetSourceArtifactSha256

`func (o *AnnouncementDataPoint) GetSourceArtifactSha256() string`

GetSourceArtifactSha256 returns the SourceArtifactSha256 field if non-nil, zero value otherwise.

### GetSourceArtifactSha256Ok

`func (o *AnnouncementDataPoint) GetSourceArtifactSha256Ok() (*string, bool)`

GetSourceArtifactSha256Ok returns a tuple with the SourceArtifactSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceArtifactSha256

`func (o *AnnouncementDataPoint) SetSourceArtifactSha256(v string)`

SetSourceArtifactSha256 sets SourceArtifactSha256 field to given value.

### HasSourceArtifactSha256

`func (o *AnnouncementDataPoint) HasSourceArtifactSha256() bool`

HasSourceArtifactSha256 returns a boolean if a field has been set.

### SetSourceArtifactSha256Nil

`func (o *AnnouncementDataPoint) SetSourceArtifactSha256Nil(b bool)`

 SetSourceArtifactSha256Nil sets the value for SourceArtifactSha256 to be an explicit nil

### UnsetSourceArtifactSha256
`func (o *AnnouncementDataPoint) UnsetSourceArtifactSha256()`

UnsetSourceArtifactSha256 ensures that no value is present for SourceArtifactSha256, not even an explicit nil
### GetCaptureTimeBasis

`func (o *AnnouncementDataPoint) GetCaptureTimeBasis() string`

GetCaptureTimeBasis returns the CaptureTimeBasis field if non-nil, zero value otherwise.

### GetCaptureTimeBasisOk

`func (o *AnnouncementDataPoint) GetCaptureTimeBasisOk() (*string, bool)`

GetCaptureTimeBasisOk returns a tuple with the CaptureTimeBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureTimeBasis

`func (o *AnnouncementDataPoint) SetCaptureTimeBasis(v string)`

SetCaptureTimeBasis sets CaptureTimeBasis field to given value.

### HasCaptureTimeBasis

`func (o *AnnouncementDataPoint) HasCaptureTimeBasis() bool`

HasCaptureTimeBasis returns a boolean if a field has been set.

### SetCaptureTimeBasisNil

`func (o *AnnouncementDataPoint) SetCaptureTimeBasisNil(b bool)`

 SetCaptureTimeBasisNil sets the value for CaptureTimeBasis to be an explicit nil

### UnsetCaptureTimeBasis
`func (o *AnnouncementDataPoint) UnsetCaptureTimeBasis()`

UnsetCaptureTimeBasis ensures that no value is present for CaptureTimeBasis, not even an explicit nil
### GetGeography

`func (o *AnnouncementDataPoint) GetGeography() string`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *AnnouncementDataPoint) GetGeographyOk() (*string, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *AnnouncementDataPoint) SetGeography(v string)`

SetGeography sets Geography field to given value.

### HasGeography

`func (o *AnnouncementDataPoint) HasGeography() bool`

HasGeography returns a boolean if a field has been set.

### SetGeographyNil

`func (o *AnnouncementDataPoint) SetGeographyNil(b bool)`

 SetGeographyNil sets the value for Geography to be an explicit nil

### UnsetGeography
`func (o *AnnouncementDataPoint) UnsetGeography()`

UnsetGeography ensures that no value is present for Geography, not even an explicit nil
### GetPeriodBasis

`func (o *AnnouncementDataPoint) GetPeriodBasis() string`

GetPeriodBasis returns the PeriodBasis field if non-nil, zero value otherwise.

### GetPeriodBasisOk

`func (o *AnnouncementDataPoint) GetPeriodBasisOk() (*string, bool)`

GetPeriodBasisOk returns a tuple with the PeriodBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBasis

`func (o *AnnouncementDataPoint) SetPeriodBasis(v string)`

SetPeriodBasis sets PeriodBasis field to given value.

### HasPeriodBasis

`func (o *AnnouncementDataPoint) HasPeriodBasis() bool`

HasPeriodBasis returns a boolean if a field has been set.

### SetPeriodBasisNil

`func (o *AnnouncementDataPoint) SetPeriodBasisNil(b bool)`

 SetPeriodBasisNil sets the value for PeriodBasis to be an explicit nil

### UnsetPeriodBasis
`func (o *AnnouncementDataPoint) UnsetPeriodBasis()`

UnsetPeriodBasis ensures that no value is present for PeriodBasis, not even an explicit nil
### GetMethodologyId

`func (o *AnnouncementDataPoint) GetMethodologyId() string`

GetMethodologyId returns the MethodologyId field if non-nil, zero value otherwise.

### GetMethodologyIdOk

`func (o *AnnouncementDataPoint) GetMethodologyIdOk() (*string, bool)`

GetMethodologyIdOk returns a tuple with the MethodologyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodologyId

`func (o *AnnouncementDataPoint) SetMethodologyId(v string)`

SetMethodologyId sets MethodologyId field to given value.

### HasMethodologyId

`func (o *AnnouncementDataPoint) HasMethodologyId() bool`

HasMethodologyId returns a boolean if a field has been set.

### SetMethodologyIdNil

`func (o *AnnouncementDataPoint) SetMethodologyIdNil(b bool)`

 SetMethodologyIdNil sets the value for MethodologyId to be an explicit nil

### UnsetMethodologyId
`func (o *AnnouncementDataPoint) UnsetMethodologyId()`

UnsetMethodologyId ensures that no value is present for MethodologyId, not even an explicit nil
### GetSourcePeriodLabelConflict

`func (o *AnnouncementDataPoint) GetSourcePeriodLabelConflict() bool`

GetSourcePeriodLabelConflict returns the SourcePeriodLabelConflict field if non-nil, zero value otherwise.

### GetSourcePeriodLabelConflictOk

`func (o *AnnouncementDataPoint) GetSourcePeriodLabelConflictOk() (*bool, bool)`

GetSourcePeriodLabelConflictOk returns a tuple with the SourcePeriodLabelConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePeriodLabelConflict

`func (o *AnnouncementDataPoint) SetSourcePeriodLabelConflict(v bool)`

SetSourcePeriodLabelConflict sets SourcePeriodLabelConflict field to given value.

### HasSourcePeriodLabelConflict

`func (o *AnnouncementDataPoint) HasSourcePeriodLabelConflict() bool`

HasSourcePeriodLabelConflict returns a boolean if a field has been set.

### SetSourcePeriodLabelConflictNil

`func (o *AnnouncementDataPoint) SetSourcePeriodLabelConflictNil(b bool)`

 SetSourcePeriodLabelConflictNil sets the value for SourcePeriodLabelConflict to be an explicit nil

### UnsetSourcePeriodLabelConflict
`func (o *AnnouncementDataPoint) UnsetSourcePeriodLabelConflict()`

UnsetSourcePeriodLabelConflict ensures that no value is present for SourcePeriodLabelConflict, not even an explicit nil
### GetSourceMethodologyStatus

`func (o *AnnouncementDataPoint) GetSourceMethodologyStatus() string`

GetSourceMethodologyStatus returns the SourceMethodologyStatus field if non-nil, zero value otherwise.

### GetSourceMethodologyStatusOk

`func (o *AnnouncementDataPoint) GetSourceMethodologyStatusOk() (*string, bool)`

GetSourceMethodologyStatusOk returns a tuple with the SourceMethodologyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMethodologyStatus

`func (o *AnnouncementDataPoint) SetSourceMethodologyStatus(v string)`

SetSourceMethodologyStatus sets SourceMethodologyStatus field to given value.

### HasSourceMethodologyStatus

`func (o *AnnouncementDataPoint) HasSourceMethodologyStatus() bool`

HasSourceMethodologyStatus returns a boolean if a field has been set.

### SetSourceMethodologyStatusNil

`func (o *AnnouncementDataPoint) SetSourceMethodologyStatusNil(b bool)`

 SetSourceMethodologyStatusNil sets the value for SourceMethodologyStatus to be an explicit nil

### UnsetSourceMethodologyStatus
`func (o *AnnouncementDataPoint) UnsetSourceMethodologyStatus()`

UnsetSourceMethodologyStatus ensures that no value is present for SourceMethodologyStatus, not even an explicit nil
### GetSelectedVintageKnownAtNs

`func (o *AnnouncementDataPoint) GetSelectedVintageKnownAtNs() string`

GetSelectedVintageKnownAtNs returns the SelectedVintageKnownAtNs field if non-nil, zero value otherwise.

### GetSelectedVintageKnownAtNsOk

`func (o *AnnouncementDataPoint) GetSelectedVintageKnownAtNsOk() (*string, bool)`

GetSelectedVintageKnownAtNsOk returns a tuple with the SelectedVintageKnownAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedVintageKnownAtNs

`func (o *AnnouncementDataPoint) SetSelectedVintageKnownAtNs(v string)`

SetSelectedVintageKnownAtNs sets SelectedVintageKnownAtNs field to given value.

### HasSelectedVintageKnownAtNs

`func (o *AnnouncementDataPoint) HasSelectedVintageKnownAtNs() bool`

HasSelectedVintageKnownAtNs returns a boolean if a field has been set.

### SetSelectedVintageKnownAtNsNil

`func (o *AnnouncementDataPoint) SetSelectedVintageKnownAtNsNil(b bool)`

 SetSelectedVintageKnownAtNsNil sets the value for SelectedVintageKnownAtNs to be an explicit nil

### UnsetSelectedVintageKnownAtNs
`func (o *AnnouncementDataPoint) UnsetSelectedVintageKnownAtNs()`

UnsetSelectedVintageKnownAtNs ensures that no value is present for SelectedVintageKnownAtNs, not even an explicit nil
### GetReplayVintageVerified

`func (o *AnnouncementDataPoint) GetReplayVintageVerified() bool`

GetReplayVintageVerified returns the ReplayVintageVerified field if non-nil, zero value otherwise.

### GetReplayVintageVerifiedOk

`func (o *AnnouncementDataPoint) GetReplayVintageVerifiedOk() (*bool, bool)`

GetReplayVintageVerifiedOk returns a tuple with the ReplayVintageVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayVintageVerified

`func (o *AnnouncementDataPoint) SetReplayVintageVerified(v bool)`

SetReplayVintageVerified sets ReplayVintageVerified field to given value.

### HasReplayVintageVerified

`func (o *AnnouncementDataPoint) HasReplayVintageVerified() bool`

HasReplayVintageVerified returns a boolean if a field has been set.

### SetReplayVintageVerifiedNil

`func (o *AnnouncementDataPoint) SetReplayVintageVerifiedNil(b bool)`

 SetReplayVintageVerifiedNil sets the value for ReplayVintageVerified to be an explicit nil

### UnsetReplayVintageVerified
`func (o *AnnouncementDataPoint) UnsetReplayVintageVerified()`

UnsetReplayVintageVerified ensures that no value is present for ReplayVintageVerified, not even an explicit nil
### GetOfficialPlannedReleaseDatetime

`func (o *AnnouncementDataPoint) GetOfficialPlannedReleaseDatetime() int64`

GetOfficialPlannedReleaseDatetime returns the OfficialPlannedReleaseDatetime field if non-nil, zero value otherwise.

### GetOfficialPlannedReleaseDatetimeOk

`func (o *AnnouncementDataPoint) GetOfficialPlannedReleaseDatetimeOk() (*int64, bool)`

GetOfficialPlannedReleaseDatetimeOk returns a tuple with the OfficialPlannedReleaseDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialPlannedReleaseDatetime

`func (o *AnnouncementDataPoint) SetOfficialPlannedReleaseDatetime(v int64)`

SetOfficialPlannedReleaseDatetime sets OfficialPlannedReleaseDatetime field to given value.

### HasOfficialPlannedReleaseDatetime

`func (o *AnnouncementDataPoint) HasOfficialPlannedReleaseDatetime() bool`

HasOfficialPlannedReleaseDatetime returns a boolean if a field has been set.

### SetOfficialPlannedReleaseDatetimeNil

`func (o *AnnouncementDataPoint) SetOfficialPlannedReleaseDatetimeNil(b bool)`

 SetOfficialPlannedReleaseDatetimeNil sets the value for OfficialPlannedReleaseDatetime to be an explicit nil

### UnsetOfficialPlannedReleaseDatetime
`func (o *AnnouncementDataPoint) UnsetOfficialPlannedReleaseDatetime()`

UnsetOfficialPlannedReleaseDatetime ensures that no value is present for OfficialPlannedReleaseDatetime, not even an explicit nil
### GetOfficialPlannedReleaseDatetimeLocal

`func (o *AnnouncementDataPoint) GetOfficialPlannedReleaseDatetimeLocal() string`

GetOfficialPlannedReleaseDatetimeLocal returns the OfficialPlannedReleaseDatetimeLocal field if non-nil, zero value otherwise.

### GetOfficialPlannedReleaseDatetimeLocalOk

`func (o *AnnouncementDataPoint) GetOfficialPlannedReleaseDatetimeLocalOk() (*string, bool)`

GetOfficialPlannedReleaseDatetimeLocalOk returns a tuple with the OfficialPlannedReleaseDatetimeLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialPlannedReleaseDatetimeLocal

`func (o *AnnouncementDataPoint) SetOfficialPlannedReleaseDatetimeLocal(v string)`

SetOfficialPlannedReleaseDatetimeLocal sets OfficialPlannedReleaseDatetimeLocal field to given value.

### HasOfficialPlannedReleaseDatetimeLocal

`func (o *AnnouncementDataPoint) HasOfficialPlannedReleaseDatetimeLocal() bool`

HasOfficialPlannedReleaseDatetimeLocal returns a boolean if a field has been set.

### SetOfficialPlannedReleaseDatetimeLocalNil

`func (o *AnnouncementDataPoint) SetOfficialPlannedReleaseDatetimeLocalNil(b bool)`

 SetOfficialPlannedReleaseDatetimeLocalNil sets the value for OfficialPlannedReleaseDatetimeLocal to be an explicit nil

### UnsetOfficialPlannedReleaseDatetimeLocal
`func (o *AnnouncementDataPoint) UnsetOfficialPlannedReleaseDatetimeLocal()`

UnsetOfficialPlannedReleaseDatetimeLocal ensures that no value is present for OfficialPlannedReleaseDatetimeLocal, not even an explicit nil
### GetOfficialActualReleaseDatetime

`func (o *AnnouncementDataPoint) GetOfficialActualReleaseDatetime() int64`

GetOfficialActualReleaseDatetime returns the OfficialActualReleaseDatetime field if non-nil, zero value otherwise.

### GetOfficialActualReleaseDatetimeOk

`func (o *AnnouncementDataPoint) GetOfficialActualReleaseDatetimeOk() (*int64, bool)`

GetOfficialActualReleaseDatetimeOk returns a tuple with the OfficialActualReleaseDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialActualReleaseDatetime

`func (o *AnnouncementDataPoint) SetOfficialActualReleaseDatetime(v int64)`

SetOfficialActualReleaseDatetime sets OfficialActualReleaseDatetime field to given value.

### HasOfficialActualReleaseDatetime

`func (o *AnnouncementDataPoint) HasOfficialActualReleaseDatetime() bool`

HasOfficialActualReleaseDatetime returns a boolean if a field has been set.

### SetOfficialActualReleaseDatetimeNil

`func (o *AnnouncementDataPoint) SetOfficialActualReleaseDatetimeNil(b bool)`

 SetOfficialActualReleaseDatetimeNil sets the value for OfficialActualReleaseDatetime to be an explicit nil

### UnsetOfficialActualReleaseDatetime
`func (o *AnnouncementDataPoint) UnsetOfficialActualReleaseDatetime()`

UnsetOfficialActualReleaseDatetime ensures that no value is present for OfficialActualReleaseDatetime, not even an explicit nil
### GetOfficialActualReleaseDatetimeLocal

`func (o *AnnouncementDataPoint) GetOfficialActualReleaseDatetimeLocal() string`

GetOfficialActualReleaseDatetimeLocal returns the OfficialActualReleaseDatetimeLocal field if non-nil, zero value otherwise.

### GetOfficialActualReleaseDatetimeLocalOk

`func (o *AnnouncementDataPoint) GetOfficialActualReleaseDatetimeLocalOk() (*string, bool)`

GetOfficialActualReleaseDatetimeLocalOk returns a tuple with the OfficialActualReleaseDatetimeLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialActualReleaseDatetimeLocal

`func (o *AnnouncementDataPoint) SetOfficialActualReleaseDatetimeLocal(v string)`

SetOfficialActualReleaseDatetimeLocal sets OfficialActualReleaseDatetimeLocal field to given value.

### HasOfficialActualReleaseDatetimeLocal

`func (o *AnnouncementDataPoint) HasOfficialActualReleaseDatetimeLocal() bool`

HasOfficialActualReleaseDatetimeLocal returns a boolean if a field has been set.

### SetOfficialActualReleaseDatetimeLocalNil

`func (o *AnnouncementDataPoint) SetOfficialActualReleaseDatetimeLocalNil(b bool)`

 SetOfficialActualReleaseDatetimeLocalNil sets the value for OfficialActualReleaseDatetimeLocal to be an explicit nil

### UnsetOfficialActualReleaseDatetimeLocal
`func (o *AnnouncementDataPoint) UnsetOfficialActualReleaseDatetimeLocal()`

UnsetOfficialActualReleaseDatetimeLocal ensures that no value is present for OfficialActualReleaseDatetimeLocal, not even an explicit nil
### GetCollectedAtNs

`func (o *AnnouncementDataPoint) GetCollectedAtNs() int64`

GetCollectedAtNs returns the CollectedAtNs field if non-nil, zero value otherwise.

### GetCollectedAtNsOk

`func (o *AnnouncementDataPoint) GetCollectedAtNsOk() (*int64, bool)`

GetCollectedAtNsOk returns a tuple with the CollectedAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectedAtNs

`func (o *AnnouncementDataPoint) SetCollectedAtNs(v int64)`

SetCollectedAtNs sets CollectedAtNs field to given value.

### HasCollectedAtNs

`func (o *AnnouncementDataPoint) HasCollectedAtNs() bool`

HasCollectedAtNs returns a boolean if a field has been set.

### SetCollectedAtNsNil

`func (o *AnnouncementDataPoint) SetCollectedAtNsNil(b bool)`

 SetCollectedAtNsNil sets the value for CollectedAtNs to be an explicit nil

### UnsetCollectedAtNs
`func (o *AnnouncementDataPoint) UnsetCollectedAtNs()`

UnsetCollectedAtNs ensures that no value is present for CollectedAtNs, not even an explicit nil
### GetCollectedAtIso

`func (o *AnnouncementDataPoint) GetCollectedAtIso() string`

GetCollectedAtIso returns the CollectedAtIso field if non-nil, zero value otherwise.

### GetCollectedAtIsoOk

`func (o *AnnouncementDataPoint) GetCollectedAtIsoOk() (*string, bool)`

GetCollectedAtIsoOk returns a tuple with the CollectedAtIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectedAtIso

`func (o *AnnouncementDataPoint) SetCollectedAtIso(v string)`

SetCollectedAtIso sets CollectedAtIso field to given value.

### HasCollectedAtIso

`func (o *AnnouncementDataPoint) HasCollectedAtIso() bool`

HasCollectedAtIso returns a boolean if a field has been set.

### SetCollectedAtIsoNil

`func (o *AnnouncementDataPoint) SetCollectedAtIsoNil(b bool)`

 SetCollectedAtIsoNil sets the value for CollectedAtIso to be an explicit nil

### UnsetCollectedAtIso
`func (o *AnnouncementDataPoint) UnsetCollectedAtIso()`

UnsetCollectedAtIso ensures that no value is present for CollectedAtIso, not even an explicit nil
### GetReleaseTiming

`func (o *AnnouncementDataPoint) GetReleaseTiming() ReleaseTiming`

GetReleaseTiming returns the ReleaseTiming field if non-nil, zero value otherwise.

### GetReleaseTimingOk

`func (o *AnnouncementDataPoint) GetReleaseTimingOk() (*ReleaseTiming, bool)`

GetReleaseTimingOk returns a tuple with the ReleaseTiming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTiming

`func (o *AnnouncementDataPoint) SetReleaseTiming(v ReleaseTiming)`

SetReleaseTiming sets ReleaseTiming field to given value.

### HasReleaseTiming

`func (o *AnnouncementDataPoint) HasReleaseTiming() bool`

HasReleaseTiming returns a boolean if a field has been set.

### GetIngestionLatencyMs

`func (o *AnnouncementDataPoint) GetIngestionLatencyMs() float64`

GetIngestionLatencyMs returns the IngestionLatencyMs field if non-nil, zero value otherwise.

### GetIngestionLatencyMsOk

`func (o *AnnouncementDataPoint) GetIngestionLatencyMsOk() (*float64, bool)`

GetIngestionLatencyMsOk returns a tuple with the IngestionLatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionLatencyMs

`func (o *AnnouncementDataPoint) SetIngestionLatencyMs(v float64)`

SetIngestionLatencyMs sets IngestionLatencyMs field to given value.

### HasIngestionLatencyMs

`func (o *AnnouncementDataPoint) HasIngestionLatencyMs() bool`

HasIngestionLatencyMs returns a boolean if a field has been set.

### SetIngestionLatencyMsNil

`func (o *AnnouncementDataPoint) SetIngestionLatencyMsNil(b bool)`

 SetIngestionLatencyMsNil sets the value for IngestionLatencyMs to be an explicit nil

### UnsetIngestionLatencyMs
`func (o *AnnouncementDataPoint) UnsetIngestionLatencyMs()`

UnsetIngestionLatencyMs ensures that no value is present for IngestionLatencyMs, not even an explicit nil
### GetIngestionLatencyReference

`func (o *AnnouncementDataPoint) GetIngestionLatencyReference() string`

GetIngestionLatencyReference returns the IngestionLatencyReference field if non-nil, zero value otherwise.

### GetIngestionLatencyReferenceOk

`func (o *AnnouncementDataPoint) GetIngestionLatencyReferenceOk() (*string, bool)`

GetIngestionLatencyReferenceOk returns a tuple with the IngestionLatencyReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionLatencyReference

`func (o *AnnouncementDataPoint) SetIngestionLatencyReference(v string)`

SetIngestionLatencyReference sets IngestionLatencyReference field to given value.

### HasIngestionLatencyReference

`func (o *AnnouncementDataPoint) HasIngestionLatencyReference() bool`

HasIngestionLatencyReference returns a boolean if a field has been set.

### SetIngestionLatencyReferenceNil

`func (o *AnnouncementDataPoint) SetIngestionLatencyReferenceNil(b bool)`

 SetIngestionLatencyReferenceNil sets the value for IngestionLatencyReference to be an explicit nil

### UnsetIngestionLatencyReference
`func (o *AnnouncementDataPoint) UnsetIngestionLatencyReference()`

UnsetIngestionLatencyReference ensures that no value is present for IngestionLatencyReference, not even an explicit nil
### GetPctChange

`func (o *AnnouncementDataPoint) GetPctChange() float64`

GetPctChange returns the PctChange field if non-nil, zero value otherwise.

### GetPctChangeOk

`func (o *AnnouncementDataPoint) GetPctChangeOk() (*float64, bool)`

GetPctChangeOk returns a tuple with the PctChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctChange

`func (o *AnnouncementDataPoint) SetPctChange(v float64)`

SetPctChange sets PctChange field to given value.

### HasPctChange

`func (o *AnnouncementDataPoint) HasPctChange() bool`

HasPctChange returns a boolean if a field has been set.

### SetPctChangeNil

`func (o *AnnouncementDataPoint) SetPctChangeNil(b bool)`

 SetPctChangeNil sets the value for PctChange to be an explicit nil

### UnsetPctChange
`func (o *AnnouncementDataPoint) UnsetPctChange()`

UnsetPctChange ensures that no value is present for PctChange, not even an explicit nil
### GetPctChangeYoy

`func (o *AnnouncementDataPoint) GetPctChangeYoy() float64`

GetPctChangeYoy returns the PctChangeYoy field if non-nil, zero value otherwise.

### GetPctChangeYoyOk

`func (o *AnnouncementDataPoint) GetPctChangeYoyOk() (*float64, bool)`

GetPctChangeYoyOk returns a tuple with the PctChangeYoy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctChangeYoy

`func (o *AnnouncementDataPoint) SetPctChangeYoy(v float64)`

SetPctChangeYoy sets PctChangeYoy field to given value.

### HasPctChangeYoy

`func (o *AnnouncementDataPoint) HasPctChangeYoy() bool`

HasPctChangeYoy returns a boolean if a field has been set.

### SetPctChangeYoyNil

`func (o *AnnouncementDataPoint) SetPctChangeYoyNil(b bool)`

 SetPctChangeYoyNil sets the value for PctChangeYoy to be an explicit nil

### UnsetPctChangeYoy
`func (o *AnnouncementDataPoint) UnsetPctChangeYoy()`

UnsetPctChangeYoy ensures that no value is present for PctChangeYoy, not even an explicit nil
### GetPctChangeQoq

`func (o *AnnouncementDataPoint) GetPctChangeQoq() float64`

GetPctChangeQoq returns the PctChangeQoq field if non-nil, zero value otherwise.

### GetPctChangeQoqOk

`func (o *AnnouncementDataPoint) GetPctChangeQoqOk() (*float64, bool)`

GetPctChangeQoqOk returns a tuple with the PctChangeQoq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctChangeQoq

`func (o *AnnouncementDataPoint) SetPctChangeQoq(v float64)`

SetPctChangeQoq sets PctChangeQoq field to given value.

### HasPctChangeQoq

`func (o *AnnouncementDataPoint) HasPctChangeQoq() bool`

HasPctChangeQoq returns a boolean if a field has been set.

### SetPctChangeQoqNil

`func (o *AnnouncementDataPoint) SetPctChangeQoqNil(b bool)`

 SetPctChangeQoqNil sets the value for PctChangeQoq to be an explicit nil

### UnsetPctChangeQoq
`func (o *AnnouncementDataPoint) UnsetPctChangeQoq()`

UnsetPctChangeQoq ensures that no value is present for PctChangeQoq, not even an explicit nil
### GetPctChangeMom

`func (o *AnnouncementDataPoint) GetPctChangeMom() float64`

GetPctChangeMom returns the PctChangeMom field if non-nil, zero value otherwise.

### GetPctChangeMomOk

`func (o *AnnouncementDataPoint) GetPctChangeMomOk() (*float64, bool)`

GetPctChangeMomOk returns a tuple with the PctChangeMom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctChangeMom

`func (o *AnnouncementDataPoint) SetPctChangeMom(v float64)`

SetPctChangeMom sets PctChangeMom field to given value.

### HasPctChangeMom

`func (o *AnnouncementDataPoint) HasPctChangeMom() bool`

HasPctChangeMom returns a boolean if a field has been set.

### SetPctChangeMomNil

`func (o *AnnouncementDataPoint) SetPctChangeMomNil(b bool)`

 SetPctChangeMomNil sets the value for PctChangeMom to be an explicit nil

### UnsetPctChangeMom
`func (o *AnnouncementDataPoint) UnsetPctChangeMom()`

UnsetPctChangeMom ensures that no value is present for PctChangeMom, not even an explicit nil
### GetPctChange12m

`func (o *AnnouncementDataPoint) GetPctChange12m() float64`

GetPctChange12m returns the PctChange12m field if non-nil, zero value otherwise.

### GetPctChange12mOk

`func (o *AnnouncementDataPoint) GetPctChange12mOk() (*float64, bool)`

GetPctChange12mOk returns a tuple with the PctChange12m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctChange12m

`func (o *AnnouncementDataPoint) SetPctChange12m(v float64)`

SetPctChange12m sets PctChange12m field to given value.

### HasPctChange12m

`func (o *AnnouncementDataPoint) HasPctChange12m() bool`

HasPctChange12m returns a boolean if a field has been set.

### SetPctChange12mNil

`func (o *AnnouncementDataPoint) SetPctChange12mNil(b bool)`

 SetPctChange12mNil sets the value for PctChange12m to be an explicit nil

### UnsetPctChange12m
`func (o *AnnouncementDataPoint) UnsetPctChange12m()`

UnsetPctChange12m ensures that no value is present for PctChange12m, not even an explicit nil
### GetRevisions

`func (o *AnnouncementDataPoint) GetRevisions() []RevisionEntry`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *AnnouncementDataPoint) GetRevisionsOk() (*[]RevisionEntry, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *AnnouncementDataPoint) SetRevisions(v []RevisionEntry)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *AnnouncementDataPoint) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### SetRevisionsNil

`func (o *AnnouncementDataPoint) SetRevisionsNil(b bool)`

 SetRevisionsNil sets the value for Revisions to be an explicit nil

### UnsetRevisions
`func (o *AnnouncementDataPoint) UnsetRevisions()`

UnsetRevisions ensures that no value is present for Revisions, not even an explicit nil
### GetOutsideRequestedWindow

`func (o *AnnouncementDataPoint) GetOutsideRequestedWindow() bool`

GetOutsideRequestedWindow returns the OutsideRequestedWindow field if non-nil, zero value otherwise.

### GetOutsideRequestedWindowOk

`func (o *AnnouncementDataPoint) GetOutsideRequestedWindowOk() (*bool, bool)`

GetOutsideRequestedWindowOk returns a tuple with the OutsideRequestedWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideRequestedWindow

`func (o *AnnouncementDataPoint) SetOutsideRequestedWindow(v bool)`

SetOutsideRequestedWindow sets OutsideRequestedWindow field to given value.

### HasOutsideRequestedWindow

`func (o *AnnouncementDataPoint) HasOutsideRequestedWindow() bool`

HasOutsideRequestedWindow returns a boolean if a field has been set.

### SetOutsideRequestedWindowNil

`func (o *AnnouncementDataPoint) SetOutsideRequestedWindowNil(b bool)`

 SetOutsideRequestedWindowNil sets the value for OutsideRequestedWindow to be an explicit nil

### UnsetOutsideRequestedWindow
`func (o *AnnouncementDataPoint) UnsetOutsideRequestedWindow()`

UnsetOutsideRequestedWindow ensures that no value is present for OutsideRequestedWindow, not even an explicit nil
### GetRequestedStartDate

`func (o *AnnouncementDataPoint) GetRequestedStartDate() string`

GetRequestedStartDate returns the RequestedStartDate field if non-nil, zero value otherwise.

### GetRequestedStartDateOk

`func (o *AnnouncementDataPoint) GetRequestedStartDateOk() (*string, bool)`

GetRequestedStartDateOk returns a tuple with the RequestedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedStartDate

`func (o *AnnouncementDataPoint) SetRequestedStartDate(v string)`

SetRequestedStartDate sets RequestedStartDate field to given value.

### HasRequestedStartDate

`func (o *AnnouncementDataPoint) HasRequestedStartDate() bool`

HasRequestedStartDate returns a boolean if a field has been set.

### SetRequestedStartDateNil

`func (o *AnnouncementDataPoint) SetRequestedStartDateNil(b bool)`

 SetRequestedStartDateNil sets the value for RequestedStartDate to be an explicit nil

### UnsetRequestedStartDate
`func (o *AnnouncementDataPoint) UnsetRequestedStartDate()`

UnsetRequestedStartDate ensures that no value is present for RequestedStartDate, not even an explicit nil
### GetRequestedEndDate

`func (o *AnnouncementDataPoint) GetRequestedEndDate() string`

GetRequestedEndDate returns the RequestedEndDate field if non-nil, zero value otherwise.

### GetRequestedEndDateOk

`func (o *AnnouncementDataPoint) GetRequestedEndDateOk() (*string, bool)`

GetRequestedEndDateOk returns a tuple with the RequestedEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedEndDate

`func (o *AnnouncementDataPoint) SetRequestedEndDate(v string)`

SetRequestedEndDate sets RequestedEndDate field to given value.

### HasRequestedEndDate

`func (o *AnnouncementDataPoint) HasRequestedEndDate() bool`

HasRequestedEndDate returns a boolean if a field has been set.

### SetRequestedEndDateNil

`func (o *AnnouncementDataPoint) SetRequestedEndDateNil(b bool)`

 SetRequestedEndDateNil sets the value for RequestedEndDate to be an explicit nil

### UnsetRequestedEndDate
`func (o *AnnouncementDataPoint) UnsetRequestedEndDate()`

UnsetRequestedEndDate ensures that no value is present for RequestedEndDate, not even an explicit nil
### GetReturnedReason

`func (o *AnnouncementDataPoint) GetReturnedReason() string`

GetReturnedReason returns the ReturnedReason field if non-nil, zero value otherwise.

### GetReturnedReasonOk

`func (o *AnnouncementDataPoint) GetReturnedReasonOk() (*string, bool)`

GetReturnedReasonOk returns a tuple with the ReturnedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedReason

`func (o *AnnouncementDataPoint) SetReturnedReason(v string)`

SetReturnedReason sets ReturnedReason field to given value.

### HasReturnedReason

`func (o *AnnouncementDataPoint) HasReturnedReason() bool`

HasReturnedReason returns a boolean if a field has been set.

### SetReturnedReasonNil

`func (o *AnnouncementDataPoint) SetReturnedReasonNil(b bool)`

 SetReturnedReasonNil sets the value for ReturnedReason to be an explicit nil

### UnsetReturnedReason
`func (o *AnnouncementDataPoint) UnsetReturnedReason()`

UnsetReturnedReason ensures that no value is present for ReturnedReason, not even an explicit nil
### GetStalenessDays

`func (o *AnnouncementDataPoint) GetStalenessDays() int64`

GetStalenessDays returns the StalenessDays field if non-nil, zero value otherwise.

### GetStalenessDaysOk

`func (o *AnnouncementDataPoint) GetStalenessDaysOk() (*int64, bool)`

GetStalenessDaysOk returns a tuple with the StalenessDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStalenessDays

`func (o *AnnouncementDataPoint) SetStalenessDays(v int64)`

SetStalenessDays sets StalenessDays field to given value.

### HasStalenessDays

`func (o *AnnouncementDataPoint) HasStalenessDays() bool`

HasStalenessDays returns a boolean if a field has been set.

### SetStalenessDaysNil

`func (o *AnnouncementDataPoint) SetStalenessDaysNil(b bool)`

 SetStalenessDaysNil sets the value for StalenessDays to be an explicit nil

### UnsetStalenessDays
`func (o *AnnouncementDataPoint) UnsetStalenessDays()`

UnsetStalenessDays ensures that no value is present for StalenessDays, not even an explicit nil
### GetCanonicalIndicator

`func (o *AnnouncementDataPoint) GetCanonicalIndicator() string`

GetCanonicalIndicator returns the CanonicalIndicator field if non-nil, zero value otherwise.

### GetCanonicalIndicatorOk

`func (o *AnnouncementDataPoint) GetCanonicalIndicatorOk() (*string, bool)`

GetCanonicalIndicatorOk returns a tuple with the CanonicalIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalIndicator

`func (o *AnnouncementDataPoint) SetCanonicalIndicator(v string)`

SetCanonicalIndicator sets CanonicalIndicator field to given value.

### HasCanonicalIndicator

`func (o *AnnouncementDataPoint) HasCanonicalIndicator() bool`

HasCanonicalIndicator returns a boolean if a field has been set.

### SetCanonicalIndicatorNil

`func (o *AnnouncementDataPoint) SetCanonicalIndicatorNil(b bool)`

 SetCanonicalIndicatorNil sets the value for CanonicalIndicator to be an explicit nil

### UnsetCanonicalIndicator
`func (o *AnnouncementDataPoint) UnsetCanonicalIndicator()`

UnsetCanonicalIndicator ensures that no value is present for CanonicalIndicator, not even an explicit nil
### GetRawIndicator

`func (o *AnnouncementDataPoint) GetRawIndicator() string`

GetRawIndicator returns the RawIndicator field if non-nil, zero value otherwise.

### GetRawIndicatorOk

`func (o *AnnouncementDataPoint) GetRawIndicatorOk() (*string, bool)`

GetRawIndicatorOk returns a tuple with the RawIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawIndicator

`func (o *AnnouncementDataPoint) SetRawIndicator(v string)`

SetRawIndicator sets RawIndicator field to given value.

### HasRawIndicator

`func (o *AnnouncementDataPoint) HasRawIndicator() bool`

HasRawIndicator returns a boolean if a field has been set.

### SetRawIndicatorNil

`func (o *AnnouncementDataPoint) SetRawIndicatorNil(b bool)`

 SetRawIndicatorNil sets the value for RawIndicator to be an explicit nil

### UnsetRawIndicator
`func (o *AnnouncementDataPoint) UnsetRawIndicator()`

UnsetRawIndicator ensures that no value is present for RawIndicator, not even an explicit nil
### GetRemapApplied

`func (o *AnnouncementDataPoint) GetRemapApplied() bool`

GetRemapApplied returns the RemapApplied field if non-nil, zero value otherwise.

### GetRemapAppliedOk

`func (o *AnnouncementDataPoint) GetRemapAppliedOk() (*bool, bool)`

GetRemapAppliedOk returns a tuple with the RemapApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemapApplied

`func (o *AnnouncementDataPoint) SetRemapApplied(v bool)`

SetRemapApplied sets RemapApplied field to given value.

### HasRemapApplied

`func (o *AnnouncementDataPoint) HasRemapApplied() bool`

HasRemapApplied returns a boolean if a field has been set.

### SetRemapAppliedNil

`func (o *AnnouncementDataPoint) SetRemapAppliedNil(b bool)`

 SetRemapAppliedNil sets the value for RemapApplied to be an explicit nil

### UnsetRemapApplied
`func (o *AnnouncementDataPoint) UnsetRemapApplied()`

UnsetRemapApplied ensures that no value is present for RemapApplied, not even an explicit nil
### GetRemapRuleId

`func (o *AnnouncementDataPoint) GetRemapRuleId() string`

GetRemapRuleId returns the RemapRuleId field if non-nil, zero value otherwise.

### GetRemapRuleIdOk

`func (o *AnnouncementDataPoint) GetRemapRuleIdOk() (*string, bool)`

GetRemapRuleIdOk returns a tuple with the RemapRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemapRuleId

`func (o *AnnouncementDataPoint) SetRemapRuleId(v string)`

SetRemapRuleId sets RemapRuleId field to given value.

### HasRemapRuleId

`func (o *AnnouncementDataPoint) HasRemapRuleId() bool`

HasRemapRuleId returns a boolean if a field has been set.

### SetRemapRuleIdNil

`func (o *AnnouncementDataPoint) SetRemapRuleIdNil(b bool)`

 SetRemapRuleIdNil sets the value for RemapRuleId to be an explicit nil

### UnsetRemapRuleId
`func (o *AnnouncementDataPoint) UnsetRemapRuleId()`

UnsetRemapRuleId ensures that no value is present for RemapRuleId, not even an explicit nil
### GetRemapSegmentId

`func (o *AnnouncementDataPoint) GetRemapSegmentId() string`

GetRemapSegmentId returns the RemapSegmentId field if non-nil, zero value otherwise.

### GetRemapSegmentIdOk

`func (o *AnnouncementDataPoint) GetRemapSegmentIdOk() (*string, bool)`

GetRemapSegmentIdOk returns a tuple with the RemapSegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemapSegmentId

`func (o *AnnouncementDataPoint) SetRemapSegmentId(v string)`

SetRemapSegmentId sets RemapSegmentId field to given value.

### HasRemapSegmentId

`func (o *AnnouncementDataPoint) HasRemapSegmentId() bool`

HasRemapSegmentId returns a boolean if a field has been set.

### SetRemapSegmentIdNil

`func (o *AnnouncementDataPoint) SetRemapSegmentIdNil(b bool)`

 SetRemapSegmentIdNil sets the value for RemapSegmentId to be an explicit nil

### UnsetRemapSegmentId
`func (o *AnnouncementDataPoint) UnsetRemapSegmentId()`

UnsetRemapSegmentId ensures that no value is present for RemapSegmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


