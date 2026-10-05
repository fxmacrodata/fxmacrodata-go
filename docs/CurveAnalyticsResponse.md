# CurveAnalyticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Research** | Pointer to **map[string]interface{}** |  | [optional] 
**Currency** | **string** |  | 
**CurveFamily** | **string** |  | 
**View** | Pointer to **string** |  | [optional] [default to "nodes"]
**Metric** | Pointer to **NullableString** |  | [optional] 
**Method** | Pointer to **NullableString** |  | [optional] 
**RequestedDate** | **string** |  | 
**AsOf** | Pointer to **NullableString** |  | [optional] 
**NodeCount** | **int64** |  | 
**SlopeCount** | Pointer to **NullableInt64** |  | [optional] 
**InvertedCount** | Pointer to **NullableInt64** |  | [optional] 
**SegmentCount** | Pointer to **NullableInt64** |  | [optional] 
**Sources** | **[]string** |  | 
**OfficialForwardSourceSupport** | [**OfficialForwardSourceSupport**](OfficialForwardSourceSupport.md) |  | 
**DataQuality** | [**DataQuality**](DataQuality.md) |  | 
**Data** | **[]map[string]interface{}** |  | 

## Methods

### NewCurveAnalyticsResponse

`func NewCurveAnalyticsResponse(currency string, curveFamily string, requestedDate string, nodeCount int64, sources []string, officialForwardSourceSupport OfficialForwardSourceSupport, dataQuality DataQuality, data []map[string]interface{}, ) *CurveAnalyticsResponse`

NewCurveAnalyticsResponse instantiates a new CurveAnalyticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurveAnalyticsResponseWithDefaults

`func NewCurveAnalyticsResponseWithDefaults() *CurveAnalyticsResponse`

NewCurveAnalyticsResponseWithDefaults instantiates a new CurveAnalyticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResearch

`func (o *CurveAnalyticsResponse) GetResearch() map[string]interface{}`

GetResearch returns the Research field if non-nil, zero value otherwise.

### GetResearchOk

`func (o *CurveAnalyticsResponse) GetResearchOk() (*map[string]interface{}, bool)`

GetResearchOk returns a tuple with the Research field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearch

`func (o *CurveAnalyticsResponse) SetResearch(v map[string]interface{})`

SetResearch sets Research field to given value.

### HasResearch

`func (o *CurveAnalyticsResponse) HasResearch() bool`

HasResearch returns a boolean if a field has been set.

### SetResearchNil

`func (o *CurveAnalyticsResponse) SetResearchNil(b bool)`

 SetResearchNil sets the value for Research to be an explicit nil

### UnsetResearch
`func (o *CurveAnalyticsResponse) UnsetResearch()`

UnsetResearch ensures that no value is present for Research, not even an explicit nil
### GetCurrency

`func (o *CurveAnalyticsResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CurveAnalyticsResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CurveAnalyticsResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCurveFamily

`func (o *CurveAnalyticsResponse) GetCurveFamily() string`

GetCurveFamily returns the CurveFamily field if non-nil, zero value otherwise.

### GetCurveFamilyOk

`func (o *CurveAnalyticsResponse) GetCurveFamilyOk() (*string, bool)`

GetCurveFamilyOk returns a tuple with the CurveFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurveFamily

`func (o *CurveAnalyticsResponse) SetCurveFamily(v string)`

SetCurveFamily sets CurveFamily field to given value.


### GetView

`func (o *CurveAnalyticsResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *CurveAnalyticsResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *CurveAnalyticsResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *CurveAnalyticsResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetMetric

`func (o *CurveAnalyticsResponse) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *CurveAnalyticsResponse) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *CurveAnalyticsResponse) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *CurveAnalyticsResponse) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### SetMetricNil

`func (o *CurveAnalyticsResponse) SetMetricNil(b bool)`

 SetMetricNil sets the value for Metric to be an explicit nil

### UnsetMetric
`func (o *CurveAnalyticsResponse) UnsetMetric()`

UnsetMetric ensures that no value is present for Metric, not even an explicit nil
### GetMethod

`func (o *CurveAnalyticsResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CurveAnalyticsResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CurveAnalyticsResponse) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CurveAnalyticsResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *CurveAnalyticsResponse) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *CurveAnalyticsResponse) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetRequestedDate

`func (o *CurveAnalyticsResponse) GetRequestedDate() string`

GetRequestedDate returns the RequestedDate field if non-nil, zero value otherwise.

### GetRequestedDateOk

`func (o *CurveAnalyticsResponse) GetRequestedDateOk() (*string, bool)`

GetRequestedDateOk returns a tuple with the RequestedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDate

`func (o *CurveAnalyticsResponse) SetRequestedDate(v string)`

SetRequestedDate sets RequestedDate field to given value.


### GetAsOf

`func (o *CurveAnalyticsResponse) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *CurveAnalyticsResponse) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *CurveAnalyticsResponse) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *CurveAnalyticsResponse) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *CurveAnalyticsResponse) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *CurveAnalyticsResponse) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetNodeCount

`func (o *CurveAnalyticsResponse) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *CurveAnalyticsResponse) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *CurveAnalyticsResponse) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.


### GetSlopeCount

`func (o *CurveAnalyticsResponse) GetSlopeCount() int64`

GetSlopeCount returns the SlopeCount field if non-nil, zero value otherwise.

### GetSlopeCountOk

`func (o *CurveAnalyticsResponse) GetSlopeCountOk() (*int64, bool)`

GetSlopeCountOk returns a tuple with the SlopeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlopeCount

`func (o *CurveAnalyticsResponse) SetSlopeCount(v int64)`

SetSlopeCount sets SlopeCount field to given value.

### HasSlopeCount

`func (o *CurveAnalyticsResponse) HasSlopeCount() bool`

HasSlopeCount returns a boolean if a field has been set.

### SetSlopeCountNil

`func (o *CurveAnalyticsResponse) SetSlopeCountNil(b bool)`

 SetSlopeCountNil sets the value for SlopeCount to be an explicit nil

### UnsetSlopeCount
`func (o *CurveAnalyticsResponse) UnsetSlopeCount()`

UnsetSlopeCount ensures that no value is present for SlopeCount, not even an explicit nil
### GetInvertedCount

`func (o *CurveAnalyticsResponse) GetInvertedCount() int64`

GetInvertedCount returns the InvertedCount field if non-nil, zero value otherwise.

### GetInvertedCountOk

`func (o *CurveAnalyticsResponse) GetInvertedCountOk() (*int64, bool)`

GetInvertedCountOk returns a tuple with the InvertedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvertedCount

`func (o *CurveAnalyticsResponse) SetInvertedCount(v int64)`

SetInvertedCount sets InvertedCount field to given value.

### HasInvertedCount

`func (o *CurveAnalyticsResponse) HasInvertedCount() bool`

HasInvertedCount returns a boolean if a field has been set.

### SetInvertedCountNil

`func (o *CurveAnalyticsResponse) SetInvertedCountNil(b bool)`

 SetInvertedCountNil sets the value for InvertedCount to be an explicit nil

### UnsetInvertedCount
`func (o *CurveAnalyticsResponse) UnsetInvertedCount()`

UnsetInvertedCount ensures that no value is present for InvertedCount, not even an explicit nil
### GetSegmentCount

`func (o *CurveAnalyticsResponse) GetSegmentCount() int64`

GetSegmentCount returns the SegmentCount field if non-nil, zero value otherwise.

### GetSegmentCountOk

`func (o *CurveAnalyticsResponse) GetSegmentCountOk() (*int64, bool)`

GetSegmentCountOk returns a tuple with the SegmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentCount

`func (o *CurveAnalyticsResponse) SetSegmentCount(v int64)`

SetSegmentCount sets SegmentCount field to given value.

### HasSegmentCount

`func (o *CurveAnalyticsResponse) HasSegmentCount() bool`

HasSegmentCount returns a boolean if a field has been set.

### SetSegmentCountNil

`func (o *CurveAnalyticsResponse) SetSegmentCountNil(b bool)`

 SetSegmentCountNil sets the value for SegmentCount to be an explicit nil

### UnsetSegmentCount
`func (o *CurveAnalyticsResponse) UnsetSegmentCount()`

UnsetSegmentCount ensures that no value is present for SegmentCount, not even an explicit nil
### GetSources

`func (o *CurveAnalyticsResponse) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CurveAnalyticsResponse) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CurveAnalyticsResponse) SetSources(v []string)`

SetSources sets Sources field to given value.


### GetOfficialForwardSourceSupport

`func (o *CurveAnalyticsResponse) GetOfficialForwardSourceSupport() OfficialForwardSourceSupport`

GetOfficialForwardSourceSupport returns the OfficialForwardSourceSupport field if non-nil, zero value otherwise.

### GetOfficialForwardSourceSupportOk

`func (o *CurveAnalyticsResponse) GetOfficialForwardSourceSupportOk() (*OfficialForwardSourceSupport, bool)`

GetOfficialForwardSourceSupportOk returns a tuple with the OfficialForwardSourceSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialForwardSourceSupport

`func (o *CurveAnalyticsResponse) SetOfficialForwardSourceSupport(v OfficialForwardSourceSupport)`

SetOfficialForwardSourceSupport sets OfficialForwardSourceSupport field to given value.


### GetDataQuality

`func (o *CurveAnalyticsResponse) GetDataQuality() DataQuality`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *CurveAnalyticsResponse) GetDataQualityOk() (*DataQuality, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *CurveAnalyticsResponse) SetDataQuality(v DataQuality)`

SetDataQuality sets DataQuality field to given value.


### GetData

`func (o *CurveAnalyticsResponse) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CurveAnalyticsResponse) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CurveAnalyticsResponse) SetData(v []map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


