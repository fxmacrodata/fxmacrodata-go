# AnnouncementProvenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Publisher** | Pointer to **NullableString** |  | [optional] 
**PublisherUrl** | Pointer to **NullableString** |  | [optional] 
**SourceSeriesId** | Pointer to **NullableString** |  | [optional] 
**SourceSeriesName** | Pointer to **NullableString** |  | [optional] 
**SourceLocalName** | Pointer to **NullableString** |  | [optional] 
**SeasonalAdjustment** | Pointer to **NullableString** |  | [optional] 
**PriceBasis** | Pointer to **NullableString** |  | [optional] 
**Storage** | Pointer to **NullableString** | Data provider label for the served series. Always &#x60;FXMacroData&#x60;. | [optional] 
**ServedBy** | Pointer to **NullableString** | FXMacroData endpoint that serves this series. | [optional] 
**TimestampField** | Pointer to **NullableString** |  | [optional] 
**ValueField** | Pointer to **NullableString** |  | [optional] 
**IsProxy** | Pointer to **NullableBool** |  | [optional] 
**ProxyNote** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAnnouncementProvenance

`func NewAnnouncementProvenance() *AnnouncementProvenance`

NewAnnouncementProvenance instantiates a new AnnouncementProvenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementProvenanceWithDefaults

`func NewAnnouncementProvenanceWithDefaults() *AnnouncementProvenance`

NewAnnouncementProvenanceWithDefaults instantiates a new AnnouncementProvenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublisher

`func (o *AnnouncementProvenance) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *AnnouncementProvenance) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *AnnouncementProvenance) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *AnnouncementProvenance) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisherNil

`func (o *AnnouncementProvenance) SetPublisherNil(b bool)`

 SetPublisherNil sets the value for Publisher to be an explicit nil

### UnsetPublisher
`func (o *AnnouncementProvenance) UnsetPublisher()`

UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
### GetPublisherUrl

`func (o *AnnouncementProvenance) GetPublisherUrl() string`

GetPublisherUrl returns the PublisherUrl field if non-nil, zero value otherwise.

### GetPublisherUrlOk

`func (o *AnnouncementProvenance) GetPublisherUrlOk() (*string, bool)`

GetPublisherUrlOk returns a tuple with the PublisherUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherUrl

`func (o *AnnouncementProvenance) SetPublisherUrl(v string)`

SetPublisherUrl sets PublisherUrl field to given value.

### HasPublisherUrl

`func (o *AnnouncementProvenance) HasPublisherUrl() bool`

HasPublisherUrl returns a boolean if a field has been set.

### SetPublisherUrlNil

`func (o *AnnouncementProvenance) SetPublisherUrlNil(b bool)`

 SetPublisherUrlNil sets the value for PublisherUrl to be an explicit nil

### UnsetPublisherUrl
`func (o *AnnouncementProvenance) UnsetPublisherUrl()`

UnsetPublisherUrl ensures that no value is present for PublisherUrl, not even an explicit nil
### GetSourceSeriesId

`func (o *AnnouncementProvenance) GetSourceSeriesId() string`

GetSourceSeriesId returns the SourceSeriesId field if non-nil, zero value otherwise.

### GetSourceSeriesIdOk

`func (o *AnnouncementProvenance) GetSourceSeriesIdOk() (*string, bool)`

GetSourceSeriesIdOk returns a tuple with the SourceSeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSeriesId

`func (o *AnnouncementProvenance) SetSourceSeriesId(v string)`

SetSourceSeriesId sets SourceSeriesId field to given value.

### HasSourceSeriesId

`func (o *AnnouncementProvenance) HasSourceSeriesId() bool`

HasSourceSeriesId returns a boolean if a field has been set.

### SetSourceSeriesIdNil

`func (o *AnnouncementProvenance) SetSourceSeriesIdNil(b bool)`

 SetSourceSeriesIdNil sets the value for SourceSeriesId to be an explicit nil

### UnsetSourceSeriesId
`func (o *AnnouncementProvenance) UnsetSourceSeriesId()`

UnsetSourceSeriesId ensures that no value is present for SourceSeriesId, not even an explicit nil
### GetSourceSeriesName

`func (o *AnnouncementProvenance) GetSourceSeriesName() string`

GetSourceSeriesName returns the SourceSeriesName field if non-nil, zero value otherwise.

### GetSourceSeriesNameOk

`func (o *AnnouncementProvenance) GetSourceSeriesNameOk() (*string, bool)`

GetSourceSeriesNameOk returns a tuple with the SourceSeriesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSeriesName

`func (o *AnnouncementProvenance) SetSourceSeriesName(v string)`

SetSourceSeriesName sets SourceSeriesName field to given value.

### HasSourceSeriesName

`func (o *AnnouncementProvenance) HasSourceSeriesName() bool`

HasSourceSeriesName returns a boolean if a field has been set.

### SetSourceSeriesNameNil

`func (o *AnnouncementProvenance) SetSourceSeriesNameNil(b bool)`

 SetSourceSeriesNameNil sets the value for SourceSeriesName to be an explicit nil

### UnsetSourceSeriesName
`func (o *AnnouncementProvenance) UnsetSourceSeriesName()`

UnsetSourceSeriesName ensures that no value is present for SourceSeriesName, not even an explicit nil
### GetSourceLocalName

`func (o *AnnouncementProvenance) GetSourceLocalName() string`

GetSourceLocalName returns the SourceLocalName field if non-nil, zero value otherwise.

### GetSourceLocalNameOk

`func (o *AnnouncementProvenance) GetSourceLocalNameOk() (*string, bool)`

GetSourceLocalNameOk returns a tuple with the SourceLocalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLocalName

`func (o *AnnouncementProvenance) SetSourceLocalName(v string)`

SetSourceLocalName sets SourceLocalName field to given value.

### HasSourceLocalName

`func (o *AnnouncementProvenance) HasSourceLocalName() bool`

HasSourceLocalName returns a boolean if a field has been set.

### SetSourceLocalNameNil

`func (o *AnnouncementProvenance) SetSourceLocalNameNil(b bool)`

 SetSourceLocalNameNil sets the value for SourceLocalName to be an explicit nil

### UnsetSourceLocalName
`func (o *AnnouncementProvenance) UnsetSourceLocalName()`

UnsetSourceLocalName ensures that no value is present for SourceLocalName, not even an explicit nil
### GetSeasonalAdjustment

`func (o *AnnouncementProvenance) GetSeasonalAdjustment() string`

GetSeasonalAdjustment returns the SeasonalAdjustment field if non-nil, zero value otherwise.

### GetSeasonalAdjustmentOk

`func (o *AnnouncementProvenance) GetSeasonalAdjustmentOk() (*string, bool)`

GetSeasonalAdjustmentOk returns a tuple with the SeasonalAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonalAdjustment

`func (o *AnnouncementProvenance) SetSeasonalAdjustment(v string)`

SetSeasonalAdjustment sets SeasonalAdjustment field to given value.

### HasSeasonalAdjustment

`func (o *AnnouncementProvenance) HasSeasonalAdjustment() bool`

HasSeasonalAdjustment returns a boolean if a field has been set.

### SetSeasonalAdjustmentNil

`func (o *AnnouncementProvenance) SetSeasonalAdjustmentNil(b bool)`

 SetSeasonalAdjustmentNil sets the value for SeasonalAdjustment to be an explicit nil

### UnsetSeasonalAdjustment
`func (o *AnnouncementProvenance) UnsetSeasonalAdjustment()`

UnsetSeasonalAdjustment ensures that no value is present for SeasonalAdjustment, not even an explicit nil
### GetPriceBasis

`func (o *AnnouncementProvenance) GetPriceBasis() string`

GetPriceBasis returns the PriceBasis field if non-nil, zero value otherwise.

### GetPriceBasisOk

`func (o *AnnouncementProvenance) GetPriceBasisOk() (*string, bool)`

GetPriceBasisOk returns a tuple with the PriceBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceBasis

`func (o *AnnouncementProvenance) SetPriceBasis(v string)`

SetPriceBasis sets PriceBasis field to given value.

### HasPriceBasis

`func (o *AnnouncementProvenance) HasPriceBasis() bool`

HasPriceBasis returns a boolean if a field has been set.

### SetPriceBasisNil

`func (o *AnnouncementProvenance) SetPriceBasisNil(b bool)`

 SetPriceBasisNil sets the value for PriceBasis to be an explicit nil

### UnsetPriceBasis
`func (o *AnnouncementProvenance) UnsetPriceBasis()`

UnsetPriceBasis ensures that no value is present for PriceBasis, not even an explicit nil
### GetStorage

`func (o *AnnouncementProvenance) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *AnnouncementProvenance) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *AnnouncementProvenance) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *AnnouncementProvenance) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *AnnouncementProvenance) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *AnnouncementProvenance) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetServedBy

`func (o *AnnouncementProvenance) GetServedBy() string`

GetServedBy returns the ServedBy field if non-nil, zero value otherwise.

### GetServedByOk

`func (o *AnnouncementProvenance) GetServedByOk() (*string, bool)`

GetServedByOk returns a tuple with the ServedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedBy

`func (o *AnnouncementProvenance) SetServedBy(v string)`

SetServedBy sets ServedBy field to given value.

### HasServedBy

`func (o *AnnouncementProvenance) HasServedBy() bool`

HasServedBy returns a boolean if a field has been set.

### SetServedByNil

`func (o *AnnouncementProvenance) SetServedByNil(b bool)`

 SetServedByNil sets the value for ServedBy to be an explicit nil

### UnsetServedBy
`func (o *AnnouncementProvenance) UnsetServedBy()`

UnsetServedBy ensures that no value is present for ServedBy, not even an explicit nil
### GetTimestampField

`func (o *AnnouncementProvenance) GetTimestampField() string`

GetTimestampField returns the TimestampField field if non-nil, zero value otherwise.

### GetTimestampFieldOk

`func (o *AnnouncementProvenance) GetTimestampFieldOk() (*string, bool)`

GetTimestampFieldOk returns a tuple with the TimestampField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampField

`func (o *AnnouncementProvenance) SetTimestampField(v string)`

SetTimestampField sets TimestampField field to given value.

### HasTimestampField

`func (o *AnnouncementProvenance) HasTimestampField() bool`

HasTimestampField returns a boolean if a field has been set.

### SetTimestampFieldNil

`func (o *AnnouncementProvenance) SetTimestampFieldNil(b bool)`

 SetTimestampFieldNil sets the value for TimestampField to be an explicit nil

### UnsetTimestampField
`func (o *AnnouncementProvenance) UnsetTimestampField()`

UnsetTimestampField ensures that no value is present for TimestampField, not even an explicit nil
### GetValueField

`func (o *AnnouncementProvenance) GetValueField() string`

GetValueField returns the ValueField field if non-nil, zero value otherwise.

### GetValueFieldOk

`func (o *AnnouncementProvenance) GetValueFieldOk() (*string, bool)`

GetValueFieldOk returns a tuple with the ValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueField

`func (o *AnnouncementProvenance) SetValueField(v string)`

SetValueField sets ValueField field to given value.

### HasValueField

`func (o *AnnouncementProvenance) HasValueField() bool`

HasValueField returns a boolean if a field has been set.

### SetValueFieldNil

`func (o *AnnouncementProvenance) SetValueFieldNil(b bool)`

 SetValueFieldNil sets the value for ValueField to be an explicit nil

### UnsetValueField
`func (o *AnnouncementProvenance) UnsetValueField()`

UnsetValueField ensures that no value is present for ValueField, not even an explicit nil
### GetIsProxy

`func (o *AnnouncementProvenance) GetIsProxy() bool`

GetIsProxy returns the IsProxy field if non-nil, zero value otherwise.

### GetIsProxyOk

`func (o *AnnouncementProvenance) GetIsProxyOk() (*bool, bool)`

GetIsProxyOk returns a tuple with the IsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxy

`func (o *AnnouncementProvenance) SetIsProxy(v bool)`

SetIsProxy sets IsProxy field to given value.

### HasIsProxy

`func (o *AnnouncementProvenance) HasIsProxy() bool`

HasIsProxy returns a boolean if a field has been set.

### SetIsProxyNil

`func (o *AnnouncementProvenance) SetIsProxyNil(b bool)`

 SetIsProxyNil sets the value for IsProxy to be an explicit nil

### UnsetIsProxy
`func (o *AnnouncementProvenance) UnsetIsProxy()`

UnsetIsProxy ensures that no value is present for IsProxy, not even an explicit nil
### GetProxyNote

`func (o *AnnouncementProvenance) GetProxyNote() string`

GetProxyNote returns the ProxyNote field if non-nil, zero value otherwise.

### GetProxyNoteOk

`func (o *AnnouncementProvenance) GetProxyNoteOk() (*string, bool)`

GetProxyNoteOk returns a tuple with the ProxyNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyNote

`func (o *AnnouncementProvenance) SetProxyNote(v string)`

SetProxyNote sets ProxyNote field to given value.

### HasProxyNote

`func (o *AnnouncementProvenance) HasProxyNote() bool`

HasProxyNote returns a boolean if a field has been set.

### SetProxyNoteNil

`func (o *AnnouncementProvenance) SetProxyNoteNil(b bool)`

 SetProxyNoteNil sets the value for ProxyNote to be an explicit nil

### UnsetProxyNote
`func (o *AnnouncementProvenance) UnsetProxyNote()`

UnsetProxyNote ensures that no value is present for ProxyNote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


