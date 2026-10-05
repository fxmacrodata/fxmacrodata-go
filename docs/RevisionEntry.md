# RevisionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Epoch** | Pointer to **NullableInt64** |  | [optional] 
**Val** | Pointer to **NullableFloat64** |  | [optional] 
**ObservedAtNs** | Pointer to **NullableInt64** |  | [optional] 
**ObservedAtNsString** | Pointer to **NullableString** |  | [optional] 
**PublicationAtNs** | Pointer to **NullableInt64** | Publisher-sourced publication instant, Unix epoch nanoseconds (UTC). Present only when the source states one; omitted otherwise. | [optional] 
**PublicationAtNsString** | Pointer to **NullableString** | &#x60;publication_at_ns&#x60; as a decimal string; present only alongside it. | [optional] 
**AnnouncementSourceUrl** | Pointer to **NullableString** |  | [optional] 
**CaptureTimeBasis** | Pointer to **NullableString** |  | [optional] 
**PublicationTimeStatus** | Pointer to **NullableString** |  | [optional] 
**PublicationTimePrecision** | Pointer to **NullableString** |  | [optional] 
**VintageStatus** | Pointer to **NullableString** |  | [optional] 
**IsFirstRelease** | Pointer to **NullableBool** |  | [optional] 
**SourceUrl** | Pointer to **NullableString** |  | [optional] 
**SourceArtifactSha256** | Pointer to **NullableString** |  | [optional] 
**Change** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewRevisionEntry

`func NewRevisionEntry() *RevisionEntry`

NewRevisionEntry instantiates a new RevisionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionEntryWithDefaults

`func NewRevisionEntryWithDefaults() *RevisionEntry`

NewRevisionEntryWithDefaults instantiates a new RevisionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpoch

`func (o *RevisionEntry) GetEpoch() int64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *RevisionEntry) GetEpochOk() (*int64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *RevisionEntry) SetEpoch(v int64)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *RevisionEntry) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### SetEpochNil

`func (o *RevisionEntry) SetEpochNil(b bool)`

 SetEpochNil sets the value for Epoch to be an explicit nil

### UnsetEpoch
`func (o *RevisionEntry) UnsetEpoch()`

UnsetEpoch ensures that no value is present for Epoch, not even an explicit nil
### GetVal

`func (o *RevisionEntry) GetVal() float64`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *RevisionEntry) GetValOk() (*float64, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *RevisionEntry) SetVal(v float64)`

SetVal sets Val field to given value.

### HasVal

`func (o *RevisionEntry) HasVal() bool`

HasVal returns a boolean if a field has been set.

### SetValNil

`func (o *RevisionEntry) SetValNil(b bool)`

 SetValNil sets the value for Val to be an explicit nil

### UnsetVal
`func (o *RevisionEntry) UnsetVal()`

UnsetVal ensures that no value is present for Val, not even an explicit nil
### GetObservedAtNs

`func (o *RevisionEntry) GetObservedAtNs() int64`

GetObservedAtNs returns the ObservedAtNs field if non-nil, zero value otherwise.

### GetObservedAtNsOk

`func (o *RevisionEntry) GetObservedAtNsOk() (*int64, bool)`

GetObservedAtNsOk returns a tuple with the ObservedAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAtNs

`func (o *RevisionEntry) SetObservedAtNs(v int64)`

SetObservedAtNs sets ObservedAtNs field to given value.

### HasObservedAtNs

`func (o *RevisionEntry) HasObservedAtNs() bool`

HasObservedAtNs returns a boolean if a field has been set.

### SetObservedAtNsNil

`func (o *RevisionEntry) SetObservedAtNsNil(b bool)`

 SetObservedAtNsNil sets the value for ObservedAtNs to be an explicit nil

### UnsetObservedAtNs
`func (o *RevisionEntry) UnsetObservedAtNs()`

UnsetObservedAtNs ensures that no value is present for ObservedAtNs, not even an explicit nil
### GetObservedAtNsString

`func (o *RevisionEntry) GetObservedAtNsString() string`

GetObservedAtNsString returns the ObservedAtNsString field if non-nil, zero value otherwise.

### GetObservedAtNsStringOk

`func (o *RevisionEntry) GetObservedAtNsStringOk() (*string, bool)`

GetObservedAtNsStringOk returns a tuple with the ObservedAtNsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAtNsString

`func (o *RevisionEntry) SetObservedAtNsString(v string)`

SetObservedAtNsString sets ObservedAtNsString field to given value.

### HasObservedAtNsString

`func (o *RevisionEntry) HasObservedAtNsString() bool`

HasObservedAtNsString returns a boolean if a field has been set.

### SetObservedAtNsStringNil

`func (o *RevisionEntry) SetObservedAtNsStringNil(b bool)`

 SetObservedAtNsStringNil sets the value for ObservedAtNsString to be an explicit nil

### UnsetObservedAtNsString
`func (o *RevisionEntry) UnsetObservedAtNsString()`

UnsetObservedAtNsString ensures that no value is present for ObservedAtNsString, not even an explicit nil
### GetPublicationAtNs

`func (o *RevisionEntry) GetPublicationAtNs() int64`

GetPublicationAtNs returns the PublicationAtNs field if non-nil, zero value otherwise.

### GetPublicationAtNsOk

`func (o *RevisionEntry) GetPublicationAtNsOk() (*int64, bool)`

GetPublicationAtNsOk returns a tuple with the PublicationAtNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationAtNs

`func (o *RevisionEntry) SetPublicationAtNs(v int64)`

SetPublicationAtNs sets PublicationAtNs field to given value.

### HasPublicationAtNs

`func (o *RevisionEntry) HasPublicationAtNs() bool`

HasPublicationAtNs returns a boolean if a field has been set.

### SetPublicationAtNsNil

`func (o *RevisionEntry) SetPublicationAtNsNil(b bool)`

 SetPublicationAtNsNil sets the value for PublicationAtNs to be an explicit nil

### UnsetPublicationAtNs
`func (o *RevisionEntry) UnsetPublicationAtNs()`

UnsetPublicationAtNs ensures that no value is present for PublicationAtNs, not even an explicit nil
### GetPublicationAtNsString

`func (o *RevisionEntry) GetPublicationAtNsString() string`

GetPublicationAtNsString returns the PublicationAtNsString field if non-nil, zero value otherwise.

### GetPublicationAtNsStringOk

`func (o *RevisionEntry) GetPublicationAtNsStringOk() (*string, bool)`

GetPublicationAtNsStringOk returns a tuple with the PublicationAtNsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationAtNsString

`func (o *RevisionEntry) SetPublicationAtNsString(v string)`

SetPublicationAtNsString sets PublicationAtNsString field to given value.

### HasPublicationAtNsString

`func (o *RevisionEntry) HasPublicationAtNsString() bool`

HasPublicationAtNsString returns a boolean if a field has been set.

### SetPublicationAtNsStringNil

`func (o *RevisionEntry) SetPublicationAtNsStringNil(b bool)`

 SetPublicationAtNsStringNil sets the value for PublicationAtNsString to be an explicit nil

### UnsetPublicationAtNsString
`func (o *RevisionEntry) UnsetPublicationAtNsString()`

UnsetPublicationAtNsString ensures that no value is present for PublicationAtNsString, not even an explicit nil
### GetAnnouncementSourceUrl

`func (o *RevisionEntry) GetAnnouncementSourceUrl() string`

GetAnnouncementSourceUrl returns the AnnouncementSourceUrl field if non-nil, zero value otherwise.

### GetAnnouncementSourceUrlOk

`func (o *RevisionEntry) GetAnnouncementSourceUrlOk() (*string, bool)`

GetAnnouncementSourceUrlOk returns a tuple with the AnnouncementSourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementSourceUrl

`func (o *RevisionEntry) SetAnnouncementSourceUrl(v string)`

SetAnnouncementSourceUrl sets AnnouncementSourceUrl field to given value.

### HasAnnouncementSourceUrl

`func (o *RevisionEntry) HasAnnouncementSourceUrl() bool`

HasAnnouncementSourceUrl returns a boolean if a field has been set.

### SetAnnouncementSourceUrlNil

`func (o *RevisionEntry) SetAnnouncementSourceUrlNil(b bool)`

 SetAnnouncementSourceUrlNil sets the value for AnnouncementSourceUrl to be an explicit nil

### UnsetAnnouncementSourceUrl
`func (o *RevisionEntry) UnsetAnnouncementSourceUrl()`

UnsetAnnouncementSourceUrl ensures that no value is present for AnnouncementSourceUrl, not even an explicit nil
### GetCaptureTimeBasis

`func (o *RevisionEntry) GetCaptureTimeBasis() string`

GetCaptureTimeBasis returns the CaptureTimeBasis field if non-nil, zero value otherwise.

### GetCaptureTimeBasisOk

`func (o *RevisionEntry) GetCaptureTimeBasisOk() (*string, bool)`

GetCaptureTimeBasisOk returns a tuple with the CaptureTimeBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureTimeBasis

`func (o *RevisionEntry) SetCaptureTimeBasis(v string)`

SetCaptureTimeBasis sets CaptureTimeBasis field to given value.

### HasCaptureTimeBasis

`func (o *RevisionEntry) HasCaptureTimeBasis() bool`

HasCaptureTimeBasis returns a boolean if a field has been set.

### SetCaptureTimeBasisNil

`func (o *RevisionEntry) SetCaptureTimeBasisNil(b bool)`

 SetCaptureTimeBasisNil sets the value for CaptureTimeBasis to be an explicit nil

### UnsetCaptureTimeBasis
`func (o *RevisionEntry) UnsetCaptureTimeBasis()`

UnsetCaptureTimeBasis ensures that no value is present for CaptureTimeBasis, not even an explicit nil
### GetPublicationTimeStatus

`func (o *RevisionEntry) GetPublicationTimeStatus() string`

GetPublicationTimeStatus returns the PublicationTimeStatus field if non-nil, zero value otherwise.

### GetPublicationTimeStatusOk

`func (o *RevisionEntry) GetPublicationTimeStatusOk() (*string, bool)`

GetPublicationTimeStatusOk returns a tuple with the PublicationTimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationTimeStatus

`func (o *RevisionEntry) SetPublicationTimeStatus(v string)`

SetPublicationTimeStatus sets PublicationTimeStatus field to given value.

### HasPublicationTimeStatus

`func (o *RevisionEntry) HasPublicationTimeStatus() bool`

HasPublicationTimeStatus returns a boolean if a field has been set.

### SetPublicationTimeStatusNil

`func (o *RevisionEntry) SetPublicationTimeStatusNil(b bool)`

 SetPublicationTimeStatusNil sets the value for PublicationTimeStatus to be an explicit nil

### UnsetPublicationTimeStatus
`func (o *RevisionEntry) UnsetPublicationTimeStatus()`

UnsetPublicationTimeStatus ensures that no value is present for PublicationTimeStatus, not even an explicit nil
### GetPublicationTimePrecision

`func (o *RevisionEntry) GetPublicationTimePrecision() string`

GetPublicationTimePrecision returns the PublicationTimePrecision field if non-nil, zero value otherwise.

### GetPublicationTimePrecisionOk

`func (o *RevisionEntry) GetPublicationTimePrecisionOk() (*string, bool)`

GetPublicationTimePrecisionOk returns a tuple with the PublicationTimePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationTimePrecision

`func (o *RevisionEntry) SetPublicationTimePrecision(v string)`

SetPublicationTimePrecision sets PublicationTimePrecision field to given value.

### HasPublicationTimePrecision

`func (o *RevisionEntry) HasPublicationTimePrecision() bool`

HasPublicationTimePrecision returns a boolean if a field has been set.

### SetPublicationTimePrecisionNil

`func (o *RevisionEntry) SetPublicationTimePrecisionNil(b bool)`

 SetPublicationTimePrecisionNil sets the value for PublicationTimePrecision to be an explicit nil

### UnsetPublicationTimePrecision
`func (o *RevisionEntry) UnsetPublicationTimePrecision()`

UnsetPublicationTimePrecision ensures that no value is present for PublicationTimePrecision, not even an explicit nil
### GetVintageStatus

`func (o *RevisionEntry) GetVintageStatus() string`

GetVintageStatus returns the VintageStatus field if non-nil, zero value otherwise.

### GetVintageStatusOk

`func (o *RevisionEntry) GetVintageStatusOk() (*string, bool)`

GetVintageStatusOk returns a tuple with the VintageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVintageStatus

`func (o *RevisionEntry) SetVintageStatus(v string)`

SetVintageStatus sets VintageStatus field to given value.

### HasVintageStatus

`func (o *RevisionEntry) HasVintageStatus() bool`

HasVintageStatus returns a boolean if a field has been set.

### SetVintageStatusNil

`func (o *RevisionEntry) SetVintageStatusNil(b bool)`

 SetVintageStatusNil sets the value for VintageStatus to be an explicit nil

### UnsetVintageStatus
`func (o *RevisionEntry) UnsetVintageStatus()`

UnsetVintageStatus ensures that no value is present for VintageStatus, not even an explicit nil
### GetIsFirstRelease

`func (o *RevisionEntry) GetIsFirstRelease() bool`

GetIsFirstRelease returns the IsFirstRelease field if non-nil, zero value otherwise.

### GetIsFirstReleaseOk

`func (o *RevisionEntry) GetIsFirstReleaseOk() (*bool, bool)`

GetIsFirstReleaseOk returns a tuple with the IsFirstRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstRelease

`func (o *RevisionEntry) SetIsFirstRelease(v bool)`

SetIsFirstRelease sets IsFirstRelease field to given value.

### HasIsFirstRelease

`func (o *RevisionEntry) HasIsFirstRelease() bool`

HasIsFirstRelease returns a boolean if a field has been set.

### SetIsFirstReleaseNil

`func (o *RevisionEntry) SetIsFirstReleaseNil(b bool)`

 SetIsFirstReleaseNil sets the value for IsFirstRelease to be an explicit nil

### UnsetIsFirstRelease
`func (o *RevisionEntry) UnsetIsFirstRelease()`

UnsetIsFirstRelease ensures that no value is present for IsFirstRelease, not even an explicit nil
### GetSourceUrl

`func (o *RevisionEntry) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *RevisionEntry) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *RevisionEntry) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *RevisionEntry) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *RevisionEntry) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *RevisionEntry) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSourceArtifactSha256

`func (o *RevisionEntry) GetSourceArtifactSha256() string`

GetSourceArtifactSha256 returns the SourceArtifactSha256 field if non-nil, zero value otherwise.

### GetSourceArtifactSha256Ok

`func (o *RevisionEntry) GetSourceArtifactSha256Ok() (*string, bool)`

GetSourceArtifactSha256Ok returns a tuple with the SourceArtifactSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceArtifactSha256

`func (o *RevisionEntry) SetSourceArtifactSha256(v string)`

SetSourceArtifactSha256 sets SourceArtifactSha256 field to given value.

### HasSourceArtifactSha256

`func (o *RevisionEntry) HasSourceArtifactSha256() bool`

HasSourceArtifactSha256 returns a boolean if a field has been set.

### SetSourceArtifactSha256Nil

`func (o *RevisionEntry) SetSourceArtifactSha256Nil(b bool)`

 SetSourceArtifactSha256Nil sets the value for SourceArtifactSha256 to be an explicit nil

### UnsetSourceArtifactSha256
`func (o *RevisionEntry) UnsetSourceArtifactSha256()`

UnsetSourceArtifactSha256 ensures that no value is present for SourceArtifactSha256, not even an explicit nil
### GetChange

`func (o *RevisionEntry) GetChange() float64`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *RevisionEntry) GetChangeOk() (*float64, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *RevisionEntry) SetChange(v float64)`

SetChange sets Change field to given value.

### HasChange

`func (o *RevisionEntry) HasChange() bool`

HasChange returns a boolean if a field has been set.

### SetChangeNil

`func (o *RevisionEntry) SetChangeNil(b bool)`

 SetChangeNil sets the value for Change to be an explicit nil

### UnsetChange
`func (o *RevisionEntry) UnsetChange()`

UnsetChange ensures that no value is present for Change, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


