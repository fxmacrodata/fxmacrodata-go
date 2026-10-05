# RiskSentimentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **NullableString** |  | [optional] 
**EndDate** | Pointer to **NullableString** |  | [optional] 
**LatestAvailableDate** | Pointer to **NullableString** |  | [optional] 
**LastUpdated** | Pointer to **NullableString** |  | [optional] 
**DataQuality** | Pointer to [**NullableDataQuality**](DataQuality.md) |  | [optional] 
**ComponentMetadata** | Pointer to [**NullableRiskSentimentComponentMetadata**](RiskSentimentComponentMetadata.md) |  | [optional] 
**Pagination** | Pointer to [**NullableIndicatorPaginationInfo**](IndicatorPaginationInfo.md) |  | [optional] 
**Data** | Pointer to [**[]RiskSentimentPoint**](RiskSentimentPoint.md) |  | [optional] 

## Methods

### NewRiskSentimentResponse

`func NewRiskSentimentResponse() *RiskSentimentResponse`

NewRiskSentimentResponse instantiates a new RiskSentimentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskSentimentResponseWithDefaults

`func NewRiskSentimentResponseWithDefaults() *RiskSentimentResponse`

NewRiskSentimentResponseWithDefaults instantiates a new RiskSentimentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *RiskSentimentResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *RiskSentimentResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *RiskSentimentResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *RiskSentimentResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *RiskSentimentResponse) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *RiskSentimentResponse) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *RiskSentimentResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *RiskSentimentResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *RiskSentimentResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *RiskSentimentResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *RiskSentimentResponse) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *RiskSentimentResponse) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetLatestAvailableDate

`func (o *RiskSentimentResponse) GetLatestAvailableDate() string`

GetLatestAvailableDate returns the LatestAvailableDate field if non-nil, zero value otherwise.

### GetLatestAvailableDateOk

`func (o *RiskSentimentResponse) GetLatestAvailableDateOk() (*string, bool)`

GetLatestAvailableDateOk returns a tuple with the LatestAvailableDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAvailableDate

`func (o *RiskSentimentResponse) SetLatestAvailableDate(v string)`

SetLatestAvailableDate sets LatestAvailableDate field to given value.

### HasLatestAvailableDate

`func (o *RiskSentimentResponse) HasLatestAvailableDate() bool`

HasLatestAvailableDate returns a boolean if a field has been set.

### SetLatestAvailableDateNil

`func (o *RiskSentimentResponse) SetLatestAvailableDateNil(b bool)`

 SetLatestAvailableDateNil sets the value for LatestAvailableDate to be an explicit nil

### UnsetLatestAvailableDate
`func (o *RiskSentimentResponse) UnsetLatestAvailableDate()`

UnsetLatestAvailableDate ensures that no value is present for LatestAvailableDate, not even an explicit nil
### GetLastUpdated

`func (o *RiskSentimentResponse) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RiskSentimentResponse) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RiskSentimentResponse) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *RiskSentimentResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *RiskSentimentResponse) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *RiskSentimentResponse) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetDataQuality

`func (o *RiskSentimentResponse) GetDataQuality() DataQuality`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *RiskSentimentResponse) GetDataQualityOk() (*DataQuality, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *RiskSentimentResponse) SetDataQuality(v DataQuality)`

SetDataQuality sets DataQuality field to given value.

### HasDataQuality

`func (o *RiskSentimentResponse) HasDataQuality() bool`

HasDataQuality returns a boolean if a field has been set.

### SetDataQualityNil

`func (o *RiskSentimentResponse) SetDataQualityNil(b bool)`

 SetDataQualityNil sets the value for DataQuality to be an explicit nil

### UnsetDataQuality
`func (o *RiskSentimentResponse) UnsetDataQuality()`

UnsetDataQuality ensures that no value is present for DataQuality, not even an explicit nil
### GetComponentMetadata

`func (o *RiskSentimentResponse) GetComponentMetadata() RiskSentimentComponentMetadata`

GetComponentMetadata returns the ComponentMetadata field if non-nil, zero value otherwise.

### GetComponentMetadataOk

`func (o *RiskSentimentResponse) GetComponentMetadataOk() (*RiskSentimentComponentMetadata, bool)`

GetComponentMetadataOk returns a tuple with the ComponentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentMetadata

`func (o *RiskSentimentResponse) SetComponentMetadata(v RiskSentimentComponentMetadata)`

SetComponentMetadata sets ComponentMetadata field to given value.

### HasComponentMetadata

`func (o *RiskSentimentResponse) HasComponentMetadata() bool`

HasComponentMetadata returns a boolean if a field has been set.

### SetComponentMetadataNil

`func (o *RiskSentimentResponse) SetComponentMetadataNil(b bool)`

 SetComponentMetadataNil sets the value for ComponentMetadata to be an explicit nil

### UnsetComponentMetadata
`func (o *RiskSentimentResponse) UnsetComponentMetadata()`

UnsetComponentMetadata ensures that no value is present for ComponentMetadata, not even an explicit nil
### GetPagination

`func (o *RiskSentimentResponse) GetPagination() IndicatorPaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RiskSentimentResponse) GetPaginationOk() (*IndicatorPaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RiskSentimentResponse) SetPagination(v IndicatorPaginationInfo)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RiskSentimentResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### SetPaginationNil

`func (o *RiskSentimentResponse) SetPaginationNil(b bool)`

 SetPaginationNil sets the value for Pagination to be an explicit nil

### UnsetPagination
`func (o *RiskSentimentResponse) UnsetPagination()`

UnsetPagination ensures that no value is present for Pagination, not even an explicit nil
### GetData

`func (o *RiskSentimentResponse) GetData() []RiskSentimentPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RiskSentimentResponse) GetDataOk() (*[]RiskSentimentPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RiskSentimentResponse) SetData(v []RiskSentimentPoint)`

SetData sets Data field to given value.

### HasData

`func (o *RiskSentimentResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *RiskSentimentResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *RiskSentimentResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


