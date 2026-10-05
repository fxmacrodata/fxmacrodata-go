# AnnouncementChangesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AnnouncementChangeEvent**](AnnouncementChangeEvent.md) |  | 
**Count** | **int64** |  | 
**NextCursor** | **string** |  | 
**HasMore** | **bool** |  | 
**RetentionSeconds** | **int64** |  | 
**Scope** | [**AnnouncementChangesScope**](AnnouncementChangesScope.md) |  | 
**FreemiumDelay** | Pointer to [**NullableFreemiumDelayNotice**](FreemiumDelayNotice.md) |  | [optional] 

## Methods

### NewAnnouncementChangesResponse

`func NewAnnouncementChangesResponse(data []AnnouncementChangeEvent, count int64, nextCursor string, hasMore bool, retentionSeconds int64, scope AnnouncementChangesScope, ) *AnnouncementChangesResponse`

NewAnnouncementChangesResponse instantiates a new AnnouncementChangesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementChangesResponseWithDefaults

`func NewAnnouncementChangesResponseWithDefaults() *AnnouncementChangesResponse`

NewAnnouncementChangesResponseWithDefaults instantiates a new AnnouncementChangesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AnnouncementChangesResponse) GetData() []AnnouncementChangeEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnnouncementChangesResponse) GetDataOk() (*[]AnnouncementChangeEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnnouncementChangesResponse) SetData(v []AnnouncementChangeEvent)`

SetData sets Data field to given value.


### GetCount

`func (o *AnnouncementChangesResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AnnouncementChangesResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AnnouncementChangesResponse) SetCount(v int64)`

SetCount sets Count field to given value.


### GetNextCursor

`func (o *AnnouncementChangesResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *AnnouncementChangesResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *AnnouncementChangesResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### GetHasMore

`func (o *AnnouncementChangesResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *AnnouncementChangesResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *AnnouncementChangesResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetRetentionSeconds

`func (o *AnnouncementChangesResponse) GetRetentionSeconds() int64`

GetRetentionSeconds returns the RetentionSeconds field if non-nil, zero value otherwise.

### GetRetentionSecondsOk

`func (o *AnnouncementChangesResponse) GetRetentionSecondsOk() (*int64, bool)`

GetRetentionSecondsOk returns a tuple with the RetentionSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionSeconds

`func (o *AnnouncementChangesResponse) SetRetentionSeconds(v int64)`

SetRetentionSeconds sets RetentionSeconds field to given value.


### GetScope

`func (o *AnnouncementChangesResponse) GetScope() AnnouncementChangesScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AnnouncementChangesResponse) GetScopeOk() (*AnnouncementChangesScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AnnouncementChangesResponse) SetScope(v AnnouncementChangesScope)`

SetScope sets Scope field to given value.


### GetFreemiumDelay

`func (o *AnnouncementChangesResponse) GetFreemiumDelay() FreemiumDelayNotice`

GetFreemiumDelay returns the FreemiumDelay field if non-nil, zero value otherwise.

### GetFreemiumDelayOk

`func (o *AnnouncementChangesResponse) GetFreemiumDelayOk() (*FreemiumDelayNotice, bool)`

GetFreemiumDelayOk returns a tuple with the FreemiumDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreemiumDelay

`func (o *AnnouncementChangesResponse) SetFreemiumDelay(v FreemiumDelayNotice)`

SetFreemiumDelay sets FreemiumDelay field to given value.

### HasFreemiumDelay

`func (o *AnnouncementChangesResponse) HasFreemiumDelay() bool`

HasFreemiumDelay returns a boolean if a field has been set.

### SetFreemiumDelayNil

`func (o *AnnouncementChangesResponse) SetFreemiumDelayNil(b bool)`

 SetFreemiumDelayNil sets the value for FreemiumDelay to be an explicit nil

### UnsetFreemiumDelay
`func (o *AnnouncementChangesResponse) UnsetFreemiumDelay()`

UnsetFreemiumDelay ensures that no value is present for FreemiumDelay, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


