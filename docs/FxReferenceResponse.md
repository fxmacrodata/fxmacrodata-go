# FxReferenceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | **string** |  | 
**Quote** | **string** |  | 
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**TimestampBasis** | **string** |  | 
**SourcePolicy** | [**FxSourcePolicy**](FxSourcePolicy.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Data** | [**[]FxReferencePoint**](FxReferencePoint.md) |  | 
**DatasetVersion** | **string** |  | 
**Coverage** | **map[string]interface{}** |  | 

## Methods

### NewFxReferenceResponse

`func NewFxReferenceResponse(base string, quote string, startDate string, endDate string, timestampBasis string, sourcePolicy FxSourcePolicy, pagination Pagination, data []FxReferencePoint, datasetVersion string, coverage map[string]interface{}, ) *FxReferenceResponse`

NewFxReferenceResponse instantiates a new FxReferenceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxReferenceResponseWithDefaults

`func NewFxReferenceResponseWithDefaults() *FxReferenceResponse`

NewFxReferenceResponseWithDefaults instantiates a new FxReferenceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *FxReferenceResponse) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *FxReferenceResponse) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *FxReferenceResponse) SetBase(v string)`

SetBase sets Base field to given value.


### GetQuote

`func (o *FxReferenceResponse) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *FxReferenceResponse) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *FxReferenceResponse) SetQuote(v string)`

SetQuote sets Quote field to given value.


### GetStartDate

`func (o *FxReferenceResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FxReferenceResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FxReferenceResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *FxReferenceResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FxReferenceResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FxReferenceResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetTimestampBasis

`func (o *FxReferenceResponse) GetTimestampBasis() string`

GetTimestampBasis returns the TimestampBasis field if non-nil, zero value otherwise.

### GetTimestampBasisOk

`func (o *FxReferenceResponse) GetTimestampBasisOk() (*string, bool)`

GetTimestampBasisOk returns a tuple with the TimestampBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampBasis

`func (o *FxReferenceResponse) SetTimestampBasis(v string)`

SetTimestampBasis sets TimestampBasis field to given value.


### GetSourcePolicy

`func (o *FxReferenceResponse) GetSourcePolicy() FxSourcePolicy`

GetSourcePolicy returns the SourcePolicy field if non-nil, zero value otherwise.

### GetSourcePolicyOk

`func (o *FxReferenceResponse) GetSourcePolicyOk() (*FxSourcePolicy, bool)`

GetSourcePolicyOk returns a tuple with the SourcePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePolicy

`func (o *FxReferenceResponse) SetSourcePolicy(v FxSourcePolicy)`

SetSourcePolicy sets SourcePolicy field to given value.


### GetPagination

`func (o *FxReferenceResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *FxReferenceResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *FxReferenceResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *FxReferenceResponse) GetData() []FxReferencePoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FxReferenceResponse) GetDataOk() (*[]FxReferencePoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FxReferenceResponse) SetData(v []FxReferencePoint)`

SetData sets Data field to given value.


### GetDatasetVersion

`func (o *FxReferenceResponse) GetDatasetVersion() string`

GetDatasetVersion returns the DatasetVersion field if non-nil, zero value otherwise.

### GetDatasetVersionOk

`func (o *FxReferenceResponse) GetDatasetVersionOk() (*string, bool)`

GetDatasetVersionOk returns a tuple with the DatasetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetVersion

`func (o *FxReferenceResponse) SetDatasetVersion(v string)`

SetDatasetVersion sets DatasetVersion field to given value.


### GetCoverage

`func (o *FxReferenceResponse) GetCoverage() map[string]interface{}`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *FxReferenceResponse) GetCoverageOk() (*map[string]interface{}, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *FxReferenceResponse) SetCoverage(v map[string]interface{})`

SetCoverage sets Coverage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


