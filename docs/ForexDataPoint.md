# ForexDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourcePermissionStatus** | Pointer to **NullableString** |  | [optional] 
**Date** | **string** |  | 
**Val** | Pointer to **NullableFloat64** |  | [optional] 
**Open** | Pointer to **NullableFloat64** |  | [optional] 
**High** | Pointer to **NullableFloat64** |  | [optional] 
**Low** | Pointer to **NullableFloat64** |  | [optional] 
**Close** | Pointer to **NullableFloat64** |  | [optional] 
**OhlcPointCount** | Pointer to **NullableInt64** |  | [optional] 
**OhlcSourceCount** | Pointer to **NullableInt64** |  | [optional] 
**OhlcTimestampStartUtc** | Pointer to **NullableString** |  | [optional] 
**OhlcTimestampEndUtc** | Pointer to **NullableString** |  | [optional] 
**OhlcType** | Pointer to **NullableString** |  | [optional] 
**AnnouncementDatetime** | Pointer to **NullableInt64** |  | [optional] 
**ObservationDatetime** | Pointer to **NullableInt64** |  | [optional] 
**ObservationDatetimeIso** | Pointer to **NullableString** |  | [optional] 
**ObservationDatetimePrecision** | Pointer to **NullableString** |  | [optional] 
**SourceType** | Pointer to **NullableString** |  | [optional] 
**SourceChartTimestamp** | Pointer to **NullableInt64** |  | [optional] 
**SourceChartTimestampMs** | Pointer to **NullableInt64** |  | [optional] 
**SourceChartPeriod** | Pointer to **NullableString** |  | [optional] 
**QuoteTimeStatus** | Pointer to **NullableString** |  | [optional] 
**SamplingMethod** | Pointer to **NullableString** |  | [optional] 
**PublicationTimeStatus** | Pointer to **NullableString** |  | [optional] 
**PointInTimeSafe** | Pointer to **NullableBool** |  | [optional] 
**ProvenanceVersion** | Pointer to **NullableInt64** |  | [optional] 
**Source** | Pointer to [**ForexPointSource**](ForexPointSource.md) |  | [optional] 
**Sma20** | Pointer to **NullableFloat64** |  | [optional] 
**Sma50** | Pointer to **NullableFloat64** |  | [optional] 
**Sma200** | Pointer to **NullableFloat64** |  | [optional] 
**Ema12** | Pointer to **NullableFloat64** |  | [optional] 
**Ema20** | Pointer to **NullableFloat64** |  | [optional] 
**Ema26** | Pointer to **NullableFloat64** |  | [optional] 
**Ema50** | Pointer to **NullableFloat64** |  | [optional] 
**Ema200** | Pointer to **NullableFloat64** |  | [optional] 
**Rsi14** | Pointer to **NullableFloat64** |  | [optional] 
**Atr14** | Pointer to **NullableFloat64** |  | [optional] 
**Adx14** | Pointer to **NullableFloat64** |  | [optional] 
**StochK14** | Pointer to **NullableFloat64** |  | [optional] 
**StochD3** | Pointer to **NullableFloat64** |  | [optional] 
**WilliamsR14** | Pointer to **NullableFloat64** |  | [optional] 
**Cci20** | Pointer to **NullableFloat64** |  | [optional] 
**DonchianUpper20** | Pointer to **NullableFloat64** |  | [optional] 
**DonchianMiddle20** | Pointer to **NullableFloat64** |  | [optional] 
**DonchianLower20** | Pointer to **NullableFloat64** |  | [optional] 
**Macd** | Pointer to **NullableFloat64** |  | [optional] 
**MacdSignal** | Pointer to **NullableFloat64** |  | [optional] 
**MacdHistogram** | Pointer to **NullableFloat64** |  | [optional] 
**BbUpper** | Pointer to **NullableFloat64** |  | [optional] 
**BbMiddle** | Pointer to **NullableFloat64** |  | [optional] 
**BbLower** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewForexDataPoint

`func NewForexDataPoint(date string, ) *ForexDataPoint`

NewForexDataPoint instantiates a new ForexDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForexDataPointWithDefaults

`func NewForexDataPointWithDefaults() *ForexDataPoint`

NewForexDataPointWithDefaults instantiates a new ForexDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourcePermissionStatus

`func (o *ForexDataPoint) GetSourcePermissionStatus() string`

GetSourcePermissionStatus returns the SourcePermissionStatus field if non-nil, zero value otherwise.

### GetSourcePermissionStatusOk

`func (o *ForexDataPoint) GetSourcePermissionStatusOk() (*string, bool)`

GetSourcePermissionStatusOk returns a tuple with the SourcePermissionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePermissionStatus

`func (o *ForexDataPoint) SetSourcePermissionStatus(v string)`

SetSourcePermissionStatus sets SourcePermissionStatus field to given value.

### HasSourcePermissionStatus

`func (o *ForexDataPoint) HasSourcePermissionStatus() bool`

HasSourcePermissionStatus returns a boolean if a field has been set.

### SetSourcePermissionStatusNil

`func (o *ForexDataPoint) SetSourcePermissionStatusNil(b bool)`

 SetSourcePermissionStatusNil sets the value for SourcePermissionStatus to be an explicit nil

### UnsetSourcePermissionStatus
`func (o *ForexDataPoint) UnsetSourcePermissionStatus()`

UnsetSourcePermissionStatus ensures that no value is present for SourcePermissionStatus, not even an explicit nil
### GetDate

`func (o *ForexDataPoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ForexDataPoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ForexDataPoint) SetDate(v string)`

SetDate sets Date field to given value.


### GetVal

`func (o *ForexDataPoint) GetVal() float64`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *ForexDataPoint) GetValOk() (*float64, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *ForexDataPoint) SetVal(v float64)`

SetVal sets Val field to given value.

### HasVal

`func (o *ForexDataPoint) HasVal() bool`

HasVal returns a boolean if a field has been set.

### SetValNil

`func (o *ForexDataPoint) SetValNil(b bool)`

 SetValNil sets the value for Val to be an explicit nil

### UnsetVal
`func (o *ForexDataPoint) UnsetVal()`

UnsetVal ensures that no value is present for Val, not even an explicit nil
### GetOpen

`func (o *ForexDataPoint) GetOpen() float64`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *ForexDataPoint) GetOpenOk() (*float64, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *ForexDataPoint) SetOpen(v float64)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *ForexDataPoint) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### SetOpenNil

`func (o *ForexDataPoint) SetOpenNil(b bool)`

 SetOpenNil sets the value for Open to be an explicit nil

### UnsetOpen
`func (o *ForexDataPoint) UnsetOpen()`

UnsetOpen ensures that no value is present for Open, not even an explicit nil
### GetHigh

`func (o *ForexDataPoint) GetHigh() float64`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *ForexDataPoint) GetHighOk() (*float64, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *ForexDataPoint) SetHigh(v float64)`

SetHigh sets High field to given value.

### HasHigh

`func (o *ForexDataPoint) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### SetHighNil

`func (o *ForexDataPoint) SetHighNil(b bool)`

 SetHighNil sets the value for High to be an explicit nil

### UnsetHigh
`func (o *ForexDataPoint) UnsetHigh()`

UnsetHigh ensures that no value is present for High, not even an explicit nil
### GetLow

`func (o *ForexDataPoint) GetLow() float64`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *ForexDataPoint) GetLowOk() (*float64, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *ForexDataPoint) SetLow(v float64)`

SetLow sets Low field to given value.

### HasLow

`func (o *ForexDataPoint) HasLow() bool`

HasLow returns a boolean if a field has been set.

### SetLowNil

`func (o *ForexDataPoint) SetLowNil(b bool)`

 SetLowNil sets the value for Low to be an explicit nil

### UnsetLow
`func (o *ForexDataPoint) UnsetLow()`

UnsetLow ensures that no value is present for Low, not even an explicit nil
### GetClose

`func (o *ForexDataPoint) GetClose() float64`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *ForexDataPoint) GetCloseOk() (*float64, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *ForexDataPoint) SetClose(v float64)`

SetClose sets Close field to given value.

### HasClose

`func (o *ForexDataPoint) HasClose() bool`

HasClose returns a boolean if a field has been set.

### SetCloseNil

`func (o *ForexDataPoint) SetCloseNil(b bool)`

 SetCloseNil sets the value for Close to be an explicit nil

### UnsetClose
`func (o *ForexDataPoint) UnsetClose()`

UnsetClose ensures that no value is present for Close, not even an explicit nil
### GetOhlcPointCount

`func (o *ForexDataPoint) GetOhlcPointCount() int64`

GetOhlcPointCount returns the OhlcPointCount field if non-nil, zero value otherwise.

### GetOhlcPointCountOk

`func (o *ForexDataPoint) GetOhlcPointCountOk() (*int64, bool)`

GetOhlcPointCountOk returns a tuple with the OhlcPointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOhlcPointCount

`func (o *ForexDataPoint) SetOhlcPointCount(v int64)`

SetOhlcPointCount sets OhlcPointCount field to given value.

### HasOhlcPointCount

`func (o *ForexDataPoint) HasOhlcPointCount() bool`

HasOhlcPointCount returns a boolean if a field has been set.

### SetOhlcPointCountNil

`func (o *ForexDataPoint) SetOhlcPointCountNil(b bool)`

 SetOhlcPointCountNil sets the value for OhlcPointCount to be an explicit nil

### UnsetOhlcPointCount
`func (o *ForexDataPoint) UnsetOhlcPointCount()`

UnsetOhlcPointCount ensures that no value is present for OhlcPointCount, not even an explicit nil
### GetOhlcSourceCount

`func (o *ForexDataPoint) GetOhlcSourceCount() int64`

GetOhlcSourceCount returns the OhlcSourceCount field if non-nil, zero value otherwise.

### GetOhlcSourceCountOk

`func (o *ForexDataPoint) GetOhlcSourceCountOk() (*int64, bool)`

GetOhlcSourceCountOk returns a tuple with the OhlcSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOhlcSourceCount

`func (o *ForexDataPoint) SetOhlcSourceCount(v int64)`

SetOhlcSourceCount sets OhlcSourceCount field to given value.

### HasOhlcSourceCount

`func (o *ForexDataPoint) HasOhlcSourceCount() bool`

HasOhlcSourceCount returns a boolean if a field has been set.

### SetOhlcSourceCountNil

`func (o *ForexDataPoint) SetOhlcSourceCountNil(b bool)`

 SetOhlcSourceCountNil sets the value for OhlcSourceCount to be an explicit nil

### UnsetOhlcSourceCount
`func (o *ForexDataPoint) UnsetOhlcSourceCount()`

UnsetOhlcSourceCount ensures that no value is present for OhlcSourceCount, not even an explicit nil
### GetOhlcTimestampStartUtc

`func (o *ForexDataPoint) GetOhlcTimestampStartUtc() string`

GetOhlcTimestampStartUtc returns the OhlcTimestampStartUtc field if non-nil, zero value otherwise.

### GetOhlcTimestampStartUtcOk

`func (o *ForexDataPoint) GetOhlcTimestampStartUtcOk() (*string, bool)`

GetOhlcTimestampStartUtcOk returns a tuple with the OhlcTimestampStartUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOhlcTimestampStartUtc

`func (o *ForexDataPoint) SetOhlcTimestampStartUtc(v string)`

SetOhlcTimestampStartUtc sets OhlcTimestampStartUtc field to given value.

### HasOhlcTimestampStartUtc

`func (o *ForexDataPoint) HasOhlcTimestampStartUtc() bool`

HasOhlcTimestampStartUtc returns a boolean if a field has been set.

### SetOhlcTimestampStartUtcNil

`func (o *ForexDataPoint) SetOhlcTimestampStartUtcNil(b bool)`

 SetOhlcTimestampStartUtcNil sets the value for OhlcTimestampStartUtc to be an explicit nil

### UnsetOhlcTimestampStartUtc
`func (o *ForexDataPoint) UnsetOhlcTimestampStartUtc()`

UnsetOhlcTimestampStartUtc ensures that no value is present for OhlcTimestampStartUtc, not even an explicit nil
### GetOhlcTimestampEndUtc

`func (o *ForexDataPoint) GetOhlcTimestampEndUtc() string`

GetOhlcTimestampEndUtc returns the OhlcTimestampEndUtc field if non-nil, zero value otherwise.

### GetOhlcTimestampEndUtcOk

`func (o *ForexDataPoint) GetOhlcTimestampEndUtcOk() (*string, bool)`

GetOhlcTimestampEndUtcOk returns a tuple with the OhlcTimestampEndUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOhlcTimestampEndUtc

`func (o *ForexDataPoint) SetOhlcTimestampEndUtc(v string)`

SetOhlcTimestampEndUtc sets OhlcTimestampEndUtc field to given value.

### HasOhlcTimestampEndUtc

`func (o *ForexDataPoint) HasOhlcTimestampEndUtc() bool`

HasOhlcTimestampEndUtc returns a boolean if a field has been set.

### SetOhlcTimestampEndUtcNil

`func (o *ForexDataPoint) SetOhlcTimestampEndUtcNil(b bool)`

 SetOhlcTimestampEndUtcNil sets the value for OhlcTimestampEndUtc to be an explicit nil

### UnsetOhlcTimestampEndUtc
`func (o *ForexDataPoint) UnsetOhlcTimestampEndUtc()`

UnsetOhlcTimestampEndUtc ensures that no value is present for OhlcTimestampEndUtc, not even an explicit nil
### GetOhlcType

`func (o *ForexDataPoint) GetOhlcType() string`

GetOhlcType returns the OhlcType field if non-nil, zero value otherwise.

### GetOhlcTypeOk

`func (o *ForexDataPoint) GetOhlcTypeOk() (*string, bool)`

GetOhlcTypeOk returns a tuple with the OhlcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOhlcType

`func (o *ForexDataPoint) SetOhlcType(v string)`

SetOhlcType sets OhlcType field to given value.

### HasOhlcType

`func (o *ForexDataPoint) HasOhlcType() bool`

HasOhlcType returns a boolean if a field has been set.

### SetOhlcTypeNil

`func (o *ForexDataPoint) SetOhlcTypeNil(b bool)`

 SetOhlcTypeNil sets the value for OhlcType to be an explicit nil

### UnsetOhlcType
`func (o *ForexDataPoint) UnsetOhlcType()`

UnsetOhlcType ensures that no value is present for OhlcType, not even an explicit nil
### GetAnnouncementDatetime

`func (o *ForexDataPoint) GetAnnouncementDatetime() int64`

GetAnnouncementDatetime returns the AnnouncementDatetime field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeOk

`func (o *ForexDataPoint) GetAnnouncementDatetimeOk() (*int64, bool)`

GetAnnouncementDatetimeOk returns a tuple with the AnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetime

`func (o *ForexDataPoint) SetAnnouncementDatetime(v int64)`

SetAnnouncementDatetime sets AnnouncementDatetime field to given value.

### HasAnnouncementDatetime

`func (o *ForexDataPoint) HasAnnouncementDatetime() bool`

HasAnnouncementDatetime returns a boolean if a field has been set.

### SetAnnouncementDatetimeNil

`func (o *ForexDataPoint) SetAnnouncementDatetimeNil(b bool)`

 SetAnnouncementDatetimeNil sets the value for AnnouncementDatetime to be an explicit nil

### UnsetAnnouncementDatetime
`func (o *ForexDataPoint) UnsetAnnouncementDatetime()`

UnsetAnnouncementDatetime ensures that no value is present for AnnouncementDatetime, not even an explicit nil
### GetObservationDatetime

`func (o *ForexDataPoint) GetObservationDatetime() int64`

GetObservationDatetime returns the ObservationDatetime field if non-nil, zero value otherwise.

### GetObservationDatetimeOk

`func (o *ForexDataPoint) GetObservationDatetimeOk() (*int64, bool)`

GetObservationDatetimeOk returns a tuple with the ObservationDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationDatetime

`func (o *ForexDataPoint) SetObservationDatetime(v int64)`

SetObservationDatetime sets ObservationDatetime field to given value.

### HasObservationDatetime

`func (o *ForexDataPoint) HasObservationDatetime() bool`

HasObservationDatetime returns a boolean if a field has been set.

### SetObservationDatetimeNil

`func (o *ForexDataPoint) SetObservationDatetimeNil(b bool)`

 SetObservationDatetimeNil sets the value for ObservationDatetime to be an explicit nil

### UnsetObservationDatetime
`func (o *ForexDataPoint) UnsetObservationDatetime()`

UnsetObservationDatetime ensures that no value is present for ObservationDatetime, not even an explicit nil
### GetObservationDatetimeIso

`func (o *ForexDataPoint) GetObservationDatetimeIso() string`

GetObservationDatetimeIso returns the ObservationDatetimeIso field if non-nil, zero value otherwise.

### GetObservationDatetimeIsoOk

`func (o *ForexDataPoint) GetObservationDatetimeIsoOk() (*string, bool)`

GetObservationDatetimeIsoOk returns a tuple with the ObservationDatetimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationDatetimeIso

`func (o *ForexDataPoint) SetObservationDatetimeIso(v string)`

SetObservationDatetimeIso sets ObservationDatetimeIso field to given value.

### HasObservationDatetimeIso

`func (o *ForexDataPoint) HasObservationDatetimeIso() bool`

HasObservationDatetimeIso returns a boolean if a field has been set.

### SetObservationDatetimeIsoNil

`func (o *ForexDataPoint) SetObservationDatetimeIsoNil(b bool)`

 SetObservationDatetimeIsoNil sets the value for ObservationDatetimeIso to be an explicit nil

### UnsetObservationDatetimeIso
`func (o *ForexDataPoint) UnsetObservationDatetimeIso()`

UnsetObservationDatetimeIso ensures that no value is present for ObservationDatetimeIso, not even an explicit nil
### GetObservationDatetimePrecision

`func (o *ForexDataPoint) GetObservationDatetimePrecision() string`

GetObservationDatetimePrecision returns the ObservationDatetimePrecision field if non-nil, zero value otherwise.

### GetObservationDatetimePrecisionOk

`func (o *ForexDataPoint) GetObservationDatetimePrecisionOk() (*string, bool)`

GetObservationDatetimePrecisionOk returns a tuple with the ObservationDatetimePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationDatetimePrecision

`func (o *ForexDataPoint) SetObservationDatetimePrecision(v string)`

SetObservationDatetimePrecision sets ObservationDatetimePrecision field to given value.

### HasObservationDatetimePrecision

`func (o *ForexDataPoint) HasObservationDatetimePrecision() bool`

HasObservationDatetimePrecision returns a boolean if a field has been set.

### SetObservationDatetimePrecisionNil

`func (o *ForexDataPoint) SetObservationDatetimePrecisionNil(b bool)`

 SetObservationDatetimePrecisionNil sets the value for ObservationDatetimePrecision to be an explicit nil

### UnsetObservationDatetimePrecision
`func (o *ForexDataPoint) UnsetObservationDatetimePrecision()`

UnsetObservationDatetimePrecision ensures that no value is present for ObservationDatetimePrecision, not even an explicit nil
### GetSourceType

`func (o *ForexDataPoint) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ForexDataPoint) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ForexDataPoint) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *ForexDataPoint) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *ForexDataPoint) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *ForexDataPoint) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetSourceChartTimestamp

`func (o *ForexDataPoint) GetSourceChartTimestamp() int64`

GetSourceChartTimestamp returns the SourceChartTimestamp field if non-nil, zero value otherwise.

### GetSourceChartTimestampOk

`func (o *ForexDataPoint) GetSourceChartTimestampOk() (*int64, bool)`

GetSourceChartTimestampOk returns a tuple with the SourceChartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChartTimestamp

`func (o *ForexDataPoint) SetSourceChartTimestamp(v int64)`

SetSourceChartTimestamp sets SourceChartTimestamp field to given value.

### HasSourceChartTimestamp

`func (o *ForexDataPoint) HasSourceChartTimestamp() bool`

HasSourceChartTimestamp returns a boolean if a field has been set.

### SetSourceChartTimestampNil

`func (o *ForexDataPoint) SetSourceChartTimestampNil(b bool)`

 SetSourceChartTimestampNil sets the value for SourceChartTimestamp to be an explicit nil

### UnsetSourceChartTimestamp
`func (o *ForexDataPoint) UnsetSourceChartTimestamp()`

UnsetSourceChartTimestamp ensures that no value is present for SourceChartTimestamp, not even an explicit nil
### GetSourceChartTimestampMs

`func (o *ForexDataPoint) GetSourceChartTimestampMs() int64`

GetSourceChartTimestampMs returns the SourceChartTimestampMs field if non-nil, zero value otherwise.

### GetSourceChartTimestampMsOk

`func (o *ForexDataPoint) GetSourceChartTimestampMsOk() (*int64, bool)`

GetSourceChartTimestampMsOk returns a tuple with the SourceChartTimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChartTimestampMs

`func (o *ForexDataPoint) SetSourceChartTimestampMs(v int64)`

SetSourceChartTimestampMs sets SourceChartTimestampMs field to given value.

### HasSourceChartTimestampMs

`func (o *ForexDataPoint) HasSourceChartTimestampMs() bool`

HasSourceChartTimestampMs returns a boolean if a field has been set.

### SetSourceChartTimestampMsNil

`func (o *ForexDataPoint) SetSourceChartTimestampMsNil(b bool)`

 SetSourceChartTimestampMsNil sets the value for SourceChartTimestampMs to be an explicit nil

### UnsetSourceChartTimestampMs
`func (o *ForexDataPoint) UnsetSourceChartTimestampMs()`

UnsetSourceChartTimestampMs ensures that no value is present for SourceChartTimestampMs, not even an explicit nil
### GetSourceChartPeriod

`func (o *ForexDataPoint) GetSourceChartPeriod() string`

GetSourceChartPeriod returns the SourceChartPeriod field if non-nil, zero value otherwise.

### GetSourceChartPeriodOk

`func (o *ForexDataPoint) GetSourceChartPeriodOk() (*string, bool)`

GetSourceChartPeriodOk returns a tuple with the SourceChartPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChartPeriod

`func (o *ForexDataPoint) SetSourceChartPeriod(v string)`

SetSourceChartPeriod sets SourceChartPeriod field to given value.

### HasSourceChartPeriod

`func (o *ForexDataPoint) HasSourceChartPeriod() bool`

HasSourceChartPeriod returns a boolean if a field has been set.

### SetSourceChartPeriodNil

`func (o *ForexDataPoint) SetSourceChartPeriodNil(b bool)`

 SetSourceChartPeriodNil sets the value for SourceChartPeriod to be an explicit nil

### UnsetSourceChartPeriod
`func (o *ForexDataPoint) UnsetSourceChartPeriod()`

UnsetSourceChartPeriod ensures that no value is present for SourceChartPeriod, not even an explicit nil
### GetQuoteTimeStatus

`func (o *ForexDataPoint) GetQuoteTimeStatus() string`

GetQuoteTimeStatus returns the QuoteTimeStatus field if non-nil, zero value otherwise.

### GetQuoteTimeStatusOk

`func (o *ForexDataPoint) GetQuoteTimeStatusOk() (*string, bool)`

GetQuoteTimeStatusOk returns a tuple with the QuoteTimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTimeStatus

`func (o *ForexDataPoint) SetQuoteTimeStatus(v string)`

SetQuoteTimeStatus sets QuoteTimeStatus field to given value.

### HasQuoteTimeStatus

`func (o *ForexDataPoint) HasQuoteTimeStatus() bool`

HasQuoteTimeStatus returns a boolean if a field has been set.

### SetQuoteTimeStatusNil

`func (o *ForexDataPoint) SetQuoteTimeStatusNil(b bool)`

 SetQuoteTimeStatusNil sets the value for QuoteTimeStatus to be an explicit nil

### UnsetQuoteTimeStatus
`func (o *ForexDataPoint) UnsetQuoteTimeStatus()`

UnsetQuoteTimeStatus ensures that no value is present for QuoteTimeStatus, not even an explicit nil
### GetSamplingMethod

`func (o *ForexDataPoint) GetSamplingMethod() string`

GetSamplingMethod returns the SamplingMethod field if non-nil, zero value otherwise.

### GetSamplingMethodOk

`func (o *ForexDataPoint) GetSamplingMethodOk() (*string, bool)`

GetSamplingMethodOk returns a tuple with the SamplingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingMethod

`func (o *ForexDataPoint) SetSamplingMethod(v string)`

SetSamplingMethod sets SamplingMethod field to given value.

### HasSamplingMethod

`func (o *ForexDataPoint) HasSamplingMethod() bool`

HasSamplingMethod returns a boolean if a field has been set.

### SetSamplingMethodNil

`func (o *ForexDataPoint) SetSamplingMethodNil(b bool)`

 SetSamplingMethodNil sets the value for SamplingMethod to be an explicit nil

### UnsetSamplingMethod
`func (o *ForexDataPoint) UnsetSamplingMethod()`

UnsetSamplingMethod ensures that no value is present for SamplingMethod, not even an explicit nil
### GetPublicationTimeStatus

`func (o *ForexDataPoint) GetPublicationTimeStatus() string`

GetPublicationTimeStatus returns the PublicationTimeStatus field if non-nil, zero value otherwise.

### GetPublicationTimeStatusOk

`func (o *ForexDataPoint) GetPublicationTimeStatusOk() (*string, bool)`

GetPublicationTimeStatusOk returns a tuple with the PublicationTimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationTimeStatus

`func (o *ForexDataPoint) SetPublicationTimeStatus(v string)`

SetPublicationTimeStatus sets PublicationTimeStatus field to given value.

### HasPublicationTimeStatus

`func (o *ForexDataPoint) HasPublicationTimeStatus() bool`

HasPublicationTimeStatus returns a boolean if a field has been set.

### SetPublicationTimeStatusNil

`func (o *ForexDataPoint) SetPublicationTimeStatusNil(b bool)`

 SetPublicationTimeStatusNil sets the value for PublicationTimeStatus to be an explicit nil

### UnsetPublicationTimeStatus
`func (o *ForexDataPoint) UnsetPublicationTimeStatus()`

UnsetPublicationTimeStatus ensures that no value is present for PublicationTimeStatus, not even an explicit nil
### GetPointInTimeSafe

`func (o *ForexDataPoint) GetPointInTimeSafe() bool`

GetPointInTimeSafe returns the PointInTimeSafe field if non-nil, zero value otherwise.

### GetPointInTimeSafeOk

`func (o *ForexDataPoint) GetPointInTimeSafeOk() (*bool, bool)`

GetPointInTimeSafeOk returns a tuple with the PointInTimeSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeSafe

`func (o *ForexDataPoint) SetPointInTimeSafe(v bool)`

SetPointInTimeSafe sets PointInTimeSafe field to given value.

### HasPointInTimeSafe

`func (o *ForexDataPoint) HasPointInTimeSafe() bool`

HasPointInTimeSafe returns a boolean if a field has been set.

### SetPointInTimeSafeNil

`func (o *ForexDataPoint) SetPointInTimeSafeNil(b bool)`

 SetPointInTimeSafeNil sets the value for PointInTimeSafe to be an explicit nil

### UnsetPointInTimeSafe
`func (o *ForexDataPoint) UnsetPointInTimeSafe()`

UnsetPointInTimeSafe ensures that no value is present for PointInTimeSafe, not even an explicit nil
### GetProvenanceVersion

`func (o *ForexDataPoint) GetProvenanceVersion() int64`

GetProvenanceVersion returns the ProvenanceVersion field if non-nil, zero value otherwise.

### GetProvenanceVersionOk

`func (o *ForexDataPoint) GetProvenanceVersionOk() (*int64, bool)`

GetProvenanceVersionOk returns a tuple with the ProvenanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenanceVersion

`func (o *ForexDataPoint) SetProvenanceVersion(v int64)`

SetProvenanceVersion sets ProvenanceVersion field to given value.

### HasProvenanceVersion

`func (o *ForexDataPoint) HasProvenanceVersion() bool`

HasProvenanceVersion returns a boolean if a field has been set.

### SetProvenanceVersionNil

`func (o *ForexDataPoint) SetProvenanceVersionNil(b bool)`

 SetProvenanceVersionNil sets the value for ProvenanceVersion to be an explicit nil

### UnsetProvenanceVersion
`func (o *ForexDataPoint) UnsetProvenanceVersion()`

UnsetProvenanceVersion ensures that no value is present for ProvenanceVersion, not even an explicit nil
### GetSource

`func (o *ForexDataPoint) GetSource() ForexPointSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ForexDataPoint) GetSourceOk() (*ForexPointSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ForexDataPoint) SetSource(v ForexPointSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ForexDataPoint) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSma20

`func (o *ForexDataPoint) GetSma20() float64`

GetSma20 returns the Sma20 field if non-nil, zero value otherwise.

### GetSma20Ok

`func (o *ForexDataPoint) GetSma20Ok() (*float64, bool)`

GetSma20Ok returns a tuple with the Sma20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSma20

`func (o *ForexDataPoint) SetSma20(v float64)`

SetSma20 sets Sma20 field to given value.

### HasSma20

`func (o *ForexDataPoint) HasSma20() bool`

HasSma20 returns a boolean if a field has been set.

### SetSma20Nil

`func (o *ForexDataPoint) SetSma20Nil(b bool)`

 SetSma20Nil sets the value for Sma20 to be an explicit nil

### UnsetSma20
`func (o *ForexDataPoint) UnsetSma20()`

UnsetSma20 ensures that no value is present for Sma20, not even an explicit nil
### GetSma50

`func (o *ForexDataPoint) GetSma50() float64`

GetSma50 returns the Sma50 field if non-nil, zero value otherwise.

### GetSma50Ok

`func (o *ForexDataPoint) GetSma50Ok() (*float64, bool)`

GetSma50Ok returns a tuple with the Sma50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSma50

`func (o *ForexDataPoint) SetSma50(v float64)`

SetSma50 sets Sma50 field to given value.

### HasSma50

`func (o *ForexDataPoint) HasSma50() bool`

HasSma50 returns a boolean if a field has been set.

### SetSma50Nil

`func (o *ForexDataPoint) SetSma50Nil(b bool)`

 SetSma50Nil sets the value for Sma50 to be an explicit nil

### UnsetSma50
`func (o *ForexDataPoint) UnsetSma50()`

UnsetSma50 ensures that no value is present for Sma50, not even an explicit nil
### GetSma200

`func (o *ForexDataPoint) GetSma200() float64`

GetSma200 returns the Sma200 field if non-nil, zero value otherwise.

### GetSma200Ok

`func (o *ForexDataPoint) GetSma200Ok() (*float64, bool)`

GetSma200Ok returns a tuple with the Sma200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSma200

`func (o *ForexDataPoint) SetSma200(v float64)`

SetSma200 sets Sma200 field to given value.

### HasSma200

`func (o *ForexDataPoint) HasSma200() bool`

HasSma200 returns a boolean if a field has been set.

### SetSma200Nil

`func (o *ForexDataPoint) SetSma200Nil(b bool)`

 SetSma200Nil sets the value for Sma200 to be an explicit nil

### UnsetSma200
`func (o *ForexDataPoint) UnsetSma200()`

UnsetSma200 ensures that no value is present for Sma200, not even an explicit nil
### GetEma12

`func (o *ForexDataPoint) GetEma12() float64`

GetEma12 returns the Ema12 field if non-nil, zero value otherwise.

### GetEma12Ok

`func (o *ForexDataPoint) GetEma12Ok() (*float64, bool)`

GetEma12Ok returns a tuple with the Ema12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEma12

`func (o *ForexDataPoint) SetEma12(v float64)`

SetEma12 sets Ema12 field to given value.

### HasEma12

`func (o *ForexDataPoint) HasEma12() bool`

HasEma12 returns a boolean if a field has been set.

### SetEma12Nil

`func (o *ForexDataPoint) SetEma12Nil(b bool)`

 SetEma12Nil sets the value for Ema12 to be an explicit nil

### UnsetEma12
`func (o *ForexDataPoint) UnsetEma12()`

UnsetEma12 ensures that no value is present for Ema12, not even an explicit nil
### GetEma20

`func (o *ForexDataPoint) GetEma20() float64`

GetEma20 returns the Ema20 field if non-nil, zero value otherwise.

### GetEma20Ok

`func (o *ForexDataPoint) GetEma20Ok() (*float64, bool)`

GetEma20Ok returns a tuple with the Ema20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEma20

`func (o *ForexDataPoint) SetEma20(v float64)`

SetEma20 sets Ema20 field to given value.

### HasEma20

`func (o *ForexDataPoint) HasEma20() bool`

HasEma20 returns a boolean if a field has been set.

### SetEma20Nil

`func (o *ForexDataPoint) SetEma20Nil(b bool)`

 SetEma20Nil sets the value for Ema20 to be an explicit nil

### UnsetEma20
`func (o *ForexDataPoint) UnsetEma20()`

UnsetEma20 ensures that no value is present for Ema20, not even an explicit nil
### GetEma26

`func (o *ForexDataPoint) GetEma26() float64`

GetEma26 returns the Ema26 field if non-nil, zero value otherwise.

### GetEma26Ok

`func (o *ForexDataPoint) GetEma26Ok() (*float64, bool)`

GetEma26Ok returns a tuple with the Ema26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEma26

`func (o *ForexDataPoint) SetEma26(v float64)`

SetEma26 sets Ema26 field to given value.

### HasEma26

`func (o *ForexDataPoint) HasEma26() bool`

HasEma26 returns a boolean if a field has been set.

### SetEma26Nil

`func (o *ForexDataPoint) SetEma26Nil(b bool)`

 SetEma26Nil sets the value for Ema26 to be an explicit nil

### UnsetEma26
`func (o *ForexDataPoint) UnsetEma26()`

UnsetEma26 ensures that no value is present for Ema26, not even an explicit nil
### GetEma50

`func (o *ForexDataPoint) GetEma50() float64`

GetEma50 returns the Ema50 field if non-nil, zero value otherwise.

### GetEma50Ok

`func (o *ForexDataPoint) GetEma50Ok() (*float64, bool)`

GetEma50Ok returns a tuple with the Ema50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEma50

`func (o *ForexDataPoint) SetEma50(v float64)`

SetEma50 sets Ema50 field to given value.

### HasEma50

`func (o *ForexDataPoint) HasEma50() bool`

HasEma50 returns a boolean if a field has been set.

### SetEma50Nil

`func (o *ForexDataPoint) SetEma50Nil(b bool)`

 SetEma50Nil sets the value for Ema50 to be an explicit nil

### UnsetEma50
`func (o *ForexDataPoint) UnsetEma50()`

UnsetEma50 ensures that no value is present for Ema50, not even an explicit nil
### GetEma200

`func (o *ForexDataPoint) GetEma200() float64`

GetEma200 returns the Ema200 field if non-nil, zero value otherwise.

### GetEma200Ok

`func (o *ForexDataPoint) GetEma200Ok() (*float64, bool)`

GetEma200Ok returns a tuple with the Ema200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEma200

`func (o *ForexDataPoint) SetEma200(v float64)`

SetEma200 sets Ema200 field to given value.

### HasEma200

`func (o *ForexDataPoint) HasEma200() bool`

HasEma200 returns a boolean if a field has been set.

### SetEma200Nil

`func (o *ForexDataPoint) SetEma200Nil(b bool)`

 SetEma200Nil sets the value for Ema200 to be an explicit nil

### UnsetEma200
`func (o *ForexDataPoint) UnsetEma200()`

UnsetEma200 ensures that no value is present for Ema200, not even an explicit nil
### GetRsi14

`func (o *ForexDataPoint) GetRsi14() float64`

GetRsi14 returns the Rsi14 field if non-nil, zero value otherwise.

### GetRsi14Ok

`func (o *ForexDataPoint) GetRsi14Ok() (*float64, bool)`

GetRsi14Ok returns a tuple with the Rsi14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsi14

`func (o *ForexDataPoint) SetRsi14(v float64)`

SetRsi14 sets Rsi14 field to given value.

### HasRsi14

`func (o *ForexDataPoint) HasRsi14() bool`

HasRsi14 returns a boolean if a field has been set.

### SetRsi14Nil

`func (o *ForexDataPoint) SetRsi14Nil(b bool)`

 SetRsi14Nil sets the value for Rsi14 to be an explicit nil

### UnsetRsi14
`func (o *ForexDataPoint) UnsetRsi14()`

UnsetRsi14 ensures that no value is present for Rsi14, not even an explicit nil
### GetAtr14

`func (o *ForexDataPoint) GetAtr14() float64`

GetAtr14 returns the Atr14 field if non-nil, zero value otherwise.

### GetAtr14Ok

`func (o *ForexDataPoint) GetAtr14Ok() (*float64, bool)`

GetAtr14Ok returns a tuple with the Atr14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtr14

`func (o *ForexDataPoint) SetAtr14(v float64)`

SetAtr14 sets Atr14 field to given value.

### HasAtr14

`func (o *ForexDataPoint) HasAtr14() bool`

HasAtr14 returns a boolean if a field has been set.

### SetAtr14Nil

`func (o *ForexDataPoint) SetAtr14Nil(b bool)`

 SetAtr14Nil sets the value for Atr14 to be an explicit nil

### UnsetAtr14
`func (o *ForexDataPoint) UnsetAtr14()`

UnsetAtr14 ensures that no value is present for Atr14, not even an explicit nil
### GetAdx14

`func (o *ForexDataPoint) GetAdx14() float64`

GetAdx14 returns the Adx14 field if non-nil, zero value otherwise.

### GetAdx14Ok

`func (o *ForexDataPoint) GetAdx14Ok() (*float64, bool)`

GetAdx14Ok returns a tuple with the Adx14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdx14

`func (o *ForexDataPoint) SetAdx14(v float64)`

SetAdx14 sets Adx14 field to given value.

### HasAdx14

`func (o *ForexDataPoint) HasAdx14() bool`

HasAdx14 returns a boolean if a field has been set.

### SetAdx14Nil

`func (o *ForexDataPoint) SetAdx14Nil(b bool)`

 SetAdx14Nil sets the value for Adx14 to be an explicit nil

### UnsetAdx14
`func (o *ForexDataPoint) UnsetAdx14()`

UnsetAdx14 ensures that no value is present for Adx14, not even an explicit nil
### GetStochK14

`func (o *ForexDataPoint) GetStochK14() float64`

GetStochK14 returns the StochK14 field if non-nil, zero value otherwise.

### GetStochK14Ok

`func (o *ForexDataPoint) GetStochK14Ok() (*float64, bool)`

GetStochK14Ok returns a tuple with the StochK14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStochK14

`func (o *ForexDataPoint) SetStochK14(v float64)`

SetStochK14 sets StochK14 field to given value.

### HasStochK14

`func (o *ForexDataPoint) HasStochK14() bool`

HasStochK14 returns a boolean if a field has been set.

### SetStochK14Nil

`func (o *ForexDataPoint) SetStochK14Nil(b bool)`

 SetStochK14Nil sets the value for StochK14 to be an explicit nil

### UnsetStochK14
`func (o *ForexDataPoint) UnsetStochK14()`

UnsetStochK14 ensures that no value is present for StochK14, not even an explicit nil
### GetStochD3

`func (o *ForexDataPoint) GetStochD3() float64`

GetStochD3 returns the StochD3 field if non-nil, zero value otherwise.

### GetStochD3Ok

`func (o *ForexDataPoint) GetStochD3Ok() (*float64, bool)`

GetStochD3Ok returns a tuple with the StochD3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStochD3

`func (o *ForexDataPoint) SetStochD3(v float64)`

SetStochD3 sets StochD3 field to given value.

### HasStochD3

`func (o *ForexDataPoint) HasStochD3() bool`

HasStochD3 returns a boolean if a field has been set.

### SetStochD3Nil

`func (o *ForexDataPoint) SetStochD3Nil(b bool)`

 SetStochD3Nil sets the value for StochD3 to be an explicit nil

### UnsetStochD3
`func (o *ForexDataPoint) UnsetStochD3()`

UnsetStochD3 ensures that no value is present for StochD3, not even an explicit nil
### GetWilliamsR14

`func (o *ForexDataPoint) GetWilliamsR14() float64`

GetWilliamsR14 returns the WilliamsR14 field if non-nil, zero value otherwise.

### GetWilliamsR14Ok

`func (o *ForexDataPoint) GetWilliamsR14Ok() (*float64, bool)`

GetWilliamsR14Ok returns a tuple with the WilliamsR14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWilliamsR14

`func (o *ForexDataPoint) SetWilliamsR14(v float64)`

SetWilliamsR14 sets WilliamsR14 field to given value.

### HasWilliamsR14

`func (o *ForexDataPoint) HasWilliamsR14() bool`

HasWilliamsR14 returns a boolean if a field has been set.

### SetWilliamsR14Nil

`func (o *ForexDataPoint) SetWilliamsR14Nil(b bool)`

 SetWilliamsR14Nil sets the value for WilliamsR14 to be an explicit nil

### UnsetWilliamsR14
`func (o *ForexDataPoint) UnsetWilliamsR14()`

UnsetWilliamsR14 ensures that no value is present for WilliamsR14, not even an explicit nil
### GetCci20

`func (o *ForexDataPoint) GetCci20() float64`

GetCci20 returns the Cci20 field if non-nil, zero value otherwise.

### GetCci20Ok

`func (o *ForexDataPoint) GetCci20Ok() (*float64, bool)`

GetCci20Ok returns a tuple with the Cci20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCci20

`func (o *ForexDataPoint) SetCci20(v float64)`

SetCci20 sets Cci20 field to given value.

### HasCci20

`func (o *ForexDataPoint) HasCci20() bool`

HasCci20 returns a boolean if a field has been set.

### SetCci20Nil

`func (o *ForexDataPoint) SetCci20Nil(b bool)`

 SetCci20Nil sets the value for Cci20 to be an explicit nil

### UnsetCci20
`func (o *ForexDataPoint) UnsetCci20()`

UnsetCci20 ensures that no value is present for Cci20, not even an explicit nil
### GetDonchianUpper20

`func (o *ForexDataPoint) GetDonchianUpper20() float64`

GetDonchianUpper20 returns the DonchianUpper20 field if non-nil, zero value otherwise.

### GetDonchianUpper20Ok

`func (o *ForexDataPoint) GetDonchianUpper20Ok() (*float64, bool)`

GetDonchianUpper20Ok returns a tuple with the DonchianUpper20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonchianUpper20

`func (o *ForexDataPoint) SetDonchianUpper20(v float64)`

SetDonchianUpper20 sets DonchianUpper20 field to given value.

### HasDonchianUpper20

`func (o *ForexDataPoint) HasDonchianUpper20() bool`

HasDonchianUpper20 returns a boolean if a field has been set.

### SetDonchianUpper20Nil

`func (o *ForexDataPoint) SetDonchianUpper20Nil(b bool)`

 SetDonchianUpper20Nil sets the value for DonchianUpper20 to be an explicit nil

### UnsetDonchianUpper20
`func (o *ForexDataPoint) UnsetDonchianUpper20()`

UnsetDonchianUpper20 ensures that no value is present for DonchianUpper20, not even an explicit nil
### GetDonchianMiddle20

`func (o *ForexDataPoint) GetDonchianMiddle20() float64`

GetDonchianMiddle20 returns the DonchianMiddle20 field if non-nil, zero value otherwise.

### GetDonchianMiddle20Ok

`func (o *ForexDataPoint) GetDonchianMiddle20Ok() (*float64, bool)`

GetDonchianMiddle20Ok returns a tuple with the DonchianMiddle20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonchianMiddle20

`func (o *ForexDataPoint) SetDonchianMiddle20(v float64)`

SetDonchianMiddle20 sets DonchianMiddle20 field to given value.

### HasDonchianMiddle20

`func (o *ForexDataPoint) HasDonchianMiddle20() bool`

HasDonchianMiddle20 returns a boolean if a field has been set.

### SetDonchianMiddle20Nil

`func (o *ForexDataPoint) SetDonchianMiddle20Nil(b bool)`

 SetDonchianMiddle20Nil sets the value for DonchianMiddle20 to be an explicit nil

### UnsetDonchianMiddle20
`func (o *ForexDataPoint) UnsetDonchianMiddle20()`

UnsetDonchianMiddle20 ensures that no value is present for DonchianMiddle20, not even an explicit nil
### GetDonchianLower20

`func (o *ForexDataPoint) GetDonchianLower20() float64`

GetDonchianLower20 returns the DonchianLower20 field if non-nil, zero value otherwise.

### GetDonchianLower20Ok

`func (o *ForexDataPoint) GetDonchianLower20Ok() (*float64, bool)`

GetDonchianLower20Ok returns a tuple with the DonchianLower20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonchianLower20

`func (o *ForexDataPoint) SetDonchianLower20(v float64)`

SetDonchianLower20 sets DonchianLower20 field to given value.

### HasDonchianLower20

`func (o *ForexDataPoint) HasDonchianLower20() bool`

HasDonchianLower20 returns a boolean if a field has been set.

### SetDonchianLower20Nil

`func (o *ForexDataPoint) SetDonchianLower20Nil(b bool)`

 SetDonchianLower20Nil sets the value for DonchianLower20 to be an explicit nil

### UnsetDonchianLower20
`func (o *ForexDataPoint) UnsetDonchianLower20()`

UnsetDonchianLower20 ensures that no value is present for DonchianLower20, not even an explicit nil
### GetMacd

`func (o *ForexDataPoint) GetMacd() float64`

GetMacd returns the Macd field if non-nil, zero value otherwise.

### GetMacdOk

`func (o *ForexDataPoint) GetMacdOk() (*float64, bool)`

GetMacdOk returns a tuple with the Macd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacd

`func (o *ForexDataPoint) SetMacd(v float64)`

SetMacd sets Macd field to given value.

### HasMacd

`func (o *ForexDataPoint) HasMacd() bool`

HasMacd returns a boolean if a field has been set.

### SetMacdNil

`func (o *ForexDataPoint) SetMacdNil(b bool)`

 SetMacdNil sets the value for Macd to be an explicit nil

### UnsetMacd
`func (o *ForexDataPoint) UnsetMacd()`

UnsetMacd ensures that no value is present for Macd, not even an explicit nil
### GetMacdSignal

`func (o *ForexDataPoint) GetMacdSignal() float64`

GetMacdSignal returns the MacdSignal field if non-nil, zero value otherwise.

### GetMacdSignalOk

`func (o *ForexDataPoint) GetMacdSignalOk() (*float64, bool)`

GetMacdSignalOk returns a tuple with the MacdSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacdSignal

`func (o *ForexDataPoint) SetMacdSignal(v float64)`

SetMacdSignal sets MacdSignal field to given value.

### HasMacdSignal

`func (o *ForexDataPoint) HasMacdSignal() bool`

HasMacdSignal returns a boolean if a field has been set.

### SetMacdSignalNil

`func (o *ForexDataPoint) SetMacdSignalNil(b bool)`

 SetMacdSignalNil sets the value for MacdSignal to be an explicit nil

### UnsetMacdSignal
`func (o *ForexDataPoint) UnsetMacdSignal()`

UnsetMacdSignal ensures that no value is present for MacdSignal, not even an explicit nil
### GetMacdHistogram

`func (o *ForexDataPoint) GetMacdHistogram() float64`

GetMacdHistogram returns the MacdHistogram field if non-nil, zero value otherwise.

### GetMacdHistogramOk

`func (o *ForexDataPoint) GetMacdHistogramOk() (*float64, bool)`

GetMacdHistogramOk returns a tuple with the MacdHistogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacdHistogram

`func (o *ForexDataPoint) SetMacdHistogram(v float64)`

SetMacdHistogram sets MacdHistogram field to given value.

### HasMacdHistogram

`func (o *ForexDataPoint) HasMacdHistogram() bool`

HasMacdHistogram returns a boolean if a field has been set.

### SetMacdHistogramNil

`func (o *ForexDataPoint) SetMacdHistogramNil(b bool)`

 SetMacdHistogramNil sets the value for MacdHistogram to be an explicit nil

### UnsetMacdHistogram
`func (o *ForexDataPoint) UnsetMacdHistogram()`

UnsetMacdHistogram ensures that no value is present for MacdHistogram, not even an explicit nil
### GetBbUpper

`func (o *ForexDataPoint) GetBbUpper() float64`

GetBbUpper returns the BbUpper field if non-nil, zero value otherwise.

### GetBbUpperOk

`func (o *ForexDataPoint) GetBbUpperOk() (*float64, bool)`

GetBbUpperOk returns a tuple with the BbUpper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbUpper

`func (o *ForexDataPoint) SetBbUpper(v float64)`

SetBbUpper sets BbUpper field to given value.

### HasBbUpper

`func (o *ForexDataPoint) HasBbUpper() bool`

HasBbUpper returns a boolean if a field has been set.

### SetBbUpperNil

`func (o *ForexDataPoint) SetBbUpperNil(b bool)`

 SetBbUpperNil sets the value for BbUpper to be an explicit nil

### UnsetBbUpper
`func (o *ForexDataPoint) UnsetBbUpper()`

UnsetBbUpper ensures that no value is present for BbUpper, not even an explicit nil
### GetBbMiddle

`func (o *ForexDataPoint) GetBbMiddle() float64`

GetBbMiddle returns the BbMiddle field if non-nil, zero value otherwise.

### GetBbMiddleOk

`func (o *ForexDataPoint) GetBbMiddleOk() (*float64, bool)`

GetBbMiddleOk returns a tuple with the BbMiddle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbMiddle

`func (o *ForexDataPoint) SetBbMiddle(v float64)`

SetBbMiddle sets BbMiddle field to given value.

### HasBbMiddle

`func (o *ForexDataPoint) HasBbMiddle() bool`

HasBbMiddle returns a boolean if a field has been set.

### SetBbMiddleNil

`func (o *ForexDataPoint) SetBbMiddleNil(b bool)`

 SetBbMiddleNil sets the value for BbMiddle to be an explicit nil

### UnsetBbMiddle
`func (o *ForexDataPoint) UnsetBbMiddle()`

UnsetBbMiddle ensures that no value is present for BbMiddle, not even an explicit nil
### GetBbLower

`func (o *ForexDataPoint) GetBbLower() float64`

GetBbLower returns the BbLower field if non-nil, zero value otherwise.

### GetBbLowerOk

`func (o *ForexDataPoint) GetBbLowerOk() (*float64, bool)`

GetBbLowerOk returns a tuple with the BbLower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbLower

`func (o *ForexDataPoint) SetBbLower(v float64)`

SetBbLower sets BbLower field to given value.

### HasBbLower

`func (o *ForexDataPoint) HasBbLower() bool`

HasBbLower returns a boolean if a field has been set.

### SetBbLowerNil

`func (o *ForexDataPoint) SetBbLowerNil(b bool)`

 SetBbLowerNil sets the value for BbLower to be an explicit nil

### UnsetBbLower
`func (o *ForexDataPoint) UnsetBbLower()`

UnsetBbLower ensures that no value is present for BbLower, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


