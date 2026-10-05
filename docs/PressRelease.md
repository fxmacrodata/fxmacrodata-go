# PressRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** |  | [optional] 
**Date** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Topics** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**RatePath** | Pointer to [**NullablePressReleaseRatePath**](PressReleaseRatePath.md) |  | [optional] 
**Relevance** | Pointer to **NullableFloat64** |  | [optional] 
**Sentiment** | Pointer to **NullableFloat64** |  | [optional] 
**AiSummary** | Pointer to **NullableString** | Model-generated one-to-two sentence FX/macro takeaway. Not an official statement. | [optional] 
**AiStance** | Pointer to **NullableString** | Model-assessed stance: &#x60;hawkish&#x60;, &#x60;dovish&#x60; or &#x60;neutral&#x60;. | [optional] 
**AiStanceScore** | Pointer to **NullableFloat64** | Model-assessed stance in [-1, 1]; positive is hawkish. | [optional] 
**AiNextMeetingAction** | Pointer to **NullableString** | Model-assessed next policy action: &#x60;hike&#x60;, &#x60;cut&#x60; or &#x60;hold&#x60;. | [optional] 
**AiNextMeetingProbability** | Pointer to **NullableFloat64** | Model-assessed probability in [0, 1] of &#x60;ai_next_meeting_action&#x60; at the next meeting. | [optional] 
**AiRationale** | Pointer to **NullableString** | Model-generated one-sentence justification for the stance. | [optional] 

## Methods

### NewPressRelease

`func NewPressRelease() *PressRelease`

NewPressRelease instantiates a new PressRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPressReleaseWithDefaults

`func NewPressReleaseWithDefaults() *PressRelease`

NewPressReleaseWithDefaults instantiates a new PressRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PressRelease) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PressRelease) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PressRelease) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PressRelease) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PressRelease) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PressRelease) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetDate

`func (o *PressRelease) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PressRelease) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PressRelease) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PressRelease) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *PressRelease) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *PressRelease) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetTitle

`func (o *PressRelease) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PressRelease) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PressRelease) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PressRelease) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PressRelease) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PressRelease) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTopics

`func (o *PressRelease) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *PressRelease) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *PressRelease) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *PressRelease) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### SetTopicsNil

`func (o *PressRelease) SetTopicsNil(b bool)`

 SetTopicsNil sets the value for Topics to be an explicit nil

### UnsetTopics
`func (o *PressRelease) UnsetTopics()`

UnsetTopics ensures that no value is present for Topics, not even an explicit nil
### GetSummary

`func (o *PressRelease) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PressRelease) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PressRelease) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PressRelease) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *PressRelease) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *PressRelease) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetCategory

`func (o *PressRelease) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PressRelease) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PressRelease) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PressRelease) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *PressRelease) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *PressRelease) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetRatePath

`func (o *PressRelease) GetRatePath() PressReleaseRatePath`

GetRatePath returns the RatePath field if non-nil, zero value otherwise.

### GetRatePathOk

`func (o *PressRelease) GetRatePathOk() (*PressReleaseRatePath, bool)`

GetRatePathOk returns a tuple with the RatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePath

`func (o *PressRelease) SetRatePath(v PressReleaseRatePath)`

SetRatePath sets RatePath field to given value.

### HasRatePath

`func (o *PressRelease) HasRatePath() bool`

HasRatePath returns a boolean if a field has been set.

### SetRatePathNil

`func (o *PressRelease) SetRatePathNil(b bool)`

 SetRatePathNil sets the value for RatePath to be an explicit nil

### UnsetRatePath
`func (o *PressRelease) UnsetRatePath()`

UnsetRatePath ensures that no value is present for RatePath, not even an explicit nil
### GetRelevance

`func (o *PressRelease) GetRelevance() float64`

GetRelevance returns the Relevance field if non-nil, zero value otherwise.

### GetRelevanceOk

`func (o *PressRelease) GetRelevanceOk() (*float64, bool)`

GetRelevanceOk returns a tuple with the Relevance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevance

`func (o *PressRelease) SetRelevance(v float64)`

SetRelevance sets Relevance field to given value.

### HasRelevance

`func (o *PressRelease) HasRelevance() bool`

HasRelevance returns a boolean if a field has been set.

### SetRelevanceNil

`func (o *PressRelease) SetRelevanceNil(b bool)`

 SetRelevanceNil sets the value for Relevance to be an explicit nil

### UnsetRelevance
`func (o *PressRelease) UnsetRelevance()`

UnsetRelevance ensures that no value is present for Relevance, not even an explicit nil
### GetSentiment

`func (o *PressRelease) GetSentiment() float64`

GetSentiment returns the Sentiment field if non-nil, zero value otherwise.

### GetSentimentOk

`func (o *PressRelease) GetSentimentOk() (*float64, bool)`

GetSentimentOk returns a tuple with the Sentiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentiment

`func (o *PressRelease) SetSentiment(v float64)`

SetSentiment sets Sentiment field to given value.

### HasSentiment

`func (o *PressRelease) HasSentiment() bool`

HasSentiment returns a boolean if a field has been set.

### SetSentimentNil

`func (o *PressRelease) SetSentimentNil(b bool)`

 SetSentimentNil sets the value for Sentiment to be an explicit nil

### UnsetSentiment
`func (o *PressRelease) UnsetSentiment()`

UnsetSentiment ensures that no value is present for Sentiment, not even an explicit nil
### GetAiSummary

`func (o *PressRelease) GetAiSummary() string`

GetAiSummary returns the AiSummary field if non-nil, zero value otherwise.

### GetAiSummaryOk

`func (o *PressRelease) GetAiSummaryOk() (*string, bool)`

GetAiSummaryOk returns a tuple with the AiSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSummary

`func (o *PressRelease) SetAiSummary(v string)`

SetAiSummary sets AiSummary field to given value.

### HasAiSummary

`func (o *PressRelease) HasAiSummary() bool`

HasAiSummary returns a boolean if a field has been set.

### SetAiSummaryNil

`func (o *PressRelease) SetAiSummaryNil(b bool)`

 SetAiSummaryNil sets the value for AiSummary to be an explicit nil

### UnsetAiSummary
`func (o *PressRelease) UnsetAiSummary()`

UnsetAiSummary ensures that no value is present for AiSummary, not even an explicit nil
### GetAiStance

`func (o *PressRelease) GetAiStance() string`

GetAiStance returns the AiStance field if non-nil, zero value otherwise.

### GetAiStanceOk

`func (o *PressRelease) GetAiStanceOk() (*string, bool)`

GetAiStanceOk returns a tuple with the AiStance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiStance

`func (o *PressRelease) SetAiStance(v string)`

SetAiStance sets AiStance field to given value.

### HasAiStance

`func (o *PressRelease) HasAiStance() bool`

HasAiStance returns a boolean if a field has been set.

### SetAiStanceNil

`func (o *PressRelease) SetAiStanceNil(b bool)`

 SetAiStanceNil sets the value for AiStance to be an explicit nil

### UnsetAiStance
`func (o *PressRelease) UnsetAiStance()`

UnsetAiStance ensures that no value is present for AiStance, not even an explicit nil
### GetAiStanceScore

`func (o *PressRelease) GetAiStanceScore() float64`

GetAiStanceScore returns the AiStanceScore field if non-nil, zero value otherwise.

### GetAiStanceScoreOk

`func (o *PressRelease) GetAiStanceScoreOk() (*float64, bool)`

GetAiStanceScoreOk returns a tuple with the AiStanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiStanceScore

`func (o *PressRelease) SetAiStanceScore(v float64)`

SetAiStanceScore sets AiStanceScore field to given value.

### HasAiStanceScore

`func (o *PressRelease) HasAiStanceScore() bool`

HasAiStanceScore returns a boolean if a field has been set.

### SetAiStanceScoreNil

`func (o *PressRelease) SetAiStanceScoreNil(b bool)`

 SetAiStanceScoreNil sets the value for AiStanceScore to be an explicit nil

### UnsetAiStanceScore
`func (o *PressRelease) UnsetAiStanceScore()`

UnsetAiStanceScore ensures that no value is present for AiStanceScore, not even an explicit nil
### GetAiNextMeetingAction

`func (o *PressRelease) GetAiNextMeetingAction() string`

GetAiNextMeetingAction returns the AiNextMeetingAction field if non-nil, zero value otherwise.

### GetAiNextMeetingActionOk

`func (o *PressRelease) GetAiNextMeetingActionOk() (*string, bool)`

GetAiNextMeetingActionOk returns a tuple with the AiNextMeetingAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiNextMeetingAction

`func (o *PressRelease) SetAiNextMeetingAction(v string)`

SetAiNextMeetingAction sets AiNextMeetingAction field to given value.

### HasAiNextMeetingAction

`func (o *PressRelease) HasAiNextMeetingAction() bool`

HasAiNextMeetingAction returns a boolean if a field has been set.

### SetAiNextMeetingActionNil

`func (o *PressRelease) SetAiNextMeetingActionNil(b bool)`

 SetAiNextMeetingActionNil sets the value for AiNextMeetingAction to be an explicit nil

### UnsetAiNextMeetingAction
`func (o *PressRelease) UnsetAiNextMeetingAction()`

UnsetAiNextMeetingAction ensures that no value is present for AiNextMeetingAction, not even an explicit nil
### GetAiNextMeetingProbability

`func (o *PressRelease) GetAiNextMeetingProbability() float64`

GetAiNextMeetingProbability returns the AiNextMeetingProbability field if non-nil, zero value otherwise.

### GetAiNextMeetingProbabilityOk

`func (o *PressRelease) GetAiNextMeetingProbabilityOk() (*float64, bool)`

GetAiNextMeetingProbabilityOk returns a tuple with the AiNextMeetingProbability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiNextMeetingProbability

`func (o *PressRelease) SetAiNextMeetingProbability(v float64)`

SetAiNextMeetingProbability sets AiNextMeetingProbability field to given value.

### HasAiNextMeetingProbability

`func (o *PressRelease) HasAiNextMeetingProbability() bool`

HasAiNextMeetingProbability returns a boolean if a field has been set.

### SetAiNextMeetingProbabilityNil

`func (o *PressRelease) SetAiNextMeetingProbabilityNil(b bool)`

 SetAiNextMeetingProbabilityNil sets the value for AiNextMeetingProbability to be an explicit nil

### UnsetAiNextMeetingProbability
`func (o *PressRelease) UnsetAiNextMeetingProbability()`

UnsetAiNextMeetingProbability ensures that no value is present for AiNextMeetingProbability, not even an explicit nil
### GetAiRationale

`func (o *PressRelease) GetAiRationale() string`

GetAiRationale returns the AiRationale field if non-nil, zero value otherwise.

### GetAiRationaleOk

`func (o *PressRelease) GetAiRationaleOk() (*string, bool)`

GetAiRationaleOk returns a tuple with the AiRationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiRationale

`func (o *PressRelease) SetAiRationale(v string)`

SetAiRationale sets AiRationale field to given value.

### HasAiRationale

`func (o *PressRelease) HasAiRationale() bool`

HasAiRationale returns a boolean if a field has been set.

### SetAiRationaleNil

`func (o *PressRelease) SetAiRationaleNil(b bool)`

 SetAiRationaleNil sets the value for AiRationale to be an explicit nil

### UnsetAiRationale
`func (o *PressRelease) UnsetAiRationale()`

UnsetAiRationale ensures that no value is present for AiRationale, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


