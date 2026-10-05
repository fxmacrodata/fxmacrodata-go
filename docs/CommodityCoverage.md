# CommodityCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** |  | 
**RequestedStart** | Pointer to **NullableString** |  | [optional] 
**RequestedEnd** | Pointer to **NullableString** |  | [optional] 
**AvailableStart** | Pointer to **NullableString** |  | [optional] 
**AvailableEnd** | Pointer to **NullableString** |  | [optional] 
**InputRowCount** | **int64** |  | 
**RowCount** | **int64** |  | 
**DistinctDateCount** | **int64** |  | 
**DuplicateDateRowCount** | **int64** |  | 
**ExcludedWeekendRowCount** | **int64** |  | 
**InvalidDateRowCount** | **int64** |  | 
**ExpectedWeekdayLabelCount** | **int64** |  | 
**MissingWeekdayLabelCount** | **int64** |  | 
**LargestMissingWeekdayRun** | **int64** |  | 
**LargestCalendarGapDays** | **int64** |  | 
**MissingWeekdayRanges** | [**[]CommodityMissingWeekdayRange**](CommodityMissingWeekdayRange.md) |  | 
**LongestContinuousWeekdaySegment** | [**CommodityCoverageSegment**](CommodityCoverageSegment.md) |  | 
**LatestContinuousWeekdaySegment** | [**CommodityCoverageSegment**](CommodityCoverageSegment.md) |  | 
**SparseHistory** | **bool** |  | 
**CoverageBasis** | **string** |  | 
**ContinuousCoverageVerified** | Pointer to **bool** |  | [optional] [default to false]
**ReferenceOnly** | Pointer to **bool** |  | [optional] [default to true]
**ExecutableQuotes** | Pointer to **bool** |  | [optional] [default to false]
**ProvenanceVersion** | **int64** |  | 
**RepairedExcludedWeekendRowCount** | Pointer to **int64** |  | [optional] [default to 0]
**RepairedInvalidDateRowCount** | Pointer to **int64** |  | [optional] [default to 0]

## Methods

### NewCommodityCoverage

`func NewCommodityCoverage(scope string, inputRowCount int64, rowCount int64, distinctDateCount int64, duplicateDateRowCount int64, excludedWeekendRowCount int64, invalidDateRowCount int64, expectedWeekdayLabelCount int64, missingWeekdayLabelCount int64, largestMissingWeekdayRun int64, largestCalendarGapDays int64, missingWeekdayRanges []CommodityMissingWeekdayRange, longestContinuousWeekdaySegment CommodityCoverageSegment, latestContinuousWeekdaySegment CommodityCoverageSegment, sparseHistory bool, coverageBasis string, provenanceVersion int64, ) *CommodityCoverage`

NewCommodityCoverage instantiates a new CommodityCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommodityCoverageWithDefaults

`func NewCommodityCoverageWithDefaults() *CommodityCoverage`

NewCommodityCoverageWithDefaults instantiates a new CommodityCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *CommodityCoverage) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CommodityCoverage) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CommodityCoverage) SetScope(v string)`

SetScope sets Scope field to given value.


### GetRequestedStart

`func (o *CommodityCoverage) GetRequestedStart() string`

GetRequestedStart returns the RequestedStart field if non-nil, zero value otherwise.

### GetRequestedStartOk

`func (o *CommodityCoverage) GetRequestedStartOk() (*string, bool)`

GetRequestedStartOk returns a tuple with the RequestedStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedStart

`func (o *CommodityCoverage) SetRequestedStart(v string)`

SetRequestedStart sets RequestedStart field to given value.

### HasRequestedStart

`func (o *CommodityCoverage) HasRequestedStart() bool`

HasRequestedStart returns a boolean if a field has been set.

### SetRequestedStartNil

`func (o *CommodityCoverage) SetRequestedStartNil(b bool)`

 SetRequestedStartNil sets the value for RequestedStart to be an explicit nil

### UnsetRequestedStart
`func (o *CommodityCoverage) UnsetRequestedStart()`

UnsetRequestedStart ensures that no value is present for RequestedStart, not even an explicit nil
### GetRequestedEnd

`func (o *CommodityCoverage) GetRequestedEnd() string`

GetRequestedEnd returns the RequestedEnd field if non-nil, zero value otherwise.

### GetRequestedEndOk

`func (o *CommodityCoverage) GetRequestedEndOk() (*string, bool)`

GetRequestedEndOk returns a tuple with the RequestedEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedEnd

`func (o *CommodityCoverage) SetRequestedEnd(v string)`

SetRequestedEnd sets RequestedEnd field to given value.

### HasRequestedEnd

`func (o *CommodityCoverage) HasRequestedEnd() bool`

HasRequestedEnd returns a boolean if a field has been set.

### SetRequestedEndNil

`func (o *CommodityCoverage) SetRequestedEndNil(b bool)`

 SetRequestedEndNil sets the value for RequestedEnd to be an explicit nil

### UnsetRequestedEnd
`func (o *CommodityCoverage) UnsetRequestedEnd()`

UnsetRequestedEnd ensures that no value is present for RequestedEnd, not even an explicit nil
### GetAvailableStart

`func (o *CommodityCoverage) GetAvailableStart() string`

GetAvailableStart returns the AvailableStart field if non-nil, zero value otherwise.

### GetAvailableStartOk

`func (o *CommodityCoverage) GetAvailableStartOk() (*string, bool)`

GetAvailableStartOk returns a tuple with the AvailableStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStart

`func (o *CommodityCoverage) SetAvailableStart(v string)`

SetAvailableStart sets AvailableStart field to given value.

### HasAvailableStart

`func (o *CommodityCoverage) HasAvailableStart() bool`

HasAvailableStart returns a boolean if a field has been set.

### SetAvailableStartNil

`func (o *CommodityCoverage) SetAvailableStartNil(b bool)`

 SetAvailableStartNil sets the value for AvailableStart to be an explicit nil

### UnsetAvailableStart
`func (o *CommodityCoverage) UnsetAvailableStart()`

UnsetAvailableStart ensures that no value is present for AvailableStart, not even an explicit nil
### GetAvailableEnd

`func (o *CommodityCoverage) GetAvailableEnd() string`

GetAvailableEnd returns the AvailableEnd field if non-nil, zero value otherwise.

### GetAvailableEndOk

`func (o *CommodityCoverage) GetAvailableEndOk() (*string, bool)`

GetAvailableEndOk returns a tuple with the AvailableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableEnd

`func (o *CommodityCoverage) SetAvailableEnd(v string)`

SetAvailableEnd sets AvailableEnd field to given value.

### HasAvailableEnd

`func (o *CommodityCoverage) HasAvailableEnd() bool`

HasAvailableEnd returns a boolean if a field has been set.

### SetAvailableEndNil

`func (o *CommodityCoverage) SetAvailableEndNil(b bool)`

 SetAvailableEndNil sets the value for AvailableEnd to be an explicit nil

### UnsetAvailableEnd
`func (o *CommodityCoverage) UnsetAvailableEnd()`

UnsetAvailableEnd ensures that no value is present for AvailableEnd, not even an explicit nil
### GetInputRowCount

`func (o *CommodityCoverage) GetInputRowCount() int64`

GetInputRowCount returns the InputRowCount field if non-nil, zero value otherwise.

### GetInputRowCountOk

`func (o *CommodityCoverage) GetInputRowCountOk() (*int64, bool)`

GetInputRowCountOk returns a tuple with the InputRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRowCount

`func (o *CommodityCoverage) SetInputRowCount(v int64)`

SetInputRowCount sets InputRowCount field to given value.


### GetRowCount

`func (o *CommodityCoverage) GetRowCount() int64`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *CommodityCoverage) GetRowCountOk() (*int64, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *CommodityCoverage) SetRowCount(v int64)`

SetRowCount sets RowCount field to given value.


### GetDistinctDateCount

`func (o *CommodityCoverage) GetDistinctDateCount() int64`

GetDistinctDateCount returns the DistinctDateCount field if non-nil, zero value otherwise.

### GetDistinctDateCountOk

`func (o *CommodityCoverage) GetDistinctDateCountOk() (*int64, bool)`

GetDistinctDateCountOk returns a tuple with the DistinctDateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctDateCount

`func (o *CommodityCoverage) SetDistinctDateCount(v int64)`

SetDistinctDateCount sets DistinctDateCount field to given value.


### GetDuplicateDateRowCount

`func (o *CommodityCoverage) GetDuplicateDateRowCount() int64`

GetDuplicateDateRowCount returns the DuplicateDateRowCount field if non-nil, zero value otherwise.

### GetDuplicateDateRowCountOk

`func (o *CommodityCoverage) GetDuplicateDateRowCountOk() (*int64, bool)`

GetDuplicateDateRowCountOk returns a tuple with the DuplicateDateRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateDateRowCount

`func (o *CommodityCoverage) SetDuplicateDateRowCount(v int64)`

SetDuplicateDateRowCount sets DuplicateDateRowCount field to given value.


### GetExcludedWeekendRowCount

`func (o *CommodityCoverage) GetExcludedWeekendRowCount() int64`

GetExcludedWeekendRowCount returns the ExcludedWeekendRowCount field if non-nil, zero value otherwise.

### GetExcludedWeekendRowCountOk

`func (o *CommodityCoverage) GetExcludedWeekendRowCountOk() (*int64, bool)`

GetExcludedWeekendRowCountOk returns a tuple with the ExcludedWeekendRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedWeekendRowCount

`func (o *CommodityCoverage) SetExcludedWeekendRowCount(v int64)`

SetExcludedWeekendRowCount sets ExcludedWeekendRowCount field to given value.


### GetInvalidDateRowCount

`func (o *CommodityCoverage) GetInvalidDateRowCount() int64`

GetInvalidDateRowCount returns the InvalidDateRowCount field if non-nil, zero value otherwise.

### GetInvalidDateRowCountOk

`func (o *CommodityCoverage) GetInvalidDateRowCountOk() (*int64, bool)`

GetInvalidDateRowCountOk returns a tuple with the InvalidDateRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidDateRowCount

`func (o *CommodityCoverage) SetInvalidDateRowCount(v int64)`

SetInvalidDateRowCount sets InvalidDateRowCount field to given value.


### GetExpectedWeekdayLabelCount

`func (o *CommodityCoverage) GetExpectedWeekdayLabelCount() int64`

GetExpectedWeekdayLabelCount returns the ExpectedWeekdayLabelCount field if non-nil, zero value otherwise.

### GetExpectedWeekdayLabelCountOk

`func (o *CommodityCoverage) GetExpectedWeekdayLabelCountOk() (*int64, bool)`

GetExpectedWeekdayLabelCountOk returns a tuple with the ExpectedWeekdayLabelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedWeekdayLabelCount

`func (o *CommodityCoverage) SetExpectedWeekdayLabelCount(v int64)`

SetExpectedWeekdayLabelCount sets ExpectedWeekdayLabelCount field to given value.


### GetMissingWeekdayLabelCount

`func (o *CommodityCoverage) GetMissingWeekdayLabelCount() int64`

GetMissingWeekdayLabelCount returns the MissingWeekdayLabelCount field if non-nil, zero value otherwise.

### GetMissingWeekdayLabelCountOk

`func (o *CommodityCoverage) GetMissingWeekdayLabelCountOk() (*int64, bool)`

GetMissingWeekdayLabelCountOk returns a tuple with the MissingWeekdayLabelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingWeekdayLabelCount

`func (o *CommodityCoverage) SetMissingWeekdayLabelCount(v int64)`

SetMissingWeekdayLabelCount sets MissingWeekdayLabelCount field to given value.


### GetLargestMissingWeekdayRun

`func (o *CommodityCoverage) GetLargestMissingWeekdayRun() int64`

GetLargestMissingWeekdayRun returns the LargestMissingWeekdayRun field if non-nil, zero value otherwise.

### GetLargestMissingWeekdayRunOk

`func (o *CommodityCoverage) GetLargestMissingWeekdayRunOk() (*int64, bool)`

GetLargestMissingWeekdayRunOk returns a tuple with the LargestMissingWeekdayRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargestMissingWeekdayRun

`func (o *CommodityCoverage) SetLargestMissingWeekdayRun(v int64)`

SetLargestMissingWeekdayRun sets LargestMissingWeekdayRun field to given value.


### GetLargestCalendarGapDays

`func (o *CommodityCoverage) GetLargestCalendarGapDays() int64`

GetLargestCalendarGapDays returns the LargestCalendarGapDays field if non-nil, zero value otherwise.

### GetLargestCalendarGapDaysOk

`func (o *CommodityCoverage) GetLargestCalendarGapDaysOk() (*int64, bool)`

GetLargestCalendarGapDaysOk returns a tuple with the LargestCalendarGapDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargestCalendarGapDays

`func (o *CommodityCoverage) SetLargestCalendarGapDays(v int64)`

SetLargestCalendarGapDays sets LargestCalendarGapDays field to given value.


### GetMissingWeekdayRanges

`func (o *CommodityCoverage) GetMissingWeekdayRanges() []CommodityMissingWeekdayRange`

GetMissingWeekdayRanges returns the MissingWeekdayRanges field if non-nil, zero value otherwise.

### GetMissingWeekdayRangesOk

`func (o *CommodityCoverage) GetMissingWeekdayRangesOk() (*[]CommodityMissingWeekdayRange, bool)`

GetMissingWeekdayRangesOk returns a tuple with the MissingWeekdayRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingWeekdayRanges

`func (o *CommodityCoverage) SetMissingWeekdayRanges(v []CommodityMissingWeekdayRange)`

SetMissingWeekdayRanges sets MissingWeekdayRanges field to given value.


### GetLongestContinuousWeekdaySegment

`func (o *CommodityCoverage) GetLongestContinuousWeekdaySegment() CommodityCoverageSegment`

GetLongestContinuousWeekdaySegment returns the LongestContinuousWeekdaySegment field if non-nil, zero value otherwise.

### GetLongestContinuousWeekdaySegmentOk

`func (o *CommodityCoverage) GetLongestContinuousWeekdaySegmentOk() (*CommodityCoverageSegment, bool)`

GetLongestContinuousWeekdaySegmentOk returns a tuple with the LongestContinuousWeekdaySegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestContinuousWeekdaySegment

`func (o *CommodityCoverage) SetLongestContinuousWeekdaySegment(v CommodityCoverageSegment)`

SetLongestContinuousWeekdaySegment sets LongestContinuousWeekdaySegment field to given value.


### GetLatestContinuousWeekdaySegment

`func (o *CommodityCoverage) GetLatestContinuousWeekdaySegment() CommodityCoverageSegment`

GetLatestContinuousWeekdaySegment returns the LatestContinuousWeekdaySegment field if non-nil, zero value otherwise.

### GetLatestContinuousWeekdaySegmentOk

`func (o *CommodityCoverage) GetLatestContinuousWeekdaySegmentOk() (*CommodityCoverageSegment, bool)`

GetLatestContinuousWeekdaySegmentOk returns a tuple with the LatestContinuousWeekdaySegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestContinuousWeekdaySegment

`func (o *CommodityCoverage) SetLatestContinuousWeekdaySegment(v CommodityCoverageSegment)`

SetLatestContinuousWeekdaySegment sets LatestContinuousWeekdaySegment field to given value.


### GetSparseHistory

`func (o *CommodityCoverage) GetSparseHistory() bool`

GetSparseHistory returns the SparseHistory field if non-nil, zero value otherwise.

### GetSparseHistoryOk

`func (o *CommodityCoverage) GetSparseHistoryOk() (*bool, bool)`

GetSparseHistoryOk returns a tuple with the SparseHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparseHistory

`func (o *CommodityCoverage) SetSparseHistory(v bool)`

SetSparseHistory sets SparseHistory field to given value.


### GetCoverageBasis

`func (o *CommodityCoverage) GetCoverageBasis() string`

GetCoverageBasis returns the CoverageBasis field if non-nil, zero value otherwise.

### GetCoverageBasisOk

`func (o *CommodityCoverage) GetCoverageBasisOk() (*string, bool)`

GetCoverageBasisOk returns a tuple with the CoverageBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageBasis

`func (o *CommodityCoverage) SetCoverageBasis(v string)`

SetCoverageBasis sets CoverageBasis field to given value.


### GetContinuousCoverageVerified

`func (o *CommodityCoverage) GetContinuousCoverageVerified() bool`

GetContinuousCoverageVerified returns the ContinuousCoverageVerified field if non-nil, zero value otherwise.

### GetContinuousCoverageVerifiedOk

`func (o *CommodityCoverage) GetContinuousCoverageVerifiedOk() (*bool, bool)`

GetContinuousCoverageVerifiedOk returns a tuple with the ContinuousCoverageVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousCoverageVerified

`func (o *CommodityCoverage) SetContinuousCoverageVerified(v bool)`

SetContinuousCoverageVerified sets ContinuousCoverageVerified field to given value.

### HasContinuousCoverageVerified

`func (o *CommodityCoverage) HasContinuousCoverageVerified() bool`

HasContinuousCoverageVerified returns a boolean if a field has been set.

### GetReferenceOnly

`func (o *CommodityCoverage) GetReferenceOnly() bool`

GetReferenceOnly returns the ReferenceOnly field if non-nil, zero value otherwise.

### GetReferenceOnlyOk

`func (o *CommodityCoverage) GetReferenceOnlyOk() (*bool, bool)`

GetReferenceOnlyOk returns a tuple with the ReferenceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOnly

`func (o *CommodityCoverage) SetReferenceOnly(v bool)`

SetReferenceOnly sets ReferenceOnly field to given value.

### HasReferenceOnly

`func (o *CommodityCoverage) HasReferenceOnly() bool`

HasReferenceOnly returns a boolean if a field has been set.

### GetExecutableQuotes

`func (o *CommodityCoverage) GetExecutableQuotes() bool`

GetExecutableQuotes returns the ExecutableQuotes field if non-nil, zero value otherwise.

### GetExecutableQuotesOk

`func (o *CommodityCoverage) GetExecutableQuotesOk() (*bool, bool)`

GetExecutableQuotesOk returns a tuple with the ExecutableQuotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutableQuotes

`func (o *CommodityCoverage) SetExecutableQuotes(v bool)`

SetExecutableQuotes sets ExecutableQuotes field to given value.

### HasExecutableQuotes

`func (o *CommodityCoverage) HasExecutableQuotes() bool`

HasExecutableQuotes returns a boolean if a field has been set.

### GetProvenanceVersion

`func (o *CommodityCoverage) GetProvenanceVersion() int64`

GetProvenanceVersion returns the ProvenanceVersion field if non-nil, zero value otherwise.

### GetProvenanceVersionOk

`func (o *CommodityCoverage) GetProvenanceVersionOk() (*int64, bool)`

GetProvenanceVersionOk returns a tuple with the ProvenanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenanceVersion

`func (o *CommodityCoverage) SetProvenanceVersion(v int64)`

SetProvenanceVersion sets ProvenanceVersion field to given value.


### GetRepairedExcludedWeekendRowCount

`func (o *CommodityCoverage) GetRepairedExcludedWeekendRowCount() int64`

GetRepairedExcludedWeekendRowCount returns the RepairedExcludedWeekendRowCount field if non-nil, zero value otherwise.

### GetRepairedExcludedWeekendRowCountOk

`func (o *CommodityCoverage) GetRepairedExcludedWeekendRowCountOk() (*int64, bool)`

GetRepairedExcludedWeekendRowCountOk returns a tuple with the RepairedExcludedWeekendRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairedExcludedWeekendRowCount

`func (o *CommodityCoverage) SetRepairedExcludedWeekendRowCount(v int64)`

SetRepairedExcludedWeekendRowCount sets RepairedExcludedWeekendRowCount field to given value.

### HasRepairedExcludedWeekendRowCount

`func (o *CommodityCoverage) HasRepairedExcludedWeekendRowCount() bool`

HasRepairedExcludedWeekendRowCount returns a boolean if a field has been set.

### GetRepairedInvalidDateRowCount

`func (o *CommodityCoverage) GetRepairedInvalidDateRowCount() int64`

GetRepairedInvalidDateRowCount returns the RepairedInvalidDateRowCount field if non-nil, zero value otherwise.

### GetRepairedInvalidDateRowCountOk

`func (o *CommodityCoverage) GetRepairedInvalidDateRowCountOk() (*int64, bool)`

GetRepairedInvalidDateRowCountOk returns a tuple with the RepairedInvalidDateRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairedInvalidDateRowCount

`func (o *CommodityCoverage) SetRepairedInvalidDateRowCount(v int64)`

SetRepairedInvalidDateRowCount sets RepairedInvalidDateRowCount field to given value.

### HasRepairedInvalidDateRowCount

`func (o *CommodityCoverage) HasRepairedInvalidDateRowCount() bool`

HasRepairedInvalidDateRowCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


