# CBTargetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Source** | **string** |  | 
**Current** | [**CBTargetEntry**](CBTargetEntry.md) |  | 
**History** | [**[]CBTargetEntry**](CBTargetEntry.md) |  | 

## Methods

### NewCBTargetInfo

`func NewCBTargetInfo(description string, source string, current CBTargetEntry, history []CBTargetEntry, ) *CBTargetInfo`

NewCBTargetInfo instantiates a new CBTargetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCBTargetInfoWithDefaults

`func NewCBTargetInfoWithDefaults() *CBTargetInfo`

NewCBTargetInfoWithDefaults instantiates a new CBTargetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CBTargetInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CBTargetInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CBTargetInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSource

`func (o *CBTargetInfo) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CBTargetInfo) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CBTargetInfo) SetSource(v string)`

SetSource sets Source field to given value.


### GetCurrent

`func (o *CBTargetInfo) GetCurrent() CBTargetEntry`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *CBTargetInfo) GetCurrentOk() (*CBTargetEntry, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *CBTargetInfo) SetCurrent(v CBTargetEntry)`

SetCurrent sets Current field to given value.


### GetHistory

`func (o *CBTargetInfo) GetHistory() []CBTargetEntry`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *CBTargetInfo) GetHistoryOk() (*[]CBTargetEntry, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *CBTargetInfo) SetHistory(v []CBTargetEntry)`

SetHistory sets History field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


