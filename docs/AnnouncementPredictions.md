# AnnouncementPredictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnouncementId** | **string** | Stable identifier of the announcement these predictions target, in the form &#x60;&#x60;{currency}_{indicator}_{date}&#x60;&#x60;. | 
**ObservationId** | Pointer to **NullableString** | Variant-safe identifier for the selected series observation. Use this with &#x60;selected_series_id&#x60; when query selectors choose a non-default series. | [optional] 
**SelectedSeriesId** | Pointer to **NullableString** | Stable identifier of the selected announcement series. | [optional] 
**Currency** | **string** |  | 
**Indicator** | **string** |  | 
**Date** | **string** | Reference period date of the forecast, &#x60;&#x60;YYYY-MM-DD&#x60;&#x60;. | 
**AnnouncementDatetime** | **int64** | Release timestamp of the forecasted announcement, Unix epoch seconds (UTC). | 
**AnnouncementDatetimeLocal** | Pointer to **NullableString** | ISO 8601 release time with timezone offset, rendered in the currency&#39;s market timezone. | [optional] 
**AnnouncementTiming** | Pointer to **NullableString** | Whether the forecasted announcement is still upcoming (&#x60;future&#x60;) or already released (&#x60;past&#x60;). | [optional] 
**Predictions** | Pointer to [**[]PredictionItem**](PredictionItem.md) | List of forecasts for this announcement, one entry per prediction source.  Sorted by &#x60;&#x60;prediction_source&#x60;&#x60; for deterministic output. | [optional] 

## Methods

### NewAnnouncementPredictions

`func NewAnnouncementPredictions(announcementId string, currency string, indicator string, date string, announcementDatetime int64, ) *AnnouncementPredictions`

NewAnnouncementPredictions instantiates a new AnnouncementPredictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementPredictionsWithDefaults

`func NewAnnouncementPredictionsWithDefaults() *AnnouncementPredictions`

NewAnnouncementPredictionsWithDefaults instantiates a new AnnouncementPredictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnouncementId

`func (o *AnnouncementPredictions) GetAnnouncementId() string`

GetAnnouncementId returns the AnnouncementId field if non-nil, zero value otherwise.

### GetAnnouncementIdOk

`func (o *AnnouncementPredictions) GetAnnouncementIdOk() (*string, bool)`

GetAnnouncementIdOk returns a tuple with the AnnouncementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementId

`func (o *AnnouncementPredictions) SetAnnouncementId(v string)`

SetAnnouncementId sets AnnouncementId field to given value.


### GetObservationId

`func (o *AnnouncementPredictions) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *AnnouncementPredictions) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *AnnouncementPredictions) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *AnnouncementPredictions) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### SetObservationIdNil

`func (o *AnnouncementPredictions) SetObservationIdNil(b bool)`

 SetObservationIdNil sets the value for ObservationId to be an explicit nil

### UnsetObservationId
`func (o *AnnouncementPredictions) UnsetObservationId()`

UnsetObservationId ensures that no value is present for ObservationId, not even an explicit nil
### GetSelectedSeriesId

`func (o *AnnouncementPredictions) GetSelectedSeriesId() string`

GetSelectedSeriesId returns the SelectedSeriesId field if non-nil, zero value otherwise.

### GetSelectedSeriesIdOk

`func (o *AnnouncementPredictions) GetSelectedSeriesIdOk() (*string, bool)`

GetSelectedSeriesIdOk returns a tuple with the SelectedSeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSeriesId

`func (o *AnnouncementPredictions) SetSelectedSeriesId(v string)`

SetSelectedSeriesId sets SelectedSeriesId field to given value.

### HasSelectedSeriesId

`func (o *AnnouncementPredictions) HasSelectedSeriesId() bool`

HasSelectedSeriesId returns a boolean if a field has been set.

### SetSelectedSeriesIdNil

`func (o *AnnouncementPredictions) SetSelectedSeriesIdNil(b bool)`

 SetSelectedSeriesIdNil sets the value for SelectedSeriesId to be an explicit nil

### UnsetSelectedSeriesId
`func (o *AnnouncementPredictions) UnsetSelectedSeriesId()`

UnsetSelectedSeriesId ensures that no value is present for SelectedSeriesId, not even an explicit nil
### GetCurrency

`func (o *AnnouncementPredictions) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AnnouncementPredictions) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AnnouncementPredictions) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetIndicator

`func (o *AnnouncementPredictions) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *AnnouncementPredictions) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *AnnouncementPredictions) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.


### GetDate

`func (o *AnnouncementPredictions) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AnnouncementPredictions) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AnnouncementPredictions) SetDate(v string)`

SetDate sets Date field to given value.


### GetAnnouncementDatetime

`func (o *AnnouncementPredictions) GetAnnouncementDatetime() int64`

GetAnnouncementDatetime returns the AnnouncementDatetime field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeOk

`func (o *AnnouncementPredictions) GetAnnouncementDatetimeOk() (*int64, bool)`

GetAnnouncementDatetimeOk returns a tuple with the AnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetime

`func (o *AnnouncementPredictions) SetAnnouncementDatetime(v int64)`

SetAnnouncementDatetime sets AnnouncementDatetime field to given value.


### GetAnnouncementDatetimeLocal

`func (o *AnnouncementPredictions) GetAnnouncementDatetimeLocal() string`

GetAnnouncementDatetimeLocal returns the AnnouncementDatetimeLocal field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeLocalOk

`func (o *AnnouncementPredictions) GetAnnouncementDatetimeLocalOk() (*string, bool)`

GetAnnouncementDatetimeLocalOk returns a tuple with the AnnouncementDatetimeLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetimeLocal

`func (o *AnnouncementPredictions) SetAnnouncementDatetimeLocal(v string)`

SetAnnouncementDatetimeLocal sets AnnouncementDatetimeLocal field to given value.

### HasAnnouncementDatetimeLocal

`func (o *AnnouncementPredictions) HasAnnouncementDatetimeLocal() bool`

HasAnnouncementDatetimeLocal returns a boolean if a field has been set.

### SetAnnouncementDatetimeLocalNil

`func (o *AnnouncementPredictions) SetAnnouncementDatetimeLocalNil(b bool)`

 SetAnnouncementDatetimeLocalNil sets the value for AnnouncementDatetimeLocal to be an explicit nil

### UnsetAnnouncementDatetimeLocal
`func (o *AnnouncementPredictions) UnsetAnnouncementDatetimeLocal()`

UnsetAnnouncementDatetimeLocal ensures that no value is present for AnnouncementDatetimeLocal, not even an explicit nil
### GetAnnouncementTiming

`func (o *AnnouncementPredictions) GetAnnouncementTiming() string`

GetAnnouncementTiming returns the AnnouncementTiming field if non-nil, zero value otherwise.

### GetAnnouncementTimingOk

`func (o *AnnouncementPredictions) GetAnnouncementTimingOk() (*string, bool)`

GetAnnouncementTimingOk returns a tuple with the AnnouncementTiming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementTiming

`func (o *AnnouncementPredictions) SetAnnouncementTiming(v string)`

SetAnnouncementTiming sets AnnouncementTiming field to given value.

### HasAnnouncementTiming

`func (o *AnnouncementPredictions) HasAnnouncementTiming() bool`

HasAnnouncementTiming returns a boolean if a field has been set.

### SetAnnouncementTimingNil

`func (o *AnnouncementPredictions) SetAnnouncementTimingNil(b bool)`

 SetAnnouncementTimingNil sets the value for AnnouncementTiming to be an explicit nil

### UnsetAnnouncementTiming
`func (o *AnnouncementPredictions) UnsetAnnouncementTiming()`

UnsetAnnouncementTiming ensures that no value is present for AnnouncementTiming, not even an explicit nil
### GetPredictions

`func (o *AnnouncementPredictions) GetPredictions() []PredictionItem`

GetPredictions returns the Predictions field if non-nil, zero value otherwise.

### GetPredictionsOk

`func (o *AnnouncementPredictions) GetPredictionsOk() (*[]PredictionItem, bool)`

GetPredictionsOk returns a tuple with the Predictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictions

`func (o *AnnouncementPredictions) SetPredictions(v []PredictionItem)`

SetPredictions sets Predictions field to given value.

### HasPredictions

`func (o *AnnouncementPredictions) HasPredictions() bool`

HasPredictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


