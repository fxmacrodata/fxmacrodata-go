# SourceProvenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Publisher** | Pointer to **NullableString** |  | [optional] 
**PublisherUrl** | Pointer to **NullableString** |  | [optional] 
**Storage** | Pointer to **NullableString** | Data provider label for the served series. Always &#x60;FXMacroData&#x60;. | [optional] 
**SourceStore** | Pointer to **NullableString** | Source series reference, when reported. | [optional] 
**ServedBy** | Pointer to **NullableString** | FXMacroData endpoint that serves this series. | [optional] 
**TimestampField** | Pointer to **NullableString** |  | [optional] 
**ValueField** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSourceProvenance

`func NewSourceProvenance() *SourceProvenance`

NewSourceProvenance instantiates a new SourceProvenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceProvenanceWithDefaults

`func NewSourceProvenanceWithDefaults() *SourceProvenance`

NewSourceProvenanceWithDefaults instantiates a new SourceProvenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublisher

`func (o *SourceProvenance) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *SourceProvenance) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *SourceProvenance) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *SourceProvenance) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisherNil

`func (o *SourceProvenance) SetPublisherNil(b bool)`

 SetPublisherNil sets the value for Publisher to be an explicit nil

### UnsetPublisher
`func (o *SourceProvenance) UnsetPublisher()`

UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
### GetPublisherUrl

`func (o *SourceProvenance) GetPublisherUrl() string`

GetPublisherUrl returns the PublisherUrl field if non-nil, zero value otherwise.

### GetPublisherUrlOk

`func (o *SourceProvenance) GetPublisherUrlOk() (*string, bool)`

GetPublisherUrlOk returns a tuple with the PublisherUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherUrl

`func (o *SourceProvenance) SetPublisherUrl(v string)`

SetPublisherUrl sets PublisherUrl field to given value.

### HasPublisherUrl

`func (o *SourceProvenance) HasPublisherUrl() bool`

HasPublisherUrl returns a boolean if a field has been set.

### SetPublisherUrlNil

`func (o *SourceProvenance) SetPublisherUrlNil(b bool)`

 SetPublisherUrlNil sets the value for PublisherUrl to be an explicit nil

### UnsetPublisherUrl
`func (o *SourceProvenance) UnsetPublisherUrl()`

UnsetPublisherUrl ensures that no value is present for PublisherUrl, not even an explicit nil
### GetStorage

`func (o *SourceProvenance) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *SourceProvenance) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *SourceProvenance) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *SourceProvenance) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *SourceProvenance) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *SourceProvenance) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetSourceStore

`func (o *SourceProvenance) GetSourceStore() string`

GetSourceStore returns the SourceStore field if non-nil, zero value otherwise.

### GetSourceStoreOk

`func (o *SourceProvenance) GetSourceStoreOk() (*string, bool)`

GetSourceStoreOk returns a tuple with the SourceStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStore

`func (o *SourceProvenance) SetSourceStore(v string)`

SetSourceStore sets SourceStore field to given value.

### HasSourceStore

`func (o *SourceProvenance) HasSourceStore() bool`

HasSourceStore returns a boolean if a field has been set.

### SetSourceStoreNil

`func (o *SourceProvenance) SetSourceStoreNil(b bool)`

 SetSourceStoreNil sets the value for SourceStore to be an explicit nil

### UnsetSourceStore
`func (o *SourceProvenance) UnsetSourceStore()`

UnsetSourceStore ensures that no value is present for SourceStore, not even an explicit nil
### GetServedBy

`func (o *SourceProvenance) GetServedBy() string`

GetServedBy returns the ServedBy field if non-nil, zero value otherwise.

### GetServedByOk

`func (o *SourceProvenance) GetServedByOk() (*string, bool)`

GetServedByOk returns a tuple with the ServedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedBy

`func (o *SourceProvenance) SetServedBy(v string)`

SetServedBy sets ServedBy field to given value.

### HasServedBy

`func (o *SourceProvenance) HasServedBy() bool`

HasServedBy returns a boolean if a field has been set.

### SetServedByNil

`func (o *SourceProvenance) SetServedByNil(b bool)`

 SetServedByNil sets the value for ServedBy to be an explicit nil

### UnsetServedBy
`func (o *SourceProvenance) UnsetServedBy()`

UnsetServedBy ensures that no value is present for ServedBy, not even an explicit nil
### GetTimestampField

`func (o *SourceProvenance) GetTimestampField() string`

GetTimestampField returns the TimestampField field if non-nil, zero value otherwise.

### GetTimestampFieldOk

`func (o *SourceProvenance) GetTimestampFieldOk() (*string, bool)`

GetTimestampFieldOk returns a tuple with the TimestampField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampField

`func (o *SourceProvenance) SetTimestampField(v string)`

SetTimestampField sets TimestampField field to given value.

### HasTimestampField

`func (o *SourceProvenance) HasTimestampField() bool`

HasTimestampField returns a boolean if a field has been set.

### SetTimestampFieldNil

`func (o *SourceProvenance) SetTimestampFieldNil(b bool)`

 SetTimestampFieldNil sets the value for TimestampField to be an explicit nil

### UnsetTimestampField
`func (o *SourceProvenance) UnsetTimestampField()`

UnsetTimestampField ensures that no value is present for TimestampField, not even an explicit nil
### GetValueField

`func (o *SourceProvenance) GetValueField() string`

GetValueField returns the ValueField field if non-nil, zero value otherwise.

### GetValueFieldOk

`func (o *SourceProvenance) GetValueFieldOk() (*string, bool)`

GetValueFieldOk returns a tuple with the ValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueField

`func (o *SourceProvenance) SetValueField(v string)`

SetValueField sets ValueField field to given value.

### HasValueField

`func (o *SourceProvenance) HasValueField() bool`

HasValueField returns a boolean if a field has been set.

### SetValueFieldNil

`func (o *SourceProvenance) SetValueFieldNil(b bool)`

 SetValueFieldNil sets the value for ValueField to be an explicit nil

### UnsetValueField
`func (o *SourceProvenance) UnsetValueField()`

UnsetValueField ensures that no value is present for ValueField, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


