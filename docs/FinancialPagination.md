# FinancialPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int64** |  | 
**Offset** | **int64** |  | 
**ReturnedCount** | **int64** |  | 
**TotalCount** | **int64** |  | 
**HasMore** | **bool** |  | 
**NextOffset** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewFinancialPagination

`func NewFinancialPagination(limit int64, offset int64, returnedCount int64, totalCount int64, hasMore bool, ) *FinancialPagination`

NewFinancialPagination instantiates a new FinancialPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialPaginationWithDefaults

`func NewFinancialPaginationWithDefaults() *FinancialPagination`

NewFinancialPaginationWithDefaults instantiates a new FinancialPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *FinancialPagination) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *FinancialPagination) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *FinancialPagination) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *FinancialPagination) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *FinancialPagination) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *FinancialPagination) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetReturnedCount

`func (o *FinancialPagination) GetReturnedCount() int64`

GetReturnedCount returns the ReturnedCount field if non-nil, zero value otherwise.

### GetReturnedCountOk

`func (o *FinancialPagination) GetReturnedCountOk() (*int64, bool)`

GetReturnedCountOk returns a tuple with the ReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedCount

`func (o *FinancialPagination) SetReturnedCount(v int64)`

SetReturnedCount sets ReturnedCount field to given value.


### GetTotalCount

`func (o *FinancialPagination) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *FinancialPagination) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *FinancialPagination) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetHasMore

`func (o *FinancialPagination) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *FinancialPagination) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *FinancialPagination) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNextOffset

`func (o *FinancialPagination) GetNextOffset() int64`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *FinancialPagination) GetNextOffsetOk() (*int64, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *FinancialPagination) SetNextOffset(v int64)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *FinancialPagination) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### SetNextOffsetNil

`func (o *FinancialPagination) SetNextOffsetNil(b bool)`

 SetNextOffsetNil sets the value for NextOffset to be an explicit nil

### UnsetNextOffset
`func (o *FinancialPagination) UnsetNextOffset()`

UnsetNextOffset ensures that no value is present for NextOffset, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


