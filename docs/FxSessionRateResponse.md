# FxSessionRateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | **string** |  | 
**Quote** | **string** |  | 
**Session** | **string** |  | 
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**SourcePolicy** | [**FxSourcePolicy**](FxSourcePolicy.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Data** | [**[]FxSessionRatePoint**](FxSessionRatePoint.md) |  | 
**DatasetVersion** | **string** |  | 
**Coverage** | **map[string]interface{}** |  | 

## Methods

### NewFxSessionRateResponse

`func NewFxSessionRateResponse(base string, quote string, session string, startDate string, endDate string, sourcePolicy FxSourcePolicy, pagination Pagination, data []FxSessionRatePoint, datasetVersion string, coverage map[string]interface{}, ) *FxSessionRateResponse`

NewFxSessionRateResponse instantiates a new FxSessionRateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxSessionRateResponseWithDefaults

`func NewFxSessionRateResponseWithDefaults() *FxSessionRateResponse`

NewFxSessionRateResponseWithDefaults instantiates a new FxSessionRateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *FxSessionRateResponse) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *FxSessionRateResponse) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *FxSessionRateResponse) SetBase(v string)`

SetBase sets Base field to given value.


### GetQuote

`func (o *FxSessionRateResponse) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *FxSessionRateResponse) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *FxSessionRateResponse) SetQuote(v string)`

SetQuote sets Quote field to given value.


### GetSession

`func (o *FxSessionRateResponse) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *FxSessionRateResponse) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *FxSessionRateResponse) SetSession(v string)`

SetSession sets Session field to given value.


### GetStartDate

`func (o *FxSessionRateResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FxSessionRateResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FxSessionRateResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *FxSessionRateResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FxSessionRateResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FxSessionRateResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetSourcePolicy

`func (o *FxSessionRateResponse) GetSourcePolicy() FxSourcePolicy`

GetSourcePolicy returns the SourcePolicy field if non-nil, zero value otherwise.

### GetSourcePolicyOk

`func (o *FxSessionRateResponse) GetSourcePolicyOk() (*FxSourcePolicy, bool)`

GetSourcePolicyOk returns a tuple with the SourcePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePolicy

`func (o *FxSessionRateResponse) SetSourcePolicy(v FxSourcePolicy)`

SetSourcePolicy sets SourcePolicy field to given value.


### GetPagination

`func (o *FxSessionRateResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *FxSessionRateResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *FxSessionRateResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *FxSessionRateResponse) GetData() []FxSessionRatePoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FxSessionRateResponse) GetDataOk() (*[]FxSessionRatePoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FxSessionRateResponse) SetData(v []FxSessionRatePoint)`

SetData sets Data field to given value.


### GetDatasetVersion

`func (o *FxSessionRateResponse) GetDatasetVersion() string`

GetDatasetVersion returns the DatasetVersion field if non-nil, zero value otherwise.

### GetDatasetVersionOk

`func (o *FxSessionRateResponse) GetDatasetVersionOk() (*string, bool)`

GetDatasetVersionOk returns a tuple with the DatasetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetVersion

`func (o *FxSessionRateResponse) SetDatasetVersion(v string)`

SetDatasetVersion sets DatasetVersion field to given value.


### GetCoverage

`func (o *FxSessionRateResponse) GetCoverage() map[string]interface{}`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *FxSessionRateResponse) GetCoverageOk() (*map[string]interface{}, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *FxSessionRateResponse) SetCoverage(v map[string]interface{})`

SetCoverage sets Coverage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


