# ForexProvenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Publisher** | Pointer to **NullableString** |  | [optional] 
**PublisherUrl** | Pointer to **NullableString** |  | [optional] 
**Storage** | Pointer to **NullableString** | Data provider label for the served series. Always &#x60;FXMacroData&#x60;. | [optional] 
**ServedBy** | Pointer to **NullableString** | FXMacroData endpoint that serves this series. | [optional] 
**TimestampField** | Pointer to **NullableString** |  | [optional] 
**ValueField** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewForexProvenance

`func NewForexProvenance() *ForexProvenance`

NewForexProvenance instantiates a new ForexProvenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForexProvenanceWithDefaults

`func NewForexProvenanceWithDefaults() *ForexProvenance`

NewForexProvenanceWithDefaults instantiates a new ForexProvenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublisher

`func (o *ForexProvenance) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *ForexProvenance) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *ForexProvenance) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *ForexProvenance) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisherNil

`func (o *ForexProvenance) SetPublisherNil(b bool)`

 SetPublisherNil sets the value for Publisher to be an explicit nil

### UnsetPublisher
`func (o *ForexProvenance) UnsetPublisher()`

UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
### GetPublisherUrl

`func (o *ForexProvenance) GetPublisherUrl() string`

GetPublisherUrl returns the PublisherUrl field if non-nil, zero value otherwise.

### GetPublisherUrlOk

`func (o *ForexProvenance) GetPublisherUrlOk() (*string, bool)`

GetPublisherUrlOk returns a tuple with the PublisherUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherUrl

`func (o *ForexProvenance) SetPublisherUrl(v string)`

SetPublisherUrl sets PublisherUrl field to given value.

### HasPublisherUrl

`func (o *ForexProvenance) HasPublisherUrl() bool`

HasPublisherUrl returns a boolean if a field has been set.

### SetPublisherUrlNil

`func (o *ForexProvenance) SetPublisherUrlNil(b bool)`

 SetPublisherUrlNil sets the value for PublisherUrl to be an explicit nil

### UnsetPublisherUrl
`func (o *ForexProvenance) UnsetPublisherUrl()`

UnsetPublisherUrl ensures that no value is present for PublisherUrl, not even an explicit nil
### GetStorage

`func (o *ForexProvenance) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ForexProvenance) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ForexProvenance) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ForexProvenance) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *ForexProvenance) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *ForexProvenance) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetServedBy

`func (o *ForexProvenance) GetServedBy() string`

GetServedBy returns the ServedBy field if non-nil, zero value otherwise.

### GetServedByOk

`func (o *ForexProvenance) GetServedByOk() (*string, bool)`

GetServedByOk returns a tuple with the ServedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedBy

`func (o *ForexProvenance) SetServedBy(v string)`

SetServedBy sets ServedBy field to given value.

### HasServedBy

`func (o *ForexProvenance) HasServedBy() bool`

HasServedBy returns a boolean if a field has been set.

### SetServedByNil

`func (o *ForexProvenance) SetServedByNil(b bool)`

 SetServedByNil sets the value for ServedBy to be an explicit nil

### UnsetServedBy
`func (o *ForexProvenance) UnsetServedBy()`

UnsetServedBy ensures that no value is present for ServedBy, not even an explicit nil
### GetTimestampField

`func (o *ForexProvenance) GetTimestampField() string`

GetTimestampField returns the TimestampField field if non-nil, zero value otherwise.

### GetTimestampFieldOk

`func (o *ForexProvenance) GetTimestampFieldOk() (*string, bool)`

GetTimestampFieldOk returns a tuple with the TimestampField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampField

`func (o *ForexProvenance) SetTimestampField(v string)`

SetTimestampField sets TimestampField field to given value.

### HasTimestampField

`func (o *ForexProvenance) HasTimestampField() bool`

HasTimestampField returns a boolean if a field has been set.

### SetTimestampFieldNil

`func (o *ForexProvenance) SetTimestampFieldNil(b bool)`

 SetTimestampFieldNil sets the value for TimestampField to be an explicit nil

### UnsetTimestampField
`func (o *ForexProvenance) UnsetTimestampField()`

UnsetTimestampField ensures that no value is present for TimestampField, not even an explicit nil
### GetValueField

`func (o *ForexProvenance) GetValueField() string`

GetValueField returns the ValueField field if non-nil, zero value otherwise.

### GetValueFieldOk

`func (o *ForexProvenance) GetValueFieldOk() (*string, bool)`

GetValueFieldOk returns a tuple with the ValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueField

`func (o *ForexProvenance) SetValueField(v string)`

SetValueField sets ValueField field to given value.

### HasValueField

`func (o *ForexProvenance) HasValueField() bool`

HasValueField returns a boolean if a field has been set.

### SetValueFieldNil

`func (o *ForexProvenance) SetValueFieldNil(b bool)`

 SetValueFieldNil sets the value for ValueField to be an explicit nil

### UnsetValueField
`func (o *ForexProvenance) UnsetValueField()`

UnsetValueField ensures that no value is present for ValueField, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


