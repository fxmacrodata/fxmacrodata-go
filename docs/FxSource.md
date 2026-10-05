# FxSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**AuthorityType** | Pointer to **NullableString** |  | [optional] 
**CountryOrArea** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**AttributionText** | Pointer to **NullableString** |  | [optional] 
**Tier** | Pointer to **NullableString** |  | [optional] 
**IsOfficial** | Pointer to **NullableBool** |  | [optional] 
**IsAggregator** | Pointer to **NullableBool** |  | [optional] 
**DefaultTimezone** | Pointer to **NullableString** |  | [optional] 
**PublicationFrequency** | Pointer to **NullableString** |  | [optional] 
**NativePairs** | Pointer to **[]string** | Every pair the authority publishes, as documentation of the source. Not a list of requestable pairs; see served_pairs. | [optional] 
**NativePairCount** | Pointer to **NullableInt64** |  | [optional] 
**ServedPairs** | Pointer to **[]string** | The pairs from this source that the public FX endpoints serve. Only these pairs can be requested; a pair listed in native_pairs but not here returns no rows. | [optional] 
**ServedPairCount** | Pointer to **NullableInt64** |  | [optional] 
**ServesPublicData** | Pointer to **NullableBool** | True when at least one pair from this source is served by the public FX endpoints (served_pair_count &gt; 0). False means the source is listed for provenance only and no data from it can be requested. | [optional] 
**CoverageNote** | Pointer to **NullableString** |  | [optional] 
**SourceUniverseEndpoint** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFxSource

`func NewFxSource() *FxSource`

NewFxSource instantiates a new FxSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxSourceWithDefaults

`func NewFxSourceWithDefaults() *FxSource`

NewFxSourceWithDefaults instantiates a new FxSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FxSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FxSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FxSource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FxSource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *FxSource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FxSource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *FxSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FxSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FxSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FxSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FxSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FxSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAuthorityType

`func (o *FxSource) GetAuthorityType() string`

GetAuthorityType returns the AuthorityType field if non-nil, zero value otherwise.

### GetAuthorityTypeOk

`func (o *FxSource) GetAuthorityTypeOk() (*string, bool)`

GetAuthorityTypeOk returns a tuple with the AuthorityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityType

`func (o *FxSource) SetAuthorityType(v string)`

SetAuthorityType sets AuthorityType field to given value.

### HasAuthorityType

`func (o *FxSource) HasAuthorityType() bool`

HasAuthorityType returns a boolean if a field has been set.

### SetAuthorityTypeNil

`func (o *FxSource) SetAuthorityTypeNil(b bool)`

 SetAuthorityTypeNil sets the value for AuthorityType to be an explicit nil

### UnsetAuthorityType
`func (o *FxSource) UnsetAuthorityType()`

UnsetAuthorityType ensures that no value is present for AuthorityType, not even an explicit nil
### GetCountryOrArea

`func (o *FxSource) GetCountryOrArea() string`

GetCountryOrArea returns the CountryOrArea field if non-nil, zero value otherwise.

### GetCountryOrAreaOk

`func (o *FxSource) GetCountryOrAreaOk() (*string, bool)`

GetCountryOrAreaOk returns a tuple with the CountryOrArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOrArea

`func (o *FxSource) SetCountryOrArea(v string)`

SetCountryOrArea sets CountryOrArea field to given value.

### HasCountryOrArea

`func (o *FxSource) HasCountryOrArea() bool`

HasCountryOrArea returns a boolean if a field has been set.

### SetCountryOrAreaNil

`func (o *FxSource) SetCountryOrAreaNil(b bool)`

 SetCountryOrAreaNil sets the value for CountryOrArea to be an explicit nil

### UnsetCountryOrArea
`func (o *FxSource) UnsetCountryOrArea()`

UnsetCountryOrArea ensures that no value is present for CountryOrArea, not even an explicit nil
### GetUrl

`func (o *FxSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FxSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FxSource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FxSource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FxSource) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FxSource) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetAttributionText

`func (o *FxSource) GetAttributionText() string`

GetAttributionText returns the AttributionText field if non-nil, zero value otherwise.

### GetAttributionTextOk

`func (o *FxSource) GetAttributionTextOk() (*string, bool)`

GetAttributionTextOk returns a tuple with the AttributionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionText

`func (o *FxSource) SetAttributionText(v string)`

SetAttributionText sets AttributionText field to given value.

### HasAttributionText

`func (o *FxSource) HasAttributionText() bool`

HasAttributionText returns a boolean if a field has been set.

### SetAttributionTextNil

`func (o *FxSource) SetAttributionTextNil(b bool)`

 SetAttributionTextNil sets the value for AttributionText to be an explicit nil

### UnsetAttributionText
`func (o *FxSource) UnsetAttributionText()`

UnsetAttributionText ensures that no value is present for AttributionText, not even an explicit nil
### GetTier

`func (o *FxSource) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *FxSource) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *FxSource) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *FxSource) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *FxSource) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *FxSource) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil
### GetIsOfficial

`func (o *FxSource) GetIsOfficial() bool`

GetIsOfficial returns the IsOfficial field if non-nil, zero value otherwise.

### GetIsOfficialOk

`func (o *FxSource) GetIsOfficialOk() (*bool, bool)`

GetIsOfficialOk returns a tuple with the IsOfficial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOfficial

`func (o *FxSource) SetIsOfficial(v bool)`

SetIsOfficial sets IsOfficial field to given value.

### HasIsOfficial

`func (o *FxSource) HasIsOfficial() bool`

HasIsOfficial returns a boolean if a field has been set.

### SetIsOfficialNil

`func (o *FxSource) SetIsOfficialNil(b bool)`

 SetIsOfficialNil sets the value for IsOfficial to be an explicit nil

### UnsetIsOfficial
`func (o *FxSource) UnsetIsOfficial()`

UnsetIsOfficial ensures that no value is present for IsOfficial, not even an explicit nil
### GetIsAggregator

`func (o *FxSource) GetIsAggregator() bool`

GetIsAggregator returns the IsAggregator field if non-nil, zero value otherwise.

### GetIsAggregatorOk

`func (o *FxSource) GetIsAggregatorOk() (*bool, bool)`

GetIsAggregatorOk returns a tuple with the IsAggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAggregator

`func (o *FxSource) SetIsAggregator(v bool)`

SetIsAggregator sets IsAggregator field to given value.

### HasIsAggregator

`func (o *FxSource) HasIsAggregator() bool`

HasIsAggregator returns a boolean if a field has been set.

### SetIsAggregatorNil

`func (o *FxSource) SetIsAggregatorNil(b bool)`

 SetIsAggregatorNil sets the value for IsAggregator to be an explicit nil

### UnsetIsAggregator
`func (o *FxSource) UnsetIsAggregator()`

UnsetIsAggregator ensures that no value is present for IsAggregator, not even an explicit nil
### GetDefaultTimezone

`func (o *FxSource) GetDefaultTimezone() string`

GetDefaultTimezone returns the DefaultTimezone field if non-nil, zero value otherwise.

### GetDefaultTimezoneOk

`func (o *FxSource) GetDefaultTimezoneOk() (*string, bool)`

GetDefaultTimezoneOk returns a tuple with the DefaultTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimezone

`func (o *FxSource) SetDefaultTimezone(v string)`

SetDefaultTimezone sets DefaultTimezone field to given value.

### HasDefaultTimezone

`func (o *FxSource) HasDefaultTimezone() bool`

HasDefaultTimezone returns a boolean if a field has been set.

### SetDefaultTimezoneNil

`func (o *FxSource) SetDefaultTimezoneNil(b bool)`

 SetDefaultTimezoneNil sets the value for DefaultTimezone to be an explicit nil

### UnsetDefaultTimezone
`func (o *FxSource) UnsetDefaultTimezone()`

UnsetDefaultTimezone ensures that no value is present for DefaultTimezone, not even an explicit nil
### GetPublicationFrequency

`func (o *FxSource) GetPublicationFrequency() string`

GetPublicationFrequency returns the PublicationFrequency field if non-nil, zero value otherwise.

### GetPublicationFrequencyOk

`func (o *FxSource) GetPublicationFrequencyOk() (*string, bool)`

GetPublicationFrequencyOk returns a tuple with the PublicationFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationFrequency

`func (o *FxSource) SetPublicationFrequency(v string)`

SetPublicationFrequency sets PublicationFrequency field to given value.

### HasPublicationFrequency

`func (o *FxSource) HasPublicationFrequency() bool`

HasPublicationFrequency returns a boolean if a field has been set.

### SetPublicationFrequencyNil

`func (o *FxSource) SetPublicationFrequencyNil(b bool)`

 SetPublicationFrequencyNil sets the value for PublicationFrequency to be an explicit nil

### UnsetPublicationFrequency
`func (o *FxSource) UnsetPublicationFrequency()`

UnsetPublicationFrequency ensures that no value is present for PublicationFrequency, not even an explicit nil
### GetNativePairs

`func (o *FxSource) GetNativePairs() []string`

GetNativePairs returns the NativePairs field if non-nil, zero value otherwise.

### GetNativePairsOk

`func (o *FxSource) GetNativePairsOk() (*[]string, bool)`

GetNativePairsOk returns a tuple with the NativePairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativePairs

`func (o *FxSource) SetNativePairs(v []string)`

SetNativePairs sets NativePairs field to given value.

### HasNativePairs

`func (o *FxSource) HasNativePairs() bool`

HasNativePairs returns a boolean if a field has been set.

### SetNativePairsNil

`func (o *FxSource) SetNativePairsNil(b bool)`

 SetNativePairsNil sets the value for NativePairs to be an explicit nil

### UnsetNativePairs
`func (o *FxSource) UnsetNativePairs()`

UnsetNativePairs ensures that no value is present for NativePairs, not even an explicit nil
### GetNativePairCount

`func (o *FxSource) GetNativePairCount() int64`

GetNativePairCount returns the NativePairCount field if non-nil, zero value otherwise.

### GetNativePairCountOk

`func (o *FxSource) GetNativePairCountOk() (*int64, bool)`

GetNativePairCountOk returns a tuple with the NativePairCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativePairCount

`func (o *FxSource) SetNativePairCount(v int64)`

SetNativePairCount sets NativePairCount field to given value.

### HasNativePairCount

`func (o *FxSource) HasNativePairCount() bool`

HasNativePairCount returns a boolean if a field has been set.

### SetNativePairCountNil

`func (o *FxSource) SetNativePairCountNil(b bool)`

 SetNativePairCountNil sets the value for NativePairCount to be an explicit nil

### UnsetNativePairCount
`func (o *FxSource) UnsetNativePairCount()`

UnsetNativePairCount ensures that no value is present for NativePairCount, not even an explicit nil
### GetServedPairs

`func (o *FxSource) GetServedPairs() []string`

GetServedPairs returns the ServedPairs field if non-nil, zero value otherwise.

### GetServedPairsOk

`func (o *FxSource) GetServedPairsOk() (*[]string, bool)`

GetServedPairsOk returns a tuple with the ServedPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedPairs

`func (o *FxSource) SetServedPairs(v []string)`

SetServedPairs sets ServedPairs field to given value.

### HasServedPairs

`func (o *FxSource) HasServedPairs() bool`

HasServedPairs returns a boolean if a field has been set.

### SetServedPairsNil

`func (o *FxSource) SetServedPairsNil(b bool)`

 SetServedPairsNil sets the value for ServedPairs to be an explicit nil

### UnsetServedPairs
`func (o *FxSource) UnsetServedPairs()`

UnsetServedPairs ensures that no value is present for ServedPairs, not even an explicit nil
### GetServedPairCount

`func (o *FxSource) GetServedPairCount() int64`

GetServedPairCount returns the ServedPairCount field if non-nil, zero value otherwise.

### GetServedPairCountOk

`func (o *FxSource) GetServedPairCountOk() (*int64, bool)`

GetServedPairCountOk returns a tuple with the ServedPairCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedPairCount

`func (o *FxSource) SetServedPairCount(v int64)`

SetServedPairCount sets ServedPairCount field to given value.

### HasServedPairCount

`func (o *FxSource) HasServedPairCount() bool`

HasServedPairCount returns a boolean if a field has been set.

### SetServedPairCountNil

`func (o *FxSource) SetServedPairCountNil(b bool)`

 SetServedPairCountNil sets the value for ServedPairCount to be an explicit nil

### UnsetServedPairCount
`func (o *FxSource) UnsetServedPairCount()`

UnsetServedPairCount ensures that no value is present for ServedPairCount, not even an explicit nil
### GetServesPublicData

`func (o *FxSource) GetServesPublicData() bool`

GetServesPublicData returns the ServesPublicData field if non-nil, zero value otherwise.

### GetServesPublicDataOk

`func (o *FxSource) GetServesPublicDataOk() (*bool, bool)`

GetServesPublicDataOk returns a tuple with the ServesPublicData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServesPublicData

`func (o *FxSource) SetServesPublicData(v bool)`

SetServesPublicData sets ServesPublicData field to given value.

### HasServesPublicData

`func (o *FxSource) HasServesPublicData() bool`

HasServesPublicData returns a boolean if a field has been set.

### SetServesPublicDataNil

`func (o *FxSource) SetServesPublicDataNil(b bool)`

 SetServesPublicDataNil sets the value for ServesPublicData to be an explicit nil

### UnsetServesPublicData
`func (o *FxSource) UnsetServesPublicData()`

UnsetServesPublicData ensures that no value is present for ServesPublicData, not even an explicit nil
### GetCoverageNote

`func (o *FxSource) GetCoverageNote() string`

GetCoverageNote returns the CoverageNote field if non-nil, zero value otherwise.

### GetCoverageNoteOk

`func (o *FxSource) GetCoverageNoteOk() (*string, bool)`

GetCoverageNoteOk returns a tuple with the CoverageNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageNote

`func (o *FxSource) SetCoverageNote(v string)`

SetCoverageNote sets CoverageNote field to given value.

### HasCoverageNote

`func (o *FxSource) HasCoverageNote() bool`

HasCoverageNote returns a boolean if a field has been set.

### SetCoverageNoteNil

`func (o *FxSource) SetCoverageNoteNil(b bool)`

 SetCoverageNoteNil sets the value for CoverageNote to be an explicit nil

### UnsetCoverageNote
`func (o *FxSource) UnsetCoverageNote()`

UnsetCoverageNote ensures that no value is present for CoverageNote, not even an explicit nil
### GetSourceUniverseEndpoint

`func (o *FxSource) GetSourceUniverseEndpoint() string`

GetSourceUniverseEndpoint returns the SourceUniverseEndpoint field if non-nil, zero value otherwise.

### GetSourceUniverseEndpointOk

`func (o *FxSource) GetSourceUniverseEndpointOk() (*string, bool)`

GetSourceUniverseEndpointOk returns a tuple with the SourceUniverseEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUniverseEndpoint

`func (o *FxSource) SetSourceUniverseEndpoint(v string)`

SetSourceUniverseEndpoint sets SourceUniverseEndpoint field to given value.

### HasSourceUniverseEndpoint

`func (o *FxSource) HasSourceUniverseEndpoint() bool`

HasSourceUniverseEndpoint returns a boolean if a field has been set.

### SetSourceUniverseEndpointNil

`func (o *FxSource) SetSourceUniverseEndpointNil(b bool)`

 SetSourceUniverseEndpointNil sets the value for SourceUniverseEndpoint to be an explicit nil

### UnsetSourceUniverseEndpoint
`func (o *FxSource) UnsetSourceUniverseEndpoint()`

UnsetSourceUniverseEndpoint ensures that no value is present for SourceUniverseEndpoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


