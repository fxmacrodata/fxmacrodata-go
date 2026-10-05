# FinancialCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** |  | 
**Available** | **bool** | Whether observations are available for this source; a listed source without observations is not yet available. | 
**FirstObservation** | Pointer to **NullableString** |  | [optional] 
**LastObservation** | Pointer to **NullableString** |  | [optional] 
**LastUpdated** | Pointer to **NullableString** |  | [optional] 
**SourceAccessedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFinancialCoverage

`func NewFinancialCoverage(sourceId string, available bool, ) *FinancialCoverage`

NewFinancialCoverage instantiates a new FinancialCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialCoverageWithDefaults

`func NewFinancialCoverageWithDefaults() *FinancialCoverage`

NewFinancialCoverageWithDefaults instantiates a new FinancialCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *FinancialCoverage) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *FinancialCoverage) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *FinancialCoverage) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetAvailable

`func (o *FinancialCoverage) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *FinancialCoverage) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *FinancialCoverage) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetFirstObservation

`func (o *FinancialCoverage) GetFirstObservation() string`

GetFirstObservation returns the FirstObservation field if non-nil, zero value otherwise.

### GetFirstObservationOk

`func (o *FinancialCoverage) GetFirstObservationOk() (*string, bool)`

GetFirstObservationOk returns a tuple with the FirstObservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstObservation

`func (o *FinancialCoverage) SetFirstObservation(v string)`

SetFirstObservation sets FirstObservation field to given value.

### HasFirstObservation

`func (o *FinancialCoverage) HasFirstObservation() bool`

HasFirstObservation returns a boolean if a field has been set.

### SetFirstObservationNil

`func (o *FinancialCoverage) SetFirstObservationNil(b bool)`

 SetFirstObservationNil sets the value for FirstObservation to be an explicit nil

### UnsetFirstObservation
`func (o *FinancialCoverage) UnsetFirstObservation()`

UnsetFirstObservation ensures that no value is present for FirstObservation, not even an explicit nil
### GetLastObservation

`func (o *FinancialCoverage) GetLastObservation() string`

GetLastObservation returns the LastObservation field if non-nil, zero value otherwise.

### GetLastObservationOk

`func (o *FinancialCoverage) GetLastObservationOk() (*string, bool)`

GetLastObservationOk returns a tuple with the LastObservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastObservation

`func (o *FinancialCoverage) SetLastObservation(v string)`

SetLastObservation sets LastObservation field to given value.

### HasLastObservation

`func (o *FinancialCoverage) HasLastObservation() bool`

HasLastObservation returns a boolean if a field has been set.

### SetLastObservationNil

`func (o *FinancialCoverage) SetLastObservationNil(b bool)`

 SetLastObservationNil sets the value for LastObservation to be an explicit nil

### UnsetLastObservation
`func (o *FinancialCoverage) UnsetLastObservation()`

UnsetLastObservation ensures that no value is present for LastObservation, not even an explicit nil
### GetLastUpdated

`func (o *FinancialCoverage) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *FinancialCoverage) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *FinancialCoverage) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *FinancialCoverage) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *FinancialCoverage) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *FinancialCoverage) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetSourceAccessedAt

`func (o *FinancialCoverage) GetSourceAccessedAt() string`

GetSourceAccessedAt returns the SourceAccessedAt field if non-nil, zero value otherwise.

### GetSourceAccessedAtOk

`func (o *FinancialCoverage) GetSourceAccessedAtOk() (*string, bool)`

GetSourceAccessedAtOk returns a tuple with the SourceAccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccessedAt

`func (o *FinancialCoverage) SetSourceAccessedAt(v string)`

SetSourceAccessedAt sets SourceAccessedAt field to given value.

### HasSourceAccessedAt

`func (o *FinancialCoverage) HasSourceAccessedAt() bool`

HasSourceAccessedAt returns a boolean if a field has been set.

### SetSourceAccessedAtNil

`func (o *FinancialCoverage) SetSourceAccessedAtNil(b bool)`

 SetSourceAccessedAtNil sets the value for SourceAccessedAt to be an explicit nil

### UnsetSourceAccessedAt
`func (o *FinancialCoverage) UnsetSourceAccessedAt()`

UnsetSourceAccessedAt ensures that no value is present for SourceAccessedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


