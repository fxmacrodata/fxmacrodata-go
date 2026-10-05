# CalendarResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**Timezone** | Pointer to **NullableString** |  | [optional] 
**RequestedTimezone** | Pointer to **NullableString** |  | [optional] 
**Indicator** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**EndDate** | Pointer to **NullableString** |  | [optional] 
**HistoryStartDate** | Pointer to **NullableString** |  | [optional] 
**HistoryTruncated** | Pointer to **NullableBool** |  | [optional] 
**DataQuality** | [**DataQuality**](DataQuality.md) |  | 
**Data** | [**[]CalendarReleaseRow**](CalendarReleaseRow.md) |  | 

## Methods

### NewCalendarResponse

`func NewCalendarResponse(currency string, dataQuality DataQuality, data []CalendarReleaseRow, ) *CalendarResponse`

NewCalendarResponse instantiates a new CalendarResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarResponseWithDefaults

`func NewCalendarResponseWithDefaults() *CalendarResponse`

NewCalendarResponseWithDefaults instantiates a new CalendarResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CalendarResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CalendarResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CalendarResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTimezone

`func (o *CalendarResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CalendarResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CalendarResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CalendarResponse) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *CalendarResponse) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *CalendarResponse) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetRequestedTimezone

`func (o *CalendarResponse) GetRequestedTimezone() string`

GetRequestedTimezone returns the RequestedTimezone field if non-nil, zero value otherwise.

### GetRequestedTimezoneOk

`func (o *CalendarResponse) GetRequestedTimezoneOk() (*string, bool)`

GetRequestedTimezoneOk returns a tuple with the RequestedTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTimezone

`func (o *CalendarResponse) SetRequestedTimezone(v string)`

SetRequestedTimezone sets RequestedTimezone field to given value.

### HasRequestedTimezone

`func (o *CalendarResponse) HasRequestedTimezone() bool`

HasRequestedTimezone returns a boolean if a field has been set.

### SetRequestedTimezoneNil

`func (o *CalendarResponse) SetRequestedTimezoneNil(b bool)`

 SetRequestedTimezoneNil sets the value for RequestedTimezone to be an explicit nil

### UnsetRequestedTimezone
`func (o *CalendarResponse) UnsetRequestedTimezone()`

UnsetRequestedTimezone ensures that no value is present for RequestedTimezone, not even an explicit nil
### GetIndicator

`func (o *CalendarResponse) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *CalendarResponse) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *CalendarResponse) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.

### HasIndicator

`func (o *CalendarResponse) HasIndicator() bool`

HasIndicator returns a boolean if a field has been set.

### SetIndicatorNil

`func (o *CalendarResponse) SetIndicatorNil(b bool)`

 SetIndicatorNil sets the value for Indicator to be an explicit nil

### UnsetIndicator
`func (o *CalendarResponse) UnsetIndicator()`

UnsetIndicator ensures that no value is present for Indicator, not even an explicit nil
### GetStartDate

`func (o *CalendarResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CalendarResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CalendarResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CalendarResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *CalendarResponse) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *CalendarResponse) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *CalendarResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CalendarResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CalendarResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CalendarResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *CalendarResponse) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *CalendarResponse) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetHistoryStartDate

`func (o *CalendarResponse) GetHistoryStartDate() string`

GetHistoryStartDate returns the HistoryStartDate field if non-nil, zero value otherwise.

### GetHistoryStartDateOk

`func (o *CalendarResponse) GetHistoryStartDateOk() (*string, bool)`

GetHistoryStartDateOk returns a tuple with the HistoryStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryStartDate

`func (o *CalendarResponse) SetHistoryStartDate(v string)`

SetHistoryStartDate sets HistoryStartDate field to given value.

### HasHistoryStartDate

`func (o *CalendarResponse) HasHistoryStartDate() bool`

HasHistoryStartDate returns a boolean if a field has been set.

### SetHistoryStartDateNil

`func (o *CalendarResponse) SetHistoryStartDateNil(b bool)`

 SetHistoryStartDateNil sets the value for HistoryStartDate to be an explicit nil

### UnsetHistoryStartDate
`func (o *CalendarResponse) UnsetHistoryStartDate()`

UnsetHistoryStartDate ensures that no value is present for HistoryStartDate, not even an explicit nil
### GetHistoryTruncated

`func (o *CalendarResponse) GetHistoryTruncated() bool`

GetHistoryTruncated returns the HistoryTruncated field if non-nil, zero value otherwise.

### GetHistoryTruncatedOk

`func (o *CalendarResponse) GetHistoryTruncatedOk() (*bool, bool)`

GetHistoryTruncatedOk returns a tuple with the HistoryTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryTruncated

`func (o *CalendarResponse) SetHistoryTruncated(v bool)`

SetHistoryTruncated sets HistoryTruncated field to given value.

### HasHistoryTruncated

`func (o *CalendarResponse) HasHistoryTruncated() bool`

HasHistoryTruncated returns a boolean if a field has been set.

### SetHistoryTruncatedNil

`func (o *CalendarResponse) SetHistoryTruncatedNil(b bool)`

 SetHistoryTruncatedNil sets the value for HistoryTruncated to be an explicit nil

### UnsetHistoryTruncated
`func (o *CalendarResponse) UnsetHistoryTruncated()`

UnsetHistoryTruncated ensures that no value is present for HistoryTruncated, not even an explicit nil
### GetDataQuality

`func (o *CalendarResponse) GetDataQuality() DataQuality`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *CalendarResponse) GetDataQualityOk() (*DataQuality, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *CalendarResponse) SetDataQuality(v DataQuality)`

SetDataQuality sets DataQuality field to given value.


### GetData

`func (o *CalendarResponse) GetData() []CalendarReleaseRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CalendarResponse) GetDataOk() (*[]CalendarReleaseRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CalendarResponse) SetData(v []CalendarReleaseRow)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


