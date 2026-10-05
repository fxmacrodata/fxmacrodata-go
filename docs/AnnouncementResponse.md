# AnnouncementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**Indicator** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**ValueName** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceUrl** | Pointer to **NullableString** |  | [optional] 
**SourceSeriesId** | Pointer to **NullableString** |  | [optional] 
**SourceSeriesName** | Pointer to **NullableString** |  | [optional] 
**SourceLocalName** | Pointer to **NullableString** |  | [optional] 
**SeasonalAdjustment** | Pointer to **NullableString** |  | [optional] 
**PriceBasis** | Pointer to **NullableString** |  | [optional] 
**IsProxy** | Pointer to **bool** |  | [optional] [default to false]
**ProxyNote** | Pointer to **NullableString** |  | [optional] 
**Provenance** | Pointer to [**AnnouncementProvenance**](AnnouncementProvenance.md) |  | [optional] 
**PolicyRole** | Pointer to **NullableString** |  | [optional] 
**PolicyStructure** | Pointer to **NullableString** |  | [optional] 
**ComparisonCompatible** | Pointer to **NullableBool** |  | [optional] 
**PolicyFamily** | Pointer to [**[]PolicyFamilyEntry**](PolicyFamilyEntry.md) |  | [optional] 
**HasOfficialForecast** | **bool** |  | 
**RequestedStartDate** | Pointer to **NullableString** |  | [optional] 
**RequestedEndDate** | Pointer to **NullableString** |  | [optional] 
**RequestedWindowHasData** | Pointer to **NullableBool** |  | [optional] 
**PageIncludesLatestAvailable** | Pointer to **NullableBool** |  | [optional] 
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**EarliestAvailableDate** | Pointer to **NullableString** |  | [optional] 
**LatestAvailableDate** | Pointer to **NullableString** |  | [optional] 
**CbTarget** | Pointer to [**NullableCBTargetInfo**](CBTargetInfo.md) |  | [optional] 
**Remap** | Pointer to [**AnnouncementRemap**](AnnouncementRemap.md) |  | [optional] 
**Filters** | Pointer to [**SeriesFilters**](SeriesFilters.md) |  | [optional] 
**SelectedSeriesId** | Pointer to **NullableString** |  | [optional] 
**SelectedSeries** | Pointer to [**SelectedSeries**](SelectedSeries.md) |  | [optional] 
**SupportedOptions** | Pointer to **map[string][]string** |  | [optional] 
**DataQuality** | [**DataQuality**](DataQuality.md) |  | 
**Replay** | Pointer to **map[string]interface{}** |  | [optional] 
**DatasetVersion** | Pointer to **NullableString** |  | [optional] 
**ValueMode** | Pointer to **string** |  | [optional] [default to "normalized"]
**ValueMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Pagination** | [**IndicatorPaginationInfo**](IndicatorPaginationInfo.md) |  | 
**FreemiumWindow** | Pointer to [**NullableFreemiumWindow**](FreemiumWindow.md) |  | [optional] 
**FreemiumDelay** | Pointer to [**NullableFreemiumDelayNotice**](FreemiumDelayNotice.md) |  | [optional] 
**NoData** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | [**[]AnnouncementDataPoint**](AnnouncementDataPoint.md) |  | 

## Methods

### NewAnnouncementResponse

`func NewAnnouncementResponse(currency string, indicator string, hasOfficialForecast bool, startDate string, endDate string, dataQuality DataQuality, pagination IndicatorPaginationInfo, data []AnnouncementDataPoint, ) *AnnouncementResponse`

NewAnnouncementResponse instantiates a new AnnouncementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementResponseWithDefaults

`func NewAnnouncementResponseWithDefaults() *AnnouncementResponse`

NewAnnouncementResponseWithDefaults instantiates a new AnnouncementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *AnnouncementResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AnnouncementResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AnnouncementResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetIndicator

`func (o *AnnouncementResponse) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *AnnouncementResponse) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *AnnouncementResponse) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.


### GetName

`func (o *AnnouncementResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnnouncementResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnnouncementResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AnnouncementResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AnnouncementResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AnnouncementResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValueName

`func (o *AnnouncementResponse) GetValueName() string`

GetValueName returns the ValueName field if non-nil, zero value otherwise.

### GetValueNameOk

`func (o *AnnouncementResponse) GetValueNameOk() (*string, bool)`

GetValueNameOk returns a tuple with the ValueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueName

`func (o *AnnouncementResponse) SetValueName(v string)`

SetValueName sets ValueName field to given value.

### HasValueName

`func (o *AnnouncementResponse) HasValueName() bool`

HasValueName returns a boolean if a field has been set.

### SetValueNameNil

`func (o *AnnouncementResponse) SetValueNameNil(b bool)`

 SetValueNameNil sets the value for ValueName to be an explicit nil

### UnsetValueName
`func (o *AnnouncementResponse) UnsetValueName()`

UnsetValueName ensures that no value is present for ValueName, not even an explicit nil
### GetSource

`func (o *AnnouncementResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AnnouncementResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AnnouncementResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AnnouncementResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *AnnouncementResponse) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AnnouncementResponse) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceUrl

`func (o *AnnouncementResponse) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *AnnouncementResponse) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *AnnouncementResponse) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *AnnouncementResponse) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *AnnouncementResponse) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *AnnouncementResponse) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSourceSeriesId

`func (o *AnnouncementResponse) GetSourceSeriesId() string`

GetSourceSeriesId returns the SourceSeriesId field if non-nil, zero value otherwise.

### GetSourceSeriesIdOk

`func (o *AnnouncementResponse) GetSourceSeriesIdOk() (*string, bool)`

GetSourceSeriesIdOk returns a tuple with the SourceSeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSeriesId

`func (o *AnnouncementResponse) SetSourceSeriesId(v string)`

SetSourceSeriesId sets SourceSeriesId field to given value.

### HasSourceSeriesId

`func (o *AnnouncementResponse) HasSourceSeriesId() bool`

HasSourceSeriesId returns a boolean if a field has been set.

### SetSourceSeriesIdNil

`func (o *AnnouncementResponse) SetSourceSeriesIdNil(b bool)`

 SetSourceSeriesIdNil sets the value for SourceSeriesId to be an explicit nil

### UnsetSourceSeriesId
`func (o *AnnouncementResponse) UnsetSourceSeriesId()`

UnsetSourceSeriesId ensures that no value is present for SourceSeriesId, not even an explicit nil
### GetSourceSeriesName

`func (o *AnnouncementResponse) GetSourceSeriesName() string`

GetSourceSeriesName returns the SourceSeriesName field if non-nil, zero value otherwise.

### GetSourceSeriesNameOk

`func (o *AnnouncementResponse) GetSourceSeriesNameOk() (*string, bool)`

GetSourceSeriesNameOk returns a tuple with the SourceSeriesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSeriesName

`func (o *AnnouncementResponse) SetSourceSeriesName(v string)`

SetSourceSeriesName sets SourceSeriesName field to given value.

### HasSourceSeriesName

`func (o *AnnouncementResponse) HasSourceSeriesName() bool`

HasSourceSeriesName returns a boolean if a field has been set.

### SetSourceSeriesNameNil

`func (o *AnnouncementResponse) SetSourceSeriesNameNil(b bool)`

 SetSourceSeriesNameNil sets the value for SourceSeriesName to be an explicit nil

### UnsetSourceSeriesName
`func (o *AnnouncementResponse) UnsetSourceSeriesName()`

UnsetSourceSeriesName ensures that no value is present for SourceSeriesName, not even an explicit nil
### GetSourceLocalName

`func (o *AnnouncementResponse) GetSourceLocalName() string`

GetSourceLocalName returns the SourceLocalName field if non-nil, zero value otherwise.

### GetSourceLocalNameOk

`func (o *AnnouncementResponse) GetSourceLocalNameOk() (*string, bool)`

GetSourceLocalNameOk returns a tuple with the SourceLocalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLocalName

`func (o *AnnouncementResponse) SetSourceLocalName(v string)`

SetSourceLocalName sets SourceLocalName field to given value.

### HasSourceLocalName

`func (o *AnnouncementResponse) HasSourceLocalName() bool`

HasSourceLocalName returns a boolean if a field has been set.

### SetSourceLocalNameNil

`func (o *AnnouncementResponse) SetSourceLocalNameNil(b bool)`

 SetSourceLocalNameNil sets the value for SourceLocalName to be an explicit nil

### UnsetSourceLocalName
`func (o *AnnouncementResponse) UnsetSourceLocalName()`

UnsetSourceLocalName ensures that no value is present for SourceLocalName, not even an explicit nil
### GetSeasonalAdjustment

`func (o *AnnouncementResponse) GetSeasonalAdjustment() string`

GetSeasonalAdjustment returns the SeasonalAdjustment field if non-nil, zero value otherwise.

### GetSeasonalAdjustmentOk

`func (o *AnnouncementResponse) GetSeasonalAdjustmentOk() (*string, bool)`

GetSeasonalAdjustmentOk returns a tuple with the SeasonalAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonalAdjustment

`func (o *AnnouncementResponse) SetSeasonalAdjustment(v string)`

SetSeasonalAdjustment sets SeasonalAdjustment field to given value.

### HasSeasonalAdjustment

`func (o *AnnouncementResponse) HasSeasonalAdjustment() bool`

HasSeasonalAdjustment returns a boolean if a field has been set.

### SetSeasonalAdjustmentNil

`func (o *AnnouncementResponse) SetSeasonalAdjustmentNil(b bool)`

 SetSeasonalAdjustmentNil sets the value for SeasonalAdjustment to be an explicit nil

### UnsetSeasonalAdjustment
`func (o *AnnouncementResponse) UnsetSeasonalAdjustment()`

UnsetSeasonalAdjustment ensures that no value is present for SeasonalAdjustment, not even an explicit nil
### GetPriceBasis

`func (o *AnnouncementResponse) GetPriceBasis() string`

GetPriceBasis returns the PriceBasis field if non-nil, zero value otherwise.

### GetPriceBasisOk

`func (o *AnnouncementResponse) GetPriceBasisOk() (*string, bool)`

GetPriceBasisOk returns a tuple with the PriceBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceBasis

`func (o *AnnouncementResponse) SetPriceBasis(v string)`

SetPriceBasis sets PriceBasis field to given value.

### HasPriceBasis

`func (o *AnnouncementResponse) HasPriceBasis() bool`

HasPriceBasis returns a boolean if a field has been set.

### SetPriceBasisNil

`func (o *AnnouncementResponse) SetPriceBasisNil(b bool)`

 SetPriceBasisNil sets the value for PriceBasis to be an explicit nil

### UnsetPriceBasis
`func (o *AnnouncementResponse) UnsetPriceBasis()`

UnsetPriceBasis ensures that no value is present for PriceBasis, not even an explicit nil
### GetIsProxy

`func (o *AnnouncementResponse) GetIsProxy() bool`

GetIsProxy returns the IsProxy field if non-nil, zero value otherwise.

### GetIsProxyOk

`func (o *AnnouncementResponse) GetIsProxyOk() (*bool, bool)`

GetIsProxyOk returns a tuple with the IsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxy

`func (o *AnnouncementResponse) SetIsProxy(v bool)`

SetIsProxy sets IsProxy field to given value.

### HasIsProxy

`func (o *AnnouncementResponse) HasIsProxy() bool`

HasIsProxy returns a boolean if a field has been set.

### GetProxyNote

`func (o *AnnouncementResponse) GetProxyNote() string`

GetProxyNote returns the ProxyNote field if non-nil, zero value otherwise.

### GetProxyNoteOk

`func (o *AnnouncementResponse) GetProxyNoteOk() (*string, bool)`

GetProxyNoteOk returns a tuple with the ProxyNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyNote

`func (o *AnnouncementResponse) SetProxyNote(v string)`

SetProxyNote sets ProxyNote field to given value.

### HasProxyNote

`func (o *AnnouncementResponse) HasProxyNote() bool`

HasProxyNote returns a boolean if a field has been set.

### SetProxyNoteNil

`func (o *AnnouncementResponse) SetProxyNoteNil(b bool)`

 SetProxyNoteNil sets the value for ProxyNote to be an explicit nil

### UnsetProxyNote
`func (o *AnnouncementResponse) UnsetProxyNote()`

UnsetProxyNote ensures that no value is present for ProxyNote, not even an explicit nil
### GetProvenance

`func (o *AnnouncementResponse) GetProvenance() AnnouncementProvenance`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *AnnouncementResponse) GetProvenanceOk() (*AnnouncementProvenance, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *AnnouncementResponse) SetProvenance(v AnnouncementProvenance)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *AnnouncementResponse) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetPolicyRole

`func (o *AnnouncementResponse) GetPolicyRole() string`

GetPolicyRole returns the PolicyRole field if non-nil, zero value otherwise.

### GetPolicyRoleOk

`func (o *AnnouncementResponse) GetPolicyRoleOk() (*string, bool)`

GetPolicyRoleOk returns a tuple with the PolicyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRole

`func (o *AnnouncementResponse) SetPolicyRole(v string)`

SetPolicyRole sets PolicyRole field to given value.

### HasPolicyRole

`func (o *AnnouncementResponse) HasPolicyRole() bool`

HasPolicyRole returns a boolean if a field has been set.

### SetPolicyRoleNil

`func (o *AnnouncementResponse) SetPolicyRoleNil(b bool)`

 SetPolicyRoleNil sets the value for PolicyRole to be an explicit nil

### UnsetPolicyRole
`func (o *AnnouncementResponse) UnsetPolicyRole()`

UnsetPolicyRole ensures that no value is present for PolicyRole, not even an explicit nil
### GetPolicyStructure

`func (o *AnnouncementResponse) GetPolicyStructure() string`

GetPolicyStructure returns the PolicyStructure field if non-nil, zero value otherwise.

### GetPolicyStructureOk

`func (o *AnnouncementResponse) GetPolicyStructureOk() (*string, bool)`

GetPolicyStructureOk returns a tuple with the PolicyStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyStructure

`func (o *AnnouncementResponse) SetPolicyStructure(v string)`

SetPolicyStructure sets PolicyStructure field to given value.

### HasPolicyStructure

`func (o *AnnouncementResponse) HasPolicyStructure() bool`

HasPolicyStructure returns a boolean if a field has been set.

### SetPolicyStructureNil

`func (o *AnnouncementResponse) SetPolicyStructureNil(b bool)`

 SetPolicyStructureNil sets the value for PolicyStructure to be an explicit nil

### UnsetPolicyStructure
`func (o *AnnouncementResponse) UnsetPolicyStructure()`

UnsetPolicyStructure ensures that no value is present for PolicyStructure, not even an explicit nil
### GetComparisonCompatible

`func (o *AnnouncementResponse) GetComparisonCompatible() bool`

GetComparisonCompatible returns the ComparisonCompatible field if non-nil, zero value otherwise.

### GetComparisonCompatibleOk

`func (o *AnnouncementResponse) GetComparisonCompatibleOk() (*bool, bool)`

GetComparisonCompatibleOk returns a tuple with the ComparisonCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonCompatible

`func (o *AnnouncementResponse) SetComparisonCompatible(v bool)`

SetComparisonCompatible sets ComparisonCompatible field to given value.

### HasComparisonCompatible

`func (o *AnnouncementResponse) HasComparisonCompatible() bool`

HasComparisonCompatible returns a boolean if a field has been set.

### SetComparisonCompatibleNil

`func (o *AnnouncementResponse) SetComparisonCompatibleNil(b bool)`

 SetComparisonCompatibleNil sets the value for ComparisonCompatible to be an explicit nil

### UnsetComparisonCompatible
`func (o *AnnouncementResponse) UnsetComparisonCompatible()`

UnsetComparisonCompatible ensures that no value is present for ComparisonCompatible, not even an explicit nil
### GetPolicyFamily

`func (o *AnnouncementResponse) GetPolicyFamily() []PolicyFamilyEntry`

GetPolicyFamily returns the PolicyFamily field if non-nil, zero value otherwise.

### GetPolicyFamilyOk

`func (o *AnnouncementResponse) GetPolicyFamilyOk() (*[]PolicyFamilyEntry, bool)`

GetPolicyFamilyOk returns a tuple with the PolicyFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyFamily

`func (o *AnnouncementResponse) SetPolicyFamily(v []PolicyFamilyEntry)`

SetPolicyFamily sets PolicyFamily field to given value.

### HasPolicyFamily

`func (o *AnnouncementResponse) HasPolicyFamily() bool`

HasPolicyFamily returns a boolean if a field has been set.

### SetPolicyFamilyNil

`func (o *AnnouncementResponse) SetPolicyFamilyNil(b bool)`

 SetPolicyFamilyNil sets the value for PolicyFamily to be an explicit nil

### UnsetPolicyFamily
`func (o *AnnouncementResponse) UnsetPolicyFamily()`

UnsetPolicyFamily ensures that no value is present for PolicyFamily, not even an explicit nil
### GetHasOfficialForecast

`func (o *AnnouncementResponse) GetHasOfficialForecast() bool`

GetHasOfficialForecast returns the HasOfficialForecast field if non-nil, zero value otherwise.

### GetHasOfficialForecastOk

`func (o *AnnouncementResponse) GetHasOfficialForecastOk() (*bool, bool)`

GetHasOfficialForecastOk returns a tuple with the HasOfficialForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOfficialForecast

`func (o *AnnouncementResponse) SetHasOfficialForecast(v bool)`

SetHasOfficialForecast sets HasOfficialForecast field to given value.


### GetRequestedStartDate

`func (o *AnnouncementResponse) GetRequestedStartDate() string`

GetRequestedStartDate returns the RequestedStartDate field if non-nil, zero value otherwise.

### GetRequestedStartDateOk

`func (o *AnnouncementResponse) GetRequestedStartDateOk() (*string, bool)`

GetRequestedStartDateOk returns a tuple with the RequestedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedStartDate

`func (o *AnnouncementResponse) SetRequestedStartDate(v string)`

SetRequestedStartDate sets RequestedStartDate field to given value.

### HasRequestedStartDate

`func (o *AnnouncementResponse) HasRequestedStartDate() bool`

HasRequestedStartDate returns a boolean if a field has been set.

### SetRequestedStartDateNil

`func (o *AnnouncementResponse) SetRequestedStartDateNil(b bool)`

 SetRequestedStartDateNil sets the value for RequestedStartDate to be an explicit nil

### UnsetRequestedStartDate
`func (o *AnnouncementResponse) UnsetRequestedStartDate()`

UnsetRequestedStartDate ensures that no value is present for RequestedStartDate, not even an explicit nil
### GetRequestedEndDate

`func (o *AnnouncementResponse) GetRequestedEndDate() string`

GetRequestedEndDate returns the RequestedEndDate field if non-nil, zero value otherwise.

### GetRequestedEndDateOk

`func (o *AnnouncementResponse) GetRequestedEndDateOk() (*string, bool)`

GetRequestedEndDateOk returns a tuple with the RequestedEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedEndDate

`func (o *AnnouncementResponse) SetRequestedEndDate(v string)`

SetRequestedEndDate sets RequestedEndDate field to given value.

### HasRequestedEndDate

`func (o *AnnouncementResponse) HasRequestedEndDate() bool`

HasRequestedEndDate returns a boolean if a field has been set.

### SetRequestedEndDateNil

`func (o *AnnouncementResponse) SetRequestedEndDateNil(b bool)`

 SetRequestedEndDateNil sets the value for RequestedEndDate to be an explicit nil

### UnsetRequestedEndDate
`func (o *AnnouncementResponse) UnsetRequestedEndDate()`

UnsetRequestedEndDate ensures that no value is present for RequestedEndDate, not even an explicit nil
### GetRequestedWindowHasData

`func (o *AnnouncementResponse) GetRequestedWindowHasData() bool`

GetRequestedWindowHasData returns the RequestedWindowHasData field if non-nil, zero value otherwise.

### GetRequestedWindowHasDataOk

`func (o *AnnouncementResponse) GetRequestedWindowHasDataOk() (*bool, bool)`

GetRequestedWindowHasDataOk returns a tuple with the RequestedWindowHasData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedWindowHasData

`func (o *AnnouncementResponse) SetRequestedWindowHasData(v bool)`

SetRequestedWindowHasData sets RequestedWindowHasData field to given value.

### HasRequestedWindowHasData

`func (o *AnnouncementResponse) HasRequestedWindowHasData() bool`

HasRequestedWindowHasData returns a boolean if a field has been set.

### SetRequestedWindowHasDataNil

`func (o *AnnouncementResponse) SetRequestedWindowHasDataNil(b bool)`

 SetRequestedWindowHasDataNil sets the value for RequestedWindowHasData to be an explicit nil

### UnsetRequestedWindowHasData
`func (o *AnnouncementResponse) UnsetRequestedWindowHasData()`

UnsetRequestedWindowHasData ensures that no value is present for RequestedWindowHasData, not even an explicit nil
### GetPageIncludesLatestAvailable

`func (o *AnnouncementResponse) GetPageIncludesLatestAvailable() bool`

GetPageIncludesLatestAvailable returns the PageIncludesLatestAvailable field if non-nil, zero value otherwise.

### GetPageIncludesLatestAvailableOk

`func (o *AnnouncementResponse) GetPageIncludesLatestAvailableOk() (*bool, bool)`

GetPageIncludesLatestAvailableOk returns a tuple with the PageIncludesLatestAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIncludesLatestAvailable

`func (o *AnnouncementResponse) SetPageIncludesLatestAvailable(v bool)`

SetPageIncludesLatestAvailable sets PageIncludesLatestAvailable field to given value.

### HasPageIncludesLatestAvailable

`func (o *AnnouncementResponse) HasPageIncludesLatestAvailable() bool`

HasPageIncludesLatestAvailable returns a boolean if a field has been set.

### SetPageIncludesLatestAvailableNil

`func (o *AnnouncementResponse) SetPageIncludesLatestAvailableNil(b bool)`

 SetPageIncludesLatestAvailableNil sets the value for PageIncludesLatestAvailable to be an explicit nil

### UnsetPageIncludesLatestAvailable
`func (o *AnnouncementResponse) UnsetPageIncludesLatestAvailable()`

UnsetPageIncludesLatestAvailable ensures that no value is present for PageIncludesLatestAvailable, not even an explicit nil
### GetStartDate

`func (o *AnnouncementResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AnnouncementResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AnnouncementResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *AnnouncementResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AnnouncementResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AnnouncementResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetEarliestAvailableDate

`func (o *AnnouncementResponse) GetEarliestAvailableDate() string`

GetEarliestAvailableDate returns the EarliestAvailableDate field if non-nil, zero value otherwise.

### GetEarliestAvailableDateOk

`func (o *AnnouncementResponse) GetEarliestAvailableDateOk() (*string, bool)`

GetEarliestAvailableDateOk returns a tuple with the EarliestAvailableDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestAvailableDate

`func (o *AnnouncementResponse) SetEarliestAvailableDate(v string)`

SetEarliestAvailableDate sets EarliestAvailableDate field to given value.

### HasEarliestAvailableDate

`func (o *AnnouncementResponse) HasEarliestAvailableDate() bool`

HasEarliestAvailableDate returns a boolean if a field has been set.

### SetEarliestAvailableDateNil

`func (o *AnnouncementResponse) SetEarliestAvailableDateNil(b bool)`

 SetEarliestAvailableDateNil sets the value for EarliestAvailableDate to be an explicit nil

### UnsetEarliestAvailableDate
`func (o *AnnouncementResponse) UnsetEarliestAvailableDate()`

UnsetEarliestAvailableDate ensures that no value is present for EarliestAvailableDate, not even an explicit nil
### GetLatestAvailableDate

`func (o *AnnouncementResponse) GetLatestAvailableDate() string`

GetLatestAvailableDate returns the LatestAvailableDate field if non-nil, zero value otherwise.

### GetLatestAvailableDateOk

`func (o *AnnouncementResponse) GetLatestAvailableDateOk() (*string, bool)`

GetLatestAvailableDateOk returns a tuple with the LatestAvailableDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAvailableDate

`func (o *AnnouncementResponse) SetLatestAvailableDate(v string)`

SetLatestAvailableDate sets LatestAvailableDate field to given value.

### HasLatestAvailableDate

`func (o *AnnouncementResponse) HasLatestAvailableDate() bool`

HasLatestAvailableDate returns a boolean if a field has been set.

### SetLatestAvailableDateNil

`func (o *AnnouncementResponse) SetLatestAvailableDateNil(b bool)`

 SetLatestAvailableDateNil sets the value for LatestAvailableDate to be an explicit nil

### UnsetLatestAvailableDate
`func (o *AnnouncementResponse) UnsetLatestAvailableDate()`

UnsetLatestAvailableDate ensures that no value is present for LatestAvailableDate, not even an explicit nil
### GetCbTarget

`func (o *AnnouncementResponse) GetCbTarget() CBTargetInfo`

GetCbTarget returns the CbTarget field if non-nil, zero value otherwise.

### GetCbTargetOk

`func (o *AnnouncementResponse) GetCbTargetOk() (*CBTargetInfo, bool)`

GetCbTargetOk returns a tuple with the CbTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbTarget

`func (o *AnnouncementResponse) SetCbTarget(v CBTargetInfo)`

SetCbTarget sets CbTarget field to given value.

### HasCbTarget

`func (o *AnnouncementResponse) HasCbTarget() bool`

HasCbTarget returns a boolean if a field has been set.

### SetCbTargetNil

`func (o *AnnouncementResponse) SetCbTargetNil(b bool)`

 SetCbTargetNil sets the value for CbTarget to be an explicit nil

### UnsetCbTarget
`func (o *AnnouncementResponse) UnsetCbTarget()`

UnsetCbTarget ensures that no value is present for CbTarget, not even an explicit nil
### GetRemap

`func (o *AnnouncementResponse) GetRemap() AnnouncementRemap`

GetRemap returns the Remap field if non-nil, zero value otherwise.

### GetRemapOk

`func (o *AnnouncementResponse) GetRemapOk() (*AnnouncementRemap, bool)`

GetRemapOk returns a tuple with the Remap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemap

`func (o *AnnouncementResponse) SetRemap(v AnnouncementRemap)`

SetRemap sets Remap field to given value.

### HasRemap

`func (o *AnnouncementResponse) HasRemap() bool`

HasRemap returns a boolean if a field has been set.

### GetFilters

`func (o *AnnouncementResponse) GetFilters() SeriesFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AnnouncementResponse) GetFiltersOk() (*SeriesFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AnnouncementResponse) SetFilters(v SeriesFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AnnouncementResponse) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSelectedSeriesId

`func (o *AnnouncementResponse) GetSelectedSeriesId() string`

GetSelectedSeriesId returns the SelectedSeriesId field if non-nil, zero value otherwise.

### GetSelectedSeriesIdOk

`func (o *AnnouncementResponse) GetSelectedSeriesIdOk() (*string, bool)`

GetSelectedSeriesIdOk returns a tuple with the SelectedSeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSeriesId

`func (o *AnnouncementResponse) SetSelectedSeriesId(v string)`

SetSelectedSeriesId sets SelectedSeriesId field to given value.

### HasSelectedSeriesId

`func (o *AnnouncementResponse) HasSelectedSeriesId() bool`

HasSelectedSeriesId returns a boolean if a field has been set.

### SetSelectedSeriesIdNil

`func (o *AnnouncementResponse) SetSelectedSeriesIdNil(b bool)`

 SetSelectedSeriesIdNil sets the value for SelectedSeriesId to be an explicit nil

### UnsetSelectedSeriesId
`func (o *AnnouncementResponse) UnsetSelectedSeriesId()`

UnsetSelectedSeriesId ensures that no value is present for SelectedSeriesId, not even an explicit nil
### GetSelectedSeries

`func (o *AnnouncementResponse) GetSelectedSeries() SelectedSeries`

GetSelectedSeries returns the SelectedSeries field if non-nil, zero value otherwise.

### GetSelectedSeriesOk

`func (o *AnnouncementResponse) GetSelectedSeriesOk() (*SelectedSeries, bool)`

GetSelectedSeriesOk returns a tuple with the SelectedSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSeries

`func (o *AnnouncementResponse) SetSelectedSeries(v SelectedSeries)`

SetSelectedSeries sets SelectedSeries field to given value.

### HasSelectedSeries

`func (o *AnnouncementResponse) HasSelectedSeries() bool`

HasSelectedSeries returns a boolean if a field has been set.

### GetSupportedOptions

`func (o *AnnouncementResponse) GetSupportedOptions() map[string][]string`

GetSupportedOptions returns the SupportedOptions field if non-nil, zero value otherwise.

### GetSupportedOptionsOk

`func (o *AnnouncementResponse) GetSupportedOptionsOk() (*map[string][]string, bool)`

GetSupportedOptionsOk returns a tuple with the SupportedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOptions

`func (o *AnnouncementResponse) SetSupportedOptions(v map[string][]string)`

SetSupportedOptions sets SupportedOptions field to given value.

### HasSupportedOptions

`func (o *AnnouncementResponse) HasSupportedOptions() bool`

HasSupportedOptions returns a boolean if a field has been set.

### SetSupportedOptionsNil

`func (o *AnnouncementResponse) SetSupportedOptionsNil(b bool)`

 SetSupportedOptionsNil sets the value for SupportedOptions to be an explicit nil

### UnsetSupportedOptions
`func (o *AnnouncementResponse) UnsetSupportedOptions()`

UnsetSupportedOptions ensures that no value is present for SupportedOptions, not even an explicit nil
### GetDataQuality

`func (o *AnnouncementResponse) GetDataQuality() DataQuality`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *AnnouncementResponse) GetDataQualityOk() (*DataQuality, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *AnnouncementResponse) SetDataQuality(v DataQuality)`

SetDataQuality sets DataQuality field to given value.


### GetReplay

`func (o *AnnouncementResponse) GetReplay() map[string]interface{}`

GetReplay returns the Replay field if non-nil, zero value otherwise.

### GetReplayOk

`func (o *AnnouncementResponse) GetReplayOk() (*map[string]interface{}, bool)`

GetReplayOk returns a tuple with the Replay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplay

`func (o *AnnouncementResponse) SetReplay(v map[string]interface{})`

SetReplay sets Replay field to given value.

### HasReplay

`func (o *AnnouncementResponse) HasReplay() bool`

HasReplay returns a boolean if a field has been set.

### SetReplayNil

`func (o *AnnouncementResponse) SetReplayNil(b bool)`

 SetReplayNil sets the value for Replay to be an explicit nil

### UnsetReplay
`func (o *AnnouncementResponse) UnsetReplay()`

UnsetReplay ensures that no value is present for Replay, not even an explicit nil
### GetDatasetVersion

`func (o *AnnouncementResponse) GetDatasetVersion() string`

GetDatasetVersion returns the DatasetVersion field if non-nil, zero value otherwise.

### GetDatasetVersionOk

`func (o *AnnouncementResponse) GetDatasetVersionOk() (*string, bool)`

GetDatasetVersionOk returns a tuple with the DatasetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetVersion

`func (o *AnnouncementResponse) SetDatasetVersion(v string)`

SetDatasetVersion sets DatasetVersion field to given value.

### HasDatasetVersion

`func (o *AnnouncementResponse) HasDatasetVersion() bool`

HasDatasetVersion returns a boolean if a field has been set.

### SetDatasetVersionNil

`func (o *AnnouncementResponse) SetDatasetVersionNil(b bool)`

 SetDatasetVersionNil sets the value for DatasetVersion to be an explicit nil

### UnsetDatasetVersion
`func (o *AnnouncementResponse) UnsetDatasetVersion()`

UnsetDatasetVersion ensures that no value is present for DatasetVersion, not even an explicit nil
### GetValueMode

`func (o *AnnouncementResponse) GetValueMode() string`

GetValueMode returns the ValueMode field if non-nil, zero value otherwise.

### GetValueModeOk

`func (o *AnnouncementResponse) GetValueModeOk() (*string, bool)`

GetValueModeOk returns a tuple with the ValueMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMode

`func (o *AnnouncementResponse) SetValueMode(v string)`

SetValueMode sets ValueMode field to given value.

### HasValueMode

`func (o *AnnouncementResponse) HasValueMode() bool`

HasValueMode returns a boolean if a field has been set.

### GetValueMetadata

`func (o *AnnouncementResponse) GetValueMetadata() map[string]interface{}`

GetValueMetadata returns the ValueMetadata field if non-nil, zero value otherwise.

### GetValueMetadataOk

`func (o *AnnouncementResponse) GetValueMetadataOk() (*map[string]interface{}, bool)`

GetValueMetadataOk returns a tuple with the ValueMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMetadata

`func (o *AnnouncementResponse) SetValueMetadata(v map[string]interface{})`

SetValueMetadata sets ValueMetadata field to given value.

### HasValueMetadata

`func (o *AnnouncementResponse) HasValueMetadata() bool`

HasValueMetadata returns a boolean if a field has been set.

### SetValueMetadataNil

`func (o *AnnouncementResponse) SetValueMetadataNil(b bool)`

 SetValueMetadataNil sets the value for ValueMetadata to be an explicit nil

### UnsetValueMetadata
`func (o *AnnouncementResponse) UnsetValueMetadata()`

UnsetValueMetadata ensures that no value is present for ValueMetadata, not even an explicit nil
### GetPagination

`func (o *AnnouncementResponse) GetPagination() IndicatorPaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AnnouncementResponse) GetPaginationOk() (*IndicatorPaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AnnouncementResponse) SetPagination(v IndicatorPaginationInfo)`

SetPagination sets Pagination field to given value.


### GetFreemiumWindow

`func (o *AnnouncementResponse) GetFreemiumWindow() FreemiumWindow`

GetFreemiumWindow returns the FreemiumWindow field if non-nil, zero value otherwise.

### GetFreemiumWindowOk

`func (o *AnnouncementResponse) GetFreemiumWindowOk() (*FreemiumWindow, bool)`

GetFreemiumWindowOk returns a tuple with the FreemiumWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreemiumWindow

`func (o *AnnouncementResponse) SetFreemiumWindow(v FreemiumWindow)`

SetFreemiumWindow sets FreemiumWindow field to given value.

### HasFreemiumWindow

`func (o *AnnouncementResponse) HasFreemiumWindow() bool`

HasFreemiumWindow returns a boolean if a field has been set.

### SetFreemiumWindowNil

`func (o *AnnouncementResponse) SetFreemiumWindowNil(b bool)`

 SetFreemiumWindowNil sets the value for FreemiumWindow to be an explicit nil

### UnsetFreemiumWindow
`func (o *AnnouncementResponse) UnsetFreemiumWindow()`

UnsetFreemiumWindow ensures that no value is present for FreemiumWindow, not even an explicit nil
### GetFreemiumDelay

`func (o *AnnouncementResponse) GetFreemiumDelay() FreemiumDelayNotice`

GetFreemiumDelay returns the FreemiumDelay field if non-nil, zero value otherwise.

### GetFreemiumDelayOk

`func (o *AnnouncementResponse) GetFreemiumDelayOk() (*FreemiumDelayNotice, bool)`

GetFreemiumDelayOk returns a tuple with the FreemiumDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreemiumDelay

`func (o *AnnouncementResponse) SetFreemiumDelay(v FreemiumDelayNotice)`

SetFreemiumDelay sets FreemiumDelay field to given value.

### HasFreemiumDelay

`func (o *AnnouncementResponse) HasFreemiumDelay() bool`

HasFreemiumDelay returns a boolean if a field has been set.

### SetFreemiumDelayNil

`func (o *AnnouncementResponse) SetFreemiumDelayNil(b bool)`

 SetFreemiumDelayNil sets the value for FreemiumDelay to be an explicit nil

### UnsetFreemiumDelay
`func (o *AnnouncementResponse) UnsetFreemiumDelay()`

UnsetFreemiumDelay ensures that no value is present for FreemiumDelay, not even an explicit nil
### GetNoData

`func (o *AnnouncementResponse) GetNoData() map[string]interface{}`

GetNoData returns the NoData field if non-nil, zero value otherwise.

### GetNoDataOk

`func (o *AnnouncementResponse) GetNoDataOk() (*map[string]interface{}, bool)`

GetNoDataOk returns a tuple with the NoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoData

`func (o *AnnouncementResponse) SetNoData(v map[string]interface{})`

SetNoData sets NoData field to given value.

### HasNoData

`func (o *AnnouncementResponse) HasNoData() bool`

HasNoData returns a boolean if a field has been set.

### SetNoDataNil

`func (o *AnnouncementResponse) SetNoDataNil(b bool)`

 SetNoDataNil sets the value for NoData to be an explicit nil

### UnsetNoData
`func (o *AnnouncementResponse) UnsetNoData()`

UnsetNoData ensures that no value is present for NoData, not even an explicit nil
### GetData

`func (o *AnnouncementResponse) GetData() []AnnouncementDataPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnnouncementResponse) GetDataOk() (*[]AnnouncementDataPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnnouncementResponse) SetData(v []AnnouncementDataPoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


