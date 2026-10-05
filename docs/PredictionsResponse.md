# PredictionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **NullableString** |  | [optional] 
**EventCompatibleOnly** | Pointer to **bool** |  | [optional] [default to false]
**Currency** | **string** |  | 
**Indicator** | Pointer to **NullableString** |  | [optional] 
**Filters** | Pointer to [**SeriesFilters**](SeriesFilters.md) |  | [optional] 
**SelectedSeriesId** | Pointer to **NullableString** |  | [optional] 
**SelectedSeries** | Pointer to [**SelectedSeries**](SelectedSeries.md) |  | [optional] 
**SupportedOptions** | Pointer to **map[string][]string** |  | [optional] 
**PredictionType** | Pointer to **NullableString** |  | [optional] 
**PredictionClass** | Pointer to **NullableString** |  | [optional] 
**PredictionSource** | Pointer to **NullableString** |  | [optional] 
**PreReleaseOnly** | Pointer to **bool** |  | [optional] [default to true]
**StartDate** | Pointer to **NullableString** |  | [optional] 
**EndDate** | Pointer to **NullableString** |  | [optional] 
**NextCursor** | Pointer to **NullableString** | Exclusive &#x60;before_date&#x60; cursor for the next page. Prefer this over deep offset pagination to keep reads bounded. | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] [default to false]
**Count** | Pointer to **int64** | Number of announcement groups returned in &#x60;&#x60;data&#x60;&#x60;. | [optional] [default to 0]
**PredictionCount** | Pointer to **int64** | Total number of per-source predictions across all groups. | [optional] [default to 0]
**DataQuality** | [**DataQuality**](DataQuality.md) |  | 
**Availability** | Pointer to **map[string]interface{}** | Why a filter returned what it returned. &#x60;&#x60;classes_with_sources&#x60;&#x60; and &#x60;&#x60;types_with_sources&#x60;&#x60; list what any caller could ever get for this pair; &#x60;&#x60;reason&#x60;&#x60; explains a requested filter that came back empty. &#x60;&#x60;no_source_for_pair&#x60;&#x60; means no publisher provides it and no plan adds it. | [optional] 
**OfficialForecasts** | Pointer to **[]map[string]interface{}** | Official forecasts and surveys for this indicator in the publisher&#39;s own period and unit (quarterly averages, annual rates, survey horizons), newest publication first, every vintage. Not merged into &#x60;&#x60;data&#x60;&#x60; because their period basis differs from the monthly announcement; &#x60;&#x60;target_period&#x60;&#x60;, &#x60;&#x60;period_basis&#x60;&#x60; and &#x60;&#x60;unit&#x60;&#x60; say exactly what each number is. Same &#x60;&#x60;prediction_class&#x60;&#x60; vocabulary as &#x60;&#x60;data&#x60;&#x60;; filtered by &#x60;&#x60;prediction_class&#x60;&#x60;, &#x60;&#x60;prediction_source&#x60;&#x60;, &#x60;&#x60;start_date&#x60;&#x60; and &#x60;&#x60;end_date&#x60;&#x60; (on the publication date). | [optional] 
**OfficialForecastCount** | Pointer to **int64** | Number of rows in &#x60;&#x60;official_forecasts&#x60;&#x60;. | [optional] [default to 0]
**OfficialForecastCalendar** | Pointer to **[]map[string]interface{}** | Confirmed official forecast publication times for this currency over the next 366 days. These are publisher releases, separate from the indicator&#39;s announcement times. | [optional] 
**OfficialForecastsTruncated** | Pointer to **bool** | True when more official forecast vintages exist than the &#x60;&#x60;official_limit&#x60;&#x60; returned; narrow &#x60;&#x60;start_date&#x60;&#x60;/&#x60;&#x60;end_date&#x60;&#x60; or &#x60;&#x60;prediction_source&#x60;&#x60; to see the rest. | [optional] [default to false]
**Data** | [**[]AnnouncementPredictions**](AnnouncementPredictions.md) |  | 

## Methods

### NewPredictionsResponse

`func NewPredictionsResponse(currency string, dataQuality DataQuality, data []AnnouncementPredictions, ) *PredictionsResponse`

NewPredictionsResponse instantiates a new PredictionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredictionsResponseWithDefaults

`func NewPredictionsResponseWithDefaults() *PredictionsResponse`

NewPredictionsResponseWithDefaults instantiates a new PredictionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *PredictionsResponse) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *PredictionsResponse) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *PredictionsResponse) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *PredictionsResponse) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *PredictionsResponse) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *PredictionsResponse) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetEventCompatibleOnly

`func (o *PredictionsResponse) GetEventCompatibleOnly() bool`

GetEventCompatibleOnly returns the EventCompatibleOnly field if non-nil, zero value otherwise.

### GetEventCompatibleOnlyOk

`func (o *PredictionsResponse) GetEventCompatibleOnlyOk() (*bool, bool)`

GetEventCompatibleOnlyOk returns a tuple with the EventCompatibleOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCompatibleOnly

`func (o *PredictionsResponse) SetEventCompatibleOnly(v bool)`

SetEventCompatibleOnly sets EventCompatibleOnly field to given value.

### HasEventCompatibleOnly

`func (o *PredictionsResponse) HasEventCompatibleOnly() bool`

HasEventCompatibleOnly returns a boolean if a field has been set.

### GetCurrency

`func (o *PredictionsResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PredictionsResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PredictionsResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetIndicator

`func (o *PredictionsResponse) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *PredictionsResponse) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *PredictionsResponse) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.

### HasIndicator

`func (o *PredictionsResponse) HasIndicator() bool`

HasIndicator returns a boolean if a field has been set.

### SetIndicatorNil

`func (o *PredictionsResponse) SetIndicatorNil(b bool)`

 SetIndicatorNil sets the value for Indicator to be an explicit nil

### UnsetIndicator
`func (o *PredictionsResponse) UnsetIndicator()`

UnsetIndicator ensures that no value is present for Indicator, not even an explicit nil
### GetFilters

`func (o *PredictionsResponse) GetFilters() SeriesFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PredictionsResponse) GetFiltersOk() (*SeriesFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PredictionsResponse) SetFilters(v SeriesFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PredictionsResponse) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSelectedSeriesId

`func (o *PredictionsResponse) GetSelectedSeriesId() string`

GetSelectedSeriesId returns the SelectedSeriesId field if non-nil, zero value otherwise.

### GetSelectedSeriesIdOk

`func (o *PredictionsResponse) GetSelectedSeriesIdOk() (*string, bool)`

GetSelectedSeriesIdOk returns a tuple with the SelectedSeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSeriesId

`func (o *PredictionsResponse) SetSelectedSeriesId(v string)`

SetSelectedSeriesId sets SelectedSeriesId field to given value.

### HasSelectedSeriesId

`func (o *PredictionsResponse) HasSelectedSeriesId() bool`

HasSelectedSeriesId returns a boolean if a field has been set.

### SetSelectedSeriesIdNil

`func (o *PredictionsResponse) SetSelectedSeriesIdNil(b bool)`

 SetSelectedSeriesIdNil sets the value for SelectedSeriesId to be an explicit nil

### UnsetSelectedSeriesId
`func (o *PredictionsResponse) UnsetSelectedSeriesId()`

UnsetSelectedSeriesId ensures that no value is present for SelectedSeriesId, not even an explicit nil
### GetSelectedSeries

`func (o *PredictionsResponse) GetSelectedSeries() SelectedSeries`

GetSelectedSeries returns the SelectedSeries field if non-nil, zero value otherwise.

### GetSelectedSeriesOk

`func (o *PredictionsResponse) GetSelectedSeriesOk() (*SelectedSeries, bool)`

GetSelectedSeriesOk returns a tuple with the SelectedSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSeries

`func (o *PredictionsResponse) SetSelectedSeries(v SelectedSeries)`

SetSelectedSeries sets SelectedSeries field to given value.

### HasSelectedSeries

`func (o *PredictionsResponse) HasSelectedSeries() bool`

HasSelectedSeries returns a boolean if a field has been set.

### GetSupportedOptions

`func (o *PredictionsResponse) GetSupportedOptions() map[string][]string`

GetSupportedOptions returns the SupportedOptions field if non-nil, zero value otherwise.

### GetSupportedOptionsOk

`func (o *PredictionsResponse) GetSupportedOptionsOk() (*map[string][]string, bool)`

GetSupportedOptionsOk returns a tuple with the SupportedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOptions

`func (o *PredictionsResponse) SetSupportedOptions(v map[string][]string)`

SetSupportedOptions sets SupportedOptions field to given value.

### HasSupportedOptions

`func (o *PredictionsResponse) HasSupportedOptions() bool`

HasSupportedOptions returns a boolean if a field has been set.

### SetSupportedOptionsNil

`func (o *PredictionsResponse) SetSupportedOptionsNil(b bool)`

 SetSupportedOptionsNil sets the value for SupportedOptions to be an explicit nil

### UnsetSupportedOptions
`func (o *PredictionsResponse) UnsetSupportedOptions()`

UnsetSupportedOptions ensures that no value is present for SupportedOptions, not even an explicit nil
### GetPredictionType

`func (o *PredictionsResponse) GetPredictionType() string`

GetPredictionType returns the PredictionType field if non-nil, zero value otherwise.

### GetPredictionTypeOk

`func (o *PredictionsResponse) GetPredictionTypeOk() (*string, bool)`

GetPredictionTypeOk returns a tuple with the PredictionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionType

`func (o *PredictionsResponse) SetPredictionType(v string)`

SetPredictionType sets PredictionType field to given value.

### HasPredictionType

`func (o *PredictionsResponse) HasPredictionType() bool`

HasPredictionType returns a boolean if a field has been set.

### SetPredictionTypeNil

`func (o *PredictionsResponse) SetPredictionTypeNil(b bool)`

 SetPredictionTypeNil sets the value for PredictionType to be an explicit nil

### UnsetPredictionType
`func (o *PredictionsResponse) UnsetPredictionType()`

UnsetPredictionType ensures that no value is present for PredictionType, not even an explicit nil
### GetPredictionClass

`func (o *PredictionsResponse) GetPredictionClass() string`

GetPredictionClass returns the PredictionClass field if non-nil, zero value otherwise.

### GetPredictionClassOk

`func (o *PredictionsResponse) GetPredictionClassOk() (*string, bool)`

GetPredictionClassOk returns a tuple with the PredictionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionClass

`func (o *PredictionsResponse) SetPredictionClass(v string)`

SetPredictionClass sets PredictionClass field to given value.

### HasPredictionClass

`func (o *PredictionsResponse) HasPredictionClass() bool`

HasPredictionClass returns a boolean if a field has been set.

### SetPredictionClassNil

`func (o *PredictionsResponse) SetPredictionClassNil(b bool)`

 SetPredictionClassNil sets the value for PredictionClass to be an explicit nil

### UnsetPredictionClass
`func (o *PredictionsResponse) UnsetPredictionClass()`

UnsetPredictionClass ensures that no value is present for PredictionClass, not even an explicit nil
### GetPredictionSource

`func (o *PredictionsResponse) GetPredictionSource() string`

GetPredictionSource returns the PredictionSource field if non-nil, zero value otherwise.

### GetPredictionSourceOk

`func (o *PredictionsResponse) GetPredictionSourceOk() (*string, bool)`

GetPredictionSourceOk returns a tuple with the PredictionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionSource

`func (o *PredictionsResponse) SetPredictionSource(v string)`

SetPredictionSource sets PredictionSource field to given value.

### HasPredictionSource

`func (o *PredictionsResponse) HasPredictionSource() bool`

HasPredictionSource returns a boolean if a field has been set.

### SetPredictionSourceNil

`func (o *PredictionsResponse) SetPredictionSourceNil(b bool)`

 SetPredictionSourceNil sets the value for PredictionSource to be an explicit nil

### UnsetPredictionSource
`func (o *PredictionsResponse) UnsetPredictionSource()`

UnsetPredictionSource ensures that no value is present for PredictionSource, not even an explicit nil
### GetPreReleaseOnly

`func (o *PredictionsResponse) GetPreReleaseOnly() bool`

GetPreReleaseOnly returns the PreReleaseOnly field if non-nil, zero value otherwise.

### GetPreReleaseOnlyOk

`func (o *PredictionsResponse) GetPreReleaseOnlyOk() (*bool, bool)`

GetPreReleaseOnlyOk returns a tuple with the PreReleaseOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreReleaseOnly

`func (o *PredictionsResponse) SetPreReleaseOnly(v bool)`

SetPreReleaseOnly sets PreReleaseOnly field to given value.

### HasPreReleaseOnly

`func (o *PredictionsResponse) HasPreReleaseOnly() bool`

HasPreReleaseOnly returns a boolean if a field has been set.

### GetStartDate

`func (o *PredictionsResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PredictionsResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PredictionsResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PredictionsResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *PredictionsResponse) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *PredictionsResponse) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *PredictionsResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PredictionsResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PredictionsResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PredictionsResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *PredictionsResponse) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *PredictionsResponse) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetNextCursor

`func (o *PredictionsResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PredictionsResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PredictionsResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *PredictionsResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *PredictionsResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *PredictionsResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetHasMore

`func (o *PredictionsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PredictionsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PredictionsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *PredictionsResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetCount

`func (o *PredictionsResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PredictionsResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PredictionsResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PredictionsResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPredictionCount

`func (o *PredictionsResponse) GetPredictionCount() int64`

GetPredictionCount returns the PredictionCount field if non-nil, zero value otherwise.

### GetPredictionCountOk

`func (o *PredictionsResponse) GetPredictionCountOk() (*int64, bool)`

GetPredictionCountOk returns a tuple with the PredictionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionCount

`func (o *PredictionsResponse) SetPredictionCount(v int64)`

SetPredictionCount sets PredictionCount field to given value.

### HasPredictionCount

`func (o *PredictionsResponse) HasPredictionCount() bool`

HasPredictionCount returns a boolean if a field has been set.

### GetDataQuality

`func (o *PredictionsResponse) GetDataQuality() DataQuality`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *PredictionsResponse) GetDataQualityOk() (*DataQuality, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *PredictionsResponse) SetDataQuality(v DataQuality)`

SetDataQuality sets DataQuality field to given value.


### GetAvailability

`func (o *PredictionsResponse) GetAvailability() map[string]interface{}`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *PredictionsResponse) GetAvailabilityOk() (*map[string]interface{}, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *PredictionsResponse) SetAvailability(v map[string]interface{})`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *PredictionsResponse) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *PredictionsResponse) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *PredictionsResponse) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetOfficialForecasts

`func (o *PredictionsResponse) GetOfficialForecasts() []map[string]interface{}`

GetOfficialForecasts returns the OfficialForecasts field if non-nil, zero value otherwise.

### GetOfficialForecastsOk

`func (o *PredictionsResponse) GetOfficialForecastsOk() (*[]map[string]interface{}, bool)`

GetOfficialForecastsOk returns a tuple with the OfficialForecasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialForecasts

`func (o *PredictionsResponse) SetOfficialForecasts(v []map[string]interface{})`

SetOfficialForecasts sets OfficialForecasts field to given value.

### HasOfficialForecasts

`func (o *PredictionsResponse) HasOfficialForecasts() bool`

HasOfficialForecasts returns a boolean if a field has been set.

### GetOfficialForecastCount

`func (o *PredictionsResponse) GetOfficialForecastCount() int64`

GetOfficialForecastCount returns the OfficialForecastCount field if non-nil, zero value otherwise.

### GetOfficialForecastCountOk

`func (o *PredictionsResponse) GetOfficialForecastCountOk() (*int64, bool)`

GetOfficialForecastCountOk returns a tuple with the OfficialForecastCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialForecastCount

`func (o *PredictionsResponse) SetOfficialForecastCount(v int64)`

SetOfficialForecastCount sets OfficialForecastCount field to given value.

### HasOfficialForecastCount

`func (o *PredictionsResponse) HasOfficialForecastCount() bool`

HasOfficialForecastCount returns a boolean if a field has been set.

### GetOfficialForecastCalendar

`func (o *PredictionsResponse) GetOfficialForecastCalendar() []map[string]interface{}`

GetOfficialForecastCalendar returns the OfficialForecastCalendar field if non-nil, zero value otherwise.

### GetOfficialForecastCalendarOk

`func (o *PredictionsResponse) GetOfficialForecastCalendarOk() (*[]map[string]interface{}, bool)`

GetOfficialForecastCalendarOk returns a tuple with the OfficialForecastCalendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialForecastCalendar

`func (o *PredictionsResponse) SetOfficialForecastCalendar(v []map[string]interface{})`

SetOfficialForecastCalendar sets OfficialForecastCalendar field to given value.

### HasOfficialForecastCalendar

`func (o *PredictionsResponse) HasOfficialForecastCalendar() bool`

HasOfficialForecastCalendar returns a boolean if a field has been set.

### GetOfficialForecastsTruncated

`func (o *PredictionsResponse) GetOfficialForecastsTruncated() bool`

GetOfficialForecastsTruncated returns the OfficialForecastsTruncated field if non-nil, zero value otherwise.

### GetOfficialForecastsTruncatedOk

`func (o *PredictionsResponse) GetOfficialForecastsTruncatedOk() (*bool, bool)`

GetOfficialForecastsTruncatedOk returns a tuple with the OfficialForecastsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialForecastsTruncated

`func (o *PredictionsResponse) SetOfficialForecastsTruncated(v bool)`

SetOfficialForecastsTruncated sets OfficialForecastsTruncated field to given value.

### HasOfficialForecastsTruncated

`func (o *PredictionsResponse) HasOfficialForecastsTruncated() bool`

HasOfficialForecastsTruncated returns a boolean if a field has been set.

### GetData

`func (o *PredictionsResponse) GetData() []AnnouncementPredictions`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PredictionsResponse) GetDataOk() (*[]AnnouncementPredictions, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PredictionsResponse) SetData(v []AnnouncementPredictions)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


