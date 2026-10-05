# RiskSentimentPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **NullableString** |  | [optional] 
**Val** | Pointer to **NullableFloat64** |  | [optional] 
**Score** | Pointer to **NullableFloat64** |  | [optional] 
**Regime** | Pointer to **NullableString** |  | [optional] 
**RiskRegime** | Pointer to **NullableString** |  | [optional] 
**Sentiment** | Pointer to **NullableString** |  | [optional] 
**Components** | Pointer to **map[string]float64** |  | [optional] 
**ComponentCoverage** | Pointer to [**NullableRiskSentimentComponentCoverage**](RiskSentimentComponentCoverage.md) |  | [optional] 
**StoredComponentCount** | Pointer to **NullableInt64** |  | [optional] 
**FinancialStressScore** | Pointer to **NullableFloat64** |  | [optional] 
**CommodityBetaScore** | Pointer to **NullableFloat64** |  | [optional] 
**SafeHavenScore** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewRiskSentimentPoint

`func NewRiskSentimentPoint() *RiskSentimentPoint`

NewRiskSentimentPoint instantiates a new RiskSentimentPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskSentimentPointWithDefaults

`func NewRiskSentimentPointWithDefaults() *RiskSentimentPoint`

NewRiskSentimentPointWithDefaults instantiates a new RiskSentimentPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *RiskSentimentPoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RiskSentimentPoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RiskSentimentPoint) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *RiskSentimentPoint) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *RiskSentimentPoint) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *RiskSentimentPoint) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetVal

`func (o *RiskSentimentPoint) GetVal() float64`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *RiskSentimentPoint) GetValOk() (*float64, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *RiskSentimentPoint) SetVal(v float64)`

SetVal sets Val field to given value.

### HasVal

`func (o *RiskSentimentPoint) HasVal() bool`

HasVal returns a boolean if a field has been set.

### SetValNil

`func (o *RiskSentimentPoint) SetValNil(b bool)`

 SetValNil sets the value for Val to be an explicit nil

### UnsetVal
`func (o *RiskSentimentPoint) UnsetVal()`

UnsetVal ensures that no value is present for Val, not even an explicit nil
### GetScore

`func (o *RiskSentimentPoint) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *RiskSentimentPoint) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *RiskSentimentPoint) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *RiskSentimentPoint) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *RiskSentimentPoint) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *RiskSentimentPoint) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetRegime

`func (o *RiskSentimentPoint) GetRegime() string`

GetRegime returns the Regime field if non-nil, zero value otherwise.

### GetRegimeOk

`func (o *RiskSentimentPoint) GetRegimeOk() (*string, bool)`

GetRegimeOk returns a tuple with the Regime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegime

`func (o *RiskSentimentPoint) SetRegime(v string)`

SetRegime sets Regime field to given value.

### HasRegime

`func (o *RiskSentimentPoint) HasRegime() bool`

HasRegime returns a boolean if a field has been set.

### SetRegimeNil

`func (o *RiskSentimentPoint) SetRegimeNil(b bool)`

 SetRegimeNil sets the value for Regime to be an explicit nil

### UnsetRegime
`func (o *RiskSentimentPoint) UnsetRegime()`

UnsetRegime ensures that no value is present for Regime, not even an explicit nil
### GetRiskRegime

`func (o *RiskSentimentPoint) GetRiskRegime() string`

GetRiskRegime returns the RiskRegime field if non-nil, zero value otherwise.

### GetRiskRegimeOk

`func (o *RiskSentimentPoint) GetRiskRegimeOk() (*string, bool)`

GetRiskRegimeOk returns a tuple with the RiskRegime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskRegime

`func (o *RiskSentimentPoint) SetRiskRegime(v string)`

SetRiskRegime sets RiskRegime field to given value.

### HasRiskRegime

`func (o *RiskSentimentPoint) HasRiskRegime() bool`

HasRiskRegime returns a boolean if a field has been set.

### SetRiskRegimeNil

`func (o *RiskSentimentPoint) SetRiskRegimeNil(b bool)`

 SetRiskRegimeNil sets the value for RiskRegime to be an explicit nil

### UnsetRiskRegime
`func (o *RiskSentimentPoint) UnsetRiskRegime()`

UnsetRiskRegime ensures that no value is present for RiskRegime, not even an explicit nil
### GetSentiment

`func (o *RiskSentimentPoint) GetSentiment() string`

GetSentiment returns the Sentiment field if non-nil, zero value otherwise.

### GetSentimentOk

`func (o *RiskSentimentPoint) GetSentimentOk() (*string, bool)`

GetSentimentOk returns a tuple with the Sentiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentiment

`func (o *RiskSentimentPoint) SetSentiment(v string)`

SetSentiment sets Sentiment field to given value.

### HasSentiment

`func (o *RiskSentimentPoint) HasSentiment() bool`

HasSentiment returns a boolean if a field has been set.

### SetSentimentNil

`func (o *RiskSentimentPoint) SetSentimentNil(b bool)`

 SetSentimentNil sets the value for Sentiment to be an explicit nil

### UnsetSentiment
`func (o *RiskSentimentPoint) UnsetSentiment()`

UnsetSentiment ensures that no value is present for Sentiment, not even an explicit nil
### GetComponents

`func (o *RiskSentimentPoint) GetComponents() map[string]*float64`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *RiskSentimentPoint) GetComponentsOk() (*map[string]*float64, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *RiskSentimentPoint) SetComponents(v map[string]*float64)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *RiskSentimentPoint) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *RiskSentimentPoint) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *RiskSentimentPoint) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
### GetComponentCoverage

`func (o *RiskSentimentPoint) GetComponentCoverage() RiskSentimentComponentCoverage`

GetComponentCoverage returns the ComponentCoverage field if non-nil, zero value otherwise.

### GetComponentCoverageOk

`func (o *RiskSentimentPoint) GetComponentCoverageOk() (*RiskSentimentComponentCoverage, bool)`

GetComponentCoverageOk returns a tuple with the ComponentCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCoverage

`func (o *RiskSentimentPoint) SetComponentCoverage(v RiskSentimentComponentCoverage)`

SetComponentCoverage sets ComponentCoverage field to given value.

### HasComponentCoverage

`func (o *RiskSentimentPoint) HasComponentCoverage() bool`

HasComponentCoverage returns a boolean if a field has been set.

### SetComponentCoverageNil

`func (o *RiskSentimentPoint) SetComponentCoverageNil(b bool)`

 SetComponentCoverageNil sets the value for ComponentCoverage to be an explicit nil

### UnsetComponentCoverage
`func (o *RiskSentimentPoint) UnsetComponentCoverage()`

UnsetComponentCoverage ensures that no value is present for ComponentCoverage, not even an explicit nil
### GetStoredComponentCount

`func (o *RiskSentimentPoint) GetStoredComponentCount() int64`

GetStoredComponentCount returns the StoredComponentCount field if non-nil, zero value otherwise.

### GetStoredComponentCountOk

`func (o *RiskSentimentPoint) GetStoredComponentCountOk() (*int64, bool)`

GetStoredComponentCountOk returns a tuple with the StoredComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredComponentCount

`func (o *RiskSentimentPoint) SetStoredComponentCount(v int64)`

SetStoredComponentCount sets StoredComponentCount field to given value.

### HasStoredComponentCount

`func (o *RiskSentimentPoint) HasStoredComponentCount() bool`

HasStoredComponentCount returns a boolean if a field has been set.

### SetStoredComponentCountNil

`func (o *RiskSentimentPoint) SetStoredComponentCountNil(b bool)`

 SetStoredComponentCountNil sets the value for StoredComponentCount to be an explicit nil

### UnsetStoredComponentCount
`func (o *RiskSentimentPoint) UnsetStoredComponentCount()`

UnsetStoredComponentCount ensures that no value is present for StoredComponentCount, not even an explicit nil
### GetFinancialStressScore

`func (o *RiskSentimentPoint) GetFinancialStressScore() float64`

GetFinancialStressScore returns the FinancialStressScore field if non-nil, zero value otherwise.

### GetFinancialStressScoreOk

`func (o *RiskSentimentPoint) GetFinancialStressScoreOk() (*float64, bool)`

GetFinancialStressScoreOk returns a tuple with the FinancialStressScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialStressScore

`func (o *RiskSentimentPoint) SetFinancialStressScore(v float64)`

SetFinancialStressScore sets FinancialStressScore field to given value.

### HasFinancialStressScore

`func (o *RiskSentimentPoint) HasFinancialStressScore() bool`

HasFinancialStressScore returns a boolean if a field has been set.

### SetFinancialStressScoreNil

`func (o *RiskSentimentPoint) SetFinancialStressScoreNil(b bool)`

 SetFinancialStressScoreNil sets the value for FinancialStressScore to be an explicit nil

### UnsetFinancialStressScore
`func (o *RiskSentimentPoint) UnsetFinancialStressScore()`

UnsetFinancialStressScore ensures that no value is present for FinancialStressScore, not even an explicit nil
### GetCommodityBetaScore

`func (o *RiskSentimentPoint) GetCommodityBetaScore() float64`

GetCommodityBetaScore returns the CommodityBetaScore field if non-nil, zero value otherwise.

### GetCommodityBetaScoreOk

`func (o *RiskSentimentPoint) GetCommodityBetaScoreOk() (*float64, bool)`

GetCommodityBetaScoreOk returns a tuple with the CommodityBetaScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommodityBetaScore

`func (o *RiskSentimentPoint) SetCommodityBetaScore(v float64)`

SetCommodityBetaScore sets CommodityBetaScore field to given value.

### HasCommodityBetaScore

`func (o *RiskSentimentPoint) HasCommodityBetaScore() bool`

HasCommodityBetaScore returns a boolean if a field has been set.

### SetCommodityBetaScoreNil

`func (o *RiskSentimentPoint) SetCommodityBetaScoreNil(b bool)`

 SetCommodityBetaScoreNil sets the value for CommodityBetaScore to be an explicit nil

### UnsetCommodityBetaScore
`func (o *RiskSentimentPoint) UnsetCommodityBetaScore()`

UnsetCommodityBetaScore ensures that no value is present for CommodityBetaScore, not even an explicit nil
### GetSafeHavenScore

`func (o *RiskSentimentPoint) GetSafeHavenScore() float64`

GetSafeHavenScore returns the SafeHavenScore field if non-nil, zero value otherwise.

### GetSafeHavenScoreOk

`func (o *RiskSentimentPoint) GetSafeHavenScoreOk() (*float64, bool)`

GetSafeHavenScoreOk returns a tuple with the SafeHavenScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeHavenScore

`func (o *RiskSentimentPoint) SetSafeHavenScore(v float64)`

SetSafeHavenScore sets SafeHavenScore field to given value.

### HasSafeHavenScore

`func (o *RiskSentimentPoint) HasSafeHavenScore() bool`

HasSafeHavenScore returns a boolean if a field has been set.

### SetSafeHavenScoreNil

`func (o *RiskSentimentPoint) SetSafeHavenScoreNil(b bool)`

 SetSafeHavenScoreNil sets the value for SafeHavenScore to be an explicit nil

### UnsetSafeHavenScore
`func (o *RiskSentimentPoint) UnsetSafeHavenScore()`

UnsetSafeHavenScore ensures that no value is present for SafeHavenScore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


