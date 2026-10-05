# FactorDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**Val** | Pointer to **NullableFloat64** |  | [optional] 
**Score** | Pointer to **NullableFloat64** |  | [optional] 
**LevelScore** | Pointer to **NullableFloat64** |  | [optional] 
**ImpulseScore** | Pointer to **NullableFloat64** |  | [optional] 
**RateRepricingScore** | Pointer to **NullableFloat64** |  | [optional] 
**MacroPressureScore** | Pointer to **NullableFloat64** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**StanceContext** | Pointer to **NullableString** |  | [optional] 
**AnnouncementDatetime** | Pointer to **NullableInt64** |  | [optional] 
**CoverageRatio** | Pointer to **NullableFloat64** |  | [optional] 
**ComponentCount** | Pointer to **NullableInt64** |  | [optional] 
**PointInTimeSafe** | Pointer to **NullableBool** | True only when every input observation behind this row carries a publication timestamp captured from its release. An input dated from that series&#39; measured publication lag clears it, because the composite could not be shown to have been knowable at its own announcement_datetime. | [optional] 
**ReleaseTimeAssumed** | Pointer to **NullableBool** | Present and true when at least one input observation carries a derived publication timestamp. Omitted when every input was captured from its release. | [optional] 
**AssumedReleaseTimeInputs** | Pointer to **NullableInt64** | Number of input observations behind this row whose publication timestamp was derived. | [optional] 
**Components** | Pointer to **map[string]interface{}** |  | [optional] 
**SourceObservations** | Pointer to **map[string]interface{}** |  | [optional] 
**SourceEndpoints** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFactorDataPoint

`func NewFactorDataPoint(date string, ) *FactorDataPoint`

NewFactorDataPoint instantiates a new FactorDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactorDataPointWithDefaults

`func NewFactorDataPointWithDefaults() *FactorDataPoint`

NewFactorDataPointWithDefaults instantiates a new FactorDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *FactorDataPoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FactorDataPoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FactorDataPoint) SetDate(v string)`

SetDate sets Date field to given value.


### GetVal

`func (o *FactorDataPoint) GetVal() float64`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *FactorDataPoint) GetValOk() (*float64, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *FactorDataPoint) SetVal(v float64)`

SetVal sets Val field to given value.

### HasVal

`func (o *FactorDataPoint) HasVal() bool`

HasVal returns a boolean if a field has been set.

### SetValNil

`func (o *FactorDataPoint) SetValNil(b bool)`

 SetValNil sets the value for Val to be an explicit nil

### UnsetVal
`func (o *FactorDataPoint) UnsetVal()`

UnsetVal ensures that no value is present for Val, not even an explicit nil
### GetScore

`func (o *FactorDataPoint) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *FactorDataPoint) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *FactorDataPoint) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *FactorDataPoint) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *FactorDataPoint) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *FactorDataPoint) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetLevelScore

`func (o *FactorDataPoint) GetLevelScore() float64`

GetLevelScore returns the LevelScore field if non-nil, zero value otherwise.

### GetLevelScoreOk

`func (o *FactorDataPoint) GetLevelScoreOk() (*float64, bool)`

GetLevelScoreOk returns a tuple with the LevelScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelScore

`func (o *FactorDataPoint) SetLevelScore(v float64)`

SetLevelScore sets LevelScore field to given value.

### HasLevelScore

`func (o *FactorDataPoint) HasLevelScore() bool`

HasLevelScore returns a boolean if a field has been set.

### SetLevelScoreNil

`func (o *FactorDataPoint) SetLevelScoreNil(b bool)`

 SetLevelScoreNil sets the value for LevelScore to be an explicit nil

### UnsetLevelScore
`func (o *FactorDataPoint) UnsetLevelScore()`

UnsetLevelScore ensures that no value is present for LevelScore, not even an explicit nil
### GetImpulseScore

`func (o *FactorDataPoint) GetImpulseScore() float64`

GetImpulseScore returns the ImpulseScore field if non-nil, zero value otherwise.

### GetImpulseScoreOk

`func (o *FactorDataPoint) GetImpulseScoreOk() (*float64, bool)`

GetImpulseScoreOk returns a tuple with the ImpulseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpulseScore

`func (o *FactorDataPoint) SetImpulseScore(v float64)`

SetImpulseScore sets ImpulseScore field to given value.

### HasImpulseScore

`func (o *FactorDataPoint) HasImpulseScore() bool`

HasImpulseScore returns a boolean if a field has been set.

### SetImpulseScoreNil

`func (o *FactorDataPoint) SetImpulseScoreNil(b bool)`

 SetImpulseScoreNil sets the value for ImpulseScore to be an explicit nil

### UnsetImpulseScore
`func (o *FactorDataPoint) UnsetImpulseScore()`

UnsetImpulseScore ensures that no value is present for ImpulseScore, not even an explicit nil
### GetRateRepricingScore

`func (o *FactorDataPoint) GetRateRepricingScore() float64`

GetRateRepricingScore returns the RateRepricingScore field if non-nil, zero value otherwise.

### GetRateRepricingScoreOk

`func (o *FactorDataPoint) GetRateRepricingScoreOk() (*float64, bool)`

GetRateRepricingScoreOk returns a tuple with the RateRepricingScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateRepricingScore

`func (o *FactorDataPoint) SetRateRepricingScore(v float64)`

SetRateRepricingScore sets RateRepricingScore field to given value.

### HasRateRepricingScore

`func (o *FactorDataPoint) HasRateRepricingScore() bool`

HasRateRepricingScore returns a boolean if a field has been set.

### SetRateRepricingScoreNil

`func (o *FactorDataPoint) SetRateRepricingScoreNil(b bool)`

 SetRateRepricingScoreNil sets the value for RateRepricingScore to be an explicit nil

### UnsetRateRepricingScore
`func (o *FactorDataPoint) UnsetRateRepricingScore()`

UnsetRateRepricingScore ensures that no value is present for RateRepricingScore, not even an explicit nil
### GetMacroPressureScore

`func (o *FactorDataPoint) GetMacroPressureScore() float64`

GetMacroPressureScore returns the MacroPressureScore field if non-nil, zero value otherwise.

### GetMacroPressureScoreOk

`func (o *FactorDataPoint) GetMacroPressureScoreOk() (*float64, bool)`

GetMacroPressureScoreOk returns a tuple with the MacroPressureScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacroPressureScore

`func (o *FactorDataPoint) SetMacroPressureScore(v float64)`

SetMacroPressureScore sets MacroPressureScore field to given value.

### HasMacroPressureScore

`func (o *FactorDataPoint) HasMacroPressureScore() bool`

HasMacroPressureScore returns a boolean if a field has been set.

### SetMacroPressureScoreNil

`func (o *FactorDataPoint) SetMacroPressureScoreNil(b bool)`

 SetMacroPressureScoreNil sets the value for MacroPressureScore to be an explicit nil

### UnsetMacroPressureScore
`func (o *FactorDataPoint) UnsetMacroPressureScore()`

UnsetMacroPressureScore ensures that no value is present for MacroPressureScore, not even an explicit nil
### GetLabel

`func (o *FactorDataPoint) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FactorDataPoint) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FactorDataPoint) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FactorDataPoint) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *FactorDataPoint) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *FactorDataPoint) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetStanceContext

`func (o *FactorDataPoint) GetStanceContext() string`

GetStanceContext returns the StanceContext field if non-nil, zero value otherwise.

### GetStanceContextOk

`func (o *FactorDataPoint) GetStanceContextOk() (*string, bool)`

GetStanceContextOk returns a tuple with the StanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStanceContext

`func (o *FactorDataPoint) SetStanceContext(v string)`

SetStanceContext sets StanceContext field to given value.

### HasStanceContext

`func (o *FactorDataPoint) HasStanceContext() bool`

HasStanceContext returns a boolean if a field has been set.

### SetStanceContextNil

`func (o *FactorDataPoint) SetStanceContextNil(b bool)`

 SetStanceContextNil sets the value for StanceContext to be an explicit nil

### UnsetStanceContext
`func (o *FactorDataPoint) UnsetStanceContext()`

UnsetStanceContext ensures that no value is present for StanceContext, not even an explicit nil
### GetAnnouncementDatetime

`func (o *FactorDataPoint) GetAnnouncementDatetime() int64`

GetAnnouncementDatetime returns the AnnouncementDatetime field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeOk

`func (o *FactorDataPoint) GetAnnouncementDatetimeOk() (*int64, bool)`

GetAnnouncementDatetimeOk returns a tuple with the AnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetime

`func (o *FactorDataPoint) SetAnnouncementDatetime(v int64)`

SetAnnouncementDatetime sets AnnouncementDatetime field to given value.

### HasAnnouncementDatetime

`func (o *FactorDataPoint) HasAnnouncementDatetime() bool`

HasAnnouncementDatetime returns a boolean if a field has been set.

### SetAnnouncementDatetimeNil

`func (o *FactorDataPoint) SetAnnouncementDatetimeNil(b bool)`

 SetAnnouncementDatetimeNil sets the value for AnnouncementDatetime to be an explicit nil

### UnsetAnnouncementDatetime
`func (o *FactorDataPoint) UnsetAnnouncementDatetime()`

UnsetAnnouncementDatetime ensures that no value is present for AnnouncementDatetime, not even an explicit nil
### GetCoverageRatio

`func (o *FactorDataPoint) GetCoverageRatio() float64`

GetCoverageRatio returns the CoverageRatio field if non-nil, zero value otherwise.

### GetCoverageRatioOk

`func (o *FactorDataPoint) GetCoverageRatioOk() (*float64, bool)`

GetCoverageRatioOk returns a tuple with the CoverageRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageRatio

`func (o *FactorDataPoint) SetCoverageRatio(v float64)`

SetCoverageRatio sets CoverageRatio field to given value.

### HasCoverageRatio

`func (o *FactorDataPoint) HasCoverageRatio() bool`

HasCoverageRatio returns a boolean if a field has been set.

### SetCoverageRatioNil

`func (o *FactorDataPoint) SetCoverageRatioNil(b bool)`

 SetCoverageRatioNil sets the value for CoverageRatio to be an explicit nil

### UnsetCoverageRatio
`func (o *FactorDataPoint) UnsetCoverageRatio()`

UnsetCoverageRatio ensures that no value is present for CoverageRatio, not even an explicit nil
### GetComponentCount

`func (o *FactorDataPoint) GetComponentCount() int64`

GetComponentCount returns the ComponentCount field if non-nil, zero value otherwise.

### GetComponentCountOk

`func (o *FactorDataPoint) GetComponentCountOk() (*int64, bool)`

GetComponentCountOk returns a tuple with the ComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCount

`func (o *FactorDataPoint) SetComponentCount(v int64)`

SetComponentCount sets ComponentCount field to given value.

### HasComponentCount

`func (o *FactorDataPoint) HasComponentCount() bool`

HasComponentCount returns a boolean if a field has been set.

### SetComponentCountNil

`func (o *FactorDataPoint) SetComponentCountNil(b bool)`

 SetComponentCountNil sets the value for ComponentCount to be an explicit nil

### UnsetComponentCount
`func (o *FactorDataPoint) UnsetComponentCount()`

UnsetComponentCount ensures that no value is present for ComponentCount, not even an explicit nil
### GetPointInTimeSafe

`func (o *FactorDataPoint) GetPointInTimeSafe() bool`

GetPointInTimeSafe returns the PointInTimeSafe field if non-nil, zero value otherwise.

### GetPointInTimeSafeOk

`func (o *FactorDataPoint) GetPointInTimeSafeOk() (*bool, bool)`

GetPointInTimeSafeOk returns a tuple with the PointInTimeSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeSafe

`func (o *FactorDataPoint) SetPointInTimeSafe(v bool)`

SetPointInTimeSafe sets PointInTimeSafe field to given value.

### HasPointInTimeSafe

`func (o *FactorDataPoint) HasPointInTimeSafe() bool`

HasPointInTimeSafe returns a boolean if a field has been set.

### SetPointInTimeSafeNil

`func (o *FactorDataPoint) SetPointInTimeSafeNil(b bool)`

 SetPointInTimeSafeNil sets the value for PointInTimeSafe to be an explicit nil

### UnsetPointInTimeSafe
`func (o *FactorDataPoint) UnsetPointInTimeSafe()`

UnsetPointInTimeSafe ensures that no value is present for PointInTimeSafe, not even an explicit nil
### GetReleaseTimeAssumed

`func (o *FactorDataPoint) GetReleaseTimeAssumed() bool`

GetReleaseTimeAssumed returns the ReleaseTimeAssumed field if non-nil, zero value otherwise.

### GetReleaseTimeAssumedOk

`func (o *FactorDataPoint) GetReleaseTimeAssumedOk() (*bool, bool)`

GetReleaseTimeAssumedOk returns a tuple with the ReleaseTimeAssumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTimeAssumed

`func (o *FactorDataPoint) SetReleaseTimeAssumed(v bool)`

SetReleaseTimeAssumed sets ReleaseTimeAssumed field to given value.

### HasReleaseTimeAssumed

`func (o *FactorDataPoint) HasReleaseTimeAssumed() bool`

HasReleaseTimeAssumed returns a boolean if a field has been set.

### SetReleaseTimeAssumedNil

`func (o *FactorDataPoint) SetReleaseTimeAssumedNil(b bool)`

 SetReleaseTimeAssumedNil sets the value for ReleaseTimeAssumed to be an explicit nil

### UnsetReleaseTimeAssumed
`func (o *FactorDataPoint) UnsetReleaseTimeAssumed()`

UnsetReleaseTimeAssumed ensures that no value is present for ReleaseTimeAssumed, not even an explicit nil
### GetAssumedReleaseTimeInputs

`func (o *FactorDataPoint) GetAssumedReleaseTimeInputs() int64`

GetAssumedReleaseTimeInputs returns the AssumedReleaseTimeInputs field if non-nil, zero value otherwise.

### GetAssumedReleaseTimeInputsOk

`func (o *FactorDataPoint) GetAssumedReleaseTimeInputsOk() (*int64, bool)`

GetAssumedReleaseTimeInputsOk returns a tuple with the AssumedReleaseTimeInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumedReleaseTimeInputs

`func (o *FactorDataPoint) SetAssumedReleaseTimeInputs(v int64)`

SetAssumedReleaseTimeInputs sets AssumedReleaseTimeInputs field to given value.

### HasAssumedReleaseTimeInputs

`func (o *FactorDataPoint) HasAssumedReleaseTimeInputs() bool`

HasAssumedReleaseTimeInputs returns a boolean if a field has been set.

### SetAssumedReleaseTimeInputsNil

`func (o *FactorDataPoint) SetAssumedReleaseTimeInputsNil(b bool)`

 SetAssumedReleaseTimeInputsNil sets the value for AssumedReleaseTimeInputs to be an explicit nil

### UnsetAssumedReleaseTimeInputs
`func (o *FactorDataPoint) UnsetAssumedReleaseTimeInputs()`

UnsetAssumedReleaseTimeInputs ensures that no value is present for AssumedReleaseTimeInputs, not even an explicit nil
### GetComponents

`func (o *FactorDataPoint) GetComponents() map[string]interface{}`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *FactorDataPoint) GetComponentsOk() (*map[string]interface{}, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *FactorDataPoint) SetComponents(v map[string]interface{})`

SetComponents sets Components field to given value.

### HasComponents

`func (o *FactorDataPoint) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *FactorDataPoint) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *FactorDataPoint) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
### GetSourceObservations

`func (o *FactorDataPoint) GetSourceObservations() map[string]interface{}`

GetSourceObservations returns the SourceObservations field if non-nil, zero value otherwise.

### GetSourceObservationsOk

`func (o *FactorDataPoint) GetSourceObservationsOk() (*map[string]interface{}, bool)`

GetSourceObservationsOk returns a tuple with the SourceObservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObservations

`func (o *FactorDataPoint) SetSourceObservations(v map[string]interface{})`

SetSourceObservations sets SourceObservations field to given value.

### HasSourceObservations

`func (o *FactorDataPoint) HasSourceObservations() bool`

HasSourceObservations returns a boolean if a field has been set.

### SetSourceObservationsNil

`func (o *FactorDataPoint) SetSourceObservationsNil(b bool)`

 SetSourceObservationsNil sets the value for SourceObservations to be an explicit nil

### UnsetSourceObservations
`func (o *FactorDataPoint) UnsetSourceObservations()`

UnsetSourceObservations ensures that no value is present for SourceObservations, not even an explicit nil
### GetSourceEndpoints

`func (o *FactorDataPoint) GetSourceEndpoints() []string`

GetSourceEndpoints returns the SourceEndpoints field if non-nil, zero value otherwise.

### GetSourceEndpointsOk

`func (o *FactorDataPoint) GetSourceEndpointsOk() (*[]string, bool)`

GetSourceEndpointsOk returns a tuple with the SourceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEndpoints

`func (o *FactorDataPoint) SetSourceEndpoints(v []string)`

SetSourceEndpoints sets SourceEndpoints field to given value.

### HasSourceEndpoints

`func (o *FactorDataPoint) HasSourceEndpoints() bool`

HasSourceEndpoints returns a boolean if a field has been set.

### SetSourceEndpointsNil

`func (o *FactorDataPoint) SetSourceEndpointsNil(b bool)`

 SetSourceEndpointsNil sets the value for SourceEndpoints to be an explicit nil

### UnsetSourceEndpoints
`func (o *FactorDataPoint) UnsetSourceEndpoints()`

UnsetSourceEndpoints ensures that no value is present for SourceEndpoints, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


