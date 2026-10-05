# CommodityObservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourcePermissionStatus** | Pointer to **NullableString** |  | [optional] 
**Date** | Pointer to **NullableString** |  | [optional] 
**Val** | Pointer to **NullableFloat64** |  | [optional] 
**AnnouncementDatetime** | Pointer to **NullableInt64** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceUrl** | Pointer to **NullableString** |  | [optional] 
**SourceType** | Pointer to **NullableString** |  | [optional] 
**SourceChartTimestamp** | Pointer to **NullableInt64** |  | [optional] 
**SourceChartTimestampMs** | Pointer to **NullableInt64** |  | [optional] 
**SourceChartPeriod** | Pointer to **NullableString** |  | [optional] 
**QuoteTimeStatus** | Pointer to **NullableString** |  | [optional] 
**SamplingMethod** | Pointer to **NullableString** |  | [optional] 
**PublicationTimeStatus** | Pointer to **NullableString** |  | [optional] 
**PointInTimeSafe** | Pointer to **NullableBool** |  | [optional] 
**ProvenanceVersion** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewCommodityObservation

`func NewCommodityObservation() *CommodityObservation`

NewCommodityObservation instantiates a new CommodityObservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommodityObservationWithDefaults

`func NewCommodityObservationWithDefaults() *CommodityObservation`

NewCommodityObservationWithDefaults instantiates a new CommodityObservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourcePermissionStatus

`func (o *CommodityObservation) GetSourcePermissionStatus() string`

GetSourcePermissionStatus returns the SourcePermissionStatus field if non-nil, zero value otherwise.

### GetSourcePermissionStatusOk

`func (o *CommodityObservation) GetSourcePermissionStatusOk() (*string, bool)`

GetSourcePermissionStatusOk returns a tuple with the SourcePermissionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePermissionStatus

`func (o *CommodityObservation) SetSourcePermissionStatus(v string)`

SetSourcePermissionStatus sets SourcePermissionStatus field to given value.

### HasSourcePermissionStatus

`func (o *CommodityObservation) HasSourcePermissionStatus() bool`

HasSourcePermissionStatus returns a boolean if a field has been set.

### SetSourcePermissionStatusNil

`func (o *CommodityObservation) SetSourcePermissionStatusNil(b bool)`

 SetSourcePermissionStatusNil sets the value for SourcePermissionStatus to be an explicit nil

### UnsetSourcePermissionStatus
`func (o *CommodityObservation) UnsetSourcePermissionStatus()`

UnsetSourcePermissionStatus ensures that no value is present for SourcePermissionStatus, not even an explicit nil
### GetDate

`func (o *CommodityObservation) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CommodityObservation) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CommodityObservation) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CommodityObservation) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *CommodityObservation) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *CommodityObservation) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetVal

`func (o *CommodityObservation) GetVal() float64`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *CommodityObservation) GetValOk() (*float64, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *CommodityObservation) SetVal(v float64)`

SetVal sets Val field to given value.

### HasVal

`func (o *CommodityObservation) HasVal() bool`

HasVal returns a boolean if a field has been set.

### SetValNil

`func (o *CommodityObservation) SetValNil(b bool)`

 SetValNil sets the value for Val to be an explicit nil

### UnsetVal
`func (o *CommodityObservation) UnsetVal()`

UnsetVal ensures that no value is present for Val, not even an explicit nil
### GetAnnouncementDatetime

`func (o *CommodityObservation) GetAnnouncementDatetime() int64`

GetAnnouncementDatetime returns the AnnouncementDatetime field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeOk

`func (o *CommodityObservation) GetAnnouncementDatetimeOk() (*int64, bool)`

GetAnnouncementDatetimeOk returns a tuple with the AnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetime

`func (o *CommodityObservation) SetAnnouncementDatetime(v int64)`

SetAnnouncementDatetime sets AnnouncementDatetime field to given value.

### HasAnnouncementDatetime

`func (o *CommodityObservation) HasAnnouncementDatetime() bool`

HasAnnouncementDatetime returns a boolean if a field has been set.

### SetAnnouncementDatetimeNil

`func (o *CommodityObservation) SetAnnouncementDatetimeNil(b bool)`

 SetAnnouncementDatetimeNil sets the value for AnnouncementDatetime to be an explicit nil

### UnsetAnnouncementDatetime
`func (o *CommodityObservation) UnsetAnnouncementDatetime()`

UnsetAnnouncementDatetime ensures that no value is present for AnnouncementDatetime, not even an explicit nil
### GetSource

`func (o *CommodityObservation) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CommodityObservation) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CommodityObservation) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CommodityObservation) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *CommodityObservation) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *CommodityObservation) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceUrl

`func (o *CommodityObservation) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *CommodityObservation) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *CommodityObservation) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *CommodityObservation) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *CommodityObservation) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *CommodityObservation) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSourceType

`func (o *CommodityObservation) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *CommodityObservation) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *CommodityObservation) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *CommodityObservation) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *CommodityObservation) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *CommodityObservation) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetSourceChartTimestamp

`func (o *CommodityObservation) GetSourceChartTimestamp() int64`

GetSourceChartTimestamp returns the SourceChartTimestamp field if non-nil, zero value otherwise.

### GetSourceChartTimestampOk

`func (o *CommodityObservation) GetSourceChartTimestampOk() (*int64, bool)`

GetSourceChartTimestampOk returns a tuple with the SourceChartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChartTimestamp

`func (o *CommodityObservation) SetSourceChartTimestamp(v int64)`

SetSourceChartTimestamp sets SourceChartTimestamp field to given value.

### HasSourceChartTimestamp

`func (o *CommodityObservation) HasSourceChartTimestamp() bool`

HasSourceChartTimestamp returns a boolean if a field has been set.

### SetSourceChartTimestampNil

`func (o *CommodityObservation) SetSourceChartTimestampNil(b bool)`

 SetSourceChartTimestampNil sets the value for SourceChartTimestamp to be an explicit nil

### UnsetSourceChartTimestamp
`func (o *CommodityObservation) UnsetSourceChartTimestamp()`

UnsetSourceChartTimestamp ensures that no value is present for SourceChartTimestamp, not even an explicit nil
### GetSourceChartTimestampMs

`func (o *CommodityObservation) GetSourceChartTimestampMs() int64`

GetSourceChartTimestampMs returns the SourceChartTimestampMs field if non-nil, zero value otherwise.

### GetSourceChartTimestampMsOk

`func (o *CommodityObservation) GetSourceChartTimestampMsOk() (*int64, bool)`

GetSourceChartTimestampMsOk returns a tuple with the SourceChartTimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChartTimestampMs

`func (o *CommodityObservation) SetSourceChartTimestampMs(v int64)`

SetSourceChartTimestampMs sets SourceChartTimestampMs field to given value.

### HasSourceChartTimestampMs

`func (o *CommodityObservation) HasSourceChartTimestampMs() bool`

HasSourceChartTimestampMs returns a boolean if a field has been set.

### SetSourceChartTimestampMsNil

`func (o *CommodityObservation) SetSourceChartTimestampMsNil(b bool)`

 SetSourceChartTimestampMsNil sets the value for SourceChartTimestampMs to be an explicit nil

### UnsetSourceChartTimestampMs
`func (o *CommodityObservation) UnsetSourceChartTimestampMs()`

UnsetSourceChartTimestampMs ensures that no value is present for SourceChartTimestampMs, not even an explicit nil
### GetSourceChartPeriod

`func (o *CommodityObservation) GetSourceChartPeriod() string`

GetSourceChartPeriod returns the SourceChartPeriod field if non-nil, zero value otherwise.

### GetSourceChartPeriodOk

`func (o *CommodityObservation) GetSourceChartPeriodOk() (*string, bool)`

GetSourceChartPeriodOk returns a tuple with the SourceChartPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChartPeriod

`func (o *CommodityObservation) SetSourceChartPeriod(v string)`

SetSourceChartPeriod sets SourceChartPeriod field to given value.

### HasSourceChartPeriod

`func (o *CommodityObservation) HasSourceChartPeriod() bool`

HasSourceChartPeriod returns a boolean if a field has been set.

### SetSourceChartPeriodNil

`func (o *CommodityObservation) SetSourceChartPeriodNil(b bool)`

 SetSourceChartPeriodNil sets the value for SourceChartPeriod to be an explicit nil

### UnsetSourceChartPeriod
`func (o *CommodityObservation) UnsetSourceChartPeriod()`

UnsetSourceChartPeriod ensures that no value is present for SourceChartPeriod, not even an explicit nil
### GetQuoteTimeStatus

`func (o *CommodityObservation) GetQuoteTimeStatus() string`

GetQuoteTimeStatus returns the QuoteTimeStatus field if non-nil, zero value otherwise.

### GetQuoteTimeStatusOk

`func (o *CommodityObservation) GetQuoteTimeStatusOk() (*string, bool)`

GetQuoteTimeStatusOk returns a tuple with the QuoteTimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTimeStatus

`func (o *CommodityObservation) SetQuoteTimeStatus(v string)`

SetQuoteTimeStatus sets QuoteTimeStatus field to given value.

### HasQuoteTimeStatus

`func (o *CommodityObservation) HasQuoteTimeStatus() bool`

HasQuoteTimeStatus returns a boolean if a field has been set.

### SetQuoteTimeStatusNil

`func (o *CommodityObservation) SetQuoteTimeStatusNil(b bool)`

 SetQuoteTimeStatusNil sets the value for QuoteTimeStatus to be an explicit nil

### UnsetQuoteTimeStatus
`func (o *CommodityObservation) UnsetQuoteTimeStatus()`

UnsetQuoteTimeStatus ensures that no value is present for QuoteTimeStatus, not even an explicit nil
### GetSamplingMethod

`func (o *CommodityObservation) GetSamplingMethod() string`

GetSamplingMethod returns the SamplingMethod field if non-nil, zero value otherwise.

### GetSamplingMethodOk

`func (o *CommodityObservation) GetSamplingMethodOk() (*string, bool)`

GetSamplingMethodOk returns a tuple with the SamplingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingMethod

`func (o *CommodityObservation) SetSamplingMethod(v string)`

SetSamplingMethod sets SamplingMethod field to given value.

### HasSamplingMethod

`func (o *CommodityObservation) HasSamplingMethod() bool`

HasSamplingMethod returns a boolean if a field has been set.

### SetSamplingMethodNil

`func (o *CommodityObservation) SetSamplingMethodNil(b bool)`

 SetSamplingMethodNil sets the value for SamplingMethod to be an explicit nil

### UnsetSamplingMethod
`func (o *CommodityObservation) UnsetSamplingMethod()`

UnsetSamplingMethod ensures that no value is present for SamplingMethod, not even an explicit nil
### GetPublicationTimeStatus

`func (o *CommodityObservation) GetPublicationTimeStatus() string`

GetPublicationTimeStatus returns the PublicationTimeStatus field if non-nil, zero value otherwise.

### GetPublicationTimeStatusOk

`func (o *CommodityObservation) GetPublicationTimeStatusOk() (*string, bool)`

GetPublicationTimeStatusOk returns a tuple with the PublicationTimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationTimeStatus

`func (o *CommodityObservation) SetPublicationTimeStatus(v string)`

SetPublicationTimeStatus sets PublicationTimeStatus field to given value.

### HasPublicationTimeStatus

`func (o *CommodityObservation) HasPublicationTimeStatus() bool`

HasPublicationTimeStatus returns a boolean if a field has been set.

### SetPublicationTimeStatusNil

`func (o *CommodityObservation) SetPublicationTimeStatusNil(b bool)`

 SetPublicationTimeStatusNil sets the value for PublicationTimeStatus to be an explicit nil

### UnsetPublicationTimeStatus
`func (o *CommodityObservation) UnsetPublicationTimeStatus()`

UnsetPublicationTimeStatus ensures that no value is present for PublicationTimeStatus, not even an explicit nil
### GetPointInTimeSafe

`func (o *CommodityObservation) GetPointInTimeSafe() bool`

GetPointInTimeSafe returns the PointInTimeSafe field if non-nil, zero value otherwise.

### GetPointInTimeSafeOk

`func (o *CommodityObservation) GetPointInTimeSafeOk() (*bool, bool)`

GetPointInTimeSafeOk returns a tuple with the PointInTimeSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeSafe

`func (o *CommodityObservation) SetPointInTimeSafe(v bool)`

SetPointInTimeSafe sets PointInTimeSafe field to given value.

### HasPointInTimeSafe

`func (o *CommodityObservation) HasPointInTimeSafe() bool`

HasPointInTimeSafe returns a boolean if a field has been set.

### SetPointInTimeSafeNil

`func (o *CommodityObservation) SetPointInTimeSafeNil(b bool)`

 SetPointInTimeSafeNil sets the value for PointInTimeSafe to be an explicit nil

### UnsetPointInTimeSafe
`func (o *CommodityObservation) UnsetPointInTimeSafe()`

UnsetPointInTimeSafe ensures that no value is present for PointInTimeSafe, not even an explicit nil
### GetProvenanceVersion

`func (o *CommodityObservation) GetProvenanceVersion() int64`

GetProvenanceVersion returns the ProvenanceVersion field if non-nil, zero value otherwise.

### GetProvenanceVersionOk

`func (o *CommodityObservation) GetProvenanceVersionOk() (*int64, bool)`

GetProvenanceVersionOk returns a tuple with the ProvenanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenanceVersion

`func (o *CommodityObservation) SetProvenanceVersion(v int64)`

SetProvenanceVersion sets ProvenanceVersion field to given value.

### HasProvenanceVersion

`func (o *CommodityObservation) HasProvenanceVersion() bool`

HasProvenanceVersion returns a boolean if a field has been set.

### SetProvenanceVersionNil

`func (o *CommodityObservation) SetProvenanceVersionNil(b bool)`

 SetProvenanceVersionNil sets the value for ProvenanceVersion to be an explicit nil

### UnsetProvenanceVersion
`func (o *CommodityObservation) UnsetProvenanceVersion()`

UnsetProvenanceVersion ensures that no value is present for ProvenanceVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


