# FxSourceLeg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pair** | Pointer to **NullableString** |  | [optional] 
**SourcePair** | Pointer to **NullableString** |  | [optional] 
**StoragePath** | Pointer to **NullableString** | Retained for compatibility. Always null. | [optional] 
**StoredSource** | Pointer to **NullableString** | Retained for compatibility. Always null. | [optional] 
**SourceProvider** | Pointer to **NullableString** |  | [optional] 
**SourceName** | Pointer to **NullableString** |  | [optional] 
**SourceUrl** | Pointer to **NullableString** |  | [optional] 
**SourceFrequency** | Pointer to **NullableString** |  | [optional] 
**SourceType** | Pointer to **NullableString** |  | [optional] 
**IsDerived** | Pointer to **NullableBool** |  | [optional] 
**DerivationMethod** | Pointer to **NullableString** |  | [optional] 
**AnchorCurrency** | Pointer to **NullableString** |  | [optional] 
**DataPolicy** | Pointer to **NullableString** |  | [optional] 
**LastUpdated** | Pointer to **NullableString** |  | [optional] 
**BackfillWindow** | Pointer to **map[string]interface{}** |  | [optional] 
**SourceLegs** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewFxSourceLeg

`func NewFxSourceLeg() *FxSourceLeg`

NewFxSourceLeg instantiates a new FxSourceLeg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxSourceLegWithDefaults

`func NewFxSourceLegWithDefaults() *FxSourceLeg`

NewFxSourceLegWithDefaults instantiates a new FxSourceLeg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPair

`func (o *FxSourceLeg) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *FxSourceLeg) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *FxSourceLeg) SetPair(v string)`

SetPair sets Pair field to given value.

### HasPair

`func (o *FxSourceLeg) HasPair() bool`

HasPair returns a boolean if a field has been set.

### SetPairNil

`func (o *FxSourceLeg) SetPairNil(b bool)`

 SetPairNil sets the value for Pair to be an explicit nil

### UnsetPair
`func (o *FxSourceLeg) UnsetPair()`

UnsetPair ensures that no value is present for Pair, not even an explicit nil
### GetSourcePair

`func (o *FxSourceLeg) GetSourcePair() string`

GetSourcePair returns the SourcePair field if non-nil, zero value otherwise.

### GetSourcePairOk

`func (o *FxSourceLeg) GetSourcePairOk() (*string, bool)`

GetSourcePairOk returns a tuple with the SourcePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePair

`func (o *FxSourceLeg) SetSourcePair(v string)`

SetSourcePair sets SourcePair field to given value.

### HasSourcePair

`func (o *FxSourceLeg) HasSourcePair() bool`

HasSourcePair returns a boolean if a field has been set.

### SetSourcePairNil

`func (o *FxSourceLeg) SetSourcePairNil(b bool)`

 SetSourcePairNil sets the value for SourcePair to be an explicit nil

### UnsetSourcePair
`func (o *FxSourceLeg) UnsetSourcePair()`

UnsetSourcePair ensures that no value is present for SourcePair, not even an explicit nil
### GetStoragePath

`func (o *FxSourceLeg) GetStoragePath() string`

GetStoragePath returns the StoragePath field if non-nil, zero value otherwise.

### GetStoragePathOk

`func (o *FxSourceLeg) GetStoragePathOk() (*string, bool)`

GetStoragePathOk returns a tuple with the StoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePath

`func (o *FxSourceLeg) SetStoragePath(v string)`

SetStoragePath sets StoragePath field to given value.

### HasStoragePath

`func (o *FxSourceLeg) HasStoragePath() bool`

HasStoragePath returns a boolean if a field has been set.

### SetStoragePathNil

`func (o *FxSourceLeg) SetStoragePathNil(b bool)`

 SetStoragePathNil sets the value for StoragePath to be an explicit nil

### UnsetStoragePath
`func (o *FxSourceLeg) UnsetStoragePath()`

UnsetStoragePath ensures that no value is present for StoragePath, not even an explicit nil
### GetStoredSource

`func (o *FxSourceLeg) GetStoredSource() string`

GetStoredSource returns the StoredSource field if non-nil, zero value otherwise.

### GetStoredSourceOk

`func (o *FxSourceLeg) GetStoredSourceOk() (*string, bool)`

GetStoredSourceOk returns a tuple with the StoredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredSource

`func (o *FxSourceLeg) SetStoredSource(v string)`

SetStoredSource sets StoredSource field to given value.

### HasStoredSource

`func (o *FxSourceLeg) HasStoredSource() bool`

HasStoredSource returns a boolean if a field has been set.

### SetStoredSourceNil

`func (o *FxSourceLeg) SetStoredSourceNil(b bool)`

 SetStoredSourceNil sets the value for StoredSource to be an explicit nil

### UnsetStoredSource
`func (o *FxSourceLeg) UnsetStoredSource()`

UnsetStoredSource ensures that no value is present for StoredSource, not even an explicit nil
### GetSourceProvider

`func (o *FxSourceLeg) GetSourceProvider() string`

GetSourceProvider returns the SourceProvider field if non-nil, zero value otherwise.

### GetSourceProviderOk

`func (o *FxSourceLeg) GetSourceProviderOk() (*string, bool)`

GetSourceProviderOk returns a tuple with the SourceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProvider

`func (o *FxSourceLeg) SetSourceProvider(v string)`

SetSourceProvider sets SourceProvider field to given value.

### HasSourceProvider

`func (o *FxSourceLeg) HasSourceProvider() bool`

HasSourceProvider returns a boolean if a field has been set.

### SetSourceProviderNil

`func (o *FxSourceLeg) SetSourceProviderNil(b bool)`

 SetSourceProviderNil sets the value for SourceProvider to be an explicit nil

### UnsetSourceProvider
`func (o *FxSourceLeg) UnsetSourceProvider()`

UnsetSourceProvider ensures that no value is present for SourceProvider, not even an explicit nil
### GetSourceName

`func (o *FxSourceLeg) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *FxSourceLeg) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *FxSourceLeg) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *FxSourceLeg) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *FxSourceLeg) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *FxSourceLeg) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetSourceUrl

`func (o *FxSourceLeg) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *FxSourceLeg) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *FxSourceLeg) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *FxSourceLeg) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *FxSourceLeg) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *FxSourceLeg) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSourceFrequency

`func (o *FxSourceLeg) GetSourceFrequency() string`

GetSourceFrequency returns the SourceFrequency field if non-nil, zero value otherwise.

### GetSourceFrequencyOk

`func (o *FxSourceLeg) GetSourceFrequencyOk() (*string, bool)`

GetSourceFrequencyOk returns a tuple with the SourceFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFrequency

`func (o *FxSourceLeg) SetSourceFrequency(v string)`

SetSourceFrequency sets SourceFrequency field to given value.

### HasSourceFrequency

`func (o *FxSourceLeg) HasSourceFrequency() bool`

HasSourceFrequency returns a boolean if a field has been set.

### SetSourceFrequencyNil

`func (o *FxSourceLeg) SetSourceFrequencyNil(b bool)`

 SetSourceFrequencyNil sets the value for SourceFrequency to be an explicit nil

### UnsetSourceFrequency
`func (o *FxSourceLeg) UnsetSourceFrequency()`

UnsetSourceFrequency ensures that no value is present for SourceFrequency, not even an explicit nil
### GetSourceType

`func (o *FxSourceLeg) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *FxSourceLeg) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *FxSourceLeg) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *FxSourceLeg) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *FxSourceLeg) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *FxSourceLeg) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetIsDerived

`func (o *FxSourceLeg) GetIsDerived() bool`

GetIsDerived returns the IsDerived field if non-nil, zero value otherwise.

### GetIsDerivedOk

`func (o *FxSourceLeg) GetIsDerivedOk() (*bool, bool)`

GetIsDerivedOk returns a tuple with the IsDerived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDerived

`func (o *FxSourceLeg) SetIsDerived(v bool)`

SetIsDerived sets IsDerived field to given value.

### HasIsDerived

`func (o *FxSourceLeg) HasIsDerived() bool`

HasIsDerived returns a boolean if a field has been set.

### SetIsDerivedNil

`func (o *FxSourceLeg) SetIsDerivedNil(b bool)`

 SetIsDerivedNil sets the value for IsDerived to be an explicit nil

### UnsetIsDerived
`func (o *FxSourceLeg) UnsetIsDerived()`

UnsetIsDerived ensures that no value is present for IsDerived, not even an explicit nil
### GetDerivationMethod

`func (o *FxSourceLeg) GetDerivationMethod() string`

GetDerivationMethod returns the DerivationMethod field if non-nil, zero value otherwise.

### GetDerivationMethodOk

`func (o *FxSourceLeg) GetDerivationMethodOk() (*string, bool)`

GetDerivationMethodOk returns a tuple with the DerivationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivationMethod

`func (o *FxSourceLeg) SetDerivationMethod(v string)`

SetDerivationMethod sets DerivationMethod field to given value.

### HasDerivationMethod

`func (o *FxSourceLeg) HasDerivationMethod() bool`

HasDerivationMethod returns a boolean if a field has been set.

### SetDerivationMethodNil

`func (o *FxSourceLeg) SetDerivationMethodNil(b bool)`

 SetDerivationMethodNil sets the value for DerivationMethod to be an explicit nil

### UnsetDerivationMethod
`func (o *FxSourceLeg) UnsetDerivationMethod()`

UnsetDerivationMethod ensures that no value is present for DerivationMethod, not even an explicit nil
### GetAnchorCurrency

`func (o *FxSourceLeg) GetAnchorCurrency() string`

GetAnchorCurrency returns the AnchorCurrency field if non-nil, zero value otherwise.

### GetAnchorCurrencyOk

`func (o *FxSourceLeg) GetAnchorCurrencyOk() (*string, bool)`

GetAnchorCurrencyOk returns a tuple with the AnchorCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCurrency

`func (o *FxSourceLeg) SetAnchorCurrency(v string)`

SetAnchorCurrency sets AnchorCurrency field to given value.

### HasAnchorCurrency

`func (o *FxSourceLeg) HasAnchorCurrency() bool`

HasAnchorCurrency returns a boolean if a field has been set.

### SetAnchorCurrencyNil

`func (o *FxSourceLeg) SetAnchorCurrencyNil(b bool)`

 SetAnchorCurrencyNil sets the value for AnchorCurrency to be an explicit nil

### UnsetAnchorCurrency
`func (o *FxSourceLeg) UnsetAnchorCurrency()`

UnsetAnchorCurrency ensures that no value is present for AnchorCurrency, not even an explicit nil
### GetDataPolicy

`func (o *FxSourceLeg) GetDataPolicy() string`

GetDataPolicy returns the DataPolicy field if non-nil, zero value otherwise.

### GetDataPolicyOk

`func (o *FxSourceLeg) GetDataPolicyOk() (*string, bool)`

GetDataPolicyOk returns a tuple with the DataPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPolicy

`func (o *FxSourceLeg) SetDataPolicy(v string)`

SetDataPolicy sets DataPolicy field to given value.

### HasDataPolicy

`func (o *FxSourceLeg) HasDataPolicy() bool`

HasDataPolicy returns a boolean if a field has been set.

### SetDataPolicyNil

`func (o *FxSourceLeg) SetDataPolicyNil(b bool)`

 SetDataPolicyNil sets the value for DataPolicy to be an explicit nil

### UnsetDataPolicy
`func (o *FxSourceLeg) UnsetDataPolicy()`

UnsetDataPolicy ensures that no value is present for DataPolicy, not even an explicit nil
### GetLastUpdated

`func (o *FxSourceLeg) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *FxSourceLeg) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *FxSourceLeg) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *FxSourceLeg) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *FxSourceLeg) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *FxSourceLeg) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetBackfillWindow

`func (o *FxSourceLeg) GetBackfillWindow() map[string]interface{}`

GetBackfillWindow returns the BackfillWindow field if non-nil, zero value otherwise.

### GetBackfillWindowOk

`func (o *FxSourceLeg) GetBackfillWindowOk() (*map[string]interface{}, bool)`

GetBackfillWindowOk returns a tuple with the BackfillWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfillWindow

`func (o *FxSourceLeg) SetBackfillWindow(v map[string]interface{})`

SetBackfillWindow sets BackfillWindow field to given value.

### HasBackfillWindow

`func (o *FxSourceLeg) HasBackfillWindow() bool`

HasBackfillWindow returns a boolean if a field has been set.

### SetBackfillWindowNil

`func (o *FxSourceLeg) SetBackfillWindowNil(b bool)`

 SetBackfillWindowNil sets the value for BackfillWindow to be an explicit nil

### UnsetBackfillWindow
`func (o *FxSourceLeg) UnsetBackfillWindow()`

UnsetBackfillWindow ensures that no value is present for BackfillWindow, not even an explicit nil
### GetSourceLegs

`func (o *FxSourceLeg) GetSourceLegs() []map[string]interface{}`

GetSourceLegs returns the SourceLegs field if non-nil, zero value otherwise.

### GetSourceLegsOk

`func (o *FxSourceLeg) GetSourceLegsOk() (*[]map[string]interface{}, bool)`

GetSourceLegsOk returns a tuple with the SourceLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLegs

`func (o *FxSourceLeg) SetSourceLegs(v []map[string]interface{})`

SetSourceLegs sets SourceLegs field to given value.

### HasSourceLegs

`func (o *FxSourceLeg) HasSourceLegs() bool`

HasSourceLegs returns a boolean if a field has been set.

### SetSourceLegsNil

`func (o *FxSourceLeg) SetSourceLegsNil(b bool)`

 SetSourceLegsNil sets the value for SourceLegs to be an explicit nil

### UnsetSourceLegs
`func (o *FxSourceLeg) UnsetSourceLegs()`

UnsetSourceLegs ensures that no value is present for SourceLegs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


