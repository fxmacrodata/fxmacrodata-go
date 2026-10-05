# FinancialSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Currency** | **string** |  | 
**Publisher** | **string** |  | 
**Url** | **string** |  | 
**LicenseUrl** | Pointer to **NullableString** |  | [optional] 
**Attribution** | **string** |  | 
**Frequency** | Pointer to **NullableString** |  | [optional] 
**PublicationFrequency** | Pointer to **NullableString** |  | [optional] 
**HistoryStart** | Pointer to **NullableString** |  | [optional] 
**Measures** | **[]string** |  | 
**DatasetName** | Pointer to **NullableString** |  | [optional] 
**License** | Pointer to **NullableString** |  | [optional] 
**ReuseConditions** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFinancialSource

`func NewFinancialSource(id string, currency string, publisher string, url string, attribution string, measures []string, ) *FinancialSource`

NewFinancialSource instantiates a new FinancialSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialSourceWithDefaults

`func NewFinancialSourceWithDefaults() *FinancialSource`

NewFinancialSourceWithDefaults instantiates a new FinancialSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FinancialSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialSource) SetId(v string)`

SetId sets Id field to given value.


### GetCurrency

`func (o *FinancialSource) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FinancialSource) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FinancialSource) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPublisher

`func (o *FinancialSource) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *FinancialSource) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *FinancialSource) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.


### GetUrl

`func (o *FinancialSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FinancialSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FinancialSource) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLicenseUrl

`func (o *FinancialSource) GetLicenseUrl() string`

GetLicenseUrl returns the LicenseUrl field if non-nil, zero value otherwise.

### GetLicenseUrlOk

`func (o *FinancialSource) GetLicenseUrlOk() (*string, bool)`

GetLicenseUrlOk returns a tuple with the LicenseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUrl

`func (o *FinancialSource) SetLicenseUrl(v string)`

SetLicenseUrl sets LicenseUrl field to given value.

### HasLicenseUrl

`func (o *FinancialSource) HasLicenseUrl() bool`

HasLicenseUrl returns a boolean if a field has been set.

### SetLicenseUrlNil

`func (o *FinancialSource) SetLicenseUrlNil(b bool)`

 SetLicenseUrlNil sets the value for LicenseUrl to be an explicit nil

### UnsetLicenseUrl
`func (o *FinancialSource) UnsetLicenseUrl()`

UnsetLicenseUrl ensures that no value is present for LicenseUrl, not even an explicit nil
### GetAttribution

`func (o *FinancialSource) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *FinancialSource) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *FinancialSource) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.


### GetFrequency

`func (o *FinancialSource) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *FinancialSource) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *FinancialSource) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *FinancialSource) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *FinancialSource) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *FinancialSource) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetPublicationFrequency

`func (o *FinancialSource) GetPublicationFrequency() string`

GetPublicationFrequency returns the PublicationFrequency field if non-nil, zero value otherwise.

### GetPublicationFrequencyOk

`func (o *FinancialSource) GetPublicationFrequencyOk() (*string, bool)`

GetPublicationFrequencyOk returns a tuple with the PublicationFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationFrequency

`func (o *FinancialSource) SetPublicationFrequency(v string)`

SetPublicationFrequency sets PublicationFrequency field to given value.

### HasPublicationFrequency

`func (o *FinancialSource) HasPublicationFrequency() bool`

HasPublicationFrequency returns a boolean if a field has been set.

### SetPublicationFrequencyNil

`func (o *FinancialSource) SetPublicationFrequencyNil(b bool)`

 SetPublicationFrequencyNil sets the value for PublicationFrequency to be an explicit nil

### UnsetPublicationFrequency
`func (o *FinancialSource) UnsetPublicationFrequency()`

UnsetPublicationFrequency ensures that no value is present for PublicationFrequency, not even an explicit nil
### GetHistoryStart

`func (o *FinancialSource) GetHistoryStart() string`

GetHistoryStart returns the HistoryStart field if non-nil, zero value otherwise.

### GetHistoryStartOk

`func (o *FinancialSource) GetHistoryStartOk() (*string, bool)`

GetHistoryStartOk returns a tuple with the HistoryStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryStart

`func (o *FinancialSource) SetHistoryStart(v string)`

SetHistoryStart sets HistoryStart field to given value.

### HasHistoryStart

`func (o *FinancialSource) HasHistoryStart() bool`

HasHistoryStart returns a boolean if a field has been set.

### SetHistoryStartNil

`func (o *FinancialSource) SetHistoryStartNil(b bool)`

 SetHistoryStartNil sets the value for HistoryStart to be an explicit nil

### UnsetHistoryStart
`func (o *FinancialSource) UnsetHistoryStart()`

UnsetHistoryStart ensures that no value is present for HistoryStart, not even an explicit nil
### GetMeasures

`func (o *FinancialSource) GetMeasures() []string`

GetMeasures returns the Measures field if non-nil, zero value otherwise.

### GetMeasuresOk

`func (o *FinancialSource) GetMeasuresOk() (*[]string, bool)`

GetMeasuresOk returns a tuple with the Measures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasures

`func (o *FinancialSource) SetMeasures(v []string)`

SetMeasures sets Measures field to given value.


### GetDatasetName

`func (o *FinancialSource) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *FinancialSource) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *FinancialSource) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.

### HasDatasetName

`func (o *FinancialSource) HasDatasetName() bool`

HasDatasetName returns a boolean if a field has been set.

### SetDatasetNameNil

`func (o *FinancialSource) SetDatasetNameNil(b bool)`

 SetDatasetNameNil sets the value for DatasetName to be an explicit nil

### UnsetDatasetName
`func (o *FinancialSource) UnsetDatasetName()`

UnsetDatasetName ensures that no value is present for DatasetName, not even an explicit nil
### GetLicense

`func (o *FinancialSource) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *FinancialSource) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *FinancialSource) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *FinancialSource) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### SetLicenseNil

`func (o *FinancialSource) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *FinancialSource) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil
### GetReuseConditions

`func (o *FinancialSource) GetReuseConditions() string`

GetReuseConditions returns the ReuseConditions field if non-nil, zero value otherwise.

### GetReuseConditionsOk

`func (o *FinancialSource) GetReuseConditionsOk() (*string, bool)`

GetReuseConditionsOk returns a tuple with the ReuseConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseConditions

`func (o *FinancialSource) SetReuseConditions(v string)`

SetReuseConditions sets ReuseConditions field to given value.

### HasReuseConditions

`func (o *FinancialSource) HasReuseConditions() bool`

HasReuseConditions returns a boolean if a field has been set.

### SetReuseConditionsNil

`func (o *FinancialSource) SetReuseConditionsNil(b bool)`

 SetReuseConditionsNil sets the value for ReuseConditions to be an explicit nil

### UnsetReuseConditions
`func (o *FinancialSource) UnsetReuseConditions()`

UnsetReuseConditions ensures that no value is present for ReuseConditions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


