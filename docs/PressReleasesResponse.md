# PressReleasesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceUrl** | Pointer to **NullableString** |  | [optional] 
**Limit** | Pointer to **NullableInt64** |  | [optional] 
**Offset** | Pointer to **NullableInt64** |  | [optional] 
**Count** | Pointer to **NullableInt64** |  | [optional] 
**Pagination** | Pointer to [**NullableIndicatorPaginationInfo**](IndicatorPaginationInfo.md) |  | [optional] 
**Data** | Pointer to [**[]PressRelease**](PressRelease.md) |  | [optional] 

## Methods

### NewPressReleasesResponse

`func NewPressReleasesResponse() *PressReleasesResponse`

NewPressReleasesResponse instantiates a new PressReleasesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPressReleasesResponseWithDefaults

`func NewPressReleasesResponseWithDefaults() *PressReleasesResponse`

NewPressReleasesResponseWithDefaults instantiates a new PressReleasesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PressReleasesResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PressReleasesResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PressReleasesResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PressReleasesResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *PressReleasesResponse) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PressReleasesResponse) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetSource

`func (o *PressReleasesResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PressReleasesResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PressReleasesResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PressReleasesResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *PressReleasesResponse) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *PressReleasesResponse) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceUrl

`func (o *PressReleasesResponse) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *PressReleasesResponse) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *PressReleasesResponse) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *PressReleasesResponse) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *PressReleasesResponse) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *PressReleasesResponse) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetLimit

`func (o *PressReleasesResponse) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PressReleasesResponse) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PressReleasesResponse) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PressReleasesResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *PressReleasesResponse) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *PressReleasesResponse) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOffset

`func (o *PressReleasesResponse) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PressReleasesResponse) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PressReleasesResponse) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PressReleasesResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *PressReleasesResponse) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *PressReleasesResponse) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetCount

`func (o *PressReleasesResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PressReleasesResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PressReleasesResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PressReleasesResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *PressReleasesResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *PressReleasesResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetPagination

`func (o *PressReleasesResponse) GetPagination() IndicatorPaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PressReleasesResponse) GetPaginationOk() (*IndicatorPaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PressReleasesResponse) SetPagination(v IndicatorPaginationInfo)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PressReleasesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### SetPaginationNil

`func (o *PressReleasesResponse) SetPaginationNil(b bool)`

 SetPaginationNil sets the value for Pagination to be an explicit nil

### UnsetPagination
`func (o *PressReleasesResponse) UnsetPagination()`

UnsetPagination ensures that no value is present for Pagination, not even an explicit nil
### GetData

`func (o *PressReleasesResponse) GetData() []PressRelease`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PressReleasesResponse) GetDataOk() (*[]PressRelease, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PressReleasesResponse) SetData(v []PressRelease)`

SetData sets Data field to given value.

### HasData

`func (o *PressReleasesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *PressReleasesResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *PressReleasesResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


