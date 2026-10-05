# PointInTimeCompleteness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowCount** | **int64** |  | 
**AnnouncementDatetimeCount** | **int64** |  | 
**MissingAnnouncementDatetimeCount** | **int64** |  | 
**AssumedReleaseTimeCount** | Pointer to **int64** |  | [optional] [default to 0]
**BoundedReleaseTimeCount** | Pointer to **int64** |  | [optional] [default to 0]
**UnboundedReleaseTimeCount** | Pointer to **int64** |  | [optional] [default to 0]
**UnverifiedVintageCount** | Pointer to **int64** |  | [optional] [default to 0]
**UnknownPublicationTimeCount** | Pointer to **int64** |  | [optional] [default to 0]
**PointInTimeSafe** | **bool** |  | 

## Methods

### NewPointInTimeCompleteness

`func NewPointInTimeCompleteness(rowCount int64, announcementDatetimeCount int64, missingAnnouncementDatetimeCount int64, pointInTimeSafe bool, ) *PointInTimeCompleteness`

NewPointInTimeCompleteness instantiates a new PointInTimeCompleteness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointInTimeCompletenessWithDefaults

`func NewPointInTimeCompletenessWithDefaults() *PointInTimeCompleteness`

NewPointInTimeCompletenessWithDefaults instantiates a new PointInTimeCompleteness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowCount

`func (o *PointInTimeCompleteness) GetRowCount() int64`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *PointInTimeCompleteness) GetRowCountOk() (*int64, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *PointInTimeCompleteness) SetRowCount(v int64)`

SetRowCount sets RowCount field to given value.


### GetAnnouncementDatetimeCount

`func (o *PointInTimeCompleteness) GetAnnouncementDatetimeCount() int64`

GetAnnouncementDatetimeCount returns the AnnouncementDatetimeCount field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeCountOk

`func (o *PointInTimeCompleteness) GetAnnouncementDatetimeCountOk() (*int64, bool)`

GetAnnouncementDatetimeCountOk returns a tuple with the AnnouncementDatetimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetimeCount

`func (o *PointInTimeCompleteness) SetAnnouncementDatetimeCount(v int64)`

SetAnnouncementDatetimeCount sets AnnouncementDatetimeCount field to given value.


### GetMissingAnnouncementDatetimeCount

`func (o *PointInTimeCompleteness) GetMissingAnnouncementDatetimeCount() int64`

GetMissingAnnouncementDatetimeCount returns the MissingAnnouncementDatetimeCount field if non-nil, zero value otherwise.

### GetMissingAnnouncementDatetimeCountOk

`func (o *PointInTimeCompleteness) GetMissingAnnouncementDatetimeCountOk() (*int64, bool)`

GetMissingAnnouncementDatetimeCountOk returns a tuple with the MissingAnnouncementDatetimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingAnnouncementDatetimeCount

`func (o *PointInTimeCompleteness) SetMissingAnnouncementDatetimeCount(v int64)`

SetMissingAnnouncementDatetimeCount sets MissingAnnouncementDatetimeCount field to given value.


### GetAssumedReleaseTimeCount

`func (o *PointInTimeCompleteness) GetAssumedReleaseTimeCount() int64`

GetAssumedReleaseTimeCount returns the AssumedReleaseTimeCount field if non-nil, zero value otherwise.

### GetAssumedReleaseTimeCountOk

`func (o *PointInTimeCompleteness) GetAssumedReleaseTimeCountOk() (*int64, bool)`

GetAssumedReleaseTimeCountOk returns a tuple with the AssumedReleaseTimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumedReleaseTimeCount

`func (o *PointInTimeCompleteness) SetAssumedReleaseTimeCount(v int64)`

SetAssumedReleaseTimeCount sets AssumedReleaseTimeCount field to given value.

### HasAssumedReleaseTimeCount

`func (o *PointInTimeCompleteness) HasAssumedReleaseTimeCount() bool`

HasAssumedReleaseTimeCount returns a boolean if a field has been set.

### GetBoundedReleaseTimeCount

`func (o *PointInTimeCompleteness) GetBoundedReleaseTimeCount() int64`

GetBoundedReleaseTimeCount returns the BoundedReleaseTimeCount field if non-nil, zero value otherwise.

### GetBoundedReleaseTimeCountOk

`func (o *PointInTimeCompleteness) GetBoundedReleaseTimeCountOk() (*int64, bool)`

GetBoundedReleaseTimeCountOk returns a tuple with the BoundedReleaseTimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundedReleaseTimeCount

`func (o *PointInTimeCompleteness) SetBoundedReleaseTimeCount(v int64)`

SetBoundedReleaseTimeCount sets BoundedReleaseTimeCount field to given value.

### HasBoundedReleaseTimeCount

`func (o *PointInTimeCompleteness) HasBoundedReleaseTimeCount() bool`

HasBoundedReleaseTimeCount returns a boolean if a field has been set.

### GetUnboundedReleaseTimeCount

`func (o *PointInTimeCompleteness) GetUnboundedReleaseTimeCount() int64`

GetUnboundedReleaseTimeCount returns the UnboundedReleaseTimeCount field if non-nil, zero value otherwise.

### GetUnboundedReleaseTimeCountOk

`func (o *PointInTimeCompleteness) GetUnboundedReleaseTimeCountOk() (*int64, bool)`

GetUnboundedReleaseTimeCountOk returns a tuple with the UnboundedReleaseTimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnboundedReleaseTimeCount

`func (o *PointInTimeCompleteness) SetUnboundedReleaseTimeCount(v int64)`

SetUnboundedReleaseTimeCount sets UnboundedReleaseTimeCount field to given value.

### HasUnboundedReleaseTimeCount

`func (o *PointInTimeCompleteness) HasUnboundedReleaseTimeCount() bool`

HasUnboundedReleaseTimeCount returns a boolean if a field has been set.

### GetUnverifiedVintageCount

`func (o *PointInTimeCompleteness) GetUnverifiedVintageCount() int64`

GetUnverifiedVintageCount returns the UnverifiedVintageCount field if non-nil, zero value otherwise.

### GetUnverifiedVintageCountOk

`func (o *PointInTimeCompleteness) GetUnverifiedVintageCountOk() (*int64, bool)`

GetUnverifiedVintageCountOk returns a tuple with the UnverifiedVintageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnverifiedVintageCount

`func (o *PointInTimeCompleteness) SetUnverifiedVintageCount(v int64)`

SetUnverifiedVintageCount sets UnverifiedVintageCount field to given value.

### HasUnverifiedVintageCount

`func (o *PointInTimeCompleteness) HasUnverifiedVintageCount() bool`

HasUnverifiedVintageCount returns a boolean if a field has been set.

### GetUnknownPublicationTimeCount

`func (o *PointInTimeCompleteness) GetUnknownPublicationTimeCount() int64`

GetUnknownPublicationTimeCount returns the UnknownPublicationTimeCount field if non-nil, zero value otherwise.

### GetUnknownPublicationTimeCountOk

`func (o *PointInTimeCompleteness) GetUnknownPublicationTimeCountOk() (*int64, bool)`

GetUnknownPublicationTimeCountOk returns a tuple with the UnknownPublicationTimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownPublicationTimeCount

`func (o *PointInTimeCompleteness) SetUnknownPublicationTimeCount(v int64)`

SetUnknownPublicationTimeCount sets UnknownPublicationTimeCount field to given value.

### HasUnknownPublicationTimeCount

`func (o *PointInTimeCompleteness) HasUnknownPublicationTimeCount() bool`

HasUnknownPublicationTimeCount returns a boolean if a field has been set.

### GetPointInTimeSafe

`func (o *PointInTimeCompleteness) GetPointInTimeSafe() bool`

GetPointInTimeSafe returns the PointInTimeSafe field if non-nil, zero value otherwise.

### GetPointInTimeSafeOk

`func (o *PointInTimeCompleteness) GetPointInTimeSafeOk() (*bool, bool)`

GetPointInTimeSafeOk returns a tuple with the PointInTimeSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeSafe

`func (o *PointInTimeCompleteness) SetPointInTimeSafe(v bool)`

SetPointInTimeSafe sets PointInTimeSafe field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


