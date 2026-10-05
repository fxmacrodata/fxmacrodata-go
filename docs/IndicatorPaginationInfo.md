# IndicatorPaginationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **NullableInt64** |  | [optional] 
**Offset** | **int64** |  | 
**ReturnedCount** | **int64** |  | 
**TotalCount** | **int64** |  | 
**HasMore** | **bool** |  | 
**NextOffset** | Pointer to **NullableInt64** |  | [optional] 
**PageIncludesLatestAvailable** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewIndicatorPaginationInfo

`func NewIndicatorPaginationInfo(offset int64, returnedCount int64, totalCount int64, hasMore bool, ) *IndicatorPaginationInfo`

NewIndicatorPaginationInfo instantiates a new IndicatorPaginationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicatorPaginationInfoWithDefaults

`func NewIndicatorPaginationInfoWithDefaults() *IndicatorPaginationInfo`

NewIndicatorPaginationInfoWithDefaults instantiates a new IndicatorPaginationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *IndicatorPaginationInfo) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *IndicatorPaginationInfo) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *IndicatorPaginationInfo) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *IndicatorPaginationInfo) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *IndicatorPaginationInfo) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *IndicatorPaginationInfo) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOffset

`func (o *IndicatorPaginationInfo) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *IndicatorPaginationInfo) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *IndicatorPaginationInfo) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetReturnedCount

`func (o *IndicatorPaginationInfo) GetReturnedCount() int64`

GetReturnedCount returns the ReturnedCount field if non-nil, zero value otherwise.

### GetReturnedCountOk

`func (o *IndicatorPaginationInfo) GetReturnedCountOk() (*int64, bool)`

GetReturnedCountOk returns a tuple with the ReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedCount

`func (o *IndicatorPaginationInfo) SetReturnedCount(v int64)`

SetReturnedCount sets ReturnedCount field to given value.


### GetTotalCount

`func (o *IndicatorPaginationInfo) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *IndicatorPaginationInfo) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *IndicatorPaginationInfo) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetHasMore

`func (o *IndicatorPaginationInfo) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *IndicatorPaginationInfo) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *IndicatorPaginationInfo) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNextOffset

`func (o *IndicatorPaginationInfo) GetNextOffset() int64`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *IndicatorPaginationInfo) GetNextOffsetOk() (*int64, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *IndicatorPaginationInfo) SetNextOffset(v int64)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *IndicatorPaginationInfo) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### SetNextOffsetNil

`func (o *IndicatorPaginationInfo) SetNextOffsetNil(b bool)`

 SetNextOffsetNil sets the value for NextOffset to be an explicit nil

### UnsetNextOffset
`func (o *IndicatorPaginationInfo) UnsetNextOffset()`

UnsetNextOffset ensures that no value is present for NextOffset, not even an explicit nil
### GetPageIncludesLatestAvailable

`func (o *IndicatorPaginationInfo) GetPageIncludesLatestAvailable() bool`

GetPageIncludesLatestAvailable returns the PageIncludesLatestAvailable field if non-nil, zero value otherwise.

### GetPageIncludesLatestAvailableOk

`func (o *IndicatorPaginationInfo) GetPageIncludesLatestAvailableOk() (*bool, bool)`

GetPageIncludesLatestAvailableOk returns a tuple with the PageIncludesLatestAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIncludesLatestAvailable

`func (o *IndicatorPaginationInfo) SetPageIncludesLatestAvailable(v bool)`

SetPageIncludesLatestAvailable sets PageIncludesLatestAvailable field to given value.

### HasPageIncludesLatestAvailable

`func (o *IndicatorPaginationInfo) HasPageIncludesLatestAvailable() bool`

HasPageIncludesLatestAvailable returns a boolean if a field has been set.

### SetPageIncludesLatestAvailableNil

`func (o *IndicatorPaginationInfo) SetPageIncludesLatestAvailableNil(b bool)`

 SetPageIncludesLatestAvailableNil sets the value for PageIncludesLatestAvailable to be an explicit nil

### UnsetPageIncludesLatestAvailable
`func (o *IndicatorPaginationInfo) UnsetPageIncludesLatestAvailable()`

UnsetPageIncludesLatestAvailable ensures that no value is present for PageIncludesLatestAvailable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


