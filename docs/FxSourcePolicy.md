# FxSourcePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicDataPolicy** | Pointer to **NullableString** |  | [optional] 
**RequestTimeUpstreamFetches** | Pointer to **NullableBool** | Whether serving a request contacts the upstream publisher. Always false: responses come from FXMacroData&#39;s collected dataset. | [optional] 
**ValueField** | Pointer to **NullableString** |  | [optional] 
**LegacyDailyForexValueField** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFxSourcePolicy

`func NewFxSourcePolicy() *FxSourcePolicy`

NewFxSourcePolicy instantiates a new FxSourcePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxSourcePolicyWithDefaults

`func NewFxSourcePolicyWithDefaults() *FxSourcePolicy`

NewFxSourcePolicyWithDefaults instantiates a new FxSourcePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicDataPolicy

`func (o *FxSourcePolicy) GetPublicDataPolicy() string`

GetPublicDataPolicy returns the PublicDataPolicy field if non-nil, zero value otherwise.

### GetPublicDataPolicyOk

`func (o *FxSourcePolicy) GetPublicDataPolicyOk() (*string, bool)`

GetPublicDataPolicyOk returns a tuple with the PublicDataPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDataPolicy

`func (o *FxSourcePolicy) SetPublicDataPolicy(v string)`

SetPublicDataPolicy sets PublicDataPolicy field to given value.

### HasPublicDataPolicy

`func (o *FxSourcePolicy) HasPublicDataPolicy() bool`

HasPublicDataPolicy returns a boolean if a field has been set.

### SetPublicDataPolicyNil

`func (o *FxSourcePolicy) SetPublicDataPolicyNil(b bool)`

 SetPublicDataPolicyNil sets the value for PublicDataPolicy to be an explicit nil

### UnsetPublicDataPolicy
`func (o *FxSourcePolicy) UnsetPublicDataPolicy()`

UnsetPublicDataPolicy ensures that no value is present for PublicDataPolicy, not even an explicit nil
### GetRequestTimeUpstreamFetches

`func (o *FxSourcePolicy) GetRequestTimeUpstreamFetches() bool`

GetRequestTimeUpstreamFetches returns the RequestTimeUpstreamFetches field if non-nil, zero value otherwise.

### GetRequestTimeUpstreamFetchesOk

`func (o *FxSourcePolicy) GetRequestTimeUpstreamFetchesOk() (*bool, bool)`

GetRequestTimeUpstreamFetchesOk returns a tuple with the RequestTimeUpstreamFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimeUpstreamFetches

`func (o *FxSourcePolicy) SetRequestTimeUpstreamFetches(v bool)`

SetRequestTimeUpstreamFetches sets RequestTimeUpstreamFetches field to given value.

### HasRequestTimeUpstreamFetches

`func (o *FxSourcePolicy) HasRequestTimeUpstreamFetches() bool`

HasRequestTimeUpstreamFetches returns a boolean if a field has been set.

### SetRequestTimeUpstreamFetchesNil

`func (o *FxSourcePolicy) SetRequestTimeUpstreamFetchesNil(b bool)`

 SetRequestTimeUpstreamFetchesNil sets the value for RequestTimeUpstreamFetches to be an explicit nil

### UnsetRequestTimeUpstreamFetches
`func (o *FxSourcePolicy) UnsetRequestTimeUpstreamFetches()`

UnsetRequestTimeUpstreamFetches ensures that no value is present for RequestTimeUpstreamFetches, not even an explicit nil
### GetValueField

`func (o *FxSourcePolicy) GetValueField() string`

GetValueField returns the ValueField field if non-nil, zero value otherwise.

### GetValueFieldOk

`func (o *FxSourcePolicy) GetValueFieldOk() (*string, bool)`

GetValueFieldOk returns a tuple with the ValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueField

`func (o *FxSourcePolicy) SetValueField(v string)`

SetValueField sets ValueField field to given value.

### HasValueField

`func (o *FxSourcePolicy) HasValueField() bool`

HasValueField returns a boolean if a field has been set.

### SetValueFieldNil

`func (o *FxSourcePolicy) SetValueFieldNil(b bool)`

 SetValueFieldNil sets the value for ValueField to be an explicit nil

### UnsetValueField
`func (o *FxSourcePolicy) UnsetValueField()`

UnsetValueField ensures that no value is present for ValueField, not even an explicit nil
### GetLegacyDailyForexValueField

`func (o *FxSourcePolicy) GetLegacyDailyForexValueField() string`

GetLegacyDailyForexValueField returns the LegacyDailyForexValueField field if non-nil, zero value otherwise.

### GetLegacyDailyForexValueFieldOk

`func (o *FxSourcePolicy) GetLegacyDailyForexValueFieldOk() (*string, bool)`

GetLegacyDailyForexValueFieldOk returns a tuple with the LegacyDailyForexValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyDailyForexValueField

`func (o *FxSourcePolicy) SetLegacyDailyForexValueField(v string)`

SetLegacyDailyForexValueField sets LegacyDailyForexValueField field to given value.

### HasLegacyDailyForexValueField

`func (o *FxSourcePolicy) HasLegacyDailyForexValueField() bool`

HasLegacyDailyForexValueField returns a boolean if a field has been set.

### SetLegacyDailyForexValueFieldNil

`func (o *FxSourcePolicy) SetLegacyDailyForexValueFieldNil(b bool)`

 SetLegacyDailyForexValueFieldNil sets the value for LegacyDailyForexValueField to be an explicit nil

### UnsetLegacyDailyForexValueField
`func (o *FxSourcePolicy) UnsetLegacyDailyForexValueField()`

UnsetLegacyDailyForexValueField ensures that no value is present for LegacyDailyForexValueField, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


