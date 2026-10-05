# FxSourcesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourcePolicy** | [**FxSourcePolicy**](FxSourcePolicy.md) |  | 
**Sources** | [**[]FxSource**](FxSource.md) |  | 

## Methods

### NewFxSourcesResponse

`func NewFxSourcesResponse(sourcePolicy FxSourcePolicy, sources []FxSource, ) *FxSourcesResponse`

NewFxSourcesResponse instantiates a new FxSourcesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxSourcesResponseWithDefaults

`func NewFxSourcesResponseWithDefaults() *FxSourcesResponse`

NewFxSourcesResponseWithDefaults instantiates a new FxSourcesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourcePolicy

`func (o *FxSourcesResponse) GetSourcePolicy() FxSourcePolicy`

GetSourcePolicy returns the SourcePolicy field if non-nil, zero value otherwise.

### GetSourcePolicyOk

`func (o *FxSourcesResponse) GetSourcePolicyOk() (*FxSourcePolicy, bool)`

GetSourcePolicyOk returns a tuple with the SourcePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePolicy

`func (o *FxSourcesResponse) SetSourcePolicy(v FxSourcePolicy)`

SetSourcePolicy sets SourcePolicy field to given value.


### GetSources

`func (o *FxSourcesResponse) GetSources() []FxSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *FxSourcesResponse) GetSourcesOk() (*[]FxSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *FxSourcesResponse) SetSources(v []FxSource)`

SetSources sets Sources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


