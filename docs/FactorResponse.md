# FactorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**Factor** | **string** |  | 
**Methodology** | **string** |  | 
**AsOf** | Pointer to **NullableString** |  | [optional] 
**Score** | Pointer to **NullableFloat64** |  | [optional] 
**LevelScore** | Pointer to **NullableFloat64** |  | [optional] 
**ImpulseScore** | Pointer to **NullableFloat64** |  | [optional] 
**RateRepricingScore** | Pointer to **NullableFloat64** |  | [optional] 
**MacroPressureScore** | Pointer to **NullableFloat64** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**StanceContext** | Pointer to **NullableString** |  | [optional] 
**LatestAvailableDate** | Pointer to **NullableString** |  | [optional] 
**LastUpdated** | Pointer to **NullableString** |  | [optional] 
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**DataQuality** | [**DataQuality**](DataQuality.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Data** | [**[]FactorDataPoint**](FactorDataPoint.md) |  | 

## Methods

### NewFactorResponse

`func NewFactorResponse(currency string, factor string, methodology string, startDate string, endDate string, dataQuality DataQuality, pagination Pagination, data []FactorDataPoint, ) *FactorResponse`

NewFactorResponse instantiates a new FactorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactorResponseWithDefaults

`func NewFactorResponseWithDefaults() *FactorResponse`

NewFactorResponseWithDefaults instantiates a new FactorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *FactorResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FactorResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FactorResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetFactor

`func (o *FactorResponse) GetFactor() string`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *FactorResponse) GetFactorOk() (*string, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *FactorResponse) SetFactor(v string)`

SetFactor sets Factor field to given value.


### GetMethodology

`func (o *FactorResponse) GetMethodology() string`

GetMethodology returns the Methodology field if non-nil, zero value otherwise.

### GetMethodologyOk

`func (o *FactorResponse) GetMethodologyOk() (*string, bool)`

GetMethodologyOk returns a tuple with the Methodology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodology

`func (o *FactorResponse) SetMethodology(v string)`

SetMethodology sets Methodology field to given value.


### GetAsOf

`func (o *FactorResponse) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *FactorResponse) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *FactorResponse) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *FactorResponse) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *FactorResponse) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *FactorResponse) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetScore

`func (o *FactorResponse) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *FactorResponse) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *FactorResponse) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *FactorResponse) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *FactorResponse) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *FactorResponse) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetLevelScore

`func (o *FactorResponse) GetLevelScore() float64`

GetLevelScore returns the LevelScore field if non-nil, zero value otherwise.

### GetLevelScoreOk

`func (o *FactorResponse) GetLevelScoreOk() (*float64, bool)`

GetLevelScoreOk returns a tuple with the LevelScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelScore

`func (o *FactorResponse) SetLevelScore(v float64)`

SetLevelScore sets LevelScore field to given value.

### HasLevelScore

`func (o *FactorResponse) HasLevelScore() bool`

HasLevelScore returns a boolean if a field has been set.

### SetLevelScoreNil

`func (o *FactorResponse) SetLevelScoreNil(b bool)`

 SetLevelScoreNil sets the value for LevelScore to be an explicit nil

### UnsetLevelScore
`func (o *FactorResponse) UnsetLevelScore()`

UnsetLevelScore ensures that no value is present for LevelScore, not even an explicit nil
### GetImpulseScore

`func (o *FactorResponse) GetImpulseScore() float64`

GetImpulseScore returns the ImpulseScore field if non-nil, zero value otherwise.

### GetImpulseScoreOk

`func (o *FactorResponse) GetImpulseScoreOk() (*float64, bool)`

GetImpulseScoreOk returns a tuple with the ImpulseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpulseScore

`func (o *FactorResponse) SetImpulseScore(v float64)`

SetImpulseScore sets ImpulseScore field to given value.

### HasImpulseScore

`func (o *FactorResponse) HasImpulseScore() bool`

HasImpulseScore returns a boolean if a field has been set.

### SetImpulseScoreNil

`func (o *FactorResponse) SetImpulseScoreNil(b bool)`

 SetImpulseScoreNil sets the value for ImpulseScore to be an explicit nil

### UnsetImpulseScore
`func (o *FactorResponse) UnsetImpulseScore()`

UnsetImpulseScore ensures that no value is present for ImpulseScore, not even an explicit nil
### GetRateRepricingScore

`func (o *FactorResponse) GetRateRepricingScore() float64`

GetRateRepricingScore returns the RateRepricingScore field if non-nil, zero value otherwise.

### GetRateRepricingScoreOk

`func (o *FactorResponse) GetRateRepricingScoreOk() (*float64, bool)`

GetRateRepricingScoreOk returns a tuple with the RateRepricingScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateRepricingScore

`func (o *FactorResponse) SetRateRepricingScore(v float64)`

SetRateRepricingScore sets RateRepricingScore field to given value.

### HasRateRepricingScore

`func (o *FactorResponse) HasRateRepricingScore() bool`

HasRateRepricingScore returns a boolean if a field has been set.

### SetRateRepricingScoreNil

`func (o *FactorResponse) SetRateRepricingScoreNil(b bool)`

 SetRateRepricingScoreNil sets the value for RateRepricingScore to be an explicit nil

### UnsetRateRepricingScore
`func (o *FactorResponse) UnsetRateRepricingScore()`

UnsetRateRepricingScore ensures that no value is present for RateRepricingScore, not even an explicit nil
### GetMacroPressureScore

`func (o *FactorResponse) GetMacroPressureScore() float64`

GetMacroPressureScore returns the MacroPressureScore field if non-nil, zero value otherwise.

### GetMacroPressureScoreOk

`func (o *FactorResponse) GetMacroPressureScoreOk() (*float64, bool)`

GetMacroPressureScoreOk returns a tuple with the MacroPressureScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacroPressureScore

`func (o *FactorResponse) SetMacroPressureScore(v float64)`

SetMacroPressureScore sets MacroPressureScore field to given value.

### HasMacroPressureScore

`func (o *FactorResponse) HasMacroPressureScore() bool`

HasMacroPressureScore returns a boolean if a field has been set.

### SetMacroPressureScoreNil

`func (o *FactorResponse) SetMacroPressureScoreNil(b bool)`

 SetMacroPressureScoreNil sets the value for MacroPressureScore to be an explicit nil

### UnsetMacroPressureScore
`func (o *FactorResponse) UnsetMacroPressureScore()`

UnsetMacroPressureScore ensures that no value is present for MacroPressureScore, not even an explicit nil
### GetLabel

`func (o *FactorResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FactorResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FactorResponse) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FactorResponse) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *FactorResponse) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *FactorResponse) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetStanceContext

`func (o *FactorResponse) GetStanceContext() string`

GetStanceContext returns the StanceContext field if non-nil, zero value otherwise.

### GetStanceContextOk

`func (o *FactorResponse) GetStanceContextOk() (*string, bool)`

GetStanceContextOk returns a tuple with the StanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStanceContext

`func (o *FactorResponse) SetStanceContext(v string)`

SetStanceContext sets StanceContext field to given value.

### HasStanceContext

`func (o *FactorResponse) HasStanceContext() bool`

HasStanceContext returns a boolean if a field has been set.

### SetStanceContextNil

`func (o *FactorResponse) SetStanceContextNil(b bool)`

 SetStanceContextNil sets the value for StanceContext to be an explicit nil

### UnsetStanceContext
`func (o *FactorResponse) UnsetStanceContext()`

UnsetStanceContext ensures that no value is present for StanceContext, not even an explicit nil
### GetLatestAvailableDate

`func (o *FactorResponse) GetLatestAvailableDate() string`

GetLatestAvailableDate returns the LatestAvailableDate field if non-nil, zero value otherwise.

### GetLatestAvailableDateOk

`func (o *FactorResponse) GetLatestAvailableDateOk() (*string, bool)`

GetLatestAvailableDateOk returns a tuple with the LatestAvailableDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAvailableDate

`func (o *FactorResponse) SetLatestAvailableDate(v string)`

SetLatestAvailableDate sets LatestAvailableDate field to given value.

### HasLatestAvailableDate

`func (o *FactorResponse) HasLatestAvailableDate() bool`

HasLatestAvailableDate returns a boolean if a field has been set.

### SetLatestAvailableDateNil

`func (o *FactorResponse) SetLatestAvailableDateNil(b bool)`

 SetLatestAvailableDateNil sets the value for LatestAvailableDate to be an explicit nil

### UnsetLatestAvailableDate
`func (o *FactorResponse) UnsetLatestAvailableDate()`

UnsetLatestAvailableDate ensures that no value is present for LatestAvailableDate, not even an explicit nil
### GetLastUpdated

`func (o *FactorResponse) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *FactorResponse) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *FactorResponse) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *FactorResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *FactorResponse) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *FactorResponse) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetStartDate

`func (o *FactorResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FactorResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FactorResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *FactorResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FactorResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FactorResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetDataQuality

`func (o *FactorResponse) GetDataQuality() DataQuality`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *FactorResponse) GetDataQualityOk() (*DataQuality, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *FactorResponse) SetDataQuality(v DataQuality)`

SetDataQuality sets DataQuality field to given value.


### GetPagination

`func (o *FactorResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *FactorResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *FactorResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *FactorResponse) GetData() []FactorDataPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FactorResponse) GetDataOk() (*[]FactorDataPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FactorResponse) SetData(v []FactorDataPoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


