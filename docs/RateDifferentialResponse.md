# RateDifferentialResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Research** | Pointer to **map[string]interface{}** |  | [optional] 
**Base** | **string** |  | 
**Quote** | **string** |  | 
**RateType** | Pointer to **string** |  | [optional] [default to "spot"]
**MeasureRequested** | **string** |  | 
**MeasureUsed** | **string** |  | 
**BaseIndicator** | Pointer to **NullableString** |  | [optional] 
**QuoteIndicator** | Pointer to **NullableString** |  | [optional] 
**CurveFamily** | Pointer to **NullableString** |  | [optional] 
**StartTenor** | Pointer to **NullableString** |  | [optional] 
**EndTenor** | Pointer to **NullableString** |  | [optional] 
**ForwardLabel** | Pointer to **NullableString** |  | [optional] 
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**MatchedPoints** | **int64** |  | 
**Unit** | **string** |  | 
**LatestSpread** | Pointer to **NullableFloat64** |  | [optional] 
**LatestSpreadBps** | Pointer to **NullableFloat64** |  | [optional] 
**LatestDifferential** | Pointer to **NullableFloat64** |  | [optional] 
**LatestDifferentialBps** | Pointer to **NullableFloat64** |  | [optional] 
**BaseLatest** | Pointer to **NullableFloat64** |  | [optional] 
**QuoteLatest** | Pointer to **NullableFloat64** |  | [optional] 
**Sources** | [**RateDifferentialSources**](RateDifferentialSources.md) |  | 
**OfficialForwardSourceSupport** | Pointer to [**OfficialForwardSourceSupport**](OfficialForwardSourceSupport.md) |  | [optional] 
**DataQuality** | [**DataQuality**](DataQuality.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Data** | [**[]RateDifferentialPoint**](RateDifferentialPoint.md) |  | 

## Methods

### NewRateDifferentialResponse

`func NewRateDifferentialResponse(base string, quote string, measureRequested string, measureUsed string, startDate string, endDate string, matchedPoints int64, unit string, sources RateDifferentialSources, dataQuality DataQuality, pagination Pagination, data []RateDifferentialPoint, ) *RateDifferentialResponse`

NewRateDifferentialResponse instantiates a new RateDifferentialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateDifferentialResponseWithDefaults

`func NewRateDifferentialResponseWithDefaults() *RateDifferentialResponse`

NewRateDifferentialResponseWithDefaults instantiates a new RateDifferentialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResearch

`func (o *RateDifferentialResponse) GetResearch() map[string]interface{}`

GetResearch returns the Research field if non-nil, zero value otherwise.

### GetResearchOk

`func (o *RateDifferentialResponse) GetResearchOk() (*map[string]interface{}, bool)`

GetResearchOk returns a tuple with the Research field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearch

`func (o *RateDifferentialResponse) SetResearch(v map[string]interface{})`

SetResearch sets Research field to given value.

### HasResearch

`func (o *RateDifferentialResponse) HasResearch() bool`

HasResearch returns a boolean if a field has been set.

### SetResearchNil

`func (o *RateDifferentialResponse) SetResearchNil(b bool)`

 SetResearchNil sets the value for Research to be an explicit nil

### UnsetResearch
`func (o *RateDifferentialResponse) UnsetResearch()`

UnsetResearch ensures that no value is present for Research, not even an explicit nil
### GetBase

`func (o *RateDifferentialResponse) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *RateDifferentialResponse) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *RateDifferentialResponse) SetBase(v string)`

SetBase sets Base field to given value.


### GetQuote

`func (o *RateDifferentialResponse) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *RateDifferentialResponse) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *RateDifferentialResponse) SetQuote(v string)`

SetQuote sets Quote field to given value.


### GetRateType

`func (o *RateDifferentialResponse) GetRateType() string`

GetRateType returns the RateType field if non-nil, zero value otherwise.

### GetRateTypeOk

`func (o *RateDifferentialResponse) GetRateTypeOk() (*string, bool)`

GetRateTypeOk returns a tuple with the RateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateType

`func (o *RateDifferentialResponse) SetRateType(v string)`

SetRateType sets RateType field to given value.

### HasRateType

`func (o *RateDifferentialResponse) HasRateType() bool`

HasRateType returns a boolean if a field has been set.

### GetMeasureRequested

`func (o *RateDifferentialResponse) GetMeasureRequested() string`

GetMeasureRequested returns the MeasureRequested field if non-nil, zero value otherwise.

### GetMeasureRequestedOk

`func (o *RateDifferentialResponse) GetMeasureRequestedOk() (*string, bool)`

GetMeasureRequestedOk returns a tuple with the MeasureRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureRequested

`func (o *RateDifferentialResponse) SetMeasureRequested(v string)`

SetMeasureRequested sets MeasureRequested field to given value.


### GetMeasureUsed

`func (o *RateDifferentialResponse) GetMeasureUsed() string`

GetMeasureUsed returns the MeasureUsed field if non-nil, zero value otherwise.

### GetMeasureUsedOk

`func (o *RateDifferentialResponse) GetMeasureUsedOk() (*string, bool)`

GetMeasureUsedOk returns a tuple with the MeasureUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUsed

`func (o *RateDifferentialResponse) SetMeasureUsed(v string)`

SetMeasureUsed sets MeasureUsed field to given value.


### GetBaseIndicator

`func (o *RateDifferentialResponse) GetBaseIndicator() string`

GetBaseIndicator returns the BaseIndicator field if non-nil, zero value otherwise.

### GetBaseIndicatorOk

`func (o *RateDifferentialResponse) GetBaseIndicatorOk() (*string, bool)`

GetBaseIndicatorOk returns a tuple with the BaseIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseIndicator

`func (o *RateDifferentialResponse) SetBaseIndicator(v string)`

SetBaseIndicator sets BaseIndicator field to given value.

### HasBaseIndicator

`func (o *RateDifferentialResponse) HasBaseIndicator() bool`

HasBaseIndicator returns a boolean if a field has been set.

### SetBaseIndicatorNil

`func (o *RateDifferentialResponse) SetBaseIndicatorNil(b bool)`

 SetBaseIndicatorNil sets the value for BaseIndicator to be an explicit nil

### UnsetBaseIndicator
`func (o *RateDifferentialResponse) UnsetBaseIndicator()`

UnsetBaseIndicator ensures that no value is present for BaseIndicator, not even an explicit nil
### GetQuoteIndicator

`func (o *RateDifferentialResponse) GetQuoteIndicator() string`

GetQuoteIndicator returns the QuoteIndicator field if non-nil, zero value otherwise.

### GetQuoteIndicatorOk

`func (o *RateDifferentialResponse) GetQuoteIndicatorOk() (*string, bool)`

GetQuoteIndicatorOk returns a tuple with the QuoteIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteIndicator

`func (o *RateDifferentialResponse) SetQuoteIndicator(v string)`

SetQuoteIndicator sets QuoteIndicator field to given value.

### HasQuoteIndicator

`func (o *RateDifferentialResponse) HasQuoteIndicator() bool`

HasQuoteIndicator returns a boolean if a field has been set.

### SetQuoteIndicatorNil

`func (o *RateDifferentialResponse) SetQuoteIndicatorNil(b bool)`

 SetQuoteIndicatorNil sets the value for QuoteIndicator to be an explicit nil

### UnsetQuoteIndicator
`func (o *RateDifferentialResponse) UnsetQuoteIndicator()`

UnsetQuoteIndicator ensures that no value is present for QuoteIndicator, not even an explicit nil
### GetCurveFamily

`func (o *RateDifferentialResponse) GetCurveFamily() string`

GetCurveFamily returns the CurveFamily field if non-nil, zero value otherwise.

### GetCurveFamilyOk

`func (o *RateDifferentialResponse) GetCurveFamilyOk() (*string, bool)`

GetCurveFamilyOk returns a tuple with the CurveFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurveFamily

`func (o *RateDifferentialResponse) SetCurveFamily(v string)`

SetCurveFamily sets CurveFamily field to given value.

### HasCurveFamily

`func (o *RateDifferentialResponse) HasCurveFamily() bool`

HasCurveFamily returns a boolean if a field has been set.

### SetCurveFamilyNil

`func (o *RateDifferentialResponse) SetCurveFamilyNil(b bool)`

 SetCurveFamilyNil sets the value for CurveFamily to be an explicit nil

### UnsetCurveFamily
`func (o *RateDifferentialResponse) UnsetCurveFamily()`

UnsetCurveFamily ensures that no value is present for CurveFamily, not even an explicit nil
### GetStartTenor

`func (o *RateDifferentialResponse) GetStartTenor() string`

GetStartTenor returns the StartTenor field if non-nil, zero value otherwise.

### GetStartTenorOk

`func (o *RateDifferentialResponse) GetStartTenorOk() (*string, bool)`

GetStartTenorOk returns a tuple with the StartTenor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTenor

`func (o *RateDifferentialResponse) SetStartTenor(v string)`

SetStartTenor sets StartTenor field to given value.

### HasStartTenor

`func (o *RateDifferentialResponse) HasStartTenor() bool`

HasStartTenor returns a boolean if a field has been set.

### SetStartTenorNil

`func (o *RateDifferentialResponse) SetStartTenorNil(b bool)`

 SetStartTenorNil sets the value for StartTenor to be an explicit nil

### UnsetStartTenor
`func (o *RateDifferentialResponse) UnsetStartTenor()`

UnsetStartTenor ensures that no value is present for StartTenor, not even an explicit nil
### GetEndTenor

`func (o *RateDifferentialResponse) GetEndTenor() string`

GetEndTenor returns the EndTenor field if non-nil, zero value otherwise.

### GetEndTenorOk

`func (o *RateDifferentialResponse) GetEndTenorOk() (*string, bool)`

GetEndTenorOk returns a tuple with the EndTenor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTenor

`func (o *RateDifferentialResponse) SetEndTenor(v string)`

SetEndTenor sets EndTenor field to given value.

### HasEndTenor

`func (o *RateDifferentialResponse) HasEndTenor() bool`

HasEndTenor returns a boolean if a field has been set.

### SetEndTenorNil

`func (o *RateDifferentialResponse) SetEndTenorNil(b bool)`

 SetEndTenorNil sets the value for EndTenor to be an explicit nil

### UnsetEndTenor
`func (o *RateDifferentialResponse) UnsetEndTenor()`

UnsetEndTenor ensures that no value is present for EndTenor, not even an explicit nil
### GetForwardLabel

`func (o *RateDifferentialResponse) GetForwardLabel() string`

GetForwardLabel returns the ForwardLabel field if non-nil, zero value otherwise.

### GetForwardLabelOk

`func (o *RateDifferentialResponse) GetForwardLabelOk() (*string, bool)`

GetForwardLabelOk returns a tuple with the ForwardLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardLabel

`func (o *RateDifferentialResponse) SetForwardLabel(v string)`

SetForwardLabel sets ForwardLabel field to given value.

### HasForwardLabel

`func (o *RateDifferentialResponse) HasForwardLabel() bool`

HasForwardLabel returns a boolean if a field has been set.

### SetForwardLabelNil

`func (o *RateDifferentialResponse) SetForwardLabelNil(b bool)`

 SetForwardLabelNil sets the value for ForwardLabel to be an explicit nil

### UnsetForwardLabel
`func (o *RateDifferentialResponse) UnsetForwardLabel()`

UnsetForwardLabel ensures that no value is present for ForwardLabel, not even an explicit nil
### GetStartDate

`func (o *RateDifferentialResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *RateDifferentialResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *RateDifferentialResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *RateDifferentialResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *RateDifferentialResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *RateDifferentialResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetMatchedPoints

`func (o *RateDifferentialResponse) GetMatchedPoints() int64`

GetMatchedPoints returns the MatchedPoints field if non-nil, zero value otherwise.

### GetMatchedPointsOk

`func (o *RateDifferentialResponse) GetMatchedPointsOk() (*int64, bool)`

GetMatchedPointsOk returns a tuple with the MatchedPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedPoints

`func (o *RateDifferentialResponse) SetMatchedPoints(v int64)`

SetMatchedPoints sets MatchedPoints field to given value.


### GetUnit

`func (o *RateDifferentialResponse) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RateDifferentialResponse) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RateDifferentialResponse) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetLatestSpread

`func (o *RateDifferentialResponse) GetLatestSpread() float64`

GetLatestSpread returns the LatestSpread field if non-nil, zero value otherwise.

### GetLatestSpreadOk

`func (o *RateDifferentialResponse) GetLatestSpreadOk() (*float64, bool)`

GetLatestSpreadOk returns a tuple with the LatestSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSpread

`func (o *RateDifferentialResponse) SetLatestSpread(v float64)`

SetLatestSpread sets LatestSpread field to given value.

### HasLatestSpread

`func (o *RateDifferentialResponse) HasLatestSpread() bool`

HasLatestSpread returns a boolean if a field has been set.

### SetLatestSpreadNil

`func (o *RateDifferentialResponse) SetLatestSpreadNil(b bool)`

 SetLatestSpreadNil sets the value for LatestSpread to be an explicit nil

### UnsetLatestSpread
`func (o *RateDifferentialResponse) UnsetLatestSpread()`

UnsetLatestSpread ensures that no value is present for LatestSpread, not even an explicit nil
### GetLatestSpreadBps

`func (o *RateDifferentialResponse) GetLatestSpreadBps() float64`

GetLatestSpreadBps returns the LatestSpreadBps field if non-nil, zero value otherwise.

### GetLatestSpreadBpsOk

`func (o *RateDifferentialResponse) GetLatestSpreadBpsOk() (*float64, bool)`

GetLatestSpreadBpsOk returns a tuple with the LatestSpreadBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSpreadBps

`func (o *RateDifferentialResponse) SetLatestSpreadBps(v float64)`

SetLatestSpreadBps sets LatestSpreadBps field to given value.

### HasLatestSpreadBps

`func (o *RateDifferentialResponse) HasLatestSpreadBps() bool`

HasLatestSpreadBps returns a boolean if a field has been set.

### SetLatestSpreadBpsNil

`func (o *RateDifferentialResponse) SetLatestSpreadBpsNil(b bool)`

 SetLatestSpreadBpsNil sets the value for LatestSpreadBps to be an explicit nil

### UnsetLatestSpreadBps
`func (o *RateDifferentialResponse) UnsetLatestSpreadBps()`

UnsetLatestSpreadBps ensures that no value is present for LatestSpreadBps, not even an explicit nil
### GetLatestDifferential

`func (o *RateDifferentialResponse) GetLatestDifferential() float64`

GetLatestDifferential returns the LatestDifferential field if non-nil, zero value otherwise.

### GetLatestDifferentialOk

`func (o *RateDifferentialResponse) GetLatestDifferentialOk() (*float64, bool)`

GetLatestDifferentialOk returns a tuple with the LatestDifferential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDifferential

`func (o *RateDifferentialResponse) SetLatestDifferential(v float64)`

SetLatestDifferential sets LatestDifferential field to given value.

### HasLatestDifferential

`func (o *RateDifferentialResponse) HasLatestDifferential() bool`

HasLatestDifferential returns a boolean if a field has been set.

### SetLatestDifferentialNil

`func (o *RateDifferentialResponse) SetLatestDifferentialNil(b bool)`

 SetLatestDifferentialNil sets the value for LatestDifferential to be an explicit nil

### UnsetLatestDifferential
`func (o *RateDifferentialResponse) UnsetLatestDifferential()`

UnsetLatestDifferential ensures that no value is present for LatestDifferential, not even an explicit nil
### GetLatestDifferentialBps

`func (o *RateDifferentialResponse) GetLatestDifferentialBps() float64`

GetLatestDifferentialBps returns the LatestDifferentialBps field if non-nil, zero value otherwise.

### GetLatestDifferentialBpsOk

`func (o *RateDifferentialResponse) GetLatestDifferentialBpsOk() (*float64, bool)`

GetLatestDifferentialBpsOk returns a tuple with the LatestDifferentialBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDifferentialBps

`func (o *RateDifferentialResponse) SetLatestDifferentialBps(v float64)`

SetLatestDifferentialBps sets LatestDifferentialBps field to given value.

### HasLatestDifferentialBps

`func (o *RateDifferentialResponse) HasLatestDifferentialBps() bool`

HasLatestDifferentialBps returns a boolean if a field has been set.

### SetLatestDifferentialBpsNil

`func (o *RateDifferentialResponse) SetLatestDifferentialBpsNil(b bool)`

 SetLatestDifferentialBpsNil sets the value for LatestDifferentialBps to be an explicit nil

### UnsetLatestDifferentialBps
`func (o *RateDifferentialResponse) UnsetLatestDifferentialBps()`

UnsetLatestDifferentialBps ensures that no value is present for LatestDifferentialBps, not even an explicit nil
### GetBaseLatest

`func (o *RateDifferentialResponse) GetBaseLatest() float64`

GetBaseLatest returns the BaseLatest field if non-nil, zero value otherwise.

### GetBaseLatestOk

`func (o *RateDifferentialResponse) GetBaseLatestOk() (*float64, bool)`

GetBaseLatestOk returns a tuple with the BaseLatest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseLatest

`func (o *RateDifferentialResponse) SetBaseLatest(v float64)`

SetBaseLatest sets BaseLatest field to given value.

### HasBaseLatest

`func (o *RateDifferentialResponse) HasBaseLatest() bool`

HasBaseLatest returns a boolean if a field has been set.

### SetBaseLatestNil

`func (o *RateDifferentialResponse) SetBaseLatestNil(b bool)`

 SetBaseLatestNil sets the value for BaseLatest to be an explicit nil

### UnsetBaseLatest
`func (o *RateDifferentialResponse) UnsetBaseLatest()`

UnsetBaseLatest ensures that no value is present for BaseLatest, not even an explicit nil
### GetQuoteLatest

`func (o *RateDifferentialResponse) GetQuoteLatest() float64`

GetQuoteLatest returns the QuoteLatest field if non-nil, zero value otherwise.

### GetQuoteLatestOk

`func (o *RateDifferentialResponse) GetQuoteLatestOk() (*float64, bool)`

GetQuoteLatestOk returns a tuple with the QuoteLatest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteLatest

`func (o *RateDifferentialResponse) SetQuoteLatest(v float64)`

SetQuoteLatest sets QuoteLatest field to given value.

### HasQuoteLatest

`func (o *RateDifferentialResponse) HasQuoteLatest() bool`

HasQuoteLatest returns a boolean if a field has been set.

### SetQuoteLatestNil

`func (o *RateDifferentialResponse) SetQuoteLatestNil(b bool)`

 SetQuoteLatestNil sets the value for QuoteLatest to be an explicit nil

### UnsetQuoteLatest
`func (o *RateDifferentialResponse) UnsetQuoteLatest()`

UnsetQuoteLatest ensures that no value is present for QuoteLatest, not even an explicit nil
### GetSources

`func (o *RateDifferentialResponse) GetSources() RateDifferentialSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *RateDifferentialResponse) GetSourcesOk() (*RateDifferentialSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *RateDifferentialResponse) SetSources(v RateDifferentialSources)`

SetSources sets Sources field to given value.


### GetOfficialForwardSourceSupport

`func (o *RateDifferentialResponse) GetOfficialForwardSourceSupport() OfficialForwardSourceSupport`

GetOfficialForwardSourceSupport returns the OfficialForwardSourceSupport field if non-nil, zero value otherwise.

### GetOfficialForwardSourceSupportOk

`func (o *RateDifferentialResponse) GetOfficialForwardSourceSupportOk() (*OfficialForwardSourceSupport, bool)`

GetOfficialForwardSourceSupportOk returns a tuple with the OfficialForwardSourceSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialForwardSourceSupport

`func (o *RateDifferentialResponse) SetOfficialForwardSourceSupport(v OfficialForwardSourceSupport)`

SetOfficialForwardSourceSupport sets OfficialForwardSourceSupport field to given value.

### HasOfficialForwardSourceSupport

`func (o *RateDifferentialResponse) HasOfficialForwardSourceSupport() bool`

HasOfficialForwardSourceSupport returns a boolean if a field has been set.

### GetDataQuality

`func (o *RateDifferentialResponse) GetDataQuality() DataQuality`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *RateDifferentialResponse) GetDataQualityOk() (*DataQuality, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *RateDifferentialResponse) SetDataQuality(v DataQuality)`

SetDataQuality sets DataQuality field to given value.


### GetPagination

`func (o *RateDifferentialResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RateDifferentialResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RateDifferentialResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *RateDifferentialResponse) GetData() []RateDifferentialPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RateDifferentialResponse) GetDataOk() (*[]RateDifferentialPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RateDifferentialResponse) SetData(v []RateDifferentialPoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


