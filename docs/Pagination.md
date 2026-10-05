# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **NullableInt64** |  | [optional] 
**Offset** | Pointer to **NullableInt64** |  | [optional] 
**ReturnedCount** | Pointer to **NullableInt64** |  | [optional] 
**TotalCount** | Pointer to **NullableInt64** |  | [optional] 
**HasMore** | Pointer to **NullableBool** |  | [optional] 
**NextOffset** | Pointer to **NullableInt64** |  | [optional] 
**PageIncludesLatestAvailable** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewPagination

`func NewPagination() *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *Pagination) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Pagination) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Pagination) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Pagination) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *Pagination) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *Pagination) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOffset

`func (o *Pagination) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Pagination) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Pagination) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Pagination) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *Pagination) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *Pagination) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetReturnedCount

`func (o *Pagination) GetReturnedCount() int64`

GetReturnedCount returns the ReturnedCount field if non-nil, zero value otherwise.

### GetReturnedCountOk

`func (o *Pagination) GetReturnedCountOk() (*int64, bool)`

GetReturnedCountOk returns a tuple with the ReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedCount

`func (o *Pagination) SetReturnedCount(v int64)`

SetReturnedCount sets ReturnedCount field to given value.

### HasReturnedCount

`func (o *Pagination) HasReturnedCount() bool`

HasReturnedCount returns a boolean if a field has been set.

### SetReturnedCountNil

`func (o *Pagination) SetReturnedCountNil(b bool)`

 SetReturnedCountNil sets the value for ReturnedCount to be an explicit nil

### UnsetReturnedCount
`func (o *Pagination) UnsetReturnedCount()`

UnsetReturnedCount ensures that no value is present for ReturnedCount, not even an explicit nil
### GetTotalCount

`func (o *Pagination) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Pagination) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Pagination) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *Pagination) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### SetTotalCountNil

`func (o *Pagination) SetTotalCountNil(b bool)`

 SetTotalCountNil sets the value for TotalCount to be an explicit nil

### UnsetTotalCount
`func (o *Pagination) UnsetTotalCount()`

UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil
### GetHasMore

`func (o *Pagination) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *Pagination) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *Pagination) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *Pagination) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### SetHasMoreNil

`func (o *Pagination) SetHasMoreNil(b bool)`

 SetHasMoreNil sets the value for HasMore to be an explicit nil

### UnsetHasMore
`func (o *Pagination) UnsetHasMore()`

UnsetHasMore ensures that no value is present for HasMore, not even an explicit nil
### GetNextOffset

`func (o *Pagination) GetNextOffset() int64`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *Pagination) GetNextOffsetOk() (*int64, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *Pagination) SetNextOffset(v int64)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *Pagination) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### SetNextOffsetNil

`func (o *Pagination) SetNextOffsetNil(b bool)`

 SetNextOffsetNil sets the value for NextOffset to be an explicit nil

### UnsetNextOffset
`func (o *Pagination) UnsetNextOffset()`

UnsetNextOffset ensures that no value is present for NextOffset, not even an explicit nil
### GetPageIncludesLatestAvailable

`func (o *Pagination) GetPageIncludesLatestAvailable() bool`

GetPageIncludesLatestAvailable returns the PageIncludesLatestAvailable field if non-nil, zero value otherwise.

### GetPageIncludesLatestAvailableOk

`func (o *Pagination) GetPageIncludesLatestAvailableOk() (*bool, bool)`

GetPageIncludesLatestAvailableOk returns a tuple with the PageIncludesLatestAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIncludesLatestAvailable

`func (o *Pagination) SetPageIncludesLatestAvailable(v bool)`

SetPageIncludesLatestAvailable sets PageIncludesLatestAvailable field to given value.

### HasPageIncludesLatestAvailable

`func (o *Pagination) HasPageIncludesLatestAvailable() bool`

HasPageIncludesLatestAvailable returns a boolean if a field has been set.

### SetPageIncludesLatestAvailableNil

`func (o *Pagination) SetPageIncludesLatestAvailableNil(b bool)`

 SetPageIncludesLatestAvailableNil sets the value for PageIncludesLatestAvailable to be an explicit nil

### UnsetPageIncludesLatestAvailable
`func (o *Pagination) UnsetPageIncludesLatestAvailable()`

UnsetPageIncludesLatestAvailable ensures that no value is present for PageIncludesLatestAvailable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


