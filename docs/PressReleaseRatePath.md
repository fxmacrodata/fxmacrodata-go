# PressReleaseRatePath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **NullableString** |  | [optional] 
**Score** | Pointer to **NullableFloat64** |  | [optional] 
**Matches** | Pointer to [**[]RatePathMatch**](RatePathMatch.md) |  | [optional] 
**RawScore** | Pointer to **NullableFloat64** |  | [optional] 
**Confidence** | Pointer to **NullableString** |  | [optional] 
**BiasAction** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPressReleaseRatePath

`func NewPressReleaseRatePath() *PressReleaseRatePath`

NewPressReleaseRatePath instantiates a new PressReleaseRatePath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPressReleaseRatePathWithDefaults

`func NewPressReleaseRatePathWithDefaults() *PressReleaseRatePath`

NewPressReleaseRatePathWithDefaults instantiates a new PressReleaseRatePath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PressReleaseRatePath) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PressReleaseRatePath) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PressReleaseRatePath) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PressReleaseRatePath) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *PressReleaseRatePath) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *PressReleaseRatePath) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetScore

`func (o *PressReleaseRatePath) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *PressReleaseRatePath) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *PressReleaseRatePath) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *PressReleaseRatePath) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *PressReleaseRatePath) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *PressReleaseRatePath) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetMatches

`func (o *PressReleaseRatePath) GetMatches() []RatePathMatch`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *PressReleaseRatePath) GetMatchesOk() (*[]RatePathMatch, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *PressReleaseRatePath) SetMatches(v []RatePathMatch)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *PressReleaseRatePath) HasMatches() bool`

HasMatches returns a boolean if a field has been set.

### SetMatchesNil

`func (o *PressReleaseRatePath) SetMatchesNil(b bool)`

 SetMatchesNil sets the value for Matches to be an explicit nil

### UnsetMatches
`func (o *PressReleaseRatePath) UnsetMatches()`

UnsetMatches ensures that no value is present for Matches, not even an explicit nil
### GetRawScore

`func (o *PressReleaseRatePath) GetRawScore() float64`

GetRawScore returns the RawScore field if non-nil, zero value otherwise.

### GetRawScoreOk

`func (o *PressReleaseRatePath) GetRawScoreOk() (*float64, bool)`

GetRawScoreOk returns a tuple with the RawScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawScore

`func (o *PressReleaseRatePath) SetRawScore(v float64)`

SetRawScore sets RawScore field to given value.

### HasRawScore

`func (o *PressReleaseRatePath) HasRawScore() bool`

HasRawScore returns a boolean if a field has been set.

### SetRawScoreNil

`func (o *PressReleaseRatePath) SetRawScoreNil(b bool)`

 SetRawScoreNil sets the value for RawScore to be an explicit nil

### UnsetRawScore
`func (o *PressReleaseRatePath) UnsetRawScore()`

UnsetRawScore ensures that no value is present for RawScore, not even an explicit nil
### GetConfidence

`func (o *PressReleaseRatePath) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *PressReleaseRatePath) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *PressReleaseRatePath) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *PressReleaseRatePath) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *PressReleaseRatePath) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *PressReleaseRatePath) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetBiasAction

`func (o *PressReleaseRatePath) GetBiasAction() string`

GetBiasAction returns the BiasAction field if non-nil, zero value otherwise.

### GetBiasActionOk

`func (o *PressReleaseRatePath) GetBiasActionOk() (*string, bool)`

GetBiasActionOk returns a tuple with the BiasAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiasAction

`func (o *PressReleaseRatePath) SetBiasAction(v string)`

SetBiasAction sets BiasAction field to given value.

### HasBiasAction

`func (o *PressReleaseRatePath) HasBiasAction() bool`

HasBiasAction returns a boolean if a field has been set.

### SetBiasActionNil

`func (o *PressReleaseRatePath) SetBiasActionNil(b bool)`

 SetBiasActionNil sets the value for BiasAction to be an explicit nil

### UnsetBiasAction
`func (o *PressReleaseRatePath) UnsetBiasAction()`

UnsetBiasAction ensures that no value is present for BiasAction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


