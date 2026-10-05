# CommoditiesLatestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**AsOf** | Pointer to **NullableString** |  | [optional] 
**Count** | Pointer to **NullableInt64** |  | [optional] 
**Data** | Pointer to [**[]CommodityLatestItem**](CommodityLatestItem.md) |  | [optional] 

## Methods

### NewCommoditiesLatestResponse

`func NewCommoditiesLatestResponse() *CommoditiesLatestResponse`

NewCommoditiesLatestResponse instantiates a new CommoditiesLatestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommoditiesLatestResponseWithDefaults

`func NewCommoditiesLatestResponseWithDefaults() *CommoditiesLatestResponse`

NewCommoditiesLatestResponseWithDefaults instantiates a new CommoditiesLatestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CommoditiesLatestResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CommoditiesLatestResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CommoditiesLatestResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CommoditiesLatestResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *CommoditiesLatestResponse) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CommoditiesLatestResponse) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetSource

`func (o *CommoditiesLatestResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CommoditiesLatestResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CommoditiesLatestResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CommoditiesLatestResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *CommoditiesLatestResponse) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *CommoditiesLatestResponse) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetAsOf

`func (o *CommoditiesLatestResponse) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *CommoditiesLatestResponse) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *CommoditiesLatestResponse) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *CommoditiesLatestResponse) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *CommoditiesLatestResponse) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *CommoditiesLatestResponse) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetCount

`func (o *CommoditiesLatestResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CommoditiesLatestResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CommoditiesLatestResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *CommoditiesLatestResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *CommoditiesLatestResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *CommoditiesLatestResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetData

`func (o *CommoditiesLatestResponse) GetData() []CommodityLatestItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CommoditiesLatestResponse) GetDataOk() (*[]CommodityLatestItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CommoditiesLatestResponse) SetData(v []CommodityLatestItem)`

SetData sets Data field to given value.

### HasData

`func (o *CommoditiesLatestResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *CommoditiesLatestResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *CommoditiesLatestResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


