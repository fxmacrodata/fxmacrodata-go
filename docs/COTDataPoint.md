# COTDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**AnnouncementDatetime** | Pointer to **NullableInt64** |  | [optional] 
**OpenInterest** | Pointer to **NullableInt64** |  | [optional] 
**NoncommercialLong** | Pointer to **NullableInt64** |  | [optional] 
**NoncommercialShort** | Pointer to **NullableInt64** |  | [optional] 
**NoncommercialNet** | Pointer to **NullableInt64** |  | [optional] 
**NoncommercialSpread** | Pointer to **NullableInt64** |  | [optional] 
**CommercialLong** | Pointer to **NullableInt64** |  | [optional] 
**CommercialShort** | Pointer to **NullableInt64** |  | [optional] 
**CommercialNet** | Pointer to **NullableInt64** |  | [optional] 
**TotalReportableLong** | Pointer to **NullableInt64** |  | [optional] 
**TotalReportableShort** | Pointer to **NullableInt64** |  | [optional] 
**NonreportableLong** | Pointer to **NullableInt64** |  | [optional] 
**NonreportableShort** | Pointer to **NullableInt64** |  | [optional] 
**OpenInterestZscore** | Pointer to **NullableFloat64** |  | [optional] 
**NoncommercialLongZscore** | Pointer to **NullableFloat64** |  | [optional] 
**NoncommercialShortZscore** | Pointer to **NullableFloat64** |  | [optional] 
**NoncommercialNetZscore** | Pointer to **NullableFloat64** |  | [optional] 
**NoncommercialSpreadZscore** | Pointer to **NullableFloat64** |  | [optional] 
**CommercialLongZscore** | Pointer to **NullableFloat64** |  | [optional] 
**CommercialShortZscore** | Pointer to **NullableFloat64** |  | [optional] 
**CommercialNetZscore** | Pointer to **NullableFloat64** |  | [optional] 
**TotalReportableLongZscore** | Pointer to **NullableFloat64** |  | [optional] 
**TotalReportableShortZscore** | Pointer to **NullableFloat64** |  | [optional] 
**NonreportableLongZscore** | Pointer to **NullableFloat64** |  | [optional] 
**NonreportableShortZscore** | Pointer to **NullableFloat64** |  | [optional] 
**ReportDate** | Pointer to **NullableString** |  | [optional] 
**CutoffDate** | Pointer to **NullableString** |  | [optional] 
**ReleaseDate** | Pointer to **NullableString** |  | [optional] 
**ReleaseDatetime** | Pointer to **NullableString** |  | [optional] 
**ReleaseDateConfirmed** | Pointer to **NullableBool** |  | [optional] 
**ReleaseTimeAssumed** | Pointer to **NullableBool** |  | [optional] 
**ReleaseSource** | Pointer to **NullableString** |  | [optional] 
**ReleaseSourceUrl** | Pointer to **NullableString** |  | [optional] 
**HolidayAdjustedRelease** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCOTDataPoint

`func NewCOTDataPoint(date string, ) *COTDataPoint`

NewCOTDataPoint instantiates a new COTDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCOTDataPointWithDefaults

`func NewCOTDataPointWithDefaults() *COTDataPoint`

NewCOTDataPointWithDefaults instantiates a new COTDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *COTDataPoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *COTDataPoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *COTDataPoint) SetDate(v string)`

SetDate sets Date field to given value.


### GetAnnouncementDatetime

`func (o *COTDataPoint) GetAnnouncementDatetime() int64`

GetAnnouncementDatetime returns the AnnouncementDatetime field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeOk

`func (o *COTDataPoint) GetAnnouncementDatetimeOk() (*int64, bool)`

GetAnnouncementDatetimeOk returns a tuple with the AnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetime

`func (o *COTDataPoint) SetAnnouncementDatetime(v int64)`

SetAnnouncementDatetime sets AnnouncementDatetime field to given value.

### HasAnnouncementDatetime

`func (o *COTDataPoint) HasAnnouncementDatetime() bool`

HasAnnouncementDatetime returns a boolean if a field has been set.

### SetAnnouncementDatetimeNil

`func (o *COTDataPoint) SetAnnouncementDatetimeNil(b bool)`

 SetAnnouncementDatetimeNil sets the value for AnnouncementDatetime to be an explicit nil

### UnsetAnnouncementDatetime
`func (o *COTDataPoint) UnsetAnnouncementDatetime()`

UnsetAnnouncementDatetime ensures that no value is present for AnnouncementDatetime, not even an explicit nil
### GetOpenInterest

`func (o *COTDataPoint) GetOpenInterest() int64`

GetOpenInterest returns the OpenInterest field if non-nil, zero value otherwise.

### GetOpenInterestOk

`func (o *COTDataPoint) GetOpenInterestOk() (*int64, bool)`

GetOpenInterestOk returns a tuple with the OpenInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInterest

`func (o *COTDataPoint) SetOpenInterest(v int64)`

SetOpenInterest sets OpenInterest field to given value.

### HasOpenInterest

`func (o *COTDataPoint) HasOpenInterest() bool`

HasOpenInterest returns a boolean if a field has been set.

### SetOpenInterestNil

`func (o *COTDataPoint) SetOpenInterestNil(b bool)`

 SetOpenInterestNil sets the value for OpenInterest to be an explicit nil

### UnsetOpenInterest
`func (o *COTDataPoint) UnsetOpenInterest()`

UnsetOpenInterest ensures that no value is present for OpenInterest, not even an explicit nil
### GetNoncommercialLong

`func (o *COTDataPoint) GetNoncommercialLong() int64`

GetNoncommercialLong returns the NoncommercialLong field if non-nil, zero value otherwise.

### GetNoncommercialLongOk

`func (o *COTDataPoint) GetNoncommercialLongOk() (*int64, bool)`

GetNoncommercialLongOk returns a tuple with the NoncommercialLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoncommercialLong

`func (o *COTDataPoint) SetNoncommercialLong(v int64)`

SetNoncommercialLong sets NoncommercialLong field to given value.

### HasNoncommercialLong

`func (o *COTDataPoint) HasNoncommercialLong() bool`

HasNoncommercialLong returns a boolean if a field has been set.

### SetNoncommercialLongNil

`func (o *COTDataPoint) SetNoncommercialLongNil(b bool)`

 SetNoncommercialLongNil sets the value for NoncommercialLong to be an explicit nil

### UnsetNoncommercialLong
`func (o *COTDataPoint) UnsetNoncommercialLong()`

UnsetNoncommercialLong ensures that no value is present for NoncommercialLong, not even an explicit nil
### GetNoncommercialShort

`func (o *COTDataPoint) GetNoncommercialShort() int64`

GetNoncommercialShort returns the NoncommercialShort field if non-nil, zero value otherwise.

### GetNoncommercialShortOk

`func (o *COTDataPoint) GetNoncommercialShortOk() (*int64, bool)`

GetNoncommercialShortOk returns a tuple with the NoncommercialShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoncommercialShort

`func (o *COTDataPoint) SetNoncommercialShort(v int64)`

SetNoncommercialShort sets NoncommercialShort field to given value.

### HasNoncommercialShort

`func (o *COTDataPoint) HasNoncommercialShort() bool`

HasNoncommercialShort returns a boolean if a field has been set.

### SetNoncommercialShortNil

`func (o *COTDataPoint) SetNoncommercialShortNil(b bool)`

 SetNoncommercialShortNil sets the value for NoncommercialShort to be an explicit nil

### UnsetNoncommercialShort
`func (o *COTDataPoint) UnsetNoncommercialShort()`

UnsetNoncommercialShort ensures that no value is present for NoncommercialShort, not even an explicit nil
### GetNoncommercialNet

`func (o *COTDataPoint) GetNoncommercialNet() int64`

GetNoncommercialNet returns the NoncommercialNet field if non-nil, zero value otherwise.

### GetNoncommercialNetOk

`func (o *COTDataPoint) GetNoncommercialNetOk() (*int64, bool)`

GetNoncommercialNetOk returns a tuple with the NoncommercialNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoncommercialNet

`func (o *COTDataPoint) SetNoncommercialNet(v int64)`

SetNoncommercialNet sets NoncommercialNet field to given value.

### HasNoncommercialNet

`func (o *COTDataPoint) HasNoncommercialNet() bool`

HasNoncommercialNet returns a boolean if a field has been set.

### SetNoncommercialNetNil

`func (o *COTDataPoint) SetNoncommercialNetNil(b bool)`

 SetNoncommercialNetNil sets the value for NoncommercialNet to be an explicit nil

### UnsetNoncommercialNet
`func (o *COTDataPoint) UnsetNoncommercialNet()`

UnsetNoncommercialNet ensures that no value is present for NoncommercialNet, not even an explicit nil
### GetNoncommercialSpread

`func (o *COTDataPoint) GetNoncommercialSpread() int64`

GetNoncommercialSpread returns the NoncommercialSpread field if non-nil, zero value otherwise.

### GetNoncommercialSpreadOk

`func (o *COTDataPoint) GetNoncommercialSpreadOk() (*int64, bool)`

GetNoncommercialSpreadOk returns a tuple with the NoncommercialSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoncommercialSpread

`func (o *COTDataPoint) SetNoncommercialSpread(v int64)`

SetNoncommercialSpread sets NoncommercialSpread field to given value.

### HasNoncommercialSpread

`func (o *COTDataPoint) HasNoncommercialSpread() bool`

HasNoncommercialSpread returns a boolean if a field has been set.

### SetNoncommercialSpreadNil

`func (o *COTDataPoint) SetNoncommercialSpreadNil(b bool)`

 SetNoncommercialSpreadNil sets the value for NoncommercialSpread to be an explicit nil

### UnsetNoncommercialSpread
`func (o *COTDataPoint) UnsetNoncommercialSpread()`

UnsetNoncommercialSpread ensures that no value is present for NoncommercialSpread, not even an explicit nil
### GetCommercialLong

`func (o *COTDataPoint) GetCommercialLong() int64`

GetCommercialLong returns the CommercialLong field if non-nil, zero value otherwise.

### GetCommercialLongOk

`func (o *COTDataPoint) GetCommercialLongOk() (*int64, bool)`

GetCommercialLongOk returns a tuple with the CommercialLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialLong

`func (o *COTDataPoint) SetCommercialLong(v int64)`

SetCommercialLong sets CommercialLong field to given value.

### HasCommercialLong

`func (o *COTDataPoint) HasCommercialLong() bool`

HasCommercialLong returns a boolean if a field has been set.

### SetCommercialLongNil

`func (o *COTDataPoint) SetCommercialLongNil(b bool)`

 SetCommercialLongNil sets the value for CommercialLong to be an explicit nil

### UnsetCommercialLong
`func (o *COTDataPoint) UnsetCommercialLong()`

UnsetCommercialLong ensures that no value is present for CommercialLong, not even an explicit nil
### GetCommercialShort

`func (o *COTDataPoint) GetCommercialShort() int64`

GetCommercialShort returns the CommercialShort field if non-nil, zero value otherwise.

### GetCommercialShortOk

`func (o *COTDataPoint) GetCommercialShortOk() (*int64, bool)`

GetCommercialShortOk returns a tuple with the CommercialShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialShort

`func (o *COTDataPoint) SetCommercialShort(v int64)`

SetCommercialShort sets CommercialShort field to given value.

### HasCommercialShort

`func (o *COTDataPoint) HasCommercialShort() bool`

HasCommercialShort returns a boolean if a field has been set.

### SetCommercialShortNil

`func (o *COTDataPoint) SetCommercialShortNil(b bool)`

 SetCommercialShortNil sets the value for CommercialShort to be an explicit nil

### UnsetCommercialShort
`func (o *COTDataPoint) UnsetCommercialShort()`

UnsetCommercialShort ensures that no value is present for CommercialShort, not even an explicit nil
### GetCommercialNet

`func (o *COTDataPoint) GetCommercialNet() int64`

GetCommercialNet returns the CommercialNet field if non-nil, zero value otherwise.

### GetCommercialNetOk

`func (o *COTDataPoint) GetCommercialNetOk() (*int64, bool)`

GetCommercialNetOk returns a tuple with the CommercialNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialNet

`func (o *COTDataPoint) SetCommercialNet(v int64)`

SetCommercialNet sets CommercialNet field to given value.

### HasCommercialNet

`func (o *COTDataPoint) HasCommercialNet() bool`

HasCommercialNet returns a boolean if a field has been set.

### SetCommercialNetNil

`func (o *COTDataPoint) SetCommercialNetNil(b bool)`

 SetCommercialNetNil sets the value for CommercialNet to be an explicit nil

### UnsetCommercialNet
`func (o *COTDataPoint) UnsetCommercialNet()`

UnsetCommercialNet ensures that no value is present for CommercialNet, not even an explicit nil
### GetTotalReportableLong

`func (o *COTDataPoint) GetTotalReportableLong() int64`

GetTotalReportableLong returns the TotalReportableLong field if non-nil, zero value otherwise.

### GetTotalReportableLongOk

`func (o *COTDataPoint) GetTotalReportableLongOk() (*int64, bool)`

GetTotalReportableLongOk returns a tuple with the TotalReportableLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReportableLong

`func (o *COTDataPoint) SetTotalReportableLong(v int64)`

SetTotalReportableLong sets TotalReportableLong field to given value.

### HasTotalReportableLong

`func (o *COTDataPoint) HasTotalReportableLong() bool`

HasTotalReportableLong returns a boolean if a field has been set.

### SetTotalReportableLongNil

`func (o *COTDataPoint) SetTotalReportableLongNil(b bool)`

 SetTotalReportableLongNil sets the value for TotalReportableLong to be an explicit nil

### UnsetTotalReportableLong
`func (o *COTDataPoint) UnsetTotalReportableLong()`

UnsetTotalReportableLong ensures that no value is present for TotalReportableLong, not even an explicit nil
### GetTotalReportableShort

`func (o *COTDataPoint) GetTotalReportableShort() int64`

GetTotalReportableShort returns the TotalReportableShort field if non-nil, zero value otherwise.

### GetTotalReportableShortOk

`func (o *COTDataPoint) GetTotalReportableShortOk() (*int64, bool)`

GetTotalReportableShortOk returns a tuple with the TotalReportableShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReportableShort

`func (o *COTDataPoint) SetTotalReportableShort(v int64)`

SetTotalReportableShort sets TotalReportableShort field to given value.

### HasTotalReportableShort

`func (o *COTDataPoint) HasTotalReportableShort() bool`

HasTotalReportableShort returns a boolean if a field has been set.

### SetTotalReportableShortNil

`func (o *COTDataPoint) SetTotalReportableShortNil(b bool)`

 SetTotalReportableShortNil sets the value for TotalReportableShort to be an explicit nil

### UnsetTotalReportableShort
`func (o *COTDataPoint) UnsetTotalReportableShort()`

UnsetTotalReportableShort ensures that no value is present for TotalReportableShort, not even an explicit nil
### GetNonreportableLong

`func (o *COTDataPoint) GetNonreportableLong() int64`

GetNonreportableLong returns the NonreportableLong field if non-nil, zero value otherwise.

### GetNonreportableLongOk

`func (o *COTDataPoint) GetNonreportableLongOk() (*int64, bool)`

GetNonreportableLongOk returns a tuple with the NonreportableLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonreportableLong

`func (o *COTDataPoint) SetNonreportableLong(v int64)`

SetNonreportableLong sets NonreportableLong field to given value.

### HasNonreportableLong

`func (o *COTDataPoint) HasNonreportableLong() bool`

HasNonreportableLong returns a boolean if a field has been set.

### SetNonreportableLongNil

`func (o *COTDataPoint) SetNonreportableLongNil(b bool)`

 SetNonreportableLongNil sets the value for NonreportableLong to be an explicit nil

### UnsetNonreportableLong
`func (o *COTDataPoint) UnsetNonreportableLong()`

UnsetNonreportableLong ensures that no value is present for NonreportableLong, not even an explicit nil
### GetNonreportableShort

`func (o *COTDataPoint) GetNonreportableShort() int64`

GetNonreportableShort returns the NonreportableShort field if non-nil, zero value otherwise.

### GetNonreportableShortOk

`func (o *COTDataPoint) GetNonreportableShortOk() (*int64, bool)`

GetNonreportableShortOk returns a tuple with the NonreportableShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonreportableShort

`func (o *COTDataPoint) SetNonreportableShort(v int64)`

SetNonreportableShort sets NonreportableShort field to given value.

### HasNonreportableShort

`func (o *COTDataPoint) HasNonreportableShort() bool`

HasNonreportableShort returns a boolean if a field has been set.

### SetNonreportableShortNil

`func (o *COTDataPoint) SetNonreportableShortNil(b bool)`

 SetNonreportableShortNil sets the value for NonreportableShort to be an explicit nil

### UnsetNonreportableShort
`func (o *COTDataPoint) UnsetNonreportableShort()`

UnsetNonreportableShort ensures that no value is present for NonreportableShort, not even an explicit nil
### GetOpenInterestZscore

`func (o *COTDataPoint) GetOpenInterestZscore() float64`

GetOpenInterestZscore returns the OpenInterestZscore field if non-nil, zero value otherwise.

### GetOpenInterestZscoreOk

`func (o *COTDataPoint) GetOpenInterestZscoreOk() (*float64, bool)`

GetOpenInterestZscoreOk returns a tuple with the OpenInterestZscore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInterestZscore

`func (o *COTDataPoint) SetOpenInterestZscore(v float64)`

SetOpenInterestZscore sets OpenInterestZscore field to given value.

### HasOpenInterestZscore

`func (o *COTDataPoint) HasOpenInterestZscore() bool`

HasOpenInterestZscore returns a boolean if a field has been set.

### SetOpenInterestZscoreNil

`func (o *COTDataPoint) SetOpenInterestZscoreNil(b bool)`

 SetOpenInterestZscoreNil sets the value for OpenInterestZscore to be an explicit nil

### UnsetOpenInterestZscore
`func (o *COTDataPoint) UnsetOpenInterestZscore()`

UnsetOpenInterestZscore ensures that no value is present for OpenInterestZscore, not even an explicit nil
### GetNoncommercialLongZscore

`func (o *COTDataPoint) GetNoncommercialLongZscore() float64`

GetNoncommercialLongZscore returns the NoncommercialLongZscore field if non-nil, zero value otherwise.

### GetNoncommercialLongZscoreOk

`func (o *COTDataPoint) GetNoncommercialLongZscoreOk() (*float64, bool)`

GetNoncommercialLongZscoreOk returns a tuple with the NoncommercialLongZscore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoncommercialLongZscore

`func (o *COTDataPoint) SetNoncommercialLongZscore(v float64)`

SetNoncommercialLongZscore sets NoncommercialLongZscore field to given value.

### HasNoncommercialLongZscore

`func (o *COTDataPoint) HasNoncommercialLongZscore() bool`

HasNoncommercialLongZscore returns a boolean if a field has been set.

### SetNoncommercialLongZscoreNil

`func (o *COTDataPoint) SetNoncommercialLongZscoreNil(b bool)`

 SetNoncommercialLongZscoreNil sets the value for NoncommercialLongZscore to be an explicit nil

### UnsetNoncommercialLongZscore
`func (o *COTDataPoint) UnsetNoncommercialLongZscore()`

UnsetNoncommercialLongZscore ensures that no value is present for NoncommercialLongZscore, not even an explicit nil
### GetNoncommercialShortZscore

`func (o *COTDataPoint) GetNoncommercialShortZscore() float64`

GetNoncommercialShortZscore returns the NoncommercialShortZscore field if non-nil, zero value otherwise.

### GetNoncommercialShortZscoreOk

`func (o *COTDataPoint) GetNoncommercialShortZscoreOk() (*float64, bool)`

GetNoncommercialShortZscoreOk returns a tuple with the NoncommercialShortZscore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoncommercialShortZscore

`func (o *COTDataPoint) SetNoncommercialShortZscore(v float64)`

SetNoncommercialShortZscore sets NoncommercialShortZscore field to given value.

### HasNoncommercialShortZscore

`func (o *COTDataPoint) HasNoncommercialShortZscore() bool`

HasNoncommercialShortZscore returns a boolean if a field has been set.

### SetNoncommercialShortZscoreNil

`func (o *COTDataPoint) SetNoncommercialShortZscoreNil(b bool)`

 SetNoncommercialShortZscoreNil sets the value for NoncommercialShortZscore to be an explicit nil

### UnsetNoncommercialShortZscore
`func (o *COTDataPoint) UnsetNoncommercialShortZscore()`

UnsetNoncommercialShortZscore ensures that no value is present for NoncommercialShortZscore, not even an explicit nil
### GetNoncommercialNetZscore

`func (o *COTDataPoint) GetNoncommercialNetZscore() float64`

GetNoncommercialNetZscore returns the NoncommercialNetZscore field if non-nil, zero value otherwise.

### GetNoncommercialNetZscoreOk

`func (o *COTDataPoint) GetNoncommercialNetZscoreOk() (*float64, bool)`

GetNoncommercialNetZscoreOk returns a tuple with the NoncommercialNetZscore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoncommercialNetZscore

`func (o *COTDataPoint) SetNoncommercialNetZscore(v float64)`

SetNoncommercialNetZscore sets NoncommercialNetZscore field to given value.

### HasNoncommercialNetZscore

`func (o *COTDataPoint) HasNoncommercialNetZscore() bool`

HasNoncommercialNetZscore returns a boolean if a field has been set.

### SetNoncommercialNetZscoreNil

`func (o *COTDataPoint) SetNoncommercialNetZscoreNil(b bool)`

 SetNoncommercialNetZscoreNil sets the value for NoncommercialNetZscore to be an explicit nil

### UnsetNoncommercialNetZscore
`func (o *COTDataPoint) UnsetNoncommercialNetZscore()`

UnsetNoncommercialNetZscore ensures that no value is present for NoncommercialNetZscore, not even an explicit nil
### GetNoncommercialSpreadZscore

`func (o *COTDataPoint) GetNoncommercialSpreadZscore() float64`

GetNoncommercialSpreadZscore returns the NoncommercialSpreadZscore field if non-nil, zero value otherwise.

### GetNoncommercialSpreadZscoreOk

`func (o *COTDataPoint) GetNoncommercialSpreadZscoreOk() (*float64, bool)`

GetNoncommercialSpreadZscoreOk returns a tuple with the NoncommercialSpreadZscore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoncommercialSpreadZscore

`func (o *COTDataPoint) SetNoncommercialSpreadZscore(v float64)`

SetNoncommercialSpreadZscore sets NoncommercialSpreadZscore field to given value.

### HasNoncommercialSpreadZscore

`func (o *COTDataPoint) HasNoncommercialSpreadZscore() bool`

HasNoncommercialSpreadZscore returns a boolean if a field has been set.

### SetNoncommercialSpreadZscoreNil

`func (o *COTDataPoint) SetNoncommercialSpreadZscoreNil(b bool)`

 SetNoncommercialSpreadZscoreNil sets the value for NoncommercialSpreadZscore to be an explicit nil

### UnsetNoncommercialSpreadZscore
`func (o *COTDataPoint) UnsetNoncommercialSpreadZscore()`

UnsetNoncommercialSpreadZscore ensures that no value is present for NoncommercialSpreadZscore, not even an explicit nil
### GetCommercialLongZscore

`func (o *COTDataPoint) GetCommercialLongZscore() float64`

GetCommercialLongZscore returns the CommercialLongZscore field if non-nil, zero value otherwise.

### GetCommercialLongZscoreOk

`func (o *COTDataPoint) GetCommercialLongZscoreOk() (*float64, bool)`

GetCommercialLongZscoreOk returns a tuple with the CommercialLongZscore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialLongZscore

`func (o *COTDataPoint) SetCommercialLongZscore(v float64)`

SetCommercialLongZscore sets CommercialLongZscore field to given value.

### HasCommercialLongZscore

`func (o *COTDataPoint) HasCommercialLongZscore() bool`

HasCommercialLongZscore returns a boolean if a field has been set.

### SetCommercialLongZscoreNil

`func (o *COTDataPoint) SetCommercialLongZscoreNil(b bool)`

 SetCommercialLongZscoreNil sets the value for CommercialLongZscore to be an explicit nil

### UnsetCommercialLongZscore
`func (o *COTDataPoint) UnsetCommercialLongZscore()`

UnsetCommercialLongZscore ensures that no value is present for CommercialLongZscore, not even an explicit nil
### GetCommercialShortZscore

`func (o *COTDataPoint) GetCommercialShortZscore() float64`

GetCommercialShortZscore returns the CommercialShortZscore field if non-nil, zero value otherwise.

### GetCommercialShortZscoreOk

`func (o *COTDataPoint) GetCommercialShortZscoreOk() (*float64, bool)`

GetCommercialShortZscoreOk returns a tuple with the CommercialShortZscore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialShortZscore

`func (o *COTDataPoint) SetCommercialShortZscore(v float64)`

SetCommercialShortZscore sets CommercialShortZscore field to given value.

### HasCommercialShortZscore

`func (o *COTDataPoint) HasCommercialShortZscore() bool`

HasCommercialShortZscore returns a boolean if a field has been set.

### SetCommercialShortZscoreNil

`func (o *COTDataPoint) SetCommercialShortZscoreNil(b bool)`

 SetCommercialShortZscoreNil sets the value for CommercialShortZscore to be an explicit nil

### UnsetCommercialShortZscore
`func (o *COTDataPoint) UnsetCommercialShortZscore()`

UnsetCommercialShortZscore ensures that no value is present for CommercialShortZscore, not even an explicit nil
### GetCommercialNetZscore

`func (o *COTDataPoint) GetCommercialNetZscore() float64`

GetCommercialNetZscore returns the CommercialNetZscore field if non-nil, zero value otherwise.

### GetCommercialNetZscoreOk

`func (o *COTDataPoint) GetCommercialNetZscoreOk() (*float64, bool)`

GetCommercialNetZscoreOk returns a tuple with the CommercialNetZscore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialNetZscore

`func (o *COTDataPoint) SetCommercialNetZscore(v float64)`

SetCommercialNetZscore sets CommercialNetZscore field to given value.

### HasCommercialNetZscore

`func (o *COTDataPoint) HasCommercialNetZscore() bool`

HasCommercialNetZscore returns a boolean if a field has been set.

### SetCommercialNetZscoreNil

`func (o *COTDataPoint) SetCommercialNetZscoreNil(b bool)`

 SetCommercialNetZscoreNil sets the value for CommercialNetZscore to be an explicit nil

### UnsetCommercialNetZscore
`func (o *COTDataPoint) UnsetCommercialNetZscore()`

UnsetCommercialNetZscore ensures that no value is present for CommercialNetZscore, not even an explicit nil
### GetTotalReportableLongZscore

`func (o *COTDataPoint) GetTotalReportableLongZscore() float64`

GetTotalReportableLongZscore returns the TotalReportableLongZscore field if non-nil, zero value otherwise.

### GetTotalReportableLongZscoreOk

`func (o *COTDataPoint) GetTotalReportableLongZscoreOk() (*float64, bool)`

GetTotalReportableLongZscoreOk returns a tuple with the TotalReportableLongZscore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReportableLongZscore

`func (o *COTDataPoint) SetTotalReportableLongZscore(v float64)`

SetTotalReportableLongZscore sets TotalReportableLongZscore field to given value.

### HasTotalReportableLongZscore

`func (o *COTDataPoint) HasTotalReportableLongZscore() bool`

HasTotalReportableLongZscore returns a boolean if a field has been set.

### SetTotalReportableLongZscoreNil

`func (o *COTDataPoint) SetTotalReportableLongZscoreNil(b bool)`

 SetTotalReportableLongZscoreNil sets the value for TotalReportableLongZscore to be an explicit nil

### UnsetTotalReportableLongZscore
`func (o *COTDataPoint) UnsetTotalReportableLongZscore()`

UnsetTotalReportableLongZscore ensures that no value is present for TotalReportableLongZscore, not even an explicit nil
### GetTotalReportableShortZscore

`func (o *COTDataPoint) GetTotalReportableShortZscore() float64`

GetTotalReportableShortZscore returns the TotalReportableShortZscore field if non-nil, zero value otherwise.

### GetTotalReportableShortZscoreOk

`func (o *COTDataPoint) GetTotalReportableShortZscoreOk() (*float64, bool)`

GetTotalReportableShortZscoreOk returns a tuple with the TotalReportableShortZscore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReportableShortZscore

`func (o *COTDataPoint) SetTotalReportableShortZscore(v float64)`

SetTotalReportableShortZscore sets TotalReportableShortZscore field to given value.

### HasTotalReportableShortZscore

`func (o *COTDataPoint) HasTotalReportableShortZscore() bool`

HasTotalReportableShortZscore returns a boolean if a field has been set.

### SetTotalReportableShortZscoreNil

`func (o *COTDataPoint) SetTotalReportableShortZscoreNil(b bool)`

 SetTotalReportableShortZscoreNil sets the value for TotalReportableShortZscore to be an explicit nil

### UnsetTotalReportableShortZscore
`func (o *COTDataPoint) UnsetTotalReportableShortZscore()`

UnsetTotalReportableShortZscore ensures that no value is present for TotalReportableShortZscore, not even an explicit nil
### GetNonreportableLongZscore

`func (o *COTDataPoint) GetNonreportableLongZscore() float64`

GetNonreportableLongZscore returns the NonreportableLongZscore field if non-nil, zero value otherwise.

### GetNonreportableLongZscoreOk

`func (o *COTDataPoint) GetNonreportableLongZscoreOk() (*float64, bool)`

GetNonreportableLongZscoreOk returns a tuple with the NonreportableLongZscore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonreportableLongZscore

`func (o *COTDataPoint) SetNonreportableLongZscore(v float64)`

SetNonreportableLongZscore sets NonreportableLongZscore field to given value.

### HasNonreportableLongZscore

`func (o *COTDataPoint) HasNonreportableLongZscore() bool`

HasNonreportableLongZscore returns a boolean if a field has been set.

### SetNonreportableLongZscoreNil

`func (o *COTDataPoint) SetNonreportableLongZscoreNil(b bool)`

 SetNonreportableLongZscoreNil sets the value for NonreportableLongZscore to be an explicit nil

### UnsetNonreportableLongZscore
`func (o *COTDataPoint) UnsetNonreportableLongZscore()`

UnsetNonreportableLongZscore ensures that no value is present for NonreportableLongZscore, not even an explicit nil
### GetNonreportableShortZscore

`func (o *COTDataPoint) GetNonreportableShortZscore() float64`

GetNonreportableShortZscore returns the NonreportableShortZscore field if non-nil, zero value otherwise.

### GetNonreportableShortZscoreOk

`func (o *COTDataPoint) GetNonreportableShortZscoreOk() (*float64, bool)`

GetNonreportableShortZscoreOk returns a tuple with the NonreportableShortZscore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonreportableShortZscore

`func (o *COTDataPoint) SetNonreportableShortZscore(v float64)`

SetNonreportableShortZscore sets NonreportableShortZscore field to given value.

### HasNonreportableShortZscore

`func (o *COTDataPoint) HasNonreportableShortZscore() bool`

HasNonreportableShortZscore returns a boolean if a field has been set.

### SetNonreportableShortZscoreNil

`func (o *COTDataPoint) SetNonreportableShortZscoreNil(b bool)`

 SetNonreportableShortZscoreNil sets the value for NonreportableShortZscore to be an explicit nil

### UnsetNonreportableShortZscore
`func (o *COTDataPoint) UnsetNonreportableShortZscore()`

UnsetNonreportableShortZscore ensures that no value is present for NonreportableShortZscore, not even an explicit nil
### GetReportDate

`func (o *COTDataPoint) GetReportDate() string`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *COTDataPoint) GetReportDateOk() (*string, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *COTDataPoint) SetReportDate(v string)`

SetReportDate sets ReportDate field to given value.

### HasReportDate

`func (o *COTDataPoint) HasReportDate() bool`

HasReportDate returns a boolean if a field has been set.

### SetReportDateNil

`func (o *COTDataPoint) SetReportDateNil(b bool)`

 SetReportDateNil sets the value for ReportDate to be an explicit nil

### UnsetReportDate
`func (o *COTDataPoint) UnsetReportDate()`

UnsetReportDate ensures that no value is present for ReportDate, not even an explicit nil
### GetCutoffDate

`func (o *COTDataPoint) GetCutoffDate() string`

GetCutoffDate returns the CutoffDate field if non-nil, zero value otherwise.

### GetCutoffDateOk

`func (o *COTDataPoint) GetCutoffDateOk() (*string, bool)`

GetCutoffDateOk returns a tuple with the CutoffDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoffDate

`func (o *COTDataPoint) SetCutoffDate(v string)`

SetCutoffDate sets CutoffDate field to given value.

### HasCutoffDate

`func (o *COTDataPoint) HasCutoffDate() bool`

HasCutoffDate returns a boolean if a field has been set.

### SetCutoffDateNil

`func (o *COTDataPoint) SetCutoffDateNil(b bool)`

 SetCutoffDateNil sets the value for CutoffDate to be an explicit nil

### UnsetCutoffDate
`func (o *COTDataPoint) UnsetCutoffDate()`

UnsetCutoffDate ensures that no value is present for CutoffDate, not even an explicit nil
### GetReleaseDate

`func (o *COTDataPoint) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *COTDataPoint) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *COTDataPoint) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *COTDataPoint) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *COTDataPoint) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *COTDataPoint) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetReleaseDatetime

`func (o *COTDataPoint) GetReleaseDatetime() string`

GetReleaseDatetime returns the ReleaseDatetime field if non-nil, zero value otherwise.

### GetReleaseDatetimeOk

`func (o *COTDataPoint) GetReleaseDatetimeOk() (*string, bool)`

GetReleaseDatetimeOk returns a tuple with the ReleaseDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDatetime

`func (o *COTDataPoint) SetReleaseDatetime(v string)`

SetReleaseDatetime sets ReleaseDatetime field to given value.

### HasReleaseDatetime

`func (o *COTDataPoint) HasReleaseDatetime() bool`

HasReleaseDatetime returns a boolean if a field has been set.

### SetReleaseDatetimeNil

`func (o *COTDataPoint) SetReleaseDatetimeNil(b bool)`

 SetReleaseDatetimeNil sets the value for ReleaseDatetime to be an explicit nil

### UnsetReleaseDatetime
`func (o *COTDataPoint) UnsetReleaseDatetime()`

UnsetReleaseDatetime ensures that no value is present for ReleaseDatetime, not even an explicit nil
### GetReleaseDateConfirmed

`func (o *COTDataPoint) GetReleaseDateConfirmed() bool`

GetReleaseDateConfirmed returns the ReleaseDateConfirmed field if non-nil, zero value otherwise.

### GetReleaseDateConfirmedOk

`func (o *COTDataPoint) GetReleaseDateConfirmedOk() (*bool, bool)`

GetReleaseDateConfirmedOk returns a tuple with the ReleaseDateConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDateConfirmed

`func (o *COTDataPoint) SetReleaseDateConfirmed(v bool)`

SetReleaseDateConfirmed sets ReleaseDateConfirmed field to given value.

### HasReleaseDateConfirmed

`func (o *COTDataPoint) HasReleaseDateConfirmed() bool`

HasReleaseDateConfirmed returns a boolean if a field has been set.

### SetReleaseDateConfirmedNil

`func (o *COTDataPoint) SetReleaseDateConfirmedNil(b bool)`

 SetReleaseDateConfirmedNil sets the value for ReleaseDateConfirmed to be an explicit nil

### UnsetReleaseDateConfirmed
`func (o *COTDataPoint) UnsetReleaseDateConfirmed()`

UnsetReleaseDateConfirmed ensures that no value is present for ReleaseDateConfirmed, not even an explicit nil
### GetReleaseTimeAssumed

`func (o *COTDataPoint) GetReleaseTimeAssumed() bool`

GetReleaseTimeAssumed returns the ReleaseTimeAssumed field if non-nil, zero value otherwise.

### GetReleaseTimeAssumedOk

`func (o *COTDataPoint) GetReleaseTimeAssumedOk() (*bool, bool)`

GetReleaseTimeAssumedOk returns a tuple with the ReleaseTimeAssumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTimeAssumed

`func (o *COTDataPoint) SetReleaseTimeAssumed(v bool)`

SetReleaseTimeAssumed sets ReleaseTimeAssumed field to given value.

### HasReleaseTimeAssumed

`func (o *COTDataPoint) HasReleaseTimeAssumed() bool`

HasReleaseTimeAssumed returns a boolean if a field has been set.

### SetReleaseTimeAssumedNil

`func (o *COTDataPoint) SetReleaseTimeAssumedNil(b bool)`

 SetReleaseTimeAssumedNil sets the value for ReleaseTimeAssumed to be an explicit nil

### UnsetReleaseTimeAssumed
`func (o *COTDataPoint) UnsetReleaseTimeAssumed()`

UnsetReleaseTimeAssumed ensures that no value is present for ReleaseTimeAssumed, not even an explicit nil
### GetReleaseSource

`func (o *COTDataPoint) GetReleaseSource() string`

GetReleaseSource returns the ReleaseSource field if non-nil, zero value otherwise.

### GetReleaseSourceOk

`func (o *COTDataPoint) GetReleaseSourceOk() (*string, bool)`

GetReleaseSourceOk returns a tuple with the ReleaseSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseSource

`func (o *COTDataPoint) SetReleaseSource(v string)`

SetReleaseSource sets ReleaseSource field to given value.

### HasReleaseSource

`func (o *COTDataPoint) HasReleaseSource() bool`

HasReleaseSource returns a boolean if a field has been set.

### SetReleaseSourceNil

`func (o *COTDataPoint) SetReleaseSourceNil(b bool)`

 SetReleaseSourceNil sets the value for ReleaseSource to be an explicit nil

### UnsetReleaseSource
`func (o *COTDataPoint) UnsetReleaseSource()`

UnsetReleaseSource ensures that no value is present for ReleaseSource, not even an explicit nil
### GetReleaseSourceUrl

`func (o *COTDataPoint) GetReleaseSourceUrl() string`

GetReleaseSourceUrl returns the ReleaseSourceUrl field if non-nil, zero value otherwise.

### GetReleaseSourceUrlOk

`func (o *COTDataPoint) GetReleaseSourceUrlOk() (*string, bool)`

GetReleaseSourceUrlOk returns a tuple with the ReleaseSourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseSourceUrl

`func (o *COTDataPoint) SetReleaseSourceUrl(v string)`

SetReleaseSourceUrl sets ReleaseSourceUrl field to given value.

### HasReleaseSourceUrl

`func (o *COTDataPoint) HasReleaseSourceUrl() bool`

HasReleaseSourceUrl returns a boolean if a field has been set.

### SetReleaseSourceUrlNil

`func (o *COTDataPoint) SetReleaseSourceUrlNil(b bool)`

 SetReleaseSourceUrlNil sets the value for ReleaseSourceUrl to be an explicit nil

### UnsetReleaseSourceUrl
`func (o *COTDataPoint) UnsetReleaseSourceUrl()`

UnsetReleaseSourceUrl ensures that no value is present for ReleaseSourceUrl, not even an explicit nil
### GetHolidayAdjustedRelease

`func (o *COTDataPoint) GetHolidayAdjustedRelease() bool`

GetHolidayAdjustedRelease returns the HolidayAdjustedRelease field if non-nil, zero value otherwise.

### GetHolidayAdjustedReleaseOk

`func (o *COTDataPoint) GetHolidayAdjustedReleaseOk() (*bool, bool)`

GetHolidayAdjustedReleaseOk returns a tuple with the HolidayAdjustedRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayAdjustedRelease

`func (o *COTDataPoint) SetHolidayAdjustedRelease(v bool)`

SetHolidayAdjustedRelease sets HolidayAdjustedRelease field to given value.

### HasHolidayAdjustedRelease

`func (o *COTDataPoint) HasHolidayAdjustedRelease() bool`

HasHolidayAdjustedRelease returns a boolean if a field has been set.

### SetHolidayAdjustedReleaseNil

`func (o *COTDataPoint) SetHolidayAdjustedReleaseNil(b bool)`

 SetHolidayAdjustedReleaseNil sets the value for HolidayAdjustedRelease to be an explicit nil

### UnsetHolidayAdjustedRelease
`func (o *COTDataPoint) UnsetHolidayAdjustedRelease()`

UnsetHolidayAdjustedRelease ensures that no value is present for HolidayAdjustedRelease, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


