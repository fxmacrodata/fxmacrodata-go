# ResearchSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**Indicator** | **string** |  | 
**MaxAgeDays** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewResearchSeries

`func NewResearchSeries(currency string, indicator string, ) *ResearchSeries`

NewResearchSeries instantiates a new ResearchSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResearchSeriesWithDefaults

`func NewResearchSeriesWithDefaults() *ResearchSeries`

NewResearchSeriesWithDefaults instantiates a new ResearchSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *ResearchSeries) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ResearchSeries) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ResearchSeries) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetIndicator

`func (o *ResearchSeries) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *ResearchSeries) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *ResearchSeries) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.


### GetMaxAgeDays

`func (o *ResearchSeries) GetMaxAgeDays() int64`

GetMaxAgeDays returns the MaxAgeDays field if non-nil, zero value otherwise.

### GetMaxAgeDaysOk

`func (o *ResearchSeries) GetMaxAgeDaysOk() (*int64, bool)`

GetMaxAgeDaysOk returns a tuple with the MaxAgeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAgeDays

`func (o *ResearchSeries) SetMaxAgeDays(v int64)`

SetMaxAgeDays sets MaxAgeDays field to given value.

### HasMaxAgeDays

`func (o *ResearchSeries) HasMaxAgeDays() bool`

HasMaxAgeDays returns a boolean if a field has been set.

### SetMaxAgeDaysNil

`func (o *ResearchSeries) SetMaxAgeDaysNil(b bool)`

 SetMaxAgeDaysNil sets the value for MaxAgeDays to be an explicit nil

### UnsetMaxAgeDays
`func (o *ResearchSeries) UnsetMaxAgeDays()`

UnsetMaxAgeDays ensures that no value is present for MaxAgeDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


