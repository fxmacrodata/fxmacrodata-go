# CommodityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**Indicator** | **string** |  | 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceUrl** | Pointer to **NullableString** |  | [optional] 
**Provenance** | [**SourceProvenance**](SourceProvenance.md) |  | 
**HasOfficialForecast** | **bool** |  | 
**LastUpdated** | Pointer to **NullableString** |  | [optional] 
**LatestAvailableDate** | Pointer to **NullableString** |  | [optional] 
**DataQuality** | [**DataQuality**](DataQuality.md) |  | 
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Data** | [**[]CommodityDataPoint**](CommodityDataPoint.md) |  | 

## Methods

### NewCommodityResponse

`func NewCommodityResponse(currency string, indicator string, provenance SourceProvenance, hasOfficialForecast bool, dataQuality DataQuality, startDate string, endDate string, pagination Pagination, data []CommodityDataPoint, ) *CommodityResponse`

NewCommodityResponse instantiates a new CommodityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommodityResponseWithDefaults

`func NewCommodityResponseWithDefaults() *CommodityResponse`

NewCommodityResponseWithDefaults instantiates a new CommodityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CommodityResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CommodityResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CommodityResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetIndicator

`func (o *CommodityResponse) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *CommodityResponse) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *CommodityResponse) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.


### GetSource

`func (o *CommodityResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CommodityResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CommodityResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CommodityResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *CommodityResponse) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *CommodityResponse) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceUrl

`func (o *CommodityResponse) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *CommodityResponse) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *CommodityResponse) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *CommodityResponse) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *CommodityResponse) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *CommodityResponse) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetProvenance

`func (o *CommodityResponse) GetProvenance() SourceProvenance`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *CommodityResponse) GetProvenanceOk() (*SourceProvenance, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *CommodityResponse) SetProvenance(v SourceProvenance)`

SetProvenance sets Provenance field to given value.


### GetHasOfficialForecast

`func (o *CommodityResponse) GetHasOfficialForecast() bool`

GetHasOfficialForecast returns the HasOfficialForecast field if non-nil, zero value otherwise.

### GetHasOfficialForecastOk

`func (o *CommodityResponse) GetHasOfficialForecastOk() (*bool, bool)`

GetHasOfficialForecastOk returns a tuple with the HasOfficialForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOfficialForecast

`func (o *CommodityResponse) SetHasOfficialForecast(v bool)`

SetHasOfficialForecast sets HasOfficialForecast field to given value.


### GetLastUpdated

`func (o *CommodityResponse) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CommodityResponse) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CommodityResponse) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CommodityResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *CommodityResponse) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *CommodityResponse) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetLatestAvailableDate

`func (o *CommodityResponse) GetLatestAvailableDate() string`

GetLatestAvailableDate returns the LatestAvailableDate field if non-nil, zero value otherwise.

### GetLatestAvailableDateOk

`func (o *CommodityResponse) GetLatestAvailableDateOk() (*string, bool)`

GetLatestAvailableDateOk returns a tuple with the LatestAvailableDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAvailableDate

`func (o *CommodityResponse) SetLatestAvailableDate(v string)`

SetLatestAvailableDate sets LatestAvailableDate field to given value.

### HasLatestAvailableDate

`func (o *CommodityResponse) HasLatestAvailableDate() bool`

HasLatestAvailableDate returns a boolean if a field has been set.

### SetLatestAvailableDateNil

`func (o *CommodityResponse) SetLatestAvailableDateNil(b bool)`

 SetLatestAvailableDateNil sets the value for LatestAvailableDate to be an explicit nil

### UnsetLatestAvailableDate
`func (o *CommodityResponse) UnsetLatestAvailableDate()`

UnsetLatestAvailableDate ensures that no value is present for LatestAvailableDate, not even an explicit nil
### GetDataQuality

`func (o *CommodityResponse) GetDataQuality() DataQuality`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *CommodityResponse) GetDataQualityOk() (*DataQuality, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *CommodityResponse) SetDataQuality(v DataQuality)`

SetDataQuality sets DataQuality field to given value.


### GetStartDate

`func (o *CommodityResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CommodityResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CommodityResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *CommodityResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CommodityResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CommodityResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetPagination

`func (o *CommodityResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CommodityResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CommodityResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *CommodityResponse) GetData() []CommodityDataPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CommodityResponse) GetDataOk() (*[]CommodityDataPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CommodityResponse) SetData(v []CommodityDataPoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


