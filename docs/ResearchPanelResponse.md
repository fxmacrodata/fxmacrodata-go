# ResearchPanelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**DatasetVersion** | **string** |  | 
**Complete** | **bool** |  | 
**Availability** | **string** |  | 
**ValueMode** | **string** |  | 
**ValueMetadata** | **map[string]interface{}** |  | 
**Series** | **[]map[string]interface{}** |  | 
**Data** | **[]map[string]interface{}** |  | 
**Quality** | **map[string]interface{}** |  | 
**Limits** | **map[string]int64** |  | 

## Methods

### NewResearchPanelResponse

`func NewResearchPanelResponse(status string, datasetVersion string, complete bool, availability string, valueMode string, valueMetadata map[string]interface{}, series []map[string]interface{}, data []map[string]interface{}, quality map[string]interface{}, limits map[string]int64, ) *ResearchPanelResponse`

NewResearchPanelResponse instantiates a new ResearchPanelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResearchPanelResponseWithDefaults

`func NewResearchPanelResponseWithDefaults() *ResearchPanelResponse`

NewResearchPanelResponseWithDefaults instantiates a new ResearchPanelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResearchPanelResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResearchPanelResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResearchPanelResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDatasetVersion

`func (o *ResearchPanelResponse) GetDatasetVersion() string`

GetDatasetVersion returns the DatasetVersion field if non-nil, zero value otherwise.

### GetDatasetVersionOk

`func (o *ResearchPanelResponse) GetDatasetVersionOk() (*string, bool)`

GetDatasetVersionOk returns a tuple with the DatasetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetVersion

`func (o *ResearchPanelResponse) SetDatasetVersion(v string)`

SetDatasetVersion sets DatasetVersion field to given value.


### GetComplete

`func (o *ResearchPanelResponse) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *ResearchPanelResponse) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *ResearchPanelResponse) SetComplete(v bool)`

SetComplete sets Complete field to given value.


### GetAvailability

`func (o *ResearchPanelResponse) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ResearchPanelResponse) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ResearchPanelResponse) SetAvailability(v string)`

SetAvailability sets Availability field to given value.


### GetValueMode

`func (o *ResearchPanelResponse) GetValueMode() string`

GetValueMode returns the ValueMode field if non-nil, zero value otherwise.

### GetValueModeOk

`func (o *ResearchPanelResponse) GetValueModeOk() (*string, bool)`

GetValueModeOk returns a tuple with the ValueMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMode

`func (o *ResearchPanelResponse) SetValueMode(v string)`

SetValueMode sets ValueMode field to given value.


### GetValueMetadata

`func (o *ResearchPanelResponse) GetValueMetadata() map[string]interface{}`

GetValueMetadata returns the ValueMetadata field if non-nil, zero value otherwise.

### GetValueMetadataOk

`func (o *ResearchPanelResponse) GetValueMetadataOk() (*map[string]interface{}, bool)`

GetValueMetadataOk returns a tuple with the ValueMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMetadata

`func (o *ResearchPanelResponse) SetValueMetadata(v map[string]interface{})`

SetValueMetadata sets ValueMetadata field to given value.


### GetSeries

`func (o *ResearchPanelResponse) GetSeries() []map[string]interface{}`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ResearchPanelResponse) GetSeriesOk() (*[]map[string]interface{}, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ResearchPanelResponse) SetSeries(v []map[string]interface{})`

SetSeries sets Series field to given value.


### GetData

`func (o *ResearchPanelResponse) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResearchPanelResponse) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResearchPanelResponse) SetData(v []map[string]interface{})`

SetData sets Data field to given value.


### GetQuality

`func (o *ResearchPanelResponse) GetQuality() map[string]interface{}`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ResearchPanelResponse) GetQualityOk() (*map[string]interface{}, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ResearchPanelResponse) SetQuality(v map[string]interface{})`

SetQuality sets Quality field to given value.


### GetLimits

`func (o *ResearchPanelResponse) GetLimits() map[string]int64`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *ResearchPanelResponse) GetLimitsOk() (*map[string]int64, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *ResearchPanelResponse) SetLimits(v map[string]int64)`

SetLimits sets Limits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


