# AnnouncementChangeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**Currency** | **string** |  | 
**Indicator** | **string** |  | 
**RecordsWritten** | Pointer to **NullableInt64** |  | [optional] 
**Timestamp** | Pointer to **NullableFloat64** |  | [optional] 
**ReleaseTimestamp** | Pointer to **NullableFloat64** |  | [optional] 
**LatestAnnouncement** | Pointer to [**NullableReleaseDeliveryAnnouncement**](ReleaseDeliveryAnnouncement.md) |  | [optional] 
**DeliveryMode** | Pointer to **NullableString** |  | [optional] 
**OriginDeliveryMode** | Pointer to **NullableString** |  | [optional] 
**SourceFreshAtNs** | Pointer to **NullableFloat64** |  | [optional] 
**StreamReadyAtNs** | Pointer to **NullableFloat64** |  | [optional] 
**ServerSentAtNs** | Pointer to **NullableFloat64** | Time FXMacroData sent this event, as Unix epoch nanoseconds (UTC). | [optional] 
**AcknowledgementEndpoint** | Pointer to **NullableString** | Endpoint for acknowledging receipt of a streamed event. Present only on &#x60;/v1/stream/events&#x60; deliveries; never set on &#x60;/v1/announcements/changes&#x60;. | [optional] 
**AgeMs** | Pointer to **NullableFloat64** | Event age in milliseconds when it was sent. Present only on &#x60;/v1/stream/events&#x60; deliveries; never set on &#x60;/v1/announcements/changes&#x60;. | [optional] 
**StaleAfterMs** | Pointer to **NullableFloat64** | Age in milliseconds after which a streamed event is considered stale. Present only on &#x60;/v1/stream/events&#x60; deliveries; never set on &#x60;/v1/announcements/changes&#x60;. | [optional] 
**Stale** | Pointer to **NullableBool** | True when the event was older than &#x60;stale_after_ms&#x60; when sent. Present only on &#x60;/v1/stream/events&#x60; deliveries; never set on &#x60;/v1/announcements/changes&#x60;. | [optional] 
**LateDelivery** | Pointer to **NullableBool** | True when the event was sent after its staleness threshold. Present only on &#x60;/v1/stream/events&#x60; deliveries; never set on &#x60;/v1/announcements/changes&#x60;. | [optional] 
**LateEventsAreDelivered** | Pointer to **NullableBool** |  | [optional] 
**ScheduledReleaseAtNs** | Pointer to **NullableFloat64** |  | [optional] 
**ScheduledToServerSendMs** | Pointer to **NullableFloat64** | Milliseconds from the scheduled release time to the time FXMacroData sent this event. | [optional] 
**SourceLate** | Pointer to **NullableBool** |  | [optional] 
**SourceLateByMs** | Pointer to **NullableFloat64** |  | [optional] 
**SourceDelayProven** | Pointer to **NullableBool** |  | [optional] 
**PlatformDeliveryAfterSourceMs** | Pointer to **NullableFloat64** |  | [optional] 
**PollingStartedAtNs** | Pointer to **NullableFloat64** |  | [optional] 
**PollingStartedLagMs** | Pointer to **NullableFloat64** |  | [optional] 
**PollingStartedBeforeScheduledRelease** | Pointer to **NullableBool** |  | [optional] 
**PreFreshStalePollCount** | Pointer to **NullableFloat64** |  | [optional] 
**LastStaleFetchCompletedAtNs** | Pointer to **NullableFloat64** |  | [optional] 
**FirstFreshResponseCompletedAtNs** | Pointer to **NullableFloat64** |  | [optional] 
**OfficialSourceStaleAfterScheduledRelease** | Pointer to **NullableBool** |  | [optional] 
**ScheduledToFirstFreshResponseMs** | Pointer to **NullableFloat64** |  | [optional] 
**SourceFreshnessObservationWindowStartAtNs** | Pointer to **NullableFloat64** |  | [optional] 
**SourceFreshnessObservationWindowEndAtNs** | Pointer to **NullableFloat64** |  | [optional] 
**SourceFreshnessObservationWindowMs** | Pointer to **NullableFloat64** |  | [optional] 
**SourceFreshnessObservationBasis** | Pointer to **NullableString** |  | [optional] 
**SourceDelayAttribution** | Pointer to **NullableString** |  | [optional] 
**FxmdDeliveryAttribution** | Pointer to **NullableString** |  | [optional] 
**FxmdAfterSourceMs** | Pointer to **NullableFloat64** |  | [optional] 
**FxmdAfterSourceSloMet** | Pointer to **NullableBool** |  | [optional] 
**CatchupSource** | Pointer to **NullableString** |  | [optional] 
**CatchupDelayMs** | Pointer to **NullableFloat64** |  | [optional] 
**RecoveryClass** | Pointer to **NullableString** |  | [optional] 
**SubsecondStatus** | Pointer to **NullableString** |  | [optional] 
**SubsecondOperationalStatus** | Pointer to **NullableString** |  | [optional] 
**SubsecondGuaranteeActive** | Pointer to **NullableBool** |  | [optional] 
**SubsecondContractOutcome** | Pointer to **NullableString** |  | [optional] 
**SubsecondContractBreached** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewAnnouncementChangeEvent

`func NewAnnouncementChangeEvent(eventId string, currency string, indicator string, ) *AnnouncementChangeEvent`

NewAnnouncementChangeEvent instantiates a new AnnouncementChangeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementChangeEventWithDefaults

`func NewAnnouncementChangeEventWithDefaults() *AnnouncementChangeEvent`

NewAnnouncementChangeEventWithDefaults instantiates a new AnnouncementChangeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *AnnouncementChangeEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *AnnouncementChangeEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *AnnouncementChangeEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetCurrency

`func (o *AnnouncementChangeEvent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AnnouncementChangeEvent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AnnouncementChangeEvent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetIndicator

`func (o *AnnouncementChangeEvent) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *AnnouncementChangeEvent) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *AnnouncementChangeEvent) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.


### GetRecordsWritten

`func (o *AnnouncementChangeEvent) GetRecordsWritten() int64`

GetRecordsWritten returns the RecordsWritten field if non-nil, zero value otherwise.

### GetRecordsWrittenOk

`func (o *AnnouncementChangeEvent) GetRecordsWrittenOk() (*int64, bool)`

GetRecordsWrittenOk returns a tuple with the RecordsWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsWritten

`func (o *AnnouncementChangeEvent) SetRecordsWritten(v int64)`

SetRecordsWritten sets RecordsWritten field to given value.

### HasRecordsWritten

`func (o *AnnouncementChangeEvent) HasRecordsWritten() bool`

HasRecordsWritten returns a boolean if a field has been set.

### SetRecordsWrittenNil

`func (o *AnnouncementChangeEvent) SetRecordsWrittenNil(b bool)`

 SetRecordsWrittenNil sets the value for RecordsWritten to be an explicit nil

### UnsetRecordsWritten
`func (o *AnnouncementChangeEvent) UnsetRecordsWritten()`

UnsetRecordsWritten ensures that no value is present for RecordsWritten, not even an explicit nil
### GetTimestamp

`func (o *AnnouncementChangeEvent) GetTimestamp() float64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AnnouncementChangeEvent) GetTimestampOk() (*float64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AnnouncementChangeEvent) SetTimestamp(v float64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AnnouncementChangeEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *AnnouncementChangeEvent) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *AnnouncementChangeEvent) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetReleaseTimestamp

`func (o *AnnouncementChangeEvent) GetReleaseTimestamp() float64`

GetReleaseTimestamp returns the ReleaseTimestamp field if non-nil, zero value otherwise.

### GetReleaseTimestampOk

`func (o *AnnouncementChangeEvent) GetReleaseTimestampOk() (*float64, bool)`

GetReleaseTimestampOk returns a tuple with the ReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTimestamp

`func (o *AnnouncementChangeEvent) SetReleaseTimestamp(v float64)`

SetReleaseTimestamp sets ReleaseTimestamp field to given value.

### HasReleaseTimestamp

`func (o *AnnouncementChangeEvent) HasReleaseTimestamp() bool`

HasReleaseTimestamp returns a boolean if a field has been set.

### SetReleaseTimestampNil

`func (o *AnnouncementChangeEvent) SetReleaseTimestampNil(b bool)`

 SetReleaseTimestampNil sets the value for ReleaseTimestamp to be an explicit nil

### UnsetReleaseTimestamp
`func (o *AnnouncementChangeEvent) UnsetReleaseTimestamp()`

UnsetReleaseTimestamp ensures that no value is present for ReleaseTimestamp, not even an explicit nil
### GetLatestAnnouncement

`func (o *AnnouncementChangeEvent) GetLatestAnnouncement() ReleaseDeliveryAnnouncement`

GetLatestAnnouncement returns the LatestAnnouncement field if non-nil, zero value otherwise.

### GetLatestAnnouncementOk

`func (o *AnnouncementChangeEvent) GetLatestAnnouncementOk() (*ReleaseDeliveryAnnouncement, bool)`

GetLatestAnnouncementOk returns a tuple with the LatestAnnouncement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAnnouncement

`func (o *AnnouncementChangeEvent) SetLatestAnnouncement(v ReleaseDeliveryAnnouncement)`

SetLatestAnnouncement sets LatestAnnouncement field to given value.

### HasLatestAnnouncement

`func (o *AnnouncementChangeEvent) HasLatestAnnouncement() bool`

HasLatestAnnouncement returns a boolean if a field has been set.

### SetLatestAnnouncementNil

`func (o *AnnouncementChangeEvent) SetLatestAnnouncementNil(b bool)`

 SetLatestAnnouncementNil sets the value for LatestAnnouncement to be an explicit nil

### UnsetLatestAnnouncement
`func (o *AnnouncementChangeEvent) UnsetLatestAnnouncement()`

UnsetLatestAnnouncement ensures that no value is present for LatestAnnouncement, not even an explicit nil
### GetDeliveryMode

`func (o *AnnouncementChangeEvent) GetDeliveryMode() string`

GetDeliveryMode returns the DeliveryMode field if non-nil, zero value otherwise.

### GetDeliveryModeOk

`func (o *AnnouncementChangeEvent) GetDeliveryModeOk() (*string, bool)`

GetDeliveryModeOk returns a tuple with the DeliveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMode

`func (o *AnnouncementChangeEvent) SetDeliveryMode(v string)`

SetDeliveryMode sets DeliveryMode field to given value.

### HasDeliveryMode

`func (o *AnnouncementChangeEvent) HasDeliveryMode() bool`

HasDeliveryMode returns a boolean if a field has been set.

### SetDeliveryModeNil

`func (o *AnnouncementChangeEvent) SetDeliveryModeNil(b bool)`

 SetDeliveryModeNil sets the value for DeliveryMode to be an explicit nil

### UnsetDeliveryMode
`func (o *AnnouncementChangeEvent) UnsetDeliveryMode()`

UnsetDeliveryMode ensures that no value is present for DeliveryMode, not even an explicit nil
### GetOriginDeliveryMode

`func (o *AnnouncementChangeEvent) GetOriginDeliveryMode() string`

GetOriginDeliveryMode returns the OriginDeliveryMode field if non-nil, zero value otherwise.

### GetOriginDeliveryModeOk

`func (o *AnnouncementChangeEvent) GetOriginDeliveryModeOk() (*string, bool)`

GetOriginDeliveryModeOk returns a tuple with the OriginDeliveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDeliveryMode

`func (o *AnnouncementChangeEvent) SetOriginDeliveryMode(v string)`

SetOriginDeliveryMode sets OriginDeliveryMode field to given value.

### HasOriginDeliveryMode

`func (o *AnnouncementChangeEvent) HasOriginDeliveryMode() bool`

HasOriginDeliveryMode returns a boolean if a field has been set.

### SetOriginDeliveryModeNil

`func (o *AnnouncementChangeEvent) SetOriginDeliveryModeNil(b bool)`

 SetOriginDeliveryModeNil sets the value for OriginDeliveryMode to be an explicit nil

### UnsetOriginDeliveryMode
`func (o *AnnouncementChangeEvent) UnsetOriginDeliveryMode()`

UnsetOriginDeliveryMode ensures that no value is present for OriginDeliveryMode, not even an explicit nil
### GetSourceFreshAtNs

`func (o *AnnouncementChangeEvent) GetSourceFreshAtNs() float64`

GetSourceFreshAtNs returns the SourceFreshAtNs field if non-nil, zero value otherwise.

### GetSourceFreshAtNsOk

`func (o *AnnouncementChangeEvent) GetSourceFreshAtNsOk() (*float64, bool)`

GetSourceFreshAtNsOk returns a tuple with the SourceFreshAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFreshAtNs

`func (o *AnnouncementChangeEvent) SetSourceFreshAtNs(v float64)`

SetSourceFreshAtNs sets SourceFreshAtNs field to given value.

### HasSourceFreshAtNs

`func (o *AnnouncementChangeEvent) HasSourceFreshAtNs() bool`

HasSourceFreshAtNs returns a boolean if a field has been set.

### SetSourceFreshAtNsNil

`func (o *AnnouncementChangeEvent) SetSourceFreshAtNsNil(b bool)`

 SetSourceFreshAtNsNil sets the value for SourceFreshAtNs to be an explicit nil

### UnsetSourceFreshAtNs
`func (o *AnnouncementChangeEvent) UnsetSourceFreshAtNs()`

UnsetSourceFreshAtNs ensures that no value is present for SourceFreshAtNs, not even an explicit nil
### GetStreamReadyAtNs

`func (o *AnnouncementChangeEvent) GetStreamReadyAtNs() float64`

GetStreamReadyAtNs returns the StreamReadyAtNs field if non-nil, zero value otherwise.

### GetStreamReadyAtNsOk

`func (o *AnnouncementChangeEvent) GetStreamReadyAtNsOk() (*float64, bool)`

GetStreamReadyAtNsOk returns a tuple with the StreamReadyAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamReadyAtNs

`func (o *AnnouncementChangeEvent) SetStreamReadyAtNs(v float64)`

SetStreamReadyAtNs sets StreamReadyAtNs field to given value.

### HasStreamReadyAtNs

`func (o *AnnouncementChangeEvent) HasStreamReadyAtNs() bool`

HasStreamReadyAtNs returns a boolean if a field has been set.

### SetStreamReadyAtNsNil

`func (o *AnnouncementChangeEvent) SetStreamReadyAtNsNil(b bool)`

 SetStreamReadyAtNsNil sets the value for StreamReadyAtNs to be an explicit nil

### UnsetStreamReadyAtNs
`func (o *AnnouncementChangeEvent) UnsetStreamReadyAtNs()`

UnsetStreamReadyAtNs ensures that no value is present for StreamReadyAtNs, not even an explicit nil
### GetServerSentAtNs

`func (o *AnnouncementChangeEvent) GetServerSentAtNs() float64`

GetServerSentAtNs returns the ServerSentAtNs field if non-nil, zero value otherwise.

### GetServerSentAtNsOk

`func (o *AnnouncementChangeEvent) GetServerSentAtNsOk() (*float64, bool)`

GetServerSentAtNsOk returns a tuple with the ServerSentAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSentAtNs

`func (o *AnnouncementChangeEvent) SetServerSentAtNs(v float64)`

SetServerSentAtNs sets ServerSentAtNs field to given value.

### HasServerSentAtNs

`func (o *AnnouncementChangeEvent) HasServerSentAtNs() bool`

HasServerSentAtNs returns a boolean if a field has been set.

### SetServerSentAtNsNil

`func (o *AnnouncementChangeEvent) SetServerSentAtNsNil(b bool)`

 SetServerSentAtNsNil sets the value for ServerSentAtNs to be an explicit nil

### UnsetServerSentAtNs
`func (o *AnnouncementChangeEvent) UnsetServerSentAtNs()`

UnsetServerSentAtNs ensures that no value is present for ServerSentAtNs, not even an explicit nil
### GetAcknowledgementEndpoint

`func (o *AnnouncementChangeEvent) GetAcknowledgementEndpoint() string`

GetAcknowledgementEndpoint returns the AcknowledgementEndpoint field if non-nil, zero value otherwise.

### GetAcknowledgementEndpointOk

`func (o *AnnouncementChangeEvent) GetAcknowledgementEndpointOk() (*string, bool)`

GetAcknowledgementEndpointOk returns a tuple with the AcknowledgementEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgementEndpoint

`func (o *AnnouncementChangeEvent) SetAcknowledgementEndpoint(v string)`

SetAcknowledgementEndpoint sets AcknowledgementEndpoint field to given value.

### HasAcknowledgementEndpoint

`func (o *AnnouncementChangeEvent) HasAcknowledgementEndpoint() bool`

HasAcknowledgementEndpoint returns a boolean if a field has been set.

### SetAcknowledgementEndpointNil

`func (o *AnnouncementChangeEvent) SetAcknowledgementEndpointNil(b bool)`

 SetAcknowledgementEndpointNil sets the value for AcknowledgementEndpoint to be an explicit nil

### UnsetAcknowledgementEndpoint
`func (o *AnnouncementChangeEvent) UnsetAcknowledgementEndpoint()`

UnsetAcknowledgementEndpoint ensures that no value is present for AcknowledgementEndpoint, not even an explicit nil
### GetAgeMs

`func (o *AnnouncementChangeEvent) GetAgeMs() float64`

GetAgeMs returns the AgeMs field if non-nil, zero value otherwise.

### GetAgeMsOk

`func (o *AnnouncementChangeEvent) GetAgeMsOk() (*float64, bool)`

GetAgeMsOk returns a tuple with the AgeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeMs

`func (o *AnnouncementChangeEvent) SetAgeMs(v float64)`

SetAgeMs sets AgeMs field to given value.

### HasAgeMs

`func (o *AnnouncementChangeEvent) HasAgeMs() bool`

HasAgeMs returns a boolean if a field has been set.

### SetAgeMsNil

`func (o *AnnouncementChangeEvent) SetAgeMsNil(b bool)`

 SetAgeMsNil sets the value for AgeMs to be an explicit nil

### UnsetAgeMs
`func (o *AnnouncementChangeEvent) UnsetAgeMs()`

UnsetAgeMs ensures that no value is present for AgeMs, not even an explicit nil
### GetStaleAfterMs

`func (o *AnnouncementChangeEvent) GetStaleAfterMs() float64`

GetStaleAfterMs returns the StaleAfterMs field if non-nil, zero value otherwise.

### GetStaleAfterMsOk

`func (o *AnnouncementChangeEvent) GetStaleAfterMsOk() (*float64, bool)`

GetStaleAfterMsOk returns a tuple with the StaleAfterMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleAfterMs

`func (o *AnnouncementChangeEvent) SetStaleAfterMs(v float64)`

SetStaleAfterMs sets StaleAfterMs field to given value.

### HasStaleAfterMs

`func (o *AnnouncementChangeEvent) HasStaleAfterMs() bool`

HasStaleAfterMs returns a boolean if a field has been set.

### SetStaleAfterMsNil

`func (o *AnnouncementChangeEvent) SetStaleAfterMsNil(b bool)`

 SetStaleAfterMsNil sets the value for StaleAfterMs to be an explicit nil

### UnsetStaleAfterMs
`func (o *AnnouncementChangeEvent) UnsetStaleAfterMs()`

UnsetStaleAfterMs ensures that no value is present for StaleAfterMs, not even an explicit nil
### GetStale

`func (o *AnnouncementChangeEvent) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *AnnouncementChangeEvent) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *AnnouncementChangeEvent) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *AnnouncementChangeEvent) HasStale() bool`

HasStale returns a boolean if a field has been set.

### SetStaleNil

`func (o *AnnouncementChangeEvent) SetStaleNil(b bool)`

 SetStaleNil sets the value for Stale to be an explicit nil

### UnsetStale
`func (o *AnnouncementChangeEvent) UnsetStale()`

UnsetStale ensures that no value is present for Stale, not even an explicit nil
### GetLateDelivery

`func (o *AnnouncementChangeEvent) GetLateDelivery() bool`

GetLateDelivery returns the LateDelivery field if non-nil, zero value otherwise.

### GetLateDeliveryOk

`func (o *AnnouncementChangeEvent) GetLateDeliveryOk() (*bool, bool)`

GetLateDeliveryOk returns a tuple with the LateDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateDelivery

`func (o *AnnouncementChangeEvent) SetLateDelivery(v bool)`

SetLateDelivery sets LateDelivery field to given value.

### HasLateDelivery

`func (o *AnnouncementChangeEvent) HasLateDelivery() bool`

HasLateDelivery returns a boolean if a field has been set.

### SetLateDeliveryNil

`func (o *AnnouncementChangeEvent) SetLateDeliveryNil(b bool)`

 SetLateDeliveryNil sets the value for LateDelivery to be an explicit nil

### UnsetLateDelivery
`func (o *AnnouncementChangeEvent) UnsetLateDelivery()`

UnsetLateDelivery ensures that no value is present for LateDelivery, not even an explicit nil
### GetLateEventsAreDelivered

`func (o *AnnouncementChangeEvent) GetLateEventsAreDelivered() bool`

GetLateEventsAreDelivered returns the LateEventsAreDelivered field if non-nil, zero value otherwise.

### GetLateEventsAreDeliveredOk

`func (o *AnnouncementChangeEvent) GetLateEventsAreDeliveredOk() (*bool, bool)`

GetLateEventsAreDeliveredOk returns a tuple with the LateEventsAreDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateEventsAreDelivered

`func (o *AnnouncementChangeEvent) SetLateEventsAreDelivered(v bool)`

SetLateEventsAreDelivered sets LateEventsAreDelivered field to given value.

### HasLateEventsAreDelivered

`func (o *AnnouncementChangeEvent) HasLateEventsAreDelivered() bool`

HasLateEventsAreDelivered returns a boolean if a field has been set.

### SetLateEventsAreDeliveredNil

`func (o *AnnouncementChangeEvent) SetLateEventsAreDeliveredNil(b bool)`

 SetLateEventsAreDeliveredNil sets the value for LateEventsAreDelivered to be an explicit nil

### UnsetLateEventsAreDelivered
`func (o *AnnouncementChangeEvent) UnsetLateEventsAreDelivered()`

UnsetLateEventsAreDelivered ensures that no value is present for LateEventsAreDelivered, not even an explicit nil
### GetScheduledReleaseAtNs

`func (o *AnnouncementChangeEvent) GetScheduledReleaseAtNs() float64`

GetScheduledReleaseAtNs returns the ScheduledReleaseAtNs field if non-nil, zero value otherwise.

### GetScheduledReleaseAtNsOk

`func (o *AnnouncementChangeEvent) GetScheduledReleaseAtNsOk() (*float64, bool)`

GetScheduledReleaseAtNsOk returns a tuple with the ScheduledReleaseAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReleaseAtNs

`func (o *AnnouncementChangeEvent) SetScheduledReleaseAtNs(v float64)`

SetScheduledReleaseAtNs sets ScheduledReleaseAtNs field to given value.

### HasScheduledReleaseAtNs

`func (o *AnnouncementChangeEvent) HasScheduledReleaseAtNs() bool`

HasScheduledReleaseAtNs returns a boolean if a field has been set.

### SetScheduledReleaseAtNsNil

`func (o *AnnouncementChangeEvent) SetScheduledReleaseAtNsNil(b bool)`

 SetScheduledReleaseAtNsNil sets the value for ScheduledReleaseAtNs to be an explicit nil

### UnsetScheduledReleaseAtNs
`func (o *AnnouncementChangeEvent) UnsetScheduledReleaseAtNs()`

UnsetScheduledReleaseAtNs ensures that no value is present for ScheduledReleaseAtNs, not even an explicit nil
### GetScheduledToServerSendMs

`func (o *AnnouncementChangeEvent) GetScheduledToServerSendMs() float64`

GetScheduledToServerSendMs returns the ScheduledToServerSendMs field if non-nil, zero value otherwise.

### GetScheduledToServerSendMsOk

`func (o *AnnouncementChangeEvent) GetScheduledToServerSendMsOk() (*float64, bool)`

GetScheduledToServerSendMsOk returns a tuple with the ScheduledToServerSendMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledToServerSendMs

`func (o *AnnouncementChangeEvent) SetScheduledToServerSendMs(v float64)`

SetScheduledToServerSendMs sets ScheduledToServerSendMs field to given value.

### HasScheduledToServerSendMs

`func (o *AnnouncementChangeEvent) HasScheduledToServerSendMs() bool`

HasScheduledToServerSendMs returns a boolean if a field has been set.

### SetScheduledToServerSendMsNil

`func (o *AnnouncementChangeEvent) SetScheduledToServerSendMsNil(b bool)`

 SetScheduledToServerSendMsNil sets the value for ScheduledToServerSendMs to be an explicit nil

### UnsetScheduledToServerSendMs
`func (o *AnnouncementChangeEvent) UnsetScheduledToServerSendMs()`

UnsetScheduledToServerSendMs ensures that no value is present for ScheduledToServerSendMs, not even an explicit nil
### GetSourceLate

`func (o *AnnouncementChangeEvent) GetSourceLate() bool`

GetSourceLate returns the SourceLate field if non-nil, zero value otherwise.

### GetSourceLateOk

`func (o *AnnouncementChangeEvent) GetSourceLateOk() (*bool, bool)`

GetSourceLateOk returns a tuple with the SourceLate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLate

`func (o *AnnouncementChangeEvent) SetSourceLate(v bool)`

SetSourceLate sets SourceLate field to given value.

### HasSourceLate

`func (o *AnnouncementChangeEvent) HasSourceLate() bool`

HasSourceLate returns a boolean if a field has been set.

### SetSourceLateNil

`func (o *AnnouncementChangeEvent) SetSourceLateNil(b bool)`

 SetSourceLateNil sets the value for SourceLate to be an explicit nil

### UnsetSourceLate
`func (o *AnnouncementChangeEvent) UnsetSourceLate()`

UnsetSourceLate ensures that no value is present for SourceLate, not even an explicit nil
### GetSourceLateByMs

`func (o *AnnouncementChangeEvent) GetSourceLateByMs() float64`

GetSourceLateByMs returns the SourceLateByMs field if non-nil, zero value otherwise.

### GetSourceLateByMsOk

`func (o *AnnouncementChangeEvent) GetSourceLateByMsOk() (*float64, bool)`

GetSourceLateByMsOk returns a tuple with the SourceLateByMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLateByMs

`func (o *AnnouncementChangeEvent) SetSourceLateByMs(v float64)`

SetSourceLateByMs sets SourceLateByMs field to given value.

### HasSourceLateByMs

`func (o *AnnouncementChangeEvent) HasSourceLateByMs() bool`

HasSourceLateByMs returns a boolean if a field has been set.

### SetSourceLateByMsNil

`func (o *AnnouncementChangeEvent) SetSourceLateByMsNil(b bool)`

 SetSourceLateByMsNil sets the value for SourceLateByMs to be an explicit nil

### UnsetSourceLateByMs
`func (o *AnnouncementChangeEvent) UnsetSourceLateByMs()`

UnsetSourceLateByMs ensures that no value is present for SourceLateByMs, not even an explicit nil
### GetSourceDelayProven

`func (o *AnnouncementChangeEvent) GetSourceDelayProven() bool`

GetSourceDelayProven returns the SourceDelayProven field if non-nil, zero value otherwise.

### GetSourceDelayProvenOk

`func (o *AnnouncementChangeEvent) GetSourceDelayProvenOk() (*bool, bool)`

GetSourceDelayProvenOk returns a tuple with the SourceDelayProven field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDelayProven

`func (o *AnnouncementChangeEvent) SetSourceDelayProven(v bool)`

SetSourceDelayProven sets SourceDelayProven field to given value.

### HasSourceDelayProven

`func (o *AnnouncementChangeEvent) HasSourceDelayProven() bool`

HasSourceDelayProven returns a boolean if a field has been set.

### SetSourceDelayProvenNil

`func (o *AnnouncementChangeEvent) SetSourceDelayProvenNil(b bool)`

 SetSourceDelayProvenNil sets the value for SourceDelayProven to be an explicit nil

### UnsetSourceDelayProven
`func (o *AnnouncementChangeEvent) UnsetSourceDelayProven()`

UnsetSourceDelayProven ensures that no value is present for SourceDelayProven, not even an explicit nil
### GetPlatformDeliveryAfterSourceMs

`func (o *AnnouncementChangeEvent) GetPlatformDeliveryAfterSourceMs() float64`

GetPlatformDeliveryAfterSourceMs returns the PlatformDeliveryAfterSourceMs field if non-nil, zero value otherwise.

### GetPlatformDeliveryAfterSourceMsOk

`func (o *AnnouncementChangeEvent) GetPlatformDeliveryAfterSourceMsOk() (*float64, bool)`

GetPlatformDeliveryAfterSourceMsOk returns a tuple with the PlatformDeliveryAfterSourceMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformDeliveryAfterSourceMs

`func (o *AnnouncementChangeEvent) SetPlatformDeliveryAfterSourceMs(v float64)`

SetPlatformDeliveryAfterSourceMs sets PlatformDeliveryAfterSourceMs field to given value.

### HasPlatformDeliveryAfterSourceMs

`func (o *AnnouncementChangeEvent) HasPlatformDeliveryAfterSourceMs() bool`

HasPlatformDeliveryAfterSourceMs returns a boolean if a field has been set.

### SetPlatformDeliveryAfterSourceMsNil

`func (o *AnnouncementChangeEvent) SetPlatformDeliveryAfterSourceMsNil(b bool)`

 SetPlatformDeliveryAfterSourceMsNil sets the value for PlatformDeliveryAfterSourceMs to be an explicit nil

### UnsetPlatformDeliveryAfterSourceMs
`func (o *AnnouncementChangeEvent) UnsetPlatformDeliveryAfterSourceMs()`

UnsetPlatformDeliveryAfterSourceMs ensures that no value is present for PlatformDeliveryAfterSourceMs, not even an explicit nil
### GetPollingStartedAtNs

`func (o *AnnouncementChangeEvent) GetPollingStartedAtNs() float64`

GetPollingStartedAtNs returns the PollingStartedAtNs field if non-nil, zero value otherwise.

### GetPollingStartedAtNsOk

`func (o *AnnouncementChangeEvent) GetPollingStartedAtNsOk() (*float64, bool)`

GetPollingStartedAtNsOk returns a tuple with the PollingStartedAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingStartedAtNs

`func (o *AnnouncementChangeEvent) SetPollingStartedAtNs(v float64)`

SetPollingStartedAtNs sets PollingStartedAtNs field to given value.

### HasPollingStartedAtNs

`func (o *AnnouncementChangeEvent) HasPollingStartedAtNs() bool`

HasPollingStartedAtNs returns a boolean if a field has been set.

### SetPollingStartedAtNsNil

`func (o *AnnouncementChangeEvent) SetPollingStartedAtNsNil(b bool)`

 SetPollingStartedAtNsNil sets the value for PollingStartedAtNs to be an explicit nil

### UnsetPollingStartedAtNs
`func (o *AnnouncementChangeEvent) UnsetPollingStartedAtNs()`

UnsetPollingStartedAtNs ensures that no value is present for PollingStartedAtNs, not even an explicit nil
### GetPollingStartedLagMs

`func (o *AnnouncementChangeEvent) GetPollingStartedLagMs() float64`

GetPollingStartedLagMs returns the PollingStartedLagMs field if non-nil, zero value otherwise.

### GetPollingStartedLagMsOk

`func (o *AnnouncementChangeEvent) GetPollingStartedLagMsOk() (*float64, bool)`

GetPollingStartedLagMsOk returns a tuple with the PollingStartedLagMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingStartedLagMs

`func (o *AnnouncementChangeEvent) SetPollingStartedLagMs(v float64)`

SetPollingStartedLagMs sets PollingStartedLagMs field to given value.

### HasPollingStartedLagMs

`func (o *AnnouncementChangeEvent) HasPollingStartedLagMs() bool`

HasPollingStartedLagMs returns a boolean if a field has been set.

### SetPollingStartedLagMsNil

`func (o *AnnouncementChangeEvent) SetPollingStartedLagMsNil(b bool)`

 SetPollingStartedLagMsNil sets the value for PollingStartedLagMs to be an explicit nil

### UnsetPollingStartedLagMs
`func (o *AnnouncementChangeEvent) UnsetPollingStartedLagMs()`

UnsetPollingStartedLagMs ensures that no value is present for PollingStartedLagMs, not even an explicit nil
### GetPollingStartedBeforeScheduledRelease

`func (o *AnnouncementChangeEvent) GetPollingStartedBeforeScheduledRelease() bool`

GetPollingStartedBeforeScheduledRelease returns the PollingStartedBeforeScheduledRelease field if non-nil, zero value otherwise.

### GetPollingStartedBeforeScheduledReleaseOk

`func (o *AnnouncementChangeEvent) GetPollingStartedBeforeScheduledReleaseOk() (*bool, bool)`

GetPollingStartedBeforeScheduledReleaseOk returns a tuple with the PollingStartedBeforeScheduledRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingStartedBeforeScheduledRelease

`func (o *AnnouncementChangeEvent) SetPollingStartedBeforeScheduledRelease(v bool)`

SetPollingStartedBeforeScheduledRelease sets PollingStartedBeforeScheduledRelease field to given value.

### HasPollingStartedBeforeScheduledRelease

`func (o *AnnouncementChangeEvent) HasPollingStartedBeforeScheduledRelease() bool`

HasPollingStartedBeforeScheduledRelease returns a boolean if a field has been set.

### SetPollingStartedBeforeScheduledReleaseNil

`func (o *AnnouncementChangeEvent) SetPollingStartedBeforeScheduledReleaseNil(b bool)`

 SetPollingStartedBeforeScheduledReleaseNil sets the value for PollingStartedBeforeScheduledRelease to be an explicit nil

### UnsetPollingStartedBeforeScheduledRelease
`func (o *AnnouncementChangeEvent) UnsetPollingStartedBeforeScheduledRelease()`

UnsetPollingStartedBeforeScheduledRelease ensures that no value is present for PollingStartedBeforeScheduledRelease, not even an explicit nil
### GetPreFreshStalePollCount

`func (o *AnnouncementChangeEvent) GetPreFreshStalePollCount() float64`

GetPreFreshStalePollCount returns the PreFreshStalePollCount field if non-nil, zero value otherwise.

### GetPreFreshStalePollCountOk

`func (o *AnnouncementChangeEvent) GetPreFreshStalePollCountOk() (*float64, bool)`

GetPreFreshStalePollCountOk returns a tuple with the PreFreshStalePollCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreFreshStalePollCount

`func (o *AnnouncementChangeEvent) SetPreFreshStalePollCount(v float64)`

SetPreFreshStalePollCount sets PreFreshStalePollCount field to given value.

### HasPreFreshStalePollCount

`func (o *AnnouncementChangeEvent) HasPreFreshStalePollCount() bool`

HasPreFreshStalePollCount returns a boolean if a field has been set.

### SetPreFreshStalePollCountNil

`func (o *AnnouncementChangeEvent) SetPreFreshStalePollCountNil(b bool)`

 SetPreFreshStalePollCountNil sets the value for PreFreshStalePollCount to be an explicit nil

### UnsetPreFreshStalePollCount
`func (o *AnnouncementChangeEvent) UnsetPreFreshStalePollCount()`

UnsetPreFreshStalePollCount ensures that no value is present for PreFreshStalePollCount, not even an explicit nil
### GetLastStaleFetchCompletedAtNs

`func (o *AnnouncementChangeEvent) GetLastStaleFetchCompletedAtNs() float64`

GetLastStaleFetchCompletedAtNs returns the LastStaleFetchCompletedAtNs field if non-nil, zero value otherwise.

### GetLastStaleFetchCompletedAtNsOk

`func (o *AnnouncementChangeEvent) GetLastStaleFetchCompletedAtNsOk() (*float64, bool)`

GetLastStaleFetchCompletedAtNsOk returns a tuple with the LastStaleFetchCompletedAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStaleFetchCompletedAtNs

`func (o *AnnouncementChangeEvent) SetLastStaleFetchCompletedAtNs(v float64)`

SetLastStaleFetchCompletedAtNs sets LastStaleFetchCompletedAtNs field to given value.

### HasLastStaleFetchCompletedAtNs

`func (o *AnnouncementChangeEvent) HasLastStaleFetchCompletedAtNs() bool`

HasLastStaleFetchCompletedAtNs returns a boolean if a field has been set.

### SetLastStaleFetchCompletedAtNsNil

`func (o *AnnouncementChangeEvent) SetLastStaleFetchCompletedAtNsNil(b bool)`

 SetLastStaleFetchCompletedAtNsNil sets the value for LastStaleFetchCompletedAtNs to be an explicit nil

### UnsetLastStaleFetchCompletedAtNs
`func (o *AnnouncementChangeEvent) UnsetLastStaleFetchCompletedAtNs()`

UnsetLastStaleFetchCompletedAtNs ensures that no value is present for LastStaleFetchCompletedAtNs, not even an explicit nil
### GetFirstFreshResponseCompletedAtNs

`func (o *AnnouncementChangeEvent) GetFirstFreshResponseCompletedAtNs() float64`

GetFirstFreshResponseCompletedAtNs returns the FirstFreshResponseCompletedAtNs field if non-nil, zero value otherwise.

### GetFirstFreshResponseCompletedAtNsOk

`func (o *AnnouncementChangeEvent) GetFirstFreshResponseCompletedAtNsOk() (*float64, bool)`

GetFirstFreshResponseCompletedAtNsOk returns a tuple with the FirstFreshResponseCompletedAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstFreshResponseCompletedAtNs

`func (o *AnnouncementChangeEvent) SetFirstFreshResponseCompletedAtNs(v float64)`

SetFirstFreshResponseCompletedAtNs sets FirstFreshResponseCompletedAtNs field to given value.

### HasFirstFreshResponseCompletedAtNs

`func (o *AnnouncementChangeEvent) HasFirstFreshResponseCompletedAtNs() bool`

HasFirstFreshResponseCompletedAtNs returns a boolean if a field has been set.

### SetFirstFreshResponseCompletedAtNsNil

`func (o *AnnouncementChangeEvent) SetFirstFreshResponseCompletedAtNsNil(b bool)`

 SetFirstFreshResponseCompletedAtNsNil sets the value for FirstFreshResponseCompletedAtNs to be an explicit nil

### UnsetFirstFreshResponseCompletedAtNs
`func (o *AnnouncementChangeEvent) UnsetFirstFreshResponseCompletedAtNs()`

UnsetFirstFreshResponseCompletedAtNs ensures that no value is present for FirstFreshResponseCompletedAtNs, not even an explicit nil
### GetOfficialSourceStaleAfterScheduledRelease

`func (o *AnnouncementChangeEvent) GetOfficialSourceStaleAfterScheduledRelease() bool`

GetOfficialSourceStaleAfterScheduledRelease returns the OfficialSourceStaleAfterScheduledRelease field if non-nil, zero value otherwise.

### GetOfficialSourceStaleAfterScheduledReleaseOk

`func (o *AnnouncementChangeEvent) GetOfficialSourceStaleAfterScheduledReleaseOk() (*bool, bool)`

GetOfficialSourceStaleAfterScheduledReleaseOk returns a tuple with the OfficialSourceStaleAfterScheduledRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialSourceStaleAfterScheduledRelease

`func (o *AnnouncementChangeEvent) SetOfficialSourceStaleAfterScheduledRelease(v bool)`

SetOfficialSourceStaleAfterScheduledRelease sets OfficialSourceStaleAfterScheduledRelease field to given value.

### HasOfficialSourceStaleAfterScheduledRelease

`func (o *AnnouncementChangeEvent) HasOfficialSourceStaleAfterScheduledRelease() bool`

HasOfficialSourceStaleAfterScheduledRelease returns a boolean if a field has been set.

### SetOfficialSourceStaleAfterScheduledReleaseNil

`func (o *AnnouncementChangeEvent) SetOfficialSourceStaleAfterScheduledReleaseNil(b bool)`

 SetOfficialSourceStaleAfterScheduledReleaseNil sets the value for OfficialSourceStaleAfterScheduledRelease to be an explicit nil

### UnsetOfficialSourceStaleAfterScheduledRelease
`func (o *AnnouncementChangeEvent) UnsetOfficialSourceStaleAfterScheduledRelease()`

UnsetOfficialSourceStaleAfterScheduledRelease ensures that no value is present for OfficialSourceStaleAfterScheduledRelease, not even an explicit nil
### GetScheduledToFirstFreshResponseMs

`func (o *AnnouncementChangeEvent) GetScheduledToFirstFreshResponseMs() float64`

GetScheduledToFirstFreshResponseMs returns the ScheduledToFirstFreshResponseMs field if non-nil, zero value otherwise.

### GetScheduledToFirstFreshResponseMsOk

`func (o *AnnouncementChangeEvent) GetScheduledToFirstFreshResponseMsOk() (*float64, bool)`

GetScheduledToFirstFreshResponseMsOk returns a tuple with the ScheduledToFirstFreshResponseMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledToFirstFreshResponseMs

`func (o *AnnouncementChangeEvent) SetScheduledToFirstFreshResponseMs(v float64)`

SetScheduledToFirstFreshResponseMs sets ScheduledToFirstFreshResponseMs field to given value.

### HasScheduledToFirstFreshResponseMs

`func (o *AnnouncementChangeEvent) HasScheduledToFirstFreshResponseMs() bool`

HasScheduledToFirstFreshResponseMs returns a boolean if a field has been set.

### SetScheduledToFirstFreshResponseMsNil

`func (o *AnnouncementChangeEvent) SetScheduledToFirstFreshResponseMsNil(b bool)`

 SetScheduledToFirstFreshResponseMsNil sets the value for ScheduledToFirstFreshResponseMs to be an explicit nil

### UnsetScheduledToFirstFreshResponseMs
`func (o *AnnouncementChangeEvent) UnsetScheduledToFirstFreshResponseMs()`

UnsetScheduledToFirstFreshResponseMs ensures that no value is present for ScheduledToFirstFreshResponseMs, not even an explicit nil
### GetSourceFreshnessObservationWindowStartAtNs

`func (o *AnnouncementChangeEvent) GetSourceFreshnessObservationWindowStartAtNs() float64`

GetSourceFreshnessObservationWindowStartAtNs returns the SourceFreshnessObservationWindowStartAtNs field if non-nil, zero value otherwise.

### GetSourceFreshnessObservationWindowStartAtNsOk

`func (o *AnnouncementChangeEvent) GetSourceFreshnessObservationWindowStartAtNsOk() (*float64, bool)`

GetSourceFreshnessObservationWindowStartAtNsOk returns a tuple with the SourceFreshnessObservationWindowStartAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFreshnessObservationWindowStartAtNs

`func (o *AnnouncementChangeEvent) SetSourceFreshnessObservationWindowStartAtNs(v float64)`

SetSourceFreshnessObservationWindowStartAtNs sets SourceFreshnessObservationWindowStartAtNs field to given value.

### HasSourceFreshnessObservationWindowStartAtNs

`func (o *AnnouncementChangeEvent) HasSourceFreshnessObservationWindowStartAtNs() bool`

HasSourceFreshnessObservationWindowStartAtNs returns a boolean if a field has been set.

### SetSourceFreshnessObservationWindowStartAtNsNil

`func (o *AnnouncementChangeEvent) SetSourceFreshnessObservationWindowStartAtNsNil(b bool)`

 SetSourceFreshnessObservationWindowStartAtNsNil sets the value for SourceFreshnessObservationWindowStartAtNs to be an explicit nil

### UnsetSourceFreshnessObservationWindowStartAtNs
`func (o *AnnouncementChangeEvent) UnsetSourceFreshnessObservationWindowStartAtNs()`

UnsetSourceFreshnessObservationWindowStartAtNs ensures that no value is present for SourceFreshnessObservationWindowStartAtNs, not even an explicit nil
### GetSourceFreshnessObservationWindowEndAtNs

`func (o *AnnouncementChangeEvent) GetSourceFreshnessObservationWindowEndAtNs() float64`

GetSourceFreshnessObservationWindowEndAtNs returns the SourceFreshnessObservationWindowEndAtNs field if non-nil, zero value otherwise.

### GetSourceFreshnessObservationWindowEndAtNsOk

`func (o *AnnouncementChangeEvent) GetSourceFreshnessObservationWindowEndAtNsOk() (*float64, bool)`

GetSourceFreshnessObservationWindowEndAtNsOk returns a tuple with the SourceFreshnessObservationWindowEndAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFreshnessObservationWindowEndAtNs

`func (o *AnnouncementChangeEvent) SetSourceFreshnessObservationWindowEndAtNs(v float64)`

SetSourceFreshnessObservationWindowEndAtNs sets SourceFreshnessObservationWindowEndAtNs field to given value.

### HasSourceFreshnessObservationWindowEndAtNs

`func (o *AnnouncementChangeEvent) HasSourceFreshnessObservationWindowEndAtNs() bool`

HasSourceFreshnessObservationWindowEndAtNs returns a boolean if a field has been set.

### SetSourceFreshnessObservationWindowEndAtNsNil

`func (o *AnnouncementChangeEvent) SetSourceFreshnessObservationWindowEndAtNsNil(b bool)`

 SetSourceFreshnessObservationWindowEndAtNsNil sets the value for SourceFreshnessObservationWindowEndAtNs to be an explicit nil

### UnsetSourceFreshnessObservationWindowEndAtNs
`func (o *AnnouncementChangeEvent) UnsetSourceFreshnessObservationWindowEndAtNs()`

UnsetSourceFreshnessObservationWindowEndAtNs ensures that no value is present for SourceFreshnessObservationWindowEndAtNs, not even an explicit nil
### GetSourceFreshnessObservationWindowMs

`func (o *AnnouncementChangeEvent) GetSourceFreshnessObservationWindowMs() float64`

GetSourceFreshnessObservationWindowMs returns the SourceFreshnessObservationWindowMs field if non-nil, zero value otherwise.

### GetSourceFreshnessObservationWindowMsOk

`func (o *AnnouncementChangeEvent) GetSourceFreshnessObservationWindowMsOk() (*float64, bool)`

GetSourceFreshnessObservationWindowMsOk returns a tuple with the SourceFreshnessObservationWindowMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFreshnessObservationWindowMs

`func (o *AnnouncementChangeEvent) SetSourceFreshnessObservationWindowMs(v float64)`

SetSourceFreshnessObservationWindowMs sets SourceFreshnessObservationWindowMs field to given value.

### HasSourceFreshnessObservationWindowMs

`func (o *AnnouncementChangeEvent) HasSourceFreshnessObservationWindowMs() bool`

HasSourceFreshnessObservationWindowMs returns a boolean if a field has been set.

### SetSourceFreshnessObservationWindowMsNil

`func (o *AnnouncementChangeEvent) SetSourceFreshnessObservationWindowMsNil(b bool)`

 SetSourceFreshnessObservationWindowMsNil sets the value for SourceFreshnessObservationWindowMs to be an explicit nil

### UnsetSourceFreshnessObservationWindowMs
`func (o *AnnouncementChangeEvent) UnsetSourceFreshnessObservationWindowMs()`

UnsetSourceFreshnessObservationWindowMs ensures that no value is present for SourceFreshnessObservationWindowMs, not even an explicit nil
### GetSourceFreshnessObservationBasis

`func (o *AnnouncementChangeEvent) GetSourceFreshnessObservationBasis() string`

GetSourceFreshnessObservationBasis returns the SourceFreshnessObservationBasis field if non-nil, zero value otherwise.

### GetSourceFreshnessObservationBasisOk

`func (o *AnnouncementChangeEvent) GetSourceFreshnessObservationBasisOk() (*string, bool)`

GetSourceFreshnessObservationBasisOk returns a tuple with the SourceFreshnessObservationBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFreshnessObservationBasis

`func (o *AnnouncementChangeEvent) SetSourceFreshnessObservationBasis(v string)`

SetSourceFreshnessObservationBasis sets SourceFreshnessObservationBasis field to given value.

### HasSourceFreshnessObservationBasis

`func (o *AnnouncementChangeEvent) HasSourceFreshnessObservationBasis() bool`

HasSourceFreshnessObservationBasis returns a boolean if a field has been set.

### SetSourceFreshnessObservationBasisNil

`func (o *AnnouncementChangeEvent) SetSourceFreshnessObservationBasisNil(b bool)`

 SetSourceFreshnessObservationBasisNil sets the value for SourceFreshnessObservationBasis to be an explicit nil

### UnsetSourceFreshnessObservationBasis
`func (o *AnnouncementChangeEvent) UnsetSourceFreshnessObservationBasis()`

UnsetSourceFreshnessObservationBasis ensures that no value is present for SourceFreshnessObservationBasis, not even an explicit nil
### GetSourceDelayAttribution

`func (o *AnnouncementChangeEvent) GetSourceDelayAttribution() string`

GetSourceDelayAttribution returns the SourceDelayAttribution field if non-nil, zero value otherwise.

### GetSourceDelayAttributionOk

`func (o *AnnouncementChangeEvent) GetSourceDelayAttributionOk() (*string, bool)`

GetSourceDelayAttributionOk returns a tuple with the SourceDelayAttribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDelayAttribution

`func (o *AnnouncementChangeEvent) SetSourceDelayAttribution(v string)`

SetSourceDelayAttribution sets SourceDelayAttribution field to given value.

### HasSourceDelayAttribution

`func (o *AnnouncementChangeEvent) HasSourceDelayAttribution() bool`

HasSourceDelayAttribution returns a boolean if a field has been set.

### SetSourceDelayAttributionNil

`func (o *AnnouncementChangeEvent) SetSourceDelayAttributionNil(b bool)`

 SetSourceDelayAttributionNil sets the value for SourceDelayAttribution to be an explicit nil

### UnsetSourceDelayAttribution
`func (o *AnnouncementChangeEvent) UnsetSourceDelayAttribution()`

UnsetSourceDelayAttribution ensures that no value is present for SourceDelayAttribution, not even an explicit nil
### GetFxmdDeliveryAttribution

`func (o *AnnouncementChangeEvent) GetFxmdDeliveryAttribution() string`

GetFxmdDeliveryAttribution returns the FxmdDeliveryAttribution field if non-nil, zero value otherwise.

### GetFxmdDeliveryAttributionOk

`func (o *AnnouncementChangeEvent) GetFxmdDeliveryAttributionOk() (*string, bool)`

GetFxmdDeliveryAttributionOk returns a tuple with the FxmdDeliveryAttribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxmdDeliveryAttribution

`func (o *AnnouncementChangeEvent) SetFxmdDeliveryAttribution(v string)`

SetFxmdDeliveryAttribution sets FxmdDeliveryAttribution field to given value.

### HasFxmdDeliveryAttribution

`func (o *AnnouncementChangeEvent) HasFxmdDeliveryAttribution() bool`

HasFxmdDeliveryAttribution returns a boolean if a field has been set.

### SetFxmdDeliveryAttributionNil

`func (o *AnnouncementChangeEvent) SetFxmdDeliveryAttributionNil(b bool)`

 SetFxmdDeliveryAttributionNil sets the value for FxmdDeliveryAttribution to be an explicit nil

### UnsetFxmdDeliveryAttribution
`func (o *AnnouncementChangeEvent) UnsetFxmdDeliveryAttribution()`

UnsetFxmdDeliveryAttribution ensures that no value is present for FxmdDeliveryAttribution, not even an explicit nil
### GetFxmdAfterSourceMs

`func (o *AnnouncementChangeEvent) GetFxmdAfterSourceMs() float64`

GetFxmdAfterSourceMs returns the FxmdAfterSourceMs field if non-nil, zero value otherwise.

### GetFxmdAfterSourceMsOk

`func (o *AnnouncementChangeEvent) GetFxmdAfterSourceMsOk() (*float64, bool)`

GetFxmdAfterSourceMsOk returns a tuple with the FxmdAfterSourceMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxmdAfterSourceMs

`func (o *AnnouncementChangeEvent) SetFxmdAfterSourceMs(v float64)`

SetFxmdAfterSourceMs sets FxmdAfterSourceMs field to given value.

### HasFxmdAfterSourceMs

`func (o *AnnouncementChangeEvent) HasFxmdAfterSourceMs() bool`

HasFxmdAfterSourceMs returns a boolean if a field has been set.

### SetFxmdAfterSourceMsNil

`func (o *AnnouncementChangeEvent) SetFxmdAfterSourceMsNil(b bool)`

 SetFxmdAfterSourceMsNil sets the value for FxmdAfterSourceMs to be an explicit nil

### UnsetFxmdAfterSourceMs
`func (o *AnnouncementChangeEvent) UnsetFxmdAfterSourceMs()`

UnsetFxmdAfterSourceMs ensures that no value is present for FxmdAfterSourceMs, not even an explicit nil
### GetFxmdAfterSourceSloMet

`func (o *AnnouncementChangeEvent) GetFxmdAfterSourceSloMet() bool`

GetFxmdAfterSourceSloMet returns the FxmdAfterSourceSloMet field if non-nil, zero value otherwise.

### GetFxmdAfterSourceSloMetOk

`func (o *AnnouncementChangeEvent) GetFxmdAfterSourceSloMetOk() (*bool, bool)`

GetFxmdAfterSourceSloMetOk returns a tuple with the FxmdAfterSourceSloMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxmdAfterSourceSloMet

`func (o *AnnouncementChangeEvent) SetFxmdAfterSourceSloMet(v bool)`

SetFxmdAfterSourceSloMet sets FxmdAfterSourceSloMet field to given value.

### HasFxmdAfterSourceSloMet

`func (o *AnnouncementChangeEvent) HasFxmdAfterSourceSloMet() bool`

HasFxmdAfterSourceSloMet returns a boolean if a field has been set.

### SetFxmdAfterSourceSloMetNil

`func (o *AnnouncementChangeEvent) SetFxmdAfterSourceSloMetNil(b bool)`

 SetFxmdAfterSourceSloMetNil sets the value for FxmdAfterSourceSloMet to be an explicit nil

### UnsetFxmdAfterSourceSloMet
`func (o *AnnouncementChangeEvent) UnsetFxmdAfterSourceSloMet()`

UnsetFxmdAfterSourceSloMet ensures that no value is present for FxmdAfterSourceSloMet, not even an explicit nil
### GetCatchupSource

`func (o *AnnouncementChangeEvent) GetCatchupSource() string`

GetCatchupSource returns the CatchupSource field if non-nil, zero value otherwise.

### GetCatchupSourceOk

`func (o *AnnouncementChangeEvent) GetCatchupSourceOk() (*string, bool)`

GetCatchupSourceOk returns a tuple with the CatchupSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchupSource

`func (o *AnnouncementChangeEvent) SetCatchupSource(v string)`

SetCatchupSource sets CatchupSource field to given value.

### HasCatchupSource

`func (o *AnnouncementChangeEvent) HasCatchupSource() bool`

HasCatchupSource returns a boolean if a field has been set.

### SetCatchupSourceNil

`func (o *AnnouncementChangeEvent) SetCatchupSourceNil(b bool)`

 SetCatchupSourceNil sets the value for CatchupSource to be an explicit nil

### UnsetCatchupSource
`func (o *AnnouncementChangeEvent) UnsetCatchupSource()`

UnsetCatchupSource ensures that no value is present for CatchupSource, not even an explicit nil
### GetCatchupDelayMs

`func (o *AnnouncementChangeEvent) GetCatchupDelayMs() float64`

GetCatchupDelayMs returns the CatchupDelayMs field if non-nil, zero value otherwise.

### GetCatchupDelayMsOk

`func (o *AnnouncementChangeEvent) GetCatchupDelayMsOk() (*float64, bool)`

GetCatchupDelayMsOk returns a tuple with the CatchupDelayMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchupDelayMs

`func (o *AnnouncementChangeEvent) SetCatchupDelayMs(v float64)`

SetCatchupDelayMs sets CatchupDelayMs field to given value.

### HasCatchupDelayMs

`func (o *AnnouncementChangeEvent) HasCatchupDelayMs() bool`

HasCatchupDelayMs returns a boolean if a field has been set.

### SetCatchupDelayMsNil

`func (o *AnnouncementChangeEvent) SetCatchupDelayMsNil(b bool)`

 SetCatchupDelayMsNil sets the value for CatchupDelayMs to be an explicit nil

### UnsetCatchupDelayMs
`func (o *AnnouncementChangeEvent) UnsetCatchupDelayMs()`

UnsetCatchupDelayMs ensures that no value is present for CatchupDelayMs, not even an explicit nil
### GetRecoveryClass

`func (o *AnnouncementChangeEvent) GetRecoveryClass() string`

GetRecoveryClass returns the RecoveryClass field if non-nil, zero value otherwise.

### GetRecoveryClassOk

`func (o *AnnouncementChangeEvent) GetRecoveryClassOk() (*string, bool)`

GetRecoveryClassOk returns a tuple with the RecoveryClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryClass

`func (o *AnnouncementChangeEvent) SetRecoveryClass(v string)`

SetRecoveryClass sets RecoveryClass field to given value.

### HasRecoveryClass

`func (o *AnnouncementChangeEvent) HasRecoveryClass() bool`

HasRecoveryClass returns a boolean if a field has been set.

### SetRecoveryClassNil

`func (o *AnnouncementChangeEvent) SetRecoveryClassNil(b bool)`

 SetRecoveryClassNil sets the value for RecoveryClass to be an explicit nil

### UnsetRecoveryClass
`func (o *AnnouncementChangeEvent) UnsetRecoveryClass()`

UnsetRecoveryClass ensures that no value is present for RecoveryClass, not even an explicit nil
### GetSubsecondStatus

`func (o *AnnouncementChangeEvent) GetSubsecondStatus() string`

GetSubsecondStatus returns the SubsecondStatus field if non-nil, zero value otherwise.

### GetSubsecondStatusOk

`func (o *AnnouncementChangeEvent) GetSubsecondStatusOk() (*string, bool)`

GetSubsecondStatusOk returns a tuple with the SubsecondStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsecondStatus

`func (o *AnnouncementChangeEvent) SetSubsecondStatus(v string)`

SetSubsecondStatus sets SubsecondStatus field to given value.

### HasSubsecondStatus

`func (o *AnnouncementChangeEvent) HasSubsecondStatus() bool`

HasSubsecondStatus returns a boolean if a field has been set.

### SetSubsecondStatusNil

`func (o *AnnouncementChangeEvent) SetSubsecondStatusNil(b bool)`

 SetSubsecondStatusNil sets the value for SubsecondStatus to be an explicit nil

### UnsetSubsecondStatus
`func (o *AnnouncementChangeEvent) UnsetSubsecondStatus()`

UnsetSubsecondStatus ensures that no value is present for SubsecondStatus, not even an explicit nil
### GetSubsecondOperationalStatus

`func (o *AnnouncementChangeEvent) GetSubsecondOperationalStatus() string`

GetSubsecondOperationalStatus returns the SubsecondOperationalStatus field if non-nil, zero value otherwise.

### GetSubsecondOperationalStatusOk

`func (o *AnnouncementChangeEvent) GetSubsecondOperationalStatusOk() (*string, bool)`

GetSubsecondOperationalStatusOk returns a tuple with the SubsecondOperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsecondOperationalStatus

`func (o *AnnouncementChangeEvent) SetSubsecondOperationalStatus(v string)`

SetSubsecondOperationalStatus sets SubsecondOperationalStatus field to given value.

### HasSubsecondOperationalStatus

`func (o *AnnouncementChangeEvent) HasSubsecondOperationalStatus() bool`

HasSubsecondOperationalStatus returns a boolean if a field has been set.

### SetSubsecondOperationalStatusNil

`func (o *AnnouncementChangeEvent) SetSubsecondOperationalStatusNil(b bool)`

 SetSubsecondOperationalStatusNil sets the value for SubsecondOperationalStatus to be an explicit nil

### UnsetSubsecondOperationalStatus
`func (o *AnnouncementChangeEvent) UnsetSubsecondOperationalStatus()`

UnsetSubsecondOperationalStatus ensures that no value is present for SubsecondOperationalStatus, not even an explicit nil
### GetSubsecondGuaranteeActive

`func (o *AnnouncementChangeEvent) GetSubsecondGuaranteeActive() bool`

GetSubsecondGuaranteeActive returns the SubsecondGuaranteeActive field if non-nil, zero value otherwise.

### GetSubsecondGuaranteeActiveOk

`func (o *AnnouncementChangeEvent) GetSubsecondGuaranteeActiveOk() (*bool, bool)`

GetSubsecondGuaranteeActiveOk returns a tuple with the SubsecondGuaranteeActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsecondGuaranteeActive

`func (o *AnnouncementChangeEvent) SetSubsecondGuaranteeActive(v bool)`

SetSubsecondGuaranteeActive sets SubsecondGuaranteeActive field to given value.

### HasSubsecondGuaranteeActive

`func (o *AnnouncementChangeEvent) HasSubsecondGuaranteeActive() bool`

HasSubsecondGuaranteeActive returns a boolean if a field has been set.

### SetSubsecondGuaranteeActiveNil

`func (o *AnnouncementChangeEvent) SetSubsecondGuaranteeActiveNil(b bool)`

 SetSubsecondGuaranteeActiveNil sets the value for SubsecondGuaranteeActive to be an explicit nil

### UnsetSubsecondGuaranteeActive
`func (o *AnnouncementChangeEvent) UnsetSubsecondGuaranteeActive()`

UnsetSubsecondGuaranteeActive ensures that no value is present for SubsecondGuaranteeActive, not even an explicit nil
### GetSubsecondContractOutcome

`func (o *AnnouncementChangeEvent) GetSubsecondContractOutcome() string`

GetSubsecondContractOutcome returns the SubsecondContractOutcome field if non-nil, zero value otherwise.

### GetSubsecondContractOutcomeOk

`func (o *AnnouncementChangeEvent) GetSubsecondContractOutcomeOk() (*string, bool)`

GetSubsecondContractOutcomeOk returns a tuple with the SubsecondContractOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsecondContractOutcome

`func (o *AnnouncementChangeEvent) SetSubsecondContractOutcome(v string)`

SetSubsecondContractOutcome sets SubsecondContractOutcome field to given value.

### HasSubsecondContractOutcome

`func (o *AnnouncementChangeEvent) HasSubsecondContractOutcome() bool`

HasSubsecondContractOutcome returns a boolean if a field has been set.

### SetSubsecondContractOutcomeNil

`func (o *AnnouncementChangeEvent) SetSubsecondContractOutcomeNil(b bool)`

 SetSubsecondContractOutcomeNil sets the value for SubsecondContractOutcome to be an explicit nil

### UnsetSubsecondContractOutcome
`func (o *AnnouncementChangeEvent) UnsetSubsecondContractOutcome()`

UnsetSubsecondContractOutcome ensures that no value is present for SubsecondContractOutcome, not even an explicit nil
### GetSubsecondContractBreached

`func (o *AnnouncementChangeEvent) GetSubsecondContractBreached() bool`

GetSubsecondContractBreached returns the SubsecondContractBreached field if non-nil, zero value otherwise.

### GetSubsecondContractBreachedOk

`func (o *AnnouncementChangeEvent) GetSubsecondContractBreachedOk() (*bool, bool)`

GetSubsecondContractBreachedOk returns a tuple with the SubsecondContractBreached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsecondContractBreached

`func (o *AnnouncementChangeEvent) SetSubsecondContractBreached(v bool)`

SetSubsecondContractBreached sets SubsecondContractBreached field to given value.

### HasSubsecondContractBreached

`func (o *AnnouncementChangeEvent) HasSubsecondContractBreached() bool`

HasSubsecondContractBreached returns a boolean if a field has been set.

### SetSubsecondContractBreachedNil

`func (o *AnnouncementChangeEvent) SetSubsecondContractBreachedNil(b bool)`

 SetSubsecondContractBreachedNil sets the value for SubsecondContractBreached to be an explicit nil

### UnsetSubsecondContractBreached
`func (o *AnnouncementChangeEvent) UnsetSubsecondContractBreached()`

UnsetSubsecondContractBreached ensures that no value is present for SubsecondContractBreached, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


