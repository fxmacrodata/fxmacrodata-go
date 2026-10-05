# ResearchPanelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | [**[]ResearchSeries**](ResearchSeries.md) |  | 
**DecisionTimes** | **[]string** |  | 
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**Availability** | Pointer to **string** |  | [optional] [default to "public"]
**DatasetVersion** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResearchPanelRequest

`func NewResearchPanelRequest(series []ResearchSeries, decisionTimes []string, startDate string, endDate string, ) *ResearchPanelRequest`

NewResearchPanelRequest instantiates a new ResearchPanelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResearchPanelRequestWithDefaults

`func NewResearchPanelRequestWithDefaults() *ResearchPanelRequest`

NewResearchPanelRequestWithDefaults instantiates a new ResearchPanelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeries

`func (o *ResearchPanelRequest) GetSeries() []ResearchSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ResearchPanelRequest) GetSeriesOk() (*[]ResearchSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ResearchPanelRequest) SetSeries(v []ResearchSeries)`

SetSeries sets Series field to given value.


### GetDecisionTimes

`func (o *ResearchPanelRequest) GetDecisionTimes() []string`

GetDecisionTimes returns the DecisionTimes field if non-nil, zero value otherwise.

### GetDecisionTimesOk

`func (o *ResearchPanelRequest) GetDecisionTimesOk() (*[]string, bool)`

GetDecisionTimesOk returns a tuple with the DecisionTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionTimes

`func (o *ResearchPanelRequest) SetDecisionTimes(v []string)`

SetDecisionTimes sets DecisionTimes field to given value.


### GetStartDate

`func (o *ResearchPanelRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ResearchPanelRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ResearchPanelRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *ResearchPanelRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ResearchPanelRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ResearchPanelRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetAvailability

`func (o *ResearchPanelRequest) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ResearchPanelRequest) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ResearchPanelRequest) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *ResearchPanelRequest) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetDatasetVersion

`func (o *ResearchPanelRequest) GetDatasetVersion() string`

GetDatasetVersion returns the DatasetVersion field if non-nil, zero value otherwise.

### GetDatasetVersionOk

`func (o *ResearchPanelRequest) GetDatasetVersionOk() (*string, bool)`

GetDatasetVersionOk returns a tuple with the DatasetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetVersion

`func (o *ResearchPanelRequest) SetDatasetVersion(v string)`

SetDatasetVersion sets DatasetVersion field to given value.

### HasDatasetVersion

`func (o *ResearchPanelRequest) HasDatasetVersion() bool`

HasDatasetVersion returns a boolean if a field has been set.

### SetDatasetVersionNil

`func (o *ResearchPanelRequest) SetDatasetVersionNil(b bool)`

 SetDatasetVersionNil sets the value for DatasetVersion to be an explicit nil

### UnsetDatasetVersion
`func (o *ResearchPanelRequest) UnsetDatasetVersion()`

UnsetDatasetVersion ensures that no value is present for DatasetVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


