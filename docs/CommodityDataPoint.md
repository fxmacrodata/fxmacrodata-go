# CommodityDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourcePermissionStatus** | Pointer to **NullableString** |  | [optional] 
**Date** | **string** |  | 
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
**PctChange** | Pointer to **NullableFloat64** |  | [optional] 
**PctChange12m** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewCommodityDataPoint

`func NewCommodityDataPoint(date string, ) *CommodityDataPoint`

NewCommodityDataPoint instantiates a new CommodityDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommodityDataPointWithDefaults

`func NewCommodityDataPointWithDefaults() *CommodityDataPoint`

NewCommodityDataPointWithDefaults instantiates a new CommodityDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourcePermissionStatus

`func (o *CommodityDataPoint) GetSourcePermissionStatus() string`

GetSourcePermissionStatus returns the SourcePermissionStatus field if non-nil, zero value otherwise.

### GetSourcePermissionStatusOk

`func (o *CommodityDataPoint) GetSourcePermissionStatusOk() (*string, bool)`

GetSourcePermissionStatusOk returns a tuple with the SourcePermissionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePermissionStatus

`func (o *CommodityDataPoint) SetSourcePermissionStatus(v string)`

SetSourcePermissionStatus sets SourcePermissionStatus field to given value.

### HasSourcePermissionStatus

`func (o *CommodityDataPoint) HasSourcePermissionStatus() bool`

HasSourcePermissionStatus returns a boolean if a field has been set.

### SetSourcePermissionStatusNil

`func (o *CommodityDataPoint) SetSourcePermissionStatusNil(b bool)`

 SetSourcePermissionStatusNil sets the value for SourcePermissionStatus to be an explicit nil

### UnsetSourcePermissionStatus
`func (o *CommodityDataPoint) UnsetSourcePermissionStatus()`

UnsetSourcePermissionStatus ensures that no value is present for SourcePermissionStatus, not even an explicit nil
### GetDate

`func (o *CommodityDataPoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CommodityDataPoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CommodityDataPoint) SetDate(v string)`

SetDate sets Date field to given value.


### GetVal

`func (o *CommodityDataPoint) GetVal() float64`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *CommodityDataPoint) GetValOk() (*float64, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *CommodityDataPoint) SetVal(v float64)`

SetVal sets Val field to given value.

### HasVal

`func (o *CommodityDataPoint) HasVal() bool`

HasVal returns a boolean if a field has been set.

### SetValNil

`func (o *CommodityDataPoint) SetValNil(b bool)`

 SetValNil sets the value for Val to be an explicit nil

### UnsetVal
`func (o *CommodityDataPoint) UnsetVal()`

UnsetVal ensures that no value is present for Val, not even an explicit nil
### GetAnnouncementDatetime

`func (o *CommodityDataPoint) GetAnnouncementDatetime() int64`

GetAnnouncementDatetime returns the AnnouncementDatetime field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeOk

`func (o *CommodityDataPoint) GetAnnouncementDatetimeOk() (*int64, bool)`

GetAnnouncementDatetimeOk returns a tuple with the AnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetime

`func (o *CommodityDataPoint) SetAnnouncementDatetime(v int64)`

SetAnnouncementDatetime sets AnnouncementDatetime field to given value.

### HasAnnouncementDatetime

`func (o *CommodityDataPoint) HasAnnouncementDatetime() bool`

HasAnnouncementDatetime returns a boolean if a field has been set.

### SetAnnouncementDatetimeNil

`func (o *CommodityDataPoint) SetAnnouncementDatetimeNil(b bool)`

 SetAnnouncementDatetimeNil sets the value for AnnouncementDatetime to be an explicit nil

### UnsetAnnouncementDatetime
`func (o *CommodityDataPoint) UnsetAnnouncementDatetime()`

UnsetAnnouncementDatetime ensures that no value is present for AnnouncementDatetime, not even an explicit nil
### GetSource

`func (o *CommodityDataPoint) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CommodityDataPoint) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CommodityDataPoint) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CommodityDataPoint) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *CommodityDataPoint) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *CommodityDataPoint) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceUrl

`func (o *CommodityDataPoint) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *CommodityDataPoint) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *CommodityDataPoint) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *CommodityDataPoint) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *CommodityDataPoint) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *CommodityDataPoint) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSourceType

`func (o *CommodityDataPoint) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *CommodityDataPoint) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *CommodityDataPoint) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *CommodityDataPoint) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *CommodityDataPoint) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *CommodityDataPoint) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetSourceChartTimestamp

`func (o *CommodityDataPoint) GetSourceChartTimestamp() int64`

GetSourceChartTimestamp returns the SourceChartTimestamp field if non-nil, zero value otherwise.

### GetSourceChartTimestampOk

`func (o *CommodityDataPoint) GetSourceChartTimestampOk() (*int64, bool)`

GetSourceChartTimestampOk returns a tuple with the SourceChartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChartTimestamp

`func (o *CommodityDataPoint) SetSourceChartTimestamp(v int64)`

SetSourceChartTimestamp sets SourceChartTimestamp field to given value.

### HasSourceChartTimestamp

`func (o *CommodityDataPoint) HasSourceChartTimestamp() bool`

HasSourceChartTimestamp returns a boolean if a field has been set.

### SetSourceChartTimestampNil

`func (o *CommodityDataPoint) SetSourceChartTimestampNil(b bool)`

 SetSourceChartTimestampNil sets the value for SourceChartTimestamp to be an explicit nil

### UnsetSourceChartTimestamp
`func (o *CommodityDataPoint) UnsetSourceChartTimestamp()`

UnsetSourceChartTimestamp ensures that no value is present for SourceChartTimestamp, not even an explicit nil
### GetSourceChartTimestampMs

`func (o *CommodityDataPoint) GetSourceChartTimestampMs() int64`

GetSourceChartTimestampMs returns the SourceChartTimestampMs field if non-nil, zero value otherwise.

### GetSourceChartTimestampMsOk

`func (o *CommodityDataPoint) GetSourceChartTimestampMsOk() (*int64, bool)`

GetSourceChartTimestampMsOk returns a tuple with the SourceChartTimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChartTimestampMs

`func (o *CommodityDataPoint) SetSourceChartTimestampMs(v int64)`

SetSourceChartTimestampMs sets SourceChartTimestampMs field to given value.

### HasSourceChartTimestampMs

`func (o *CommodityDataPoint) HasSourceChartTimestampMs() bool`

HasSourceChartTimestampMs returns a boolean if a field has been set.

### SetSourceChartTimestampMsNil

`func (o *CommodityDataPoint) SetSourceChartTimestampMsNil(b bool)`

 SetSourceChartTimestampMsNil sets the value for SourceChartTimestampMs to be an explicit nil

### UnsetSourceChartTimestampMs
`func (o *CommodityDataPoint) UnsetSourceChartTimestampMs()`

UnsetSourceChartTimestampMs ensures that no value is present for SourceChartTimestampMs, not even an explicit nil
### GetSourceChartPeriod

`func (o *CommodityDataPoint) GetSourceChartPeriod() string`

GetSourceChartPeriod returns the SourceChartPeriod field if non-nil, zero value otherwise.

### GetSourceChartPeriodOk

`func (o *CommodityDataPoint) GetSourceChartPeriodOk() (*string, bool)`

GetSourceChartPeriodOk returns a tuple with the SourceChartPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChartPeriod

`func (o *CommodityDataPoint) SetSourceChartPeriod(v string)`

SetSourceChartPeriod sets SourceChartPeriod field to given value.

### HasSourceChartPeriod

`func (o *CommodityDataPoint) HasSourceChartPeriod() bool`

HasSourceChartPeriod returns a boolean if a field has been set.

### SetSourceChartPeriodNil

`func (o *CommodityDataPoint) SetSourceChartPeriodNil(b bool)`

 SetSourceChartPeriodNil sets the value for SourceChartPeriod to be an explicit nil

### UnsetSourceChartPeriod
`func (o *CommodityDataPoint) UnsetSourceChartPeriod()`

UnsetSourceChartPeriod ensures that no value is present for SourceChartPeriod, not even an explicit nil
### GetQuoteTimeStatus

`func (o *CommodityDataPoint) GetQuoteTimeStatus() string`

GetQuoteTimeStatus returns the QuoteTimeStatus field if non-nil, zero value otherwise.

### GetQuoteTimeStatusOk

`func (o *CommodityDataPoint) GetQuoteTimeStatusOk() (*string, bool)`

GetQuoteTimeStatusOk returns a tuple with the QuoteTimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTimeStatus

`func (o *CommodityDataPoint) SetQuoteTimeStatus(v string)`

SetQuoteTimeStatus sets QuoteTimeStatus field to given value.

### HasQuoteTimeStatus

`func (o *CommodityDataPoint) HasQuoteTimeStatus() bool`

HasQuoteTimeStatus returns a boolean if a field has been set.

### SetQuoteTimeStatusNil

`func (o *CommodityDataPoint) SetQuoteTimeStatusNil(b bool)`

 SetQuoteTimeStatusNil sets the value for QuoteTimeStatus to be an explicit nil

### UnsetQuoteTimeStatus
`func (o *CommodityDataPoint) UnsetQuoteTimeStatus()`

UnsetQuoteTimeStatus ensures that no value is present for QuoteTimeStatus, not even an explicit nil
### GetSamplingMethod

`func (o *CommodityDataPoint) GetSamplingMethod() string`

GetSamplingMethod returns the SamplingMethod field if non-nil, zero value otherwise.

### GetSamplingMethodOk

`func (o *CommodityDataPoint) GetSamplingMethodOk() (*string, bool)`

GetSamplingMethodOk returns a tuple with the SamplingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingMethod

`func (o *CommodityDataPoint) SetSamplingMethod(v string)`

SetSamplingMethod sets SamplingMethod field to given value.

### HasSamplingMethod

`func (o *CommodityDataPoint) HasSamplingMethod() bool`

HasSamplingMethod returns a boolean if a field has been set.

### SetSamplingMethodNil

`func (o *CommodityDataPoint) SetSamplingMethodNil(b bool)`

 SetSamplingMethodNil sets the value for SamplingMethod to be an explicit nil

### UnsetSamplingMethod
`func (o *CommodityDataPoint) UnsetSamplingMethod()`

UnsetSamplingMethod ensures that no value is present for SamplingMethod, not even an explicit nil
### GetPublicationTimeStatus

`func (o *CommodityDataPoint) GetPublicationTimeStatus() string`

GetPublicationTimeStatus returns the PublicationTimeStatus field if non-nil, zero value otherwise.

### GetPublicationTimeStatusOk

`func (o *CommodityDataPoint) GetPublicationTimeStatusOk() (*string, bool)`

GetPublicationTimeStatusOk returns a tuple with the PublicationTimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationTimeStatus

`func (o *CommodityDataPoint) SetPublicationTimeStatus(v string)`

SetPublicationTimeStatus sets PublicationTimeStatus field to given value.

### HasPublicationTimeStatus

`func (o *CommodityDataPoint) HasPublicationTimeStatus() bool`

HasPublicationTimeStatus returns a boolean if a field has been set.

### SetPublicationTimeStatusNil

`func (o *CommodityDataPoint) SetPublicationTimeStatusNil(b bool)`

 SetPublicationTimeStatusNil sets the value for PublicationTimeStatus to be an explicit nil

### UnsetPublicationTimeStatus
`func (o *CommodityDataPoint) UnsetPublicationTimeStatus()`

UnsetPublicationTimeStatus ensures that no value is present for PublicationTimeStatus, not even an explicit nil
### GetPointInTimeSafe

`func (o *CommodityDataPoint) GetPointInTimeSafe() bool`

GetPointInTimeSafe returns the PointInTimeSafe field if non-nil, zero value otherwise.

### GetPointInTimeSafeOk

`func (o *CommodityDataPoint) GetPointInTimeSafeOk() (*bool, bool)`

GetPointInTimeSafeOk returns a tuple with the PointInTimeSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeSafe

`func (o *CommodityDataPoint) SetPointInTimeSafe(v bool)`

SetPointInTimeSafe sets PointInTimeSafe field to given value.

### HasPointInTimeSafe

`func (o *CommodityDataPoint) HasPointInTimeSafe() bool`

HasPointInTimeSafe returns a boolean if a field has been set.

### SetPointInTimeSafeNil

`func (o *CommodityDataPoint) SetPointInTimeSafeNil(b bool)`

 SetPointInTimeSafeNil sets the value for PointInTimeSafe to be an explicit nil

### UnsetPointInTimeSafe
`func (o *CommodityDataPoint) UnsetPointInTimeSafe()`

UnsetPointInTimeSafe ensures that no value is present for PointInTimeSafe, not even an explicit nil
### GetProvenanceVersion

`func (o *CommodityDataPoint) GetProvenanceVersion() int64`

GetProvenanceVersion returns the ProvenanceVersion field if non-nil, zero value otherwise.

### GetProvenanceVersionOk

`func (o *CommodityDataPoint) GetProvenanceVersionOk() (*int64, bool)`

GetProvenanceVersionOk returns a tuple with the ProvenanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenanceVersion

`func (o *CommodityDataPoint) SetProvenanceVersion(v int64)`

SetProvenanceVersion sets ProvenanceVersion field to given value.

### HasProvenanceVersion

`func (o *CommodityDataPoint) HasProvenanceVersion() bool`

HasProvenanceVersion returns a boolean if a field has been set.

### SetProvenanceVersionNil

`func (o *CommodityDataPoint) SetProvenanceVersionNil(b bool)`

 SetProvenanceVersionNil sets the value for ProvenanceVersion to be an explicit nil

### UnsetProvenanceVersion
`func (o *CommodityDataPoint) UnsetProvenanceVersion()`

UnsetProvenanceVersion ensures that no value is present for ProvenanceVersion, not even an explicit nil
### GetPctChange

`func (o *CommodityDataPoint) GetPctChange() float64`

GetPctChange returns the PctChange field if non-nil, zero value otherwise.

### GetPctChangeOk

`func (o *CommodityDataPoint) GetPctChangeOk() (*float64, bool)`

GetPctChangeOk returns a tuple with the PctChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctChange

`func (o *CommodityDataPoint) SetPctChange(v float64)`

SetPctChange sets PctChange field to given value.

### HasPctChange

`func (o *CommodityDataPoint) HasPctChange() bool`

HasPctChange returns a boolean if a field has been set.

### SetPctChangeNil

`func (o *CommodityDataPoint) SetPctChangeNil(b bool)`

 SetPctChangeNil sets the value for PctChange to be an explicit nil

### UnsetPctChange
`func (o *CommodityDataPoint) UnsetPctChange()`

UnsetPctChange ensures that no value is present for PctChange, not even an explicit nil
### GetPctChange12m

`func (o *CommodityDataPoint) GetPctChange12m() float64`

GetPctChange12m returns the PctChange12m field if non-nil, zero value otherwise.

### GetPctChange12mOk

`func (o *CommodityDataPoint) GetPctChange12mOk() (*float64, bool)`

GetPctChange12mOk returns a tuple with the PctChange12m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctChange12m

`func (o *CommodityDataPoint) SetPctChange12m(v float64)`

SetPctChange12m sets PctChange12m field to given value.

### HasPctChange12m

`func (o *CommodityDataPoint) HasPctChange12m() bool`

HasPctChange12m returns a boolean if a field has been set.

### SetPctChange12mNil

`func (o *CommodityDataPoint) SetPctChange12mNil(b bool)`

 SetPctChange12mNil sets the value for PctChange12m to be an explicit nil

### UnsetPctChange12m
`func (o *CommodityDataPoint) UnsetPctChange12m()`

UnsetPctChange12m ensures that no value is present for PctChange12m, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


