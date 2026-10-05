# LatestAnnouncementsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**Provenance** | Pointer to [**NullableLatestAnnouncementsProvenance**](LatestAnnouncementsProvenance.md) |  | [optional] 
**AsOf** | Pointer to **NullableString** |  | [optional] 
**Count** | Pointer to **NullableInt64** |  | [optional] 
**Data** | Pointer to [**[]LatestAnnouncementItem**](LatestAnnouncementItem.md) |  | [optional] 
**FreemiumDelay** | Pointer to [**NullableFreemiumDelayNotice**](FreemiumDelayNotice.md) |  | [optional] 

## Methods

### NewLatestAnnouncementsResponse

`func NewLatestAnnouncementsResponse() *LatestAnnouncementsResponse`

NewLatestAnnouncementsResponse instantiates a new LatestAnnouncementsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLatestAnnouncementsResponseWithDefaults

`func NewLatestAnnouncementsResponseWithDefaults() *LatestAnnouncementsResponse`

NewLatestAnnouncementsResponseWithDefaults instantiates a new LatestAnnouncementsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *LatestAnnouncementsResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LatestAnnouncementsResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LatestAnnouncementsResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *LatestAnnouncementsResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *LatestAnnouncementsResponse) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *LatestAnnouncementsResponse) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetSource

`func (o *LatestAnnouncementsResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LatestAnnouncementsResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LatestAnnouncementsResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LatestAnnouncementsResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *LatestAnnouncementsResponse) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *LatestAnnouncementsResponse) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetProvenance

`func (o *LatestAnnouncementsResponse) GetProvenance() LatestAnnouncementsProvenance`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *LatestAnnouncementsResponse) GetProvenanceOk() (*LatestAnnouncementsProvenance, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *LatestAnnouncementsResponse) SetProvenance(v LatestAnnouncementsProvenance)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *LatestAnnouncementsResponse) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### SetProvenanceNil

`func (o *LatestAnnouncementsResponse) SetProvenanceNil(b bool)`

 SetProvenanceNil sets the value for Provenance to be an explicit nil

### UnsetProvenance
`func (o *LatestAnnouncementsResponse) UnsetProvenance()`

UnsetProvenance ensures that no value is present for Provenance, not even an explicit nil
### GetAsOf

`func (o *LatestAnnouncementsResponse) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *LatestAnnouncementsResponse) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *LatestAnnouncementsResponse) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *LatestAnnouncementsResponse) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *LatestAnnouncementsResponse) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *LatestAnnouncementsResponse) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetCount

`func (o *LatestAnnouncementsResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LatestAnnouncementsResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LatestAnnouncementsResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *LatestAnnouncementsResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *LatestAnnouncementsResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *LatestAnnouncementsResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetData

`func (o *LatestAnnouncementsResponse) GetData() []LatestAnnouncementItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LatestAnnouncementsResponse) GetDataOk() (*[]LatestAnnouncementItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LatestAnnouncementsResponse) SetData(v []LatestAnnouncementItem)`

SetData sets Data field to given value.

### HasData

`func (o *LatestAnnouncementsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *LatestAnnouncementsResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *LatestAnnouncementsResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetFreemiumDelay

`func (o *LatestAnnouncementsResponse) GetFreemiumDelay() FreemiumDelayNotice`

GetFreemiumDelay returns the FreemiumDelay field if non-nil, zero value otherwise.

### GetFreemiumDelayOk

`func (o *LatestAnnouncementsResponse) GetFreemiumDelayOk() (*FreemiumDelayNotice, bool)`

GetFreemiumDelayOk returns a tuple with the FreemiumDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreemiumDelay

`func (o *LatestAnnouncementsResponse) SetFreemiumDelay(v FreemiumDelayNotice)`

SetFreemiumDelay sets FreemiumDelay field to given value.

### HasFreemiumDelay

`func (o *LatestAnnouncementsResponse) HasFreemiumDelay() bool`

HasFreemiumDelay returns a boolean if a field has been set.

### SetFreemiumDelayNil

`func (o *LatestAnnouncementsResponse) SetFreemiumDelayNil(b bool)`

 SetFreemiumDelayNil sets the value for FreemiumDelay to be an explicit nil

### UnsetFreemiumDelay
`func (o *LatestAnnouncementsResponse) UnsetFreemiumDelay()`

UnsetFreemiumDelay ensures that no value is present for FreemiumDelay, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


