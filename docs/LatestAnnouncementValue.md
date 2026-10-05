# LatestAnnouncementValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **NullableString** |  | [optional] 
**Val** | Pointer to **NullableFloat64** |  | [optional] 
**AnnouncementDatetime** | Pointer to **NullableInt64** |  | [optional] 
**PublicationTimeStatus** | Pointer to **NullableString** | What &#x60;announcement_datetime&#x60; is: &#x60;confirmed&#x60;, &#x60;unverified&#x60;, &#x60;assumed_historical&#x60; or &#x60;unknown&#x60;. Same values as on &#x60;/v1/announcements/{currency}/{indicator}&#x60; rows; only &#x60;confirmed&#x60; is evidence of when the value became public. | [optional] 
**PublicationTimePrecision** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceUrl** | Pointer to **NullableString** |  | [optional] 
**SourceUrlScope** | Pointer to **NullableString** |  | [optional] 
**OriginalVal** | Pointer to **NullableFloat64** |  | [optional] 
**OriginalUnit** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLatestAnnouncementValue

`func NewLatestAnnouncementValue() *LatestAnnouncementValue`

NewLatestAnnouncementValue instantiates a new LatestAnnouncementValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLatestAnnouncementValueWithDefaults

`func NewLatestAnnouncementValueWithDefaults() *LatestAnnouncementValue`

NewLatestAnnouncementValueWithDefaults instantiates a new LatestAnnouncementValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *LatestAnnouncementValue) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *LatestAnnouncementValue) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *LatestAnnouncementValue) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *LatestAnnouncementValue) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *LatestAnnouncementValue) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *LatestAnnouncementValue) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetVal

`func (o *LatestAnnouncementValue) GetVal() float64`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *LatestAnnouncementValue) GetValOk() (*float64, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *LatestAnnouncementValue) SetVal(v float64)`

SetVal sets Val field to given value.

### HasVal

`func (o *LatestAnnouncementValue) HasVal() bool`

HasVal returns a boolean if a field has been set.

### SetValNil

`func (o *LatestAnnouncementValue) SetValNil(b bool)`

 SetValNil sets the value for Val to be an explicit nil

### UnsetVal
`func (o *LatestAnnouncementValue) UnsetVal()`

UnsetVal ensures that no value is present for Val, not even an explicit nil
### GetAnnouncementDatetime

`func (o *LatestAnnouncementValue) GetAnnouncementDatetime() int64`

GetAnnouncementDatetime returns the AnnouncementDatetime field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeOk

`func (o *LatestAnnouncementValue) GetAnnouncementDatetimeOk() (*int64, bool)`

GetAnnouncementDatetimeOk returns a tuple with the AnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetime

`func (o *LatestAnnouncementValue) SetAnnouncementDatetime(v int64)`

SetAnnouncementDatetime sets AnnouncementDatetime field to given value.

### HasAnnouncementDatetime

`func (o *LatestAnnouncementValue) HasAnnouncementDatetime() bool`

HasAnnouncementDatetime returns a boolean if a field has been set.

### SetAnnouncementDatetimeNil

`func (o *LatestAnnouncementValue) SetAnnouncementDatetimeNil(b bool)`

 SetAnnouncementDatetimeNil sets the value for AnnouncementDatetime to be an explicit nil

### UnsetAnnouncementDatetime
`func (o *LatestAnnouncementValue) UnsetAnnouncementDatetime()`

UnsetAnnouncementDatetime ensures that no value is present for AnnouncementDatetime, not even an explicit nil
### GetPublicationTimeStatus

`func (o *LatestAnnouncementValue) GetPublicationTimeStatus() string`

GetPublicationTimeStatus returns the PublicationTimeStatus field if non-nil, zero value otherwise.

### GetPublicationTimeStatusOk

`func (o *LatestAnnouncementValue) GetPublicationTimeStatusOk() (*string, bool)`

GetPublicationTimeStatusOk returns a tuple with the PublicationTimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationTimeStatus

`func (o *LatestAnnouncementValue) SetPublicationTimeStatus(v string)`

SetPublicationTimeStatus sets PublicationTimeStatus field to given value.

### HasPublicationTimeStatus

`func (o *LatestAnnouncementValue) HasPublicationTimeStatus() bool`

HasPublicationTimeStatus returns a boolean if a field has been set.

### SetPublicationTimeStatusNil

`func (o *LatestAnnouncementValue) SetPublicationTimeStatusNil(b bool)`

 SetPublicationTimeStatusNil sets the value for PublicationTimeStatus to be an explicit nil

### UnsetPublicationTimeStatus
`func (o *LatestAnnouncementValue) UnsetPublicationTimeStatus()`

UnsetPublicationTimeStatus ensures that no value is present for PublicationTimeStatus, not even an explicit nil
### GetPublicationTimePrecision

`func (o *LatestAnnouncementValue) GetPublicationTimePrecision() string`

GetPublicationTimePrecision returns the PublicationTimePrecision field if non-nil, zero value otherwise.

### GetPublicationTimePrecisionOk

`func (o *LatestAnnouncementValue) GetPublicationTimePrecisionOk() (*string, bool)`

GetPublicationTimePrecisionOk returns a tuple with the PublicationTimePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationTimePrecision

`func (o *LatestAnnouncementValue) SetPublicationTimePrecision(v string)`

SetPublicationTimePrecision sets PublicationTimePrecision field to given value.

### HasPublicationTimePrecision

`func (o *LatestAnnouncementValue) HasPublicationTimePrecision() bool`

HasPublicationTimePrecision returns a boolean if a field has been set.

### SetPublicationTimePrecisionNil

`func (o *LatestAnnouncementValue) SetPublicationTimePrecisionNil(b bool)`

 SetPublicationTimePrecisionNil sets the value for PublicationTimePrecision to be an explicit nil

### UnsetPublicationTimePrecision
`func (o *LatestAnnouncementValue) UnsetPublicationTimePrecision()`

UnsetPublicationTimePrecision ensures that no value is present for PublicationTimePrecision, not even an explicit nil
### GetSource

`func (o *LatestAnnouncementValue) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LatestAnnouncementValue) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LatestAnnouncementValue) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LatestAnnouncementValue) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *LatestAnnouncementValue) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *LatestAnnouncementValue) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceUrl

`func (o *LatestAnnouncementValue) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *LatestAnnouncementValue) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *LatestAnnouncementValue) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *LatestAnnouncementValue) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *LatestAnnouncementValue) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *LatestAnnouncementValue) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSourceUrlScope

`func (o *LatestAnnouncementValue) GetSourceUrlScope() string`

GetSourceUrlScope returns the SourceUrlScope field if non-nil, zero value otherwise.

### GetSourceUrlScopeOk

`func (o *LatestAnnouncementValue) GetSourceUrlScopeOk() (*string, bool)`

GetSourceUrlScopeOk returns a tuple with the SourceUrlScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrlScope

`func (o *LatestAnnouncementValue) SetSourceUrlScope(v string)`

SetSourceUrlScope sets SourceUrlScope field to given value.

### HasSourceUrlScope

`func (o *LatestAnnouncementValue) HasSourceUrlScope() bool`

HasSourceUrlScope returns a boolean if a field has been set.

### SetSourceUrlScopeNil

`func (o *LatestAnnouncementValue) SetSourceUrlScopeNil(b bool)`

 SetSourceUrlScopeNil sets the value for SourceUrlScope to be an explicit nil

### UnsetSourceUrlScope
`func (o *LatestAnnouncementValue) UnsetSourceUrlScope()`

UnsetSourceUrlScope ensures that no value is present for SourceUrlScope, not even an explicit nil
### GetOriginalVal

`func (o *LatestAnnouncementValue) GetOriginalVal() float64`

GetOriginalVal returns the OriginalVal field if non-nil, zero value otherwise.

### GetOriginalValOk

`func (o *LatestAnnouncementValue) GetOriginalValOk() (*float64, bool)`

GetOriginalValOk returns a tuple with the OriginalVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalVal

`func (o *LatestAnnouncementValue) SetOriginalVal(v float64)`

SetOriginalVal sets OriginalVal field to given value.

### HasOriginalVal

`func (o *LatestAnnouncementValue) HasOriginalVal() bool`

HasOriginalVal returns a boolean if a field has been set.

### SetOriginalValNil

`func (o *LatestAnnouncementValue) SetOriginalValNil(b bool)`

 SetOriginalValNil sets the value for OriginalVal to be an explicit nil

### UnsetOriginalVal
`func (o *LatestAnnouncementValue) UnsetOriginalVal()`

UnsetOriginalVal ensures that no value is present for OriginalVal, not even an explicit nil
### GetOriginalUnit

`func (o *LatestAnnouncementValue) GetOriginalUnit() string`

GetOriginalUnit returns the OriginalUnit field if non-nil, zero value otherwise.

### GetOriginalUnitOk

`func (o *LatestAnnouncementValue) GetOriginalUnitOk() (*string, bool)`

GetOriginalUnitOk returns a tuple with the OriginalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalUnit

`func (o *LatestAnnouncementValue) SetOriginalUnit(v string)`

SetOriginalUnit sets OriginalUnit field to given value.

### HasOriginalUnit

`func (o *LatestAnnouncementValue) HasOriginalUnit() bool`

HasOriginalUnit returns a boolean if a field has been set.

### SetOriginalUnitNil

`func (o *LatestAnnouncementValue) SetOriginalUnitNil(b bool)`

 SetOriginalUnitNil sets the value for OriginalUnit to be an explicit nil

### UnsetOriginalUnit
`func (o *LatestAnnouncementValue) UnsetOriginalUnit()`

UnsetOriginalUnit ensures that no value is present for OriginalUnit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


