# ReleaseTiming

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledAt** | Pointer to **NullableString** | Official scheduled release time (ISO 8601, UTC). | [optional] 
**PublishedAfter** | Pointer to **NullableString** | The official source still showed the previous period after this instant (ISO 8601, UTC), or its &#x60;Last-Modified&#x60; header placed the update after it. Null when no such read was made. | [optional] 
**PublishedBy** | Pointer to **NullableString** | The first read of the official source that showed this period completed at this instant (ISO 8601, UTC), so the release was public by then. Null when the window could not be bounded. | [optional] 
**PublicationWindowBasis** | Pointer to **NullableString** | How the window was bounded: &#x60;poll_bracket&#x60; (a read showing the previous period, then a read showing this one), &#x60;last_modified&#x60; (the source&#39;s &#x60;Last-Modified&#x60; header sets the lower bound), or &#x60;first_new_period_read_only&#x60; (only the upper bound is known). | [optional] 
**CollectedAt** | Pointer to **NullableString** | When FXMacroData first captured this value from the official source (ISO 8601, UTC). Same instant as &#x60;collected_at_iso&#x60;. | [optional] 
**AvailableAt** | Pointer to **NullableString** | When this value was committed to the FXMacroData store and readable through the API (ISO 8601, UTC, microsecond precision). Null on rows captured before this time was recorded. | [optional] 
**PickupMsMin** | Pointer to **NullableFloat64** | Milliseconds from &#x60;published_by&#x60; to &#x60;available_at&#x60;: the shortest time the value can have taken to become available after publication. Null when either time is missing or &#x60;available_at&#x60; precedes &#x60;published_by&#x60;. | [optional] 
**PickupMsMax** | Pointer to **NullableFloat64** | Milliseconds from &#x60;published_after&#x60; to &#x60;available_at&#x60;: the value was available through the API within this long of publication. Null when no read showed the previous period still in place, or when &#x60;available_at&#x60; is missing. | [optional] 

## Methods

### NewReleaseTiming

`func NewReleaseTiming() *ReleaseTiming`

NewReleaseTiming instantiates a new ReleaseTiming object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseTimingWithDefaults

`func NewReleaseTimingWithDefaults() *ReleaseTiming`

NewReleaseTimingWithDefaults instantiates a new ReleaseTiming object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledAt

`func (o *ReleaseTiming) GetScheduledAt() string`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *ReleaseTiming) GetScheduledAtOk() (*string, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *ReleaseTiming) SetScheduledAt(v string)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *ReleaseTiming) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### SetScheduledAtNil

`func (o *ReleaseTiming) SetScheduledAtNil(b bool)`

 SetScheduledAtNil sets the value for ScheduledAt to be an explicit nil

### UnsetScheduledAt
`func (o *ReleaseTiming) UnsetScheduledAt()`

UnsetScheduledAt ensures that no value is present for ScheduledAt, not even an explicit nil
### GetPublishedAfter

`func (o *ReleaseTiming) GetPublishedAfter() string`

GetPublishedAfter returns the PublishedAfter field if non-nil, zero value otherwise.

### GetPublishedAfterOk

`func (o *ReleaseTiming) GetPublishedAfterOk() (*string, bool)`

GetPublishedAfterOk returns a tuple with the PublishedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAfter

`func (o *ReleaseTiming) SetPublishedAfter(v string)`

SetPublishedAfter sets PublishedAfter field to given value.

### HasPublishedAfter

`func (o *ReleaseTiming) HasPublishedAfter() bool`

HasPublishedAfter returns a boolean if a field has been set.

### SetPublishedAfterNil

`func (o *ReleaseTiming) SetPublishedAfterNil(b bool)`

 SetPublishedAfterNil sets the value for PublishedAfter to be an explicit nil

### UnsetPublishedAfter
`func (o *ReleaseTiming) UnsetPublishedAfter()`

UnsetPublishedAfter ensures that no value is present for PublishedAfter, not even an explicit nil
### GetPublishedBy

`func (o *ReleaseTiming) GetPublishedBy() string`

GetPublishedBy returns the PublishedBy field if non-nil, zero value otherwise.

### GetPublishedByOk

`func (o *ReleaseTiming) GetPublishedByOk() (*string, bool)`

GetPublishedByOk returns a tuple with the PublishedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedBy

`func (o *ReleaseTiming) SetPublishedBy(v string)`

SetPublishedBy sets PublishedBy field to given value.

### HasPublishedBy

`func (o *ReleaseTiming) HasPublishedBy() bool`

HasPublishedBy returns a boolean if a field has been set.

### SetPublishedByNil

`func (o *ReleaseTiming) SetPublishedByNil(b bool)`

 SetPublishedByNil sets the value for PublishedBy to be an explicit nil

### UnsetPublishedBy
`func (o *ReleaseTiming) UnsetPublishedBy()`

UnsetPublishedBy ensures that no value is present for PublishedBy, not even an explicit nil
### GetPublicationWindowBasis

`func (o *ReleaseTiming) GetPublicationWindowBasis() string`

GetPublicationWindowBasis returns the PublicationWindowBasis field if non-nil, zero value otherwise.

### GetPublicationWindowBasisOk

`func (o *ReleaseTiming) GetPublicationWindowBasisOk() (*string, bool)`

GetPublicationWindowBasisOk returns a tuple with the PublicationWindowBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationWindowBasis

`func (o *ReleaseTiming) SetPublicationWindowBasis(v string)`

SetPublicationWindowBasis sets PublicationWindowBasis field to given value.

### HasPublicationWindowBasis

`func (o *ReleaseTiming) HasPublicationWindowBasis() bool`

HasPublicationWindowBasis returns a boolean if a field has been set.

### SetPublicationWindowBasisNil

`func (o *ReleaseTiming) SetPublicationWindowBasisNil(b bool)`

 SetPublicationWindowBasisNil sets the value for PublicationWindowBasis to be an explicit nil

### UnsetPublicationWindowBasis
`func (o *ReleaseTiming) UnsetPublicationWindowBasis()`

UnsetPublicationWindowBasis ensures that no value is present for PublicationWindowBasis, not even an explicit nil
### GetCollectedAt

`func (o *ReleaseTiming) GetCollectedAt() string`

GetCollectedAt returns the CollectedAt field if non-nil, zero value otherwise.

### GetCollectedAtOk

`func (o *ReleaseTiming) GetCollectedAtOk() (*string, bool)`

GetCollectedAtOk returns a tuple with the CollectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectedAt

`func (o *ReleaseTiming) SetCollectedAt(v string)`

SetCollectedAt sets CollectedAt field to given value.

### HasCollectedAt

`func (o *ReleaseTiming) HasCollectedAt() bool`

HasCollectedAt returns a boolean if a field has been set.

### SetCollectedAtNil

`func (o *ReleaseTiming) SetCollectedAtNil(b bool)`

 SetCollectedAtNil sets the value for CollectedAt to be an explicit nil

### UnsetCollectedAt
`func (o *ReleaseTiming) UnsetCollectedAt()`

UnsetCollectedAt ensures that no value is present for CollectedAt, not even an explicit nil
### GetAvailableAt

`func (o *ReleaseTiming) GetAvailableAt() string`

GetAvailableAt returns the AvailableAt field if non-nil, zero value otherwise.

### GetAvailableAtOk

`func (o *ReleaseTiming) GetAvailableAtOk() (*string, bool)`

GetAvailableAtOk returns a tuple with the AvailableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAt

`func (o *ReleaseTiming) SetAvailableAt(v string)`

SetAvailableAt sets AvailableAt field to given value.

### HasAvailableAt

`func (o *ReleaseTiming) HasAvailableAt() bool`

HasAvailableAt returns a boolean if a field has been set.

### SetAvailableAtNil

`func (o *ReleaseTiming) SetAvailableAtNil(b bool)`

 SetAvailableAtNil sets the value for AvailableAt to be an explicit nil

### UnsetAvailableAt
`func (o *ReleaseTiming) UnsetAvailableAt()`

UnsetAvailableAt ensures that no value is present for AvailableAt, not even an explicit nil
### GetPickupMsMin

`func (o *ReleaseTiming) GetPickupMsMin() float64`

GetPickupMsMin returns the PickupMsMin field if non-nil, zero value otherwise.

### GetPickupMsMinOk

`func (o *ReleaseTiming) GetPickupMsMinOk() (*float64, bool)`

GetPickupMsMinOk returns a tuple with the PickupMsMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupMsMin

`func (o *ReleaseTiming) SetPickupMsMin(v float64)`

SetPickupMsMin sets PickupMsMin field to given value.

### HasPickupMsMin

`func (o *ReleaseTiming) HasPickupMsMin() bool`

HasPickupMsMin returns a boolean if a field has been set.

### SetPickupMsMinNil

`func (o *ReleaseTiming) SetPickupMsMinNil(b bool)`

 SetPickupMsMinNil sets the value for PickupMsMin to be an explicit nil

### UnsetPickupMsMin
`func (o *ReleaseTiming) UnsetPickupMsMin()`

UnsetPickupMsMin ensures that no value is present for PickupMsMin, not even an explicit nil
### GetPickupMsMax

`func (o *ReleaseTiming) GetPickupMsMax() float64`

GetPickupMsMax returns the PickupMsMax field if non-nil, zero value otherwise.

### GetPickupMsMaxOk

`func (o *ReleaseTiming) GetPickupMsMaxOk() (*float64, bool)`

GetPickupMsMaxOk returns a tuple with the PickupMsMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupMsMax

`func (o *ReleaseTiming) SetPickupMsMax(v float64)`

SetPickupMsMax sets PickupMsMax field to given value.

### HasPickupMsMax

`func (o *ReleaseTiming) HasPickupMsMax() bool`

HasPickupMsMax returns a boolean if a field has been set.

### SetPickupMsMaxNil

`func (o *ReleaseTiming) SetPickupMsMaxNil(b bool)`

 SetPickupMsMaxNil sets the value for PickupMsMax to be an explicit nil

### UnsetPickupMsMax
`func (o *ReleaseTiming) UnsetPickupMsMax()`

UnsetPickupMsMax ensures that no value is present for PickupMsMax, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


