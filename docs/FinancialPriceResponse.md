# FinancialPriceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**Sources** | [**[]FinancialSource**](FinancialSource.md) |  | 
**Coverage** | [**[]FinancialCoverage**](FinancialCoverage.md) |  | 
**Pagination** | [**FinancialPagination**](FinancialPagination.md) |  | 
**Data** | [**[]FinancialObservation**](FinancialObservation.md) |  | 

## Methods

### NewFinancialPriceResponse

`func NewFinancialPriceResponse(currency string, startDate string, endDate string, sources []FinancialSource, coverage []FinancialCoverage, pagination FinancialPagination, data []FinancialObservation, ) *FinancialPriceResponse`

NewFinancialPriceResponse instantiates a new FinancialPriceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialPriceResponseWithDefaults

`func NewFinancialPriceResponseWithDefaults() *FinancialPriceResponse`

NewFinancialPriceResponseWithDefaults instantiates a new FinancialPriceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *FinancialPriceResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FinancialPriceResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FinancialPriceResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetStartDate

`func (o *FinancialPriceResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FinancialPriceResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FinancialPriceResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *FinancialPriceResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FinancialPriceResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FinancialPriceResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetSources

`func (o *FinancialPriceResponse) GetSources() []FinancialSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *FinancialPriceResponse) GetSourcesOk() (*[]FinancialSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *FinancialPriceResponse) SetSources(v []FinancialSource)`

SetSources sets Sources field to given value.


### GetCoverage

`func (o *FinancialPriceResponse) GetCoverage() []FinancialCoverage`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *FinancialPriceResponse) GetCoverageOk() (*[]FinancialCoverage, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *FinancialPriceResponse) SetCoverage(v []FinancialCoverage)`

SetCoverage sets Coverage field to given value.


### GetPagination

`func (o *FinancialPriceResponse) GetPagination() FinancialPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *FinancialPriceResponse) GetPaginationOk() (*FinancialPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *FinancialPriceResponse) SetPagination(v FinancialPagination)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *FinancialPriceResponse) GetData() []FinancialObservation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FinancialPriceResponse) GetDataOk() (*[]FinancialObservation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FinancialPriceResponse) SetData(v []FinancialObservation)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


