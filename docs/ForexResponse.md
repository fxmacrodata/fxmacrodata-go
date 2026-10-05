# ForexResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | **string** |  | 
**Quote** | **string** |  | 
**Source** | **string** |  | 
**Provenance** | [**ForexProvenance**](ForexProvenance.md) |  | 
**PairMetadata** | Pointer to [**ForexPairMetadata**](ForexPairMetadata.md) |  | [optional] 
**DataQuality** | [**DataQuality**](DataQuality.md) |  | 
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Data** | [**[]ForexDataPoint**](ForexDataPoint.md) |  | 
**Indicators** | Pointer to **map[string]interface{}** |  | [optional] 
**DailyOhlcBasis** | Pointer to **map[string]interface{}** |  | [optional] 
**TechnicalIndicatorBasis** | Pointer to **map[string]interface{}** |  | [optional] 
**DatasetVersion** | **string** |  | 
**Coverage** | **map[string]interface{}** |  | 

## Methods

### NewForexResponse

`func NewForexResponse(base string, quote string, source string, provenance ForexProvenance, dataQuality DataQuality, startDate string, endDate string, pagination Pagination, data []ForexDataPoint, datasetVersion string, coverage map[string]interface{}, ) *ForexResponse`

NewForexResponse instantiates a new ForexResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForexResponseWithDefaults

`func NewForexResponseWithDefaults() *ForexResponse`

NewForexResponseWithDefaults instantiates a new ForexResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *ForexResponse) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *ForexResponse) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *ForexResponse) SetBase(v string)`

SetBase sets Base field to given value.


### GetQuote

`func (o *ForexResponse) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *ForexResponse) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *ForexResponse) SetQuote(v string)`

SetQuote sets Quote field to given value.


### GetSource

`func (o *ForexResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ForexResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ForexResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetProvenance

`func (o *ForexResponse) GetProvenance() ForexProvenance`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *ForexResponse) GetProvenanceOk() (*ForexProvenance, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *ForexResponse) SetProvenance(v ForexProvenance)`

SetProvenance sets Provenance field to given value.


### GetPairMetadata

`func (o *ForexResponse) GetPairMetadata() ForexPairMetadata`

GetPairMetadata returns the PairMetadata field if non-nil, zero value otherwise.

### GetPairMetadataOk

`func (o *ForexResponse) GetPairMetadataOk() (*ForexPairMetadata, bool)`

GetPairMetadataOk returns a tuple with the PairMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairMetadata

`func (o *ForexResponse) SetPairMetadata(v ForexPairMetadata)`

SetPairMetadata sets PairMetadata field to given value.

### HasPairMetadata

`func (o *ForexResponse) HasPairMetadata() bool`

HasPairMetadata returns a boolean if a field has been set.

### GetDataQuality

`func (o *ForexResponse) GetDataQuality() DataQuality`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *ForexResponse) GetDataQualityOk() (*DataQuality, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *ForexResponse) SetDataQuality(v DataQuality)`

SetDataQuality sets DataQuality field to given value.


### GetStartDate

`func (o *ForexResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ForexResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ForexResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *ForexResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ForexResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ForexResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetPagination

`func (o *ForexResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ForexResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ForexResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *ForexResponse) GetData() []ForexDataPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ForexResponse) GetDataOk() (*[]ForexDataPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ForexResponse) SetData(v []ForexDataPoint)`

SetData sets Data field to given value.


### GetIndicators

`func (o *ForexResponse) GetIndicators() map[string]interface{}`

GetIndicators returns the Indicators field if non-nil, zero value otherwise.

### GetIndicatorsOk

`func (o *ForexResponse) GetIndicatorsOk() (*map[string]interface{}, bool)`

GetIndicatorsOk returns a tuple with the Indicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicators

`func (o *ForexResponse) SetIndicators(v map[string]interface{})`

SetIndicators sets Indicators field to given value.

### HasIndicators

`func (o *ForexResponse) HasIndicators() bool`

HasIndicators returns a boolean if a field has been set.

### SetIndicatorsNil

`func (o *ForexResponse) SetIndicatorsNil(b bool)`

 SetIndicatorsNil sets the value for Indicators to be an explicit nil

### UnsetIndicators
`func (o *ForexResponse) UnsetIndicators()`

UnsetIndicators ensures that no value is present for Indicators, not even an explicit nil
### GetDailyOhlcBasis

`func (o *ForexResponse) GetDailyOhlcBasis() map[string]interface{}`

GetDailyOhlcBasis returns the DailyOhlcBasis field if non-nil, zero value otherwise.

### GetDailyOhlcBasisOk

`func (o *ForexResponse) GetDailyOhlcBasisOk() (*map[string]interface{}, bool)`

GetDailyOhlcBasisOk returns a tuple with the DailyOhlcBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyOhlcBasis

`func (o *ForexResponse) SetDailyOhlcBasis(v map[string]interface{})`

SetDailyOhlcBasis sets DailyOhlcBasis field to given value.

### HasDailyOhlcBasis

`func (o *ForexResponse) HasDailyOhlcBasis() bool`

HasDailyOhlcBasis returns a boolean if a field has been set.

### SetDailyOhlcBasisNil

`func (o *ForexResponse) SetDailyOhlcBasisNil(b bool)`

 SetDailyOhlcBasisNil sets the value for DailyOhlcBasis to be an explicit nil

### UnsetDailyOhlcBasis
`func (o *ForexResponse) UnsetDailyOhlcBasis()`

UnsetDailyOhlcBasis ensures that no value is present for DailyOhlcBasis, not even an explicit nil
### GetTechnicalIndicatorBasis

`func (o *ForexResponse) GetTechnicalIndicatorBasis() map[string]interface{}`

GetTechnicalIndicatorBasis returns the TechnicalIndicatorBasis field if non-nil, zero value otherwise.

### GetTechnicalIndicatorBasisOk

`func (o *ForexResponse) GetTechnicalIndicatorBasisOk() (*map[string]interface{}, bool)`

GetTechnicalIndicatorBasisOk returns a tuple with the TechnicalIndicatorBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalIndicatorBasis

`func (o *ForexResponse) SetTechnicalIndicatorBasis(v map[string]interface{})`

SetTechnicalIndicatorBasis sets TechnicalIndicatorBasis field to given value.

### HasTechnicalIndicatorBasis

`func (o *ForexResponse) HasTechnicalIndicatorBasis() bool`

HasTechnicalIndicatorBasis returns a boolean if a field has been set.

### SetTechnicalIndicatorBasisNil

`func (o *ForexResponse) SetTechnicalIndicatorBasisNil(b bool)`

 SetTechnicalIndicatorBasisNil sets the value for TechnicalIndicatorBasis to be an explicit nil

### UnsetTechnicalIndicatorBasis
`func (o *ForexResponse) UnsetTechnicalIndicatorBasis()`

UnsetTechnicalIndicatorBasis ensures that no value is present for TechnicalIndicatorBasis, not even an explicit nil
### GetDatasetVersion

`func (o *ForexResponse) GetDatasetVersion() string`

GetDatasetVersion returns the DatasetVersion field if non-nil, zero value otherwise.

### GetDatasetVersionOk

`func (o *ForexResponse) GetDatasetVersionOk() (*string, bool)`

GetDatasetVersionOk returns a tuple with the DatasetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetVersion

`func (o *ForexResponse) SetDatasetVersion(v string)`

SetDatasetVersion sets DatasetVersion field to given value.


### GetCoverage

`func (o *ForexResponse) GetCoverage() map[string]interface{}`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *ForexResponse) GetCoverageOk() (*map[string]interface{}, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *ForexResponse) SetCoverage(v map[string]interface{})`

SetCoverage sets Coverage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


