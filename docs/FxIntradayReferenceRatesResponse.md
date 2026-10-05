# FxIntradayReferenceRatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pair** | **string** |  | 
**StartTime** | **string** |  | 
**EndTime** | **string** |  | 
**Data** | [**[]FxIntradayReferenceRatePoint**](FxIntradayReferenceRatePoint.md) |  | 

## Methods

### NewFxIntradayReferenceRatesResponse

`func NewFxIntradayReferenceRatesResponse(pair string, startTime string, endTime string, data []FxIntradayReferenceRatePoint, ) *FxIntradayReferenceRatesResponse`

NewFxIntradayReferenceRatesResponse instantiates a new FxIntradayReferenceRatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxIntradayReferenceRatesResponseWithDefaults

`func NewFxIntradayReferenceRatesResponseWithDefaults() *FxIntradayReferenceRatesResponse`

NewFxIntradayReferenceRatesResponseWithDefaults instantiates a new FxIntradayReferenceRatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPair

`func (o *FxIntradayReferenceRatesResponse) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *FxIntradayReferenceRatesResponse) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *FxIntradayReferenceRatesResponse) SetPair(v string)`

SetPair sets Pair field to given value.


### GetStartTime

`func (o *FxIntradayReferenceRatesResponse) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *FxIntradayReferenceRatesResponse) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *FxIntradayReferenceRatesResponse) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *FxIntradayReferenceRatesResponse) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *FxIntradayReferenceRatesResponse) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *FxIntradayReferenceRatesResponse) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetData

`func (o *FxIntradayReferenceRatesResponse) GetData() []FxIntradayReferenceRatePoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FxIntradayReferenceRatesResponse) GetDataOk() (*[]FxIntradayReferenceRatePoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FxIntradayReferenceRatesResponse) SetData(v []FxIntradayReferenceRatePoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


